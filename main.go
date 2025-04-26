package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"os"
	"os/exec"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/gofiber/fiber/v2"
)

const (
	POOL_LIMIT = 100
)

var (
	wg sync.WaitGroup
)

var (
	manifest	= flag.String("manifest", "manifest.json", "Manifest file location.")
	poolLimit = flag.Int("poollimit", POOL_LIMIT, "Download pool limit.")
)

func downloadModFile(projectID, fileID string, pool chan bool, time uint8) {
	defer wg.Done()
	wg.Add(1)

	agent := fiber.Get(fmt.Sprintf("https://www.curseforge.com/api/v1/mods/%v/files/%v/download", projectID, fileID))
	_, body, err := agent.String()
	if err != nil && time < 3 {
		go downloadModFile(projectID, fileID, pool, time + 1)
		return
	}
	if err != nil && time >= 3 {
		log.Fatalf("Failed to download mod file. %v", err)
	}
	i := strings.Index(body, "https")
	errs := exec.Command("curl", "-LO", body[i:]).Run()
	if err != nil && time < 3 {
		go downloadModFile(projectID, fileID, pool, time + 1)
		return
	}
	if errs != nil && time >= 3 {
		log.Fatalf("Failed to download mod file. %v", errs)
	}
	pool <- true
}

func main() {
	flag.Parse()
	
	// Open file
	file, err := os.Open(*manifest)
	if err != nil {
		log.Fatalf("Failed to open manifest file. %v", err)
	}
	defer file.Close()

	// Get file info
	fileInfo, err := file.Stat()
	if err != nil {
		log.Fatalf("Failed to stat manifest file. %v", err)
	}
	
	// Read data from file and store to a variable
	data := make([]byte, fileInfo.Size())
	_, err = file.Read(data)
	if err != nil {
		log.Fatalf("Failed to read manifest file. %v", err)
	}
	var value map[string]interface{}
	json.Unmarshal(data, &value)
	
	// Download files
	pool := make(chan bool, *poolLimit)
	for range *poolLimit {
		pool <- true
	}
	fmt.Println("Init Pool size:", len(pool))
	for finish, file := range value["files"].([]interface{}) {
		MainLoop:
		for {
			select {
			case <-pool:
				go downloadModFile(strconv.Itoa(int(file.(map[string]interface{})["projectID"].(float64))), strconv.Itoa(int(file.(map[string]interface{})["fileID"].(float64))), pool, 0)
				fmt.Printf("%v/%v: Pool size: %v\n", finish, len(value["files"].([]interface{})), len(pool))
				break MainLoop
			default :
				if finish == len(value["files"].([]interface{})){
					break MainLoop
				}
				time.Sleep(time.Millisecond * 100)
				fmt.Printf("%v/%v: Pool size: %v\n", finish, len(value["files"].([]interface{})), len(pool))
			}
		}
	}
	// Wait until all goroutine finished
	wg.Wait()
	fmt.Printf("%v/%v: Pool size: %v\n", len(value["files"].([]interface{})), len(value["files"].([]interface{})), len(pool))
	fmt.Println("Complete.")
}

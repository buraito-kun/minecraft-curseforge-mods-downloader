# A program to download Minecraft mods from CurseForge

## How to install
### 1. Run the following command in terminal to install
``` bash
go install github.com/buraito-kun/minecraft-curseforge-mods-downloader@latest
```

### 2. (optional) Try the following command to ensure it's working
``` bash
minecraft-curseforge-mods-downloader -help
```
<br>

## How to use
### 1. Head to modpack you want to download in CurseForge website
### 2. Click files menu
### 3. Select version and download file
### 4. Open downloaded file and copy/cut manifest.json to mods folder
### 5. Open terminal in mods folder and Run the following command
``` bash
minecraft-forge-mods-downloader
```
<br>

## Program options
### manifest (string) 
- Manifest file location. (default "manifest.json")
### poollimit (int)
- Download pool limit. (default 100)

<br>

## Example of manifest.json file
``` json
{
  "minecraft": {
    "version": "1.20.1",
    "modLoaders": [
      {
        "id": "forge-47.4.0",
        "primary": true
      }
    ],
    "recommendedRam": 8416
  },
  "manifestType": "minecraftModpack",
  "manifestVersion": 1,
  "name": "Beyond Depth",
  "version": "Ver10.8.1",
  "author": "blueversal",
  "files": [
    {
      "projectID": 831663,
      "fileID": 4894789,
      "required": true
    },
    {
      "projectID": 300297,
      "fileID": 6427959,
      "required": true
    },
    .
    .
    .
  ],
  "overrides": "overrides"
}
```

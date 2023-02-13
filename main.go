package main

import (
	"errors"
	"fmt"
	"os"
)

const (
	RegionInfoDownloadUrl string = "https://raw.githubusercontent.com/jerem2772/LeCrew-Server/main/regionInfo.json"
	RegionInfoFileName    string = "regionInfo.json"
)

func main() {
	DS := string(os.PathSeparator)
	UserHomeDir, err := os.UserHomeDir()
	if err != nil {
		ShowError(err)
		Stop()
		return
	}
	AmongUsRegionInfoFolderPath := fmt.Sprintf("%s%sAppData%sLocalLow%sInnersloth%sAmong Us%s", UserHomeDir, DS, DS, DS, DS, DS)
	if !dirExists(AmongUsRegionInfoFolderPath) {
		ShowError(errors.New(fmt.Sprintf("File \"%s\" does not exist", AmongUsRegionInfoFolderPath)))
		Stop()
		return
	}
	DownloadFile(RegionInfoDownloadUrl, AmongUsRegionInfoFolderPath)
	ShowSuccess(fmt.Sprintf("The %s file has been successfully updated", RegionInfoFileName))
	Stop()
}

package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

func settingsWeb(rw http.ResponseWriter, req *http.Request) {
	blob, err := json.Marshal(&SETTINGS)
	if err != nil {
		fmt.Println(err, err.Error())
	}
	io.WriteString(rw, string(blob))
}

func settingsReloadWeb(rw http.ResponseWriter, req *http.Request) {
	longest := 0
	for i := range STATUS.Targets {
		if STATUS.Targets[i].PollingInterval > longest {
			longest = STATUS.Targets[i].PollingInterval
		}
	}
	if time.Now().Unix()-SETTINGS.LastReload.Unix() < int64(longest+5) {
		io.WriteString(rw, fmt.Sprintf("Settings are still being reloaded. New settings will be applied once the longest-running application monitor checks in.  This could take up to %d seconds.", longest+5))
	} else {
		SETTINGS.LastReload = time.Now()
		WG.Add(1)
		go SETTINGS.reloadSettings()
		// show caller new settings
		io.WriteString(rw, fmt.Sprintf("Settings are being reloaded. This could take up to %d seconds.", longest+5))
	}
}

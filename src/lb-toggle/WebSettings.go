package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"runtime"
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
	RELOADSETTINGS = true

	// wait for all target gorountines to exit, leaving only main and http
	for runtime.NumGoroutine() != 2 {
		time.Sleep(time.Duration(1) * time.Second)
	}

	// set tagets to empty slice
	SETTINGS.Targets = []Target{}

	// repopulate targets from config file, presumably updated with new stuff (this calls popualteTargets)
	SETTINGS.parseSettingsFile()

	// undo routine kill condition
	RELOADSETTINGS = false

	// resume motoring with new targets and settings
	STATUS.startMonitor()

	// show caller new settings
	blob, err := json.Marshal(&SETTINGS)
	if err != nil {
		fmt.Println(err, err.Error())
	}
	io.WriteString(rw, string(blob))
}

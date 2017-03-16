package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"runtime"
	"time"

	"github.com/davecgh/go-spew/spew"
)

//Settings contains the config.json information for configuring the listening port, monitored application details, etc
type Settings struct {
	State   State    `json:"state"`
	Targets []Target `json:"targets"`
	Service struct {
		HTTPPort string `json:"http_port"` // port to listen on for web interface (5704)
		Debug    bool   `json:"debug"`
	} `json:"service"`
}

func (s *Settings) parseSettingsFile() {
	confFile := "../../init/config.json"
	if len(os.Args) > 1 {
		confFile = os.Args[1]
	}

	file, err := os.Open(confFile)
	defer file.Close()
	if err != nil {
		log.Fatal(err)
	}

	fileContent, err := os.Open(confFile)
	if err != nil {
		fmt.Println("Could not open config file")
	}

	jsonParser := json.NewDecoder(fileContent)
	if err = jsonParser.Decode(&s); err != nil {
		fmt.Println("Could not load config file. Check JSON formatting.")
	}

	// Populate global STATUS with targets from config file
	s.populateTargets()

	if SETTINGS.Service.Debug {
		spew.Dump(SETTINGS)
	}
}

func (s *Settings) populateTargets() {
	STATUS.Version = VERSION
	for i := range s.Targets {
		s.Targets[i].ID = i
		fmt.Println("Initializing:", s.Targets[i].ID, s.Targets[i].PollingInterval, s.Targets[i].Name, s.Targets[i].Endpoint)
		STATUS.Targets = append(STATUS.Targets, s.Targets[i])
	}
}

func (s *Settings) reloadSettings() {
	RELOADSETTINGS = true

	// wait for all target gorountines to exit, leaving only main and http
	for runtime.NumGoroutine() > 5 {
		time.Sleep(time.Duration(1) * time.Second)
	}

	// set tagets to empty slice
	SETTINGS.Targets = []Target{}

	// repopulate targets from config file, presumably updated with new stuff
	// (this calls popualteTargets)
	SETTINGS.parseSettingsFile()

	// undo routine kill condition
	RELOADSETTINGS = false

	// resume motoring with new targets and settings
	STATUS.startMonitor()
	WG.Done()
}

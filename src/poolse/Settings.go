package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"time"
)

//Settings contains the config.json information for configuring the listening port, monitored application details, etc
type Settings struct {
	State      State    `json:"state"`
	Targets    []Target `json:"targets"`
	LastReload time.Time
	Service    struct {
		StateFileName   string `json:"state_file_name"`
		HTTPPort        string `json:"http_port"` // port to listen on for web interface (5704)
		Debug           bool   `json:"debug"`
		ShowHTTPLog     bool   `json:"show_http_log"`
		FollowRedirects bool   `json:"follow_redirects"`
	} `json:"service"`
}

func (s *Settings) load() {
	// fill settings struct from config file
	s.parseSettingsFile()

	// copy Settings struct targets to Status
	s.populateTargets()
}

func (s *Settings) parseSettingsFile() {
	SettingsMu.Lock()
	defer SettingsMu.Unlock()

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

	if s.Service.StateFileName == "" {
		s.Service.StateFileName = "state.dat"
	}
	if s.State.PersistState {
		s.State.loadState(s.Service.StateFileName)
	}

	// apply the settings state to the STATUS state
	STATUS.State = s.State
}

func (s *Settings) populateTargets() {
	StatusMu.Lock()
	SettingsMu.Lock()

	defer StatusMu.Unlock()
	defer SettingsMu.Unlock()

	STATUS.Version = VERSION
	for i := range s.Targets {
		s.Targets[i].ID = i
		if s.Targets[i].DownCountThreshold <= 0 {
			s.Targets[i].DownCountThreshold = 1
		}
		if s.Targets[i].UpCountThreshold <= 0 {
			s.Targets[i].UpCountThreshold = 1
		}
		fmt.Println("Initializing:",
			s.Targets[i].ID,
			s.Targets[i].PollingInterval,
			s.Targets[i].Name,
			s.Targets[i].Endpoint)
		STATUS.Targets = append(
			STATUS.Targets,
			s.Targets[i])
	}
}

func (s *Settings) reloadSettings() {
	GlobalWaitGroupHelper(true)
	StopChan <- true
	// repopulate targets from config file, presumably updated with new stuff
	s.parseSettingsFile()
	s.populateTargets()

	// resume motoring with new targets and settings
	STATUS.startMonitor(StopChan)
	GlobalWaitGroupHelper(false)
}

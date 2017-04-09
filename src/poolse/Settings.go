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
		StateFileName string `json:"state_file_name"`
		HTTPPort      string `json:"http_port"` // port to listen on for web interface (5704)
		Debug         bool   `json:"debug"`
		ShowHTTPLog   bool   `json:"show_http_log"`
	} `json:"service"`
}

func (s *Settings) load() {
	s.parseSettingsFile()

	// Populate global STATUS with targets from config file
	s.populateTargets()
}

func (s *Settings) checkStartupState() {
	WG.Add(1)
	go func() {
		if s.State.StartupState {
			// give the targets a bit to catch up
			time.Sleep(time.Duration(findLongestPollingInterval(s.Targets)+3) * time.Second)
			if STATUS.isOk() {
				STATUSMUTEX.Lock()
				STATUS.State.OK = true
				STATUSMUTEX.Unlock()
			}
		}
		WG.Done()
	}()
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

	s.stopAllTargetMonitors()

	// repopulate targets from config file, presumably updated with new stuff
	s.parseSettingsFile()
	s.populateTargets()

	// resume motoring with new targets and settings
	STATUS.startMonitor()
	s.checkStartupState()
	WG.Done()
}

func (s *Settings) stopAllTargetMonitors() {
	RTMUTEX.Lock()
	for i := range s.Targets {
		RTARGETS = append(RTARGETS, s.Targets[i].ID)
	}
	RTMUTEX.Unlock()
	waitingForReset := true
	// wait for all target goroutines to exit, leaving only main and http
	for !waitingForReset {
		RTMUTEX.Lock()
		defer RTMUTEX.Unlock()
		if !(len(RTARGETS) > 0) {
			RTMUTEX.Unlock()
			time.Sleep(time.Duration(1) * time.Second)
		} else {
			RTMUTEX.Unlock()
			waitingForReset = false
		}

	}

	// set SETTINGS and STATUS to empty struct
	SETTINGSMUTEX.Lock()
	defer SETTINGSMUTEX.Unlock()
	SETTINGS = Settings{}

	STATUSMUTEX.Lock()
	defer STATUSMUTEX.Unlock()
	STATUS = Status{}
}

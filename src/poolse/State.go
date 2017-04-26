package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

// State represents the way the application handles administratively on/off
type State struct {
	OK                  bool   `json:"ok"`                   // value looked at for true app status
	StartupState        bool   `json:"startup_state"`        // whether State.OK should be true or false upon monitor startup
	PersistState        bool   `json:"persist_state"`        // if administratively up or down, will store that in a file and pull it in upon settings reload
	AdministrativeState string `json:"administrative_state"` // something
}

func (s *State) saveState(stateFile string) {

	// save STATUS.State.AdministrativeState to file
	f, err := os.Create(stateFile)
	defer f.Close()
	if err != nil {
		fmt.Println(err)
	}
	_, err = f.WriteString(STATUS.State.AdministrativeState)
	if err != nil {
		fmt.Println(err)
	}
	// Issue a Sync to flush writes to stable storage.
	f.Sync()
}

func (s *State) loadState(stateFile string) {
	// load administrative_state from file, and impart value to the SETTINGS.State.AdministrativeState field
	v, err := ioutil.ReadFile(stateFile)
	if err != nil {
		fmt.Println(err)
	}
	s.AdministrativeState = string(v)
	STATUS.State.AdministrativeState = string(v)
}

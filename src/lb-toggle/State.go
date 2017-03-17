package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

// State represents the way the application handles administratively on/off
type State struct {
	OK                  bool   `'json:"ok"`                  // value looked at for true app status
	StartupState        string `json:"startup_state"`        // whether State.OK should be true or false upon monitor startup
	PersistState        bool   `json:"persist_state"`        // if administratively up or down, will store that in a file and pull it in upon settings reload
	AdministrativeState string `json:"administrative_state"` // something
}

func (s *State) saveState() {
	// load administrative_state from file, and impart value to the SETTINGS.State.AdministrativeState field
	f, err := os.Create("state.dat")
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

func (s *State) loadState() {
	// save STATUS.State.AdministrativeState to file
	v, err := ioutil.ReadFile("state.dat")
	if err != nil {
		fmt.Println(err)
	}
	s.AdministrativeState = string(v)
	STATUS.State.AdministrativeState = string(v)
}

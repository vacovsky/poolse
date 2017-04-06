package main

import (
	"fmt"
)

// Status indicates the state of the application being monitored
type Status struct {
	State   State
	Targets []Target
	Version string
}

func (s *Status) startMonitor() {
	for i := range STATUS.Targets {
		if SETTINGS.Service.Debug {
			fmt.Println("Starting ",
				STATUS.Targets[i].Name,
				STATUS.Targets[i].Endpoint)
		}
		WG.Add(1)
		go STATUS.Targets[i].Monitor()
	}
}

func (s *Status) toggleOn() {
	// check all endpoints, and if all pass the checks, set STATUS.State to true
	safe := true
	for _, t := range s.Targets {
		if !t.OK {
			safe = false
		}
	}
	s.State.OK = safe
}

func (s *Status) toggleOff() {
	s.State.OK = false
}

func (s *Status) toggle() {
	if !s.State.OK {
		safe := true
		for _, t := range STATUS.Targets {
			if !t.OK {
				safe = false
			}
		}
		s.State.OK = safe
	} else {
		s.State.OK = false
	}
}

func (s Status) isOk() bool {
	ok := true
	for _, t := range s.Targets {
		if !t.OK {
			ok = false
		}
	}
	//if ok && s.State.StartupState {
	//	s.State.OK = true
	//} else if ok && !s.State.StartupState {
	//	s.State.OK = false
	//} else {
	//	s.State.OK = false
	//}
	return ok
}

func (s *Status) toggleAdminStateOff() {
	s.State.AdministrativeState = "AdminOff"
	if s.State.PersistState {
		s.State.saveState(SETTINGS.Service.StateFileName)
	}
}

func (s *Status) toggleAdminStateOn() {
	s.State.AdministrativeState = "AdminOn"
	if s.State.PersistState {
		s.State.saveState(SETTINGS.Service.StateFileName)
	}
}

func (s *Status) toggleResetAdminState() {
	s.State.AdministrativeState = ""
	if s.State.PersistState {
		s.State.saveState(SETTINGS.Service.StateFileName)
	}
}

func (s *Status) checkStatusByID(id int) bool {
	if (STATUS.Targets[id].OK &&
		!(STATUS.State.AdministrativeState == "AdminOff")) ||
		STATUS.State.AdministrativeState == "AdminOn" {
		return true
	}
	return false
}

func (s *Status) checkStatus() bool {
	if (STATUS.isOk() && STATUS.State.OK &&
		!(STATUS.State.AdministrativeState == "AdminOff")) ||
		STATUS.State.AdministrativeState == "AdminOn" {
		return true
	}
	return false
}

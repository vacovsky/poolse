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
			fmt.Println("Starting ", STATUS.Targets[i].Name, STATUS.Targets[i].Endpoint)
		}
		WG.Add(1)
		go STATUS.Targets[i].Monitor()
	}
}

func (s *Status) toggleOn() {
	// TODO: check all endpoints, and if all pass the checks, set STATUS.State to true
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
	return ok
}

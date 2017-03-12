package main

// Status indicates the state of the application being monitored
type Status struct {
	State   bool
	Targets []Target
	Version string
}

func (s *Status) startMonitor() {
	for _, target := range STATUS.Targets {
		WG.Add(1)
		go target.Monitor()
	}
}

func (s *Status) toggleOn() {
	// check all endpoints, and if all pass the checks, set STATUS.State to true
}

func (s *Status) toggleOff() {
	s.State = false
}

func (s *Status) toggle() {
	if !s.State {
		if s.HealthStatus.OK && s.SmokeStatus.OK {
			s.State = !s.State
		}
	} else {
		s.State = !s.State
	}
}

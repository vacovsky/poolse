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
	// TODO: check all endpoints, and if all pass the checks, set STATUS.State to true
	safe := true
	for _, t := range s.Targets {
		if !t.OK {
			safe = false
		}
	}
	s.State = safe
}

func (s *Status) toggleOff() {
	s.State = false
}

func (s *Status) toggle() {
	if !s.State {
		safe := true
		for _, t := range STATUS.Targets {
			if !t.OK {
				safe = false
			}
		}
		s.State = safe
	} else {
		s.State = false
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

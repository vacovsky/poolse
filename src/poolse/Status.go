package main

// Status indicates the state of the application being monitored
type Status struct {
	State   State
	Targets []Target
	Version string
}

func counterCleaner(tt *Target) {
	if tt.DownCount > 2147483640 {
		tt.DownCount = tt.DownCountThreshold + 1
	}
	if tt.UpCount > 2147483640 {
		tt.UpCount = tt.UpCountThreshold + 1
	}
}

func (s *Status) startMonitor() {
	updater := make(chan *Target)
	go func() {
		GlobalWaitGroupHelper(true)
		defer GlobalWaitGroupHelper(false)

		for !TARGETSTOP {
			// receive the target pointer from the channel
			var tt = <-updater

			func() {
				StatusMu.Lock()
				defer StatusMu.Unlock()
				if STATUS.State.StartupState {
					if STATUS.isOk() {
						STATUS.State.OK = true
					} else {
						STATUS.State.OK = false
					}
				}
			}()

			// ensure counter isn't going to break
			counterCleaner(tt)

			// send it back to monitor stuff
			go tt.Monitor(updater)
		}
		close(updater)
	}()

	// loop over targets and kick them off to get the ball rolling.  after this, the lambda handles it
	for i := range STATUS.Targets {
		go s.Targets[i].Monitor(updater)
	}
}

func (s *Status) toggleOn() {
	// check all endpoints, and if all pass the checks, set STATUS.State to true
	s.State.OK = s.isOk()
}

func (s *Status) toggleOff() {
	s.State.OK = false
}

func (s *Status) toggle() {
	s.State.OK = s.isOk()
}

func (s Status) isOk() bool {
	ok := true
	for _, t := range s.Targets {
		if !t.OK {
			ok = false
			break
		}
	}
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
	if (s.Targets[id].OK &&
		!(s.State.AdministrativeState == "AdminOff")) ||
		s.State.AdministrativeState == "AdminOn" {
		return true
	}
	return false
}

func (s *Status) checkStatus() bool {
	StatusMu.Lock()
	defer StatusMu.Unlock()
	if (s.isOk() && s.State.OK &&
		!(s.State.AdministrativeState == "AdminOff")) ||
		s.State.AdministrativeState == "AdminOn" {
		return true
	}
	return false
}

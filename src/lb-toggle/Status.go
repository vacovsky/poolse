package main

import (
	"fmt"
	"net/http"
	"time"
)

// Status indicates the state of the application being monitored
type Status struct {
	State        bool
	HealthStatus struct {
		OK       bool
		Last     time.Time
		Endpoint string
		Interval int
	}
	SmokeStatus struct {
		OK       bool
		Last     time.Time
		Endpoint string
		Interval int
	}
}

func (s *Status) checkAppSmoke() {
	r, err := http.Get(SETTINGS.Target.SmokeEndpoint)
	if err != nil {
		fmt.Println(err.Error())
	}
	if r.StatusCode == 200 {
		STATUS.SmokeStatus.OK = true
		// bodyBytes, err2 := ioutil.ReadAll(r.Body)
		// bodyString = string(bodyBytes)
	} else {
		STATUS.HealthStatus.OK = false
	}
	STATUS.SmokeStatus.Last = time.Now()
}

func (s *Status) checkAppHealth() {
	r, err := http.Get(SETTINGS.Target.HealthEndpoint)
	if err != nil {
		fmt.Println(err.Error())
	}
	defer r.Body.Close()
	if r.StatusCode == 200 {
		STATUS.HealthStatus.OK = true
	} else {
		STATUS.HealthStatus.OK = false
	}
	STATUS.HealthStatus.Last = time.Now()
}

func (s *Status) startHealthMonitor() {
	for {
		s.checkAppSmoke()
		time.Sleep(time.Duration(SETTINGS.Service.HealthInterval) * time.Second)
	}
}

func (s *Status) startSmokeMonitor() {
	for {
		s.checkAppHealth()
		time.Sleep(time.Duration(SETTINGS.Service.SmokeInterval) * time.Second)
	}
}

func (s *Status) toggleOn() {
	s.checkAppHealth()
	s.checkAppSmoke()
	if s.SmokeStatus.OK && s.HealthStatus.OK {
		s.State = true
	}
}

func (s *Status) toggleOff() {
	s.State = false
}

func (s *Status) toggle() {
	if !s.State {
		s.checkAppHealth()
		s.checkAppSmoke()
		if s.HealthStatus.OK && s.SmokeStatus.OK {
			s.State = !s.State
		}
	} else {
		s.State = !s.State
	}
}

package main

import (
	"net/http"
	"time"
)

// Target models the information required to perform a status check against an HTTP endpoint at interval
type Target struct {
	ID                       int       `json:"id"`
	Name                     string    `json:"name"`
	Endpoint                 string    `json:"endpoint"`
	PollingInterval          int       `json:"polling_interval"`
	ExpectedStatusCode       int       `json:"expected_status_code"`
	ExpectedResponseString   string    `json:"expected_response_string"`
	UnexpectedResponseString string    `json:"unexpected_response_string"`
	LastOK                   time.Time `json:"last_ok"`
	LastChecked              time.Time `json:"laste_checked"`
	OK                       bool      `json:"ok"`
}

// Monitor initiates the target montitor using target properties
func (t *Target) Monitor() {
	defer WG.Done()
	for {
		// get response body

		// take a snooze
		time.Sleep(time.Duration(t.PollingInterval) * time.Second)
		r, err := http.Get(t.Endpoint)
		if err != nil {
			t.OK = false
			t.LastChecked = time.Now()
			r.Body.Close()
		} else {
			s := string(byteArray[:])
			t.validateResultBody()
		}

		r.Body.Close()
		// 	if r.StatusCode == 200 {
		// 		STATUS.HealthStatus.OK = true
		// 	} else {
		// 		STATUS.HealthStatus.OK = false
		// 	}
		// 	STATUS.HealthStatus.Last = time.Now()
		// }
	}
}

func (t *Target) validateResultBody(body string) {

}

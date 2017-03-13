package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
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
	LastChecked              time.Time `json:"last_checked"`
	OK                       bool      `json:"ok"`
}

// Monitor initiates the target montitor using target properties
func (t *Target) Monitor() {
	defer WG.Done()
	for {
		thisIterState := true
		bodyString := ""
		// get response body
		r, err := http.Get(t.Endpoint)

		// if unable to connect, mark failed and move on
		if err != nil {
			r.Body.Close()
			thisIterState = false
		} else {
			r.Body.Close()
			if r.StatusCode == t.ExpectedStatusCode {
				bodyBytes, err := ioutil.ReadAll(r.Body)
				if err != nil {
					fmt.Println(err)
				}
				bodyString = string(bodyBytes)
				if t.ExpectedResponseString != "" {
					thisIterState = t.validateResultBody(true, bodyString)
				}
				if t.UnexpectedResponseString != "" {
					thisIterState = t.validateResultBody(false, bodyString)
				}
			} else {
				thisIterState = false
			}
		}
		t.OK = thisIterState
		t.LastChecked = time.Now()
		if t.OK {
			t.LastOK = t.LastChecked
		}
		if SETTINGS.Service.Debug {
			fmt.Println(t.ID, ":::", "OK:", t.OK, ":::", "Last Checked:", t.LastChecked)
		}
		// take a snooze
		time.Sleep(time.Duration(t.PollingInterval) * time.Second)
	}
}

func (t *Target) validateResultBody(shouldFind bool, body string) bool {
	r := false
	if shouldFind && strings.Contains(body, t.ExpectedResponseString) {
		r = true
	} else {
		r = false
	}

	if !shouldFind && strings.Contains(body, t.UnexpectedResponseString) {
		r = false
	} else {
		r = true
	}
	return r
}

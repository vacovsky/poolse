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
	ID                        int       `json:"id"`
	Name                      string    `json:"name"`
	Endpoint                  string    `json:"endpoint"`
	PollingInterval           int       `json:"polling_interval"`
	ExpectedStatusCode        int       `json:"expected_status_code"`
	ExpectedResponseStrings   []string  `json:"expected_response_strings"`
	UnexpectedResponseStrings []string  `json:"unexpected_response_strings"`
	LastOK                    time.Time `json:"last_ok"`
	LastChecked               time.Time `json:"last_checked"`
	OK                        bool      `json:"ok"`
}

func (t *Target) shouldReload() bool {
	ok := false
	for i := range RTARGETS {
		if t.ID == RTARGETS[i] {
			ok = true
			// remove this index from the reload targets list
			RTARGETS = append(RTARGETS[:i], RTARGETS[i+1:]...)
		}
	}
	return ok
}

// Monitor initiates the target montitor using target properties
func (t *Target) Monitor() {
	defer WG.Done()
	for !t.shouldReload() {
		thisIterState := true
		bodyString := ""
		// get response body
		r, err := http.Get(t.Endpoint)

		// if unable to connect, mark failed and move on
		if err != nil {
			r.Body.Close()
			thisIterState = false
		} else {
			if r.StatusCode == t.ExpectedStatusCode {
				bodyBytes, err := ioutil.ReadAll(r.Body)
				if err != nil {
					fmt.Println(err)
				}
				bodyString = string(bodyBytes)
				thisIterState = t.validateResultBody(bodyString)
			} else {
				thisIterState = false
			}
		}
		r.Body.Close()

		t.OK = thisIterState
		t.LastChecked = time.Now()
		if t.OK {
			t.LastOK = t.LastChecked
		}
		if SETTINGS.Service.Debug {
			fmt.Println(t.ID, ":::", "Last Checked:", t.LastChecked, t.Name, ":::", "OK:", t.OK)
		}
		// take a snooze
		time.Sleep(time.Duration(t.PollingInterval) * time.Second)
	}
}

func (t *Target) validateResultBody(body string) bool {
	r := true

	if len(t.ExpectedResponseStrings) > 0 {
		for s := range t.ExpectedResponseStrings {
			if SETTINGS.Service.Debug {
				fmt.Println(body, t.ExpectedResponseStrings[s])
			}
			if strings.Contains(body, t.ExpectedResponseStrings[s]) {
				r = true
			} else {
				r = false
			}
		}
	}
	if len(t.UnexpectedResponseStrings) > 0 {
		for s := range t.UnexpectedResponseStrings {
			if SETTINGS.Service.Debug {
				fmt.Println(body, t.UnexpectedResponseStrings[s])
			}
			if strings.Contains(body, t.UnexpectedResponseStrings[s]) {
				r = false
			} else {
				r = true
			}
		}
	}
	return r
}

package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"time"

	"github.com/davecgh/go-spew/spew"
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
	bodyString := ""
	for {
		// get response body
		r, err := http.Get(t.Endpoint)

		// if unable to connect, mark failed and move on
		if err != nil {
			defer r.Body.Close()
			t.OK = false
		} else {
			defer r.Body.Close()
			if r.StatusCode == t.ExpectedStatusCode {
				bodyBytes, err := ioutil.ReadAll(r.Body)
				if err != nil {
					fmt.Println(err)
				}
				bodyString = string(bodyBytes)
				if t.ExpectedResponseString != "" {
					t.validateResultBody(true, bodyString)
				}
				if t.UnexpectedResponseString != "" {
					t.validateResultBody(false, bodyString)
				}
			}
		}
		if SETTINGS.Service.Debug {
			spew.Dump(t)
		}
		if t.OK {
			t.LastOK = time.Now()
		}
		t.LastChecked = time.Now()
		// take a snooze
		time.Sleep(time.Duration(t.PollingInterval) * time.Second)
	}
}

func (t *Target) validateResultBody(shouldFind bool, body string) {
	if strings.Contains(body, t.ExpectedResponseString) && shouldFind {
		t.OK = true
	} else {
		t.OK = false
	}

	if strings.Contains(body, t.UnexpectedResponseString) && !shouldFind {
		t.OK = false
	} else {
		t.OK = true
	}
}

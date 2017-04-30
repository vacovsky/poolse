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
	UpCount                   int64     `json:"up_count"`
	UpCountThreshold          int64     `json:"up_count_threshold"` // this many UpCounts before marked OK again
	DownCount                 int64     `json:"down_count"`
	DownCountThreshold        int64     `json:"down_count_threshold"` // this many DownCounts before marked offline
}

// Monitor initiates the target monitor using target properties
func (t *Target) Monitor(ch chan *Target) {
	GlobalWaitGroupHelper(true)
	defer GlobalWaitGroupHelper(false)

	// this is the call to procedurally perform all checks
	thisIterState := t.checkHealth()
	thisIterState = t.validateUpDownThresholds(thisIterState)

	func() {
		StatusMu.Lock()
		defer StatusMu.Unlock()
		t.OK = thisIterState
		t.LastChecked = time.Now()
		if t.OK {
			t.LastOK = t.LastChecked
		}
	}()

	// take a snooze based on PollingInterval value
	time.Sleep(time.Duration(t.PollingInterval) * time.Second)

	// return Target pointer to channel and await next call
	ch <- t
}

func (t *Target) checkHealth() bool {
	// get response body
	var client = &http.Client{}
	
	if !SETTINGS.Service.FollowRedirects {
		client = &http.Client{
			CheckRedirect: func(req *http.Request, via []*http.Request) error {
				return http.ErrUseLastResponse
			},
		}
	}

	req, err := http.NewRequest("GET", t.Endpoint, nil)
	if err != nil {
		if SETTINGS.Service.Debug {
			fmt.Println(err)
		}
		return false
	}
	req.Header.Set("User-Agent", APPNAME+"/"+VERSION)
	r, err := client.Do(req)

	// if unable to connect, mark failed and move on
	if err != nil {
		return false
	}
	defer r.Body.Close()

	if r == nil && r.Body == nil {
		return false
	}

	if !t.validateResponseStatusCode(r) {
		return false
	}

	bodyBytes, _ := ioutil.ReadAll(r.Body)
	if !t.validateResultBody(string(bodyBytes)) {
		return false
	}

	return true
}

func (t *Target) validateUpDownThresholds(curState bool) bool {
	StatusMu.Lock()
	defer StatusMu.Unlock()
	newState := true
	if curState {
		t.DownCount = 0
		t.UpCount++
		if !(t.UpCount >= t.UpCountThreshold && t.UpCountThreshold > 0) {
			newState = false
		}
	} else {
		t.UpCount = 0
		t.DownCount++
		if t.DownCount >= t.DownCountThreshold && t.DownCountThreshold > 0 {
			newState = false
		}
	}
	return newState
}

func (t *Target) validateResponseStatusCode(r *http.Response) bool {
	match := true
	if t.ExpectedStatusCode != r.StatusCode {
		match = false
	}
	return match
}

func (t *Target) validateResultBody(body string) bool {
	r := true

	if len(t.ExpectedResponseStrings) > 0 {
		for s := range t.ExpectedResponseStrings {
			if !strings.Contains(body, t.ExpectedResponseStrings[s]) {
				r = false
			}
		}
	}
	if len(t.UnexpectedResponseStrings) > 0 {
		for s := range t.UnexpectedResponseStrings {
			if strings.Contains(body, t.UnexpectedResponseStrings[s]) {
				r = false
			}
		}
	}
	return r
}

package main

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strings"
	"time"

	"github.com/davecgh/go-spew/spew"
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
	MembersEndpoint           string    `json:"members_endpoint"`
	Members                   []Member  `json:"members"`
}

func (t *Target) loadMembers() {
	type tempMembers struct {
		Servers struct {
			Server []Member `json:"server"`
		} `json:"servers"`
	}
	var client = &http.Client{}
	StatusMu.Lock()
	var e = t.Endpoint + t.MembersEndpoint
	StatusMu.Unlock()
	req, err := http.NewRequest("GET", e, nil)
	if err != nil {
		return
	}
	req.Header.Set("User-Agent", APPNAME+"/"+VERSION)
	r, err := client.Do(req)

	var tm tempMembers
	body, err := ioutil.ReadAll(r.Body)
	err = json.Unmarshal(body, &tm)

	StatusMu.Lock()
	defer StatusMu.Unlock()
	t.Members = tm.Servers.Server
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

	StatusMu.Lock()
	checkMembers := t.MembersEndpoint != ""
	timeout := t.PollingInterval
	StatusMu.Unlock()

	if checkMembers {
		t.loadMembers()
		spew.Dump(t)
	}

	// take a snooze based on PollingInterval value
	time.Sleep(time.Duration(timeout) * time.Second)

	// return Target pointer to channel and await next call
	ch <- t
}

func (t *Target) checkHealth() bool {
	// get response body
	var client = &http.Client{}

	SettingsMu.Lock()
	if !SETTINGS.Service.FollowRedirects {
		client = &http.Client{
			CheckRedirect: func(req *http.Request, via []*http.Request) error {
				return http.ErrUseLastResponse
			},
		}
	}
	SettingsMu.Unlock()

	StatusMu.Lock()
	e := t.Endpoint
	StatusMu.Unlock()

	req, err := http.NewRequest("GET", e, nil)
	if err != nil {
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
	StatusMu.Lock()
	defer StatusMu.Unlock()

	match := true
	if t.ExpectedStatusCode != r.StatusCode {
		match = false
	}
	return match
}

func (t *Target) validateResultBody(body string) bool {
	StatusMu.Lock()
	defer StatusMu.Unlock()

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

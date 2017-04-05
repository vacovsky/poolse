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

func (t *Target) shouldReload() bool {
	ok := false
	for i := range RTARGETS {
		if t.ID == RTARGETS[i] && !ok {
			ok = true
			// remove this index from the reload targets list
			// RTARGETS = RTARGETS[:i+copy(RTARGETS[i:], RTARGETS[i+1:])]
			RTARGETS[i] = -1
			RTNULLIFY++
		}
	}
	return ok
}

// Monitor initiates the target montitor using target properties
func (t *Target) Monitor() {
	defer WG.Done()
	for !t.shouldReload() {
		thisIterState := t.checkHealth()

		if thisIterState {
			t.DownCount = 0
			t.UpCount++
			if t.UpCount >= t.UpCountThreshold && t.UpCountThreshold > 0 {
				thisIterState = true
			} else {
				thisIterState = false
			}
		} else {
			t.UpCount = 0
			t.DownCount++
			if t.DownCount >= t.DownCountThreshold && t.DownCountThreshold > 0 {
				thisIterState = false
			} else {
				thisIterState = true
			}
		}

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

	if RTNULLIFY == len(RTARGETS) {
		RTARGETS = []int{}
		RTNULLIFY = 0
	}

	if t.DownCount > 2147483640 {
		t.DownCount = 0
	}
	if t.UpCount > 2147483640 {
		t.UpCount = 0
	}
	return
}

func (t *Target) checkHealth() bool {
	// get response body
	client := &http.Client{}
	req, err := http.NewRequest("GET", t.Endpoint, nil)
	if err != nil {
		if SETTINGS.Service.Debug {
			fmt.Println(err)
		}
		return false
	}
	req.Header.Set("User-Agent", "Go-Healthcheck/"+VERSION)

	r, err := client.Do(req)
	if err != nil {
		if SETTINGS.Service.Debug {
			fmt.Println(err)
		}
		return false
	}

	// if unable to connect, mark failed and move on
	if err != nil {
		if r != nil && r.Body != nil {
			r.Body.Close()
			return false
		}

		if !t.validateResponseStatusCode(r) {
			return false
		}

		bodyBytes, _ := ioutil.ReadAll(r.Body)
		if !t.validateResultBody(string(bodyBytes)) {
			return false
		}
	}
	return true
}

func (t *Target) validateResponseStatusCode(r *http.Response) bool {
	if t.ExpectedStatusCode == r.StatusCode {
		return true
	}
	return false
}

func (t *Target) validateResultBody(body string) bool {
	r := true

	if len(t.ExpectedResponseStrings) > 0 {
		for s := range t.ExpectedResponseStrings {
			if SETTINGS.Service.Debug {
				fmt.Println(body, t.ExpectedResponseStrings[s])
			}
			if !strings.Contains(body, t.ExpectedResponseStrings[s]) {
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
			}
		}
	}
	return r

}

package main

import (
	"fmt"
	"testing"
)

func TestFindLongestRefresh(t *testing.T) {
	ts := []Target{
		Target{
			PollingInterval: 7,
		},
		Target{
			PollingInterval: 15,
		},
		Target{
			PollingInterval: 9,
		},
	}
	fmt.Println(findLongestPollingInterval(ts))
	if findLongestPollingInterval(ts) != 15 {
		t.Errorf("longest polling interval should be 15")
	}
}

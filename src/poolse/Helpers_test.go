package main

import "testing"

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
	if findLongestPollingInterval(ts) != 15 {
		t.Errorf("longest polling interval should be 15")
	}
}

func TestShowVersion(t *testing.T) {
	if showVersion() != APPNAME+" "+VERSION {
		t.Errorf("showVersion should return Foo Bar")
	}
}


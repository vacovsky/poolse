package main

import (
	"runtime"
	"testing"
)

func TestFindLongestRefresh(t *testing.T) {
	logTraffic()
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
	logTraffic()
	if showVersion() != APPNAME+" "+VERSION {
		t.Errorf("showVersion should return Foo Bar")
	}
}

func logTraffic() string {
	pc, _, _, _ := runtime.Caller(1)
	return runtime.FuncForPC(pc).Name()
}

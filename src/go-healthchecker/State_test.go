package main

import (
	"testing"
)

// TestToggleFailsBecauseHealthStatusIsFalse
func TestFullStateCycle(t *testing.T) {
	// setup state
	STATUS.State.AdministrativeState = "foo"
	SETTINGS.Service.StateFileName = "state_unittest.dat"

	// persist foo to state_unittest.dat
	STATUS.State.saveState(SETTINGS.Service.StateFileName)
	STATUS.State.AdministrativeState = "bar"

	// load state "foo" from file and stuff it into
	STATUS.State.loadState(SETTINGS.Service.StateFileName)

	if STATUS.State.AdministrativeState != "foo" {
		t.Errorf("STATUS.State.AdministrativeState should be foobar.")
	}
}

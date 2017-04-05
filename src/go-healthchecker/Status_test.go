package main

import "testing"

// TestToggleOff ensures the Status object is off when
// it is manually set to off through the method
func TestToggleOff(t *testing.T) {
	STATUS = Status{
		State: State{
			OK: true,
		},
		Targets: []Target{
			Target{
				OK: true,
			},
		},
	}
	STATUS.toggleOff()
	if STATUS.State.OK != false {
		t.Errorf("State.OK should return false, but returned true.")
	}
}

// TestToggleFailsBecauseHealthStatusIsFalse
func TestToggleFailsBecauseHealthStatusIsFalse(t *testing.T) {
	STATUS = Status{
		State: State{
			OK: true,
		},
		Targets: []Target{
			Target{
				OK: false,
			},
		},
	}
	STATUS.toggle()

	if STATUS.checkStatus() != false {
		t.Errorf("STATUS.checkStatus() should return false, but returned true.")
	}
}

// TestToggleAdminOff
func TestToggleAdminOff(t *testing.T) {
	STATUS = Status{
		State: State{
			AdministrativeState: "AdminOn",
			OK:                  true,
		},
		Targets: []Target{
			Target{
				OK: false,
			},
		},
	}
	STATUS.toggleAdminStateOff()
	if STATUS.State.AdministrativeState != "AdminOff" {
		t.Errorf("State.AdministrativeState should return \"AdminOff\", but returned " +
			STATUS.State.AdministrativeState)
	}
	if STATUS.checkStatus() != false {
		t.Errorf("STATUS.checkStatus() should return false, but returned true.")
	}
}

// TestToggleAdminOn
func TestToggleAdminOn(t *testing.T) {
	STATUS = Status{
		State: State{
			AdministrativeState: "AdminOff",
			OK:                  true,
		},
		Targets: []Target{
			Target{
				OK: false,
			},
		},
	}
	STATUS.toggleAdminStateOn()
	if STATUS.State.AdministrativeState != "AdminOn" {
		t.Errorf("State.AdministrativeState should return \"AdminOn\", but returned " +
			STATUS.State.AdministrativeState)
	}
	if STATUS.checkStatus() != true {
		t.Errorf("STATUS.checkStatus() should return true, but returned false.")
	}
}

// TestToggleAdminResetWithOK
func TestToggleAdminResetWithOK(t *testing.T) {
	STATUS = Status{
		State: State{
			AdministrativeState: "AdminOff",
			OK:                  true,
		},
		Targets: []Target{
			Target{
				OK: true,
			},
		},
	}
	STATUS.toggleResetAdminState()
	if STATUS.State.AdministrativeState != "" {
		t.Errorf("State.AdministrativeState should return \"\", but returned " +
			STATUS.State.AdministrativeState)
	}
	if STATUS.checkStatus() != true {
		t.Errorf("STATUS.checkStatus() should return true, but returned false.")
	}
}

// TestToggleAdminResetWithoutOK
func TestToggleAdminResetWithoutOK(t *testing.T) {
	STATUS = Status{
		State: State{
			AdministrativeState: "AdminOff",
			OK:                  false,
		},
		Targets: []Target{
			Target{
				OK: false,
			},
		},
	}
	STATUS.toggleResetAdminState()
	if STATUS.State.AdministrativeState != "" {
		t.Errorf("State.AdministrativeState should return \"\", but returned " +
			STATUS.State.AdministrativeState)
	}
	if STATUS.checkStatus() != false {
		t.Errorf("STATUS.checkStatus() should return false, but returned true.")
	}
}

// TestStatusOfSingleTargetByIDWithAdminReset
func TestStatusOfSingleTargetByIDWithAdminReset(t *testing.T) {
	STATUS = Status{
		State: State{
			AdministrativeState: "AdminOn",
			OK:                  true,
		},
		Targets: []Target{
			Target{
				OK: false,
			},
		},
	}
	STATUS.toggleResetAdminState()
	if STATUS.State.AdministrativeState != "" {
		t.Errorf("State.AdministrativeState should return \"\", but returned " +
			STATUS.State.AdministrativeState)
	}
	if STATUS.checkStatusByID(0) != false {
		t.Errorf("STATUS.checkStatus() should return false, but returned true.")
	}
}

// TestStatusOfSingleTargetByIDWithAdminOn
func TestStatusOfSingleTargetByIDWithAdminOn(t *testing.T) {
	STATUS = Status{
		State: State{
			AdministrativeState: "",
			OK:                  false,
		},
		Targets: []Target{
			Target{
				OK: false,
			},
		},
	}
	STATUS.toggleAdminStateOn()
	if STATUS.State.AdministrativeState != "AdminOn" {
		t.Errorf("State.AdministrativeState should return \"AdminOn\", but returned " +
			STATUS.State.AdministrativeState)
	}
	if STATUS.checkStatusByID(0) != true {
		t.Errorf("STATUS.checkStatus() should return true, but returned false.")
	}
}

// TestStatusOfSingleTargetByIDWithAdminOff
func TestStatusOfSingleTargetByIDWithAdminOff(t *testing.T) {
	STATUS = Status{
		State: State{
			AdministrativeState: "",
			OK:                  false,
		},
		Targets: []Target{
			Target{
				OK: true,
			},
		},
	}
	STATUS.toggleAdminStateOff()
	if STATUS.State.AdministrativeState != "AdminOff" {
		t.Errorf("State.AdministrativeState should return \"AdminOff\", but returned " +
			STATUS.State.AdministrativeState)
	}
	if STATUS.checkStatusByID(0) != false {
		t.Errorf("STATUS.checkStatus() should return false, but returned true.")
	}
}

package main

import (
	"fmt"
	"testing"
)

// TestToggleOff ensures the Status object is off when
// it is manually set to off through the method
func TestToggleOff(*testing.T) {
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
	fmt.Println(STATUS.State)

	// Output:
	// false
}

// TestToggleFailsBecauseHealthStatusIsFalse
func TestToggleFailsBecauseHealthStatusIsFalse(*testing.T) {
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
	fmt.Println(STATUS.State)
	// Output:
	// false
}

// TestToggleAdminOff
func TestToggleAdminOff(*testing.T) {
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
	fmt.Println(STATUS.State.AdministrativeState)
	// Output:
	// "AdminOff"
}

// TestToggleAdminOn
func TestToggleAdminOn(*testing.T) {
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
	fmt.Println(STATUS.State.AdministrativeState)
	fmt.Println(STATUS.checkStatus())

	// Output:
	// "AdminOn"
	// true
}

// TestToggleAdminResetWithOK
func TestToggleAdminResetWithOK(*testing.T) {
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
	STATUS.toggleResetAdminState()
	fmt.Println(STATUS.State.AdministrativeState)
	fmt.Println(STATUS.checkStatus())

	// Output:
	// ""
	// true
}

// TestToggleAdminResetWithoutOK
func TestToggleAdminResetWithoutOK(*testing.T) {
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
	fmt.Println(STATUS.State.AdministrativeState)
	fmt.Println(STATUS.checkStatus())
	// Output:
	// ""
	// false
}

// TestStatusOfSingleTargetByIDWithAdminReset
func TestStatusOfSingleTargetByIDWithAdminReset(*testing.T) {
	STATUS = Status{
		State: State{
			AdministrativeState: "AdminOn",
			OK:                  false,
		},
		Targets: []Target{
			Target{
				OK: false,
			},
		},
	}
	STATUS.toggleAdminStateOff()
	fmt.Println(STATUS.State.AdministrativeState)
	fmt.Println(STATUS.checkStatusByID(0))
	// Output:
	// ""
	// true
}

// TestStatusOfSingleTargetByIDWithAdminOn
func TestStatusOfSingleTargetByIDWithAdminOn(*testing.T) {
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
	STATUS.toggleResetAdminState()
	fmt.Println(STATUS.State.AdministrativeState)
	fmt.Println(STATUS.checkStatus())
	// Output:
	// ""
	// true
}

// TestStatusOfSingleTargetByIDWithAdminOff
func TestStatusOfSingleTargetByIDWithAdminOff(*testing.T) {
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
	fmt.Println(STATUS.checkStatusByID(0))
	// Output:
	// false
}

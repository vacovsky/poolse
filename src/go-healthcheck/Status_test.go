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
	// Output:
	// "AdminOn"
}

// TestToggleAdminReset
func TestToggleAdminReset(*testing.T) {
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
	// Output:
	// ""
}

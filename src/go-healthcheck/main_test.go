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

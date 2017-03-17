package main

import (
	"fmt"
	"testing"
)

// TestToggleOff ensures the Status object is off when
// it is manually set to off through the method
func TestToggleOff(*testing.T) {
	STATUS = Status{}
	STATUS.State = true
	STATUS.HealthStatus.OK = true
	STATUS.SmokeStatus.OK = true

	STATUS.toggleOff()

	fmt.Println(STATUS.State)
	// Output:
	// false
}

// TestToggleFailsBecauseHealthStatusIsFalse
func TestToggleFailsBecauseHealthStatusIsFalse(*testing.T) {
	STATUS = Status{}
	STATUS.State = false
	STATUS.HealthStatus.OK = false
	STATUS.SmokeStatus.OK = true

	// STATUS.toggle()

	fmt.Println(STATUS.State)
	// Output:
	// false
}

package main

import (
	"testing"
)

// TestValidateResultBodyExpectedResult returns false because the multiple expected strings are not found.
func TestValidateResultBodyMultipleExpectedResult(t *testing.T) {
	target := Target{
		ID:      0,
		UpCount: 1,
		OK:      true,
	}

	target.shouldReload()
}

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

	target.ExpectedResponseStrings = []string{"we should", "icanteventell"}

	fakeBody := `
		thisisagiant
		wallofnonsense
		maybe we should look
		for this string
		but who can besure
		icanteventell
		something
		`
	target.validateResultBody(fakeBody)

	if target.OK != true {
		t.Errorf("Target.OK should be true, but returned false.")
	}

	//target.shouldReload()
}

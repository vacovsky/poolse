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
	if !target.validateResultBody(fakeBody) {

		t.Errorf("Target.OK should be true, but returned false.")
	}

	//target.shouldReload()
}

// TestValidateResultBodyExpectedResultOneNotFound returns false because the multiple expected strings are not found.
func TestValidateResultBodyMultipleExpectedResultOneNotFound(t *testing.T) {
	target := Target{
		ID:      0,
		UpCount: 1,
		OK:      true,
	}

	target.ExpectedResponseStrings = []string{"we should", "write more unit tests"}

	fakeBody := `
		thisisagiant
		wallofnonsense
		maybe we should look
		for this string
		but who can besure
		icanteventell
		something
		`
	if target.validateResultBody(fakeBody) {
		t.Errorf("Target.OK should be false, but returned true.")
	}

	//target.shouldReload()
}

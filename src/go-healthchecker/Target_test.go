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

		t.Errorf("Should be true, but returned false.")
	}
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
		t.Errorf("Should be false, but returned true.")
	}
}

// TestValidateResultBodyUnexpectedResult returns false because not all expected strings are not found.
func TestValidateResultBodyMultipleUnexpectedResult(t *testing.T) {
	target := Target{
		ID:      0,
		UpCount: 1,
		OK:      true,
	}

	target.UnexpectedResponseStrings = []string{"we should", "write more unit tests"}

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
		t.Errorf("Should be false, but returned true.")
	}
}

// TestValidateResultBodyUnexpectedResultOneFound returns true because at least one of multiple unexpected string are found.
func TestValidateResultBodyMultipleUnexpectedResultOneFound(t *testing.T) {
	target := Target{
		ID:      0,
		UpCount: 1,
		OK:      true,
	}

	target.UnexpectedResponseStrings = []string{"we fburigbdiugbudiblu", "write more unit tests"}

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
}

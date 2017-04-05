package main

import (
	"net/http/httptest"
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

func TestStatusCodeComparisonSuccess(t *testing.T) {
	target := Target{
		ExpectedStatusCode: 200,
	}
	h := httptest.NewRecorder().Result()
	h.StatusCode = 200

	if !target.validateResponseStatusCode(h) {
		t.Errorf("Should return true - status code and ExpectedStatusCode matches expected.")
	}
}

func TestStatusCodeComparisonFails(t *testing.T) {
	target := Target{
		ExpectedStatusCode: 200,
	}
	h := httptest.NewRecorder().Result()
	h.StatusCode = 500

	if target.validateResponseStatusCode(h) {
		t.Errorf("Should return false - status code does not match expected.")
	}
}

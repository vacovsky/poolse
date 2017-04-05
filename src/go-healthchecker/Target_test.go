package main

import (
	"net/http/httptest"
	"testing"
)

// TestValidateResultBodyExpectedResult returns false because the multiple
// expected strings are not found.
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

// TestValidateResultBodyExpectedResultOneNotFound returns false because the
// multiple expected strings are not found.
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

// TestValidateResultBodyUnexpectedResult returns false because not
// all expected strings are not found.
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

// TestValidateResultBodyUnexpectedResultOneFound returns true because at
// least one of multiple unexpected string are found.
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

// TestStatusCodeComparison* tests that comparing status code returns the
// correct boolean value
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

// TestValidateUpDownThreshold* tests that incrementing success and failures
// are correctly calculated
func TestValidateUpDownThresholdWithUnmetUpLimit(t *testing.T) {
	target := Target{
		UpCountThreshold:   3,
		UpCount:            0,
		DownCount:          78,
		DownCountThreshold: 3,
	}
	o := target.validateUpDownThresholds(true)
	if o {
		t.Errorf("Should return false because Up threshold was not met")
	}
	if target.DownCount != 0 || target.UpCount != 1 {
		t.Errorf("UpCount should have been incremented by 1 and DownCount should remain at 0.")
	}

}

func TestValidateUpDownThresholdWithMetUpLimit(t *testing.T) {
	target := Target{
		UpCountThreshold:   3,
		UpCount:            3,
		DownCount:          0,
		DownCountThreshold: 3,
	}
	o := target.validateUpDownThresholds(true)
	if !o {
		t.Errorf("Should return true because Up threshold was met")
	}
	if target.DownCount != 0 || target.UpCount != 4 {
		t.Errorf("UpCount should have been incremented by 1 and DownCount should remain at 0.")
	}
}

func TestValidateUpDownThresholdWithUnmetDownLimit(t *testing.T) {
	target := Target{
		UpCountThreshold:   3,
		UpCount:            5,
		DownCount:          0,
		DownCountThreshold: 3,
	}
	o := target.validateUpDownThresholds(false)
	if !o {
		t.Errorf("Should return true because Down threshold was not met")
	}
	if target.UpCount > 0 || target.DownCount != 1 {
		t.Errorf("UpCount should be reset to 0 upon first detected down; DownCount should be incremented by one.")
	}
}

func TestValidateUpDownThresholdWithMetDownLimit(t *testing.T) {
	target := Target{
		UpCountThreshold:   3,
		UpCount:            0,
		DownCount:          3,
		DownCountThreshold: 3,
	}
	o := target.validateUpDownThresholds(false)
	if o {
		t.Errorf("Should return false because Down threshold was met")
	}
	if target.UpCount != 0 || target.DownCount != 4 {
		t.Errorf("UpCount should remain at 0 and Downcount should be incremented by 1.")
	}
}

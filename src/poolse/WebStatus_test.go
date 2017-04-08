package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func Test_StatusSimpleWeb_200(t *testing.T) {
	STATUS := Status{
		Targets: []Target{
			Target{
				OK: true,
			},
		},
	}
	STATUS.State.OK = true

	STATUS.State.OK = true
	STATUS.State.AdministrativeState = "AdminOn"
	req, err := http.NewRequest("GET", "/status/simple", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(statusSimpleWeb)
	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	// Check the response body is what we expect.
	//expected := `{"alive": true}`
	//if rr.Body.String() != expected {
	//	t.Errorf("handler returned unexpected body: got %v want %v",
	//		rr.Body.String(), expected)
	//}

}

func Test_StatusSimpleWeb_503(t *testing.T) {
	STATUS := Status{
		Targets: []Target{
			Target{
				OK: false,
			},
		},
	}
	STATUS.State.OK = false

	req, err := http.NewRequest("GET", "/status/simple", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(statusSimpleWeb)
	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusServiceUnavailable {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusServiceUnavailable)
	}
}

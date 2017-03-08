package main

import (
	"net/http"
)

func toggleWeb(rw http.ResponseWriter, req *http.Request) {
	STATUS.toggle()
	statusWeb(rw, req)
}

func toggleOnWeb(rw http.ResponseWriter, req *http.Request) {
	STATUS.toggleOn()
	statusWeb(rw, req)
}

func toggleOffWeb(rw http.ResponseWriter, req *http.Request) {
	STATUS.toggleOff()
	statusWeb(rw, req)
}

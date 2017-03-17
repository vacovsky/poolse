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

func toggleAdminStateOffWeb(rw http.ResponseWriter, req *http.Request) {
	STATUS.toggleAdminStateOff()
	statusWeb(rw, req)
}

func toggleAdminStateOnWeb(rw http.ResponseWriter, req *http.Request) {
	STATUS.toggleAdminStateOn()
	statusWeb(rw, req)
}

func toggleResetAdminStateWeb(rw http.ResponseWriter, req *http.Request) {
	STATUS.toggleResetAdminState()
	statusWeb(rw, req)
}

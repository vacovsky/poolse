package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"strconv"
)

func statusWeb(rw http.ResponseWriter, req *http.Request) {
	if SETTINGS.Service.ShowHTTPLog {
		SERVEDCOUNT++
		logRequest(req)
	}
	// do we have a query param?
	req.ParseForm()
	ppid := req.Form.Get("id")
	id, err := strconv.Atoi(ppid)
	if err == nil && id >= 0 && id <= len(STATUS.Targets) {
		blob, _ := json.Marshal(&STATUS.Targets[id])
		io.WriteString(rw, string(blob))
	} else {
		blob, err := json.Marshal(&STATUS)
		if err != nil {
			fmt.Println(err, err.Error())
		}
		io.WriteString(rw, string(blob))
	}
}

func statusSimpleWeb(rw http.ResponseWriter, req *http.Request) {
	result := false
	id := -1

	if SETTINGS.Service.ShowHTTPLog {
		SERVEDCOUNT++
		logRequest(req)
	}
	req.ParseForm()
	ppid := req.Form.Get("id")
	id, err = strconv.Atoi(ppid)

	if err == nil && id >= 0 && id < len(STATUS.Targets) {
		result = STATUS.checkStatusByID(id)
	} else {
		result = STATUS.checkStatus()
	}

	if result {
		rw.WriteHeader(http.StatusOK)
	} else {
		http.Error(rw, http.StatusText(http.StatusServiceUnavailable), http.StatusServiceUnavailable)
	}
}

func statusSimple2Web(rw http.ResponseWriter, req *http.Request) {
	if SETTINGS.Service.ShowHTTPLog {
		SERVEDCOUNT++
		logRequest(req)
	}
	req.ParseForm()
	ppid := req.Form.Get("id")
	id, err := strconv.Atoi(ppid)
	if err == nil && id >= 0 && id <= len(STATUS.Targets) {
		if (STATUS.Targets[id].OK && !(STATUS.State.AdministrativeState == "AdminOff")) || STATUS.State.AdministrativeState == "AdminOn" {
			rw.WriteHeader(http.StatusOK)
		} else {
			var err = errors.New("intentionally erroring for sake of not returning anything to the caller")
			panic(err)
		}
	} else {
		if (STATUS.isOk() && STATUS.State.OK && !(STATUS.State.AdministrativeState == "AdminOff")) || STATUS.State.AdministrativeState == "AdminOn" {
			rw.WriteHeader(http.StatusOK)
		} else {
			var err = errors.New("intentionally erroring for sake of not returning anything to the caller")
			panic(err)
		}
	}
}

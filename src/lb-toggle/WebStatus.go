package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
)

func statusWeb(rw http.ResponseWriter, req *http.Request) {
	blob, err := json.Marshal(&STATUS)
	if err != nil {
		fmt.Println(err, err.Error())
	}
	io.WriteString(rw, string(blob))
}

func statusSimpleWeb(rw http.ResponseWriter, req *http.Request) {
	if (STATUS.isOk() && STATUS.State.OK && !(STATUS.State.AdministrativeState == "down")) || STATUS.State.AdministrativeState == "up" {
		rw.WriteHeader(http.StatusOK)
	} else {
		http.Error(rw, http.StatusText(http.StatusServiceUnavailable), http.StatusServiceUnavailable)
	}
}

func statusSimple2Web(rw http.ResponseWriter, req *http.Request) {
	if (STATUS.isOk() && STATUS.State.OK) || STATUS.State.AdministrativeState == "up" {
		rw.WriteHeader(http.StatusOK)
	} else {
		var err = errors.New("intentionally erroring for sake of not returning anything to the caller")
		panic(err)
	}
}

package main

import (
	"encoding/json"
	"fmt"
	"io"
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

func toggleAdminOffWeb(rw http.ResponseWriter, req *http.Request) {
	blob, err := json.Marshal(&STATUS)
	if err != nil {
		fmt.Println(err, err.Error())
	}
	io.WriteString(rw, string(blob))
}

func toggleAdminOnWeb(rw http.ResponseWriter, req *http.Request) {
	blob, err := json.Marshal(&STATUS)
	if err != nil {
		fmt.Println(err, err.Error())
	}
	io.WriteString(rw, string(blob))
}

package main

import (
	"io"
	"net/http"
)

func fakeSmoke(rw http.ResponseWriter, req *http.Request) {
	rw.WriteHeader(200)
}

func fakeHealth(rw http.ResponseWriter, req *http.Request) {
	rw.WriteHeader(200)
}

func fakeExpected(rw http.ResponseWriter, req *http.Request) {
	blob := "{\"is_working\": true}"
	io.WriteString(rw, blob)
}

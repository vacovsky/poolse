package main

import "net/http"

func fakeSmoke(rw http.ResponseWriter, req *http.Request) {
	rw.WriteHeader(200)
}

func fakeHealth(rw http.ResponseWriter, req *http.Request) {
	rw.WriteHeader(200)
}

package main

import (
	"io"
	"net/http"
)

func displayReadme(rw http.ResponseWriter, req *http.Request) {
	helpinfo := `Please go to https://github.com/vacoj/go-healthchecker for the latest source code and details.`
	io.WriteString(rw, helpinfo)

}

package main

import (
	"fmt"
	"net/http"
	"time"
)

// func logRequest(handle string, routes map[string]func(http.ResponseWriter, *http.Request)) func(http.ResponseWriter, *http.Request) {
// 	return routes[handle]
// }

func logRequest(r *http.Request, result bool) {
	SERVEDCOUNT++
	fmt.Println(SERVEDCOUNT, r.RemoteAddr, time.Now(), "/"+r.URL.Path[1:], " OK: ", result)
}

func startWeb() {
	routes := map[string]func(http.ResponseWriter, *http.Request){
		// return status
		"/status":         statusWeb,
		"/status/simple":  statusSimpleWeb,
		"/status/simple2": statusSimple2Web,

		// turn on/off application (for benefit of monitor or LB)
		"/toggle/on":         toggleOnWeb,
		"/toggle/adminon":    toggleAdminStateOnWeb,
		"/toggle/adminoff":   toggleAdminStateOffWeb,
		"/toggle/adminreset": toggleResetAdminStateWeb,
		"/toggle/off":        toggleOffWeb,

		"/settings": settingsWeb,

		// disabled this method because there is a bug, and it's safer this way until it gets fixed
		// "/settings/reload": settingsReloadWeb,

		"/fakesmoke":    fakeSmoke,
		"/fakehealth":   fakeHealth,
		"/fakeexpected": fakeExpected,

		// show status as landing page?
		"/help": displayReadme,
		"/":     statusWeb,
	}

	// register routes
	for k, v := range routes {
		http.HandleFunc(k, v)
	}

	// In case we want static content
	http.HandleFunc("/static/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, r.URL.Path[1:])
	})

	// Host web server
	panic(http.ListenAndServe(":"+SETTINGS.Service.HTTPPort, nil))
}

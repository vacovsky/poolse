package main

import "time"

func main() {
	showVersion()

	SETTINGS.load()

	// Monitor application for health status
	STATUS.startMonitor()

	// Start the Web application.
	go startWeb()

	SETTINGSMUTEX.Lock()
	ss := SETTINGS.State.StartupState
	SETTINGSMUTEX.Unlock()
	tt := len(STATUS.Targets)

	if ss {
		// give the targets a bit to catch up
		time.Sleep(time.Duration(tt) * time.Second)
		okay := STATUS.isOk()
		if okay {
			STATUSMUTEX.Lock()
			STATUS.State.OK = true
			STATUSMUTEX.Unlock()
		}
	}
	WG.Wait()
}

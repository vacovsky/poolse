package main

import "time"

func main() {
	showVersion()

	SETTINGS.load()

	// Monitor application for health status
	STATUS.startMonitor()

	// Start the Web application.
	go startWeb()

	if SETTINGS.State.StartupState {
		// give the targets a bit to catch up
		time.Sleep(time.Duration(len(STATUS.Targets)) * time.Second)
		if STATUS.isOk() {
			STATUSMUTEX.Lock()
			STATUS.State.OK = true
			STATUSMUTEX.Unlock()
		}
	}
	WG.Wait()
}

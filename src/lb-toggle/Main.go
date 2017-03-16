package main

import (
	"fmt"
	"runtime"
	"time"
)

func showVersion() {
	name := "LB-Toggle " + VERSION
	fmt.Println(name)
}

func main() {
	SETTINGS.parseSettingsFile()

	// Monitor application for health status
	STATUS.startMonitor()

	// Start the Web application.
	go startWeb()

	for {
		if SETTINGS.Service.Debug {
			grc := runtime.NumGoroutine()
			fmt.Println(grc, "Active goroutines as of", time.Now())
		}
		time.Sleep(time.Duration(1) * time.Second)
	}
	WG.Wait()
}

package main

import (
	"fmt"
)

func showVersion() {
	name := "Goggle " + VERSION
	fmt.Println(name)
}

func main() {
	SETTINGS.parseSettingsFile()

	// Initialize the wait group so threads don't exit
	WG.Add(2)

	// Start the Web application.
	go startWeb()

	// Monitor application for health status
	go startMonitor()
	WG.Wait()
}

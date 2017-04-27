package main

func main() {
	showVersion()

	SETTINGS.load()

	// Monitor application for health status
	STATUS.startMonitor()

	// Start the Web application.
	go startWeb()

	WG.Wait()
}

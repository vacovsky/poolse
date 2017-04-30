package main

func main() {

	showVersion()

	SETTINGS.load()

	// Monitor application for health status
	GlobalWaitGroupHelper(true)
	go STATUS.startMonitor(make(chan bool))
	// STATUS.startMonitor()

	// Start the Web application.
	go startWeb()

	WG.Wait()
}

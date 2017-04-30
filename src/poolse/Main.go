package main

func main() {

	showVersion()

	SETTINGS.load()

	// Monitor application for health status
	GlobalWaitGroupHelper(true)
	go STATUS.startMonitor(StopChan)

	// Start the Web application.
	go startWeb()

	WG.Wait()
}

package main

// State represents the way the application handles administratively on/off
type State struct {
	OK           bool   `'json:"ok"`
	StartupState string `json:"startup_state"`
	PersistState bool   `json:"persist_state"`
}

func persistState() {

}

func loadState() {

}

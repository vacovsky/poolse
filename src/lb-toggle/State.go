package main

// State represents the way the application handles administratively on/off
type State struct {
	OK                  bool   `'json:"ok"`                  // value looked at for true app status
	StartupState        string `json:"startup_state"`        // whether State.OK should be true or false upon monitor startup
	PersistState        bool   `json:"persist_state"`        // if administratively up or down, will store that in a file and pull it in upon settings reload
	AdministrativeState bool   `json:"administrative_state"` // something
}

func persistState() {
	// save administrative_state to file

}

func loadState() {
	// load administrative_state from file, and impart value to the SETTINGS.State.AdministrativeState field

}

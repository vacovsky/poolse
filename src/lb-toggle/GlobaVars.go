package main

import (
	"sync"
)

// VERSION of application
var VERSION = "0.2.1"

// WG contains list of running goroutines
var WG sync.WaitGroup

// SETTINGS Contains the loaded settings for the application
var SETTINGS Settings

// STATUS hold the last results of a status poll of the target application
var STATUS Status

// RELOADSETTINGS is how we tell goroutines to stop and reasses the configuration
var RELOADSETTINGS bool

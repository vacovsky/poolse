package main

import (
	"sync"
)

// VERSION of application
var VERSION = "0.1.0"

// WG contains list of running goroutines
var WG sync.WaitGroup

// SETTINGS Contains the loaded settings for the application
var SETTINGS Settings

// STATUS hold the last results of a status poll of the target application
var STATUS Status

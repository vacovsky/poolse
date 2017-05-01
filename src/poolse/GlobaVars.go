package main

import (
	"sync"
)

const (
	// VERSION of application
	VERSION = "0.5.2"

	// APPNAME of application.  One place to change it everywhere else.  :shrug:
	APPNAME = "Poolse"
)

var (
	// WG contains list of running goroutines
	WG sync.WaitGroup

	// SETTINGS Contains the loaded settings for the application
	SETTINGS Settings

	// STATUS hold the last results of a status poll of the target application
	STATUS Status

	// SERVEDCOUNT is the running counter of requests served
	SERVEDCOUNT int64

	// WGMUTEX protects the global waitgroup from race issues
	WGMUTEX sync.Mutex

	// StatusMu protects the global Status struct from race issues
	StatusMu sync.Mutex

	// SettingsMu protects the global Status struct from race issues
	SettingsMu sync.Mutex

	// StopChan is used to tell workers to start or stop
	StopChan = make(chan bool)
)

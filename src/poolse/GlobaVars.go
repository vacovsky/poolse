package main

import (
	"sync"
)

const (
	// VERSION of application
	VERSION = "0.4.2"

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

	// RTARGETS is how we tell goroutines to stop and reassess the configuration
	RTARGETS []int

	// RTNULLIFY when reloading settings, if this matches the len() of RTARGETS, zero both out.
	RTNULLIFY int

	// RTMUTEX is the thread safety for RTARGETS
	RTMUTEX = &sync.Mutex{}

	// SERVEDCOUNT is the running counter of requests served
	SERVEDCOUNT int64
)

package main

import (
	"sync"
)

// VERSION of application
const VERSION = "0.3.14"

// APPNAME of application.  One place to change it everywhere else.  :shrug:
const APPNAME = "Go-Healthchecker"

// WG contains list of running goroutines
var WG sync.WaitGroup

// SETTINGS Contains the loaded settings for the application
var SETTINGS Settings

// STATUS hold the last results of a status poll of the target application
var STATUS Status

// RTARGETS is how we tell goroutines to stop and reasses the configuration
var RTARGETS []int

// RTNULLIFY when reloading settings, if this matches the len() of RTARGETS, zero both out.
var RTNULLIFY int

// SERVEDCOUNT is the running counter of requests served
var SERVEDCOUNT int64

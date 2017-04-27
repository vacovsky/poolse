package main

import "fmt"

func findLongestPollingInterval(targets []Target) int64 {
	var longest int64
	for i := range targets {
		if int64(targets[i].PollingInterval) > longest {
			longest = int64(targets[i].PollingInterval)
		}
	}
	return longest
}

func findLargestUpThreshold(targets []Target) int64 {
	var largest int64
	for i := range targets {
		if targets[i].UpCountThreshold > largest {
			largest = targets[i].UpCountThreshold
		}
	}
	return largest
}

func showVersion() string {
	name := APPNAME + " " + VERSION
	fmt.Println(name)
	return name
}

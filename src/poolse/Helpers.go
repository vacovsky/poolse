package main

func findLongestPollingInterval(targets []Target) int {
	longest := 0
	for i := range targets {
		if targets[i].PollingInterval > longest {
			longest = targets[i].PollingInterval
		}
	}
	return longest
}

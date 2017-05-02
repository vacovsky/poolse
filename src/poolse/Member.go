package main

// Member is a member of a load balanced pool
type Member struct {
	Fall     int    `json:"fall"`
	Index    int    `json:"index"`
	Name     string `json:"name"`
	Port     int    `json:"port"`
	Rise     int    `json:"rise"`
	Status   string `json:"status"`
	Type     string `json:"type"`
	Upstream string `json:"upstream"`
}

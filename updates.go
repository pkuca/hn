package hn

// Updates is based on the response example here: https://github.com/HackerNews/API#changed-items-and-profiles
type Updates struct {
	Items    []int    `json:"items"`
	Profiles []string `json:"profiles"`
}

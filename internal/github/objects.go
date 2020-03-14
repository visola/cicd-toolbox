package github

// Object can represents different Git data objects, normally it is a commit
type Object struct {
	SHA  string `json:"sha"`
	Type string `json:"type"`
	URL  string `json:"url"`
}

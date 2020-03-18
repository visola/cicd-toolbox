package github

// Commit represents a GitHub commit object
type Commit struct {
	Commit GitCommit `json:"commit"`
	SHA    string    `json:"sha"`
}

// GitCommit represents a Git database commit object
type GitCommit struct {
	Message string `json:"message"`
}

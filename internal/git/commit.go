package git

// Commit represents a Git commit object
type Commit struct {
	Message string `json:"message"`
	SHA     string `json:"sha"`
}

// ShortSHA returns the short version of the commit's SHA
func (commit Commit) ShortSHA() string {
	return commit.SHA[:8]
}

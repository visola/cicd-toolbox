package git

// Commit represents a Git commit object
type Commit struct {
	Message string `json:"message"`
	SHA     string `json:"sha"`
}

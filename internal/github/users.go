package github

// User represents a GitHub user
type User struct {
	Email string `json:"email"`
	ID    int64  `json:"id"`
	Login string `json:"login"`
	Name  string `json:"name"`
}

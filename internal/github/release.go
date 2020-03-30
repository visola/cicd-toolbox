package github

// Tag represents a GitHub tag
type Tag struct {
	Message string `json:"message"`
	Object  string `json:"object"`
	Tag     string `json:"tag"`
	Tagger  Tagger `json:"tagger"`
	Type    string `json:"type"`
}

// Tagger represents the owner of a Release/Tag
type Tagger struct {
	Date  string `json:"date"`
	Email string `json:"email"`
	Name  string `json:"name"`
}

func CreateRelease(commit Commit) {

}

package nodejs

import "path/filepath"

type Introspector struct{}

func (i *Introspector) Matches(files []string) bool {
	for _, file := range files {
		filename := filepath.Base(file)
		if filename == "package.json" {
			return true
		}
	}
	return false
}

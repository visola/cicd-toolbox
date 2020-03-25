package arrayutil

// ContainsString checks if an array of strings contains the specified string or not
func ContainsString(searchIn []string, searchFor string) bool {
	for _, val := range searchIn {
		if val == searchFor {
			return true
		}
	}
	return false
}

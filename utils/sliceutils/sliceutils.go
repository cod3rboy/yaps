package sliceutils

func ContainsString(slice []string, searchValue string) bool {
	for _, val := range slice {
		if val == searchValue {
			return true
		}
	}
	return false
}

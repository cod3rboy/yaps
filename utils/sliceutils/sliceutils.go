// Package sliceutils contains common utility functions for go slices.
package sliceutils

// ContainsString returns true if searchValue is present in slice, otherwise it returns false.
func ContainsString(slice []string, searchValue string) bool {
	for _, val := range slice {
		if val == searchValue {
			return true
		}
	}
	return false
}

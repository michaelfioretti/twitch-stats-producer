package utils

// Basic "contains string" function as a placeholder - this will be filled in later
func ContainsString(slice []string, value string) bool {
	for _, item := range slice {
		if item == value {
			return true
		}
	}
	return false
}

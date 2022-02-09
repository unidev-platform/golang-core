package xcollection

import "math/rand"

// StringBoolMapKeys - extract map keys as slice
func StringBoolMapKeys(m map[string]bool) []string {
	keys := make([]string, 0, len(m))
	for kk := range m {
		keys = append(keys, kk)
	}
	return keys
}

// StringRandomElement - fetch random element from string slice
func StringRandomElement(slice []string) string {
	randomIndex := rand.Intn(len(slice))
	return slice[randomIndex]
}

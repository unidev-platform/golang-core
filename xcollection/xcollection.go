package xcollection

import "math/rand"

// MapKeys - extract map keys as slice
func MapKeys(m map[string]bool) []string {
	keys := make([]string, 0, len(m))
	for kk := range m {
		keys = append(keys, kk)
	}
	return keys
}

// Random - fetch random element from string slice
func Random(slice []string) string {
	randomIndex := rand.Intn(len(slice))
	return slice[randomIndex]
}

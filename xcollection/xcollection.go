package xcollection

// StringBoolMapKeys - extract map keys as slice
func StringBoolMapKeys(m map[string]bool) []string {
	keys := make([]string, 0, len(m))
	for kk := range m {
		keys = append(keys, kk)
	}
	return keys
}

package xenv

import "os"

// Env - get environment variable value or default value if missing.
func Env(name string, def string) string {
	v := os.Getenv(name)
	if v == "" {
		return def
	}
	return v
}

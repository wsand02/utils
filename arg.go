package utils

import "os"

// GetArgWDefault returns a commandline argument with a default value incase it doesn't exist.
func GetArgWDefault(index int, defaultVal string) string {
	if len(os.Args) <= index {
		return defaultVal
	}
	valStr := os.Args[index]
	return valStr
}

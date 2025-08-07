package utils

import "os"

// GetArgWDefault returns a commandline argument with a default value incase it doesn't exist.
func GetArgWDefault(index int, defaultVal string) string {
	valStr := os.Args[index]
	if len(valStr) != 0 {
		return valStr
	}
	return defaultVal
}

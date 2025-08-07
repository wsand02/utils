package utils

import (
	"os"
	"strconv"
)

// GetEnvAsBool is taken from https://gist.github.com/craicoverflow/f487b83a5bae318a7a94ba6d326219c1#file-config-go
func GetEnvAsBool(name string, defaultVal bool) bool {
	valStr := os.Getenv(name)
	if val, err := strconv.ParseBool(valStr); err == nil {
		return val
	}

	return defaultVal
}

// GetEnvAsInt64 takes a environmental variable and returns it as int64
func GetEnvAsInt64(name string, defaultVal int64) int64 {
	valStr := os.Getenv(name)
	if val, err := strconv.ParseInt(valStr, 10, 64); err == nil {
		return val
	}

	return defaultVal
}

// GetEnvAsUint64 takes a environmental variable and returns it as uint64
func GetEnvAsUint64(name string, defaultVal uint64) uint64 {
	valStr := os.Getenv(name)
	if val, err := strconv.ParseUint(valStr, 10, 64); err == nil {
		return val
	}

	return defaultVal
}

// GetEnvWDefault takes a environmental variable and if it doesn't exist it returns the default value that is set.
func GetEnvWDefault(name string, defaultVal string) string {
	valStr := os.Getenv(name)
	if len(valStr) != 0 {
		return valStr
	}
	return defaultVal
}

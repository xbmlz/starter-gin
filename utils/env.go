package utils

import (
	"os"
	"strconv"
)

func GetEnvString(key, defaultVal string) string {
	val := os.Getenv(key)
	if val == "" {
		return defaultVal
	}
	return val
}

func GetEnvInteger(key string, defaultVal int) int {
	val := os.Getenv(key)
	if val == "" {
		return defaultVal
	}
	valInt, err := strconv.Atoi(val)
	if err != nil {
		return defaultVal
	}
	return valInt
}

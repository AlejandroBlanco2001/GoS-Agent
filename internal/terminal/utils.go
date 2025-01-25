package terminal

import (
	"runtime"
	"strings"
)

func RemoveOutputCommandPrefix(commandOutputPrefix []byte) string {
	cleanString := strings.TrimSpace(string(commandOutputPrefix))

	if runtime.GOOS == "windows" {
		cleanString = strings.ReplaceAll(cleanString, "\r", "")
		return cleanString // Windows based systems
	}

	return cleanString // Unix based systems
}

func BytesToMB(bytes int64) float64 {
	return float64(bytes) / 1024 / 1024
}

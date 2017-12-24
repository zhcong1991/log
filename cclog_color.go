package cclog

import (
	"fmt"
	"runtime"
)

// White returns a white string
func White(message string) string {
	if "windows" == runtime.GOOS {
		return message
	}
	return fmt.Sprintf("\x1b[37m%s\x1b[0m", message)
}

// Green returns a green string
func Green(message string) string {
	if "windows" == runtime.GOOS {
		return message
	}
	return fmt.Sprintf("\x1b[32m%s\x1b[0m", message)
}

// Red returns a red string
func Red(message string) string {
	if "windows" == runtime.GOOS {
		return message
	}
	return fmt.Sprintf("\x1b[31m%s\x1b[0m", message)
}
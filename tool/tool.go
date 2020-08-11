package tool

import (
	"McDailyAutoRun/config"
)

// GetUserIndex ...
func GetUserIndex(LineID string) int {
	for p, v := range config.User {
		if v.LineID == LineID {
			return p
		}
	}
	return -1
}

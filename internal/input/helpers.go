package input

import (
	"fmt"
	"strings"
	"time"
)

func parseTimestamp(line string) (time.Time, error) {
	parts := strings.Split(line, ",")
	if len(parts) != 2 {
		return time.Time{}, fmt.Errorf("invalid line format")
	}
	return time.Parse(time.RFC3339, parts[1])
}

package service

import (
	"fmt"
	"strings"
	"time"
)

// parseDateFromString takes a date string in the format "YYYY-MM-DD" and attempts to parse it into a time.Time object.
func (s *Service) parseDateFromString(dateStr string) (time.Time, error) {
	const layout = "2006-01-02"
	parsedDate, err := time.Parse(layout, dateStr)
	if err != nil {
		return time.Time{}, fmt.Errorf("invalid date format: %v, expected YYYY-MM-DD", err)
	}

	return parsedDate, nil
}

// FindMostActiveCookiesForDate analyzes a file to find the most active cookies for a given date.
// It requires a file path and a target date string in "YYYY-MM-DD" format.
// Returns a slice of the most active cookie(s) or an error if any step fails.
func (s *Service) FindMostActiveCookiesForDate(filepath, date string) ([]string, error) {
	targetDate, err := s.parseDateFromString(date)
	if err != nil {
		return nil, err
	}

	logsForDate, err := s.Input.FileInRecordsForDate(filepath, targetDate)
	if err != nil {
		return nil, fmt.Errorf("error_reading_file: %v", err)
	}

	cookieCounts := make(map[string]int)
	for _, log := range logsForDate {
		cookie := strings.Split(log, ",")[0]
		cookieCounts[cookie]++
	}

	var mostActiveCookies []string
	maxCount := 0
	for cookie, count := range cookieCounts {
		cookie := strings.Split(cookie, ",")[0]
		if count > maxCount {
			mostActiveCookies = []string{cookie}
			maxCount = count
		} else if count == maxCount {
			mostActiveCookies = append(mostActiveCookies, cookie)
		}
	}

	return mostActiveCookies, nil
}

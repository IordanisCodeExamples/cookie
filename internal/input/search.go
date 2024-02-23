package input

import (
	"bufio"
	"fmt"
	"os"
	"time"
)

// FileInRecordsForDate opens a file at the specified filepath and searches for records that match the targetDate.
// It returns a slice of strings containing all matching records or an error if the file cannot be opened,
// if there's an issue reading the file, or if no records are found for the given date.
func (i *Input) FileInRecordsForDate(filepath string, targetDate time.Time) ([]string, error) {
	file, err := os.Open(filepath)
	if err != nil {
		return nil, fmt.Errorf("error_opening_file: %v", err)
	}
	defer file.Close()

	return i.findRecordsForDate(file, targetDate)
}

// binarySearchFirst finds the first occurrence of the target date.
func (i *Input) binarySearchFirst(lines []string, targetDate time.Time) int {
	low, high := 0, len(lines)-1
	result := -1

	for low <= high {
		mid := low + (high-low)/2
		midTime, err := parseTimestamp(lines[mid])
		if err != nil {
			return result
		}
		if midTime.Format(dateFormat) == targetDate.Format(dateFormat) {
			result = mid
			high = mid - 1
		} else if midTime.Before(targetDate) {
			high = mid - 1
		} else {
			low = mid + 1
		}
	}
	return result
}

// binarySearchLast finds the last occurrence of the target date.
func (i *Input) binarySearchLast(lines []string, targetDate time.Time) int {
	low, high := 0, len(lines)-1
	result := -1

	for low <= high {
		mid := low + (high-low)/2
		midTime, err := parseTimestamp(lines[mid])
		if err != nil {
			return result
		}

		if midTime.Format(dateFormat) == targetDate.Format(dateFormat) {
			result = mid
			low = mid + 1
		} else if midTime.Before(targetDate) {
			high = mid - 1
		} else {
			low = mid + 1
		}
	}
	return result
}

// findRecordsForDate searches for records that match the target date within the file.
func (i *Input) findRecordsForDate(file *os.File, targetDate time.Time) ([]string, error) {
	scanner := bufio.NewScanner(file)
	var lines []string

	firstLine := true // Add a flag to track the first line
	for scanner.Scan() {
		if firstLine {
			firstLine = false // Skip the first line
			continue
		}
		line := scanner.Text()
		lines = append(lines, line)
	}

	if len(lines) < 2 {
		return nil, fmt.Errorf("input_findrecordsfordate_nodata")
	}

	low := i.binarySearchFirst(lines, targetDate)
	high := i.binarySearchLast(lines, targetDate)
	if low == -1 {
		return nil, fmt.Errorf("input_findrecordsfordate_notfound")
	}
	return lines[low : high+1], nil
}

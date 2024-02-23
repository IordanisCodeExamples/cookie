// Package service processes data from files based on dates. It includes an Input interface for reading records
// by date and a Service struct for analyzing data, such as finding the most active cookies for a given date.
// The package ensures data can be flexibly accessed and analyzed with clean separation between data retrieval
// and business logic.
package service

import (
	"errors"
	"time"
)

// Input defines an interface for retrieving records from a file based on a specific date.
type Input interface {
	// FileInRecordsForDate searches for and returns all records matching the targetDate
	// within the file specified by filepath. If no matching records are found, or an error
	// occurs during file processing, an error is returned.
	FileInRecordsForDate(filepath string, targetDate time.Time) ([]string, error)
}

// Service encapsulates the business logic of the application.
// It relies on the Input interface to access data
type Service struct {
	Input Input
}

// New creates a new instance of the Service with the provided Input implementation.
// It returns an error if the provided Input is nil, ensuring the Service is properly
// initialized with a valid data source before use.
func New(input Input) (*Service, error) {
	if input == nil {
		return nil, errors.New("input_is_nil") // Error if input is not provided.
	}
	return &Service{
		Input: input,
	}, nil
}

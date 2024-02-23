// Package input provides functionality to read and process records from files based on specific dates.
// It includes mechanisms for opening files, searching for records by date using a binary search algorithm,
// and parsing timestamps from the records.
package input

const (
	dateFormat = "2006-01-02"
)

// Input is a struct that encapsulates the functionality for processing file-based records.
type Input struct{}

// New creates and returns a new instance of the Input struct.
func New() *Input {
	return &Input{}
}

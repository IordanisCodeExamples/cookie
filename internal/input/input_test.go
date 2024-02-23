package input_test

import (
	"io"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"

	"cookie/internal/input"
)

func createTestFile(t *testing.T, content string) *os.File {
	t.Helper()
	file, err := os.CreateTemp("", "testfile")
	assert.NoError(t, err)
	_, err = file.WriteString(content)
	assert.NoError(t, err)
	_, err = file.Seek(0, io.SeekStart) // Rewind to start of the file for reading
	assert.NoError(t, err)
	return file
}

func getInput(t *testing.T) *input.Input {
	t.Helper()
	inputObj := input.New()
	return inputObj
}

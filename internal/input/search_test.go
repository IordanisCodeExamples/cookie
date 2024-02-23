package input_test

import (
	"errors"
	"os"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func content() string {
	return `cookie,timestamp
ACqkGMZMyOl5ZgFN,2021-05-23T18:36:58+00:00
BZfllfS7zGt3CPE9,2021-05-23T17:15:00+00:00
ACqkGMZMyOl5ZgFN,2021-05-23T16:49:00+00:00
YmFZn3Nd12zVz7Tk,2021-05-23T15:45:00+00:00
Dkj12Zf34lG5Zf9N,2021-05-22T14:05:00+00:00
YmFZn3Nd12zVz7Tk,2021-05-22T13:45:00+00:00
ACqkGMZMyOl5ZgFN,2021-05-22T12:17:00+00:00
BZfllfS7zGt3CPE9,2021-05-22T11:15:00+00:00
YmFZn3Nd12zVz7Tk,2021-05-21T10:45:00+00:00
Dkj12Zf34lG5Zf9N,2021-05-21T09:05:00+00:00
ACqkGMZMyOl5ZgFN,2021-05-21T08:17:00+00:00
BZfllfS7zGt3CPE9,2021-05-20T07:15:00+00:00
YmFZn3Nd12zVz7Tk,2021-05-19T06:45:00+00:00
Dkj12Zf34lG5Zf9N,2021-05-18T05:05:00+00:00
ACqkGMZMyOl5ZgFN,2021-05-17T04:17:00+00:00`
}

func secondTypeOfContent() string {
	return `cookie,timestamp
	AtY0laUfhglK3lC7,2018-12-09T14:19:00+00:00
	SAZuXPGUrfbcn5UA,2018-12-09T10:13:00+00:00
	5UAVanZf6UtGyKVS,2018-12-09T07:25:00+00:00
	AtY0laUffslK3lC7,2018-12-09T06:19:00+00:00
	4sMM2LxV07bPJzwf,2018-12-09T06:14:00+00:00
	SAZuXPGUrfbcn5UA,2018-12-08T22:03:00+00:00
	SAZuXPGUrfbcn5UA,2018-12-08T22:02:00+00:00
	4sMM2LxV07bPJzwf,2018-12-08T21:30:00+00:00
	4sMM2LxV07bPJzwf,2018-12-08T21:29:00+00:00
	fbcn5UAVanZf6UtG,2018-12-08T09:30:00+00:00
	4sMM2LxV07bPJzwf,2018-12-07T23:30:00+00:00`
}

func brokenContent() string {
	return `cookie,timestamp
ACqkGMZMyOl5ZgFN,2021-05-23T18:36:58+00:00
BZfllfS7zGt3CPE9
ACqkGMZMyOl5ZgFN,2021-05-23T16:49:00+00:00`
}

func TestFileInRecordsForDate(t *testing.T) {
	inputObj := getInput(t) // Assuming getInput initializes your input object

	type testCase struct {
		name          string
		fileName      string
		targetDate    time.Time
		file          *os.File
		expectedLen   int
		expectedError error
	}

	tests := []testCase{
		{
			name:          "Test for 2018-12-08",
			targetDate:    time.Date(2018, 12, 8, 0, 0, 0, 0, time.UTC),
			file:          createTestFile(t, secondTypeOfContent()),
			fileName:      createTestFile(t, secondTypeOfContent()).Name(),
			expectedLen:   5,
			expectedError: nil,
		},
		{
			name:          "Test for 2021-05-22",
			targetDate:    time.Date(2021, 05, 22, 0, 0, 0, 0, time.UTC),
			file:          createTestFile(t, content()),
			fileName:      createTestFile(t, content()).Name(),
			expectedLen:   4,
			expectedError: nil,
		},
		{
			name:          "Test for 2021-05-23",
			targetDate:    time.Date(2021, 05, 23, 0, 0, 0, 0, time.UTC),
			file:          createTestFile(t, content()),
			fileName:      createTestFile(t, content()).Name(),
			expectedLen:   4,
			expectedError: nil,
		},
		{
			name:          "Test for 1 record",
			targetDate:    time.Date(2021, 05, 19, 0, 0, 0, 0, time.UTC),
			file:          createTestFile(t, content()),
			fileName:      createTestFile(t, content()).Name(),
			expectedLen:   1,
			expectedError: nil,
		},
		{
			name:          "Test for another date",
			targetDate:    time.Date(2021, 05, 23, 0, 0, 0, 0, time.UTC),
			file:          createTestFile(t, ""),
			fileName:      createTestFile(t, "").Name(),
			expectedLen:   0,
			expectedError: errors.New("input_findrecordsfordate_nodata"),
		},
		{
			name:       "Test no dates found",
			targetDate: time.Date(2020, 05, 24, 0, 0, 0, 0, time.UTC),
			file:       createTestFile(t, content()),
			fileName:   createTestFile(t, content()).Name(),

			expectedLen:   0,
			expectedError: errors.New("input_findrecordsfordate_notfound"),
		},
		{
			name:          "Test for file open error",
			targetDate:    time.Date(2021, 05, 23, 0, 0, 0, 0, time.UTC),
			file:          createTestFile(t, content()),
			fileName:      "non_existent_file",
			expectedLen:   0,
			expectedError: errors.New("error_opening_file: open non_existent_file: no such file or directory"),
		},
		{
			name:          "Test for file read error",
			targetDate:    time.Date(2021, 05, 23, 0, 0, 0, 0, time.UTC),
			file:          createTestFile(t, brokenContent()),
			fileName:      createTestFile(t, brokenContent()).Name(),
			expectedLen:   0,
			expectedError: errors.New("input_findrecordsfordate_notfound"),
		},
	}

	// Iterate over the test cases
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			records, err := inputObj.FileInRecordsForDate(tc.fileName, tc.targetDate)
			assert.Equal(t, tc.expectedError, err)
			assert.Len(t, records, tc.expectedLen)
		})
	}
}

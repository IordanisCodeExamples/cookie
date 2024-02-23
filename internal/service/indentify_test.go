package service_test

import (
	"path/filepath"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestFindMostActiveCookiesForDate(t *testing.T) {
	date := "2021-05-22"
	filepath := filepath.Join("testdata", "cookie_log.csv")
	const layout = "2006-01-02"
	t.Run("Success with one most common cookie", func(t *testing.T) {
		mocks := getMocks(t)
		parsedDate, err := time.Parse(layout, date)
		assert.NoError(t, err)
		mocks.input.EXPECT().
			FileInRecordsForDate(filepath, parsedDate).
			Return(
				[]string{
					"ACqkGMZMyOl5ZgFN,2021-05-23T12:17:00+00:00",
					"BZfllfS7zGt3CPE9,2021-05-23T11:15:00+00:00",
					"ACqkGMZMyOl5ZgFN,2021-05-23T10:49:00+00:00",
					"YmFZn3Nd12zVz7Tk.2021-05-23T09:45:00+00:00"},
				nil,
			)

		srv := getService(t, mocks)
		cookies, err := srv.FindMostActiveCookiesForDate(filepath, date)
		assert.NoError(t, err)
		assert.Equal(t, 1, len(cookies))
		assert.Equal(t, "ACqkGMZMyOl5ZgFN", cookies[0])
	})

	t.Run("Success with multiple most common cookies", func(t *testing.T) {
		mocks := getMocks(t)
		parsedDate, err := time.Parse(layout, date)
		assert.NoError(t, err)
		mocks.input.EXPECT().
			FileInRecordsForDate(filepath, parsedDate).
			Return(
				[]string{
					"ACqkGMZMyOl5ZgFN,2021-05-23T12:17:00+00:00",
					"BZfllfS7zGt3CPE9,2021-05-23T11:15:00+00:00",
					"ACqkGMZMyOl5ZgFN,2021-05-23T10:49:00+00:00",
					"YmFZn3Nd12zVz7Tk.2021-05-23T09:45:00+00:00",
					"BZfllfS7zGt3CPE9,2021-05-23T11:15:00+00:00",
				},
				nil,
			)

		srv := getService(t, mocks)
		cookies, err := srv.FindMostActiveCookiesForDate(filepath, date)
		assert.NoError(t, err)
		assert.Equal(t, 2, len(cookies))
		assert.Equal(t, "ACqkGMZMyOl5ZgFN", cookies[0])
		assert.Equal(t, "BZfllfS7zGt3CPE9", cookies[1])
	})

	t.Run("Error when parsing date", func(t *testing.T) {
		mocks := getMocks(t)
		mocks.input.EXPECT().
			FileInRecordsForDate(filepath, time.Time{}).
			Times(0)

		srv := getService(t, mocks)
		cookies, err := srv.FindMostActiveCookiesForDate(filepath, "2021-05-22-")
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "invalid date format")
		assert.Nil(t, cookies)
	})

	t.Run("Error when parsing date because of layout", func(t *testing.T) {
		mocks := getMocks(t)
		mocks.input.EXPECT().
			FileInRecordsForDate(filepath, time.Time{}).
			Times(0)

		srv := getService(t, mocks)
		cookies, err := srv.FindMostActiveCookiesForDate(filepath, "2021-22-05")
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "invalid date format")
		assert.Nil(t, cookies)
	})

	t.Run("Error on file in record date", func(t *testing.T) {
		mocks := getMocks(t)
		parsedDate, err := time.Parse(layout, date)
		assert.NoError(t, err)
		mocks.input.EXPECT().
			FileInRecordsForDate(filepath, parsedDate).
			Return(nil, assert.AnError)

		srv := getService(t, mocks)

		cookies, err := srv.FindMostActiveCookiesForDate(filepath, date)
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "error_reading_file")
		assert.Nil(t, cookies)
	})
}

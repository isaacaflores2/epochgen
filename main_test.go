package main

import (
	"errors"
	"testing"
	"time"
)

func Test_addDuration(t *testing.T) {
	testCases := []struct {
		name      string
		epochTime int64
		duration  time.Duration
		expected  int64
	}{
		{
			name:      "Sub 30 minutes",
			epochTime: 1659114000000, // 27JUL2022 10am
			duration:  -30 * time.Minute,
		},
		{
			name:      "Sub 10 minutes",
			epochTime: 1659114000000, // 27JUL2022 10am
			duration:  -10 * time.Minute,
		},
		{
			name:      "Sub 7 minutes",
			epochTime: 1659114000000, // 27JUL2022 10am
			duration:  -7 * time.Minute,
		},
		{
			name:      "Sub 5 minutes",
			epochTime: 1659114000000, // 27JUL2022 10am
			duration:  -5 * time.Minute,
		},
		{
			name:      "Add 1 hour and 30 minutes",
			epochTime: 1659114000000, // 27JUL2022 10am
			duration:  90 * time.Minute,
		},
		{
			name:      "Add 2 hours and 5 minutes",
			epochTime: 1659114000000, // 27JUL2022 10am
			duration:  125 * time.Minute,
		},
		{
			name:      "Add 2 hours and 10 minutes",
			epochTime: 1659114000000, // 27JUL2022 10am
			duration:  130 * time.Minute,
		},
		{
			name:      "Add 3600 seconds",
			epochTime: 1664528400000,
			duration:  3600 * time.Second,
		},
	}
	for _, td := range testCases {
		t.Run(td.name, func(t *testing.T) {
			updatedEpochTime := addDuration(td.epochTime, td.duration)
			t.Log("Updated epoch time:", updatedEpochTime, "->", time.UnixMilli(updatedEpochTime).Local())

			if updatedEpochTime == 0 {
				t.Errorf("The updated time should not be zero")
			}
		})
	}
}

func Test_parseTime(t *testing.T) {
	testCases := []struct {
		name        string
		input       string
		format      string
		expectedErr error
	}{
		{
			name:        "Parse valid RFC3339 time",
			input:       "2022-08-18T00:00:00.000Z",
			format:      time.RFC3339,
			expectedErr: nil,
		},
		{
			name:        "Parse invalid RFC3339 time",
			input:       "2022-08-",
			format:      time.RFC3339,
			expectedErr: errParseTime,
		},
		{
			name:        "Parse empty time",
			input:       "",
			format:      time.RFC3339,
			expectedErr: errParseTime,
		},
	}
	for _, td := range testCases {
		t.Run(td.name, func(t *testing.T) {
			parsedTime, err := parseTime(td.format, td.input)
			if !errors.Is(err, td.expectedErr) {
				t.Errorf("parseTime() expected error: %s, instead found %s", td.expectedErr, err)
			}
			if td.expectedErr == nil && parsedTime.IsZero() {
				t.Errorf("parseTime() parsedTime should not be zero")
			}
		})
	}
}

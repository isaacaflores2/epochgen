package main

import (
	"errors"
	"flag"
	"fmt"
	"time"
)

var (
	errParseTime = errors.New("failed to parse time")
	timeStr      string
	format       string
	epochTime    int64
	duration     string
)

func init() {
	flag.StringVar(&timeStr, "t", "", "time string to parse, example:  2022-08-18T00:00:00.000Z")
	flag.StringVar(&format, "f", time.RFC3339, "format for time string parse.")
	flag.Int64Var(&epochTime, "e", 0, "epoch time in milliseconds")
	flag.StringVar(&duration, "d", "", "duration to add to epoch time. Example: \"300ms\", \"-1.5h\" or \"2h45m\"")
}

func main() {
	flag.Parse()

	// parse time string
	if len(timeStr) != 0 {
		parsedTime, err := parseTime(format, timeStr)
		if err != nil {
			fmt.Printf("❌ failed to parse time: %s using format: %s. Error: %s \n", timeStr, format, err)
			return
		}

		fmt.Printf("Successfully Parsed time: %s  ⏱ \n - local time: %s \n - epoch in milliseconds: %d \n", timeStr, toLocalTime(parsedTime), parsedTime.UnixMilli())
		return
	}

	// add duration to epoch time
	if epochTime != 0 && len(duration) != 0 {
		parsedDuration, err := time.ParseDuration(duration)
		if err != nil {
			fmt.Printf("❌ failed to parse duration: %s. Error: %s .\nDuration must use one of these valid time units: \"ns\", \"us\" (or \"µs\"), \"ms\", \"s\", \"m\", \"h\" \n", timeStr, err)
			return
		}
		newEpochTime := addDuration(epochTime, parsedDuration)
		fmt.Printf("Successfully added %s to epoch time: %d ⏱ \n - New time: %d \n", parsedDuration, epochTime, newEpochTime)
		return
	}

	defaultEpochGen()
}

func defaultEpochGen() {
	currEpochMs := time.Now().UnixMilli()
	fourtyEightHoursFromNowEpochMs := time.Now().Add(48 * time.Hour).UnixMilli()
	twoWeeksFromNowEpochMs := time.Now().AddDate(0, 0, 14).UnixMilli()
	fmt.Printf("All epoch times are milliseconds ⏱ \n - current epoch: %d \n - epoch 48 hours from now: %d \n - epoch 14 days from now: %d \n",
		currEpochMs, fourtyEightHoursFromNowEpochMs, twoWeeksFromNowEpochMs)
}

func addDuration(epochTime int64, duration time.Duration) int64 {
	t := time.UnixMilli(epochTime)
	return t.Add(duration).UnixMilli()
}

func parseTime(format, timeStr string) (time.Time, error) {
	if len(format) == 0 {
		format = time.RFC3339
	}

	parsedTime, err := time.Parse(format, timeStr)
	if err != nil {
		return time.Time{}, fmt.Errorf("%w : %s", errParseTime, err)
	}
	return parsedTime, nil
}

func toLocalTime(t time.Time) string {
	loc, err := time.LoadLocation("America/Los_Angeles")
	if err != nil {
		// could not load PST, fallback to Local
		return t.Local().Format(time.RFC822)
	}

	return t.In(loc).Format(time.RFC822)
}

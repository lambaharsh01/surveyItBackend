package utils

import "time"

func TimeDiffMinuet(before time.Time, after time.Time) int {
	return int(after.Sub(before).Minutes())
}

func ParseDate(dateStr string, timePointer *time.Time) error {
	parsedTime, err := time.Parse("2006-01-02", dateStr)
	if err != nil {
		return err
	}
	*timePointer = parsedTime
	return nil
}

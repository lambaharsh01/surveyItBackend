package utils

import "time"

func TimeDiffMinuet(before time.Time, after time.Time) int {
	return int(after.Sub(before).Minutes())
}
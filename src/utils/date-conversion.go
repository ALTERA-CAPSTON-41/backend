package utils

import "time"

func ConvertDate(time time.Time) string {
	return time.Format("2006-01-02")
}

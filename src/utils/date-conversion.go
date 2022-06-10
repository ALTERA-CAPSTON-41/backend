package utils

import "time"

func ConvertDateToString(time time.Time) string {
	return time.Format("2006-01-02")
}

func ConvertStringToDate(date string) time.Time {
	t, _ := time.Parse("2006-01-02", date)
	return t
}

func ConvertStringToDatetime(datetime string) time.Time {
	t, _ := time.Parse("2006-01-02 15:04:05", datetime)
	return t
}
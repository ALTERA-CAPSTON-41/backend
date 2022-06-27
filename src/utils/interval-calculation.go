package utils

import "time"

func CountIntervalByYearRoundDown(origin time.Time, target time.Time) (years int) {
	if target.Before(origin) {
		return 0
	}

	target = target.In(origin.Location())

	targetYear, targetMonth, targetDate := target.Date()
	target = time.Date(targetYear, targetMonth, targetDate, 0, 0, 0, 0, time.UTC)

	originYear, originMonth, originDate := origin.Date()
	origin = time.Date(originYear, originMonth, originDate, 0, 0, 0, 0, time.UTC)

	yearsPassed := targetYear - originYear
	anniversary := origin.AddDate(yearsPassed, 0, 0)
	if anniversary.After(target) {
		yearsPassed--
	}
	return yearsPassed
}

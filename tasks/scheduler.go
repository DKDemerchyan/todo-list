package tasks

import (
	"errors"
	"strconv"
	"strings"
	"time"
)

const DateFormat = "20060102"

func NextDate(now time.Time, date string, repeat string) (string, error) {
	now = now.Truncate(24 * time.Hour) // Избавляюсь от точного времени для корректного сравнения дат
	taskStartDate, err := time.Parse(DateFormat, date)
	if err != nil {
		return "", err
	}

	switch {
	case strings.HasPrefix(repeat, "d "):
		daysToAdd, err := strconv.Atoi(strings.TrimPrefix(repeat, "d "))
		if err != nil || daysToAdd < 1 || daysToAdd > 400 {
			return "", err
		}

		nextDate := taskStartDate
		for nextDate.Before(now) && !nextDate.Equal(now) {
			nextDate = nextDate.AddDate(0, 0, daysToAdd)
		}
		return nextDate.Format(DateFormat), nil

	case repeat == "y":
		nextDate := taskStartDate.AddDate(1, 0, 0)
		for nextDate.Before(now) {
			nextDate = nextDate.AddDate(1, 0, 0)
		}
		return nextDate.Format(DateFormat), nil

	default:
		return "", errors.New("invalid repeat pattern")
	}
}

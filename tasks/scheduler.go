package tasks

import (
	"errors"
	"strconv"
	"strings"
	"time"
)

const DateFormat = "20060102"

func NextDate(now time.Time, date string, repeat string) (string, error) {
	taskStartDate, err := time.Parse(DateFormat, date)
	if err != nil {
		return "", err
	}

	switch {
	case strings.HasPrefix(repeat, "d "):
		daysToAdd, err := strconv.Atoi(strings.TrimPrefix(repeat, "d "))
		if err != nil || daysToAdd < 1 || daysToAdd > 400 {
			return "", errors.New("invalid date number of days to add")
		}

		nextDate := taskStartDate.AddDate(0, 0, daysToAdd)
		for nextDate.Before(now) {
			nextDate = nextDate.AddDate(0, 0, daysToAdd)
		}
		return nextDate.Format(DateFormat), nil

	case strings.HasPrefix(repeat, "w "):
		weekdays := strings.Split(strings.TrimPrefix(repeat, "w "), ",")
		var validWeekdays []int

		for _, strDay := range weekdays {
			if strDay == "" {
				continue
			}
			day, err := strconv.Atoi(strDay)
			if err != nil || day < 0 || day > 7 {
				return "", errors.New("invalid weekday number in " + repeat)
			}
			validWeekdays = append(validWeekdays, day)
		}

		var nextDate time.Time
		if taskStartDate.Before(now) {
			nextDate = now
		} else {
			nextDate = taskStartDate
		}
		for {
			nextDate = nextDate.AddDate(0, 0, 1)
			for _, day := range validWeekdays {
				if int(nextDate.Weekday()) == (day % 7) {
					return nextDate.Format(DateFormat), nil
				}
			}
		}

	case strings.HasPrefix(repeat, "m "):
		parts := strings.Split(strings.TrimPrefix(repeat, "m "), " ")
		if len(parts) != 1 && len(parts) != 2 {
			return "", errors.New("invalid repeat pattern")
		}

		var validDays []int
		var validMonths []int

		days := strings.Split(parts[0], ",")
		for _, d := range days {
			day, err := strconv.Atoi(d)
			if err != nil || day < -2 || day > 31 {
				return "", errors.New("invalid day number in " + repeat)
			}
			validDays = append(validDays, day)
		}

		if len(parts) == 2 {
			months := strings.Split(parts[1], ",")
			for _, m := range months {
				month, err := strconv.Atoi(m)
				if err != nil || month < 1 || month > 12 {
					return "", errors.New("invalid month number in " + repeat)
				}
				validMonths = append(validMonths, month)
			}
		}

		var nextDate time.Time
		if taskStartDate.Before(now) {
			nextDate = now
		} else {
			nextDate = taskStartDate
		}

		for {
			nextDate = nextDate.AddDate(0, 0, 1)
			for _, day := range validDays {
				var targetDay int
				if day == -1 {
					targetDay = getLastDayOfMonth(nextDate)
				} else if day == -2 {
					targetDay = getLastDayOfMonth(nextDate) - 1
				} else {
					targetDay = day
				}

				if nextDate.Day() == targetDay {
					if len(validMonths) == 0 {
						return nextDate.Format(DateFormat), nil
					}
					for _, month := range validMonths {
						if int(nextDate.Month()) == month {
							return nextDate.Format(DateFormat), nil
						}
					}
				}
			}
		}

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

func getLastDayOfMonth(date time.Time) int {
	year, month, _ := date.Date()
	return time.Date(year, month+1, 0, 0, 0, 0, 0, date.Location()).Day()
}

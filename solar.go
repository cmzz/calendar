package calendar

import (
	"errors"
	"time"
)

type Solar struct {
}

func Calendar(start time.Time, end time.Time) ([]*Year, error) {
	if start.After(end) {
		return nil, errors.New("start time must be before end time")
	}
	if end.Sub(start) < time.Hour*24*365 {
		return nil, errors.New("calendar must be at least one year")
	}

	var years []*Year
	for {
		if end.After(start) {
			break
		}

		years = append(years, NewYear(start))
		start = start.AddDate(1, 0, 0)
	}

	return years, nil
}

func MonthCalendar(t time.Time) *Month {
	return NewMonth(t)
}

func WeekCalendar(t time.Time) *Week {
	return NewWeek(t, false)
}

func YearCalendar(t time.Time) *Year {
	return NewYear(t)
}

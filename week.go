package calendar

import (
	"time"
)

type Week struct {
	date        time.Time
	isoWeekday  bool
	year        int
	weekOfYear  int
	weekOfMonth int
	weekdays    [7]*WeekDay
	month       int
	dayFormat   string
	dateFormat  string
}

func (w *Week) SetDateFormat(format string) {
	w.dateFormat = format
}

func (w *Week) SetDayFormat(format string) {
	w.dayFormat = format
}

func (w *Week) Start() time.Time {
	for _, weekday := range w.weekdays {
		if weekday != nil {
			return weekday.day.Start()
		}
	}

	return time.Time{}
}
func (w *Week) End() time.Time {
	return w.weekdays[6].day.End()
}

func (w *Week) WeekOfYear() int {
	_, week := w.date.ISOWeek()
	return week
}

func (w *Week) Year() *Year {
	return w.weekdays[0].day.Year()
}

func (w *Week) NYear() int {
	for _, weekday := range w.weekdays {
		if weekday != nil {
			return weekday.day.NYear()
		}
	}

	return 0
}

func (w *Week) Print() {
	header := []string{"Sun", "Mon", "Tue", "Wed", "Thu", "Fri", "Sat"}
	var row []string
	for _, wd := range w.weekdays {
		if wd != nil {
			row = append(row, wd.day.String())
		} else {
			row = append(row, "")
		}
	}
	PrintTable(header, [][]string{row}, nil)
}

func (w *Week) Month() *Month {
	return w.weekdays[0].day.Month()
}

func (w *Week) Months() []*Month {
	var months []*Month
	m := &Month{}

	for _, wd := range w.weekdays {
		if wd != nil {
			if wd.day != nil {
				if m == nil {
					m = wd.day.Month()
					months = append(months, wd.day.Month())
				}

				if m != wd.day.Month() {
					m = wd.day.Month()
					months = append(months, wd.day.Month())
				}
			}
		}
	}
	return months
}

func (w *Week) InitWeekdays(acrossMonth bool) {
	var startDate time.Time
	var dayOfWeek int

	w.weekdays = [7]*WeekDay{}

	dayOfWeek = int(w.date.Weekday())
	startDate = w.date.AddDate(0, 0, -(dayOfWeek))

	for i := 0; i < 7; i++ {
		if acrossMonth {
			d := startDate.AddDate(0, 0, i)

			if d.Month() == w.date.Month() {
				_day := NewDay(d)
				_day.SetDayFormat(w.dayFormat)
				_day.SetDateFormat(w.dateFormat)
				w.weekdays[i] = &WeekDay{
					dayOfWeek: i,
					day:       _day,
				}
			}
		} else {
			_day := NewDay(startDate.AddDate(0, 0, i))
			_day.SetDayFormat(w.dayFormat)
			_day.SetDateFormat(w.dateFormat)
			w.weekdays[i] = &WeekDay{
				dayOfWeek: i,
				day:       _day,
			}
		}
	}
}

func NewWeek(t time.Time, acrossMonth bool) *Week {
	w := &Week{}

	w.date = t
	w.weekOfMonth = int(t.Day())
	w.month = int(t.Month())
	w.year = int(t.Year())

	w.InitWeekdays(acrossMonth)

	return w
}

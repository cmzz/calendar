package calendar

import "time"

type Year struct {
	year       int
	months     []*Month
	days       []*Day
	dateFormat string
	dayFormat  string
	startTime  time.Time
	endTime    time.Time
	daysInYear int
}

func (y *Year) SetDateFormat(format string) {
	y.dateFormat = format
}

func (y *Year) SetDayFormat(format string) {
	y.dayFormat = format
}

func (y *Year) Days() []*Day {
	return y.days
}

func (y *Year) DaysInYear() int {
	return y.daysInYear
}

func (y *Year) Start() *Day {
	return y.days[0]
}

func (y *Year) End() *Day {
	return y.days[len(y.days)-1]
}

func (y *Year) InitDays() {
	y.days = make([]*Day, 0)
	
	for i := 0; i < len(y.months); i++ {
		for j := 0; j < y.months[i].DaysInMonth(); j++ {
			d := NewDay(y.months[i].days[j].Date())
			d.SetDayFormat(y.dayFormat)
			d.SetDateFormat(y.dateFormat)
			y.days = append(y.days, d)
		}
	}
}

func (y *Year) InitMonth() {
	y.months = make([]*Month, 0)

	for i := 1; i < 13; i++ {
		m := NewMonth(time.Date(y.year, time.Month(i), 1, 0, 0, 0, 0, time.Local))
		m.SetDateFormat(y.dateFormat)
		m.SetDayFormat(y.dayFormat)
		m.InitDays()
		m.InitWeeks()

		y.months = append(y.months, m)
	}
}

func NewYear(t time.Time) *Year {
	year := &Year{}
	year.year = t.Year()
	year.startTime = time.Date(year.year, 1, 1, 0, 0, 0, 0, time.Local)
	year.endTime = time.Date(year.year, 12, 31, 23, 59, 59, 0, time.Local)

	year.InitDays()
	year.InitMonth()

	return year
}

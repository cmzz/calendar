package calendar

import "time"

type Day struct {
	year       int
	month      int
	day        int
	dayOfWeek  int
	dayOfMonth int
	dayOfYear  int
	date       time.Time
	dateFormat string
	dayFormat  string
	start      time.Time
	end        time.Time
}

func (d *Day) SetDayFormat(format string) {
	if format != "" {
		d.dayFormat = format
	}
}

func (d *Day) SetDateFormat(format string) {
	if format != "" {
		d.dateFormat = format
	}
}

func (d *Day) String() string {
	return d.date.Format(d.dayFormat)
}

func (d *Day) DateString() string {
	return d.date.Format(d.dateFormat)
}

func (d *Day) Date() time.Time {
	return d.date
}

func (d *Day) Start() time.Time {
	return d.start
}

func (d *Day) WeekDay() int {
	return int(d.date.Weekday())
}

func (d *Day) End() time.Time {
	return d.end
}

func (d *Day) DayOfYear() int {
	return d.dayOfYear
}

func (d *Day) NYear() int {
	return d.year
}

func (d *Day) NMonth() int {
	return d.month
}

func (d *Day) DayOfMonth() int {
	return d.dayOfMonth
}

func (d *Day) DayOfWeek() int {
	return int(d.date.Weekday())
}

func (d *Day) Year() *Year {
	return NewYear(d.date)
}

func (d *Day) Month() *Month {
	return NewMonth(d.date)
}

func (d *Day) Week() *Week {
	return NewWeek(d.date, false)
}

func (d *Day) init() {
	t := d.date
	d.year = t.Year()
	d.month = int(t.Month())
	d.dayOfYear = t.YearDay()
	d.dayOfMonth = t.Day()
	d.dayOfWeek = int(t.Weekday())

	year, month, _day := t.Date()
	d.start = time.Date(year, month, _day, 0, 0, 0, 0, t.Location())
	d.end = time.Date(year, month, _day, 23, 59, 59, 0, t.Location())
}

func NewDay(t time.Time) *Day {
	day := &Day{}
	day.dateFormat = "2006-01-02"
	day.dayFormat = "2"
	day.date = t

	day.init()

	return day
}

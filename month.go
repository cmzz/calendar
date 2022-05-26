package calendar

import (
	"time"
)

type Month struct {
	date         time.Time
	year         int
	monthOfYear  int
	month        time.Month
	days         []*Day
	weeks        []*Week
	daysInMonth  int
	startOfMonth time.Time
	endOfMonth   time.Time
	dayFormat    string
	dateFormat   string
}

func (m *Month) SetDateFormat(format string) {
	m.dateFormat = format
}

func (m *Month) SetDayFormat(format string) {
	m.dayFormat = format
}

func (m *Month) String() string {
	return m.date.Month().String()
}

func (m *Month) End() *Day {
	return m.days[len(m.days)-1]
}

func (m *Month) Start() *Day {
	return m.days[0]
}

func (m *Month) DaysInMonth() int {
	return m.daysInMonth
}

func (m *Month) InitDays() {
	m.days = make([]*Day, 0)
	for i := 0; i < m.DaysInMonth(); i++ {
		d := NewDay(m.startOfMonth.AddDate(0, 0, i))
		d.SetDayFormat(m.dayFormat)
		d.SetDateFormat(m.dateFormat)
		m.days = append(m.days, d)
	}
}

func (m *Month) InitWeeks() {
	var lw *Week
	m.weeks = make([]*Week, 0)

	for i := 0; i < m.DaysInMonth(); i++ {
		w := NewWeek(m.days[i].Date(), true)
		w.SetDayFormat(m.dayFormat)
		w.SetDateFormat(m.dateFormat)

		if lw == nil {
			m.weeks = append(m.weeks, w)
			lw = w
		} else {
			if lw.Start() != w.Start() {
				m.weeks = append(m.weeks, w)
				lw = w
			}
		}
	}
}

func NewMonth(t time.Time) *Month {
	m := &Month{}

	m.date = t
	m.year = t.Year()
	m.monthOfYear = int(t.Month())
	m.month = t.Month()
	m.weeks = []*Week{}

	m.startOfMonth = time.Date(m.year, m.month, 1, 0, 0, 0, 0, m.date.Location())
	m.endOfMonth = m.startOfMonth.AddDate(0, 1, -1)
	m.daysInMonth = m.endOfMonth.Day()

	m.InitDays()
	m.InitWeeks()

	return m
}

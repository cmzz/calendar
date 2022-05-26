package calendar

type WeekDay struct {
	dayOfWeek int
	day       *Day
}

func (wd *WeekDay) Day() *Day {
	return wd.day
}

func (wd *WeekDay) Name() string {
	return wd.day.Date().Weekday().String()
}

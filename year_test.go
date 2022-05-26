package calendar

import (
	"fmt"
	"testing"
	"time"
)

func (y *Year) Print() {
	for i := 0; i < len(y.months); i++ {
		fmt.Println(y.months[i].String())
		y.months[i].Print()
	}
}

func TestYearPrint(t *testing.T) {
	year := NewYear(time.Date(2018, time.January, 1, 0, 0, 0, 0, time.UTC))
	year.SetDayFormat("02")
	year.SetDateFormat("2006-01-02")

	year.InitDays()
	year.InitMonth()

	fmt.Println(fmt.Sprintf("%d 年年历", year.Start().NYear()))
	year.Print()
}

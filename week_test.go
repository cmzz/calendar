package calendar

import (
	"fmt"
	"testing"
	"time"
)

func TestPrint(t *testing.T) {
	week := NewWeek(time.Date(2014, time.November, 1, 0, 0, 0, 0, time.UTC), true)
	week.SetDayFormat("02")

	months := week.Months()
	monthStr := ""

	if len(months) > 1 {
		monthStr = fmt.Sprintf("%s - %s", months[0].String(), months[len(months)-1].String())
	} else {
		monthStr = months[0].String()
	}

	fmt.Println(fmt.Sprintf("%d 年第 %d 周周历 [%s]", week.NYear(), week.WeekOfYear(), monthStr))
	week.Print()
}

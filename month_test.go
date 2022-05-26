package calendar

import (
	"fmt"
	"testing"
	"time"
)

func (m *Month) Print() {
	header := []string{"Sun", "Mon", "Tue", "Wed", "Thu", "Fri", "Sat"}
	var rows [][]string
	for _, week := range m.weeks {
		row := make([]string, 0)
		for _, wd := range week.weekdays {
			if wd != nil {
				row = append(row, wd.day.String())
			} else {
				row = append(row, "")
			}
		}

		rows = append(rows, row)
	}

	PrintTable(header, rows, nil)
}

func TestMonthPrint(t *testing.T) {
	month := NewMonth(time.Date(2018, time.April, 1, 0, 0, 0, 0, time.UTC))
	month.SetDayFormat("02")

	fmt.Println(fmt.Sprintf("%d 年第 %d 月月历", month.Start().NYear(), month.Start().NMonth()))
	month.Print()

	month = NewMonth(time.Date(2022, time.May, 26, 0, 0, 0, 0, time.UTC))
	month.SetDayFormat("02")

	fmt.Println(fmt.Sprintf("%d 年第 %d 月月历", month.Start().NYear(), month.Start().NMonth()))
	month.Print()
}

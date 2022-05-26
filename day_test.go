package calendar

import (
	"testing"
	"time"
)

var td = time.Date(2015, time.April, 19, 0, 0, 0, 0, time.UTC)
var day = NewDay(td)

func TestDateString(t *testing.T) {
	if day.DateString() != "2015-04-19" {
		t.Errorf("Expected 2015-04-19, got %s", day.DateString())
	}

	day.SetDateFormat("2006/01/02")
	if day.DateString() != "2015/04/19" {
		t.Errorf("Expected 2015/04/19, got %s", day.DateString())
	}
}

func TestString(t *testing.T) {
	if day.String() != "19" {
		t.Errorf("Expected 19, got %s", day.String())
	}

	day.SetDayFormat("02")
	if day.String() != "19" {
		t.Errorf("Expected 19, got %s", day.String())
	}
}

func TestDayOfMonth(t *testing.T) {
	if day.DayOfMonth() != 19 {
		t.Errorf("Expected 19, got %d", day.DayOfMonth())
	}
}

func TestDayOfYear(t *testing.T) {
	if day.DayOfYear() != 109 {
		t.Errorf("Expected 109, got %d", day.DayOfYear())
	}
}

func TestDayOfWeek(t *testing.T) {
	if day.DayOfWeek() != 0 {
		t.Errorf("Expected 0, got %d", day.DayOfWeek())
	}
}

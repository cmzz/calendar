# Calendar

The Calendar package is implemented based on the Go time package and aims to increase the calendar functions of the time package. 

The Calendar package can generate weekly, monthly, and annual calendars based on any point in time.

### Install

```shell
go get -u github.com/cmzz/calendar
```

### Example

#### Annual calendar

```go   
year := NewYear(time.Date(2018, time.January, 1, 0, 0, 0, 0, time.UTC))
year.SetDayFormat("02")
year.SetDateFormat("2006-01-02")

year.InitDays()
year.InitMonth()

fmt.Println(fmt.Sprintf("%d 年年历", year.Start().NYear()))
year.Print()
```

Output:

```shell
2018 年年历
January
+-----+-----+-----+-----+-----+-----+-----+
| SUN | MON | TUE | WED | THU | FRI | SAT |
+-----+-----+-----+-----+-----+-----+-----+
|     | 1   | 2   | 3   | 4   | 5   | 6   |
| 7   | 8   | 9   | 10  | 11  | 12  | 13  |
| 14  | 15  | 16  | 17  | 18  | 19  | 20  |
| 21  | 22  | 23  | 24  | 25  | 26  | 27  |
| 28  | 29  | 30  | 31  |     |     |     |
+-----+-----+-----+-----+-----+-----+-----+
February
+-----+-----+-----+-----+-----+-----+-----+
| SUN | MON | TUE | WED | THU | FRI | SAT |
+-----+-----+-----+-----+-----+-----+-----+
|     |     |     |     | 1   | 2   | 3   |
| 4   | 5   | 6   | 7   | 8   | 9   | 10  |
| 11  | 12  | 13  | 14  | 15  | 16  | 17  |
| 18  | 19  | 20  | 21  | 22  | 23  | 24  |
| 25  | 26  | 27  | 28  |     |     |     |
+-----+-----+-----+-----+-----+-----+-----+

...
```

#### Monthly calendar

```go
month := NewMonth(time.Date(2018, time.April, 1, 0, 0, 0, 0, time.UTC))
month.SetDayFormat("02")

fmt.Println(fmt.Sprintf("%d 年第 %d 月月历", month.Start().NYear(), month.Start().NMonth()))
month.Print()
```

Output:

```shell
2018 年第 4 月月历
+-----+-----+-----+-----+-----+-----+-----+
| SUN | MON | TUE | WED | THU | FRI | SAT |
+-----+-----+-----+-----+-----+-----+-----+
| 1   | 2   | 3   | 4   | 5   | 6   | 7   |
| 8   | 9   | 10  | 11  | 12  | 13  | 14  |
| 15  | 16  | 17  | 18  | 19  | 20  | 21  |
| 22  | 23  | 24  | 25  | 26  | 27  | 28  |
| 29  | 30  |     |     |     |     |     |
+-----+-----+-----+-----+-----+-----+-----+
```

#### Weekly calendar

```go
week := NewWeek(time.Date(2014, time.November, 1, 0, 0, 0, 0, time.UTC), true)
week.SetDayFormat("02")

months := week.Months()
monthStr := ""

fmt.Println(fmt.Sprintf("%d 年第 %d 周周历", week.NYear(), week.WeekOfYear()))
week.Print()
```

Output:

```shell
2014 年第 44 周周历 [November]
+-----+-----+-----+-----+-----+-----+-----+
| SUN | MON | TUE | WED | THU | FRI | SAT |
+-----+-----+-----+-----+-----+-----+-----+
|     |     |     |     |     |     | 1   |
+-----+-----+-----+-----+-----+-----+-----+
```
package main

import (
	"fmt"
	"math"
	"strings"
	"time"
	"flag"

	"github.com/jinzhu/now"
	. "github.com/logrusorgru/aurora/v3"
)

const (
	blank = "  "
)

func main() {
	nextMonth := flag.Bool("n", false, "Display this month and next month")
	flag.Parse()

	s := ""
	s += getHeader()
	today := time.Now()
	s += getDates(today, true)
	if *nextMonth {
		nextMonth := now.EndOfMonth().Add(time.Hour * 24)
		s += getDates(nextMonth, false)
	}
	fmt.Println(s)
}

func getHeader() string {
	year, month, _ := time.Now().Date()
	s := ""
	headerLen := 20
	leftSpaces := int(math.Floor(float64((headerLen - 4 - len(month.String())) / 2)))
	for i := 0; i < leftSpaces; i++ {
		s += " "
	}

	s += fmt.Sprintf("%v %v\n", month, year)
	s += "Su Mo Tu We Th Fr Sa\n"
	return s
}

func getDates(date time.Time, highlightDt bool) string {
	firstOfMonth := now.With(date).BeginningOfMonth()
	lastOfMonth := now.With(date).EndOfMonth()

	var days [][]string
	weekCount := 0

	dt := firstOfMonth
	for dt.Before(lastOfMonth) {
		if len(days) == weekCount {
			days = append(days, []string{blank, blank, blank, blank, blank, blank, blank})
		}

		str := fmt.Sprintf("%2d", dt.Day())
		if datesEqual(dt, date) && highlightDt {
			str = Sprintf(Reverse(White(str)))
		}
		days[weekCount][dt.Weekday()] = str
		if dt.Weekday() == 6 {
			weekCount += 1
		}
		dt = dt.Add(time.Hour * 24)
	}

	s := ""
	for _, week := range days {
		s += strings.Join(week, " ") + "\n"
	}
	return s
}

func datesEqual(t1 time.Time, t2 time.Time) bool {
	y1, m1, d1 := t1.Date()
	y2, m2, d2 := t2.Date()
	return y1 == y2 && m1 == m2 && d1 == d2
}

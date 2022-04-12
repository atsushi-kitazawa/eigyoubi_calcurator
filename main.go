package main

import (
	"flag"
	"log"
	"testing"
	"time"

	holiday "github.com/holiday-jp/holiday_jp-go"
)

var (
	from string
	to   string
)

func init() {
	flag.StringVar(&from, "from", time.Now().Format("2006/01/02"), "from day")
	flag.StringVar(&to, "to", "", "to day")
	testing.Init()
	flag.Parse()
}

func main() {
	// log.Printf("from=%s, to=%s", from, to)
	doMain()
}

func doMain() {
	fromTime, _ := time.Parse("2006/01/02", from)
	toTime, _ := time.Parse("2006/01/02", to)

	subDays := toTime.Sub(fromTime).Hours() / 24
	var count int
	for i := 0; i <= int(subDays); i++ {
		day := fromTime.AddDate(0, 0, i)
		if isWeekend(day) || isHoliday(day) {
			continue
		}
		count++
	}
	log.Printf("%v working days", count)
}

func isWeekend(day time.Time) bool {
	weekday := day.Local().Weekday()
	if weekday == time.Saturday || weekday == time.Sunday {
		return true
	}
	return false
}

func isHoliday(day time.Time) bool {
	return holiday.IsHoliday(day)
}

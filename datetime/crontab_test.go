package datetime

import (
	"fmt"
	"testing"
	"time"
)

func TestCrontabSecond(t *testing.T) {
	crontab, err := NewCrontab("* * * * * *", func() {
		fmt.Println("call func1", time.Now())
	}, func() {
		fmt.Println("call func2", time.Now())
	})
	if err != nil {
		fmt.Println(err)
		t.Fatal()
	}
	defer crontab.Stop()
	go crontab.Start()
	time.Sleep(time.Second * 10)
}

//func TestCrontabSecond(t *testing.T) {
//	crontab, err := New("*/5 * * * * *")
//	if err != nil {
//		fmt.Println(err)
//		t.Fatal()
//	}
//	defer crontab.Stop()
//	go crontab.Start()
//	time.Sleep(time.Second * 10)
//}

//func TestCrontabMinute(t *testing.T) {
//	crontab, err := New("* */1 * * * *")
//	if err != nil {
//		fmt.Println(err)
//		t.Fatal()
//	}
//	defer crontab.Stop()
//	go crontab.Start()
//	time.Sleep(time.Minute * 10)
//}

//func TestCrontabTicker(t *testing.T) {
//	crontab, err := New("1,5,7 1-59 * * * *")
//	if err != nil {
//		fmt.Println(err)
//		t.Fatal()
//	}
//	defer crontab.Stop()
//	go crontab.Start()
//	time.Sleep(time.Hour * 10)
//}

//func TestParseWeek(t *testing.T) {
//	weekdays, _, err := parseWeek("Mon-Wed")
//	panicError(err)
//	if weekdays.Size() != 3 {
//		t.Fatal("期望 3")
//	}
//
//	weekdays, _, err = parseWeek("3-0,5")
//	panicError(err)
//	if weekdays.Size() != 5 {
//		t.Fatal("期望 5")
//	}
//
//	weekdays, _, err = parseWeek("3-0,Mon")
//	panicError(err)
//	if weekdays.Size() != 6 {
//		t.Fatal("期望 6")
//	}
//}
//
//func TestParseMonth(t *testing.T) {
//	monthes, _, err := parseMonth("Jan-May")
//	panicError(err)
//	if monthes.Size() != 5 {
//		t.Fatal("期望 5")
//	}
//
//	monthes, _, err = parseMonth("3-1,5")
//	panicError(err)
//	if monthes.Size() != 11 {
//		t.Fatal("期望 11")
//	}
//
//	monthes, _, err = parseMonth("3-1,Jan")
//	panicError(err)
//	if monthes.Size() != 11 {
//		t.Fatal("期望 11")
//	}
//}
//
//func TestParseDay(t *testing.T) {
//	days, _, err := parseDay("1-5")
//	panicError(err)
//	if days.Size() != 5 {
//		t.Fatal("期望 5")
//	}
//
//	days, _, err = parseDay("31-1,5")
//	panicError(err)
//	if days.Size() != 3 {
//		t.Fatal("期望 3")
//	}
//
//	days, _, err = parseDay("1,3,5")
//	panicError(err)
//	if days.Size() != 3 {
//		t.Fatal("期望 3")
//	}
//}
//
//func TestParseHour(t *testing.T) {
//	hours, _, err := parseHour("1-5")
//	panicError(err)
//	if hours.Size() != 5 {
//		t.Fatal("期望 5")
//	}
//
//	hours, _, err = parseHour("3-1,5")
//	panicError(err)
//	if hours.Size() != 23 {
//		t.Fatal("期望 23")
//	}
//
//	hours, _, err = parseHour("1,3,5")
//	panicError(err)
//	if hours.Size() != 3 {
//		t.Fatal("期望 3")
//	}
//}
//
//func panicError(err error) {
//	if err != nil {
//		panic(err)
//	}
//}

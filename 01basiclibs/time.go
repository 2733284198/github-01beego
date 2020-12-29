package main

import (
	"fmt"
	"time"
)

func timeDemo() {
	now := time.Now()

	// 当前日期时间
	fmt.Println("Current time %v\n", now)

	year := now.Year()
	month := now.Month()
	day := now.Day()
	hour := now.Hour()
	minute := now.Minute()
	second := now.Second()

	fmt.Println("%d-%02d-%02d %02d:%02d:%02d\n", year, month, day, hour, minute, second)

}

func timestamp2Date(timestamp int64) {
	var now time.Time

	if timestamp == 0 {
		now := time.Unix(timestamp, 0)
		fmt.Println(now)
	} else {
		now := time.Now()

		// 当前日期时间
		fmt.Println("Current time %v\n", now)
	}

	year := now.Year()
	month := now.Month()
	day := now.Day()
	hour := now.Hour()
	minute := now.Minute()
	second := now.Second()

	fmt.Println("%2d-%2d-%2d %02d:%02d:%02d\n", year, month, day, hour, minute, second)

}

func main() {
	//timeDemo()

	now := time.Now()
	later := now.Add(time.Hour)

	fmt.Println(now)
	fmt.Println(now.Format("2006/01/02 15:04"))
	fmt.Println(later)
}

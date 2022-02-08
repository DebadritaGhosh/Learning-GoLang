package main

import (
	"fmt"
	"time"
)

func main() {

	fmt.Println("Handlint time in GoLang")
	presentTime := time.Now()
	fmt.Println(presentTime) // will show us present time(YYYY-DD-MM HH:MM:SS.SSSS....)

	fmt.Println(presentTime.Format("01-02-2006")) // Will show us only date

	fmt.Println(presentTime.Format("01-02-2006 Monday")) // Will show us current day and day of week (02-08-2022 Tuesday)

	fmt.Println(presentTime.Format("01-02-2006 15:04:05 Monday")) // 02-08-2022 28:32:11 Tuesday

	createDate := time.Date(2020, time.June, 30, 23, 23, 0, 0, time.UTC)
	fmt.Println(createDate) //2020-06-30 23:23:00 +0000 UTC
}

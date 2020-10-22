package main

import (
	"time"
	"fmt"
)


func main() {
	year, month, date := time.Now().Date()
	day := time.Now().Weekday()
	fmt.Println(year, month, date, day)


}


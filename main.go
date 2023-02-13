package main

import (
	"fmt"
	"time"
)

func main() {
	location, _ := time.LoadLocation("Asia/Shanghai")
	date := time.Date(2022, 1, 1, 0, 0, 0, 0, location)
	fmt.Println(date.Unix())
}

package main

import (
	"fmt"
	"time"
)

func main() {
	withNanos := "2006-01-02 15:04:05"

	time1 := "2020-03-07 15:00:00"

	t1,_ := time.Parse(withNanos, time1)

	fmt.Println(t1.Sub(time.Now()) - 8 * time.Hour)
}
package main

import (
	"fmt"
	"time"
)

func main() {
	t := time.Now()
	var time string = fmt.Sprintf("%d-%02d-%02d %02d:%02d", t.Year(), t.Month(), t.Day(), t.Hour(), t.Minute())
	fmt.Println(time)
}

package main

import (
	"fmt"
	"os"
	"time"
)

func main() {
	delay, err := time.ParseDuration(os.Args[1])
	if err != nil || len(os.Args) != 2 {
		panic("Expected a single duration argument")
	}
	now := time.Now()
	t := now.Add(delay)
	fmt.Printf("It is %v I will wake and exit at %v in %v\n", now, t, -time.Since(t))
	<-time.After(-time.Since(t))
	fmt.Printf("I was supposed to wake at %v, it is now %v\n", t, time.Now())
}

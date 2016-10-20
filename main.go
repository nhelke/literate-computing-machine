package main

import (
	"fmt"
	"os"
	"time"
)

func main() {
	if len(os.Args) != 2 {
		panic("Expected a single duration argument")
	}
	delay, err := time.ParseDuration(os.Args[1])
	if err != nil {
		panic(err)
	}
	now := time.Now()
	t := now.Add(delay)
	fmt.Printf("%v %v: should wake at %v in %v\n", delay, now, t, delay)
	<-time.After(delay)
	nownow := time.Now()
	fmt.Printf("%v %v: did wake with delta %v\n", delay, nownow, nownow.Sub(t))
}

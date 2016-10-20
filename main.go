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
	fmt.Printf("%v %v: should wake at %v in %v\n", delay, now, t, -time.Since(t))
	<-time.After(-time.Since(t))
	nownow := time.Now()
	fmt.Printf("%v %v: did wake with delta %v\n", delay, nownow, nownow.Sub(t))
}

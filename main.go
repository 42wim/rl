package main

import (
	"bufio"
	"flag"
	"fmt"
	"github.com/juju/ratelimit"
	"os"
	"time"
)

var (
	drops    int
	flagRate int64
	flagKeep bool
)

func init() {
	flag.Int64Var(&flagRate, "r", 5, "limit to r messages per second (drops those exceeding the limit)")
	flag.BoolVar(&flagKeep, "k", false, "keep the messages instead of dropping them")
	flag.Parse()
}

func reset() {
	if drops > 0 {
		fmt.Fprintf(os.Stderr, "dropped %d messages\n", drops)
		drops = 0
	}
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	c := make(chan string)
	done := make(chan bool)
	go func(c chan string, done chan bool) {
		for scanner.Scan() {
			c <- scanner.Text()
		}
		close(done)
	}(c, done)
	l := ratelimit.NewBucket(time.Second/time.Duration(flagRate), flagRate)
	timer := time.NewTicker(time.Second)
	for {
		select {
		case line := <-c:
			if flagKeep {
				l.Wait(1)
				fmt.Println(line)
			} else {
				if l.TakeAvailable(1) > 0 {
					fmt.Println(line)
					reset()
				} else {
					drops++
				}
			}
		case <-timer.C:
			if l.Available() == flagRate {
				reset()
			}
		case <-done:
			reset()
			return
		}
	}
}

package main

import (
	"fmt"
	"io"
	"os"
	"time"
)

const sleep = "sleep"
const write = "write"

type Sleeper interface {
	Sleep()
}

type SpySleeper struct {
	Calls int
}

func (s *SpySleeper) Sleep() {
	s.Calls++
}

type DefaultSleeper struct {
}

func (d *DefaultSleeper) Sleep() {
	time.Sleep(1 * time.Second)
}

type CountdownOperationSpy struct {
	Calls []string
}

func (c *CountdownOperationSpy) Sleep() {
	c.Calls = append(c.Calls, sleep)
}

func (c *CountdownOperationSpy) Write(p []byte) (n int, err error) {
	c.Calls = append(c.Calls, write)
	return
}

func Countdown(w io.Writer, s Sleeper) {
	for i := 3; i > 0; i-- {
		s.Sleep()
		fmt.Fprintln(w, i)
	}
	s.Sleep()
	fmt.Fprint(w, "Go!")
}

func main() {
	s := &DefaultSleeper{}
	Countdown(os.Stdout, s)
}

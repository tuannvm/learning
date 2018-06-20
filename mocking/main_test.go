package main

import (
	"bytes"
	"reflect"
	"testing"
)

func TestCountdown(t *testing.T) {
	t.Run("count to 3", func(t *testing.T) {
		buffer := &bytes.Buffer{}

		Countdown(buffer, &CountdownOperationSpy{})

		got := buffer.String()
		want := `3
2
1
Go!`

		if got != want {
			t.Errorf("got %s want %s", got, want)
		}

	})

	t.Run("sleep after every print", func(t *testing.T) {
		spySleeperPrinter := &CountdownOperationSpy{}
		Countdown(spySleeperPrinter, spySleeperPrinter)

		want := []string{
			sleep,
			write,
			sleep,
			write,
			sleep,
			write,
			sleep,
			write,
		}

		if !reflect.DeepEqual(spySleeperPrinter.Calls, want) {
			t.Errorf("want %v got %v", want, spySleeperPrinter.Calls)
		}
	})
}

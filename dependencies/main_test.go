package main

import (
	"bytes"
	"testing"
)

func TestGreet(t *testing.T) {
	buffer := bytes.Buffer{}
	Greet(&buffer, "Tuan")
	got := buffer.String()
	want := "Hello, Tuan"

	if got != want {
		t.Errorf("got '%s' want '%s'", got, want)
	}
}

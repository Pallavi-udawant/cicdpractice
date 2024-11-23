package main

import (
	"testing"
)

func Test_printHi(t *testing.T) {
	name := "Abhishek."
	want := "Hi Abhishek."

	if got := src.printHi(name); got != want {
		t.Errorf("hello() = %q, want %q", got, want)
	}
}

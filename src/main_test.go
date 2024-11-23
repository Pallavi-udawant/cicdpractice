package main

import (
	"fmt"
	"testing"
)

func Test_printHi(t *testing.T) {
	name := "Abhishek."
	want := "Hi Abhishek."

	if got := printHi(name); got != want {
		fmt.Println("success")
	}
}

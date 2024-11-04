package main

import (
	"testing"
	"strings"
)

func TestGreeting(t *testing.T) {
	name := "Alice"
	expected := "Hello, Alice!"
	response := "Hello, " + strings.Join([]string{name}, " ") + "!"

	if response != expected {
		t.Errorf("Expected %s, but got %s", expected, response)
	}
}

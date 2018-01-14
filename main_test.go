package main

import (
	"testing"
)

func TestHello(t *testing.T) {
	result := hello()
	if result != "Hello World" {
		t.Errorf("hello() = %s, want %s", result, "Hello World")
	}
}

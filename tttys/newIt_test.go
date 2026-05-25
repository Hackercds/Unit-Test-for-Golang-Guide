package tttys_test

import (
	"DeepTest/tttys"
	"testing"
)

func TestGreeting(t *testing.T) {
	got := tttys.Greeting()
	if got == "" {
		t.Error("Greeting() returned empty string")
	}
}

func TestInitialize(t *testing.T) {
	got := tttys.Initialize()
	if got == "" {
		t.Error("Initialize() returned empty string")
	}
}

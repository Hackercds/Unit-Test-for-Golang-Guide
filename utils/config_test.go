package utils_test

import (
	"DeepTest/utils"
	"testing"
)

func TestConfigMessage(t *testing.T) {
	got := utils.ConfigMessage()
	if got == "" {
		t.Error("ConfigMessage() returned empty string")
	}
}

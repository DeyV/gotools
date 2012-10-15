package gotools

import (
	"testing"
)

func TestNumerFormat(t *testing.T) {
	if a, _ := NumerFormat(1234, 0, ",", " "); a != "1 234" {
		t.Error("1234 => 1 234")
	}

	if a, _ := NumerFormat(1234.123, 2, ",", " "); a != "1 234,12" {
		t.Error("1234.123 => 1 234,12")
	}
}

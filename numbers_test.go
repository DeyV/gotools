package gotools

import (
	"testing"
)

type testNumberForm struct {
	Required               string
	Number                 float64
	Decimals               int
	DecPoint, ThousandsSep string
}

func TestNumerFormat(t *testing.T) {
	testdata := []testNumberForm{
		{"123 123,123", 123123.123, 3, ",", " "},
		{"-123 123,123", -123123.123, 3, ",", " "},
		{"-123 123,123", -123123.1234, 3, ",", " "},
		{"1 123,12", 1123.123, 2, ",", " "},
		{"123 123.1", 123123.123, 1, ".", " "},
		{"1233,12300", 1233.123, 5, ",", ""},
		{"1233.12300", 1233.123, 5, ".", ""},
		{"1233.12300", 1233.123, 5, "", ""},
	}

	for _, data := range testdata {
		assertNumer(t, data)
	}
}

func assertNumer(t *testing.T, data testNumberForm) {
	if a, _ := NumberFormat(data.Number, data.Decimals, data.DecPoint, data.ThousandsSep); a != data.Required {
		t.Errorf("%f => %s  but is %s", data.Number, data.Required, a)
	}

}

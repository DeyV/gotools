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
		{"1233.12300", 1233.123, 5, "", ""},

		{"233,12", 233.123, 2, ",", " "},
		{"23,12", 23.123, 2, ",", " "},
		{"3,12", 3.123, 2, ",", " "},
		{"0,12", 0.123, 2, ",", " "},
		{"0,02", 0.023, 2, ",", " "},

		{"0,3", 0.26, 1, ",", " "},
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

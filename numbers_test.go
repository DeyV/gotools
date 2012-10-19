package gotools

import (
	"math"
	"strconv"
	"testing"
)

type testNumberForm struct {
	Required               string
	Number                 float64
	Decimals               int
	DecPoint, ThousandsSep string
}

type testRound struct {
	Required string
	Number   float64
	Decimals int
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

		{"123312300", 1233.123, 5, "", ""},
		{"123312300", 1233.123, 5, "", ""},

		{"233,12", 233.123, 2, ",", " "},
		{"23,12", 23.123, 2, ",", " "},
		{"3,12", 3.123, 2, ",", " "},
		{"0,12", 0.123, 2, ",", " "},
		{"0,02", 0.023, 2, ",", " "},

		{"0,3", 0.26, 1, ",", " "},

		{"0,0", math.NaN(), 1, ",", " "},
		{"0,0", math.Inf(0), 1, ",", " "},
	}

	for _, data := range testdata {
		assertNumerFormat(t, data)
	}
}

var testdataRoundPrec = []testRound{
	{"123123.123", 123123.123, 3},
	{"-123123.123", -123123.123, 3},

	{"123123.1", 123123.1, 1},
	{"123123.2", 123123.2, 1},

	{"123123.46", 123123.458, 2},
	{"123123.24", 123123.239, 2},
	{"123123.23", 123123.232, 2},

	{"123123.1", 123123.123, 1},
	{"123123.2", 123123.223, 1},

	{"123123.23", 123123.22499999999999, 2},
	{"123123.23", 123123.22499999999993, 2}, // float precision problem, the same like fmt.Sprintf( "%.2f")
	{"123123.22", 123123.2249993, 2},        // correct behavior, smaller precision

	{"454.44", 454.444, 2}, //  
	{"454.45", 454.445, 2}, //  
	{"454.45", 454.446, 2}, //  

	{"-454.44", -454.444, 2}, //  
	{"-454.45", -454.445, 2}, //  
	{"-454.45", -454.446, 2}, //  

	{"454", 454.444, 0}, //  
	{"454", 454.445, 0}, //  
	{"454", 454.446, 0}, // 

	{"-454", -454.444, 0}, //  
	{"-454", -454.445, 0}, //  
	{"-454", -454.446, 0}, // 

	{"18446744073709551616.00", 1 << 64, 2},   //  
	{"-18446744073709551616.00", -1 << 64, 2}, //  
}

func TestRoundPrec(t *testing.T) {
	for _, data := range testdataRoundPrec {
		assertRoundPrec(t, data)
	}

	var val float64
	val = math.NaN()
	if out := RoundPrec(val, 0); !math.IsNaN(out) {
		t.Errorf("Round from NaN should be Nan, but is %v", out)
	}
	val = math.Inf(1)
	if out := RoundPrec(val, 0); !math.IsInf(out, 0) {
		t.Errorf("Round from Inf should be Inf, but is %v", out)
	}
}

func TestRound(t *testing.T) {
	var val float64

	val = 12.9
	if out := Round(val); out != 13 {
		t.Errorf("%f => %s  but is %s", val, 13, out)
	}

	val = 12.4
	if out := Round(val); out != 12 {
		t.Errorf("%f => %s  but is %s", val, 12, out)
	}

	val = -12.9
	if out := Round(val); out != -13 {
		t.Errorf("%f => %s  but is %s", val, -13, out)
	}
	val = math.NaN()
	if out := Round(val); out != 0 {
		t.Errorf("Round from NaN should be 0, but is %v", out)
	}
	val = math.Inf(0)
	if out := Round(val); out != 0 {
		t.Errorf("Round from Inf should be 0, but is %v", out)
	}
}

func assertNumerFormat(t *testing.T, data testNumberForm) {
	if a := NumberFormat(data.Number, data.Decimals, data.DecPoint, data.ThousandsSep); a != data.Required {
		t.Errorf("%f => %s  but is %s", data.Number, data.Required, a)
	}
}

func assertRoundPrec(t *testing.T, data testRound) {
	a := RoundPrec(data.Number, data.Decimals)
	out := strconv.FormatFloat(a, 'f', data.Decimals, 64)
	if out != data.Required {
		t.Errorf("%f => %s  but is %s", data.Number, data.Required, out)
	}
}

func BenchmarkNumberFormat(b *testing.B) {
	for i := float64(0); i <= float64(b.N); i++ {
		NumberFormat(i, 3, ",", " ")
	}
}

func BenchmarkRoundPrec(b *testing.B) {
	for i := float64(0); i <= float64(b.N); i++ {
		RoundPrec(i/1234, 3)
	}
}

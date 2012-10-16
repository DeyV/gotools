// gotools is set of simple math and test helpers.
//
// ## Install
//	go get github.com/DeyV/gotools
package gotools

import (
	"math"
	"strconv"
)

// NumberFormat convert float or int to string
// in local format, for example:
//	NumberFormat( 123456.12, 4, ",", " " )  
//	>> 123 456,1200
//
// Special cases are: 
//	NumberFormat(±Inf) = formatted 0.0
//	NumberFormat(NaN) = formatted 0.0
func NumberFormat(number float64, decimals int, decPoint, thousandsSep string) string {
	if math.IsNaN(number) || math.IsInf(number, 0) {
		number = 0
	}

	var ret string
	var negative bool

	if number < 0 {
		number *= -1
		negative = true
	}

	d, fract := math.Modf(number)

	if decimals <= 0 {
		fract = 0
	} else {
		pow := math.Pow(10, float64(decimals))
		fract = RoundPrec(fract*pow, 0)
	}

	if thousandsSep == "" {
		ret = strconv.FormatFloat(d, 'f', 0, 64)
	} else if d >= 1 {
		var x float64
		for d >= 1 {
			d, x = math.Modf(d / 1000)
			x = x * 1000
			ret = strconv.FormatFloat(x, 'f', 0, 64) + ret
			if d >= 1 {
				ret = thousandsSep + ret
			}
		}
	} else {
		ret = "0"
	}

	fracts := strconv.FormatFloat(fract, 'f', 0, 64)

	// "0" pad left
	for i := len(fracts); i < decimals; i++ {
		fracts = "0" + fracts
	}

	ret += decPoint + fracts

	if negative {
		ret = "-" + ret
	}
	return ret
}

// Round convert x to int with correct math round 
//
// Special cases are:
//	Round(±0) = ±0
//	Round(±Inf) = ±0
//	Round(NaN) = 0
func Round(x float64) int {
	val := RoundPrec(x, 0)
	if math.IsNaN(val) || math.IsInf(val, 0) {
		return 0
	}

	return int(val)
}

// RoundPrec return rounded version of x with prec precision.
//
// Special cases are:
//	Round(±0) = ±0
//	Round(±Inf) = ±Inf
//	Round(NaN) = NaN
func RoundPrec(x float64, prec int) float64 {
	pow := math.Pow(10, float64(prec))
	intermed := x * pow
	_, frac := math.Modf(intermed)

	var rounder float64
	if frac >= 0.5 {
		rounder = math.Ceil(intermed)
	} else {
		rounder = math.Floor(intermed)
	}

	return rounder / pow
}

package gotools

import (
	"fmt"
	"math"
)

// NumberFormat convert float or int to string
// in local format, for example 123 456,1200
// from float 123456.12
func NumberFormat(number float64, decimals int, decPoint, thousandsSep string) (string, error) {
	if decPoint == "" {
		decPoint = "."
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
		fract, _ = math.Modf(fract * pow)
	}

	var x float64
	if thousandsSep == "" {
		ret = fmt.Sprintf("%.0f", d)
	} else if d >= 1 {
		for d >= 1 {
			d, x = math.Modf(d / 1000)
			x = x * 1000
			ret = fmt.Sprintf("%.0f", x) + ret
			if d >= 1 {
				ret = thousandsSep + ret
			}
		}
	} else {
		ret = "0"
	}

	if fract > 0 {
		fmt.Println(fract)
		ret += decPoint + fmt.Sprintf("%.0f", fract)
	}

	if negative {
		ret = "-" + ret
	}
	return ret, nil
}

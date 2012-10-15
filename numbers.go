package gotools

import (
	"fmt"
	"math"
)

func NumerFormat(number interface{}, decimals int, decPoint, thousandsSep string) (string, error) {
	var num float64

	if decPoint == "" {
		decPoint = "."
	}

	switch i := number.(type) {
	case int:
		num = float64(number.(int))
	case float64, float32:
		num = i.(float64)
	case nil:
		num = 0
	}

	var ret string

	var negative bool
	if num < 0 {
		num *= -1
		negative = true
	}

	d, fract := math.Modf(num)

	if decimals <= 0 {
		fract = 0
	} else {
		pow := math.Pow(10, float64(decimals))
		fract, _ = math.Modf(fract * pow)
	}

	var x float64
	if thousandsSep == "" {
		ret = fmt.Sprintf("%.0f", d)
	} else {
		for d >= 1 {
			d, x = math.Modf(d / 1000)
			x = x * 1000
			ret = fmt.Sprintf("%.0f", x) + ret
			if d >= 1 {
				ret = thousandsSep + ret
			}
		}
	}

	//
	if fract > 0 {
		ret += decPoint + fmt.Sprintf("%.0f", fract)
	}

	if negative {
		ret = "-" + ret
	}
	// fmt.Println("...")
	return ret, nil
}

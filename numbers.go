package gotools

import (
	"fmt"
)

func NumerFormat(number interface{}, decimals int, decPoint, thousandsSep string) (string, error) {
	var num float64

	switch i := number.(type) {
	case int:
		num = float64(number.(int))
	case float64, float32:
		num = i.(float64)
	case nil:
		num = 0
	}

	var format string
	if decimals > 0 {
		format = fmt.Sprint("%%.%df", decimals)
	} else {
		format = "%f"
	}

	return fmt.Sprintf(format, num), nil
}

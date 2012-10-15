package gotools

import (
	"fmt"
	"strings"
)

func NumerFormat(number interface{}, decimals int, decPoint, thousandsSep string) (string, error) {
	var num float64
	var format string

	if decimals > 0 {
		format = fmt.Sprint("%%.%df", decimals)
	} else {
		format = "%f"
	}

	num, _ = caseToFloat(number)

	if thousandsSep == "" {
		out, _ := fmt.Sprintf(format, num), nil
		return strings.Replace(out, ".", decPoint, 1), nil
	}

	return fmt.Sprintf(format, num), nil
}

func caseToFloat(number interface{}) (float64, error) {
	var num float64

	switch i := number.(type) {
	case int:
		num = float64(number.(int))
	case float64, float32:
		num = i.(float64)
	case nil:
		num = 0
	}
	return num, nil
}

# gotools
--

## Usage

#### func  NumberFormat

```go
func NumberFormat(number interface{}, decimals int, decPoint, thousandsSep string) (string, error)
```
NumberFormat convert float or int to string in local format, for example 123
456,1200 from float 123456.12

# gotools
--
gotools is set of simple math and test helpers.

## Install
	go get github.com/DeyV/gotools

## Usage

```go
const (
	B  = iota // ignore first value by assigning to blank identifier
	KB = 1 << (10 * iota)
	MB
	GB
	TB
	PB
	EB
	ZB
	YB
)
```

#### func  MemoryFormat

```go
func MemoryFormat(x float64) string
```
MemoryFormat retun memory size (in Bytes) in verbalize form, for example 2.01MB

#### func  NumberFormat

```go
func NumberFormat(number float64, decimals int, decPoint, thousandsSep string) string
```
NumberFormat convert float or int to string (like PHP number_format() ) in local
format, for example:

    NumberFormat( 123456.12, 4, ",", " " )
    >> 123 456,1200

Special cases are:

    NumberFormat(±Inf) = formatted 0.0
    NumberFormat(NaN) = formatted 0.0

#### func  Round

```go
func Round(x float64) int
```
Round convert x to int with correct math round

Special cases are:

    Round(±0) = ±0
    Round(±Inf) = ±0
    Round(NaN) = 0

#### func  RoundPrec

```go
func RoundPrec(x float64, prec int) float64
```
RoundPrec return rounded version of x with prec precision.

Special cases are:

    Round(±0) = ±0
    Round(±Inf) = ±Inf
    Round(NaN) = NaN

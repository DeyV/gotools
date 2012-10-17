package gotools

import (
	"testing"
)

func TestMemorySizes(t *testing.T) {
	testData := map[float64]string{
		(2 << 9):               "1.00KB",
		(2 << 9) + 24:          "1.02KB",
		(2 << 9) * 10:          "10.00KB",
		(2 << 19):              "1.00MB",
		(2 << 19) * 10:         "10.00MB",
		(2<<19)*10 + (2<<9)*10: "10.01MB",
		(2 << 29):              "1.00GB",
	}

	for in, out := range testData {
		assertkByte(t, in, out)
	}

}

func assertkByte(t *testing.T, in float64, out string) {
	ret := MemoryFormat(in)

	if ret != out {
		t.Errorf("%f bytes should be %v - now is %v", in, out, ret)
	}
}

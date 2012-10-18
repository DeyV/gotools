package gotools

import (
	"fmt"
)

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

// MemoryFormat retun memory size (in Bytes) in verbalize form, 
// for example 2.01MB.
func MemoryFormat(x float64) string {
	switch {
	case x >= YB:
		return fmt.Sprintf("%.2fYB", x/YB)
	case x >= ZB:
		return fmt.Sprintf("%.2fZB", x/ZB)
	case x >= EB:
		return fmt.Sprintf("%.2fEB", x/EB)
	case x >= PB:
		return fmt.Sprintf("%.2fPB", x/PB)
	case x >= TB:
		return fmt.Sprintf("%.2fTB", x/TB)
	case x >= GB:
		return fmt.Sprintf("%.2fGB", x/GB)
	case x >= MB:
		return fmt.Sprintf("%.2fMB", x/MB)
	case x >= KB:
		return fmt.Sprintf("%.2fKB", x/KB)
	}
	return fmt.Sprintf("%.0fB", x)
}

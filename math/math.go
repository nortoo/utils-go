package math

import (
	"fmt"
	"strconv"

	"github.com/nortoo/utils-go/types"
)

// Decimal returns a specific precise decimal.
func Decimal(value float64, n uint32) float64 {
	format := "%." + types.Int64ToString(int64(n)) + "f"
	value, _ = strconv.ParseFloat(fmt.Sprintf(format, value), 64)
	return value
}

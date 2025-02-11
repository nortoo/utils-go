package math

import (
	"fmt"
	"github.com/nortoo/utils-go/type"
	"strconv"
)

// Decimal returns a specific precise decimal.
func Decimal(value float64, n uint32) float64 {
	format := "%." + _type.Int64ToString(int64(n)) + "f"
	value, _ = strconv.ParseFloat(fmt.Sprintf(format, value), 64)
	return value
}

package utils

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

// ConvertToDesiredUnit converts a value with a unit prefix (e.g., "2G") to the specified desired unit (e.g., "m").
// It returns a Unit struct containing the converted value in both float and integer representations.
// Int value truncates the decimal part and does not round the value
func ConvertToDesiredUnit(value, desiredUnit string) Unit {
	units := "KMGTPE"
	prefix := strings.ToUpper(string(value[len(value)-1]))
	valueFloat, _ := strconv.ParseFloat(value[:len(value)-1], 32)
	prefixId := 0
	for idx := 0; idx < len(units); idx++ {
		if string(units[idx]) == prefix {
			prefixId = idx
			break
		}
	}

	desiredUnit = strings.ToUpper(desiredUnit)
	desiredUnitId := 0
	for idx := 0; idx < len(units); idx++ {
		if string(units[idx]) == desiredUnit {
			desiredUnitId = idx
			break
		}
	}

	switch {
	case prefixId > desiredUnitId:
		for idx := 0; idx < prefixId-desiredUnitId; idx++ {
			valueFloat *= 1024
		}
	case prefixId < desiredUnitId:
		for idx := 0; idx < desiredUnitId-prefixId; idx++ {
			valueFloat /= 1024
		}
	}

	intVal := int(valueFloat)
	return Unit{
		FloatStr:   fmt.Sprintf("%.1f%s", valueFloat, desiredUnit),
		FloatValue: math.Round(valueFloat*10) / 10,
		IntStr:     fmt.Sprintf("%d%s", intVal, desiredUnit),
		IntValue:   intVal,
	}
}

func ByteCountIEC(b int64) string {
	const unit = 1024
	if b < unit {
		return fmt.Sprintf("%d B", b)
	}
	div, exp := int64(unit), 0
	for n := b / unit; n >= unit; n /= unit {
		div *= unit
		exp++
	}
	return fmt.Sprintf("%.1f %ciB",
		float64(b)/float64(div), "KMGTPE"[exp])
}

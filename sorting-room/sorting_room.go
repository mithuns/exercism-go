package sorting

import (
	"fmt"
	"strconv"
)

// DescribeNumber should return a string describing the number.
func DescribeNumber(f float64) string {
	return fmt.Sprintf("This is the number %.1f", f)
}

type NumberBox interface {
	Number() int
}

// DescribeNumberBox should return a string describing the NumberBox.
func DescribeNumberBox(nb NumberBox) string {
	return fmt.Sprintf("This is a box containing the number %.1f", float32(nb.Number()))
}

type FancyNumber struct {
	n string
}

func (i FancyNumber) Value() string {
	return i.n
}

type FancyNumberBox interface {
	Value() string
}

// ExtractFancyNumber should return the integer value for a FancyNumber
// and 0 if any other FancyNumberBox is supplied.
func ExtractFancyNumber(fnb FancyNumberBox) int {
	switch v := fnb.(type) {
	case FancyNumber:
		result, _ := strconv.Atoi(v.Value())
		return result
	default:
		return 0
	}
}

// DescribeFancyNumberBox should return a string describing the FancyNumberBox.
func DescribeFancyNumberBox(fnb FancyNumberBox) string {
	return fmt.Sprintf("This is a fancy box containing the number %.1f", float64(ExtractFancyNumber(fnb)))
}

// DescribeAnything should return a string describing whatever it contains.
func DescribeAnything(i interface{}) string {
	switch v := i.(type) {
	case FancyNumber:
		result, _ := strconv.Atoi(v.Value())
		return fmt.Sprintf("This is a fancy box containing the number %.1f", float64(result))
	case FancyNumberBox:
		result, _ := strconv.Atoi(v.Value())
		return fmt.Sprintf("This is a fancy box containing the number %.1f", float64(result))
	case NumberBox:
		return fmt.Sprintf("This is a box containing the number %.1f", float64(v.Number()))
	case int:
		return fmt.Sprintf("This is the number %.1f", float64(i.(int)))
	case float32:
		return fmt.Sprintf("This is the number %.1f", float64(i.(float32)))
	case float64:
		return fmt.Sprintf("This is the number %.1f", float64(i.(float64)))
	}
	return "Return to sender"
}

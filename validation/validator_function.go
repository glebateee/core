package validation

import (
	"fmt"
	"strconv"
)

func required(fieldName string, val any, arg string) (bool, error) {
	if str, ok := val.(string); ok {
		if str == "" {
			return false, fmt.Errorf("value is required: %v", str)
		}
		return true, nil
	}
	return false, fmt.Errorf("'required' is only for strings")
}

func min(fieldName string, val any, arg string) (bool, error) {
	minVal, err := strconv.Atoi(arg)
	if err != nil {
		return false, fmt.Errorf("invalid argument for validator: %v", arg)
	}
	err = fmt.Errorf("'min' value (int) is: %v", minVal)
	if i, ok := val.(int); ok {
		if i >= minVal {
			return true, nil
		}
		return false, err
	}
	if f, ok := val.(float64); ok {
		if f >= float64(minVal) {
			return true, nil
		}
		return false, err
	}
	if s, ok := val.(string); ok {
		if len(s) >= minVal {
			return true, nil
		}
		return false, err
	}
	return false, fmt.Errorf("'min' validator is for int, float64, string only: %T", val)
}

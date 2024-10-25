package go_policy_enforcer

import (
	"fmt"
	"reflect"

	"github.com/kmesiab/go-policy-enforcer/internal/utils"
)

// EqualsPolicyCheckOperator checks if two values are equal.
// Returns true if leftVal is equal to rightVal.
var EqualsPolicyCheckOperator = func(leftVal, rightVal any) bool {
	leftVal = utils.CoerceToComparable(leftVal)
	rightVal = utils.CoerceToComparable(rightVal)
	return reflect.DeepEqual(leftVal, rightVal)
}

// NotEqualsPolicyCheckOperator checks if two values are not equal.
// Returns true if leftVal is not equal to rightVal.
var NotEqualsPolicyCheckOperator = func(leftVal, rightVal any) bool {
	leftVal = utils.CoerceToComparable(leftVal)
	rightVal = utils.CoerceToComparable(rightVal)
	return !reflect.DeepEqual(leftVal, rightVal)
}

// GreaterThanPolicyCheckOperator checks if the left value is greater than the right value.
// Assumes both values are integers. Returns true if leftVal is greater than rightVal.
var GreaterThanPolicyCheckOperator = func(leftVal, rightVal any) bool {
	switch left := leftVal.(type) {
	case int:
		// Handle int vs float64
		switch right := rightVal.(type) {
		case int:
			return left > right
		case float64:
			return float64(left) > right
		default:
			return false
		}
	case float64:
		// Handle float64 vs int
		switch right := rightVal.(type) {
		case int:
			return left > float64(right)
		case float64:
			return left > right
		default:
			return false
		}
	case string:
		if right, ok := rightVal.(string); ok {
			return left > right
		}
		return false
	default:
		return false
	}
}

// GreaterThanOrEqualsPolicyCheckOperator checks if the left value is greater than or equal to the right value.
// Assumes both values are integers. Returns true if leftVal is greater than or equal to rightVal.
var GreaterThanOrEqualsPolicyCheckOperator = func(leftVal, rightVal any) bool {
	switch left := leftVal.(type) {
	case int:
		switch right := rightVal.(type) {
		case int:
			return left >= right
		case float64:
			return float64(left) >= right
		default:
			return false
		}
	case float64:
		switch right := rightVal.(type) {
		case int:
			return left >= float64(right)
		case float64:
			return left >= right
		default:
			return false
		}
	case string:
		if right, ok := rightVal.(string); ok {
			return left >= right
		}
		return false
	default:
		return false
	}
}

// LessThanPolicyCheckOperator checks if the left value is less than the right value.
// Assumes both values are integers. Returns true if leftVal is less than rightVal.
var LessThanPolicyCheckOperator = func(leftVal, rightVal any) bool {
	switch left := leftVal.(type) {
	case int:
		switch right := rightVal.(type) {
		case int:
			return left < right
		case float64:
			return float64(left) < right
		default:
			return false
		}
	case float64:
		switch right := rightVal.(type) {
		case int:
			return left < float64(right)
		case float64:
			return left < right
		default:
			return false
		}
	case string:
		if right, ok := rightVal.(string); ok {
			return left < right
		}
		return false
	default:
		return false
	}
}

// LessThanOrEqualsPolicyCheckOperator checks if the left value is less than or
// equal to the right value.
var LessThanOrEqualsPolicyCheckOperator = func(leftVal, rightVal any) bool {
	switch left := leftVal.(type) {
	case int:
		switch right := rightVal.(type) {
		case int:
			return left <= right
		case float64:
			return float64(left) <= right
		default:
			return false
		}
	case float64:
		switch right := rightVal.(type) {
		case int:
			return left <= float64(right)
		case float64:
			return left <= right
		default:
			return false
		}
	case string:
		if right, ok := rightVal.(string); ok {
			return left <= right
		}
		return false
	default:
		return false
	}
}

// InPolicyCheckOperator checks if the left value exists in a slice of right values.
var InPolicyCheckOperator = func(leftVal, rightVal any) bool {
	rightSlice, ok := utils.ToStringSlice(rightVal)
	if !ok {
		fmt.Printf("Right value %v is not a valid slice\n", rightVal)
		return false
	}

	leftStr := fmt.Sprintf("%v", leftVal)
	fmt.Printf("Checking if %s is in %v\n", leftStr, rightSlice)

	for _, v := range rightSlice {
		if v == leftStr {
			fmt.Printf("%s is in %v\n", leftStr, rightSlice)
			return true
		}
	}

	fmt.Printf("%s is not in %v\n", leftStr, rightSlice)
	return false
}

package go_policy_enforcer

import (
	"fmt"
	"reflect"

	"github.com/kmesiab/go-policy-enforcer/internal/utils"
)

// equalsPolicyCheckOperator checks if two values are equal.
// Returns true if leftVal is equal to rightVal.
var equalsPolicyCheckOperator = func(leftVal, rightVal any) bool {
	leftVal = utils.CoerceToComparable(leftVal)
	rightVal = utils.CoerceToComparable(rightVal)
	return reflect.DeepEqual(leftVal, rightVal)
}

// notEqualsPolicyCheckOperator checks if two values are not equal.
// Returns true if leftVal is not equal to rightVal.
var notEqualsPolicyCheckOperator = func(leftVal, rightVal any) bool {
	leftVal = utils.CoerceToComparable(leftVal)
	rightVal = utils.CoerceToComparable(rightVal)
	return !reflect.DeepEqual(leftVal, rightVal)
}

// greaterThanPolicyCheckOperator checks if the left value is greater than the right value.
// Assumes both values are integers. Returns true if leftVal is greater than rightVal.
var greaterThanPolicyCheckOperator = func(leftVal, rightVal any) bool {
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

// greaterThanOrEqualsPolicyCheckOperator checks if the left value is greater than or equal to the right value.
// Assumes both values are integers. Returns true if leftVal is greater than or equal to rightVal.
var greaterThanOrEqualsPolicyCheckOperator = func(leftVal, rightVal any) bool {
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

// lessThanPolicyCheckOperator checks if the left value is less than the right value.
// Assumes both values are integers. Returns true if leftVal is less than rightVal.
var lessThanPolicyCheckOperator = func(leftVal, rightVal any) bool {
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

// lessThanOrEqualsPolicyCheckOperator checks if the left value is less than or
// equal to the right value.
var lessThanOrEqualsPolicyCheckOperator = func(leftVal, rightVal any) bool {
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

// inPolicyCheckOperator checks if the left value exists in a slice of right values.
var inPolicyCheckOperator = func(leftVal, rightVal any) bool {
	rightSlice, ok := utils.ToStringSlice(rightVal)
	if !ok {
		return false
	}

	leftStr := fmt.Sprintf("%v", leftVal)

	for _, v := range rightSlice {
		if v == leftStr {
			return true
		}
	}
	return false
}

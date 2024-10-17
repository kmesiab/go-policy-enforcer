package go_policy_enforcer

import (
	"fmt"
	"reflect"
	"strconv"
)

// PolicyCheckOperator is a function type that accepts two values of any type
// and returns a boolean result based on a comparison of the two values.
type PolicyCheckOperator func(leftVal, rightVal any) bool

// EqualsPolicyCheckOperator checks if two values are equal.
// Returns true if leftVal is equal to rightVal.
var EqualsPolicyCheckOperator = func(leftVal, rightVal any) bool {
	leftVal = coerceToComparable(leftVal)
	rightVal = coerceToComparable(rightVal)
	return reflect.DeepEqual(leftVal, rightVal)
}

// NotEqualsPolicyCheckOperator checks if two values are not equal.
// Returns true if leftVal is not equal to rightVal.
var NotEqualsPolicyCheckOperator = func(leftVal, rightVal any) bool {
	leftVal = coerceToComparable(leftVal)
	rightVal = coerceToComparable(rightVal)
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

// DeepEqualsPolicyCheckOperator checks if two slices or maps are deeply equal.
// Returns true if the left value is deeply equal to the right value.
var DeepEqualsPolicyCheckOperator = func(leftVal, rightVal any) bool {
	leftSlice, leftIsSlice := toStringSlice(leftVal)
	rightSlice, rightIsSlice := toStringSlice(rightVal)

	if leftIsSlice && rightIsSlice {
		return reflect.DeepEqual(leftSlice, rightSlice)
	}

	leftMap, leftIsMap := toMap(leftVal)
	rightMap, rightIsMap := toMap(rightVal)

	if leftIsMap && rightIsMap {
		return reflect.DeepEqual(leftMap, rightMap)
	}

	return false
}

// NotDeepEqualsPolicyCheckOperator checks if two slices or maps are not deeply equal.
// Returns true if the left value is not deeply equal to the right value.
var NotDeepEqualsPolicyCheckOperator = func(leftVal, rightVal any) bool {
	return !DeepEqualsPolicyCheckOperator(leftVal, rightVal)
}

// InPolicyCheckOperator checks if the left value exists in a slice of right values.
var InPolicyCheckOperator = func(leftVal, rightVal any) bool {
	rightSlice, ok := toStringSlice(rightVal)
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

// policyCheckOperatorMap maps string representations of comparison operators
// to their corresponding PolicyCheckOperator functions.
var policyCheckOperatorMap = map[string]PolicyCheckOperator{
	"==":  EqualsPolicyCheckOperator,
	"!=":  NotEqualsPolicyCheckOperator,
	"===": DeepEqualsPolicyCheckOperator,
	"!==": NotDeepEqualsPolicyCheckOperator,
	">":   GreaterThanPolicyCheckOperator,
	">=":  GreaterThanOrEqualsPolicyCheckOperator,
	"<":   LessThanPolicyCheckOperator,
	"<=":  LessThanOrEqualsPolicyCheckOperator,
	"in":  InPolicyCheckOperator,
}

// Helper function to convert any value to a map for comparison
func toMap(val any) (map[string]interface{}, bool) {
	v := reflect.ValueOf(val)
	if v.Kind() == reflect.Map {
		mapVal := make(map[string]interface{})
		for _, key := range v.MapKeys() {
			mapVal[fmt.Sprintf("%v", key.Interface())] = v.MapIndex(key).Interface()
		}
		return mapVal, true
	}
	return nil, false
}

// GetPolicyCheckOperator retrieves the appropriate PolicyCheckOperator function
// based on the provided operator string.
func GetPolicyCheckOperator(operator string) PolicyCheckOperator {
	if op, exists := policyCheckOperatorMap[operator]; exists {
		return op
	}
	// Return nil if the operator doesn't exist
	return nil
}

// EvaluatePolicyCheckOperator takes a string operator, a left value, and a right value,
// retrieves the corresponding PolicyCheckOperator function, and evaluates it with the given values.
// Returns the result of the comparison as a boolean.
func EvaluatePolicyCheckOperator(operator string, leftVal, rightVal any) bool {
	opFunc := GetPolicyCheckOperator(operator)
	if opFunc == nil {
		fmt.Printf("Operator '%s' is not supported\n", operator)
		return false
	}

	leftVal = dereferencePointer(leftVal)
	rightVal = dereferencePointer(rightVal)

	// Handle slice comparisons
	if reflect.TypeOf(leftVal).Kind() == reflect.Slice || reflect.TypeOf(rightVal).Kind() == reflect.Slice {
		return evaluateSliceComparison(leftVal, rightVal, operator)
	}

	leftVal = coerceToComparable(leftVal)
	rightVal = coerceToComparable(rightVal)

	return opFunc(leftVal, rightVal)
}

// dereferencePointer takes a value of any type and attempts to dereference it if it is a pointer.
// If the input value is a pointer and not nil, the function returns the value it points to.
// If the input value is not a pointer or is nil, the function returns the original value as is.
//
// Parameters:
// - val: The input value of any type.
//
// Returns:
// - The dereferenced value if the input value is a pointer and not nil.
// - The original value if the input value is not a pointer or is nil.
func dereferencePointer(val any) any {
	v := reflect.ValueOf(val)
	if v.Kind() == reflect.Ptr && !v.IsNil() {
		return v.Elem().Interface()
	}
	return val
}

// coerceToComparable takes a value of any type and attempts to coerce it into a comparable type.
// The function supports coercing strings to integers, floats, or strings if the conversion fails.
// If the input value is already an integer or float64, it is returned as is.
// If the input value is a string, the function attempts to convert it to an integer or float64.
// If the conversion is successful, the corresponding numeric type is returned.
// If the conversion fails, the function returns the original string value.
// If the input value is of an unsupported type, the function returns a string representation of the value.
func coerceToComparable(val any) any {
	switch v := val.(type) {
	case string:
		// Try to convert string to int or float
		if intValue, err := strconv.Atoi(v); err == nil {
			return intValue
		}
		if floatValue, err := strconv.ParseFloat(v, 64); err == nil {
			return floatValue
		}
		// If the string cannot be converted, return it as-is
		return v
	case int, float64:
		// If it's already a number, return it
		return v
	default:
		// If it's any other type, return it as-is
		return fmt.Sprintf("%v", v) // convert other types to string for comparison
	}
}

// Helper function to convert any slice to []string for comparison
// Helper function to convert any slice to []string for comparison
func toStringSlice(val any) ([]string, bool) {
	v := reflect.ValueOf(val)
	if v.Kind() == reflect.Slice {
		strSlice := make([]string, v.Len())
		for i := 0; i < v.Len(); i++ {
			strSlice[i] = fmt.Sprintf("%v", v.Index(i).Interface())
		}
		return strSlice, true
	}
	return nil, false
}

func evaluateSliceComparison(leftVal, rightVal any, operator string) bool {
	leftSlice, leftIsSlice := toStringSlice(leftVal)
	rightSlice, rightIsSlice := toStringSlice(rightVal)

	switch operator {
	case "==":
		// Check if both slices contain the same elements (order-agnostic)
		if leftIsSlice && rightIsSlice {
			return slicesContainSameElements(leftSlice, rightSlice)
		}
		// For a slice and a single item, return false (== is only valid between slices or single values)
		return false

	case "===":
		// Deep equality check (order matters)
		if leftIsSlice && rightIsSlice {
			return reflect.DeepEqual(leftSlice, rightSlice)
		}
		return false

	case "!=":
		// Negation of ==, check if slices do not contain the same elements (order-agnostic)
		if leftIsSlice && rightIsSlice {
			return !slicesContainSameElements(leftSlice, rightSlice)
		}
		return true // If one is a slice and the other isn't, they are not equal

	case "!==":
		// Negation of ===, deep equality check (order matters)
		if leftIsSlice && rightIsSlice {
			return !reflect.DeepEqual(leftSlice, rightSlice)
		}
		return true // If one is a slice and the other isn't, they are not deeply equal

	case "in":
		// Check if the left value is contained in the right slice
		if rightIsSlice && !leftIsSlice {
			for _, item := range rightSlice {
				if item == fmt.Sprintf("%v", leftVal) {
					return true
				}
			}
			return false
		}
		// Check if the right value is contained in the left slice
		if leftIsSlice && !rightIsSlice {
			for _, item := range leftSlice {
				if item == fmt.Sprintf("%v", rightVal) {
					return true
				}
			}
			return false
		}
	}

	return false
}

// Helper function to check if two slices contain the same elements (order-agnostic)
func slicesContainSameElements(slice1, slice2 []string) bool {
	if len(slice1) != len(slice2) {
		return false
	}

	elementCount := make(map[string]int)

	for _, elem := range slice1 {
		elementCount[elem]++
	}

	for _, elem := range slice2 {
		if count, exists := elementCount[elem]; !exists || count == 0 {
			return false
		}
		elementCount[elem]--
	}

	for _, count := range elementCount {
		if count != 0 {
			return false
		}
	}

	return true
}

package custom_operators

import (
	"reflect"

	"github.com/kmesiab/go-policy-enforcer/internal/utils"
)

// deepEqualsPolicyCheckFunc is a generic function that checks if two values are deeply equal.
// It supports comparing slices and maps of any comparable type.
//
// Parameters:
// - leftVal: The first value to compare.
// - rightVal: The second value to compare.
//
// Returns:
//   - A boolean indicating whether the two values are deeply equal.
//     Returns true if the left value is deeply equal to the right value.
//     Returns false otherwise.
func deepEqualsPolicyCheckFunc[T comparable](leftVal, rightVal any) bool {
	leftSlice, leftIsSlice := utils.TryConvertGenericSoTypedSlice[T](leftVal)
	rightSlice, rightIsSlice := utils.TryConvertGenericSoTypedSlice[T](rightVal)

	if leftIsSlice && rightIsSlice {
		return reflect.DeepEqual(leftSlice, rightSlice)
	}

	leftMap, leftIsMap := utils.ToMap(leftVal)
	rightMap, rightIsMap := utils.ToMap(rightVal)

	if leftIsMap && rightIsMap {
		return reflect.DeepEqual(leftMap, rightMap)
	}

	return false
}

// DeepEqualsPolicyCheckOperator checks if two values are deeply equal.
// It leverages the deepEqualsPolicyCheckFunc to determine equality, which can compare slices or maps deeply.
// Returns true if the left and right values are deeply equal, false otherwise.
var DeepEqualsPolicyCheckOperator = func(leftVal, rightVal any) bool {
	return deepEqualsPolicyCheckFunc[any](leftVal, rightVal)
}

// NotDeepEqualsPolicyCheckOperator checks if two slices or maps are not deeply equal.
// Returns true if the left value is not deeply equal to the right value.
var NotDeepEqualsPolicyCheckOperator = func(leftVal, rightVal any) bool {
	return !DeepEqualsPolicyCheckOperator(leftVal, rightVal)
}

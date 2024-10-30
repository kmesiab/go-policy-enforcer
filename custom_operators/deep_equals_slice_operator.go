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
	// Explicitly handle nil slices to preserve their nilness distinction
	if leftVal == nil && rightVal == nil {
		return true
	}
	// If they're both empty, they are the same
	if (utils.Len(leftVal) + utils.Len(rightVal)) == 0 {
		return true
	}

	// Check if both are slices
	leftSlice, leftIsSlice := utils.TryConvertGenericToTypedSlice[T](leftVal)
	rightSlice, rightIsSlice := utils.TryConvertGenericToTypedSlice[T](rightVal)
	if leftIsSlice && rightIsSlice {
		// If one slice is nil and the other is empty, treat them as unequal
		if (leftSlice == nil && len(rightSlice) == 0) || (rightSlice == nil && len(leftSlice) == 0) {
			return false
		}
		// Use reflect.DeepEqual for exact comparison of slices
		return reflect.DeepEqual(leftSlice, rightSlice)
	}

	// Check if both are maps
	leftMap, leftIsMap := utils.ToMap(leftVal)
	rightMap, rightIsMap := utils.ToMap(rightVal)
	if leftIsMap && rightIsMap {
		// Use reflect.DeepEqual for exact comparison of maps
		return reflect.DeepEqual(leftMap, rightMap)
	}

	// If neither both slices nor both maps, they are not the same thing
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

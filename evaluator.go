package go_policy_enforcer

import (
	"reflect"

	"github.com/kmesiab/go-policy-enforcer/internal/utils"
)

// PolicyCheckOperator is a function type that accepts two values of any type
// and returns a boolean result based on a comparison of the two values.
type PolicyCheckOperator[T comparable] func(T, T) bool

// evaluatePolicyCheckOperator takes a string operator, a left value, and a right value,
// retrieves the corresponding PolicyCheckOperator function, and evaluates it with the given values.
// Returns the result of the comparison as a boolean.
func evaluatePolicyCheckOperator(operator string, leftVal, rightVal any) (bool, error) {
	opFunc, err := getPolicyCheckOperator(operator)

	if opFunc == nil || err != nil {
		return false, err
	}

	leftVal = utils.DereferencePointer(leftVal)
	rightVal = utils.DereferencePointer(rightVal)

	// Handle slice comparisons
	if reflect.TypeOf(leftVal).Kind() == reflect.Slice || reflect.TypeOf(rightVal).Kind() == reflect.Slice {
		return EvaluateSliceComparison[any](leftVal, rightVal, operator), nil
	}

	leftVal = utils.CoerceToComparable(leftVal)
	rightVal = utils.CoerceToComparable(rightVal)

	return opFunc(leftVal, rightVal), nil
}

// EvaluateSliceComparison compares two slices or checks if a value is within a slice based on the given operator.
// It supports generic types, allowing it to work with slices of any comparable type.
//
// Parameters:
// - leftVal: The left-hand side value for comparison. It can be a slice or a single value.
// - rightVal: The right-hand side value for comparison. It can be a slice or a single value.
// - operator: The comparison operator. Supported operators are:
//   - "==": Checks if two slices contain the same elements.
//   - "!=": Checks if two slices do not contain the same elements.
//   - "===": Checks if two slices are deeply equal.
//   - "!==": Checks if two slices are not deeply equal.
//   - "in": Checks if a value is within a slice.
//   - "not in": Checks if a value is not within a slice.
//
// Returns:
// - bool: A boolean indicating the result of the comparison.
func EvaluateSliceComparison[T comparable](leftVal, rightVal any, operator string) bool {
	// Determine if left or right is the slice comparison
	leftSlice, leftIsSlice := utils.TryConvertGenericSoTypedSlice[T](leftVal)
	rightSlice, rightIsSlice := utils.TryConvertGenericSoTypedSlice[T](rightVal)

	// Comparing two slices
	if leftIsSlice && rightIsSlice {
		switch operator {
		case "==":
			return utils.SlicesContainSameElements(leftSlice, rightSlice)
		case "!=":
			return !utils.SlicesContainSameElements(leftSlice, rightSlice)
		case "===":
			return reflect.DeepEqual(leftSlice, rightSlice)
		case "!==":
			return !reflect.DeepEqual(leftSlice, rightSlice)
		}
	} else {
		// Check if a value is within a slice
		switch operator {

		case "in":
			if leftIsSlice {
				return utils.SliceContainsElement(rightVal.(T), leftSlice)
			} else if rightIsSlice {
				return utils.SliceContainsElement(leftVal.(T), rightSlice)
			}
		case "not in":
			if leftIsSlice {
				return !utils.SliceContainsElement(rightVal.(T), leftSlice)
			} else if rightIsSlice {
				return !utils.SliceContainsElement(leftVal.(T), rightSlice)
			}
		}
	}

	return false
}

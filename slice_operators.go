// slice_helpers.go

package go_policy_enforcer

import (
	"fmt"
	"reflect"
)

// slicesContainSameElements checks if two slices contain the same elements (order-agnostic).
//
// The function uses a map to count the occurrences of each element in the first slice.
// Then, it iterates over the second slice and checks if each element exists in the map
// with a non-zero count. If an element is not found or its count becomes zero, the function returns false.
// If all elements are found with non-zero counts, the function returns true.
//
// Parameters:
// - slice1: The first slice to compare. Must be a slice of strings.
// - slice2: The second slice to compare. Must be a slice of strings.
//
// Returns:
// - bool: Returns true if the two slices contain the same elements (order-agnostic), otherwise returns false.
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

// evaluateSliceComparison compares two values based on a given operator and returns a boolean result.
// The function supports comparing slices and values within slices.
//
// Parameters:
// - leftVal: The left-hand side value for comparison. Can be of any type.
// - rightVal: The right-hand side value for comparison. Can be of any type.
// - operator: The comparison operator to use. Supported operators are:
//   - "==": Checks if the two slices contain the same elements (order-agnostic).
//   - "!=": Checks if the two slices do not contain the same elements (order-agnostic).
//   - "===": Checks if the two slices are deeply equal.
//   - "!==": Checks if the two slices are not deeply equal.
//   - "in": Checks if the right value exists in the left slice.
//   - "not in": Checks if the right value does not exist in the left slice.
//
// Returns:
// - bool: Returns true if the comparison is successful, otherwise returns false.
func evaluateSliceComparison(leftVal, rightVal any, operator string) bool {
	leftSlice, leftIsSlice := toStringSlice(leftVal)
	rightSlice, rightIsSlice := toStringSlice(rightVal)

	// Compare a slice to a slice
	if leftIsSlice && rightIsSlice {
		switch operator {
		case "==":
			return slicesContainSameElements(leftSlice, rightSlice)
		case "!=":
			return !slicesContainSameElements(leftSlice, rightSlice)
		case "===":
			return reflect.DeepEqual(leftSlice, rightSlice)
		case "!==":
			return !reflect.DeepEqual(leftSlice, rightSlice)
		}
	} else {
		// Check a value within a slice
		switch operator {

		case "in":

			if leftIsSlice {
				return sliceContainsElement(rightVal, leftSlice)
			} else if rightIsSlice {
				return sliceContainsElement(leftVal, rightSlice)
			}

		case "not in":

			if leftIsSlice {
				return !sliceContainsElement(rightVal, leftSlice)
			} else if rightIsSlice {
				return !sliceContainsElement(leftVal, rightSlice)
			}
		}
	}

	return false
}

// sliceContainsElement checks if a given needle exists in a haystack slice.
// The function uses reflect.DeepEqual to compare the needle with each element in the haystack.
//
// Parameters:
// - needle: The value to search for in the haystack slice. Can be of any type.
// - haystack: The slice to search within. Must be a slice of strings.
//
// Returns:
// - bool: Returns true if the needle is found in the haystack, otherwise returns false.
func sliceContainsElement(needle any, haystack []string) bool {

	needleString := fmt.Sprintf("%s", needle)

	for _, hay := range haystack {
		if reflect.DeepEqual(needleString, hay) {
			return true
		}
	}
	return false
}

// toStringSlice converts any slice to []string for comparison.
// It uses reflection to handle different types of slices and converts each element to a string.
//
// Parameters:
//   - val: The input value that needs to be converted to a slice of strings.
//     It can be of any type that implements the reflect.Slice interface.
//
// Returns:
//   - []string: A slice of strings representing the input value.
//   - bool: A boolean indicating whether the conversion was successful.
//     Returns true if the input value is a slice, otherwise returns false.
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

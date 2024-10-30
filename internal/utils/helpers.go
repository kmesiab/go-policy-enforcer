package utils

import (
	"fmt"
	"reflect"
	"strconv"
)

// SlicesContainSameElements checks if two slices contain the same elements, regardless of their order.
// It uses a map to count the occurrences of each element in both slices and compares the counts.
//
// The function supports generic types, allowing it to work with slices of any comparable type.
//
// Parameters:
// - slice1: The first slice to compare. It can be of any comparable type.
// - slice2: The second slice to compare. It can be of any comparable type.
//
// Returns:
// - bool: A boolean indicating whether the two slices contain the same elements (true) or not (false).
func SlicesContainSameElements[T comparable](slice1, slice2 []T) bool {
	if len(slice1) != len(slice2) {
		return false
	}

	elementCount := make(map[T]int)
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

// SliceContainsElement checks if a given element exists within a slice.
//
// The function supports generic types, allowing it to work with slices of any comparable type.
//
// Parameters:
// - needle: The element to search for within the slice. It can be of any comparable type.
// - haystack: The slice to search within. It can be of any comparable type.
//
// Returns:
// - bool: A boolean indicating whether the needle element was found in the haystack slice (true) or not (false).
func SliceContainsElement[T comparable](needle T, haystack []T) bool {
	for _, elem := range haystack {
		if elem == needle {
			return true
		}
	}
	return false
}

// ToStringSlice attempts to convert any value into a slice of
// strings. If the input is a slice, each element is converted
// to a string. If the input is not a slice, it is converted
// to a single-element slice containing the string representation
// of the input.
//
// Parameters:
//   - val: The input value that needs to be converted to a slice of strings.
//     It can be of any type that implements the reflect.Slice interface.
//
// Returns:
//   - []string: A slice of strings representing the input value.
//   - bool: A boolean indicating whether the conversion was successful.
//     Returns true if the input value is a slice, otherwise returns false.
func ToStringSlice(val any) ([]string, bool) {
	v := reflect.ValueOf(val)
	if v.Kind() == reflect.Slice {
		strSlice := make([]string, v.Len())
		for i := 0; i < v.Len(); i++ {
			strSlice[i] = fmt.Sprintf("%v", v.Index(i).Interface())
		}
		return strSlice, true
	}
	return []string{fmt.Sprintf("%v", val)}, false
}

// TryConvertGenericSoTypedSlice attempts to convert a generic slice to a typed
// slice of T. This function uses reflection to determine if the input is a
// slice, then tries to convert each element to type T.
//
// If the input value is not a slice or if any element cannot be converted to
// type T, the function returns nil and false. The function returns false if
// any element cannot be converted to type T, avoiding runtime panics.
//
// Parameters:
//   - val: The value to convert, expected to be a slice of elements that can be
//     cast to type T.
//
// Returns:
//   - []T: A typed slice containing the converted elements from the input
//     slice. Returns nil if the conversion fails.
//   - bool: Indicates success (true) if all elements were successfully converted
//     to type T; false otherwise.
func TryConvertGenericSoTypedSlice[T any](val any) ([]T, bool) {
	// Check if the input value is actually a slice and if it's nil
	v := reflect.ValueOf(val)
	if v.Kind() != reflect.Slice || !v.IsValid() || v.IsNil() {
		return nil, false
	}

	// Initialize a slice of type T with the same length as the input slice
	slice := make([]T, v.Len())

	// Attempt to convert each element in the slice to type T
	for i := 0; i < v.Len(); i++ {
		elem, ok := v.Index(i).Interface().(T)
		if !ok {
			return nil, false // Element is not of type T, so conversion fails
		}
		slice[i] = elem
	}

	return slice, true
}

// ToMap converts any value to a map for comparison.
// It only supports converting values of type map[any]any.
//
// Parameters:
// - val: The input value to be converted. It must be of type map[any]any.
//
// Returns:
// - map[string]interface{}: The resulting map with string keys and interface{} values.
// - bool: A boolean indicating whether the input value was a map (true) or not (false).
//
// Example usage:
//
//	inputMap := map[string]int{"key1": 1, "key2": 2}
//	convertedMap, isMap := toMap(inputMap)
//	fmt.Println(convertedMap) // Output: map[key1:1 key2:2]
//	fmt.Println(isMap)       // Output: true
//
//	inputNotMap := "not a map"
//	convertedNotMap, isNotMap := toMap(inputNotMap)
//	fmt.Println(convertedNotMap) // Output: nil
//	fmt.Println(isNotMap)       // Output: false
func ToMap(val any) (map[string]interface{}, bool) {
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

// DereferencePointer takes a value of any type and attempts to dereference it if it is a pointer.
// If the input value is a pointer and not nil, the function returns the value it points to.
// If the input value is not a pointer or is nil, the function returns the original value as is.
//
// Parameters:
// - val: The input value of any type.
//
// Returns:
// - The dereferenced value if the input value is a pointer and not nil.
// - The original value if the input value is not a pointer or is nil.
func DereferencePointer(val any) any {
	v := reflect.ValueOf(val)

	// Handle pointers to interfaces safely
	for v.Kind() == reflect.Ptr && !v.IsNil() {
		v = v.Elem()
	}

	// Return final dereferenced value
	return v.Interface()
}

// CoerceToComparable takes a value of any type and attempts to coerce it into a comparable type.
// It supports converting strings to integers or floats, returning the original string if conversion fails.
// If the input value is already an integer or float64, it is returned as is.
// For unsupported types, the function returns a string representation of the value.
func CoerceToComparable(val any) any {
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
	case int, int8, int16, int32, int64:
		// If it's an integer type, return it as is
		return v
	case uint, uint8, uint16, uint32, uint64:
		// If it's an unsigned integer type, return it as is
		return v
	case float32, float64:
		// If it's a float type, return it as is
		return v
	default:
		// If it's any other type, return it as-is
		return fmt.Sprintf("%v", v) // convert other types to string for comparison
	}
}

// Len is a generic function that calculates the length of a given value.
// It supports arrays, slices, strings, and maps. For other types, it returns 0.
//
// Parameters:
// - value: The input value of any type.
//
// Returns:
//   - int: The length of the input value. If the input value is not an array, slice, string, or map,
//     the function returns 0.
func Len[T any](value T) int {
	v := reflect.ValueOf(value)
	switch v.Kind() {
	case reflect.Array, reflect.Slice, reflect.String, reflect.Map:
		return v.Len()
	default:
		return 0
	}
}

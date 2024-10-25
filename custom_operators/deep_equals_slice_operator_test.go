package custom_operators

import "testing"

func TestDeepEqualsPolicyCheckFunc_TwoSlicesOfComparableTypes(t *testing.T) {
	leftVal := []int{1, 2, 3}
	rightVal := []int{1, 2, 3}

	result := deepEqualsPolicyCheckFunc[int](leftVal, rightVal)

	if !result {
		t.Errorf("expected deepEqualsPolicyCheckFunc to return true for two slices of comparable types, but got false")
	}
}

func TestDeepEqualsPolicyCheckFunc_TwoSlicesOfDifferentComparableTypes(t *testing.T) {
	leftVal := []string{"apple", "banana", "cherry"}
	rightVal := []string{"apple", "banana", "cherry"}

	result := deepEqualsPolicyCheckFunc[string](leftVal, rightVal)

	if !result {
		t.Errorf("expected deepEqualsPolicyCheckFunc to return true for two slices of different comparable types, but got false")
	}
}

func TestDeepEqualsPolicyCheckFunc_TwoMapsOfComparableTypes(t *testing.T) {
	leftVal := map[string]int{"apple": 1, "banana": 2, "cherry": 3}
	rightVal := map[string]int{"apple": 1, "banana": 2, "cherry": 3}

	result := deepEqualsPolicyCheckFunc[int](leftVal, rightVal)

	if !result {
		t.Errorf("expected deepEqualsPolicyCheckFunc to return true for two maps of comparable types, but got false")
	}
}

func TestDeepEqualsPolicyCheckFunc_TwoMapsOfDifferentComparableTypes(t *testing.T) {
	leftVal := map[string]int{"apple": 1, "banana": 2, "cherry": 3}
	rightVal := map[string]string{"apple": "1", "banana": "2", "cherry": "3"}

	result := deepEqualsPolicyCheckFunc[any](leftVal, rightVal)

	if result {
		t.Errorf("expected deepEqualsPolicyCheckFunc to return false for two maps of different comparable types, but got true")
	}
}

func TestDeepEqualsPolicyCheckFunc_SliceAndMap(t *testing.T) {
	leftVal := []int{1, 2, 3}
	rightVal := map[string]int{"one": 1, "two": 2, "three": 3}

	result := deepEqualsPolicyCheckFunc[int](leftVal, rightVal)

	if result {
		t.Errorf("expected deepEqualsPolicyCheckFunc to return false for a slice and a map of comparable types, but got true")
	}
}

func TestDeepEqualsPolicyCheckFunc_MapAndNonComparableType(t *testing.T) {
	leftVal := map[string]int{"apple": 1, "banana": 2, "cherry": 3}
	rightVal := "non-comparable"

	result := deepEqualsPolicyCheckFunc[any](leftVal, rightVal)

	if result {
		t.Errorf("expected deepEqualsPolicyCheckFunc to return false for a map and a non-comparable type, but got true")
	}
}

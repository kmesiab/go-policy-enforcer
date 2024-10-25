package utils

import "testing"

func TestSlicesContainSameElementsDifferentLengths(t *testing.T) {
	slice1 := []int{1, 2, 3}
	slice2 := []int{1, 2}

	result := SlicesContainSameElements(slice1, slice2)

	if result {
		t.Errorf("Expected false, but got true for slices of different lengths")
	}
}

func TestSlicesContainSameElements_DuplicateElements(t *testing.T) {
	slice1 := []int{1, 2, 3, 2, 1}
	slice2 := []int{3, 2, 1, 2, 1}

	result := SlicesContainSameElements(slice1, slice2)

	if !result {
		t.Errorf("Expected SlicesContainSameElements to return true, but got false")
	}
}

func TestSlicesContainSameElements_NegativeNumbers(t *testing.T) {
	slice1 := []int{-1, 0, 2, -3}
	slice2 := []int{2, -3, 0, -1}

	result := SlicesContainSameElements(slice1, slice2)
	if !result {
		t.Errorf("Expected slices to contain the same elements, but got false")
	}
}

func TestSlicesContainSameElements_Floats(t *testing.T) {
	slice1 := []float64{1.1, 2.2, 3.3}
	slice2 := []float64{3.3, 1.1, 2.2}

	result := SlicesContainSameElements(slice1, slice2)

	if !result {
		t.Errorf("Expected SlicesContainSameElements to return true for slices containing floating point numbers, but got false")
	}
}

func TestSlicesContainSameElements_EmptyElements(t *testing.T) {
	slice1 := []int{1, 2, 0, 3}
	slice2 := []int{0, 1, 2, 3}

	result := SlicesContainSameElements(slice1, slice2)

	if !result {
		t.Errorf("Expected SlicesContainSameElements to return true, got false")
	}
}

func TestSlicesDoNotContainSameElementsWithDiffNil(t *testing.T) {
	var nilInt *int

	slice1 := []interface{}{1, nilInt, "abc"}
	slice2 := []interface{}{1, nil, "abc"}

	result := SlicesContainSameElements(slice1, slice2)
	if result {
		t.Errorf("Expected SlicesContainSameElements to return false for slices with nil values, but got true")
	}
}

func TestSlicesContainSameElementsNonComparableTypes(t *testing.T) {
	type customType struct {
		name string
	}

	slice1 := []customType{
		{name: "apple"},
		{name: "banana"},
		{name: "cherry"},
	}

	slice2 := []customType{
		{name: "banana"},
		{name: "cherry"},
		{name: "apple"},
	}

	result := SlicesContainSameElements(slice1, slice2)
	if !result {
		t.Errorf("Expected SlicesContainSameElements to return true for non-comparable types, but got false")
	}
}

func TestSlicesContainSameElementsLargeNumbers(t *testing.T) {
	slice1 := []int{1000000000, 2000000000, 3000000000}
	slice2 := []int{3000000000, 1000000000, 2000000000}

	result := SlicesContainSameElements(slice1, slice2)

	if !result {
		t.Errorf("Expected SlicesContainSameElements to return true for large numbers, but got false")
	}
}

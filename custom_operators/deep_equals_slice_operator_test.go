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

func TestDeepEqualsPolicyCheckFunc_SliceEdgeCases(t *testing.T) {
	tests := []struct {
		name     string
		left     []int
		right    []int
		expected bool
	}{
		{"different lengths", []int{1, 2, 3}, []int{1, 2}, false},
		{"different content", []int{1, 2, 3}, []int{1, 2, 4}, false},
		{"empty slices", []int{}, []int{}, true},
		{"nil slices", nil, nil, true},
		{"nil and empty", nil, []int{}, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := deepEqualsPolicyCheckFunc[int](tt.left, tt.right)
			if result != tt.expected {
				t.Errorf("got %v, want %v", result, tt.expected)
			}
		})
	}
}

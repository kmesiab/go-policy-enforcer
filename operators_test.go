package go_policy_enforcer

import (
	"reflect"
	"testing"

	"github.com/kmesiab/go-policy-enforcer/internal/utils"
)

// TestEqualsPolicyCheckOperator tests the equalsPolicyCheckOperator
func TestEqualsPolicyCheckOperator(t *testing.T) {
	tests := []struct {
		leftVal  any
		rightVal any
		expected bool
	}{
		{1, 1, true},
		{"abc", "abc", true},
		{1.0, 1.0, true},
		{1, 2, false},
		{"abc", "xyz", false},
	}

	for _, test := range tests {
		result := equalsPolicyCheckOperator(test.leftVal, test.rightVal)
		if result != test.expected {
			t.Errorf("EqualsPolicyCheckOperator(%v, %v) = %v; want %v", test.leftVal, test.rightVal, result, test.expected)
		}
	}
}

// TestNotEqualsPolicyCheckOperator tests the notEqualsPolicyCheckOperator
func TestNotEqualsPolicyCheckOperator(t *testing.T) {
	tests := []struct {
		leftVal  any
		rightVal any
		expected bool
	}{
		{1, 1, false},
		{"abc", "abc", false},
		{1.0, 2.0, true},
		{1, 2, true},
		{"abc", "xyz", true},
	}

	for _, test := range tests {
		result := notEqualsPolicyCheckOperator(test.leftVal, test.rightVal)
		if result != test.expected {
			t.Errorf("NotEqualsPolicyCheckOperator(%v, %v) = %v; want %v", test.leftVal, test.rightVal, result, test.expected)
		}
	}
}

// TestGreaterThanPolicyCheckOperator tests the greaterThanPolicyCheckOperator
func TestGreaterThanPolicyCheckOperator(t *testing.T) {
	tests := []struct {
		leftVal  any
		rightVal any
		expected bool
	}{
		{2, 1, true},
		{1, 1, false},
		{0, 1, false},
	}

	for _, test := range tests {
		result := greaterThanPolicyCheckOperator(test.leftVal, test.rightVal)
		if result != test.expected {
			t.Errorf("GreaterThanPolicyCheckOperator(%v, %v) = %v; want %v", test.leftVal, test.rightVal, result, test.expected)
		}
	}
}

// TestGreaterThanOrEqualsPolicyCheckOperator tests the greaterThanOrEqualsPolicyCheckOperator
func TestGreaterThanOrEqualsPolicyCheckOperator(t *testing.T) {
	tests := []struct {
		leftVal  any
		rightVal any
		expected bool
	}{
		{2, 1, true},
		{1, 1, true},
		{0, 1, false},
	}

	for _, test := range tests {
		result := greaterThanOrEqualsPolicyCheckOperator(test.leftVal, test.rightVal)
		if result != test.expected {
			t.Errorf("GreaterThanOrEqualsPolicyCheckOperator(%v, %v) = %v; want %v", test.leftVal, test.rightVal, result, test.expected)
		}
	}
}

// TestLessThanPolicyCheckOperator tests the lessThanPolicyCheckOperator
func TestLessThanPolicyCheckOperator(t *testing.T) {
	tests := []struct {
		leftVal  any
		rightVal any
		expected bool
	}{
		{1, 2, true},
		{1, 1, false},
		{2, 1, false},
	}

	for _, test := range tests {
		result := lessThanPolicyCheckOperator(test.leftVal, test.rightVal)
		if result != test.expected {
			t.Errorf("LessThanPolicyCheckOperator(%v, %v) = %v; want %v", test.leftVal, test.rightVal, result, test.expected)
		}
	}
}

// TestEvaluatePolicyCheckOperator tests the evaluatePolicyCheckOperator function
func TestEvaluatePolicyCheckOperator(t *testing.T) {
	tests := []struct {
		operator string
		leftVal  any
		rightVal any
		expected bool
	}{
		{"==", 1, 1, true},
		{"==", "abc", "abc", true},
		{"!=", 1, 2, true},
		{">", 2, 1, true},
		{">", "b", "a", true}, // String comparison
		{">=", 2, 1, true},
		{"<", 1, 2, true},
		{"<", "a", "b", true}, // String comparison
		{"<", 1, 1, false},
	}

	for _, test := range tests {
		result := evaluatePolicyCheckOperator(test.operator, test.leftVal, test.rightVal)
		if result != test.expected {
			t.Errorf("EvaluatePolicyCheckOperator(%v, %v, %v) = %v; want %v", test.operator, test.leftVal, test.rightVal, result, test.expected)
		}
	}
}

func TestGetPolicyCheckOperator_NonExistingOperator(t *testing.T) {

	nonExistingOperator := "non_existing_operator"
	opFunc, err := getPolicyCheckOperator(nonExistingOperator)

	if err == nil {
		t.Errorf("Expected error for non existing operator but got: %v", err)
	}

	if opFunc != nil {
		t.Errorf("Expected nil function for operator '%s', but got nil", nonExistingOperator)
	}
}

func TestGetPolicyCheckOperator_CaseSensitivity(t *testing.T) {
	tests := []struct {
		operator     string
		expectedFunc func(any, any) bool
		expectNil    bool
		expectError  bool
	}{
		{"==", equalsPolicyCheckOperator, false, false},
		{"!=", notEqualsPolicyCheckOperator, false, false},
		{">=", greaterThanOrEqualsPolicyCheckOperator, false, false},
		{"in", inPolicyCheckOperator, false, false},
		{"IN", nil, true, true}, // Case-sensitive mismatch should return nil without error
	}

	for _, tt := range tests {
		t.Run(tt.operator, func(t *testing.T) {
			opFunc, err := getPolicyCheckOperator(tt.operator)

			// Check for unexpected errors
			if (err != nil) != tt.expectError {
				t.Errorf("Unexpected error for operator '%s': got %v, want error: %v", tt.operator, err, tt.expectError)
			}

			// Check for expected nil or non-nil function
			if (opFunc == nil) != tt.expectNil {
				t.Errorf("Expected nil: %v for operator '%s', but got: %v", tt.expectNil, tt.operator, opFunc)
			}

			// Compare function behavior if a non-nil function was expected
			if !tt.expectNil && opFunc != nil && !compareFunctions(opFunc, tt.expectedFunc) {
				t.Errorf("Expected function behavior for operator '%s' does not match", tt.operator)
			}
		})
	}
}
func TestEvaluatePolicyCheckOperator_NilValues(t *testing.T) {
	operator := "=="

	// Test case 1: leftVal is nil, rightVal is a non-nil string
	var leftVal *string
	rightVal := "abc"

	result := evaluatePolicyCheckOperator(operator, leftVal, rightVal)
	if result {
		t.Errorf("EvaluatePolicyCheckOperator(%s, nil, %v) = %v; want %v", operator, rightVal, result, false)
	}

	// Test case 2: leftVal is a non-nil string, rightVal is nil
	leftValStr := "abc"
	leftVal = &leftValStr
	var rightValNil *string

	result = evaluatePolicyCheckOperator(operator, leftVal, rightValNil)
	if result {
		t.Errorf("EvaluatePolicyCheckOperator(%s, %v, nil) = %v; want %v", operator, leftValStr, result, false)
	}

	// Test case 3: both leftVal and rightVal are nil
	leftVal = nil
	rightValNil = nil

	result = evaluatePolicyCheckOperator(operator, leftVal, rightValNil)
	if !result {
		t.Errorf("EvaluatePolicyCheckOperator(%s, nil, nil) = %v; want %v", operator, result, true)
	}
}

// TestNonNumericValues tests the behavior of the comparison operators with non-numeric values
func TestNonNumericValues(t *testing.T) {
	tests := []struct {
		operator string
		leftVal  any
		rightVal any
		expected bool
	}{
		{">", "abc", "xyz", false},
		{">", "abc", "abc", false},
		{">=", "abc", "xyz", false},
		{">=", "abc", "abc", true},
		{"<", "abc", "xyz", true},
		{"<", "abc", "abc", false},
		{"<=", "abc", "xyz", true},
		{"<=", "abc", "abc", true},
	}

	for _, test := range tests {
		result := evaluatePolicyCheckOperator(test.operator, test.leftVal, test.rightVal)
		if result != test.expected {
			t.Errorf("EvaluatePolicyCheckOperator(%s, %v, %v) = %v; want %v", test.operator, test.leftVal, test.rightVal, result, test.expected)
		}
	}
}

func TestGetPolicyCheckOperator_NonStringValues(t *testing.T) {
	tests := []struct {
		operator string
		leftVal  any
		rightVal any
		expected bool
	}{
		{"==", 1, 1, true},
		{"==", 1.0, 1.0, true},
		{"==", "1", 1, true},
		{"==", "1.0", 1.0, true},
		{"!=", 1, 2, true},
		{"!=", 1.0, 2.0, true},
		{"!=", "1", 2, true},
		{"!=", "1.0", 2.0, true},
	}

	for _, test := range tests {
		opFunc, err := getPolicyCheckOperator(test.operator)

		if err != nil {
			t.Errorf("Expected no error for non existing operator but got: %v", err)
		}

		if opFunc == nil {
			t.Errorf("Expected non-nil function for operator '%s', but got nil", test.operator)
		} else {
			result := opFunc(test.leftVal, test.rightVal)
			if result != test.expected {
				t.Errorf("GetPolicyCheckOperator(%s, %v, %v) = %v; want %v", test.operator, test.leftVal, test.rightVal, result, test.expected)
			}
		}
	}
}

func TestEvaluatePolicyCheckOperator_UnsupportedOperator(t *testing.T) {
	operator := "unsupported_operator"
	leftVal := 10
	rightVal := 20

	result := evaluatePolicyCheckOperator(operator, leftVal, rightVal)
	if result {
		t.Errorf("Expected EvaluatePolicyCheckOperator('%s', %v, %v) to return false for unsupported operator, but got true", operator, leftVal, rightVal)
	}
}

func TestEvaluatePolicyCheckOperator_NonNumericValues(t *testing.T) {
	tests := []struct {
		operator string
		leftVal  any
		rightVal any
		expected bool
	}{
		{">", "abc", "xyz", false},
		{">", "abc", "abc", false},
		{">=", "abc", "xyz", false},
		{">=", "abc", "abc", true},
		{"<", "abc", "xyz", true},
		{"<", "abc", "abc", false},
		{"<=", "abc", "xyz", true},
		{"<=", "abc", "abc", true},
	}

	for _, test := range tests {
		result := evaluatePolicyCheckOperator(test.operator, test.leftVal, test.rightVal)
		if result != test.expected {
			t.Errorf("EvaluatePolicyCheckOperator(%s, %v, %v) = %v; want %v", test.operator, test.leftVal, test.rightVal, result, test.expected)
		}
	}
}

func TestEvaluatePolicyCheckOperator_DifferentDataTypes(t *testing.T) {
	sameSlice := []string{"abc", "xyz"}

	tests := []struct {
		operator string
		leftVal  any
		rightVal any
		expected bool
	}{
		{
			operator: "==",
			leftVal:  []string{"apple", "banana"},
			rightVal: []string{"apple", "banana"},
			expected: true,
		},
		{
			operator: "==",
			leftVal:  []string{"apple", "banana"},
			rightVal: []string{"banana", "apple"},
			expected: true,
		},
		{
			operator: "==",
			leftVal:  []string{"apple", "banana"},
			rightVal: "banana",
			expected: false,
		},
		{
			operator: "===",
			leftVal:  []string{"apple", "banana"},
			rightVal: []string{"banana", "apple"},
			expected: false,
		},
		{
			operator: "===",
			leftVal:  sameSlice,
			rightVal: sameSlice,
			expected: true,
		},
		{
			operator: "!=",
			leftVal:  []string{"apple", "banana"},
			rightVal: []string{"apple", "banana"},
			expected: false,
		},
		{
			operator: "!=",
			leftVal:  []string{"apple", "banana"},
			rightVal: []string{"banana", "apple"},
			expected: false,
		},
		{
			operator: "!=",
			leftVal:  []string{"apple", "banana"},
			rightVal: "bananas",
			expected: false,
		},
		{
			operator: "in",
			leftVal:  "banana",
			rightVal: []string{"apple", "banana"},
			expected: true,
		},
		{
			operator: "in",
			leftVal:  "orange",
			rightVal: []string{"apple", "banana"},
			expected: false,
		},
	}

	for _, test := range tests {
		result := evaluatePolicyCheckOperator(test.operator, test.leftVal, test.rightVal)
		if result != test.expected {
			t.Errorf("EvaluatePolicyCheckOperator(%s, %v, %v) = %v; want %v", test.operator, test.leftVal, test.rightVal, result, test.expected)
		}
	}
}

func TestEvaluatePolicyCheckOperator_SliceComparisons(t *testing.T) {
	deepSlice := []string{"apple", "banana"}

	tests := []struct {
		operator string
		leftVal  any
		rightVal any
		expected bool
	}{
		{"==", []string{"apple", "banana"}, "banana", false},                    // Not two slices
		{"==", []string{"apple", "banana"}, []string{"apple", "banana"}, true},  // Same elements same order
		{"==", []string{"apple"}, []string{"orange"}, false},                    // Not same elements
		{"!=", []string{"apple", "banana"}, []string{"apple", "banana"}, false}, // Same elements same order (negative)
		{"===", deepSlice, deepSlice, true},                                     // Deep equals
		{"===", deepSlice, []string{"apple"}, false},                            // Deep equals not same positive match
		{"!==", deepSlice, []string{"apple"}, true},                             // Deep equals not same (negative)_
		{"in", "orange", []string{"apple", "banana"}, false},                    // Single not in slice
	}

	for _, test := range tests {
		result := evaluatePolicyCheckOperator(test.operator, test.leftVal, test.rightVal)
		if result != test.expected {
			t.Errorf("EvaluatePolicyCheckOperator(%v, %v, %v) = %v; want %v", test.operator, test.leftVal, test.rightVal, result, test.expected)
		}
	}
}

func TestToStringSlice_NonNumericStringValues(t *testing.T) {
	// Test case: non-numeric string values in 'in' operator
	input := []interface{}{
		"apple",
		123,
		"banana",
		true,
		[]int{4, 5, 6},
	}

	expected := []string{"apple", "123", "banana", "true", "[4 5 6]"}

	result, ok := utils.ToStringSlice(input)

	if !ok {
		t.Errorf("Expected toStringSlice to return true, but got false")
	}

	if len(result) != len(expected) {
		t.Errorf("Expected result length %d, but got %d", len(expected), len(result))
	}

	for i := range result {
		if result[i] != expected[i] {
			t.Errorf("Expected result[%d] = %s, but got %s", i, expected[i], result[i])
		}
	}
}

func TestToStringSlice_EmptySlice(t *testing.T) {
	emptySlice := make([]interface{}, 0)
	var expectedResult []string

	result, ok := utils.ToStringSlice(emptySlice)

	if !ok {
		t.Errorf("Expected toStringSlice to return true, but got false")
	}

	if len(result) != len(expectedResult) {
		t.Errorf("Expected result length %d, but got %d", len(expectedResult), len(result))
	}

	for i := range result {
		if result[i] != expectedResult[i] {
			t.Errorf("Expected result[%d] = %s, but got %s", i, expectedResult[i], result[i])
		}
	}
}

type CustomTestType struct {
	value string
}

// Implement the fmt.Stringer interface for CustomType
func (c CustomTestType) String() string {
	return c.value
}

func TestToStringSlice_DifferentDataTypes(t *testing.T) {

	// Define CustomType with a String method

	tests := []struct {
		input    any
		expected []string
		isSlice  bool
	}{
		{
			input:    []CustomTestType{{value: "test"}, {value: "example"}},
			expected: []string{"test", "example"},
			isSlice:  true,
		},
		{
			input:    []int{1, 2, 3},
			expected: []string{"1", "2", "3"},
			isSlice:  true,
		},
		{
			input:    []float64{1.1, 2.2, 3.3},
			expected: []string{"1.1", "2.2", "3.3"},
			isSlice:  true,
		},
		{
			input:    []bool{true, false, true},
			expected: []string{"true", "false", "true"},
			isSlice:  true,
		},
		{
			input:    []interface{}{1, 2.2, "three", true},
			expected: []string{"1", "2.2", "three", "true"},
			isSlice:  true,
		},
		{
			input:    123,
			expected: nil,
			isSlice:  false,
		},
		{
			input:    "hello",
			expected: nil,
			isSlice:  false,
		},
		{
			input:    []interface{}{}, // Test empty slice
			expected: []string{},
			isSlice:  true,
		},
		{
			input:    []interface{}{nil, nil}, // Test slice with nil elements
			expected: []string{"<nil>", "<nil>"},
			isSlice:  true,
		},
	}

	for _, test := range tests {
		result, isSlice := utils.ToStringSlice(test.input)

		if isSlice != test.isSlice {
			t.Errorf("Input: %v\nExpected isSlice: %v, but got: %v", test.input, test.isSlice, isSlice)
		}

		if isSlice {
			if !reflect.DeepEqual(result, test.expected) {
				t.Errorf("Input: %v\nExpected: %v, but got: %v", test.input, test.expected, result)
			}
		} else {
			if result != nil {
				t.Errorf("Input: %v\nExpected result to be nil, but got: %v", test.input, result)
			}
		}
	}
}

// Helper function to compare two functions by applying them on the same test data
func compareFunctions(f1, f2 PolicyCheckOperator[any]) bool {
	// Test with sample values, adjust these according to the operator being tested
	testValue1 := 10
	testValue2 := 10
	return f1(testValue1, testValue2) == f2(testValue1, testValue2)
}

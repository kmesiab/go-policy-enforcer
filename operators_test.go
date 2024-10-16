package go_policy_enforcer

import (
	"testing"
)

// TestEqualsPolicyCheckOperator tests the EqualsPolicyCheckOperator
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
		result := EqualsPolicyCheckOperator(test.leftVal, test.rightVal)
		if result != test.expected {
			t.Errorf("EqualsPolicyCheckOperator(%v, %v) = %v; want %v", test.leftVal, test.rightVal, result, test.expected)
		}
	}
}

// TestNotEqualsPolicyCheckOperator tests the NotEqualsPolicyCheckOperator
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
		result := NotEqualsPolicyCheckOperator(test.leftVal, test.rightVal)
		if result != test.expected {
			t.Errorf("NotEqualsPolicyCheckOperator(%v, %v) = %v; want %v", test.leftVal, test.rightVal, result, test.expected)
		}
	}
}

// TestGreaterThanPolicyCheckOperator tests the GreaterThanPolicyCheckOperator
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
		result := GreaterThanPolicyCheckOperator(test.leftVal, test.rightVal)
		if result != test.expected {
			t.Errorf("GreaterThanPolicyCheckOperator(%v, %v) = %v; want %v", test.leftVal, test.rightVal, result, test.expected)
		}
	}
}

// TestGreaterThanOrEqualsPolicyCheckOperator tests the GreaterThanOrEqualsPolicyCheckOperator
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
		result := GreaterThanOrEqualsPolicyCheckOperator(test.leftVal, test.rightVal)
		if result != test.expected {
			t.Errorf("GreaterThanOrEqualsPolicyCheckOperator(%v, %v) = %v; want %v", test.leftVal, test.rightVal, result, test.expected)
		}
	}
}

// TestLessThanPolicyCheckOperator tests the LessThanPolicyCheckOperator
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
		result := LessThanPolicyCheckOperator(test.leftVal, test.rightVal)
		if result != test.expected {
			t.Errorf("LessThanPolicyCheckOperator(%v, %v) = %v; want %v", test.leftVal, test.rightVal, result, test.expected)
		}
	}
}

// TestEvaluatePolicyCheckOperator tests the EvaluatePolicyCheckOperator function
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
		{">=", 2, 1, true},
		{"<", 1, 2, true},
		{"<", 1, 1, false},
	}

	for _, test := range tests {
		result := EvaluatePolicyCheckOperator(test.operator, test.leftVal, test.rightVal)
		if result != test.expected {
			t.Errorf("EvaluatePolicyCheckOperator(%v, %v, %v) = %v; want %v", test.operator, test.leftVal, test.rightVal, result, test.expected)
		}
	}
}

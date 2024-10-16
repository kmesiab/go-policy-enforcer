package go_policy_enforcer

// PolicyCheckOperator is a function type that accepts two values of any type
// and returns a boolean result based on a comparison of the two values.
type PolicyCheckOperator func(leftVal, rightVal any) bool

// EqualsPolicyCheckOperator checks if two values are equal.
// Returns true if leftVal is equal to rightVal.
var EqualsPolicyCheckOperator = func(leftVal, rightVal any) bool {
	return leftVal == rightVal
}

// NotEqualsPolicyCheckOperator checks if two values are not equal.
// Returns true if leftVal is not equal to rightVal.
var NotEqualsPolicyCheckOperator = func(leftVal, rightVal any) bool {
	return leftVal != rightVal
}

// GreaterThanPolicyCheckOperator checks if the left value is greater than the right value.
// Assumes both values are integers. Returns true if leftVal is greater than rightVal.
var GreaterThanPolicyCheckOperator = func(leftVal, rightVal any) bool {
	return leftVal.(int) > rightVal.(int)
}

// GreaterThanOrEqualsPolicyCheckOperator checks if the left value is greater than or equal to the right value.
// Assumes both values are integers. Returns true if leftVal is greater than or equal to rightVal.
var GreaterThanOrEqualsPolicyCheckOperator = func(leftVal, rightVal any) bool {
	return leftVal.(int) >= rightVal.(int)
}

// LessThanPolicyCheckOperator checks if the left value is less than the right value.
// Assumes both values are integers. Returns true if leftVal is less than rightVal.
var LessThanPolicyCheckOperator = func(leftVal, rightVal any) bool {
	return leftVal.(int) < rightVal.(int)
}

// policyCheckOperatorMap maps string representations of comparison operators
// to their corresponding PolicyCheckOperator functions.
var policyCheckOperatorMap = map[string]PolicyCheckOperator{
	"==": EqualsPolicyCheckOperator,
	"!=": NotEqualsPolicyCheckOperator,
	">":  GreaterThanPolicyCheckOperator,
	">=": GreaterThanOrEqualsPolicyCheckOperator,
	"<":  LessThanPolicyCheckOperator,
}

// GetPolicyCheckOperator retrieves the appropriate PolicyCheckOperator function
// based on the provided operator string.
func GetPolicyCheckOperator(operator string) PolicyCheckOperator {
	return policyCheckOperatorMap[operator]
}

// EvaluatePolicyCheckOperator takes a string operator, a left value, and a right value,
// retrieves the corresponding PolicyCheckOperator function, and evaluates it with the given values.
// Returns the result of the comparison as a boolean.
func EvaluatePolicyCheckOperator(operator string, leftVal, rightVal any) bool {
	return GetPolicyCheckOperator(operator)(leftVal, rightVal)
}

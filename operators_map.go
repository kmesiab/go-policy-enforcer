package go_policy_enforcer

import (
	"github.com/kmesiab/go-policy-enforcer/custom_operators"
)

// PolicyCheckOperatorMap maps string representations of comparison operators
// to their corresponding PolicyCheckOperator functions.
var policyCheckOperatorMap = map[string]PolicyCheckOperator[any]{
	"==":  PolicyCheckOperator[any](EqualsPolicyCheckOperator),
	"!=":  PolicyCheckOperator[any](NotEqualsPolicyCheckOperator),
	"===": PolicyCheckOperator[any](custom_operators.DeepEqualsPolicyCheckOperator),
	"!==": PolicyCheckOperator[any](custom_operators.NotDeepEqualsPolicyCheckOperator),
	">":   PolicyCheckOperator[any](GreaterThanPolicyCheckOperator),
	">=":  PolicyCheckOperator[any](GreaterThanOrEqualsPolicyCheckOperator),
	"<":   PolicyCheckOperator[any](LessThanPolicyCheckOperator),
	"<=":  PolicyCheckOperator[any](LessThanOrEqualsPolicyCheckOperator),
	"in":  PolicyCheckOperator[any](InPolicyCheckOperator),
}

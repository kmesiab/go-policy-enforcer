package go_policy_enforcer

import (
	"github.com/kmesiab/go-policy-enforcer/custom_operators"
)

// PolicyCheckOperatorMap maps string representations of comparison operators
// to their corresponding PolicyCheckOperator functions.
var PolicyCheckOperatorMap = map[string]PolicyCheckOperator{
	"==":  EqualsPolicyCheckOperator,
	"!=":  NotEqualsPolicyCheckOperator,
	"===": custom_operators.DeepEqualsPolicyCheckOperator,
	"!==": custom_operators.NotDeepEqualsPolicyCheckOperator,
	">":   GreaterThanPolicyCheckOperator,
	">=":  GreaterThanOrEqualsPolicyCheckOperator,
	"<":   LessThanPolicyCheckOperator,
	"<=":  LessThanOrEqualsPolicyCheckOperator,
	"in":  InPolicyCheckOperator,
}

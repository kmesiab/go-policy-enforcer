package go_policy_enforcer

import (
	"github.com/kmesiab/go-policy-enforcer/custom_operators"
)

// PolicyCheckOperatorMap maps string representations of comparison operators
// to their corresponding PolicyCheckOperator functions.
var policyCheckOperatorMap = map[string]PolicyCheckOperator[any]{
	"==":  PolicyCheckOperator[any](equalsPolicyCheckOperator),
	"!=":  PolicyCheckOperator[any](notEqualsPolicyCheckOperator),
	"===": PolicyCheckOperator[any](custom_operators.DeepEqualsPolicyCheckOperator),
	"!==": PolicyCheckOperator[any](custom_operators.NotDeepEqualsPolicyCheckOperator),
	">":   PolicyCheckOperator[any](greaterThanPolicyCheckOperator),
	">=":  PolicyCheckOperator[any](greaterThanOrEqualsPolicyCheckOperator),
	"<":   PolicyCheckOperator[any](lessThanPolicyCheckOperator),
	"<=":  PolicyCheckOperator[any](lessThanOrEqualsPolicyCheckOperator),
	"in":  PolicyCheckOperator[any](inPolicyCheckOperator),
}

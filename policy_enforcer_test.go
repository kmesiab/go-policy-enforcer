package go_policy_enforcer

import (
	"testing"
)

func TestPolicyEnforcer_Enforce_HappyPath(t *testing.T) {
	// Mock policies that should pass
	policies := &[]Policy{
		{
			Name: "TestPolicy1",
			Rules: []Rule{
				{Field: "Age", Operator: "==", Value: 30},
			},
		},
		{
			Name: "TestPolicy2",
			Rules: []Rule{
				{Field: "Active", Operator: "==", Value: true},
			},
		},
	}

	// Create the resource that should pass all policies
	resource := struct {
		Age    int
		Active bool
	}{
		Age:    30,
		Active: true,
	}

	// Create the policy enforcer
	enforcer := NewPolicyEnforcer(policies)

	// Test Enforce
	if !enforcer.Enforce(resource) {
		t.Errorf("expected policy enforcement to return true, but got false")
	}
}

func TestPolicyEnforcer_Enforce_FailsPolicy(t *testing.T) {
	// Mock policies, one of which should fail
	policies := &[]Policy{
		{
			Name: "TestPolicy1",
			Rules: []Rule{
				{Field: "Age", Operator: "==", Value: 30},
			},
		},
		{
			Name: "TestPolicy2",
			Rules: []Rule{
				{Field: "Active", Operator: "==", Value: true},
			},
		},
	}

	// Create the resource that should pass all policies
	resource := struct {
		Age    int
		Active bool
	}{
		Age:    25,
		Active: true,
	}

	// Create the policy enforcer
	enforcer := NewPolicyEnforcer(policies)

	// Test Enforce
	if enforcer.Enforce(resource) {
		t.Errorf("expected policy enforcement to return false, but got true")
	}
}

func TestNewPolicyEnforcer(t *testing.T) {
	// Mock policies
	policies := &[]Policy{
		{Name: "TestPolicy"},
	}

	// Test that NewPolicyEnforcer returns a valid PolicyEnforcerInterface
	enforcer := NewPolicyEnforcer(policies)
	if enforcer == nil {
		t.Errorf("expected non-nil PolicyEnforcerInterface, but got nil")
	}
}

func TestPolicyEnforcer_Enforce_EmptyPolicies(t *testing.T) {
	// No policies
	policies := &[]Policy{}

	// Create the policy enforcer
	enforcer := NewPolicyEnforcer(policies)

	// Create the resource that should pass all policies
	resource := struct {
		Age    int
		Active bool
	}{
		Age:    30,
		Active: false,
	}

	// Test Enforce
	// With no policies, it should return false because no policy is enforcing
	if enforcer.Enforce(resource) {
		t.Errorf("expected policy enforcement to return false, but got true")
	}
}

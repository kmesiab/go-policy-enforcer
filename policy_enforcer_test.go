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

func TestPolicyEnforcer_Enforce_PartialMatch(t *testing.T) {
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

	resource := struct {
		Age    int
		Active bool
	}{
		Age:    30,    // Matches
		Active: false, // Fails
	}

	enforcer := NewPolicyEnforcer(policies)

	// Test Enforce should return false because one policy fails
	if enforcer.Enforce(resource) {
		t.Errorf("expected policy enforcement to return false due to partial match, but got true")
	}
}

func TestPolicyEnforcer_Enforce_NoMatch(t *testing.T) {
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

	resource := struct {
		Age    int
		Active bool
	}{
		Age:    25,    // Does not match
		Active: false, // Does not match
	}

	enforcer := NewPolicyEnforcer(policies)

	// Test Enforce should return false because resource does not match any policy
	if enforcer.Enforce(resource) {
		t.Errorf("expected policy enforcement to return false due to no match, but got true")
	}
}

func TestPolicyEnforcer_Enforce_MatchesAllPolicies(t *testing.T) {
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

	resource := struct {
		Age    int
		Active bool
	}{
		Age:    30,
		Active: true,
	}

	enforcer := PolicyEnforcer{
		Policies: policies,
	}

	// Act
	result := enforcer.Enforce(resource)

	// Assert
	if !result {
		t.Errorf("expected Enforce to return true when resource matches all policies, but got false")
	}
}

func TestPolicyEnforcer_Enforce_MatchesSomePolicies(t *testing.T) {
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

	resource := struct {
		Age    int
		Active bool
	}{
		Age:    30,    // Matches
		Active: false, // Does not match
	}

	enforcer := PolicyEnforcer{
		Policies: policies,
	}

	// Act
	result := enforcer.Enforce(resource)

	// Assert
	if result {
		t.Errorf("expected Enforce to return false when resource matches only some policies, but got true")
	}
}

func TestPolicyEnforcer_Enforce_NilResource(t *testing.T) {
	// Arrange
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

	enforcer := PolicyEnforcer{
		Policies: policies,
	}

	// Act
	result := enforcer.Enforce(nil)

	// Assert
	if result {
		t.Errorf("expected Enforce to return false when resource is nil, but got true")
	}
}

func TestPolicyEnforcer_Enforce_NilPolicies(t *testing.T) {
	// Arrange
	enforcer := PolicyEnforcer{
		Policies: nil,
	}

	resource := struct {
		Age    int
		Active bool
	}{
		Age:    30,
		Active: true,
	}

	// Act
	result := enforcer.Enforce(resource)

	// Assert
	if result {
		t.Errorf("expected Enforce to return false when policies are nil, but got true")
	}
}

func TestPolicyEnforcer_Enforce_NilRules(t *testing.T) {
	// Arrange
	policies := &[]Policy{
		{
			Name:  "TestPolicy1",
			Rules: nil, // Nil rules
		},
		{
			Name: "TestPolicy2",
			Rules: []Rule{
				{Field: "Active", Operator: "==", Value: true},
			},
		},
	}

	resource := struct {
		Active bool
	}{
		Active: true,
	}

	enforcer := PolicyEnforcer{
		Policies: policies,
	}

	// Act
	result := enforcer.Enforce(resource)

	// Assert
	// Since the first policy has nil rules, it should not affect the enforcement result.
	// The enforcement result should be based on the second policy.
	// In this case, the resource matches the second policy, so the result should be true.
	if !result {
		t.Errorf("expected Enforce to return true when resource matches a policy with nil rules, but got false")
	}
}

func TestPolicyEnforcer_Enforce_EmptyRules(t *testing.T) {
	// Arrange
	policies := &[]Policy{
		{
			Name:  "TestPolicy1",
			Rules: nil, // Empty rules
		},
		{
			Name: "TestPolicy2",
			Rules: []Rule{
				{Field: "Active", Operator: "==", Value: true},
			},
		},
	}

	resource := struct {
		Active bool
	}{
		Active: true,
	}

	enforcer := PolicyEnforcer{
		Policies: policies,
	}

	// Act
	result := enforcer.Enforce(resource)

	// Assert
	// Since the first policy has empty rules, it should not affect the enforcement result.
	// The enforcement result should be based on the second policy.
	// In this case, the resource matches the second policy, so the result should be true.
	if !result {
		t.Errorf("expected Enforce to return true when resource matches a policy with empty rules, but got false")
	}
}

func TestPolicyEnforcer_Match_NestedRules(t *testing.T) {
	// Arrange
	policies := &[]Policy{
		{
			Name: "TestPolicy1",
			Rules: []Rule{
				{Field: "Age", Operator: ">", Value: 25},
				{
					Field:    "Nested",
					Operator: "any",
					Value: []Rule{
						{Field: "Status", Operator: "==", Value: "active"},
						{Field: "Status", Operator: "==", Value: "inactive"},
					},
				},
			},
		},
		{
			Name: "TestPolicy2",
			Rules: []Rule{
				{Field: "Active", Operator: "==", Value: true},
			},
		},
	}

	// Set the Active field to false to avoid matching TestPolicy2
	resource := struct {
		Age    int
		Nested []struct {
			Status string
		}
		Active bool
	}{
		Age: 30,
		Nested: []struct {
			Status string
		}{
			{Status: "active"},
			{Status: "inactive"},
		},
		Active: false, // This prevents TestPolicy2 from matching
	}

	enforcer := PolicyEnforcer{
		Policies: policies,
	}

	// Act
	matchedPolicies := enforcer.Match(resource)

	// Assert
	expectedMatchedPoliciesCount := 1
	if len(matchedPolicies) != expectedMatchedPoliciesCount {
		t.Errorf("expected %d matched policies, but got %d", expectedMatchedPoliciesCount, len(matchedPolicies))
	}

	expectedMatchedPolicyName := "TestPolicy1"
	if matchedPolicies[0].Name != expectedMatchedPolicyName {
		t.Errorf("expected matched policy name to be '%s', but got '%s'", expectedMatchedPolicyName, matchedPolicies[0].Name)
	}
}

func TestPolicyEnforcer_Match_MatchesSomePoliciesWithNestedRules(t *testing.T) {
	// Arrange
	policies := &[]Policy{
		{
			Name: "TestPolicy1",
			Rules: []Rule{
				{Field: "Age", Operator: ">", Value: 25},
				{
					Field:    "Nested",
					Operator: "any",
					Value: []Rule{
						{Field: "Status", Operator: "==", Value: "active"},
						{Field: "Status", Operator: "==", Value: "inactive"},
					},
				},
			},
		},
		{
			Name: "TestPolicy2",
			Rules: []Rule{
				{Field: "Active", Operator: "==", Value: true},
			},
		},
	}

	resource := struct {
		Age    int
		Nested []struct {
			Status string
		}
		Active bool
	}{
		Age: 30,
		Nested: []struct {
			Status string
		}{
			{Status: "active"},
			{Status: "inactive"},
		},
		Active: false, // Does not match TestPolicy2
	}

	enforcer := PolicyEnforcer{
		Policies: policies,
	}

	// Act
	matchedPolicies := enforcer.Match(resource)

	// Assert
	expectedMatchedPoliciesCount := 1
	if len(matchedPolicies) != expectedMatchedPoliciesCount {
		t.Errorf("expected %d matched policies, but got %d", expectedMatchedPoliciesCount, len(matchedPolicies))
	}

	expectedMatchedPolicyName := "TestPolicy1"
	if matchedPolicies[0].Name != expectedMatchedPolicyName {
		t.Errorf("expected matched policy name to be '%s', but got '%s'", expectedMatchedPolicyName, matchedPolicies[0].Name)
	}
}

func TestPolicyEnforcer_Match_NoMatchesWithNestedRules(t *testing.T) {
	// Arrange
	policies := &[]Policy{
		{
			Name: "TestPolicy1",
			Rules: []Rule{
				{Field: "Age", Operator: ">", Value: 25},
				{
					Field:    "Nested",
					Operator: "any",
					Value: []Rule{
						{Field: "Status", Operator: "==", Value: "active"},
						{Field: "Status", Operator: "==", Value: "inactive"},
					},
				},
			},
		},
		{
			Name: "TestPolicy2",
			Rules: []Rule{
				{Field: "Active", Operator: "==", Value: true},
			},
		},
	}

	resource := struct {
		Age    int
		Nested []struct {
			Status string
		}
		Active bool
	}{
		Age: 20, // Does not match TestPolicy1
		Nested: []struct {
			Status string
		}{
			{Status: "unknown"}, // Does not match any rule in TestPolicy1
		},
		Active: false, // Does not match TestPolicy2
	}

	enforcer := PolicyEnforcer{
		Policies: policies,
	}

	// Act
	matchedPolicies := enforcer.Match(resource)

	// Assert
	expectedMatchedPoliciesCount := 0
	if len(matchedPolicies) != expectedMatchedPoliciesCount {
		t.Errorf("expected %d matched policies, but got %d", expectedMatchedPoliciesCount, len(matchedPolicies))
	}
}

func TestPolicyEnforcer_Enforce_MatchesSomePoliciesWithMultipleOperators(t *testing.T) {
	// Arrange
	policies := &[]Policy{
		{
			Name: "TestPolicy1",
			Rules: []Rule{
				{Field: "Age", Operator: ">", Value: 25},
				{
					Field:    "Status",
					Operator: "in",
					Value:    []string{"active", "inactive"},
				},
			},
		},
		{
			Name: "TestPolicy2",
			Rules: []Rule{
				{Field: "Active", Operator: "==", Value: true},
			},
		},
	}

	resource := struct {
		Age    int
		Status string
		Active bool
	}{
		Age:    30,
		Status: "active",
		Active: false, // Does not match TestPolicy2
	}

	enforcer := PolicyEnforcer{
		Policies: policies,
	}

	// Act
	result := enforcer.Enforce(resource)

	// Assert
	if result {
		t.Errorf("expected Enforce to return false when resource matches only some policies with multiple operators, but got true")
	}
}

func TestPolicyEnforcer_Enforce_MatchesNoPoliciesWithMultipleOperators(t *testing.T) {
	// Arrange
	policies := &[]Policy{
		{
			Name: "TestPolicy1",
			Rules: []Rule{
				{Field: "Age", Operator: ">", Value: 25},
				{
					Field:    "Status",
					Operator: "in",
					Value:    []string{"active", "inactive"},
				},
			},
		},
		{
			Name: "TestPolicy2",
			Rules: []Rule{
				{Field: "Active", Operator: "==", Value: true},
			},
		},
	}

	resource := struct {
		Age    int
		Status string
		Active bool
	}{
		Age:    20,
		Status: "unknown",
		Active: false,
	}

	enforcer := PolicyEnforcer{
		Policies: policies,
	}

	// Act
	result := enforcer.Enforce(resource)

	// Assert
	if result {
		t.Errorf("expected Enforce to return false when resource matches no policies with multiple operators, but got true")
	}
}

func TestPolicyEnforcer_Match_NonExistentField(t *testing.T) {
	// Arrange
	policies := &[]Policy{
		{
			Name: "TestPolicy1",
			Rules: []Rule{
				{Field: "NonExistentField", Operator: "==", Value: "value"},
			},
		},
		{
			Name: "TestPolicy2",
			Rules: []Rule{
				{Field: "Active", Operator: "==", Value: true},
			},
		},
	}

	resource := struct {
		Active bool
	}{
		Active: true,
	}

	enforcer := PolicyEnforcer{
		Policies: policies,
	}

	// Act
	matchedPolicies := enforcer.Match(resource)

	// Assert
	// Since the resource does not have a field called "NonExistentField",
	// it should not match TestPolicy1.
	// Therefore, only TestPolicy2 should be in the matched policies slice.
	expectedMatchedPoliciesCount := 1
	if len(matchedPolicies) != expectedMatchedPoliciesCount {
		t.Errorf("expected %d matched policies, but got %d", expectedMatchedPoliciesCount, len(matchedPolicies))
	}

	expectedMatchedPolicyName := "TestPolicy2"
	if matchedPolicies[0].Name != expectedMatchedPolicyName {
		t.Errorf("expected matched policy name to be '%s', but got '%s'", expectedMatchedPolicyName, matchedPolicies[0].Name)
	}
}

func TestPolicyEnforcer_Enforce_MatchesNonExistentOperator(t *testing.T) {
	// Arrange
	policies := &[]Policy{
		{
			Name: "TestPolicy1",
			Rules: []Rule{
				{Field: "Age", Operator: "non-existent", Value: 25},
			},
		},
	}

	resource := struct {
		Age int
	}{
		Age: 25,
	}

	enforcer := PolicyEnforcer{
		Policies: policies,
	}

	// Act
	result := enforcer.Enforce(resource)

	// Assert
	// Since the operator "non-existent" does not exist, the resource should not match the policy.
	// Therefore, the result should be false.
	if result {
		t.Errorf("expected Enforce to return false when resource matches a policy with a non-existent operator, but got true")
	}
}

func TestPolicyEnforcer_Enforce_MatchesNonExistentValue(t *testing.T) {
	// Arrange
	policies := &[]Policy{
		{
			Name: "TestPolicy1",
			Rules: []Rule{
				{Field: "Status", Operator: "==", Value: "active"},
			},
		},
		{
			Name: "TestPolicy2",
			Rules: []Rule{
				{Field: "Status", Operator: "==", Value: "inactive"},
			},
		},
	}

	resource := struct {
		Status string
	}{
		Status: "unknown", // This status does not match any policy
	}

	enforcer := PolicyEnforcer{
		Policies: policies,
	}

	// Act
	result := enforcer.Enforce(resource)

	// Assert
	// Since the resource does not match any policy, the result should be false.
	if result {
		t.Errorf("expected Enforce to return false when resource matches a policy with a non-existent value, but got true")
	}
}

func TestPolicyEnforcer_Enforce_NilValue(t *testing.T) {
	// Arrange
	policies := &[]Policy{
		{
			Name: "TestPolicy1",
			Rules: []Rule{
				{Field: "Age", Operator: ">", Value: 25},
				{
					Field:    "Nested",
					Operator: "any",
					Value: []Rule{
						{Field: "Status", Operator: "==", Value: "active"},
						{Field: "Status", Operator: "==", Value: "inactive"},
					},
				},
			},
		},
		{
			Name: "TestPolicy2",
			Rules: []Rule{
				{Field: "Active", Operator: "==", Value: true},
			},
		},
	}

	// Set the Nested field to nil to match TestPolicy1
	resource := struct {
		Age    int
		Nested []struct {
			Status string
		}
		Active bool
	}{
		Age:    30,
		Nested: nil,
		Active: false, // This prevents TestPolicy2 from matching
	}

	enforcer := PolicyEnforcer{
		Policies: policies,
	}

	// Act
	result := enforcer.Enforce(resource)

	// Assert
	// Since the Nested field in the resource is nil, it should not affect the enforcement result.
	// The enforcement result should be based on the TestPolicy1.
	// In this case, the resource does not match TestPolicy1, so the result should be false.
	if result {
		t.Errorf("expected Enforce to return false when resource matches a policy with a nil value, but got true")
	}
}

func TestPolicyEnforcer_Match_MultipleOperators(t *testing.T) {
	// Arrange
	policies := &[]Policy{
		{
			Name: "TestPolicy1",
			Rules: []Rule{
				{Field: "Age", Operator: ">", Value: 25},
				{
					Field:    "Status",
					Operator: "in",
					Value:    []string{"active", "inactive"},
				},
			},
		},
		{
			Name: "TestPolicy2",
			Rules: []Rule{
				{Field: "Active", Operator: "==", Value: true},
			},
		},
	}

	resource := struct {
		Age    int
		Status string
		Active bool
	}{
		Age:    30,
		Status: "active",
		Active: true,
	}

	enforcer := NewPolicyEnforcer(policies)

	// Act
	matchedPolicies := enforcer.Match(resource)

	// Assert
	expectedMatchedPolicies := 2
	if len(matchedPolicies) != expectedMatchedPolicies {
		t.Errorf("Expected %d matched policies, but got %d", expectedMatchedPolicies, len(matchedPolicies))
	}
}

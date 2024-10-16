package go_policy_enforcer

import (
	"encoding/json"
	"os"
	"testing"
)

func TestPolicy_Evaluate_HappyPath(t *testing.T) {
	policy := Policy{
		Name: "TestPolicy",
		Rules: []Rule{
			{Field: "Age", Operator: "==", Value: 30},
			{Field: "Active", Operator: "==", Value: true},
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

	if !policy.Evaluate(resource) {
		t.Errorf("expected policy evaluation to return true, but got false")
	}
}

func TestPolicy_Evaluate_NegativePath(t *testing.T) {
	policy := Policy{
		Name: "TestPolicy",
		Rules: []Rule{
			{Field: "Age", Operator: "==", Value: 30},
			{Field: "Active", Operator: "==", Value: true},
		},
	}

	// Create the resource that should pass all policies
	resource := struct {
		Age    int
		Active bool
	}{
		Age:    30,
		Active: false,
	}

	if policy.Evaluate(resource) {
		t.Errorf("expected policy evaluation to return false, but got true")
	}
}

func TestLoadPolicy_HappyPath(t *testing.T) {
	// Create a mock policy file
	mockPolicy := Policy{
		Name: "TestPolicy",
		Rules: []Rule{
			{Field: "Age", Operator: "==", Value: 30},
		},
	}

	mockPolicyData, _ := json.Marshal(mockPolicy)
	mockPolicyFile := "test_policy.json"

	if err := os.WriteFile(mockPolicyFile, mockPolicyData, 0644); err != nil {
		t.Fatalf("failed to create mock policy file: %v", err)
	}

	defer func(name string) {
		err := os.Remove(name)
		if err != nil {
			t.Errorf("failed to remove mock policy file: %v", err)
		}
	}(mockPolicyFile) // Clean up after test

	// Load the policy
	policy, err := LoadPolicy(mockPolicyFile)
	if err != nil {
		t.Errorf("unexpected error loading policy: %v", err)
	}

	// Validate loaded policy
	if policy.Name != mockPolicy.Name {
		t.Errorf("expected policy name %s, got %s", mockPolicy.Name, policy.Name)
	}
}

func TestLoadPolicy_FileNotFound(t *testing.T) {
	_, err := LoadPolicy("non_existent_policy.json")
	if err == nil {
		t.Errorf("expected error when loading non-existent policy file, but got none")
	}
}

func TestLoadPolicy_InvalidJSON(t *testing.T) {
	// Create a mock invalid policy file
	invalidPolicyFile := "invalid_policy.json"
	if err := os.WriteFile(invalidPolicyFile, []byte(`invalid-json`), 0644); err != nil {
		t.Fatalf("failed to create mock invalid policy file: %v", err)
	}

	defer func(name string) {
		err := os.Remove(name)
		if err != nil {
			t.Errorf("failed to remove mock invalid policy file: %v", err)
		}
	}(invalidPolicyFile) // Clean up after test

	_, err := LoadPolicy(invalidPolicyFile)
	if err == nil {
		t.Errorf("expected error when loading invalid JSON policy file, but got none")
	}
}

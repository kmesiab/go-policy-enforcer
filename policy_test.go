package go_policy_enforcer

import (
	"encoding/json"
	"os"
	"reflect"
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

	// Create the resource that should fail the second rule
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

func TestPolicy_Match_NestedStructs(t *testing.T) {
	nestedPolicy := Policy{
		Name: "NestedPolicy",
		Rules: []Rule{
			{Field: "NestedField.NestedField", Operator: "==", Value: "nestedValue"},
		},
	}

	parentPolicy := Policy{
		Name: "ParentPolicy",
		Rules: []Rule{
			{Field: "ParentField", Operator: "==", Value: "parentValue"},
		},
	}

	policies := &[]Policy{nestedPolicy, parentPolicy}
	enforcer := NewPolicyEnforcer(policies)

	resource := struct {
		ParentField string
		NestedField struct {
			NestedField string
		}
	}{
		ParentField: "parentValue",
		NestedField: struct {
			NestedField string
		}{
			NestedField: "nestedValue",
		},
	}

	matchedPolicies := enforcer.Match(resource)

	if len(matchedPolicies) != 2 {
		t.Fatalf("expected 2 matched policies, but got %d", len(matchedPolicies))
	}

	expectedPolicyNames := []string{nestedPolicy.Name, parentPolicy.Name}
	for _, policy := range matchedPolicies {
		found := false
		for _, expectedName := range expectedPolicyNames {
			if policy.Name == expectedName {
				found = true
				break
			}
		}
		if !found {
			t.Errorf("expected matched policy to be either %s or %s, but got %s", nestedPolicy.Name, parentPolicy.Name, policy.Name)
		}
	}
}

func TestPolicy_Evaluate_NestedField_MultipleLevels(t *testing.T) {
	policy := Policy{
		Name: "TestPolicy",
		Rules: []Rule{
			{Field: "Level1.Level2.Field", Operator: "==", Value: "nestedValue"},
		},
	}

	resource := struct {
		Level1 struct {
			Level2 struct {
				Field string
			}
		}
	}{
		Level1: struct {
			Level2 struct {
				Field string
			}
		}{
			Level2: struct {
				Field string
			}{
				Field: "nestedValue",
			},
		},
	}

	if !policy.Evaluate(resource) {
		t.Errorf("expected policy evaluation to return true for nested field, but got false")
	}
}

func TestPolicy_Evaluate_NestedField_MixedTypes(t *testing.T) {
	policy := Policy{
		Name: "TestPolicy",
		Rules: []Rule{
			{Field: "Level1.Level2.Field", Operator: "==", Value: "nestedValue"},
			{Field: "Level1.PrimitiveField", Operator: "==", Value: 100},
		},
	}

	resource := struct {
		Level1 struct {
			Level2 struct {
				Field string
			}
			PrimitiveField int
		}
	}{
		Level1: struct {
			Level2 struct {
				Field string
			}
			PrimitiveField int
		}{
			Level2: struct {
				Field string
			}{
				Field: "nestedValue",
			},
			PrimitiveField: 100,
		},
	}

	if !policy.Evaluate(resource) {
		t.Errorf("expected policy evaluation to return true for nested field, but got false")
	}
}

func TestGetNestedField_NonExistentField(t *testing.T) {
	type TestStruct struct {
		Field1 string
		Field2 struct {
			Field3 string
		}
	}

	testStruct := TestStruct{
		Field1: "value1",
		Field2: struct {
			Field3 string
		}{
			Field3: "value2",
		},
	}

	v := reflect.ValueOf(testStruct)
	fieldPath := "NonExistentField"

	_, err := getNestedField(v, fieldPath)
	if err == nil {
		t.Errorf("expected error when accessing non-existent field %s, but got none", fieldPath)
	}
}

func TestGetNestedField_PointerToStruct(t *testing.T) {
	type TestStruct struct {
		Field1 string
		Field2 *struct {
			Field3 string
		}
	}

	testStruct := TestStruct{
		Field1: "value1",
		Field2: &struct {
			Field3 string
		}{
			Field3: "value2",
		},
	}

	v := reflect.ValueOf(testStruct)
	fieldPath := "Field2.Field3"

	nestedField, err := getNestedField(v, fieldPath)
	if err != nil {
		t.Fatalf("unexpected error accessing nested field %s: %v", fieldPath, err)
	}

	expectedValue := "value2"
	actualValue := nestedField.Interface().(string)
	if actualValue != expectedValue {
		t.Errorf("expected nested field value %s, got %s", expectedValue, actualValue)
	}
}

func TestGetNestedField_PointerToPrimitiveType(t *testing.T) {
	type TestStruct struct {
		Field1 string
		Field2 *int
	}

	testStruct := TestStruct{
		Field1: "value1",
		Field2: func() *int {
			value := 100
			return &value
		}(),
	}

	v := reflect.ValueOf(testStruct)
	fieldPath := "Field2"

	nestedField, err := getNestedField(v, fieldPath)
	if err != nil {
		t.Fatalf("unexpected error accessing nested field %s: %v", fieldPath, err)
	}

	expectedValue := 100
	actualValue := nestedField.Interface().(*int)
	if *actualValue != expectedValue {
		t.Errorf("expected nested field value %d, got %d", expectedValue, *actualValue)
	}
}

func TestGetNestedField_SliceOfStructs(t *testing.T) {
	type InnerStruct struct {
		InnerField string
	}

	type TestStruct struct {
		OuterField []InnerStruct
	}

	testStruct := TestStruct{
		OuterField: []InnerStruct{
			{InnerField: "innerValue1"},
			{InnerField: "innerValue2"},
		},
	}

	v := reflect.ValueOf(testStruct)
	// Specify the index explicitly
	fieldPath := "OuterField[0].InnerField"

	nestedField, err := getNestedField(v, fieldPath)
	if err != nil {
		t.Fatalf("unexpected error accessing nested field %s: %v", fieldPath, err)
	}

	expectedValue := "innerValue1"
	actualValue := nestedField.Interface().(string)
	if actualValue != expectedValue {
		t.Errorf("expected nested field value %s, got %s", expectedValue, actualValue)
	}
}

func TestGetNestedField_SliceOfPrimitiveTypes(t *testing.T) {
	type TestStruct struct {
		Field1 string
		Field2 []int
	}

	testStruct := TestStruct{
		Field1: "value1",
		Field2: []int{100, 200, 300},
	}

	v := reflect.ValueOf(testStruct)
	fieldPath := "Field2[1]"

	nestedField, err := getNestedField(v, fieldPath)
	if err != nil {
		t.Fatalf("unexpected error accessing nested field %s: %v", fieldPath, err)
	}

	expectedValue := 200
	actualValue := nestedField.Interface().(int)
	if actualValue != expectedValue {
		t.Errorf("expected nested field value %d, got %d", expectedValue, actualValue)
	}
}

func TestGetNestedField_MapOfStructs(t *testing.T) {
	type InnerStruct struct {
		InnerField string
	}

	type TestStruct struct {
		OuterField map[string]InnerStruct
	}

	testStruct := TestStruct{
		OuterField: map[string]InnerStruct{
			"key1": {InnerField: "innerValue1"},
			"key2": {InnerField: "innerValue2"},
		},
	}

	v := reflect.ValueOf(testStruct)
	fieldPath := "OuterField.key1.InnerField"

	nestedField, err := getNestedField(v, fieldPath)
	if err != nil {
		t.Fatalf("unexpected error accessing nested field %s: %v", fieldPath, err)
	}

	expectedValue := "innerValue1"
	actualValue := nestedField.Interface().(string)
	if actualValue != expectedValue {
		t.Errorf("expected nested field value %s, got %s", expectedValue, actualValue)
	}
}

func TestGetNestedField_MapOfPrimitiveTypes(t *testing.T) {
	type TestStruct struct {
		Field1 string
		Field2 map[string]int
	}

	testStruct := TestStruct{
		Field1: "value1",
		Field2: map[string]int{
			"key1": 100,
			"key2": 200,
		},
	}

	v := reflect.ValueOf(testStruct)
	fieldPath := "Field2.key1"

	nestedField, err := getNestedField(v, fieldPath)
	if err != nil {
		t.Fatalf("unexpected error accessing nested field %s: %v", fieldPath, err)
	}

	expectedValue := 100
	actualValue := nestedField.Interface().(int)
	if actualValue != expectedValue {
		t.Errorf("expected nested field value %d, got %d", expectedValue, actualValue)
	}
}

func TestGetNestedField_UnexportedFields(t *testing.T) {
	type InnerStruct struct {
		innerField string // Unexported field
	}

	type TestStruct struct {
		OuterField InnerStruct // Exported field
	}

	testStruct := TestStruct{
		OuterField: InnerStruct{
			innerField: "innerValue",
		},
	}

	v := reflect.ValueOf(testStruct)
	fieldPath := "OuterField.innerField" // Accessing exported outer field, but unexported inner field

	_, err := getNestedField(v, fieldPath)
	if err == nil {
		t.Fatalf("expected error when accessing unexported field %s, but got none", fieldPath)
	}

	expectedError := "field innerField is unexported and cannot be accessed"
	if err.Error() != expectedError {
		t.Errorf("expected error %s, but got %s", expectedError, err.Error())
	}
}

func TestPolicy_Evaluate_ComplexResource(t *testing.T) {
	type InnerStruct struct {
		InnerField string
	}

	type TestStruct struct {
		OuterField   []InnerStruct
		MapField     map[string]int
		PointerField *int
	}

	policy := Policy{
		Name: "ComplexPolicy",
		Rules: []Rule{
			{Field: "OuterField[0].InnerField", Operator: "==", Value: "value1"},
			{Field: "MapField.key1", Operator: ">=", Value: 100},
			{Field: "PointerField", Operator: "<", Value: 200},
		},
	}

	resource := TestStruct{
		OuterField: []InnerStruct{
			{InnerField: "value1"},
			{InnerField: "value2"},
		},
		MapField: map[string]int{
			"key1": 100,
			"key2": 200,
		},
		PointerField: func() *int {
			value := 150
			return &value
		}(),
	}

	if !policy.Evaluate(resource) {
		t.Error("Policy evaluation failed for complex resource")
	}
}

func TestPolicy_Evaluate_MixedTypesInNestedFields(t *testing.T) {
	type InnerStruct struct {
		InnerField string
		Primitive  int
	}

	type TestStruct struct {
		OuterField InnerStruct
	}

	policy := Policy{
		Name: "MixedTypesPolicy",
		Rules: []Rule{
			{Field: "OuterField.InnerField", Operator: "==", Value: "value1"},
			{Field: "OuterField.Primitive", Operator: ">=", Value: 100},
		},
	}

	resource := TestStruct{
		OuterField: InnerStruct{
			InnerField: "value1",
			Primitive:  150,
		},
	}

	if !policy.Evaluate(resource) {
		t.Error("Policy evaluation failed for mixed types in nested fields")
	}
}

func TestPolicy_Evaluate_NonExistentField(t *testing.T) {
	type InnerStruct struct {
		InnerField string
	}

	type TestStruct struct {
		OuterField InnerStruct
	}

	policy := Policy{
		Name: "NonExistentFieldPolicy",
		Rules: []Rule{
			{Field: "OuterField.NonExistentField", Operator: "==", Value: "value1"},
		},
	}

	resource := TestStruct{
		OuterField: InnerStruct{
			InnerField: "value2",
		},
	}

	if policy.Evaluate(resource) {
		t.Error("Policy evaluation should fail for non-existent field in nested struct")
	}
}

package go_policy_enforcer

import (
	"encoding/json"
	"fmt"
	"os"
	"reflect"
	"strconv"
	"strings"
)

// Policy represents a set of rules that define access control or behavior.
// It is used to enforce policies on resources, such as data or actions.
//
// The Policy struct has the following fields:
// - Name: A string representing the name of the policy.
// - Rules: A slice of Rule structs representing the rules that define the policy.
type Policy struct {
	Name  string
	Rules []Rule
}

// Evaluate checks if the given resource adheres to the policy's rules.
// The resource must be a struct, and its fields are evaluated against the policy's rules.
// If any rule fails, the function returns false. Otherwise, it returns true.
//
// Parameters:
// - resource: The resource to be evaluated. It must be a struct.
//
// Return:
// - bool: Returns true if the resource adheres to all policy rules, false otherwise.
func (p *Policy) Evaluate(resource any) bool {
	v := reflect.ValueOf(resource)

	// Handle pointers by dereferencing them
	if v.Kind() == reflect.Ptr {
		v = v.Elem()
	}

	// Ensure we're working with a struct
	if v.Kind() != reflect.Struct {
		fmt.Println("The resource provided is not a struct")
		return false
	}

	for _, rule := range p.Rules {
		// Handle nested rules
		if nestedRules, ok := rule.Value.([]Rule); ok {
			fieldValue, err := getNestedField(v, rule.Field)
			if err != nil {
				fmt.Println(err)
				return false
			}

			if fieldValue.Kind() == reflect.Slice {
				matched := false
				for i := 0; i < fieldValue.Len(); i++ {
					elem := fieldValue.Index(i).Interface()
					for _, nestedRule := range nestedRules {

						// Perform the pattern match
						if EvaluatePolicyCheckOperator(
							nestedRule.Operator,
							reflect.ValueOf(elem).FieldByName(nestedRule.Field).Interface(),
							nestedRule.Value,
						) {

							matched = true
							break
						}
					}
					if matched {
						break
					}
				}
				if !matched {
					return false
				}
				continue
			}
		}

		// Handle regular policy checks
		fieldValue, err := getNestedField(v, rule.Field)
		if err != nil {
			fmt.Println(err)
			return false
		}

		if !fieldValue.CanInterface() {
			fmt.Printf("Field %s is unexported and cannot be accessed\n", rule.Field)
			return false
		}

		if !EvaluatePolicyCheckOperator(rule.Operator, fieldValue.Interface(), rule.Value) {
			return false
		}
	}

	return true
}

func getNestedField(v reflect.Value, fieldPath string) (reflect.Value, error) {
	fields := strings.Split(fieldPath, ".")
	for i, field := range fields {
		// Handle map access
		if v.Kind() == reflect.Map {
			key := reflect.ValueOf(field)
			v = v.MapIndex(key)
			if !v.IsValid() {
				return reflect.Value{}, fmt.Errorf("key %s not found in map", field)
			}
		} else {
			// Handle slice indexing (e.g., Field[0])
			if strings.Contains(field, "[") && strings.Contains(field, "]") {
				parts := strings.Split(field, "[")
				fieldName := parts[0]
				indexStr := strings.TrimSuffix(parts[1], "]")

				// Get the field by name
				v = v.FieldByName(fieldName)
				if !v.IsValid() {
					return reflect.Value{}, fmt.Errorf("field %s not found", fieldName)
				}

				// Ensure it's a slice or array
				if v.Kind() != reflect.Slice && v.Kind() != reflect.Array {
					return reflect.Value{}, fmt.Errorf("field %s is not a slice or array", fieldName)
				}

				// Convert the index to an integer
				index, err := strconv.Atoi(indexStr)
				if err != nil {
					return reflect.Value{}, fmt.Errorf("invalid slice index %s", indexStr)
				}

				// Check if the index is within bounds
				if index >= v.Len() {
					return reflect.Value{}, fmt.Errorf("index %d out of bounds for slice %s", index, fieldName)
				}

				// Get the indexed value
				v = v.Index(index)
			} else {
				// Regular struct field access
				v = v.FieldByName(field)
				if !v.IsValid() {
					return reflect.Value{}, fmt.Errorf("field %s not found", field)
				}
			}
		}

		// Check if the field is exported (CanInterface returns false for unexported fields)
		if !v.CanInterface() {
			return reflect.Value{}, fmt.Errorf("field %s is unexported and cannot be accessed", field)
		}

		// Only dereference if this is not the last field in the path
		if v.Kind() == reflect.Ptr && !v.IsNil() && i < len(fields)-1 {
			v = v.Elem()
		}
	}
	return v, nil
}

func LoadPolicy(policyFile string) (*Policy, error) {
	var (
		err          error
		policyString []byte
		policy       = &Policy{}
	)

	// Can't read the file
	if policyString, err = os.ReadFile(policyFile); err != nil {
		return nil, fmt.Errorf("failed to read policy file: %v", err)
	}

	// Invalid policy JSON
	if err = json.Unmarshal(policyString, policy); err != nil {
		return nil, fmt.Errorf("invalid policy json: %v", err)
	}

	return policy, nil
}

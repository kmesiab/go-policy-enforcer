package go_policy_enforcer

import (
	"encoding/json"
	"fmt"
	"os"
	"reflect"
)

type Policy struct {
	Name  string
	Rules []Rule
}

func (p *Policy) Evaluate(resource any) bool {
	v := reflect.ValueOf(resource)

	// Handle pointers by dereferencing them to get the underlying value
	if v.Kind() == reflect.Ptr {
		v = v.Elem()
	}

	// If the underlying value is not a struct, reject it
	if v.Kind() != reflect.Struct {
		fmt.Println("The resource provided is not a struct")
		return false
	}

	// Iterate over the rules and check the field values
	for _, rule := range p.Rules {
		// Get the field by name
		field := v.FieldByName(rule.Field)

		// If the field doesn't exist, report it
		if !field.IsValid() {
			fmt.Printf("Field %s not found in the struct\n", rule.Field)
			return false
		}

		// Check if the field is exported and can be accessed
		if !field.CanInterface() {
			fmt.Printf("Field %s is unexported and cannot be accessed\n", rule.Field)
			return false
		}

		// Compare the field value with the rule value
		if !EvaluatePolicyCheckOperator(rule.Operator, field.Interface(), rule.Value) {
			return false
		}
	}

	return true
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

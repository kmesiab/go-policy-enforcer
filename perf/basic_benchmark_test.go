package perf

import (
	"testing"

	pe "github.com/kmesiab/go-policy-enforcer"
)

type SimpleStruct struct {
	Field int
}

type ComplexStruct struct {
	Field SimpleStruct
}

func BenchmarkNestedStruct(b *testing.B) {

	testPolicies := &[]pe.Policy{
		{
			Name: "test",
			Rules: []pe.Rule{
				{
					Field:    "Field1",
					Operator: "==",
					Value:    1,
				},
			},
		},
	}

	var resource = ComplexStruct{
		Field: SimpleStruct{
			Field: 1,
		},
	}

	var enforcer = pe.NewPolicyEnforcer(testPolicies)

	for i := 0; i < b.N; i++ {
		_ = enforcer.Enforce(resource)
	}
}

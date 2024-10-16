package go_policy_enforcer

type PolicyEnforcerInterface interface {
	Enforce(resource any) bool
}

type PolicyEnforcer struct {
	PolicyEnforcerInterface
	Policies *[]Policy
}

// NewPolicyEnforcer creates a new instance of PolicyEnforcer with the provided
// policies.
//
// The function takes a pointer to a slice of Policy structs as input and returns
// a PolicyEnforcerInterface.
//
// The PolicyEnforcer struct contains a pointer to the slice of policies and
// implements the PolicyEnforcerInterface.
//
// Parameters:
// - policies: A pointer to a slice of Policy structs. Each Policy represents a
// set of rules or conditions that need to be enforced.
//
// Returns:
// - PolicyEnforcerInterface: An interface that provides the Enforce method to
// check if a resource complies with the policies.
func NewPolicyEnforcer(policies *[]Policy) PolicyEnforcerInterface {
	return PolicyEnforcer{
		Policies: policies,
	}
}

// Enforce checks if a given resource complies with all the policies.
//
// The function iterates over each policy in the PolicyEnforcer's policies slice.
//
// For each policy, it calls the Evaluate method with the provided resource.
//
// If the Evaluate method returns false, the function immediately returns false,
// indicating that the resource does not comply with the policy.
//
// If the Evaluate method returns true for all policies, the function returns true,
// indicating that the resource complies with all policies.
//
// Parameters:
// - resource: The resource to be evaluated against the policies. The type can be any valid Go type.
//
// Returns:
// - bool: A boolean value indicating whether the resource complies with all the policies.
//   - true: The resource complies with all the policies.
//   - false: The resource does not comply with at least one policy.
func (e PolicyEnforcer) Enforce(resource any) bool {

	// If there are no policies, return false immediately, as no policy can be enforced.
	if e.Policies == nil || len(*e.Policies) == 0 {
		return false
	}

	for _, p := range *e.Policies {
		if !p.Evaluate(resource) {
			return false
		}
	}
	return true
}

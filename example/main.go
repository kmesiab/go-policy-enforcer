package main

import (
	"fmt"
	"log"

	gopolicyenforcer "github.com/kmesiab/go-policy-enforcer"
)

const (
	finalizedPolicyExampleFile  = "./example/policies/finalized_policy.json"
	idRequiredPolicyExampleFile = "./example/policies/id_required_policy.json"
)

func main() {

	// Load a few policies from JSON files (./policies folder)
	policyList, err := loadPolicies()
	if err != nil {

		log.Fatalf("error loading policies: %s", err)
	}

	// Create assets to test enforcement
	allowedAsset := &Asset{ID: 1, Type: "asset", Finalized: true}
	deniedAsset := &Asset{ID: 2, Type: "asset", Finalized: false}
	assetList := []*Asset{allowedAsset, deniedAsset}

	// Create a PolicyEnforcer instance with the policies
	e := gopolicyenforcer.NewPolicyEnforcer(policyList)

	// Enforce the policies on the assets and print results
	for _, asset := range assetList {
		if e.Enforce(asset) {
			fmt.Printf("Asset %v is allowed\n", asset)
		} else {
			fmt.Printf("Asset %v is not allowed\n", asset)
		}
	}

}

func loadPolicies() (*[]gopolicyenforcer.Policy, error) {

	finalizedPolicy, err := gopolicyenforcer.LoadPolicy(finalizedPolicyExampleFile)

	if err != nil {
		return nil, err
	}

	idRequiredPolicy, err := gopolicyenforcer.LoadPolicy(idRequiredPolicyExampleFile)

	if err != nil {
		fmt.Println("Error loading policy:", err)

	}

	return &[]gopolicyenforcer.Policy{
		*finalizedPolicy,
		*idRequiredPolicy,
	}, nil

}

package main

import (
	"encoding/json"
	"fmt"
	"log"

	gopolicyenforcer "github.com/kmesiab/go-policy-enforcer"
)

const (
	finalizedPolicyExampleFile  = "./policies/finalized_policy.json"
	idRequiredPolicyExampleFile = "./policies/id_required_policy.json"
)

var (
	allowedAsset = &Asset{ID: 1, Type: "asset", Finalized: true}
	deniedAsset  = &Asset{ID: 2, Type: "asset", Finalized: false}
	assetList    = []*Asset{allowedAsset, deniedAsset}
)

type Asset struct {
	ID        int    `json:"id"`
	Type      string `json:"type"`
	Finalized bool   `json:"state"`
}

func main() {

	policyList, err := loadPolicies()
	if err != nil {
		log.Fatalf("error loading policies: %s", err)
	}

	testPolicies(policyList)
	matchPolicies(policyList)

}

func matchPolicies(policyList *[]gopolicyenforcer.Policy) {

	// Create a PolicyEnforcer instance with the policies
	e := gopolicyenforcer.NewPolicyEnforcer(policyList)

	fmt.Printf("---------------MATCHES-----------------\n")

	// Match the policies on the assets and print results
	for _, asset := range assetList {

		matchedPolicies := e.Match(asset)

		jsonString, _ := json.MarshalIndent(matchedPolicies, "", "  ")
		fmt.Printf("Matches: %s\n", jsonString)
	}

}

func testPolicies(policyList *[]gopolicyenforcer.Policy) {

	// Create a PolicyEnforcer instance with the policies
	e := gopolicyenforcer.NewPolicyEnforcer(policyList)

	// Enforce the policies on the assets and print results
	for _, asset := range assetList {

		if e.Enforce(asset) {

			fmt.Printf("Asset %d is allowed\n", asset.ID)
		} else {

			fmt.Printf("Asset %+v is not allowed\n", asset.ID)
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

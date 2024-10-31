package main

import (
	"fmt"

	pe "github.com/kmesiab/go-policy-enforcer"
)

type user struct {
	IsLoggedIn bool
	Roles      []string
}

var mustBeAdminRule = pe.Rule{
	Field:    "Roles",
	Operator: "in",
	Value:    "admin",
}

var mustBeLoggedInRule = pe.Rule{
	Field:    "IsLoggedIn",
	Operator: "==",
	Value:    true,
}

var canEditUserPolicy = pe.Policy{
	Rules: []pe.Rule{
		mustBeLoggedInRule,
		mustBeAdminRule,
	},
}

func main() {

	guest := user{IsLoggedIn: false}
	admin := user{IsLoggedIn: false, Roles: []string{"admin"}}
	loggedInAdmin := user{IsLoggedIn: true, Roles: []string{"admin"}}

	enforcer := pe.NewPolicyEnforcer(&[]pe.Policy{canEditUserPolicy})

	if enforcer.Enforce(guest) {
		fmt.Println("Guest can edit user")
	} else {
		fmt.Println("Guest cannot edit user")
	}

	if enforcer.Enforce(admin) {
		fmt.Println("Logged out Admin can edit user")
	} else {
		fmt.Println("Logged out Admin cannot edit user")
	}

	if enforcer.Enforce(loggedInAdmin) {
		fmt.Println("Logged in admin can edit user")
	} else {
		fmt.Println("Logged in admin cannot edit user")
	}
}

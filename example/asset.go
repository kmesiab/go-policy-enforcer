package main

type Asset struct {
	ID        string `json:"id"`
	Type      string `json:"type"`
	Finalized bool   `json:"state"`
}

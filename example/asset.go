package main

type Asset struct {
	ID        int    `json:"id"`
	Type      string `json:"type"`
	Finalized bool   `json:"state"`
}

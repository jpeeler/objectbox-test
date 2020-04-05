package main

import (
	"github.com/plaid/plaid-go/plaid"
)

//go:generate go run github.com/objectbox/objectbox-go/cmd/objectbox-gogen

type MyInstitution struct {
	ID uint64 `objectbox:"id"`
	plaid.Institution
}

type MyInstitions []MyInstitution


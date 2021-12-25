package main

import (
	"math/big"
	"time"
)

// Account holds the name and balance
type Account struct {
	Name    string
	Balance *big.Rat
}

type sortAccounts []*Account

func (s sortAccounts) Len() int      { return len(s) }
func (s sortAccounts) Swap(i, j int) { s[i], s[j] = s[j], s[i] }

type sortAccountsByName struct{ sortAccounts }

func (s sortAccountsByName) Less(i, j int) bool {
	return s.sortAccounts[i].Name < s.sortAccounts[j].Name
}

// Transaction is the basis of a ledger. The ledger holds a list of transactions.
// A Transaction has a Payee, Date (with no time, or to pu
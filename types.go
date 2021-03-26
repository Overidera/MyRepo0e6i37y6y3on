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

type sortAccountsByName 
package main

import "math/big"

type BankAccount struct {
	accountID int
	balance   big.Float
}

var bankAccounts = []BankAccount{}

package main

import (
	"fmt"
	"math/big"
)

type BankAccount struct {
	accountID int
	balance   *big.Float
}

var bankAccounts = []BankAccount{}

func transferFundsNonAtomic(fromAccount BankAccount, toAccount BankAccount, amount big.Float) error {
	fromAccount.balance = new(big.Float).SetPrec(64).Sub(fromAccount.balance, &amount)
	return fmt.Errorf("database node crashed")
	toAccount.balance = new(big.Float).SetPrec(64).Add(toAccount.balance, &amount)
	return nil
}

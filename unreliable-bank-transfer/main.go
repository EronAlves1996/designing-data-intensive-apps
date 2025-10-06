package main

import (
	"fmt"
	"math/big"
	"runtime"
	"sync"
	"time"
)

type BankAccount struct {
	accountID int
	balance   *big.Float
}

func transferFundsNonAtomic(fromAccount *BankAccount, toAccount *BankAccount, amount big.Float) error {
	fromAccount.balance = new(big.Float).SetPrec(64).Sub(fromAccount.balance, &amount)
	return fmt.Errorf("database node crashed")
	toAccount.balance = new(big.Float).SetPrec(64).Add(toAccount.balance, &amount)
	return nil
}

func transferFundsDelay(fromAccount *BankAccount, toAccount *BankAccount, amount big.Float) error {
	fromAccount.balance = new(big.Float).SetPrec(64).Sub(fromAccount.balance, &amount)
	time.Sleep(time.Millisecond * 2000)
	toAccount.balance = new(big.Float).SetPrec(64).Add(toAccount.balance, &amount)
	return nil
}

func getTotalBalance(a *BankAccount, b *BankAccount) *big.Float {
	return new(big.Float).SetPrec(64).Add(a.balance, b.balance)
}

func main() {
	aBalance := new(big.Float).SetPrec(64).SetFloat64(100)
	a := BankAccount{
		accountID: 1,
		balance:   aBalance,
	}
	b := BankAccount{
		accountID: 2,
		balance:   new(big.Float).SetPrec(64).SetFloat64(200),
	}

	transferFundsNonAtomic(&a, &b, *new(big.Float).SetFloat64(50))

	// Answer to question 2
	fmt.Printf("Account A balance: %.2f\n", a.balance)
	fmt.Printf("Account B balance: %.2f\n", b.balance)
	// The amount was subtracted from the balance of a, but was not added to
	// the balance of b. It's an atomicity violation problem, because the
	// transfer should have occurred indivisibly. Instead, it was partially
	// suceeded

	// Answer to question 3
	// Here, the total balance is correctly 250
	fmt.Printf("Total Balance: %.2f\n", getTotalBalance(&a, &b))
	var wg sync.WaitGroup
	wg.Go(func() {
		transferFundsDelay(&b, &a, *new(big.Float).SetFloat64(50))
	})
	wg.Go(func() {
		runtime.Gosched()
		// But here, the total balance is 200, as the money get lost in the transaction
		fmt.Printf("Total Balance: %.2f\n", getTotalBalance(&a, &b))
	})
	wg.Wait()
	// During the transfer, the total amount should never change. While the operation is
	// running, the database should return the last value, as the operation never started
	fmt.Printf("Account A balance: %.2f\n", a.balance)
	fmt.Printf("Account B balance: %.2f\n", b.balance)
}

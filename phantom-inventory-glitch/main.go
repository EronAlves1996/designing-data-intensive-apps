package main

import (
	"errors"
	"fmt"
	"math/rand"
	"sync"
	"sync/atomic"
	"time"
)

var inventory = make(map[string]int)
var outOfStockError = errors.New("OUT_OF_STOCK")
var dontExistsError = errors.New("ITEM_DONT_EXISTS")

type InventoryDB interface {
	Query(itemID string) (int, error)
	Update(itemID string, quantity int) error
	GetTransaction()
	Commit()
}

type InstantDB struct {
	lock *sync.Mutex
}

func (i *InstantDB) Query(itemID string) (int, error) {
	qtd, exists := inventory[itemID]
	if !exists {
		return 0, dontExistsError
	}
	return qtd, nil
}

func (i *InstantDB) Update(itemId string, quantity int) error {
	inventory[itemId] = quantity
	return nil
}

func (i *InstantDB) GetTransaction() {
	i.lock.Lock()
}

func (i *InstantDB) Commit() {
	i.lock.Unlock()
}

type NetworkLagDB struct {
	db InventoryDB
}

func (i *NetworkLagDB) Query(itemID string) (int, error) {
	delayMs := rand.Int31n(150) + 50
	<-time.After(time.Duration(delayMs) * time.Millisecond)
	return i.db.Query(itemID)
}

func (i *NetworkLagDB) Update(itemID string, quantity int) error {
	delayMs := rand.Int31n(150) + 50
	<-time.After(time.Duration(delayMs) * time.Millisecond)
	return i.db.Update(itemID, quantity)
}

func (i *NetworkLagDB) GetTransaction() {
	delayMs := rand.Int31n(150) + 50
	<-time.After(time.Duration(delayMs) * time.Millisecond)
	i.db.GetTransaction()
}

func (i *NetworkLagDB) Commit() {
	delayMs := rand.Int31n(150) + 50
	<-time.After(time.Duration(delayMs) * time.Millisecond)
	i.db.Commit()
}

func placeOrder(db InventoryDB, itemId string, quantity int) (bool, error) {
	db.GetTransaction()
	defer db.Commit()
	qtd, err := db.Query(itemId)
	if err != nil {
		return false, err
	}

	if qtd < quantity {
		return false, outOfStockError
	}

	newQtd := qtd - quantity

	db.Update(itemId, newQtd)

	return true, nil
}

func init() {
	inventory["boots"] = 100
}

func main() {
	m := sync.Mutex{}
	db := NetworkLagDB{
		db: &InstantDB{
			lock: &m,
		},
	}

	var wg sync.WaitGroup
	var placed int32 = 0

	for range 20 {
		wg.Go(func() {
			v, err := placeOrder(&db, "boots", 10)
			if err != nil && v {
				atomic.AddInt32(&placed, 10)
			}
		})
	}

	wg.Wait()

	fmt.Printf("item '%s', quantity '%d'. Placed %d\n", "boots", inventory["boots"], placed)
}

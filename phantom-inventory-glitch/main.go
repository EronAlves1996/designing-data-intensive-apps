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
}

type InstantDB struct {
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

func placeOrder(db InventoryDB, itemId string, quantity int) (bool, error) {
	qtd, err := db.Query(itemId)
	if err != nil {
		return false, err
	}

	<-time.After(time.Duration(rand.Int31n(50)) * time.Millisecond)

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
	db := NetworkLagDB{
		db: &InstantDB{},
	}

	var wg sync.WaitGroup
	var placed int32 = 0

	for range 20 {
		wg.Go(func() {
			v, err := placeOrder(&db, "boots", 10)
			if err == nil && v {
				atomic.AddInt32(&placed, 10)
			}
		})
	}

	wg.Wait()

	fmt.Printf("item '%s', quantity '%d'. Placed %d\n", "boots", inventory["boots"], placed)
}

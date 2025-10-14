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
var invalidReservation = errors.New("INVALID_RESERVATION")

type InventoryDB interface {
	Query(itemID string) (int, error)
	TryReserve(itemID string, quantity int) error
	Update(itemID string, quantity int) error
	ReleaseReservation(itemID string, quantity int) error
	ConfirmReservation(itemID string, quantity int) error
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

func (i *InstantDB) TryReserve(itemId string, quantity int) error {
	qtd, err := i.Query(itemId)

	if err != nil {
		return err
	}

	if qtd < quantity {
		return outOfStockError
	}

	inventory[fmt.Sprintf("reserved_%s", itemId)] += quantity
	inventory[itemId] -= quantity

	return nil
}

func (i *InstantDB) ReleaseReservation(itemId string, quantity int) error {
	qtd, exists := inventory[fmt.Sprintf("reserved_%s", itemId)]
	if !exists {
		return dontExistsError
	}

	if qtd < quantity {
		return invalidReservation
	}

	inventory[fmt.Sprintf("reserved_%s", itemId)] -= quantity
	inventory[itemId] += quantity

	return nil
}

func (i *InstantDB) ConfirmReservation(itemID string, quantity int) error {
	qtd, exists := inventory[fmt.Sprintf("reserved_%s", itemID)]
	if !exists {
		return dontExistsError
	}

	if qtd < quantity {
		return invalidReservation
	}

	inventory[fmt.Sprintf("reserved_%s", itemID)] -= quantity

	return nil
}

type NetworkLagDB struct {
	db InventoryDB
}

func (i *NetworkLagDB) Query(itemID string) (int, error) {
	i.delay()
	return i.db.Query(itemID)
}

func (*NetworkLagDB) delay() {
	delayMs := rand.Int31n(150) + 50
	<-time.After(time.Duration(delayMs) * time.Millisecond)
}

func (i *NetworkLagDB) Update(itemID string, quantity int) error {
	i.delay()
	return i.db.Update(itemID, quantity)
}

func (i *NetworkLagDB) TryReserve(itemId string, quantity int) error {
	i.delay()
	return i.db.TryReserve(itemId, quantity)
}

func (i *NetworkLagDB) ReleaseReservation(itemId string, quantity int) error {
	i.delay()
	return i.db.ReleaseReservation(itemId, quantity)
}

func (i *NetworkLagDB) ConfirmReservation(itemID string, quantity int) error {
	i.delay()
	return i.db.ConfirmReservation(itemID, quantity)
}

func placeOrder(db InventoryDB, itemId string, quantity int) (bool, error) {
	qtd, err := db.Query(itemId)
	if err != nil {
		return false, err
	}

	if qtd < quantity {
		return false, outOfStockError
	}

	if err = db.TryReserve(itemId, quantity); err != nil {
		return false, err
	}

	<-time.After(time.Duration(rand.Int31n(50)) * time.Millisecond)

	if err = db.ConfirmReservation(itemId, quantity); err != nil {
		db.ReleaseReservation(itemId, quantity)
		return false, err
	}

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

package main

import (
	"errors"
	"fmt"
	"math/rand"
	"sync"
	"sync/atomic"
	"time"
)

var reservationId int32 = 0
var reservations = make(map[string]int)
var inventory = make(map[string]int)
var outOfStockError = errors.New("OUT_OF_STOCK")
var dontExistsError = errors.New("ITEM_DONT_EXISTS")
var invalidReservation = errors.New("INVALID_RESERVATION")

type InventoryDB interface {
	Query(itemID string) (int, error)
	TryReserve(itemID string, quantity int) (int32, error)
	Update(itemID string, quantity int) error
	ReleaseReservation(itemID string, reservationID int32) error
	ConfirmReservation(itemID string, reservationID int32) error
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

func buildReservationKey(itemID string, reservationID int32) string {
	return fmt.Sprintf("%s_%d", itemID, reservationID)
}

func (i *InstantDB) TryReserve(itemId string, quantity int) (int32, error) {
	qtd, err := i.Query(itemId)

	if err != nil {
		return 0, err
	}

	if qtd < quantity {
		return 0, outOfStockError
	}

	rid := atomic.AddInt32(&reservationId, 1)
	inventory[buildReservationKey(itemId, rid)] += quantity
	inventory[itemId] -= quantity

	return rid, nil
}

func (i *InstantDB) ReleaseReservation(itemId string, rid int32) error {
	rk := buildReservationKey(itemId, rid)
	qtd, exists := reservations[rk]
	if !exists {
		return invalidReservation
	}

	delete(inventory, rk)
	inventory[itemId] += qtd

	return nil
}

func (i *InstantDB) ConfirmReservation(itemID string, rid int32) error {
	rk := buildReservationKey(itemID, rid)
	_, exists := inventory[rk]
	if !exists {
		return invalidReservation
	}

	delete(inventory, rk)

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

func (i *NetworkLagDB) TryReserve(itemId string, quantity int) (int32, error) {
	i.delay()
	return i.db.TryReserve(itemId, quantity)
}

func (i *NetworkLagDB) ReleaseReservation(itemId string, rid int32) error {
	i.delay()
	return i.db.ReleaseReservation(itemId, rid)
}

func (i *NetworkLagDB) ConfirmReservation(itemID string, rid int32) error {
	i.delay()
	return i.db.ConfirmReservation(itemID, rid)
}

func placeOrder(db InventoryDB, itemId string, quantity int) (bool, error) {
	qtd, err := db.Query(itemId)
	if err != nil {
		return false, err
	}

	if qtd < quantity {
		return false, outOfStockError
	}

	rid, err := db.TryReserve(itemId, quantity)
	if err != nil {
		return false, err
	}
	defer db.ReleaseReservation(itemId, rid)

	<-time.After(time.Duration(rand.Int31n(50)) * time.Millisecond)

	if err = db.ConfirmReservation(itemId, rid); err != nil {
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

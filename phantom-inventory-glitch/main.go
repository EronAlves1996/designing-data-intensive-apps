package main

import "errors"

var inventory = make(map[string]int)
var outOfStockError = errors.New("OUT_OF_STOCK")
var dontExistsError = errors.New("ITEM_DONT_EXISTS")

func placeOrder(itemId string, quantity int) (bool, error) {
	qtd, exists := inventory[itemId]

	if !exists {
		return false, dontExistsError
	}

	if qtd < quantity {
		return false, outOfStockError
	}

	inventory[itemId] -= quantity
	return true, nil
}

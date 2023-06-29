package main

import (
	"errors"
)

type Product struct {
	Code     int
	Name     string
	Price    float64
	Quantity int
}

// Função para atualizar o estoque de um produto com base nas vendas realizadas
func updateStock(product *Product, sales map[int]int) error {
	if product == nil {
		return errors.New("product object is nil")
	}

	for code, quantity := range sales {
		if quantity <= 0 {
			return errors.New("invalid quantity sold")
		}

		if code == product.Code {
			newQuantity := product.Quantity - quantity
			if newQuantity < 0 {
				return errors.New("insufficient stock")
			}
			product.Quantity = newQuantity
		}
	}

	return nil
}

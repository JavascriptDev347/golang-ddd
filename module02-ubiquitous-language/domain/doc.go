// Package domain -- Module 02 (Ubiquitous Language va Bounded Context) uchun tayyor misollar shu yerga qo'shiladi.
// Hozircha bo'sh -- placeholder, repo compile bo'lishi uchun.
package domain

import "fmt"

// Bad code - Kod tili va biznes tili uzilgan xolatda
type Order struct {
	Status  int // 0,1,2,3,4 - bu nimani anglatishini aniq bilmaydi
	Status2 OrderStatus
}

func (o *Order) SetFlag(f int) {
	o.Status = f
}

// Endi yaxshilanganda agar Ubiquitous Language
type OrderStatus string

const (
	OrderStatusPlaced    OrderStatus = "placed"
	OrderStatusConfirmed OrderStatus = "confirmed"
	OrderStatusShipped   OrderStatus = "shipped"
	OrderStatusCancelled OrderStatus = "cancelled"
)

func (o *Order) Confirm() error {
	if o.Status2 != OrderStatusPlaced {
		return fmt.Errorf("order is not placed")
	}
	o.Status2 = OrderStatusConfirmed
	return nil
}

// Bounded Context

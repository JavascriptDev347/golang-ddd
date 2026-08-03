package domain

import (
	"errors"
	"time"
)

type OrderID string

type OrderStatus string

const (
	StatusPending   OrderStatus = "pending"
	StatusPaid      OrderStatus = "paid"
	StatusShipped   OrderStatus = "shipped"
	StatusCancelled OrderStatus = "cancelled"
)

type Order struct {
	id         OrderID
	customerID OrderID
	total      Money
	status     OrderStatus
	createdAt  time.Time
}

func NewOrder(id OrderID, total Money) *Order {
	return &Order{
		id:        id,
		total:     total,
		status:    StatusPending,
		createdAt: time.Now(),
	}
}

// MarkAsPaid — bu Entity'ning "lifecycle" qismi: holat vaqt bilan o'zgaradi.
// Value Object'da bunday metod BO'LMAYDI — chunki agar o'zgarsa, bu YANGI
// Value Object bo'lishi kerak, eskisi o'zgarmasligi kerak.
func (o *Order) MarkAsPaid() error {
	if o.status != StatusPending {
		return errors.New("order faqat pending holatidan paid holatiga o'tishi mumkin")
	}
	o.status = StatusPaid
	return nil
}

// Equals — Entity uchun tenglik FAQAT ID orqali tekshiriladi.
// Maydonlar boshqacha bo'lsa ham (masalan, biri eski snapshot, ikkinchisi
// yangilangan), agar ID bir xil bo'lsa — bu BIR XIL order.
func (o *Order) Equals(other *Order) bool {
	if other == nil {
		return false
	}
	return o.id == other.id
}

func (o *Order) ID() OrderID         { return o.id }
func (o *Order) Status() OrderStatus { return o.status }

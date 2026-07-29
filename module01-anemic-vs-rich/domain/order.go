// Package domain — Module 1: Rich Domain Model to'liq ishlaydigan misol.
//
// Bu fayl "domain layer" deb ataladi — bu yerda faqat biznes qoidalar
// (invariant'lar) yashaydi. Hech qanday HTTP, DB, JSON haqida bilim yo'q.
// Bu qasddan shunday — domain layer TASHQI DUNYODAN mustaqil bo'lishi kerak
// (Module 8'da "Hexagonal Architecture"da buning nega muhimligini chuqur ko'ramiz).
package domain

import (
	"errors"
	"fmt"
)

// ---------- Value Object: Money ----------
// Money hali "to'liq" Value Object emas (Module 3'da Equals() kabi metodlar
// bilan boyitamiz), lekin asosiy g'oya shu: pul — bu oddiy son emas, balki
// o'zining qoidalariga ega tur.
type Money struct {
	amount   int64 // tiyin/cent'larda — float64 ASLO ishlatilmaydi (rounding xatolari)
	currency string
}

func NewMoney(amount int64, currency string) (Money, error) {
	if amount < 0 {
		return Money{}, errors.New("money: amount cannot be negative")
	}
	if currency == "" {
		return Money{}, errors.New("money: currency is required")
	}
	return Money{amount: amount, currency: currency}, nil
}

func (m Money) Amount() int64     { return m.amount }
func (m Money) Currency() string  { return m.currency }
func (m Money) Add(other Money) (Money, error) {
	if m.currency != other.currency {
		return Money{}, fmt.Errorf("money: currency mismatch %s vs %s", m.currency, other.currency)
	}
	return Money{amount: m.amount + other.amount, currency: m.currency}, nil
}

// ---------- Entity: OrderItem ----------
type OrderItem struct {
	ProductID string
	Quantity  int
	UnitPrice Money
}

func (i OrderItem) subtotal() (Money, error) {
	total, err := NewMoney(i.UnitPrice.amount*int64(i.Quantity), i.UnitPrice.currency)
	if err != nil {
		return Money{}, err
	}
	return total, nil
}

// ---------- OrderStatus ----------
type OrderStatus string

const (
	StatusPending   OrderStatus = "pending"
	StatusPaid      OrderStatus = "paid"
	StatusShipped   OrderStatus = "shipped"
	StatusCancelled OrderStatus = "cancelled"
)

var (
	ErrEmptyOrder        = errors.New("order: must contain at least one item")
	ErrInvalidTransition = errors.New("order: invalid status transition")
	ErrInsufficientPay   = errors.New("order: payment amount is insufficient")
	ErrCurrencyMismatch  = errors.New("order: currency mismatch")
)

// ---------- Aggregate Root: Order ----------
// Diqqat: BARCHA maydonlar KICHIK harf bilan. Bu qasddan qilingan.
// Tashqi kod `order.status = "paid"` deb yoza OLMAYDI — compile xatosi beradi.
// Status faqat Pay() / Ship() / Cancel() orqaligina o'zgaradi.
type Order struct {
	id     string
	status OrderStatus
	items  []OrderItem
}

// NewOrder — yagona "eshik" Order yaratish uchun. Invariant shu yerda tekshiriladi.
func NewOrder(id string, items []OrderItem) (*Order, error) {
	if len(items) == 0 {
		return nil, ErrEmptyOrder
	}
	for _, it := range items {
		if it.Quantity <= 0 {
			return nil, fmt.Errorf("order: item %s has invalid quantity %d", it.ProductID, it.Quantity)
		}
	}
	return &Order{
		id:     id,
		status: StatusPending,
		items:  items,
	}, nil
}

// Total — barcha item'larning umumiy summasini hisoblaydi.
func (o *Order) Total() (Money, error) {
	if len(o.items) == 0 {
		return Money{}, ErrEmptyOrder
	}
	total, err := o.items[0].subtotal()
	if err != nil {
		return Money{}, err
	}
	for _, it := range o.items[1:] {
		sub, err := it.subtotal()
		if err != nil {
			return Money{}, err
		}
		total, err = total.Add(sub)
		if err != nil {
			return Money{}, err
		}
	}
	return total, nil
}

// Pay — biznes qoida shu yerda "yashaydi". Bu metod bo'lmasa,
// har bir developer bu qoidani boshqacha yozishi (yoki unutishi) mumkin edi.
func (o *Order) Pay(amount Money) error {
	if o.status != StatusPending {
		return fmt.Errorf("%w: cannot pay order in status %q", ErrInvalidTransition, o.status)
	}
	total, err := o.Total()
	if err != nil {
		return err
	}
	if amount.currency != total.currency {
		return ErrCurrencyMismatch
	}
	if amount.amount < total.amount {
		return fmt.Errorf("%w: got %d, need %d", ErrInsufficientPay, amount.amount, total.amount)
	}
	o.status = StatusPaid
	return nil
}

// Ship — faqat "paid" holatidan o'tish mumkin.
func (o *Order) Ship() error {
	if o.status != StatusPaid {
		return fmt.Errorf("%w: cannot ship order in status %q", ErrInvalidTransition, o.status)
	}
	o.status = StatusShipped
	return nil
}

// Cancel — "shipped" bo'lgan orderni bekor qilib bo'lmaydi (real-life qoida).
func (o *Order) Cancel() error {
	if o.status == StatusShipped {
		return fmt.Errorf("%w: cannot cancel a shipped order", ErrInvalidTransition)
	}
	if o.status == StatusCancelled {
		return fmt.Errorf("%w: order already cancelled", ErrInvalidTransition)
	}
	o.status = StatusCancelled
	return nil
}

// ---------- Faqat o'qish uchun getterlar (yozish uchun setter YO'Q) ----------
func (o *Order) ID() string         { return o.id }
func (o *Order) Status() OrderStatus { return o.status }
func (o *Order) Items() []OrderItem {
	// Diqqat: slice'ning nusxasini qaytaramiz, aslini emas —
	// aks holda chaqiruvchi `order.Items()[0].Quantity = 999` deb yozib,
	// aggregate ichini tashqaridan buzishi mumkin edi.
	cp := make([]OrderItem, len(o.items))
	copy(cp, o.items)
	return cp
}

package domain

import (
	"errors"
	"fmt"
)

type ProductID string
type OrderID string
type CustomerID string
type OrderStatus string

const (
	StatusPending   OrderStatus = "pending"
	StatusPaid      OrderStatus = "paid"
	StatusShipped   OrderStatus = "shipped"
	StatusCancelled OrderStatus = "cancelled"
)

// OrderLine — bu aggregate ICHIDAGI Entity. Diqqat: uning o'z ID'si bor
// (chunki bitta buyurtmada bir xil mahsulot ikki marta bo'lishi mumkin,
// va ularni farqlash kerak bo'lishi mumkin — masalan chegirma qo'llashda),
// lekin u HECH QACHON o'z-o'zicha, Order'dan tashqarida yashamaydi.
// Shuning uchun uning konstruktori ham, mutatsiya metodlari ham — unexported
// yoki faqat Order orqali chaqiriladi

type OrderLine struct {
	id        string
	productID ProductID
	quantity  int
	unitPrice Money
}

func (l OrderLine) Subtotal() (Money, error) {
	total := int64(0)
	for i := 0; i < l.quantity; i++ {
		total += l.unitPrice.amount
	}
	m, err := NewMoney(total, l.unitPrice.currency)
	if err != nil {
		return Money{}, err
	}
	return m, nil
}

// Order — AGGREGATE ROOT. U endi shunchaki Entity emas — u butun
// klasterning (o'zi + OrderLine'lar) yagona kirish nuqtasi.
type Order struct {
	id       OrderID
	customer CustomerID
	lines    []OrderLine // eksport qilinmagan — tashqaridan to'g'ridan-to'g'ri o'zgartirib bo'lmaydi
	status   OrderStatus
	maxLines int
}

func NewOrder(id OrderID, customer CustomerID) *Order {
	return &Order{
		id:       id,
		customer: customer,
		status:   StatusPending,
		maxLines: 10, // invariant: bitta buyurtmada 10 tadan ortiq turdagi mahsulot bo'lmasin
	}
}

func (o *Order) AddLine(productID ProductID, quantity int, price Money) error {
	if o.status != StatusPending {
		return errors.New("faqat pending buyurtmaga qator qo'shish mumkin")
	}
	if quantity <= 0 {
		return errors.New("miqdor musbat bo'lishi kerak")
	}
	if len(o.lines) >= o.maxLines {
		return fmt.Errorf("bitta buyurtmada maksimal %d ta qator bo'lishi mumkin", o.maxLines)
	}
	line := OrderLine{
		id:        fmt.Sprintf("%s-line-%d", o.id, len(o.lines)+1),
		productID: productID,
		quantity:  quantity,
		unitPrice: price,
	}
	o.lines = append(o.lines, line)
	return nil
}

func (o *Order) Total() (Money, error) {
	total, err := NewMoney(0, "USD")
	if err != nil {
		return Money{}, err
	}
	for _, l := range o.lines {
		sub, err := l.Subtotal()
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

func (o *Order) LineCount() int { return len(o.lines) }

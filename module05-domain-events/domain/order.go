package domain

import (
	"time"
)

type OrderID string
type ProductID string
type CustomerID string

type DomainEvent interface {
	EventName() string
	OccurredAt() time.Time
}

// OrderPaid — konkret event. Diqqat: bu struct IMMUTABLE (Value Object
// kabi) — chunki bu "o'tmishda sodir bo'lgan fakt", uni o'zgartirib
// bo'lmaydi, faqat o'qish mumkin.
type OrderPaid struct {
	orderId    string
	amount     Money
	occurredAt time.Time
}

func (e OrderPaid) EventName() string {
	return "order.paid"
}

func (e OrderPaid) OccurredAt() time.Time {
	return e.occurredAt
}

func (e OrderPaid) OrderId() string {
	return e.orderId
}

func (e OrderPaid) Amount() Money {
	return e.amount
}

type OrderLineAdded struct {
	orderID    OrderID
	productID  ProductID
	quantity   int
	occurredAt time.Time
}

func (e OrderLineAdded) EventName() string     { return "order.line_added" }
func (e OrderLineAdded) OccurredAt() time.Time { return e.occurredAt }

// type Order struct {
// 	id       OrderID
// 	customer CustomerID
// 	lines    []OrderLine
// 	status   OrderStatus
// 	maxLines int

// 	// pendingEvents — bu aggregate hali "e'lon qilinmagan" hodisalarni
// 	// vaqtincha saqlaydi. Bu ham unexported — tashqi kod to'g'ridan-to'g'ri
// 	// bu slice'ga event qo'sha olmaydi, faqat Order o'zi, o'z metodlari
// 	// ichida, biznes qoida bajarilganda qo'shadi.
// 	pendingEvents []DomainEvent
// }

// func (o *Order) AddLine(productID ProductID, quantity int, price Money) error {
// 	if o.status != StatusPending {
// 		return errors.New("faqat pending buyurtmaga qator qo'shish mumkin")
// 	}
// 	if quantity <= 0 {
// 		return errors.New("miqdor musbat bo'lishi kerak")
// 	}
// 	if len(o.lines) >= o.maxLines {
// 		return fmt.Errorf("bitta buyurtmada maksimal %d ta qator bo'lishi mumkin", o.maxLines)
// 	}
// 	line := OrderLine{
// 		id:        fmt.Sprintf("%s-line-%d", o.id, len(o.lines)+1),
// 		productID: productID,
// 		quantity:  quantity,
// 		unitPrice: price,
// 	}
// 	o.lines = append(o.lines, line)

// 	// Invariant bajarildi — endi FAKT sifatida event yig'amiz.
// 	o.pendingEvents = append(o.pendingEvents, OrderLineAdded{
// 		orderID:    o.id,
// 		productID:  productID,
// 		quantity:   quantity,
// 		occurredAt: time.Now(),
// 	})
// 	return nil
// }

// func (o *Order) MarkAsPaid(amount Money) error {
// 	if o.status != StatusPending {
// 		return errors.New("order faqat pending holatidan paid holatiga o'tishi mumkin")
// 	}
// 	o.status = StatusPaid
// 	o.pendingEvents = append(o.pendingEvents, OrderPaid{
// 		orderID:    o.id,
// 		amount:     amount,
// 		occurredAt: time.Now(),
// 	})
// 	return nil
// }

// // PullEvents — yig'ilgan eventlarni chiqarib oladi VA ichki listni tozalaydi.
// // "Pull", "push" emas — chunki Order hech kimga o'zi yubormaydi, kimdir
// // (application layer) kelib so'raydi.
// func (o *Order) PullEvents() []DomainEvent {
// 	events := o.pendingEvents
// 	o.pendingEvents = nil
// 	return events
// }

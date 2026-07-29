package domain

import (
	"errors"
	"testing"
)

// Bu testlar — SENING kodingni tekshirish uchun emas, balki
// "domain" papkasidagi TAYYOR yechimning qanday ishlashini ko'rsatish uchun.
// `go test ./module01-anemic-vs-rich/domain/...` -v bilan ishga tushirib,
// har bir test nimani tekshirayotganini o'qi.

func mustMoney(t *testing.T, amount int64, currency string) Money {
	t.Helper()
	m, err := NewMoney(amount, currency)
	if err != nil {
		t.Fatalf("unexpected error creating money: %v", err)
	}
	return m
}

func TestNewOrder_EmptyItems_ReturnsError(t *testing.T) {
	_, err := NewOrder("order-1", []OrderItem{})
	if !errors.Is(err, ErrEmptyOrder) {
		t.Fatalf("expected ErrEmptyOrder, got %v", err)
	}
}

func TestOrder_Total_CalculatesCorrectly(t *testing.T) {
	items := []OrderItem{
		{ProductID: "p1", Quantity: 2, UnitPrice: mustMoney(t, 1000, "UZS")}, // 2000
		{ProductID: "p2", Quantity: 1, UnitPrice: mustMoney(t, 500, "UZS")},  // 500
	}
	order, err := NewOrder("order-1", items)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	total, err := order.Total()
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if total.Amount() != 2500 {
		t.Fatalf("expected total 2500, got %d", total.Amount())
	}
}

func TestOrder_Pay_InsufficientAmount_ReturnsError(t *testing.T) {
	items := []OrderItem{{ProductID: "p1", Quantity: 1, UnitPrice: mustMoney(t, 1000, "UZS")}}
	order, _ := NewOrder("order-1", items)

	err := order.Pay(mustMoney(t, 500, "UZS"))
	if !errors.Is(err, ErrInsufficientPay) {
		t.Fatalf("expected ErrInsufficientPay, got %v", err)
	}
	// MUHIM: xato bo'lganda status o'zgarmasligi kerak — bu invariant.
	if order.Status() != StatusPending {
		t.Fatalf("status should remain pending after failed payment, got %s", order.Status())
	}
}

func TestOrder_Pay_Success_ChangesStatus(t *testing.T) {
	items := []OrderItem{{ProductID: "p1", Quantity: 1, UnitPrice: mustMoney(t, 1000, "UZS")}}
	order, _ := NewOrder("order-1", items)

	if err := order.Pay(mustMoney(t, 1000, "UZS")); err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if order.Status() != StatusPaid {
		t.Fatalf("expected status paid, got %s", order.Status())
	}
}

func TestOrder_Pay_Twice_ReturnsError(t *testing.T) {
	items := []OrderItem{{ProductID: "p1", Quantity: 1, UnitPrice: mustMoney(t, 1000, "UZS")}}
	order, _ := NewOrder("order-1", items)
	_ = order.Pay(mustMoney(t, 1000, "UZS"))

	err := order.Pay(mustMoney(t, 1000, "UZS"))
	if !errors.Is(err, ErrInvalidTransition) {
		t.Fatalf("expected ErrInvalidTransition, got %v", err)
	}
}

func TestOrder_Cancel_AfterShipped_ReturnsError(t *testing.T) {
	items := []OrderItem{{ProductID: "p1", Quantity: 1, UnitPrice: mustMoney(t, 1000, "UZS")}}
	order, _ := NewOrder("order-1", items)
	_ = order.Pay(mustMoney(t, 1000, "UZS"))
	_ = order.Ship()

	err := order.Cancel()
	if !errors.Is(err, ErrInvalidTransition) {
		t.Fatalf("expected ErrInvalidTransition when cancelling shipped order, got %v", err)
	}
}

func TestOrder_Items_ReturnsCopyNotOriginal(t *testing.T) {
	items := []OrderItem{{ProductID: "p1", Quantity: 1, UnitPrice: mustMoney(t, 1000, "UZS")}}
	order, _ := NewOrder("order-1", items)

	got := order.Items()
	got[0].Quantity = 999 // bu faqat nusxani o'zgartirishi kerak

	original := order.Items()
	if original[0].Quantity != 1 {
		t.Fatalf("aggregate internal state was mutated from outside! quantity = %d", original[0].Quantity)
	}
}

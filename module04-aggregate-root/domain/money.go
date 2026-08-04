package domain

import (
	"errors"
	"fmt"
)

type Money struct {
	amount   int64
	currency string
}

func NewMoney(amount int64, currency string) (Money, error) {
	if amount < 0 {
		return Money{}, errors.New("summa manfiy bo'lishi mumkin emas")
	}
	if currency == "" {
		return Money{}, errors.New("valyuta bo'sh bo'lishi mumkin emas")
	}
	return Money{
		amount:   amount,
		currency: currency,
	}, nil
}

func (m Money) Add(other Money) (Money, error) {
	if m.currency != other.currency {
		return Money{}, fmt.Errorf("valyutalar mos kelmaydi: %s vs %s", m.currency, other.currency)
	}

	return Money{amount: m.amount + other.amount, currency: m.currency}, nil
}

func (m Money) Equals(other Money) bool {
	return m.amount == other.amount && m.currency == other.currency
}

func (m Money) String() string {
	return fmt.Sprintf("%.2f %s", float64(m.amount)/100, m.currency)
}

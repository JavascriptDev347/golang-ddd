package domain

import "errors"

type Money struct {
	amount   int64 // tiyin/cent'larda saqlanadi — float bilan pul hisoblash antipattern
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

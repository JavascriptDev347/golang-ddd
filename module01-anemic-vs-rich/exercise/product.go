// Package exercise — SENING TOPSHIRIG'ING.
//
// Vazifa: quyidagi ANEMIC Product structini RICH DOMAIN MODEL'ga aylantir.
// Har bir TODO belgisini top va kodni yoz. Tayyor bo'lgach:
//
//	go test ./module01-anemic-vs-rich/exercise/...
//
// buyrug'ini ishga tushir. Barcha testlar yashil (PASS) bo'lguncha ishla.
// Yordam kerak bo'lsa domain/order.go faylidagi tayyor yechimga qara —
// lekin avval o'zing urinib ko'r, shunda haqiqiy o'rganish sodir bo'ladi.
package exercise

import (
	"errors"
)

// TODO 1: Quyidagi maydonlarni PRIVATE (kichik harf) qil.
// Hozircha ular public — bu Anemic Model belgisi.
type Product struct {
	id       string
	name     string
	price    int64 // tiyinlarda
	stock    int
	isactive bool
}

// Quyidagi xatolarni ishlat (allaqachon tayyor, o'zgartirmasdan foydalan):
var (
	ErrInvalidName       = errors.New("product: name is required")
	ErrInvalidPrice      = errors.New("product: price must be positive")
	ErrInvalidQuantity   = errors.New("product: quantity must be positive")
	ErrInsufficientStock = errors.New("product: insufficient stock")
	ErrProductInactive   = errors.New("product: cannot sell inactive product")
)

// TODO 2: NewProduct constructor'ni yoz.
// Talablar:
//   - name bo'sh bo'lmasligi kerak (ErrInvalidName)
//   - price > 0 bo'lishi kerak (ErrInvalidPrice)
//   - yangi mahsulot har doim IsActive = true holatida yaratiladi
func NewProduct(id, name string, price int64, stock int) (*Product, error) {
	if name == "" {
		return nil, ErrInvalidName
	}
	if price < 0 {
		return nil, ErrInvalidPrice
	}
	if stock < 0 {
		return nil, ErrInvalidQuantity
	}

	return &Product{
		id:       id,
		name:     name,
		price:    price,
		stock:    stock,
		isactive: true,
	}, nil
}

// TODO 3: ReduceStock metodini yoz.
// Qoidalar:
//   - qty <= 0 bo'lsa -> ErrInvalidQuantity
//   - !IsActive bo'lsa -> ErrProductInactive
//   - stock - qty < 0 bo'lsa -> ErrInsufficientStock
//   - aks holda: stock kamayadi, nil qaytadi
func (p *Product) ReduceStock(qty int) error {
	if qty <= 0 {
		return ErrInvalidQuantity
	}
	if !p.isactive {
		return ErrProductInactive
	}

	if p.stock-qty < 0 {
		return ErrInsufficientStock
	}

	p.stock -= qty
	return nil
}

// TODO 4: Deactivate metodini yoz (parametrsiz, IsActive'ni false qiladi)
func (p *Product) Deactivate() {
	p.isactive = false
}

// TODO 5: Kerakli getter metodlarini yoz: Stock() int, Price() int64, IsActive() bool, ID() string, Name() string
func (p *Product) Stock() int {
	return p.stock
}

func (p *Product) Price() int64 {
	return p.price
}

func (p *Product) IsActive() bool {
	return p.isactive
}

func (p *Product) ID() string {
	return p.id
}

func (p *Product) Name() string {
	return p.name
}

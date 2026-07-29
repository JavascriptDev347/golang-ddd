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
	ID       string
	Name     string
	Price    int64 // tiyinlarda
	Stock    int
	IsActive bool
}

// Quyidagi xatolarni ishlat (allaqachon tayyor, o'zgartirmasdan foydalan):
var (
	ErrInvalidName     = errors.New("product: name is required")
	ErrInvalidPrice    = errors.New("product: price must be positive")
	ErrInvalidQuantity = errors.New("product: quantity must be positive")
	ErrInsufficientStock = errors.New("product: insufficient stock")
	ErrProductInactive = errors.New("product: cannot sell inactive product")
)

// TODO 2: NewProduct constructor'ni yoz.
// Talablar:
//   - name bo'sh bo'lmasligi kerak (ErrInvalidName)
//   - price > 0 bo'lishi kerak (ErrInvalidPrice)
//   - yangi mahsulot har doim IsActive = true holatida yaratiladi
//
// func NewProduct(id, name string, price int64, stock int) (*Product, error) {
//     // TODO: shu yerga yoz
// }

// TODO 3: ReduceStock metodini yoz.
// Qoidalar:
//   - qty <= 0 bo'lsa -> ErrInvalidQuantity
//   - !IsActive bo'lsa -> ErrProductInactive
//   - stock - qty < 0 bo'lsa -> ErrInsufficientStock
//   - aks holda: stock kamayadi, nil qaytadi
//
// func (p *Product) ReduceStock(qty int) error {
//     // TODO: shu yerga yoz
// }

// TODO 4: Deactivate metodini yoz (parametrsiz, IsActive'ni false qiladi)
//
// func (p *Product) Deactivate() error {
//     // TODO: shu yerga yoz
// }

// TODO 5: Kerakli getter metodlarini yoz: Stock() int, Price() int64, IsActive() bool, ID() string, Name() string

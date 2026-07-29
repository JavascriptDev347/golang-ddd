# 📘 MODULE 1: DDD nima va nega kerak? Anemic vs Rich Domain Model

## 1️⃣ Nima o'rganamiz va nega bu shunday ishlaydi

### Muammoning ildizi (root cause)

Sen NestJS/TypeScript fonida ishlagansan. Odatiy NestJS loyihada shunday struktura ko'rasan:

```typescript
// TypeScript - "Anemic Domain Model" (odatiy NestJS uslubi)
export class Order {
  id: string;
  status: string; // "pending" | "paid" | "shipped" | "cancelled"
  totalAmount: number;
  items: OrderItem[];
}

// Butun logika Service qatlamida:
@Injectable()
export class OrderService {
  async payOrder(orderId: string, amount: number) {
    const order = await this.orderRepo.findById(orderId);
    
    // Biznes qoidalar Service ichida yozilgan
    if (order.status !== 'pending') {
      throw new Error('Cannot pay non-pending order');
    }
    if (amount < order.totalAmount) {
      throw new Error('Insufficient amount');
    }
    
    order.status = 'paid';
    await this.orderRepo.save(order);
  }
}
```

Bu yerda **Order** — bu shunchaki ma'lumot qutisi (data bag). Uning o'zida hech qanday himoya, hech qanday "aqli" yo'q. Har qanday joydan kimdir shunday yozishi mumkin:

```typescript
order.status = 'paid'; // hech qanday tekshiruvsiz, to'g'ridan-to'g'ri
```

**Bu — Anemic Domain Model (Kamqonli Domen Modeli) deyiladi.** Martin Fowler buni "anti-pattern" deb ataydi, sababi:

- Biznes qoidalar butun kodbazaga tarqalib ketadi (Service'larda, Controller'larda, hatto frontend'da)
- Order obyektining o'zi hech qanday "invariant" (har doim to'g'ri bo'lishi kerak bo'lgan qoida)ni himoya qilmaydi
- 50ta joyda `order.status = 'paid'` yozilishi mumkin — va bittasi ham tekshiruvsiz qolsa, production'da bug

### DDD yechimi — Rich Domain Model

Go'da DDD yondashuvida xuddi shu narsa boshqacha ko'rinishga ega:

```go
package domain

import (
	"errors"
	"fmt"
)

type OrderStatus string

const (
	StatusPending   OrderStatus = "pending"
	StatusPaid      OrderStatus = "paid"
	StatusShipped   OrderStatus = "shipped"
	StatusCancelled OrderStatus = "cancelled"
)

// Order — bu endi shunchaki struct emas, balki "aql"ga ega domen obyekti.
// E'tibor ber: barcha maydonlar KICHIK harf bilan (private) — 
// tashqaridan to'g'ridan-to'g'ri o'zgartirib bo'lmaydi.
type Order struct {
	id          string
	status      OrderStatus
	totalAmount Money // Value Object — Module 3'da chuqur o'rganamiz
	items       []OrderItem
}

// Money — oddiy float64 emas, chunki pul bilan ishlashda 
// rounding xatolari va valyuta aralashmasligi kerak
type Money struct {
	amount   int64 // tiyin/cent'larda saqlanadi, float xatolarining oldini olish uchun
	currency string
}

// NewOrder — bu "constructor" funksiya. Faqat SHU orqali Order yaratiladi.
// Bu yerda invariant'larni (qoidalarni) darhol tekshiramiz.
func NewOrder(id string, items []OrderItem) (*Order, error) {
	if len(items) == 0 {
		return nil, errors.New("order must have at least one item")
	}

	total := calculateTotal(items)

	return &Order{
		id:          id,
		status:      StatusPending,
		totalAmount: total,
		items:       items,
	}, nil
}

// Pay — bu METOD, lekin muhimi: bu yerda BIZNES QOIDA yashaydi.
// Order o'zini o'zi himoya qiladi — hech kim uni "noto'g'ri holat"ga o'tkaza olmaydi.
func (o *Order) Pay(amount Money) error {
	if o.status != StatusPending {
		return fmt.Errorf("cannot pay order in status %s: only pending orders can be paid", o.status)
	}
	if amount.amount < o.totalAmount.amount {
		return fmt.Errorf("insufficient payment: got %d, need %d", amount.amount, o.totalAmount.amount)
	}
	if amount.currency != o.totalAmount.currency {
		return fmt.Errorf("currency mismatch: got %s, need %s", amount.currency, o.totalAmount.currency)
	}

	o.status = StatusPaid
	return nil
}

// Status() — o'qish uchun public getter. Lekin YOZISH uchun getter YO'Q.
// Bu qasddan qilingan — status faqat Pay(), Ship(), Cancel() metodlari orqaligina o'zgaradi.
func (o *Order) Status() OrderStatus {
	return o.status
}

func (o *Order) ID() string {
	return o.id
}
```

### Nega bu ishlaydi — Go'ning o'ziga xos jihati

TypeScript'da `private` keyword bor, lekin u faqat **compile-time**da tekshiriladi va runtime'da (masalan JSON orqali) baribir chetlab o'tish mumkin. Go'da esa:

1. **Kichik harf bilan boshlangan maydon** (`status`, `totalAmount`) — bu **package-level encapsulation**. Bu `orderTestExamplePackage` tashqarisidan **hech qanday tarzda** — na reflection, na JSON marshalling orqali to'g'ridan-to'g'ri o'zgartirib bo'lmaydi (JSON marshalling uchun alohida `MarshalJSON` metodi kerak bo'ladi, bu ham nazorat ostida).
2. Bu **package chegarasi orqali himoya** — sen buni "esda tutgan kelishuv" (convention) emas, **til darajasidagi qat'iy qoida** sifatida olasan.

**Bu — sening avvalgi tushunchang bilan taqqoslash:** Xotirangda `lowercase = TS class-private kabi` degan tasavvur bor edi — bu **to'g'ri emas edi ammo hozir to'g'ri kontekstda**: TS'da private "shu klass ichida" degani, Go'da esa "shu **package** ichida" degani. Farq katta: agar `Order` va uni ishlatuvchi kod bir xil package'da bo'lsa, private maydonlar baribir ko'rinadi. Shuning uchun DDD'da har bir **Aggregate uchun alohida package** yaratish odatiy amaliyot (buni Module 4'da chuqur ko'ramiz).

---

## 2️⃣ Real misol va sen bajaradigan task

### ✅ Real ishlatilish misoli: Bank hisobi (Bank Account)

Bu klassik DDD misoli, chunki bank domeni tabiiy ravishda ko'p invariant'larga ega:

```go
package domain

import (
	"errors"
	"fmt"
)

var (
	ErrInsufficientFunds = errors.New("insufficient funds")
	ErrNegativeAmount    = errors.New("amount must be positive")
	ErrAccountFrozen     = errors.New("account is frozen")
)

type AccountStatus string

const (
	AccountActive AccountStatus = "active"
	AccountFrozen AccountStatus = "frozen"
	AccountClosed AccountStatus = "closed"
)

type BankAccount struct {
	id      string
	balance int64 // tiyinlarda (cents), float64 ASLO ishlatilmaydi pul uchun
	status  AccountStatus
	owner   string
}

func NewBankAccount(id, owner string) *BankAccount {
	return &BankAccount{
		id:      id,
		balance: 0,
		status:  AccountActive,
		owner:   owner,
	}
}

// Withdraw — pulni yechish. Bu yerda 3 xil invariant tekshiriladi.
// Bularning HAMMASI Anemic modelda Service qatlamida yozilgan bo'lardi,
// va har safar yangi Service metodi yozilganda birortasi UNUTILIB QOLISHI mumkin edi.
func (a *BankAccount) Withdraw(amount int64) error {
	if amount <= 0 {
		return ErrNegativeAmount
	}
	if a.status != AccountActive {
		return fmt.Errorf("%w: account status is %s", ErrAccountFrozen, a.status)
	}
	if a.balance < amount {
		return fmt.Errorf("%w: balance %d, requested %d", ErrInsufficientFunds, a.balance, amount)
	}

	a.balance -= amount
	return nil
}

func (a *BankAccount) Deposit(amount int64) error {
	if amount <= 0 {
		return ErrNegativeAmount
	}
	if a.status != AccountActive {
		return ErrAccountFrozen
	}

	a.balance += amount
	return nil
}

func (a *BankAccount) Balance() int64 {
	return a.balance
}

func (a *BankAccount) Freeze() {
	a.status = AccountFrozen
}
```

**Diqqat qil:** endi `BankAccount` obyektini ishlatuvchi HAR QANDAY kod — HTTP handler, gRPC handler, CLI, background job — bir xil qoidalarga bo'ysunadi, chunki qoidalar **bitta joyda** — domen obyektining o'zida. Bu DDD'ning eng katta amaliy foydasi: **"single source of truth" biznes logikasi uchun.**

### 📝 Sening topshiriging (o'zing bajarasan)

Quyidagi **Anemic** kodni ol va uni **Rich Domain Model**ga aylantir:

```go
// BU KOD ANEMIC — sen buni to'g'irlashing kerak
type Product struct {
	ID       string
	Name     string
	Price    int64 // tiyinlarda
	Stock    int
	IsActive bool
}

// Bu funksiyalar hozir alohida "service" qatlamida:
func ReduceStock(p *Product, qty int) {
	p.Stock = p.Stock - qty
}

func Deactivate(p *Product) {
	p.IsActive = false
}
```

**Talablar:**
1. `Product` structining barcha maydonlarini **private** qil (kichik harf)
2. `NewProduct(...)` constructor yoz — kamida shu qoidani tekshir: `price > 0` va `name` bo'sh bo'lmasligi kerak
3. `ReduceStock(qty int) error` metodini yoz. Qoidalar:
   - `qty <= 0` bo'lsa — xato qaytar
   - Agar `stock - qty < 0` bo'lsa — xato qaytar (**manfiy stock bo'lishi mumkin emas**)
   - Agar `IsActive == false` bo'lsa — xato qaytar ("cannot sell inactive product")
4. `Deactivate()` metodini yoz — hech qanday parametrsiz, oddiy holatni o'zgartiradi
5. Kerakli getter metodlarini qo'sh (`Stock()`, `Price()`, `IsActive()`)

Bu kodni menga yubor — men tekshiraman va **nima uchun** har bir qaror to'g'ri yoki noto'g'ri ekanini tushuntiraman.

---

## 3️⃣ Nega bu modul shunchalik muhim

Bu — butun DDD falsafasining **poydevori**. Agar sen shu paradigma o'zgarishini (Anemic → Rich) chuqur his qilmasang:

- **Module 3 (Value Object)** senga tushunarsiz bo'lib qoladi — nega `Money`ni oddiy `int64` sifatida emas, alohida tur sifatida yozamiz?
- **Module 4 (Aggregate)** senga qiyin bo'ladi — nega "aggregate ichida faqat root orqali o'zgartirish mumkin" degan qoida bor?
- **Real production kod bazasi**da (ayniqsa fintech, e-commerce, logistika kabi sohalarda) Anemic model — bu **eng ko'p uchraydigan texnik qarz manbai**. Middle/Senior intervyularida "DDD bilasizmi" degan savolga aksariyat odamlar faqat atamalarni yodlagan, lekin **nega** degan savolga javob berolmaydi. Sen esa aynan shu "nega"ni tushunib olyapsan — bu seni boshqalardan ajratib turadi.

**Muhim ogohlantirish (senior maslahat):** DDD — bu **har bir loyihaga kerak emas**. Agar sening domening oddiy CRUD bo'lsa (masalan, blog, oddiy todo-list), Rich Domain Model ortiqcha murakkablik (over-engineering) qo'shadi. DDD **murakkab biznes qoidalari bo'lgan** domenlar uchun (bank, sug'urta, logistika, marketplace) mos keladi. Bu haqda ko'proq Module 10'da (anti-patterns) gaplashamiz.

---

## 4️⃣ O'zini tekshirish: Task + 5 ta savol

### 🎯 Yakuniy Task
Yuqoridagi **Product** topshirig'ini bajar va menga to'liq kodni yubor.

### ❓ 5 ta savol (javoblaringni yoz, men tekshiraman)

1. **Nima uchun** `Order` structidagi maydonlar kichik harf bilan yozilgan, va bu TypeScript'dagi `private` keyword'idan nima bilan farq qiladi?

2. Quyidagi kodda qanday muammo bor? Tushuntir:
```go
type Order struct {
	ID     string
	Status string
}
// boshqa faylda:
order.Status = "paid"
```

3. Nega `Money` uchun `float64` emas, `int64` (tiyinlarda) ishlatiladi? Bu qanday real muammoning oldini oladi?

4. Agar `BankAccount` structida `Withdraw` metodi bo'lmasa va uning o'rniga tashqi service quyidagicha yozsa:
```go
account.balance -= amount // balance private bo'lsa, bu compile bo'lmaydi
```
Bu compile xatosi DDD nuqtai nazaridan **nega yaxshi narsa** (feature, bug emas)?

5. Anemic Domain Model qachon **to'g'ri tanlov** bo'lishi mumkin? Har doim Rich Domain Model ishlatish kerakmi? Fikringni asosla.

---

**Javoblaringni va Product task kodingni yubor — men har birini alohida tahlil qilib, xato bo'lsa aniq ko'rsataman, to'g'ri bo'lsa nega to'g'riligini chuqurlashtiraman. Shundan keyin Module 2 (Ubiquitous Language va Bounded Context) ga o'tamiz.**

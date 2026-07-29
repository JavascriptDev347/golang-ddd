# Module 1: Anemic vs Rich Domain Model

Nazariya to'liq: [`../docs/01-anemic-vs-rich-domain-model.md`](../docs/01-anemic-vs-rich-domain-model.md)

## Papka strukturasi

```
module01-anemic-vs-rich/
├── domain/              <- TAYYOR yechim (Order aggregate). O'qi, testlarni ishga tushir, tushunib ol.
│   ├── order.go
│   └── order_test.go
└── exercise/             <- SENING topshiring (Product). TODO'larni bajar.
    ├── product.go
    └── product_test.go
```

## Ishga tushirish

Tayyor misolni ko'rish uchun:
```bash
go test ./module01-anemic-vs-rich/domain/... -v
```
Barcha testlar PASS bo'lishi kerak — bu senga Rich Domain Model qanday ishlashini ko'rsatadi.

O'zingning topshirig'ing uchun:
```bash
go test ./module01-anemic-vs-rich/exercise/... -v
```
Boshida bu **compile xatosi** beradi — bu normal holat (TDD "qizil holat").
`exercise/product.go` faylidagi TODO'larni birma-bir bajarib bor, har safar
testni qayta ishga tushirib tekshirib tur, toki barchasi yashil (PASS) bo'lguncha.

## Tayyor bo'lgach

Kodingni menga yubor (yoki shu faylni ko'rsat) — men review qilaman:
nima to'g'ri, nima antipattern, va nega.

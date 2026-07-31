# 📘 MODULE 2: Ubiquitous Language va Bounded Context

## Nazariya (qisqacha)

**Ubiquitous Language** — domen ekspertlari (biznes) ishlatadigan atamalar
kodda **aynan o'sha nom bilan** yashashi kerak. Agar sotuv menejeri
"qualify" deydi — kodda `Qualify()` bo'lishi kerak, `Approve()` yoki
`SetFlag(1)` emas. Bu — kod bilan biznes o'rtasidagi tarjima yo'qotilishini
oldini oladi.

**Bounded Context** — katta tizimda BITTA universal `Customer`/`User` model
bo'lmaydi. Har bir context (Sales, Billing, Support...) bir xil real
ob'yektni **o'ziga kerak bo'lgan qismi va o'z tili bilan** modellashtiradi.
Bu context'lar orasidagi aloqa faqat **Anti-Corruption Layer (ACL)** orqali,
DTO'lar bilan, hech qachon bir-birining domain structini import qilmasdan
amalga oshadi.

Ushbu modulda ikkita Bounded Context bilan ishlaysan:

```
domain/
├── salescontext/     — SALES context: "Lead", Qualify/Disqualify
└── billingcontext/   — BILLING context: "Payer", TaxID/PaymentMethod
                         + acl.go — ikki context orasidagi YAGONA ko'prik
```

`Lead` va `Payer` — bir xil real inson haqida, lekin ikkalasi ham bir-birini
IMPORT QILMAYDI. Bog'lanish faqat `LeadSnapshot` DTO orqali, `acl.go` faylida.

## Struktura

```
module02-ubiquitous-language-bounded-context/
├── README.md              <- shu fayl
├── domain/                <- TAYYOR YECHIM (agar tiqilib qolsang, shu yerga qara)
│   ├── salescontext/
│   │   ├── lead.go
│   │   └── lead_test.go
│   └── billingcontext/
│       ├── payer.go
│       ├── acl.go
│       └── payer_test.go
└── exercise/               <- SENING TOPSHIRIG'ING (TODO'lar bilan)
    ├── salescontext/
    │   ├── lead.go
    │   └── lead_test.go
    └── billingcontext/
        ├── payer.go
        ├── acl.go
        └── payer_test.go
```

## Vazifa

`exercise/` papkasidagi har bir TODO'ni bajar. Tayyor bo'lgach:

```bash
go test ./module02-ubiquitous-language-bounded-context/exercise/... -v
```

Barcha testlar PASS bo'lishi kerak. Hozircha ular **RED** (build failed) —
bu normal holat, chunki `TODO 1` (fieldlarni private qilish) va getter
metodlari hali bir-biriga zid.

## Qo'shimcha vazifa (chuqurroq tushunish uchun)

`exercise/billingcontext/acl.go` faylini yozib bo'lgach, o'zingdan so'ra:

> Agar men shu yerda `NewPayerFromQualifiedLead` funksiyasiga
> `*salescontext.Lead` turini to'g'ridan-to'g'ri parametr qilib bersam,
> nima buziladi?

Javob: Sales context o'z `Lead` structini o'zgartirganda (masalan yangi
majburiy maydon qo'shsa), Billing context ham compile bo'lmay qoladi —
garchi Billing bilan sales o'zgarishi umuman aloqador bo'lmasa ham. Aynan
shu — Bounded Context chegarasi buzilishining o'zi.

## Tekshiruv savollari

1. Nega Sales va Billing context'lari bir xil odam uchun ikkita alohida
   struct (`Lead` va `Payer`) ishlatadi — bu kod takrorlanishi emasmi?
2. `acl.go` faylida nega `salescontext` paketi import qilinmagan?
3. `LeadSnapshot` — bu Value Object'mi yoki DTO'mi? Farqi nimada?
4. Agar Support bo'limi ham shu odam bilan ishlashni boshlasa (masalan
   `SupportContact` degan o'z modeli bilan), bu nechta Bounded Context
   bo'lib qoladi?
5. `wasQualified bool` parametrini olib tashlab, o'rniga
   `snap.LeadStatus == "qualified"` deb tekshirsak nima uchun bu yomon
   dizayn bo'lardi?

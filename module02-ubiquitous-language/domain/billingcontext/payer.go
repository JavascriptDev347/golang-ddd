// Package billingcontext — BILLING BOUNDED CONTEXT (tayyor yechim).
//
// Diqqat: bu yerda "Lead" so'zi yo'q, "qualified" degan tushuncha ham yo'q.
// Billing uchun "u odam" — Payer. Billing faqat: kimga hisob-faktura
// yuborish, qanday to'lov usuli, soliq raqami — shular bilan qiziqadi.
//
// Ubiquitous Language (Billing/Finance domenidan): Payer, TaxID,
// PaymentMethod, Deactivate.
package billingcontext

import "errors"

var (
	ErrInvalidBillingName         = errors.New("payer: billing name is required")
	ErrInvalidTaxID               = errors.New("payer: tax id is required")
	ErrInvalidPaymentMethod       = errors.New("payer: unsupported payment method")
	ErrPayerInactive              = errors.New("payer: cannot charge an inactive payer")
	ErrLeadNotQualifiedForBilling = errors.New("payer: cannot create payer from a non-qualified lead")
)

type PaymentMethod string

const (
	PaymentMethodUnset PaymentMethod = ""
	PaymentMethodCard  PaymentMethod = "card"
	PaymentMethodBank  PaymentMethod = "bank_transfer"
)

// Payer — Billing Bounded Context'ning Aggregate Root'i.
// Bu struct salescontext.Lead haqida BILMAYDI ham — bu ataylab qilingan.
// Ikki context orasidagi bog'liqlik faqat quyidagi acl.go fayli orqali,
// va faqat bitta yo'nalishda (billing -> tashqi ma'lumot) o'tadi.
type Payer struct {
	id            string
	billingName   string
	taxID         string
	paymentMethod PaymentMethod
	active        bool
}

func NewPayer(id, billingName, taxID string) (*Payer, error) {
	if billingName == "" {
		return nil, ErrInvalidBillingName
	}
	if taxID == "" {
		return nil, ErrInvalidTaxID
	}
	return &Payer{
		id:            id,
		billingName:   billingName,
		taxID:         taxID,
		paymentMethod: PaymentMethodUnset,
		active:        true,
	}, nil
}

// SetPaymentMethod — faqat billing contextiga tegishli qoida: qo'llab-
// quvvatlanadigan usullar ro'yxati bilan cheklangan.
func (p *Payer) SetPaymentMethod(m PaymentMethod) error {
	if m != PaymentMethodCard && m != PaymentMethodBank {
		return ErrInvalidPaymentMethod
	}
	p.paymentMethod = m
	return nil
}

// Charge — soddalashtirilgan misol: nofaol Payer'dan pul yechib bo'lmaydi.
func (p *Payer) Charge(amount int64) error {
	if !p.active {
		return ErrPayerInactive
	}
	if p.paymentMethod == PaymentMethodUnset {
		return ErrInvalidPaymentMethod
	}
	// haqiqiy loyihada bu yerda to'lov provayderi chaqirilardi
	return nil
}

func (p *Payer) Deactivate() {
	p.active = false
}

func (p *Payer) ID() string                   { return p.id }
func (p *Payer) BillingName() string          { return p.billingName }
func (p *Payer) TaxID() string                { return p.taxID }
func (p *Payer) PaymentMethod() PaymentMethod { return p.paymentMethod }
func (p *Payer) IsActive() bool               { return p.active }

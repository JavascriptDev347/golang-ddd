// Package salescontext — SALES BOUNDED CONTEXT (tayyor yechim).
//
// Bu context faqat "Lead" (potentsial mijoz) haqida qayg'uradi.
// E'tibor ber: bu yerda "mijoz" so'zi umuman ishlatilmaydi — chunki
// Sales bo'limi uchun bu odam hali mijoz emas, u — Lead.
//
// Ubiquitous Language (Sales domenidan): Lead, Qualify, Disqualify, Source.
// Bu terminlar aynan shu tarzda domen ekspertlari (sotuv menejerlari)
// bilan suhbatda ishlatiladi — kod ularni o'zgartirmaydi.
package salescontext

import "errors"

var (
	ErrInvalidFullName      = errors.New("lead: full name is required")
	ErrInvalidSource        = errors.New("lead: source is required")
	ErrNegativeValue        = errors.New("lead: estimated value cannot be negative")
	ErrLeadAlreadyQualified = errors.New("lead: already qualified")
	ErrLeadDisqualified     = errors.New("lead: cannot qualify a disqualified lead")
	ErrEmptyReason          = errors.New("lead: disqualify reason is required")
)

// LeadStatus — Sales contextining o'z hayot sikli. Bu status Billing
// contextiga umuman aloqador emas — Billing "qualified" nima ekanini
// bilishi shart emas, u faqat Payer haqida qayg'uradi.
type LeadStatus string

const (
	LeadStatusNew          LeadStatus = "new"
	LeadStatusQualified    LeadStatus = "qualified"
	LeadStatusDisqualified LeadStatus = "disqualified"
)

// Lead — Sales Bounded Context'ning Aggregate Root'i.
// Diqqat: bu struct billing yoki support context'lari uchun umuman
// mos emas va mos bo'lishi ham shart emas — u faqat Sales uchun ishlaydi.
type Lead struct {
	id             string
	fullName       string
	source         string // masalan: "referral", "ads", "cold-call"
	estimatedValue int64  // tiyinlarda
	status         LeadStatus
}

// NewLead — yangi Lead har doim LeadStatusNew holatida tug'iladi.
func NewLead(id, fullName, source string, estimatedValue int64) (*Lead, error) {
	if fullName == "" {
		return nil, ErrInvalidFullName
	}
	if source == "" {
		return nil, ErrInvalidSource
	}
	if estimatedValue < 0 {
		return nil, ErrNegativeValue
	}
	return &Lead{
		id:             id,
		fullName:       fullName,
		source:         source,
		estimatedValue: estimatedValue,
		status:         LeadStatusNew,
	}, nil
}

// Qualify — "bu lead haqiqiy mijoz bo'la oladi" degan sotuv qarori.
// Business rule: disqualified lead qayta qualify qilinmaydi (bu ataylab
// qilingan — real hayotda "signal yo'qotilgan" degani, qayta tiklash uchun
// alohida jarayon kerak, oddiy qayta-qualify emas).
func (l *Lead) Qualify() error {
	if l.status == LeadStatusDisqualified {
		return ErrLeadDisqualified
	}
	if l.status == LeadStatusQualified {
		return ErrLeadAlreadyQualified
	}
	l.status = LeadStatusQualified
	return nil
}

// Disqualify — reason har doim talab qilinadi, chunki sotuv jamoasi
// keyinchalik "nega yo'qotdik" degan tahlil qiladi (ubiquitous language:
// bu "loss reason" deb ataladi sotuv bo'limida).
func (l *Lead) Disqualify(reason string) error {
	if reason == "" {
		return ErrEmptyReason
	}
	l.status = LeadStatusDisqualified
	return nil
}

func (l *Lead) ID() string            { return l.id }
func (l *Lead) FullName() string      { return l.fullName }
func (l *Lead) Source() string        { return l.source }
func (l *Lead) EstimatedValue() int64 { return l.estimatedValue }
func (l *Lead) Status() LeadStatus    { return l.status }
func (l *Lead) IsQualified() bool     { return l.status == LeadStatusQualified }

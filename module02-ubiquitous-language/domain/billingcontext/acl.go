// Anti-Corruption Layer (ACL) — bu fayl ikkita Bounded Context orasidagi
// YAGONA ko'prik.
//
// MUHIM QOIDA: bu faylda salescontext paketi HECH QACHON import qilinmaydi.
// Agar biz shunday qilsak (masalan `func NewPayerFromLead(l *salescontext.Lead)`),
// billing context sales context'ning ichki tuzilishiga BOG'LANIB QOLADI —
// va sales context o'z Lead structini o'zgartirganda (masalan yangi maydon
// qo'shsa yoki Source turini o'zgartirsa), billing context ham buzilib qoladi.
// Bu — "Bounded Context chegarasi buzilgan" degani, va aynan shu narsa
// katta monolitlarni asta-sekin "Big Ball of Mud"ga aylantiradi.
//
// TO'G'RI YECHIM: chaqiruvchi tomon (application/use-case layer) sales
// context'dan kerakli ma'lumotni oladi va uni oddiy DTO (LeadSnapshot)
// ko'rinishida shu yerga "tarjima qilib" uzatadi. Billing context faqat
// shu DTO bilan ishlaydi — sales context haqida hech narsa bilmaydi.
package billingcontext

// LeadSnapshot — Anti-Corruption Layer uchun DTO (Data Transfer Object).
// Bu salescontext.Lead EMAS — bu faqat billing uchun zarur bo'lgan
// ma'lumotlarning suratga olingan nusxasi (masalan alohida so'rov orqali
// olingan TaxID bilan birga).
type LeadSnapshot struct {
	LeadID   string
	FullName string
	TaxID    string
}

// NewPayerFromQualifiedLead — Payer yaratishning YAGONA yo'li shu orqali,
// va bu funksiya billing contextining O'Z qoidasini qo'llaydi: faqat
// "qualified" statusidagi lead uchun Payer yaratishga ruxsat beriladi.
//
// Diqqat qil: "qualified" ekanligini bu funksiya o'zi solesscontext orqali
// TEKSHIRMAYDI (chunki u sales contextini import qilmaydi!). Buning o'rniga,
// chaqiruvchi (application layer) qaror qabul qiladi va faqat kerakli
// ma'lumotni uzatadi. wasQualified flag orqali bu qaror aniq ko'rinadi.
func NewPayerFromQualifiedLead(snap LeadSnapshot, wasQualified bool) (*Payer, error) {
	if !wasQualified {
		return nil, ErrLeadNotQualifiedForBilling
	}
	return NewPayer(snap.LeadID, snap.FullName, snap.TaxID)
}

package salescontext

import (
	"errors"
	"testing"
)

func TestNewLead(t *testing.T) {
	tests := []struct {
		name    string
		id      string
		full    string
		source  string
		value   int64
		wantErr error
	}{
		{"valid", "l1", "Aziz Karimov", "referral", 100000, nil},
		{"empty name", "l2", "", "referral", 100000, ErrInvalidFullName},
		{"empty source", "l3", "Aziz", "", 100000, ErrInvalidSource},
		{"negative value", "l4", "Aziz", "referral", -1, ErrNegativeValue},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			lead, err := NewLead(tt.id, tt.full, tt.source, tt.value)
			if !errors.Is(err, tt.wantErr) {
				t.Fatalf("got err=%v, want=%v", err, tt.wantErr)
			}
			if tt.wantErr == nil {
				if lead.Status() != LeadStatusNew {
					t.Fatalf("new lead must start as New, got %v", lead.Status())
				}
			}
		})
	}
}

func TestLead_Qualify(t *testing.T) {
	lead, _ := NewLead("l1", "Aziz", "ads", 1000)

	if err := lead.Qualify(); err != nil {
		t.Fatalf("unexpected error qualifying new lead: %v", err)
	}
	if !lead.IsQualified() {
		t.Fatalf("lead should be qualified")
	}

	if err := lead.Qualify(); !errors.Is(err, ErrLeadAlreadyQualified) {
		t.Fatalf("re-qualifying should fail with ErrLeadAlreadyQualified, got %v", err)
	}
}

func TestLead_DisqualifyThenQualifyFails(t *testing.T) {
	lead, _ := NewLead("l1", "Aziz", "ads", 1000)

	if err := lead.Disqualify(""); !errors.Is(err, ErrEmptyReason) {
		t.Fatalf("empty reason must fail, got %v", err)
	}

	if err := lead.Disqualify("budget cut"); err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if err := lead.Qualify(); !errors.Is(err, ErrLeadDisqualified) {
		t.Fatalf("qualifying a disqualified lead must fail, got %v", err)
	}
}

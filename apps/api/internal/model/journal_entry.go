package model

import (
	"time"
)

type JournalEntry struct {
	JournalId string `json:"journal_id"`
	EntryNumber string `json:"entry_number"`
	EntryDate time.Time `json:"entry_date"`
	Description string `json:"description"`
	SourceModule string `json:"source_module"`
	SourceReferenceId *string `json:"source_reference_id,omitempty"`
	IsPosted *bool `json:"is_posted,omitempty"`
	CreatedBy *string `json:"created_by,omitempty"`
	UpdatedBy *string `json:"updated_by,omitempty"`
	DeletedBy *string `json:"deleted_by"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt *time.Time `json:"updated_at,omitempty"`
	DeletedAt *time.Time `json:"deleted_at"`
	Lines []JournalLines `json:"lines,omitempty"`
}

type JournalLines struct {
	LineId string `json:"line_id"`
	JournalId string `json:"journal_id"`
	AccountId string `json:"account_id"`
	Debit *float64 `json:"debit,omitempty"`
	Credit *float64 `json:"credit,omitempty"`
	Memo *string `json:"memo,omitempty"`
	CreatedAt time.Time `json:"created_at"`
}

type JournalEntryRequest struct {
	EntryNumber *string `json:"entry_number"`
	EntryDate *time.Time `json:"entry_date"`
	Description *string `json:"description"`
	SourceModule *string `json:"source_module"`
	SourceReferenceId *string `json:"source_reference_id"`
	IsPosted *bool `json:"is_posted"`
	Lines []JournalLines `json:"lines"`
}

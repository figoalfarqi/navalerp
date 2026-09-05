package model

import (
	"time"
)

type MilitaryRank struct {
	RankId string `json:"rank_id"`
	RankCode string `json:"rank_code"`
	RankName string `json:"rank_name"`
	RankCategory string `json:"rank_category"`
	NatoRankCode *string `json:"nato_rank_code,omitempty"`
	SeniorityOrder int `json:"seniority_order"`
	IsActive *bool `json:"is_active,omitempty"`
	CreatedAt time.Time `json:"created_at"`
}

type MilitaryRankRequest struct {
	RankCode *string `json:"rank_code"`
	RankName *string `json:"rank_name"`
	RankCategory *string `json:"rank_category"`
	NatoRankCode *string `json:"nato_rank_code"`
	SeniorityOrder *int `json:"seniority_order"`
	IsActive *bool `json:"is_active"`
}

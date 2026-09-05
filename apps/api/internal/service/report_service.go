package service

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/figoalfarqi/navalerp/internal/model"
	"github.com/figoalfarqi/navalerp/internal/repository"
)

type ReportService struct{ Repo *repository.ReportRepository }

func NewReportService(repo *repository.ReportRepository) *ReportService {
	return &ReportService{Repo: repo}
}

func (s *ReportService) Get(ctx context.Context, projectID *int, period, dateValue string) (*model.ReportResponse, error) {
	jakarta := time.FixedZone("Asia/Jakarta", 7*60*60)
	if period == "" {
		period = "monthly"
	}
	if period != "weekly" && period != "monthly" {
		return nil, errors.New("period must be weekly or monthly")
	}
	if dateValue == "" {
		dateValue = time.Now().In(jakarta).Format("2006-01-02")
	}
	date, err := time.ParseInLocation("2006-01-02", dateValue, jakarta)
	if err != nil {
		return nil, errors.New("date must use YYYY-MM-DD")
	}
	var from, to time.Time
	var label string
	if period == "weekly" {
		offset := (int(date.Weekday()) + 6) % 7
		from = date.AddDate(0, 0, -offset)
		to = from.AddDate(0, 0, 7)
		label = fmt.Sprintf("%s - %s", from.Format("02 Jan 2006"), to.AddDate(0, 0, -1).Format("02 Jan 2006"))
	} else {
		from = time.Date(date.Year(), date.Month(), 1, 0, 0, 0, 0, date.Location())
		to = from.AddDate(0, 1, 0)
		label = from.Format("January 2006")
	}
	return s.Repo.Build(ctx, projectID, from, to, period, label)
}

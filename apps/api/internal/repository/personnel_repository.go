package repository

import (
	"context"
	"fmt"
	"strings"

	"github.com/figoalfarqi/navalerp/internal/model"
	"github.com/jackc/pgx/v5/pgxpool"
)

type PersonnelRepository struct {
	DB *pgxpool.Pool
}

func NewPersonnelRepository(db *pgxpool.Pool) *PersonnelRepository {
	return &PersonnelRepository{DB: db}
}

// Get retrieves a single personnel by personnel_id
func (r *PersonnelRepository) Get(ctx context.Context, id string) (*model.Personnel, error) {
	query := `SELECT personnel_id, nrp, full_name, rank_id, corps_id, current_unit_id, current_position, birth_place, birth_date, gender, blood_type, religion, education_level, service_entry_date, user_id, status, created_by, updated_by, deleted_by, created_at, updated_at, deleted_at FROM hcm_personnel WHERE personnel_id = $1 AND deleted_at IS NULL`

	var m model.Personnel
	err := r.DB.QueryRow(ctx, query, id).Scan(&m.PersonnelId, &m.Nrp, &m.FullName, &m.RankId, &m.CorpsId, &m.CurrentUnitId, &m.CurrentPosition, &m.BirthPlace, &m.BirthDate, &m.Gender, &m.BloodType, &m.Religion, &m.EducationLevel, &m.ServiceEntryDate, &m.UserId, &m.Status, &m.CreatedBy, &m.UpdatedBy, &m.DeletedBy, &m.CreatedAt, &m.UpdatedAt, &m.DeletedAt)
	if err != nil {
		return nil, err
	}

	// Load ServiceRecords
	childRowsServiceRecords, err := r.DB.Query(ctx, `SELECT record_id, personnel_id, order_letter_number, assignment_type, from_unit_id, to_unit_id, position_title, start_date, end_date, remarks, created_by, created_at FROM hcm_service_records WHERE personnel_id = $1`, id)
	if err == nil {
		defer childRowsServiceRecords.Close()
		for childRowsServiceRecords.Next() {
			var item model.ServiceRecords
			if err := childRowsServiceRecords.Scan(&item.RecordId, &item.PersonnelId, &item.OrderLetterNumber, &item.AssignmentType, &item.FromUnitId, &item.ToUnitId, &item.PositionTitle, &item.StartDate, &item.EndDate, &item.Remarks, &item.CreatedBy, &item.CreatedAt); err == nil {
				m.ServiceRecords = append(m.ServiceRecords, item)
			}
		}
	}

	// Load Qualifications
	childRowsQualifications, err := r.DB.Query(ctx, `SELECT personnel_qual_id, personnel_id, qualification_id, certificate_number, obtained_date, valid_until, is_active, created_at FROM hcm_personnel_qualifications WHERE personnel_id = $1`, id)
	if err == nil {
		defer childRowsQualifications.Close()
		for childRowsQualifications.Next() {
			var item model.PersonnelQualifications
			if err := childRowsQualifications.Scan(&item.PersonnelQualId, &item.PersonnelId, &item.QualificationId, &item.CertificateNumber, &item.ObtainedDate, &item.ValidUntil, &item.IsActive, &item.CreatedAt); err == nil {
				m.Qualifications = append(m.Qualifications, item)
			}
		}
	}

	// Load MedicalReadiness
	childRowsMedicalReadiness, err := r.DB.Query(ctx, `SELECT medical_id, personnel_id, examination_date, stakes_category, physical_fitness_score, vision_status, dental_status, cardio_status, general_health_status, doctor_remarks, valid_until, created_by, created_at FROM hcm_medical_readiness WHERE personnel_id = $1`, id)
	if err == nil {
		defer childRowsMedicalReadiness.Close()
		for childRowsMedicalReadiness.Next() {
			var item model.MedicalReadiness
			if err := childRowsMedicalReadiness.Scan(&item.MedicalId, &item.PersonnelId, &item.ExaminationDate, &item.StakesCategory, &item.PhysicalFitnessScore, &item.VisionStatus, &item.DentalStatus, &item.CardioStatus, &item.GeneralHealthStatus, &item.DoctorRemarks, &item.ValidUntil, &item.CreatedBy, &item.CreatedAt); err == nil {
				m.MedicalReadiness = append(m.MedicalReadiness, item)
			}
		}
	}

	return &m, nil
}

// List retrieves paginated personnel records
func (r *PersonnelRepository) List(ctx context.Context, opts model.ListOptions) ([]model.Personnel, int, error) {
	whereClauses := []string{"1=1"}
	var args []any
	argPos := 1

	whereClauses = append(whereClauses, "deleted_at IS NULL")
	if opts.Search != "" {
		searchPattern := "%" + opts.Search + "%"
		whereClauses = append(whereClauses, fmt.Sprintf("(nrp ILIKE $%[1]d OR full_name ILIKE $%[1]d OR current_position ILIKE $%[1]d)", argPos))
		args = append(args, searchPattern)
		argPos++
	}

	whereSql := strings.Join(whereClauses, " AND ")
	countQuery := fmt.Sprintf("SELECT COUNT(*) FROM hcm_personnel WHERE %s", whereSql)
	var total int
	if err := r.DB.QueryRow(ctx, countQuery, args...).Scan(&total); err != nil {
		return nil, 0, err
	}

	limit := opts.Limit
	if limit <= 0 { limit = 10 }
	offset := opts.Offset
	if offset < 0 { offset = 0 }

	listQuery := fmt.Sprintf("SELECT personnel_id, nrp, full_name, rank_id, corps_id, current_unit_id, current_position, birth_place, birth_date, gender, blood_type, religion, education_level, service_entry_date, user_id, status, created_by, updated_by, deleted_by, created_at, updated_at, deleted_at FROM hcm_personnel WHERE %s ORDER BY created_at DESC LIMIT $%d OFFSET $%d", whereSql, argPos, argPos+1)
	args = append(args, limit, offset)

	rows, err := r.DB.Query(ctx, listQuery, args...)
	if err != nil {
		return nil, 0, err
	}
	defer rows.Close()

	var items []model.Personnel
	for rows.Next() {
		var m model.Personnel
		if err := rows.Scan(&m.PersonnelId, &m.Nrp, &m.FullName, &m.RankId, &m.CorpsId, &m.CurrentUnitId, &m.CurrentPosition, &m.BirthPlace, &m.BirthDate, &m.Gender, &m.BloodType, &m.Religion, &m.EducationLevel, &m.ServiceEntryDate, &m.UserId, &m.Status, &m.CreatedBy, &m.UpdatedBy, &m.DeletedBy, &m.CreatedAt, &m.UpdatedAt, &m.DeletedAt); err != nil {
			return nil, 0, err
		}
		items = append(items, m)
	}

	return items, total, nil
}

// Create inserts a new personnel with optional child items
func (r *PersonnelRepository) Create(ctx context.Context, m *model.Personnel) (string, error) {
	tx, err := r.DB.Begin(ctx)
	if err != nil {
		return "", err
	}
	defer tx.Rollback(ctx)

	insertQuery := `INSERT INTO hcm_personnel (nrp, full_name, rank_id, corps_id, current_unit_id, current_position, birth_place, birth_date, gender, blood_type, religion, education_level, service_entry_date, user_id, status, created_by, updated_by, deleted_by) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9::personnel_gender_type, $10::blood_type_enum, $11::religion_type, $12, $13, $14, $15::personnel_status_type, $16, $17, $18) RETURNING personnel_id`
	var newID string
	err = tx.QueryRow(ctx, insertQuery, m.Nrp, m.FullName, m.RankId, m.CorpsId, m.CurrentUnitId, m.CurrentPosition, m.BirthPlace, m.BirthDate, m.Gender, m.BloodType, m.Religion, m.EducationLevel, m.ServiceEntryDate, m.UserId, m.Status, m.CreatedBy, m.UpdatedBy, m.DeletedBy).Scan(&newID)
	if err != nil {
		return "", err
	}

	for _, item := range m.ServiceRecords {
		_, err := tx.Exec(ctx, `INSERT INTO hcm_service_records (personnel_id, order_letter_number, assignment_type, from_unit_id, to_unit_id, position_title, start_date, end_date, remarks, created_by) VALUES ($1, $2, $3::service_assignment_type, $4, $5, $6, $7, $8, $9, $10)`, newID, item.OrderLetterNumber, item.AssignmentType, item.FromUnitId, item.ToUnitId, item.PositionTitle, item.StartDate, item.EndDate, item.Remarks, item.CreatedBy)
		if err != nil {
			return "", err
		}
	}

	for _, item := range m.Qualifications {
		_, err := tx.Exec(ctx, `INSERT INTO hcm_personnel_qualifications (personnel_id, qualification_id, certificate_number, obtained_date, valid_until, is_active) VALUES ($1, $2, $3, $4, $5, $6)`, newID, item.QualificationId, item.CertificateNumber, item.ObtainedDate, item.ValidUntil, item.IsActive)
		if err != nil {
			return "", err
		}
	}

	for _, item := range m.MedicalReadiness {
		_, err := tx.Exec(ctx, `INSERT INTO hcm_medical_readiness (personnel_id, examination_date, stakes_category, physical_fitness_score, vision_status, dental_status, cardio_status, general_health_status, doctor_remarks, valid_until, created_by) VALUES ($1, $2, $3::urikes_stakes_type, $4, $5, $6, $7, $8, $9, $10, $11)`, newID, item.ExaminationDate, item.StakesCategory, item.PhysicalFitnessScore, item.VisionStatus, item.DentalStatus, item.CardioStatus, item.GeneralHealthStatus, item.DoctorRemarks, item.ValidUntil, item.CreatedBy)
		if err != nil {
			return "", err
		}
	}

	if err := tx.Commit(ctx); err != nil {
		return "", err
	}
	return newID, nil
}

// Update updates an existing personnel
func (r *PersonnelRepository) Update(ctx context.Context, id string, m *model.Personnel) error {
	tx, err := r.DB.Begin(ctx)
	if err != nil {
		return err
	}
	defer tx.Rollback(ctx)

	updateQuery := `UPDATE hcm_personnel SET nrp = $1, full_name = $2, rank_id = $3, corps_id = $4, current_unit_id = $5, current_position = $6, birth_place = $7, birth_date = $8, gender = $9::personnel_gender_type, blood_type = $10::blood_type_enum, religion = $11::religion_type, education_level = $12, service_entry_date = $13, user_id = $14, status = $15::personnel_status_type, updated_by = $16, updated_at = CURRENT_TIMESTAMP WHERE personnel_id = $17`
	_, err = tx.Exec(ctx, updateQuery, m.Nrp, m.FullName, m.RankId, m.CorpsId, m.CurrentUnitId, m.CurrentPosition, m.BirthPlace, m.BirthDate, m.Gender, m.BloodType, m.Religion, m.EducationLevel, m.ServiceEntryDate, m.UserId, m.Status, m.UpdatedBy, id)
	if err != nil {
		return err
	}

	if len(m.ServiceRecords) > 0 {
		_, _ = tx.Exec(ctx, `DELETE FROM hcm_service_records WHERE personnel_id = $1`, id)
		for _, item := range m.ServiceRecords {
			_, err := tx.Exec(ctx, `INSERT INTO hcm_service_records (personnel_id, order_letter_number, assignment_type, from_unit_id, to_unit_id, position_title, start_date, end_date, remarks, created_by) VALUES ($1, $2, $3::service_assignment_type, $4, $5, $6, $7, $8, $9, $10)`, id, item.OrderLetterNumber, item.AssignmentType, item.FromUnitId, item.ToUnitId, item.PositionTitle, item.StartDate, item.EndDate, item.Remarks, item.CreatedBy)
			if err != nil {
				return err
			}
		}
	}

	if len(m.Qualifications) > 0 {
		_, _ = tx.Exec(ctx, `DELETE FROM hcm_personnel_qualifications WHERE personnel_id = $1`, id)
		for _, item := range m.Qualifications {
			_, err := tx.Exec(ctx, `INSERT INTO hcm_personnel_qualifications (personnel_id, qualification_id, certificate_number, obtained_date, valid_until, is_active) VALUES ($1, $2, $3, $4, $5, $6)`, id, item.QualificationId, item.CertificateNumber, item.ObtainedDate, item.ValidUntil, item.IsActive)
			if err != nil {
				return err
			}
		}
	}

	if len(m.MedicalReadiness) > 0 {
		_, _ = tx.Exec(ctx, `DELETE FROM hcm_medical_readiness WHERE personnel_id = $1`, id)
		for _, item := range m.MedicalReadiness {
			_, err := tx.Exec(ctx, `INSERT INTO hcm_medical_readiness (personnel_id, examination_date, stakes_category, physical_fitness_score, vision_status, dental_status, cardio_status, general_health_status, doctor_remarks, valid_until, created_by) VALUES ($1, $2, $3::urikes_stakes_type, $4, $5, $6, $7, $8, $9, $10, $11)`, id, item.ExaminationDate, item.StakesCategory, item.PhysicalFitnessScore, item.VisionStatus, item.DentalStatus, item.CardioStatus, item.GeneralHealthStatus, item.DoctorRemarks, item.ValidUntil, item.CreatedBy)
			if err != nil {
				return err
			}
		}
	}

	return tx.Commit(ctx)
}

// Delete removes or soft-deletes personnel
func (r *PersonnelRepository) Delete(ctx context.Context, id string) error {
	query := `UPDATE hcm_personnel SET deleted_at = CURRENT_TIMESTAMP WHERE personnel_id = $1 AND deleted_at IS NULL`
	_, err := r.DB.Exec(ctx, query, id)
	return err
}

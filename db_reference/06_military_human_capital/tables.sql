-- =============================================================================
-- MODUL 6: MANAJEMEN PERSONEL & AWAK KAPAL (MILITARY HUMAN CAPITAL)
-- FILE: 06_military_human_capital/tables.sql
-- =============================================================================

CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

-- -----------------------------------------------------------------------------
-- ENUM TYPES: MODUL 6
-- -----------------------------------------------------------------------------
DO $$ BEGIN
    CREATE TYPE rank_category_type AS ENUM (
        'PATI',     -- Perwira Tinggi
        'PAMEN',    -- Perwira Menengah
        'PAMA',     -- Perwira Pertama
        'BINTARA',
        'TAMTAMA'
    );
EXCEPTION WHEN duplicate_object THEN null; END $$;

DO $$ BEGIN
    CREATE TYPE corps_code_type AS ENUM (
        'P',   -- Pelaut
        'T',   -- Teknik
        'E',   -- Elektronika
        'S',   -- Suplai
        'M',   -- Marinir
        'K',   -- Kesehatan
        'KH',  -- Khusus
        'PM'   -- Polisi Militer
    );
EXCEPTION WHEN duplicate_object THEN null; END $$;

DO $$ BEGIN
    CREATE TYPE personnel_gender_type AS ENUM (
        'MALE',
        'FEMALE'
    );
EXCEPTION WHEN duplicate_object THEN null; END $$;

DO $$ BEGIN
    CREATE TYPE blood_type_enum AS ENUM (
        'A',
        'B',
        'AB',
        'O'
    );
EXCEPTION WHEN duplicate_object THEN null; END $$;

DO $$ BEGIN
    CREATE TYPE religion_type AS ENUM (
        'ISLAM',
        'PROTESTAN',
        'KATOLIK',
        'HINDU',
        'BUDDHA',
        'KONGHUCU'
    );
EXCEPTION WHEN duplicate_object THEN null; END $$;

DO $$ BEGIN
    CREATE TYPE personnel_status_type AS ENUM (
        'ACTIVE',
        'ON_LEAVE',
        'RETIRED',
        'DECEASED',
        'WIA',
        'KIA'
    );
EXCEPTION WHEN duplicate_object THEN null; END $$;

DO $$ BEGIN
    CREATE TYPE service_assignment_type AS ENUM (
        'MUTATION',
        'PROMOTION',
        'MISSION_DEPLOYMENT',
        'EDUCATION',
        'COMMENDATION'
    );
EXCEPTION WHEN duplicate_object THEN null; END $$;

DO $$ BEGIN
    CREATE TYPE qualification_category_type AS ENUM (
        'BREVET',
        'CERTIFICATION',
        'SPECIALIZATION',
        'SEAMANSHIP'
    );
EXCEPTION WHEN duplicate_object THEN null; END $$;

DO $$ BEGIN
    CREATE TYPE urikes_stakes_type AS ENUM (
        'STAKES_I',   -- Siap Tempur Penuh Tanpa Syarat
        'STAKES_II',  -- Siap Tempur Terbatas / Bersyarat Ringan
        'STAKES_III', -- Rawat Jalan / Tidak Siap Tempur Sementara
        'STAKES_IV'   -- Tidak Siap Tempur Tetap / Invalid
    );
EXCEPTION WHEN duplicate_object THEN null; END $$;

DO $$ BEGIN
    CREATE TYPE crew_department_type AS ENUM (
        'DEPOPS',   -- Departemen Operasi
        'DEPSIN',   -- Departemen Mesin
        'DEPLOG',   -- Departemen Logistik
        'DEPSENAU'  -- Departemen Senjata & Bahari
    );
EXCEPTION WHEN duplicate_object THEN null; END $$;

DO $$ BEGIN
    CREATE TYPE watch_bill_duty_type AS ENUM (
        'VIGOUR_A',
        'VIGOUR_B',
        'SIAGA_1',
        'COMBAT_STATION',
        'COMMAND_POST',
        'ENGINEERING_CONTROL_ROOM'
    );
EXCEPTION WHEN duplicate_object THEN null; END $$;

DO $$ BEGIN
    CREATE TYPE allowance_payment_status_type AS ENUM (
        'PENDING',
        'APPROVED',
        'PAID'
    );
EXCEPTION WHEN duplicate_object THEN null; END $$;

-- -----------------------------------------------------------------------------
-- TABEL: hcm_ranks
-- -----------------------------------------------------------------------------
CREATE TABLE IF NOT EXISTS hcm_ranks (
    rank_id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    rank_code VARCHAR(20) UNIQUE NOT NULL,
    rank_name VARCHAR(100) NOT NULL,
    rank_category rank_category_type NOT NULL,
    nato_rank_code VARCHAR(10),
    seniority_order INT NOT NULL,
    is_active BOOLEAN DEFAULT TRUE,
    created_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP
);

-- -----------------------------------------------------------------------------
-- TABEL: hcm_corps
-- -----------------------------------------------------------------------------
CREATE TABLE IF NOT EXISTS hcm_corps (
    corps_id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    corps_code corps_code_type UNIQUE NOT NULL,
    corps_name VARCHAR(100) NOT NULL,
    description TEXT,
    created_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP
);

-- -----------------------------------------------------------------------------
-- TABEL: hcm_personnel
-- -----------------------------------------------------------------------------
CREATE TABLE IF NOT EXISTS hcm_personnel (
    personnel_id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    nrp VARCHAR(30) UNIQUE NOT NULL,
    full_name VARCHAR(150) NOT NULL,
    rank_id UUID NOT NULL REFERENCES hcm_ranks(rank_id),
    corps_id UUID NOT NULL REFERENCES hcm_corps(corps_id),
    current_unit_id UUID NOT NULL REFERENCES org_units(unit_id),
    current_position VARCHAR(150) NOT NULL,
    birth_place VARCHAR(100),
    birth_date DATE NOT NULL,
    gender personnel_gender_type DEFAULT 'MALE',
    blood_type blood_type_enum,
    religion religion_type,
    education_level VARCHAR(50),
    service_entry_date DATE NOT NULL,
    user_id UUID UNIQUE REFERENCES sys_users(user_id),
    status personnel_status_type DEFAULT 'ACTIVE',
    created_by UUID,
    updated_by UUID,
    deleted_by UUID,
    created_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMPTZ,
    deleted_at TIMESTAMPTZ
);

-- -----------------------------------------------------------------------------
-- TABEL: hcm_service_records
-- -----------------------------------------------------------------------------
CREATE TABLE IF NOT EXISTS hcm_service_records (
    record_id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    personnel_id UUID NOT NULL REFERENCES hcm_personnel(personnel_id) ON DELETE CASCADE,
    order_letter_number VARCHAR(100) NOT NULL,
    assignment_type service_assignment_type NOT NULL,
    from_unit_id UUID REFERENCES org_units(unit_id),
    to_unit_id UUID NOT NULL REFERENCES org_units(unit_id),
    position_title VARCHAR(150) NOT NULL,
    start_date DATE NOT NULL,
    end_date DATE,
    remarks TEXT,
    created_by UUID,
    created_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP
);

-- -----------------------------------------------------------------------------
-- TABEL: hcm_qualifications
-- -----------------------------------------------------------------------------
CREATE TABLE IF NOT EXISTS hcm_qualifications (
    qualification_id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    qualification_code VARCHAR(50) UNIQUE NOT NULL,
    qualification_name VARCHAR(150) NOT NULL,
    qualification_category qualification_category_type NOT NULL,
    issuing_institution VARCHAR(150) NOT NULL,
    validity_years INT DEFAULT 5,
    description TEXT,
    created_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP
);

-- -----------------------------------------------------------------------------
-- TABEL: hcm_personnel_qualifications
-- -----------------------------------------------------------------------------
CREATE TABLE IF NOT EXISTS hcm_personnel_qualifications (
    personnel_qual_id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    personnel_id UUID NOT NULL REFERENCES hcm_personnel(personnel_id) ON DELETE CASCADE,
    qualification_id UUID NOT NULL REFERENCES hcm_qualifications(qualification_id),
    certificate_number VARCHAR(100),
    obtained_date DATE NOT NULL,
    valid_until DATE,
    is_active BOOLEAN DEFAULT TRUE,
    created_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
    CONSTRAINT uq_pers_qual UNIQUE (personnel_id, qualification_id)
);

-- -----------------------------------------------------------------------------
-- TABEL: hcm_medical_readiness
-- -----------------------------------------------------------------------------
CREATE TABLE IF NOT EXISTS hcm_medical_readiness (
    medical_id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    personnel_id UUID NOT NULL REFERENCES hcm_personnel(personnel_id) ON DELETE CASCADE,
    examination_date DATE NOT NULL,
    stakes_category urikes_stakes_type NOT NULL,
    physical_fitness_score NUMERIC(5, 2),
    vision_status VARCHAR(50) DEFAULT 'NORMAL',
    dental_status VARCHAR(50) DEFAULT 'FIT',
    cardio_status VARCHAR(50) DEFAULT 'FIT',
    general_health_status VARCHAR(50) NOT NULL,
    doctor_remarks TEXT,
    valid_until DATE NOT NULL,
    created_by UUID,
    created_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP
);

-- -----------------------------------------------------------------------------
-- TABEL: hcm_crew_assignments
-- -----------------------------------------------------------------------------
CREATE TABLE IF NOT EXISTS hcm_crew_assignments (
    assignment_id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    ship_id UUID NOT NULL REFERENCES mro_ships(ship_id),
    personnel_id UUID NOT NULL REFERENCES hcm_personnel(personnel_id),
    crew_role VARCHAR(100) NOT NULL,
    department crew_department_type NOT NULL,
    watch_bill_duty watch_bill_duty_type,
    assigned_date DATE NOT NULL,
    relieved_date DATE,
    is_active BOOLEAN DEFAULT TRUE,
    created_by UUID,
    created_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMPTZ
);

-- -----------------------------------------------------------------------------
-- TABEL: hcm_sea_duty_allowances
-- -----------------------------------------------------------------------------
CREATE TABLE IF NOT EXISTS hcm_sea_duty_allowances (
    allowance_id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    personnel_id UUID NOT NULL REFERENCES hcm_personnel(personnel_id),
    ship_id UUID NOT NULL REFERENCES mro_ships(ship_id),
    mission_name VARCHAR(150) NOT NULL,
    start_date DATE NOT NULL,
    end_date DATE NOT NULL,
    days_at_sea INT NOT NULL,
    daily_allowance_rate NUMERIC(18, 2) NOT NULL,
    total_allowance NUMERIC(18, 2) GENERATED ALWAYS AS (days_at_sea * daily_allowance_rate) STORED,
    payment_status allowance_payment_status_type DEFAULT 'APPROVED',
    payment_reference_no VARCHAR(50),
    created_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP
);

-- Indexes
CREATE INDEX IF NOT EXISTS idx_hcm_pers_nrp ON hcm_personnel(nrp) WHERE deleted_at IS NULL;
CREATE INDEX IF NOT EXISTS idx_hcm_pers_unit ON hcm_personnel(current_unit_id) WHERE deleted_at IS NULL;
CREATE INDEX IF NOT EXISTS idx_hcm_pers_rank ON hcm_personnel(rank_id) WHERE deleted_at IS NULL;
CREATE INDEX IF NOT EXISTS idx_hcm_crew_ship ON hcm_crew_assignments(ship_id) WHERE is_active = TRUE;
CREATE INDEX IF NOT EXISTS idx_hcm_crew_pers ON hcm_crew_assignments(personnel_id);
CREATE INDEX IF NOT EXISTS idx_hcm_med_pers ON hcm_medical_readiness(personnel_id);

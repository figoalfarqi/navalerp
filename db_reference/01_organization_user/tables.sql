-- =============================================================================
-- MODUL 1: ORGANISASI, PENGGUNA & COMMON MASTER DATA (NAVALERP)
-- FILE: 01_organization_user/tables.sql
-- =============================================================================

CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

-- -----------------------------------------------------------------------------
-- ENUM TYPES: MODUL 1
-- -----------------------------------------------------------------------------
DO $$ BEGIN
    CREATE TYPE org_unit_type AS ENUM (
        'HEADQUARTERS',
        'FLEET',
        'LANTAMAL',
        'LANAL',
        'SQUADRON',
        'SHIP_UNIT',
        'SHORE_BASE',
        'DEPOT',
        'FASHARKAN'
    );
EXCEPTION WHEN duplicate_object THEN null; END $$;

DO $$ BEGIN
    CREATE TYPE military_rank_type AS ENUM (
        'LAKSAMANA_TNI',
        'LAKSAMANA_MADYA_TNI',
        'LAKSAMANA_MUDA_TNI',
        'LAKSAMANA_PERTAMA_TNI',
        'KOLONEL_LAUT',
        'LETKOL_LAUT',
        'MAYOR_LAUT',
        'KAPTEN_LAUT',
        'LETTU_LAUT',
        'LETDA_LAUT',
        'PELTU',
        'PELDA',
        'SERMA',
        'SERKA',
        'SERTU',
        'SERDA',
        'KOPKA',
        'KOPTU',
        'KOPDA',
        'KLK',
        'KLS',
        'KLD',
        'PNS_TNI'
    );
EXCEPTION WHEN duplicate_object THEN null; END $$;

DO $$ BEGIN
    CREATE TYPE department_type AS ENUM (
        'SRENA',        -- Staf Perencanaan & Anggaran
        'SOPS',         -- Staf Operasi
        'SLOG',         -- Staf Logistik
        'SPERS',        -- Staf Personel
        'SPOTMAR',      -- Staf Potensi Maritim
        'KOMANDO',      -- Pimpinan / Panglima / Komandan
        'DEPOPS',       -- Departemen Operasi Kapal KRI
        'DEPSIN',       -- Departemen Mesin Kapal KRI
        'DEPLOG',       -- Departemen Logistik Kapal KRI
        'DEPSENAU',     -- Departemen Senjata & Bahari Kapal KRI
        'FASHARKAN',    -- Fasilitas Pemeliharaan & Perbaikan
        'DISBEKAL',     -- Dinas Perbekalan
        'DISLAIKMATAL'  -- Dinas Kelaikan Material
    );
EXCEPTION WHEN duplicate_object THEN null; END $$;

DO $$ BEGIN
    CREATE TYPE user_role_type AS ENUM (
        'SUPER_ADMIN',
        'COMMAND_OFFICER',
        'KRI_COMMANDER',
        'LOGISTICS_OFFICER',
        'MAINTENANCE_OFFICER',
        'PERSONNEL_OFFICER',
        'FINANCE_OFFICER',
        'OPERATOR'
    );
EXCEPTION WHEN duplicate_object THEN null; END $$;

DO $$ BEGIN
    CREATE TYPE audit_action_type AS ENUM (
        'CREATE',
        'UPDATE',
        'DELETE',
        'LOGIN',
        'LOGOUT',
        'APPROVAL',
        'EXPORT',
        'INITIALIZE'
    );
EXCEPTION WHEN duplicate_object THEN null; END $$;

-- -----------------------------------------------------------------------------
-- TABEL: org_units
-- -----------------------------------------------------------------------------
CREATE TABLE IF NOT EXISTS org_units (
    unit_id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    parent_unit_id UUID REFERENCES org_units(unit_id),
    unit_code VARCHAR(50) UNIQUE NOT NULL,
    unit_name VARCHAR(150) NOT NULL,
    unit_type org_unit_type NOT NULL,
    description TEXT,
    command_level INT DEFAULT 1,     -- 1: Mabesal, 2: Kotama/Armada, 3: Lantamal, 4: Lanal/Satuan, 5: KRI/Posal
    latitude NUMERIC(10, 7),
    longitude NUMERIC(10, 7),
    address TEXT,
    phone VARCHAR(50),
    is_active BOOLEAN DEFAULT TRUE,
    created_by UUID,
    updated_by UUID,
    deleted_by UUID,
    created_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMPTZ,
    deleted_at TIMESTAMPTZ
);

-- -----------------------------------------------------------------------------
-- TABEL: sys_users (1 User = 1 Role Type, Department Type, Rank Type)
-- -----------------------------------------------------------------------------
CREATE TABLE IF NOT EXISTS sys_users (
    user_id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    unit_id UUID NOT NULL REFERENCES org_units(unit_id),
    username VARCHAR(50) UNIQUE NOT NULL,
    password_hash VARCHAR(255) NOT NULL,
    full_name VARCHAR(150) NOT NULL,
    email VARCHAR(100) UNIQUE,
    phone VARCHAR(30),
    military_id VARCHAR(50) UNIQUE,  -- NRP (Nomor Registrasi Prajurit)
    rank_title military_rank_type,   -- ENUM: Pangkat Militer
    department department_type,      -- ENUM: Korps / Departemen Dinas
    role user_role_type NOT NULL DEFAULT 'OPERATOR', -- ENUM: 1 User Hanya 1 Role
    is_active BOOLEAN DEFAULT TRUE,
    last_login_at TIMESTAMPTZ,
    failed_login_attempts INT DEFAULT 0,
    created_by UUID,
    updated_by UUID,
    deleted_by UUID,
    created_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMPTZ,
    deleted_at TIMESTAMPTZ
);

-- -----------------------------------------------------------------------------
-- TABEL: sys_audit_logs
-- -----------------------------------------------------------------------------
CREATE TABLE IF NOT EXISTS sys_audit_logs (
    log_id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    user_id UUID REFERENCES sys_users(user_id) ON DELETE SET NULL,
    action audit_action_type NOT NULL,
    entity_table VARCHAR(100) NOT NULL,
    entity_id VARCHAR(100),
    old_values JSONB,
    new_values JSONB,
    ip_address VARCHAR(45),
    user_agent TEXT,
    created_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP
);

-- Indexes
CREATE INDEX IF NOT EXISTS idx_org_units_parent ON org_units(parent_unit_id) WHERE deleted_at IS NULL;
CREATE INDEX IF NOT EXISTS idx_org_units_type ON org_units(unit_type) WHERE deleted_at IS NULL;
CREATE INDEX IF NOT EXISTS idx_sys_users_unit ON sys_users(unit_id) WHERE deleted_at IS NULL;
CREATE INDEX IF NOT EXISTS idx_sys_users_role ON sys_users(role) WHERE deleted_at IS NULL;
CREATE INDEX IF NOT EXISTS idx_sys_users_military_id ON sys_users(military_id) WHERE deleted_at IS NULL;
CREATE INDEX IF NOT EXISTS idx_sys_audit_logs_user ON sys_audit_logs(user_id, created_at DESC);
CREATE INDEX IF NOT EXISTS idx_sys_audit_logs_table ON sys_audit_logs(entity_table, entity_id);

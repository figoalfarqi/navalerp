-- =============================================================================
-- MODUL 7: INFRASTRUKTUR PANGKALAN & LABUH (BASE & PORT INFRASTRUCTURE)
-- FILE: 07_base_infrastructure/tables.sql
-- =============================================================================

CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

-- -----------------------------------------------------------------------------
-- ENUM TYPES: MODUL 7
-- -----------------------------------------------------------------------------
DO $$ BEGIN
    CREATE TYPE facility_type_enum AS ENUM (
        'BERTH_JETTY',
        'GRAVING_DOCK',
        'FLOATING_DOCK',
        'FUEL_STORAGE',
        'AMMO_DEPOT',
        'WORKSHOP_FASHARKAN',
        'SHORE_POWER_STATION'
    );
EXCEPTION WHEN duplicate_object THEN null; END $$;

DO $$ BEGIN
    CREATE TYPE facility_status_type AS ENUM (
        'OPERATIONAL',
        'MAINTENANCE',
        'OCCUPIED',
        'UNDER_CONSTRUCTION'
    );
EXCEPTION WHEN duplicate_object THEN null; END $$;

DO $$ BEGIN
    CREATE TYPE berth_booking_purpose_type AS ENUM (
        'BERTHING_REST',
        'LOGISTIC_REPLENISHMENT',
        'MRO_REPAIR',
        'EMBARKATION',
        'VIP_CEREMONY'
    );
EXCEPTION WHEN duplicate_object THEN null; END $$;

DO $$ BEGIN
    CREATE TYPE berth_booking_status_type AS ENUM (
        'SCHEDULED',
        'BERTHED',
        'COMPLETED',
        'CANCELLED'
    );
EXCEPTION WHEN duplicate_object THEN null; END $$;

DO $$ BEGIN
    CREATE TYPE fuel_bunker_type AS ENUM (
        'HSD_MILSPEC',
        'B35_NAVAL',
        'AVTUR_JET_A1'
    );
EXCEPTION WHEN duplicate_object THEN null; END $$;

DO $$ BEGIN
    CREATE TYPE facility_maint_type AS ENUM (
        'DREDGING_ALUR',
        'CRANE_INSPECTION',
        'SHORE_POWER_CALIB',
        'DOCK_GATE_OVERHAUL',
        'FENDER_REPLACEMENT'
    );
EXCEPTION WHEN duplicate_object THEN null; END $$;

DO $$ BEGIN
    CREATE TYPE facility_maint_status_type AS ENUM (
        'SCHEDULED',
        'IN_PROGRESS',
        'COMPLETED'
    );
EXCEPTION WHEN duplicate_object THEN null; END $$;

-- -----------------------------------------------------------------------------
-- TABEL: infra_facilities
-- -----------------------------------------------------------------------------
CREATE TABLE IF NOT EXISTS infra_facilities (
    facility_id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    base_unit_id UUID NOT NULL REFERENCES org_units(unit_id),
    facility_code VARCHAR(50) UNIQUE NOT NULL,
    facility_name VARCHAR(150) NOT NULL,
    facility_type facility_type_enum NOT NULL,
    length_meters NUMERIC(8, 2),
    draft_depth_meters NUMERIC(5, 2),
    max_displacement_tonnage NUMERIC(10, 2),
    has_shore_power BOOLEAN DEFAULT TRUE,
    has_fresh_water BOOLEAN DEFAULT TRUE,
    has_fuel_bunker_line BOOLEAN DEFAULT TRUE,
    status facility_status_type DEFAULT 'OPERATIONAL',
    created_by UUID,
    updated_by UUID,
    deleted_by UUID,
    created_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMPTZ,
    deleted_at TIMESTAMPTZ
);

-- -----------------------------------------------------------------------------
-- TABEL: infra_berth_bookings
-- -----------------------------------------------------------------------------
CREATE TABLE IF NOT EXISTS infra_berth_bookings (
    booking_id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    facility_id UUID NOT NULL REFERENCES infra_facilities(facility_id),
    ship_id UUID NOT NULL REFERENCES mro_ships(ship_id),
    booking_purpose berth_booking_purpose_type NOT NULL,
    eta TIMESTAMPTZ NOT NULL,
    etd TIMESTAMPTZ NOT NULL,
    actual_berth_time TIMESTAMPTZ,
    actual_unberth_time TIMESTAMPTZ,
    shore_power_kwh_used NUMERIC(12, 2) DEFAULT 0,
    fresh_water_ton_used NUMERIC(10, 2) DEFAULT 0,
    status berth_booking_status_type DEFAULT 'SCHEDULED',
    approved_by_user_id UUID REFERENCES sys_users(user_id),
    remarks TEXT,
    created_by UUID,
    updated_by UUID,
    deleted_by UUID,
    created_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMPTZ,
    deleted_at TIMESTAMPTZ
);

-- -----------------------------------------------------------------------------
-- TABEL: infra_fuel_bunker_records
-- -----------------------------------------------------------------------------
CREATE TABLE IF NOT EXISTS infra_fuel_bunker_records (
    bunker_id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    ship_id UUID NOT NULL REFERENCES mro_ships(ship_id),
    facility_id UUID REFERENCES infra_facilities(facility_id),
    fuel_type fuel_bunker_type NOT NULL,
    quantity_liters NUMERIC(12, 2) NOT NULL,
    density_15c NUMERIC(6, 4) DEFAULT 0.8400,
    flow_rate_lph NUMERIC(10, 2),
    bunkering_start_time TIMESTAMPTZ NOT NULL,
    bunkering_end_time TIMESTAMPTZ NOT NULL,
    receipt_voucher_no VARCHAR(50) UNIQUE NOT NULL,
    authorised_by_user_id UUID NOT NULL REFERENCES sys_users(user_id),
    created_by UUID,
    updated_by UUID,
    deleted_by UUID,
    created_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMPTZ,
    deleted_at TIMESTAMPTZ
);

-- -----------------------------------------------------------------------------
-- TABEL: infra_facility_maintenances
-- -----------------------------------------------------------------------------
CREATE TABLE IF NOT EXISTS infra_facility_maintenances (
    maint_id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    facility_id UUID NOT NULL REFERENCES infra_facilities(facility_id) ON DELETE CASCADE,
    maintenance_type facility_maint_type NOT NULL,
    start_date DATE NOT NULL,
    end_date DATE,
    cost NUMERIC(18, 2) DEFAULT 0,
    performed_by VARCHAR(150) NOT NULL,
    status facility_maint_status_type DEFAULT 'COMPLETED',
    remarks TEXT,
    created_by UUID,
    created_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP
);

-- Indexes
CREATE INDEX IF NOT EXISTS idx_infra_fac_unit ON infra_facilities(base_unit_id) WHERE deleted_at IS NULL;
CREATE INDEX IF NOT EXISTS idx_infra_book_fac ON infra_berth_bookings(facility_id) WHERE deleted_at IS NULL;
CREATE INDEX IF NOT EXISTS idx_infra_book_ship ON infra_berth_bookings(ship_id) WHERE deleted_at IS NULL;
CREATE INDEX IF NOT EXISTS idx_infra_bunker_ship ON infra_fuel_bunker_records(ship_id) WHERE deleted_at IS NULL;

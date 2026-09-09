-- =============================================================================
-- MODUL 11: INFRASTRUKTUR BAWAH LAUT KRITIS (CRITICAL UNDERWATER INFRASTRUCTURE - CUI)
-- FILE: 11_cui_underwater_infrastructure/tables.sql
-- =============================================================================

CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

-- -----------------------------------------------------------------------------
-- ENUM TYPES: MODUL 11 (CUI)
-- -----------------------------------------------------------------------------
DO $$ BEGIN
    CREATE TYPE cui_asset_type_enum AS ENUM (
        'SUBMARINE_CABLE',
        'SUBSEA_PIPELINE',
        'LANDING_STATION',
        'OFFSHORE_ENERGY',
        'MONITORING_SYSTEM',
        'OTHER'
    );
EXCEPTION WHEN duplicate_object THEN null; END $$;

DO $$ BEGIN
    CREATE TYPE cui_status_enum AS ENUM (
        'ACTIVE_MONITORED',
        'INSPECTION_REQUIRED',
        'UNDER_MAINTENANCE',
        'ALERT_ANOMALY'
    );
EXCEPTION WHEN duplicate_object THEN null; END $$;

DO $$ BEGIN
    CREATE TYPE cui_priority_enum AS ENUM (
        'CRITICAL_TIER_1',
        'HIGH_TIER_2',
        'MEDIUM_TIER_3'
    );
EXCEPTION WHEN duplicate_object THEN null; END $$;

DO $$ BEGIN
    CREATE TYPE cui_alert_type_enum AS ENUM (
        'VESSEL_ANCHOR_DRAG_RISK',
        'SEISMIC_DISTURBANCE',
        'PRESSURE_DROP',
        'ACOUSTIC_ANOMALY',
        'UNAUTHORIZED_SUBMERSIBLE'
    );
EXCEPTION WHEN duplicate_object THEN null; END $$;

DO $$ BEGIN
    CREATE TYPE cui_alert_severity_enum AS ENUM (
        'CRITICAL',
        'HIGH',
        'MEDIUM',
        'LOW'
    );
EXCEPTION WHEN duplicate_object THEN null; END $$;

DO $$ BEGIN
    CREATE TYPE cui_alert_status_enum AS ENUM (
        'ACTIVE',
        'INVESTIGATING',
        'DISPATCHED',
        'RESOLVED',
        'FALSE_ALARM'
    );
EXCEPTION WHEN duplicate_object THEN null; END $$;

DO $$ BEGIN
    CREATE TYPE cui_inspection_method_enum AS ENUM (
        'ROV_SUBMERSIBLE',
        'DIVER_TEAM',
        'SIDE_SCAN_SONAR',
        'MAGNETOMETER'
    );
EXCEPTION WHEN duplicate_object THEN null; END $$;

DO $$ BEGIN
    CREATE TYPE cui_condition_enum AS ENUM (
        'EXCELLENT',
        'GOOD',
        'FAIR',
        'DAMAGED_CRITICAL'
    );
EXCEPTION WHEN duplicate_object THEN null; END $$;

-- -----------------------------------------------------------------------------
-- TABLE 1: cui_assets (Aset Bawah Laut Kritis Nasional)
-- -----------------------------------------------------------------------------
CREATE TABLE IF NOT EXISTS cui_assets (
    cui_asset_id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    asset_code VARCHAR(64) UNIQUE NOT NULL,
    asset_name VARCHAR(255) NOT NULL,
    asset_type cui_asset_type_enum NOT NULL DEFAULT 'SUBMARINE_CABLE',
    operator_name VARCHAR(150) NOT NULL,
    theater_id UUID REFERENCES ops_theaters(theater_id) ON DELETE SET NULL,
    depth_meters NUMERIC(10,2) DEFAULT 0,
    length_km NUMERIC(10,2) DEFAULT 0,
    latitude NUMERIC(10,6) NOT NULL DEFAULT 0,
    longitude NUMERIC(10,6) NOT NULL DEFAULT 0,
    start_coordinates VARCHAR(100),
    end_coordinates VARCHAR(100),
    status cui_status_enum NOT NULL DEFAULT 'ACTIVE_MONITORED',
    health_score INT NOT NULL DEFAULT 100,
    protection_priority cui_priority_enum NOT NULL DEFAULT 'CRITICAL_TIER_1',
    last_inspected_at TIMESTAMPTZ,
    next_inspection_due TIMESTAMPTZ,
    notes TEXT,
    created_by VARCHAR(100),
    updated_by VARCHAR(100),
    deleted_by VARCHAR(100),
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ,
    deleted_at TIMESTAMPTZ
);

-- -----------------------------------------------------------------------------
-- TABLE 2: cui_monitoring_logs (Log Sensor Pemantauan Bawah Laut)
-- -----------------------------------------------------------------------------
CREATE TABLE IF NOT EXISTS cui_monitoring_logs (
    log_id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    cui_asset_id UUID NOT NULL REFERENCES cui_assets(cui_asset_id) ON DELETE CASCADE,
    sensor_code VARCHAR(64) NOT NULL,
    sensor_type VARCHAR(64) NOT NULL DEFAULT 'ACOUSTIC_SONAR',
    log_time TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    metric_value NUMERIC(12,2) NOT NULL DEFAULT 0,
    metric_unit VARCHAR(32) NOT NULL DEFAULT 'dB',
    status VARCHAR(32) NOT NULL DEFAULT 'NORMAL',
    vessel_proximity_mmsi VARCHAR(32),
    anomaly_score NUMERIC(5,3) DEFAULT 0,
    description TEXT,
    created_by VARCHAR(100),
    updated_by VARCHAR(100),
    deleted_by VARCHAR(100),
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ,
    deleted_at TIMESTAMPTZ
);

-- -----------------------------------------------------------------------------
-- TABLE 3: cui_alerts (Peringatan Dini & Anomali Keamanan CUI)
-- -----------------------------------------------------------------------------
CREATE TABLE IF NOT EXISTS cui_alerts (
    alert_id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    alert_code VARCHAR(64) UNIQUE NOT NULL,
    cui_asset_id UUID NOT NULL REFERENCES cui_assets(cui_asset_id) ON DELETE CASCADE,
    alert_type cui_alert_type_enum NOT NULL DEFAULT 'VESSEL_ANCHOR_DRAG_RISK',
    severity cui_alert_severity_enum NOT NULL DEFAULT 'HIGH',
    detected_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    assigned_ship_id UUID REFERENCES mro_ships(ship_id) ON DELETE SET NULL,
    status cui_alert_status_enum NOT NULL DEFAULT 'ACTIVE',
    ai_confidence NUMERIC(5,3) NOT NULL DEFAULT 0.90,
    recommended_action TEXT,
    resolution_notes TEXT,
    resolved_at TIMESTAMPTZ,
    created_by VARCHAR(100),
    updated_by VARCHAR(100),
    deleted_by VARCHAR(100),
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ,
    deleted_at TIMESTAMPTZ
);

-- -----------------------------------------------------------------------------
-- TABLE 4: cui_inspections (Inspeksi Terjadwal / Darurat CUI)
-- -----------------------------------------------------------------------------
CREATE TABLE IF NOT EXISTS cui_inspections (
    inspection_id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    inspection_number VARCHAR(64) UNIQUE NOT NULL,
    cui_asset_id UUID NOT NULL REFERENCES cui_assets(cui_asset_id) ON DELETE CASCADE,
    ship_id UUID REFERENCES mro_ships(ship_id) ON DELETE SET NULL,
    inspection_date DATE NOT NULL DEFAULT CURRENT_DATE,
    inspector_officer_id UUID REFERENCES hcm_personnel(personnel_id) ON DELETE SET NULL,
    method cui_inspection_method_enum NOT NULL DEFAULT 'ROV_SUBMERSIBLE',
    condition_rating cui_condition_enum NOT NULL DEFAULT 'GOOD',
    findings TEXT,
    remedial_action_required BOOLEAN NOT NULL DEFAULT FALSE,
    next_inspection_date DATE,
    created_by VARCHAR(100),
    updated_by VARCHAR(100),
    deleted_by VARCHAR(100),
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ,
    deleted_at TIMESTAMPTZ
);


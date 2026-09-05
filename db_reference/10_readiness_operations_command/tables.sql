-- =============================================================================
-- MODUL 10: KESIAPAN OPERASI & KOMANDO TEMPUR (READINESS & OPERATIONS)
-- FILE: 10_readiness_operations_command/tables.sql
-- =============================================================================

CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

-- -----------------------------------------------------------------------------
-- ENUM TYPES: MODUL 10
-- -----------------------------------------------------------------------------
DO $$ BEGIN
    CREATE TYPE ops_defcon_level_type AS ENUM (
        'DEFCON_1',
        'DEFCON_2',
        'DEFCON_3',
        'DEFCON_4',
        'DEFCON_5'
    );
EXCEPTION WHEN duplicate_object THEN null; END $$;

DO $$ BEGIN
    CREATE TYPE ops_mission_type_enum AS ENUM (
        'COMBAT_PATROL',
        'JOINT_EXERCISE',
        'COUNTER_PIRACY',
        'SAR_HUMANITARIAN',
        'DIPLOMACY'
    );
EXCEPTION WHEN duplicate_object THEN null; END $$;

DO $$ BEGIN
    CREATE TYPE ops_mission_status_type AS ENUM (
        'PLANNING',
        'ACTIVE',
        'COMPLETED',
        'SUSPENDED'
    );
EXCEPTION WHEN duplicate_object THEN null; END $$;

DO $$ BEGIN
    CREATE TYPE ops_task_force_role_type AS ENUM (
        'FLAGSHIP',
        'ESCORT_SURFACE',
        'ASW_SCREEN',
        'AIR_DEFENSE',
        'LOGISTICS_SUPPORT'
    );
EXCEPTION WHEN duplicate_object THEN null; END $$;

DO $$ BEGIN
    CREATE TYPE ops_ship_assignment_status_type AS ENUM (
        'ACTIVE',
        'RELIEVED',
        'DAMAGED_RTB'
    );
EXCEPTION WHEN duplicate_object THEN null; END $$;

DO $$ BEGIN
    CREATE TYPE readiness_category_type AS ENUM (
        'C-1', -- Siap Tempur Penuh
        'C-2', -- Siap Tempur Terbatas
        'C-3', -- Rawat Jalan
        'C-4'  -- Docking / Non-Operasional
    );
EXCEPTION WHEN duplicate_object THEN null; END $$;

DO $$ BEGIN
    CREATE TYPE alert_severity_type AS ENUM (
        'CRITICAL',
        'HIGH',
        'MEDIUM',
        'LOW'
    );
EXCEPTION WHEN duplicate_object THEN null; END $$;

DO $$ BEGIN
    CREATE TYPE alert_type_enum AS ENUM (
        'CASREP_DEFECT',
        'CRITICAL_SPARE_DEFICIT',
        'CREW_SHORTAGE',
        'EXPIRED_KELAIKAN'
    );
EXCEPTION WHEN duplicate_object THEN null; END $$;

-- -----------------------------------------------------------------------------
-- TABEL: ops_theaters
-- -----------------------------------------------------------------------------
CREATE TABLE IF NOT EXISTS ops_theaters (
    theater_id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    theater_code VARCHAR(50) UNIQUE NOT NULL,
    theater_name VARCHAR(150) NOT NULL,
    responsible_command_unit_id UUID NOT NULL REFERENCES org_units(unit_id),
    threat_level ops_defcon_level_type DEFAULT 'DEFCON_3',
    description TEXT,
    created_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP
);

-- -----------------------------------------------------------------------------
-- TABEL: ops_missions
-- -----------------------------------------------------------------------------
CREATE TABLE IF NOT EXISTS ops_missions (
    mission_id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    theater_id UUID NOT NULL REFERENCES ops_theaters(theater_id),
    mission_code VARCHAR(50) UNIQUE NOT NULL,
    mission_name VARCHAR(200) NOT NULL,
    mission_type ops_mission_type_enum NOT NULL,
    start_date DATE NOT NULL,
    end_date DATE,
    commanding_officer_user_id UUID REFERENCES sys_users(user_id),
    mission_status ops_mission_status_type DEFAULT 'ACTIVE',
    created_by UUID,
    updated_by UUID,
    deleted_by UUID,
    created_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMPTZ,
    deleted_at TIMESTAMPTZ
);

-- -----------------------------------------------------------------------------
-- TABEL: ops_mission_ship_assignments
-- -----------------------------------------------------------------------------
CREATE TABLE IF NOT EXISTS ops_mission_ship_assignments (
    assignment_id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    mission_id UUID NOT NULL REFERENCES ops_missions(mission_id) ON DELETE CASCADE,
    ship_id UUID NOT NULL REFERENCES mro_ships(ship_id),
    tactical_callsign VARCHAR(50),
    role_in_task_force ops_task_force_role_type NOT NULL,
    joined_date DATE NOT NULL,
    released_date DATE,
    status ops_ship_assignment_status_type DEFAULT 'ACTIVE',
    created_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
    CONSTRAINT uq_mission_ship UNIQUE (mission_id, ship_id)
);

-- -----------------------------------------------------------------------------
-- TABEL: ops_daily_logs
-- -----------------------------------------------------------------------------
CREATE TABLE IF NOT EXISTS ops_daily_logs (
    log_id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    ship_id UUID NOT NULL REFERENCES mro_ships(ship_id),
    log_date DATE NOT NULL,
    latitude NUMERIC(10, 6) NOT NULL,
    longitude NUMERIC(10, 6) NOT NULL,
    heading_degrees INT CHECK (heading_degrees >= 0 AND heading_degrees < 360),
    speed_knots NUMERIC(4, 1) NOT NULL,
    sea_state INT CHECK (sea_state >= 0 AND sea_state <= 9),
    weather_condition VARCHAR(50),
    fuel_remaining_liters NUMERIC(12, 2) NOT NULL,
    fresh_water_remaining_tons NUMERIC(8, 2) NOT NULL,
    tactical_summary TEXT,
    logged_by_user_id UUID NOT NULL REFERENCES sys_users(user_id),
    created_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP
);

-- -----------------------------------------------------------------------------
-- TABEL: ops_ship_readiness_snapshots
-- -----------------------------------------------------------------------------
CREATE TABLE IF NOT EXISTS ops_ship_readiness_snapshots (
    snapshot_id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    ship_id UUID NOT NULL REFERENCES mro_ships(ship_id),
    snapshot_timestamp TIMESTAMPTZ NOT NULL,
    readiness_category readiness_category_type NOT NULL,
    mro_readiness_score NUMERIC(5, 2) NOT NULL,
    personnel_manning_score NUMERIC(5, 2) NOT NULL,
    logistics_supply_score NUMERIC(5, 2) NOT NULL,
    composite_readiness_index NUMERIC(5, 2) GENERATED ALWAYS AS (
        (mro_readiness_score * 0.40) + (personnel_manning_score * 0.30) + (logistics_supply_score * 0.30)
    ) STORED,
    remarks TEXT,
    created_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP
);

-- -----------------------------------------------------------------------------
-- TABEL: ops_readiness_alerts
-- -----------------------------------------------------------------------------
CREATE TABLE IF NOT EXISTS ops_readiness_alerts (
    alert_id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    ship_id UUID NOT NULL REFERENCES mro_ships(ship_id),
    equipment_id UUID REFERENCES mro_equipments(equipment_id),
    severity alert_severity_type NOT NULL,
    alert_type alert_type_enum NOT NULL,
    alert_message TEXT NOT NULL,
    is_acknowledged BOOLEAN DEFAULT FALSE,
    acknowledged_by_user_id UUID REFERENCES sys_users(user_id),
    acknowledged_at TIMESTAMPTZ,
    created_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP
);

-- Indexes
CREATE INDEX IF NOT EXISTS idx_ops_missions_code ON ops_missions(mission_code) WHERE deleted_at IS NULL;
CREATE INDEX IF NOT EXISTS idx_ops_ship_assign_ship ON ops_mission_ship_assignments(ship_id);
CREATE INDEX IF NOT EXISTS idx_ops_daily_logs_ship ON ops_daily_logs(ship_id, log_date);
CREATE INDEX IF NOT EXISTS idx_ops_readiness_ship ON ops_ship_readiness_snapshots(ship_id, snapshot_timestamp);
CREATE INDEX IF NOT EXISTS idx_ops_alerts_ship ON ops_readiness_alerts(ship_id) WHERE is_acknowledged = FALSE;

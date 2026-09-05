-- =============================================================================
-- MODUL 2: FLEET, ASSET & HIERARKI PERALATAN KAPAL (MRO & ASSET) (NAVALERP)
-- FILE: 02_fleet_asset_mro/tables.sql
-- =============================================================================

CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

-- -----------------------------------------------------------------------------
-- ENUM TYPES: MODUL 2
-- -----------------------------------------------------------------------------
DO $$ BEGIN
    CREATE TYPE ship_category_type AS ENUM (
        'FRIGATE',
        'CORVETTE',
        'SUBMARINE',
        'LPD',
        'PATROL',
        'AIRCRAFT',
        'FAST_ATTACK',
        'AUXILIARY'
    );
EXCEPTION WHEN duplicate_object THEN null; END $$;

DO $$ BEGIN
    CREATE TYPE ship_status_type AS ENUM (
        'ACTIVE',
        'DOCKED',
        'RETIRED',
        'DEPLOYED',
        'STANDBY'
    );
EXCEPTION WHEN duplicate_object THEN null; END $$;

DO $$ BEGIN
    CREATE TYPE ship_readiness_status_type AS ENUM (
        'FULLY_MISSION_CAPABLE',
        'PARTIALLY_MISSION_CAPABLE',
        'NON_MISSION_CAPABLE'
    );
EXCEPTION WHEN duplicate_object THEN null; END $$;

DO $$ BEGIN
    CREATE TYPE system_category_type AS ENUM (
        'PROPULSION',
        'ELECTRICAL',
        'SENSOR_RADAR',
        'WEAPON',
        'NAVIGATION',
        'AUXILIARY',
        'DAMAGE_CONTROL'
    );
EXCEPTION WHEN duplicate_object THEN null; END $$;

DO $$ BEGIN
    CREATE TYPE system_level_type AS ENUM (
        'SYSTEM',
        'SUBSYSTEM',
        'EQUIPMENT',
        'COMPONENT'
    );
EXCEPTION WHEN duplicate_object THEN null; END $$;

DO $$ BEGIN
    CREATE TYPE equipment_criticality_type AS ENUM (
        'CRITICAL_SAFETY',
        'MISSION_ESSENTIAL',
        'ROUTINE'
    );
EXCEPTION WHEN duplicate_object THEN null; END $$;

DO $$ BEGIN
    CREATE TYPE equipment_health_type AS ENUM (
        'OPERATIONAL',
        'DEGRADED',
        'CRITICAL',
        'NON_OPERATIONAL'
    );
EXCEPTION WHEN duplicate_object THEN null; END $$;

DO $$ BEGIN
    CREATE TYPE telemetry_status_type AS ENUM (
        'NORMAL',
        'WARNING',
        'ALARM'
    );
EXCEPTION WHEN duplicate_object THEN null; END $$;

DO $$ BEGIN
    CREATE TYPE failure_severity_type AS ENUM (
        'CAT1', -- Immediate Mission Stop
        'CAT2', -- Mission Degraded
        'CAT3', -- Minor
        'CAT4'  -- Routine
    );
EXCEPTION WHEN duplicate_object THEN null; END $$;

DO $$ BEGIN
    CREATE TYPE failure_status_type AS ENUM (
        'PENDING',
        'ASSESSED',
        'WORK_ORDER_CREATED',
        'CLOSED'
    );
EXCEPTION WHEN duplicate_object THEN null; END $$;

DO $$ BEGIN
    CREATE TYPE work_order_type_enum AS ENUM (
        'CORRECTIVE',
        'PREVENTIVE',
        'DOCKING',
        'DEPOT_LEVEL',
        'EMERGENCY'
    );
EXCEPTION WHEN duplicate_object THEN null; END $$;

DO $$ BEGIN
    CREATE TYPE work_order_priority_type AS ENUM (
        'EMERGENCY',
        'URGENT',
        'ROUTINE'
    );
EXCEPTION WHEN duplicate_object THEN null; END $$;

DO $$ BEGIN
    CREATE TYPE work_order_status_type AS ENUM (
        'DRAFT',
        'APPROVED',
        'IN_PROGRESS',
        'WAITING_PARTS',
        'COMPLETED',
        'INSPECTED',
        'CANCELLED'
    );
EXCEPTION WHEN duplicate_object THEN null; END $$;

DO $$ BEGIN
    CREATE TYPE docking_type_enum AS ENUM (
        'ANNUAL_DOCKING',
        'SPECIAL_DOCKING',
        'REPAIR_DOCKING',
        'MODERNIZATION'
    );
EXCEPTION WHEN duplicate_object THEN null; END $$;

-- -----------------------------------------------------------------------------
-- TABEL: mro_ship_classes
-- -----------------------------------------------------------------------------
CREATE TABLE IF NOT EXISTS mro_ship_classes (
    class_id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    class_code VARCHAR(50) UNIQUE NOT NULL,
    class_name VARCHAR(100) NOT NULL,
    category ship_category_type NOT NULL,
    specifications JSONB,
    builder VARCHAR(150),
    total_built INT DEFAULT 0,
    created_by UUID,
    updated_by UUID,
    deleted_by UUID,
    created_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMPTZ,
    deleted_at TIMESTAMPTZ
);

-- -----------------------------------------------------------------------------
-- TABEL: mro_ships
-- -----------------------------------------------------------------------------
CREATE TABLE IF NOT EXISTS mro_ships (
    ship_id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    class_id UUID NOT NULL REFERENCES mro_ship_classes(class_id),
    assigned_unit_id UUID NOT NULL REFERENCES org_units(unit_id),
    hull_number VARCHAR(20) UNIQUE NOT NULL,      -- misal: 331, 332, 365, 590
    ship_name VARCHAR(100) NOT NULL,             -- misal: KRI Raden Eddy Martadinata
    call_sign VARCHAR(30),
    commission_date DATE,
    home_port VARCHAR(100),
    length_m NUMERIC(6, 2),
    beam_m NUMERIC(6, 2),
    draft_m NUMERIC(6, 2),
    displacement_tons NUMERIC(10, 2),
    max_speed_knots NUMERIC(5, 2),
    cruise_range_nm NUMERIC(10, 2),
    crew_capacity INT,
    fuel_capacity_liters NUMERIC(14, 2),
    fresh_water_capacity_liters NUMERIC(14, 2),
    status ship_status_type DEFAULT 'ACTIVE',
    current_readiness_status ship_readiness_status_type DEFAULT 'FULLY_MISSION_CAPABLE',
    created_by UUID,
    updated_by UUID,
    deleted_by UUID,
    created_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMPTZ,
    deleted_at TIMESTAMPTZ
);

-- -----------------------------------------------------------------------------
-- TABEL: mro_systems
-- -----------------------------------------------------------------------------
CREATE TABLE IF NOT EXISTS mro_systems (
    system_id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    ship_id UUID NOT NULL REFERENCES mro_ships(ship_id) ON DELETE CASCADE,
    parent_system_id UUID REFERENCES mro_systems(system_id),
    system_code VARCHAR(50) NOT NULL,
    system_name VARCHAR(100) NOT NULL,
    system_category system_category_type NOT NULL,
    system_level system_level_type NOT NULL,
    description TEXT,
    created_by UUID,
    updated_by UUID,
    deleted_by UUID,
    created_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMPTZ,
    deleted_at TIMESTAMPTZ
);

-- -----------------------------------------------------------------------------
-- TABEL: mro_equipments
-- -----------------------------------------------------------------------------
CREATE TABLE IF NOT EXISTS mro_equipments (
    equipment_id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    system_id UUID NOT NULL REFERENCES mro_systems(system_id) ON DELETE CASCADE,
    serial_number VARCHAR(100) NOT NULL,
    equipment_tag VARCHAR(50),
    equipment_name VARCHAR(150) NOT NULL,
    manufacturer VARCHAR(100),
    model_number VARCHAR(100),
    country_of_origin VARCHAR(100),
    installation_date DATE,
    total_operating_hours NUMERIC(10, 2) DEFAULT 0,
    design_life_hours NUMERIC(10, 2),
    criticality_level equipment_criticality_type DEFAULT 'MISSION_ESSENTIAL',
    health_status equipment_health_type DEFAULT 'OPERATIONAL',
    created_by UUID,
    updated_by UUID,
    deleted_by UUID,
    created_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMPTZ,
    deleted_at TIMESTAMPTZ
);

-- -----------------------------------------------------------------------------
-- TABEL: mro_equipment_parameters
-- -----------------------------------------------------------------------------
CREATE TABLE IF NOT EXISTS mro_equipment_parameters (
    param_id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    equipment_id UUID NOT NULL REFERENCES mro_equipments(equipment_id) ON DELETE CASCADE,
    recorded_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
    rpm NUMERIC(8, 2),
    temperature_celsius NUMERIC(6, 2),
    pressure_bar NUMERIC(6, 2),
    vibration_level NUMERIC(6, 3),
    oil_pressure_bar NUMERIC(6, 2),
    running_hours_snapshot NUMERIC(10, 2),
    status_flag telemetry_status_type DEFAULT 'NORMAL',
    recorded_by_user_id UUID REFERENCES sys_users(user_id)
);

-- -----------------------------------------------------------------------------
-- TABEL: mro_pm_schedules
-- -----------------------------------------------------------------------------
CREATE TABLE IF NOT EXISTS mro_pm_schedules (
    pm_id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    equipment_id UUID NOT NULL REFERENCES mro_equipments(equipment_id) ON DELETE CASCADE,
    pm_code VARCHAR(50) UNIQUE NOT NULL,
    pm_title VARCHAR(150) NOT NULL,
    interval_hours INT,              -- misal tiap 250, 500, 1000, 5000 jam operasi
    interval_days INT,               -- atau interval kalender (misal 180 hari)
    last_performed_at TIMESTAMPTZ,
    next_due_at TIMESTAMPTZ,
    task_instructions TEXT NOT NULL,
    estimated_duration_hours NUMERIC(6, 2) DEFAULT 4,
    is_active BOOLEAN DEFAULT TRUE,
    created_by UUID,
    updated_by UUID,
    deleted_by UUID,
    created_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMPTZ,
    deleted_at TIMESTAMPTZ
);

-- -----------------------------------------------------------------------------
-- TABEL: mro_failure_reports
-- -----------------------------------------------------------------------------
CREATE TABLE IF NOT EXISTS mro_failure_reports (
    report_id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    equipment_id UUID NOT NULL REFERENCES mro_equipments(equipment_id),
    reported_by_user_id UUID NOT NULL REFERENCES sys_users(user_id),
    report_number VARCHAR(50) UNIQUE NOT NULL,
    incident_date TIMESTAMPTZ NOT NULL,
    severity failure_severity_type NOT NULL,
    failure_mode VARCHAR(100),
    description TEXT NOT NULL,
    operational_impact TEXT,
    immediate_action_taken TEXT,
    status failure_status_type DEFAULT 'PENDING',
    created_by UUID,
    updated_by UUID,
    deleted_by UUID,
    created_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMPTZ,
    deleted_at TIMESTAMPTZ
);

-- -----------------------------------------------------------------------------
-- TABEL: mro_work_orders
-- -----------------------------------------------------------------------------
CREATE TABLE IF NOT EXISTS mro_work_orders (
    work_order_id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    failure_report_id UUID REFERENCES mro_failure_reports(report_id),
    pm_schedule_id UUID REFERENCES mro_pm_schedules(pm_id),
    equipment_id UUID NOT NULL REFERENCES mro_equipments(equipment_id),
    work_order_number VARCHAR(50) UNIQUE NOT NULL,
    work_order_type work_order_type_enum NOT NULL,
    priority work_order_priority_type NOT NULL,
    scheduled_start_date DATE,
    scheduled_end_date DATE,
    actual_start_date DATE,
    actual_end_date DATE,
    lead_engineer_user_id UUID REFERENCES sys_users(user_id),
    assigned_facility VARCHAR(150),       -- misal: Fasharkan Surabaya, Bengkel Mesin Ujung, Galangan PT PAL
    status work_order_status_type DEFAULT 'DRAFT',
    total_labor_hours NUMERIC(8, 2) DEFAULT 0,
    estimated_cost NUMERIC(18, 2) DEFAULT 0,
    actual_cost NUMERIC(18, 2) DEFAULT 0,
    completion_notes TEXT,
    created_by UUID,
    updated_by UUID,
    deleted_by UUID,
    created_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMPTZ,
    deleted_at TIMESTAMPTZ
);

-- -----------------------------------------------------------------------------
-- TABEL: mro_work_order_tasks
-- -----------------------------------------------------------------------------
CREATE TABLE IF NOT EXISTS mro_work_order_tasks (
    task_id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    work_order_id UUID NOT NULL REFERENCES mro_work_orders(work_order_id) ON DELETE CASCADE,
    step_number INT NOT NULL,
    task_description TEXT NOT NULL,
    estimated_minutes INT DEFAULT 60,
    actual_minutes INT,
    is_completed BOOLEAN DEFAULT FALSE,
    completed_by_user_id UUID REFERENCES sys_users(user_id),
    notes TEXT,
    created_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP
);

-- -----------------------------------------------------------------------------
-- TABEL: mro_work_order_items
-- -----------------------------------------------------------------------------
CREATE TABLE IF NOT EXISTS mro_work_order_items (
    wo_item_id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    work_order_id UUID NOT NULL REFERENCES mro_work_orders(work_order_id) ON DELETE CASCADE,
    material_id UUID NOT NULL, -- Merujuk ke inv_materials
    quantity_required NUMERIC(10, 2) NOT NULL,
    quantity_issued NUMERIC(10, 2) DEFAULT 0,
    unit_cost NUMERIC(18, 2) DEFAULT 0,
    total_cost NUMERIC(18, 2) DEFAULT 0,
    is_critical_spare BOOLEAN DEFAULT FALSE,
    created_by UUID,
    updated_by UUID,
    deleted_by UUID,
    created_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMPTZ,
    deleted_at TIMESTAMPTZ
);

-- -----------------------------------------------------------------------------
-- TABEL: mro_docking_records
-- -----------------------------------------------------------------------------
CREATE TABLE IF NOT EXISTS mro_docking_records (
    docking_id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    ship_id UUID NOT NULL REFERENCES mro_ships(ship_id) ON DELETE CASCADE,
    shipyard_name VARCHAR(150) NOT NULL,
    docking_type docking_type_enum NOT NULL,
    entry_date DATE NOT NULL,
    scheduled_exit_date DATE NOT NULL,
    actual_exit_date DATE,
    sea_trial_passed BOOLEAN DEFAULT FALSE,
    classification_surveyor VARCHAR(100),
    certificate_number VARCHAR(100),
    total_docking_cost NUMERIC(18, 2) DEFAULT 0,
    docking_summary TEXT,
    created_by UUID,
    updated_by UUID,
    deleted_by UUID,
    created_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMPTZ,
    deleted_at TIMESTAMPTZ
);

-- Indexes
CREATE INDEX IF NOT EXISTS idx_mro_ships_class ON mro_ships(class_id) WHERE deleted_at IS NULL;
CREATE INDEX IF NOT EXISTS idx_mro_ships_unit ON mro_ships(assigned_unit_id) WHERE deleted_at IS NULL;
CREATE INDEX IF NOT EXISTS idx_mro_systems_ship ON mro_systems(ship_id) WHERE deleted_at IS NULL;
CREATE INDEX IF NOT EXISTS idx_mro_equipments_system ON mro_equipments(system_id) WHERE deleted_at IS NULL;
CREATE INDEX IF NOT EXISTS idx_mro_equipments_health ON mro_equipments(health_status) WHERE deleted_at IS NULL;
CREATE INDEX IF NOT EXISTS idx_mro_work_orders_equipment ON mro_work_orders(equipment_id) WHERE deleted_at IS NULL;
CREATE INDEX IF NOT EXISTS idx_mro_work_orders_status ON mro_work_orders(status) WHERE deleted_at IS NULL;
CREATE INDEX IF NOT EXISTS idx_mro_docking_ship ON mro_docking_records(ship_id) WHERE deleted_at IS NULL;

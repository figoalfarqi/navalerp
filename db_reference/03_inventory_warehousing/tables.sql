-- =============================================================================
-- MODUL 3: LOGISTIK, INVENTARISASI & WAREHOUSING (NAVALERP)
-- FILE: 03_inventory_warehousing/tables.sql
-- =============================================================================

CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

-- -----------------------------------------------------------------------------
-- ENUM TYPES: MODUL 3
-- -----------------------------------------------------------------------------
DO $$ BEGIN
    CREATE TYPE warehouse_type_enum AS ENUM (
        'SHORE_DEPOT',
        'SHIPBOARD_STORE',
        'AMMUNITION_BUNKER',
        'FUEL_DEPOT',
        'COLD_STORAGE'
    );
EXCEPTION WHEN duplicate_object THEN null; END $$;

DO $$ BEGIN
    CREATE TYPE material_category_type AS ENUM (
        'TECHNICAL_SPARES',
        'CONSUMABLES',
        'FUEL_ENERGY',
        'AMMUNITION',
        'GENERAL_SUPPLIES',
        'STRATEGIC_RESERVE'
    );
EXCEPTION WHEN duplicate_object THEN null; END $$;

DO $$ BEGIN
    CREATE TYPE uom_type AS ENUM (
        'UNIT',
        'SET',
        'LITER',
        'KG',
        'TON',
        'METER',
        'BOX',
        'DRUM',
        'CRATE'
    );
EXCEPTION WHEN duplicate_object THEN null; END $$;

DO $$ BEGIN
    CREATE TYPE item_condition_type AS ENUM (
        'SERVICEABLE',
        'UNSERVICEABLE',
        'CONDEMNED',
        'IN_REPAIR'
    );
EXCEPTION WHEN duplicate_object THEN null; END $$;

DO $$ BEGIN
    CREATE TYPE stock_movement_type AS ENUM (
        'FLEET_RESUPPLY',
        'BASE_TRANSFER',
        'DEPOT_DISPATCH',
        'EMERGENCY_AIRDROP'
    );
EXCEPTION WHEN duplicate_object THEN null; END $$;

DO $$ BEGIN
    CREATE TYPE transfer_status_type AS ENUM (
        'PLANNED',
        'IN_TRANSIT',
        'RECEIVED',
        'REJECTED'
    );
EXCEPTION WHEN duplicate_object THEN null; END $$;

DO $$ BEGIN
    CREATE TYPE adjustment_reason_type AS ENUM (
        'STOCK_OPNAME_VARIANCE',
        'DAMAGE',
        'EXPIRY',
        'SALVAGE'
    );
EXCEPTION WHEN duplicate_object THEN null; END $$;

DO $$ BEGIN
    CREATE TYPE adjustment_status_type AS ENUM (
        'DRAFT',
        'APPROVED',
        'CANCELLED'
    );
EXCEPTION WHEN duplicate_object THEN null; END $$;

-- -----------------------------------------------------------------------------
-- TABEL: inv_warehouses
-- -----------------------------------------------------------------------------
CREATE TABLE IF NOT EXISTS inv_warehouses (
    warehouse_id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    unit_id UUID NOT NULL REFERENCES org_units(unit_id),
    warehouse_code VARCHAR(50) UNIQUE NOT NULL,
    warehouse_name VARCHAR(150) NOT NULL,
    warehouse_type warehouse_type_enum NOT NULL,
    capacity_m3 NUMERIC(12, 2),
    manager_user_id UUID REFERENCES sys_users(user_id),
    location_address TEXT,
    is_active BOOLEAN DEFAULT TRUE,
    created_by UUID,
    updated_by UUID,
    deleted_by UUID,
    created_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMPTZ,
    deleted_at TIMESTAMPTZ
);

-- -----------------------------------------------------------------------------
-- TABEL: inv_storage_locations
-- -----------------------------------------------------------------------------
CREATE TABLE IF NOT EXISTS inv_storage_locations (
    location_id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    warehouse_id UUID NOT NULL REFERENCES inv_warehouses(warehouse_id) ON DELETE CASCADE,
    zone_name VARCHAR(50) NOT NULL,     -- misal Zona A, Zona B-Amunisi, Tangki Bunker 1
    aisle VARCHAR(20),
    rack VARCHAR(20),
    shelf VARCHAR(20),
    bin_code VARCHAR(50) NOT NULL,
    capacity_kg NUMERIC(10, 2),
    is_hazardous_zone BOOLEAN DEFAULT FALSE,
    created_by UUID,
    updated_by UUID,
    deleted_by UUID,
    created_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMPTZ,
    deleted_at TIMESTAMPTZ,
    UNIQUE(warehouse_id, bin_code)
);

-- -----------------------------------------------------------------------------
-- TABEL: inv_materials
-- -----------------------------------------------------------------------------
CREATE TABLE IF NOT EXISTS inv_materials (
    material_id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    material_code VARCHAR(50) UNIQUE NOT NULL,
    nsn VARCHAR(50) UNIQUE,              -- NATO / Military Stock Number (misal: 2910-12-384-9021)
    part_number VARCHAR(100),
    oem_name VARCHAR(100),
    material_name VARCHAR(150) NOT NULL,
    category material_category_type NOT NULL,
    uom uom_type NOT NULL,
    weight_kg NUMERIC(8, 2),
    min_stock_level NUMERIC(12, 2) DEFAULT 0,
    max_stock_level NUMERIC(12, 2) DEFAULT 0,
    reorder_point NUMERIC(12, 2) DEFAULT 0,
    safety_stock NUMERIC(12, 2) DEFAULT 0,
    shelf_life_days INT,
    is_controlled_item BOOLEAN DEFAULT FALSE, -- Alutsista / Munisi / Bersandi
    unit_price_idR NUMERIC(18, 2) DEFAULT 0,
    created_by UUID,
    updated_by UUID,
    deleted_by UUID,
    created_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMPTZ,
    deleted_at TIMESTAMPTZ
);

-- -----------------------------------------------------------------------------
-- TABEL: inv_material_equipment_links
-- -----------------------------------------------------------------------------
CREATE TABLE IF NOT EXISTS inv_material_equipment_links (
    link_id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    material_id UUID NOT NULL REFERENCES inv_materials(material_id) ON DELETE CASCADE,
    system_id UUID REFERENCES mro_systems(system_id) ON DELETE SET NULL,
    equipment_id UUID REFERENCES mro_equipments(equipment_id) ON DELETE SET NULL,
    is_mandatory_spare BOOLEAN DEFAULT FALSE,
    interchangeability_code VARCHAR(50),
    created_by UUID,
    updated_by UUID,
    deleted_by UUID,
    created_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMPTZ,
    deleted_at TIMESTAMPTZ
);

-- -----------------------------------------------------------------------------
-- TABEL: inv_stock_balances
-- -----------------------------------------------------------------------------
CREATE TABLE IF NOT EXISTS inv_stock_balances (
    balance_id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    warehouse_id UUID NOT NULL REFERENCES inv_warehouses(warehouse_id),
    location_id UUID REFERENCES inv_storage_locations(location_id),
    material_id UUID NOT NULL REFERENCES inv_materials(material_id),
    quantity_on_hand NUMERIC(14, 2) NOT NULL DEFAULT 0,
    quantity_reserved NUMERIC(14, 2) NOT NULL DEFAULT 0,
    quantity_in_transit NUMERIC(14, 2) NOT NULL DEFAULT 0,
    quantity_available NUMERIC(14, 2) GENERATED ALWAYS AS (quantity_on_hand - quantity_reserved) STORED,
    last_count_date DATE,
    created_by UUID,
    updated_by UUID,
    deleted_by UUID,
    created_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMPTZ,
    deleted_at TIMESTAMPTZ,
    UNIQUE(warehouse_id, material_id)
);

-- -----------------------------------------------------------------------------
-- TABEL: inv_item_instances
-- -----------------------------------------------------------------------------
CREATE TABLE IF NOT EXISTS inv_item_instances (
    instance_id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    warehouse_id UUID NOT NULL REFERENCES inv_warehouses(warehouse_id),
    location_id UUID REFERENCES inv_storage_locations(location_id),
    material_id UUID NOT NULL REFERENCES inv_materials(material_id),
    batch_number VARCHAR(100),
    serial_number VARCHAR(100),
    lot_number VARCHAR(100),
    expiry_date DATE,
    manufactured_date DATE,
    condition item_condition_type DEFAULT 'SERVICEABLE',
    inspection_due_date DATE,
    quantity NUMERIC(12, 2) NOT NULL DEFAULT 1,
    created_by UUID,
    updated_by UUID,
    deleted_by UUID,
    created_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMPTZ,
    deleted_at TIMESTAMPTZ
);

-- -----------------------------------------------------------------------------
-- TABEL: inv_stock_transfers
-- -----------------------------------------------------------------------------
CREATE TABLE IF NOT EXISTS inv_stock_transfers (
    transfer_id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    transfer_number VARCHAR(50) UNIQUE NOT NULL,
    from_warehouse_id UUID NOT NULL REFERENCES inv_warehouses(warehouse_id),
    to_warehouse_id UUID NOT NULL REFERENCES inv_warehouses(warehouse_id),
    movement_type stock_movement_type NOT NULL,
    scheduled_departure TIMESTAMPTZ,
    actual_departure TIMESTAMPTZ,
    scheduled_arrival TIMESTAMPTZ,
    actual_arrival TIMESTAMPTZ,
    transporter_unit VARCHAR(100),
    status transfer_status_type DEFAULT 'PLANNED',
    created_by UUID,
    updated_by UUID,
    deleted_by UUID,
    created_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMPTZ,
    deleted_at TIMESTAMPTZ
);

-- -----------------------------------------------------------------------------
-- TABEL: inv_stock_transfer_items
-- -----------------------------------------------------------------------------
CREATE TABLE IF NOT EXISTS inv_stock_transfer_items (
    transfer_item_id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    transfer_id UUID NOT NULL REFERENCES inv_stock_transfers(transfer_id) ON DELETE CASCADE,
    material_id UUID NOT NULL REFERENCES inv_materials(material_id),
    quantity_shipped NUMERIC(12, 2) NOT NULL,
    quantity_received NUMERIC(12, 2) DEFAULT 0,
    condition_on_receipt item_condition_type DEFAULT 'SERVICEABLE',
    created_by UUID,
    updated_by UUID,
    deleted_by UUID,
    created_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMPTZ,
    deleted_at TIMESTAMPTZ
);

-- -----------------------------------------------------------------------------
-- TABEL: inv_stock_adjustments
-- -----------------------------------------------------------------------------
CREATE TABLE IF NOT EXISTS inv_stock_adjustments (
    adjustment_id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    warehouse_id UUID NOT NULL REFERENCES inv_warehouses(warehouse_id),
    adjustment_number VARCHAR(50) UNIQUE NOT NULL,
    adjustment_date DATE NOT NULL,
    conducted_by_user_id UUID REFERENCES sys_users(user_id),
    reason adjustment_reason_type NOT NULL,
    status adjustment_status_type DEFAULT 'DRAFT',
    remarks TEXT,
    created_by UUID,
    updated_by UUID,
    deleted_by UUID,
    created_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMPTZ,
    deleted_at TIMESTAMPTZ
);

-- -----------------------------------------------------------------------------
-- TABEL: inv_stock_adjustment_items
-- -----------------------------------------------------------------------------
CREATE TABLE IF NOT EXISTS inv_stock_adjustment_items (
    adj_item_id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    adjustment_id UUID NOT NULL REFERENCES inv_stock_adjustments(adjustment_id) ON DELETE CASCADE,
    material_id UUID NOT NULL REFERENCES inv_materials(material_id),
    book_quantity NUMERIC(12, 2) NOT NULL,
    physical_quantity NUMERIC(12, 2) NOT NULL,
    difference_quantity NUMERIC(12, 2) GENERATED ALWAYS AS (physical_quantity - book_quantity) STORED,
    unit_cost NUMERIC(18, 2) DEFAULT 0,
    total_adjustment_value NUMERIC(18, 2) DEFAULT 0,
    created_by UUID,
    updated_by UUID,
    deleted_by UUID,
    created_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMPTZ,
    deleted_at TIMESTAMPTZ
);

-- Indexes
CREATE INDEX IF NOT EXISTS idx_inv_warehouses_unit ON inv_warehouses(unit_id) WHERE deleted_at IS NULL;
CREATE INDEX IF NOT EXISTS idx_inv_locations_warehouse ON inv_storage_locations(warehouse_id) WHERE deleted_at IS NULL;
CREATE INDEX IF NOT EXISTS idx_inv_materials_nsn ON inv_materials(nsn) WHERE deleted_at IS NULL;
CREATE INDEX IF NOT EXISTS idx_inv_materials_category ON inv_materials(category) WHERE deleted_at IS NULL;
CREATE INDEX IF NOT EXISTS idx_inv_stock_balances_wh_mat ON inv_stock_balances(warehouse_id, material_id) WHERE deleted_at IS NULL;
CREATE INDEX IF NOT EXISTS idx_inv_stock_transfers_status ON inv_stock_transfers(status) WHERE deleted_at IS NULL;
CREATE INDEX IF NOT EXISTS idx_inv_item_instances_sn ON inv_item_instances(serial_number) WHERE deleted_at IS NULL;

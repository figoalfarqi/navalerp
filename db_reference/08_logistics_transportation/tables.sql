-- =============================================================================
-- MODUL 8: LOGISTIK & TRANSPORTASI MILITER (LOGISTICS & TRANSPORTATION)
-- FILE: 08_logistics_transportation/tables.sql
-- =============================================================================

CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

-- -----------------------------------------------------------------------------
-- ENUM TYPES: MODUL 8
-- -----------------------------------------------------------------------------
DO $$ BEGIN
    CREATE TYPE transport_unit_type AS ENUM (
        'NAVAL_AUXILIARY_VESSEL',
        'LAND_TRUCK',
        'AIR_CARGO',
        'CARGO_BARGE'
    );
EXCEPTION WHEN duplicate_object THEN null; END $$;

DO $$ BEGIN
    CREATE TYPE transport_status_type AS ENUM (
        'AVAILABLE',
        'ON_MISSION',
        'MAINTENANCE',
        'DECOMMISSIONED'
    );
EXCEPTION WHEN duplicate_object THEN null; END $$;

DO $$ BEGIN
    CREATE TYPE sea_route_risk_level_type AS ENUM (
        'NORMAL',
        'HIGH_SEA',
        'COMBAT_RISK'
    );
EXCEPTION WHEN duplicate_object THEN null; END $$;

DO $$ BEGIN
    CREATE TYPE escort_security_level_type AS ENUM (
        'UNESCORTED',
        'STANDARD_CONVOY',
        'WARSHIP_ESCORT'
    );
EXCEPTION WHEN duplicate_object THEN null; END $$;

DO $$ BEGIN
    CREATE TYPE shipment_status_type AS ENUM (
        'PLANNED',
        'LOADING',
        'IN_TRANSIT',
        'DELIVERED',
        'CANCELLED'
    );
EXCEPTION WHEN duplicate_object THEN null; END $$;

DO $$ BEGIN
    CREATE TYPE shipment_packaging_type AS ENUM (
        'CRATE',
        'PALLET',
        'DRUM',
        'AMMO_BOX',
        'ISO_CONTAINER'
    );
EXCEPTION WHEN duplicate_object THEN null; END $$;

-- -----------------------------------------------------------------------------
-- TABEL: log_transport_units
-- -----------------------------------------------------------------------------
CREATE TABLE IF NOT EXISTS log_transport_units (
    transport_unit_id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    unit_code VARCHAR(50) UNIQUE NOT NULL,
    unit_name VARCHAR(150) NOT NULL,
    transport_type transport_unit_type NOT NULL,
    cargo_capacity_tons NUMERIC(10, 2) NOT NULL,
    fuel_capacity_liters NUMERIC(12, 2),
    operating_unit_id UUID NOT NULL REFERENCES org_units(unit_id),
    status transport_status_type DEFAULT 'AVAILABLE',
    created_by UUID,
    updated_by UUID,
    deleted_by UUID,
    created_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMPTZ,
    deleted_at TIMESTAMPTZ
);

-- -----------------------------------------------------------------------------
-- TABEL: log_routes
-- -----------------------------------------------------------------------------
CREATE TABLE IF NOT EXISTS log_routes (
    route_id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    route_code VARCHAR(50) UNIQUE NOT NULL,
    route_name VARCHAR(150) NOT NULL,
    origin_facility_id UUID NOT NULL REFERENCES infra_facilities(facility_id),
    destination_facility_id UUID NOT NULL REFERENCES infra_facilities(facility_id),
    distance_nautical_miles NUMERIC(8, 2) NOT NULL,
    estimated_transit_hours NUMERIC(6, 2) NOT NULL,
    risk_level sea_route_risk_level_type DEFAULT 'NORMAL',
    created_by UUID,
    updated_by UUID,
    deleted_by UUID,
    created_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMPTZ,
    deleted_at TIMESTAMPTZ
);

-- -----------------------------------------------------------------------------
-- TABEL: log_shipments
-- -----------------------------------------------------------------------------
CREATE TABLE IF NOT EXISTS log_shipments (
    shipment_id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    manifest_number VARCHAR(50) UNIQUE NOT NULL,
    route_id UUID NOT NULL REFERENCES log_routes(route_id),
    transport_unit_id UUID NOT NULL REFERENCES log_transport_units(transport_unit_id),
    origin_warehouse_id UUID NOT NULL REFERENCES inv_warehouses(warehouse_id),
    destination_warehouse_id UUID NOT NULL REFERENCES inv_warehouses(warehouse_id),
    departure_date TIMESTAMPTZ NOT NULL,
    arrival_date TIMESTAMPTZ,
    escort_security_level escort_security_level_type DEFAULT 'STANDARD_CONVOY',
    status shipment_status_type DEFAULT 'IN_TRANSIT',
    authorized_by_user_id UUID REFERENCES sys_users(user_id),
    remarks TEXT,
    created_by UUID,
    updated_by UUID,
    deleted_by UUID,
    created_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMPTZ,
    deleted_at TIMESTAMPTZ
);

-- -----------------------------------------------------------------------------
-- TABEL: log_shipment_items
-- -----------------------------------------------------------------------------
CREATE TABLE IF NOT EXISTS log_shipment_items (
    shipment_item_id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    shipment_id UUID NOT NULL REFERENCES log_shipments(shipment_id) ON DELETE CASCADE,
    material_id UUID NOT NULL REFERENCES inv_materials(material_id),
    quantity_dispatched NUMERIC(12, 2) NOT NULL,
    quantity_received NUMERIC(12, 2) DEFAULT 0,
    packaging_type shipment_packaging_type NOT NULL,
    weight_kg NUMERIC(10, 2),
    notes TEXT,
    created_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP
);

-- Indexes
CREATE INDEX IF NOT EXISTS idx_log_trans_status ON log_transport_units(status) WHERE deleted_at IS NULL;
CREATE INDEX IF NOT EXISTS idx_log_routes_code ON log_routes(route_code) WHERE deleted_at IS NULL;
CREATE INDEX IF NOT EXISTS idx_log_shipments_manifest ON log_shipments(manifest_number) WHERE deleted_at IS NULL;
CREATE INDEX IF NOT EXISTS idx_log_shipment_items_mat ON log_shipment_items(material_id);

-- =============================================================================
-- MODUL 4: PENGADAAN PERTAHANAN (DEFENCE PROCUREMENT) (NAVALERP)
-- FILE: 04_defence_procurement/tables.sql
-- =============================================================================

CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

-- -----------------------------------------------------------------------------
-- ENUM TYPES: MODUL 4
-- -----------------------------------------------------------------------------
DO $$ BEGIN
    CREATE TYPE security_clearance_type AS ENUM (
        'NATO_SECRET',
        'RESTRICTED',
        'UNCLASSIFIED',
        'RAHASIA_NEGARA'
    );
EXCEPTION WHEN duplicate_object THEN null; END $$;

DO $$ BEGIN
    CREATE TYPE requisition_priority_type AS ENUM (
        'EMERGENCY',
        'HIGH',
        'REGULAR'
    );
EXCEPTION WHEN duplicate_object THEN null; END $$;

DO $$ BEGIN
    CREATE TYPE requisition_approval_status_type AS ENUM (
        'PENDING',
        'APPROVED',
        'REJECTED'
    );
EXCEPTION WHEN duplicate_object THEN null; END $$;

DO $$ BEGIN
    CREATE TYPE procurement_category_type AS ENUM (
        'ALUTSISTA',
        'SPARE_PARTS',
        'MRO_SERVICE',
        'FACILITY'
    );
EXCEPTION WHEN duplicate_object THEN null; END $$;

DO $$ BEGIN
    CREATE TYPE procurement_method_type AS ENUM (
        'DIRECT_APPOINTMENT',
        'LIMITED_TENDER',
        'OPEN_TENDER',
        'G2G'
    );
EXCEPTION WHEN duplicate_object THEN null; END $$;

DO $$ BEGIN
    CREATE TYPE tender_status_type AS ENUM (
        'DRAFT',
        'PUBLISHED',
        'EVALUATING',
        'AWARDED',
        'CANCELLED'
    );
EXCEPTION WHEN duplicate_object THEN null; END $$;

DO $$ BEGIN
    CREATE TYPE contract_status_type AS ENUM (
        'DRAFT',
        'SIGNED',
        'ACTIVE',
        'COMPLETED',
        'TERMINATED'
    );
EXCEPTION WHEN duplicate_object THEN null; END $$;

DO $$ BEGIN
    CREATE TYPE po_status_type AS ENUM (
        'ISSUED',
        'PARTIAL_RECEIVED',
        'COMPLETED',
        'CANCELLED'
    );
EXCEPTION WHEN duplicate_object THEN null; END $$;

-- -----------------------------------------------------------------------------
-- TABEL: proc_vendors
-- -----------------------------------------------------------------------------
CREATE TABLE IF NOT EXISTS proc_vendors (
    vendor_id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    vendor_code VARCHAR(50) UNIQUE NOT NULL,
    vendor_name VARCHAR(150) NOT NULL,
    tax_number VARCHAR(50),
    security_clearance_level security_clearance_type DEFAULT 'RAHASIA_NEGARA',
    defence_industry_license_no VARCHAR(100), -- Izin Usaha Industri Pertahanan (KKIP)
    country VARCHAR(100),
    contact_person VARCHAR(100),
    email VARCHAR(100),
    phone VARCHAR(50),
    bank_account_name VARCHAR(150),
    bank_account_no VARCHAR(50),
    bank_name VARCHAR(100),
    performance_rating NUMERIC(3, 2) DEFAULT 0,
    is_approved BOOLEAN DEFAULT TRUE,
    created_by UUID,
    updated_by UUID,
    deleted_by UUID,
    created_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMPTZ,
    deleted_at TIMESTAMPTZ
);

-- -----------------------------------------------------------------------------
-- TABEL: proc_vendor_ratings
-- -----------------------------------------------------------------------------
CREATE TABLE IF NOT EXISTS proc_vendor_ratings (
    rating_id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    vendor_id UUID NOT NULL REFERENCES proc_vendors(vendor_id) ON DELETE CASCADE,
    evaluation_date DATE NOT NULL,
    evaluator_user_id UUID REFERENCES sys_users(user_id),
    quality_score NUMERIC(5, 2) NOT NULL,
    delivery_time_score NUMERIC(5, 2) NOT NULL,
    service_score NUMERIC(5, 2) NOT NULL,
    price_score NUMERIC(5, 2) NOT NULL,
    overall_score NUMERIC(5, 2) NOT NULL,
    remarks TEXT,
    created_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP
);

-- -----------------------------------------------------------------------------
-- TABEL: proc_requisitions
-- -----------------------------------------------------------------------------
CREATE TABLE IF NOT EXISTS proc_requisitions (
    requisition_id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    requisition_number VARCHAR(50) UNIQUE NOT NULL,
    origin_unit_id UUID NOT NULL REFERENCES org_units(unit_id),
    work_order_id UUID REFERENCES mro_work_orders(work_order_id),
    priority requisition_priority_type NOT NULL,
    requested_date DATE NOT NULL,
    required_by_date DATE,
    approval_status requisition_approval_status_type DEFAULT 'PENDING',
    approved_by_user_id UUID REFERENCES sys_users(user_id),
    approved_at TIMESTAMPTZ,
    total_estimated_cost NUMERIC(18, 2) DEFAULT 0,
    justification TEXT,
    created_by UUID,
    updated_by UUID,
    deleted_by UUID,
    created_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMPTZ,
    deleted_at TIMESTAMPTZ
);

CREATE TABLE IF NOT EXISTS proc_requisition_items (
    req_item_id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    requisition_id UUID NOT NULL REFERENCES proc_requisitions(requisition_id) ON DELETE CASCADE,
    material_id UUID NOT NULL REFERENCES inv_materials(material_id),
    quantity NUMERIC(12, 2) NOT NULL,
    estimated_unit_price NUMERIC(18, 2) NOT NULL,
    estimated_total_price NUMERIC(18, 2) GENERATED ALWAYS AS (quantity * estimated_unit_price) STORED,
    notes TEXT,
    created_by UUID,
    updated_by UUID,
    deleted_by UUID,
    created_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMPTZ,
    deleted_at TIMESTAMPTZ
);

-- -----------------------------------------------------------------------------
-- TABEL: proc_tenders
-- -----------------------------------------------------------------------------
CREATE TABLE IF NOT EXISTS proc_tenders (
    tender_id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    tender_number VARCHAR(50) UNIQUE NOT NULL,
    title VARCHAR(200) NOT NULL,
    procurement_category procurement_category_type NOT NULL,
    estimated_budget NUMERIC(18, 2) NOT NULL,
    procurement_method procurement_method_type NOT NULL,
    start_date DATE NOT NULL,
    closing_date DATE NOT NULL,
    status tender_status_type DEFAULT 'DRAFT',
    winner_vendor_id UUID REFERENCES proc_vendors(vendor_id),
    created_by UUID,
    updated_by UUID,
    deleted_by UUID,
    created_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMPTZ,
    deleted_at TIMESTAMPTZ
);

CREATE TABLE IF NOT EXISTS proc_tender_bids (
    bid_id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    tender_id UUID NOT NULL REFERENCES proc_tenders(tender_id) ON DELETE CASCADE,
    vendor_id UUID NOT NULL REFERENCES proc_vendors(vendor_id),
    bid_amount NUMERIC(18, 2) NOT NULL,
    submission_date TIMESTAMPTZ NOT NULL,
    technical_score NUMERIC(5, 2),
    commercial_score NUMERIC(5, 2),
    is_winner BOOLEAN DEFAULT FALSE,
    remarks TEXT,
    created_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP
);

-- -----------------------------------------------------------------------------
-- TABEL: proc_contracts
-- -----------------------------------------------------------------------------
CREATE TABLE IF NOT EXISTS proc_contracts (
    contract_id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    tender_id UUID REFERENCES proc_tenders(tender_id),
    contract_number VARCHAR(100) UNIQUE NOT NULL,
    vendor_id UUID NOT NULL REFERENCES proc_vendors(vendor_id),
    contract_title VARCHAR(200) NOT NULL,
    contract_value NUMERIC(18, 2) NOT NULL,
    currency VARCHAR(10) DEFAULT 'IDR',
    start_date DATE NOT NULL,
    end_date DATE NOT NULL,
    procurement_method procurement_method_type,
    warranty_period_months INT DEFAULT 12,
    tot_clause_summary TEXT,
    status contract_status_type DEFAULT 'ACTIVE',
    created_by UUID,
    updated_by UUID,
    deleted_by UUID,
    created_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMPTZ,
    deleted_at TIMESTAMPTZ
);

CREATE TABLE IF NOT EXISTS proc_contract_amendments (
    amendment_id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    contract_id UUID NOT NULL REFERENCES proc_contracts(contract_id) ON DELETE CASCADE,
    amendment_number VARCHAR(50) NOT NULL,
    amendment_date DATE NOT NULL,
    description TEXT NOT NULL,
    additional_value NUMERIC(18, 2) DEFAULT 0,
    extended_end_date DATE,
    created_by UUID,
    created_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP
);

-- -----------------------------------------------------------------------------
-- TABEL: proc_purchase_orders
-- -----------------------------------------------------------------------------
CREATE TABLE IF NOT EXISTS proc_purchase_orders (
    po_id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    po_number VARCHAR(50) UNIQUE NOT NULL,
    contract_id UUID REFERENCES proc_contracts(contract_id),
    vendor_id UUID NOT NULL REFERENCES proc_vendors(vendor_id),
    issuing_unit_id UUID NOT NULL REFERENCES org_units(unit_id),
    order_date DATE NOT NULL,
    delivery_deadline DATE,
    destination_warehouse_id UUID REFERENCES inv_warehouses(warehouse_id),
    total_amount NUMERIC(18, 2) NOT NULL,
    tax_amount NUMERIC(18, 2) DEFAULT 0,
    grand_total NUMERIC(18, 2) GENERATED ALWAYS AS (total_amount + tax_amount) STORED,
    status po_status_type DEFAULT 'ISSUED',
    created_by UUID,
    updated_by UUID,
    deleted_by UUID,
    created_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMPTZ,
    deleted_at TIMESTAMPTZ
);

CREATE TABLE IF NOT EXISTS proc_purchase_order_items (
    po_item_id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    po_id UUID NOT NULL REFERENCES proc_purchase_orders(po_id) ON DELETE CASCADE,
    material_id UUID NOT NULL REFERENCES inv_materials(material_id),
    quantity NUMERIC(12, 2) NOT NULL,
    unit_price NUMERIC(18, 2) NOT NULL,
    total_price NUMERIC(18, 2) GENERATED ALWAYS AS (quantity * unit_price) STORED,
    notes TEXT,
    created_by UUID,
    updated_by UUID,
    deleted_by UUID,
    created_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMPTZ,
    deleted_at TIMESTAMPTZ
);

-- -----------------------------------------------------------------------------
-- TABEL: proc_goods_receipts
-- -----------------------------------------------------------------------------
CREATE TABLE IF NOT EXISTS proc_goods_receipts (
    receipt_id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    receipt_number VARCHAR(50) UNIQUE NOT NULL,
    po_id UUID NOT NULL REFERENCES proc_purchase_orders(po_id),
    warehouse_id UUID NOT NULL REFERENCES inv_warehouses(warehouse_id),
    received_date TIMESTAMPTZ NOT NULL,
    delivery_order_number VARCHAR(100),
    inspected_by_user_id UUID NOT NULL REFERENCES sys_users(user_id),
    inspection_passed BOOLEAN DEFAULT TRUE,
    remarks TEXT,
    created_by UUID,
    updated_by UUID,
    deleted_by UUID,
    created_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMPTZ,
    deleted_at TIMESTAMPTZ
);

CREATE TABLE IF NOT EXISTS proc_goods_receipt_items (
    receipt_item_id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    receipt_id UUID NOT NULL REFERENCES proc_goods_receipts(receipt_id) ON DELETE CASCADE,
    po_item_id UUID NOT NULL REFERENCES proc_purchase_order_items(po_item_id),
    material_id UUID NOT NULL REFERENCES inv_materials(material_id),
    quantity_received NUMERIC(12, 2) NOT NULL,
    quantity_accepted NUMERIC(12, 2) NOT NULL,
    quantity_rejected NUMERIC(12, 2) DEFAULT 0,
    rejection_reason TEXT,
    created_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP
);

-- Indexes
CREATE INDEX IF NOT EXISTS idx_proc_vendors_code ON proc_vendors(vendor_code) WHERE deleted_at IS NULL;
CREATE INDEX IF NOT EXISTS idx_proc_requisitions_unit ON proc_requisitions(origin_unit_id) WHERE deleted_at IS NULL;
CREATE INDEX IF NOT EXISTS idx_proc_contracts_vendor ON proc_contracts(vendor_id) WHERE deleted_at IS NULL;
CREATE INDEX IF NOT EXISTS idx_proc_po_contract ON proc_purchase_orders(contract_id) WHERE deleted_at IS NULL;
CREATE INDEX IF NOT EXISTS idx_proc_goods_receipts_po ON proc_goods_receipts(po_id) WHERE deleted_at IS NULL;

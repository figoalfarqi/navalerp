-- =============================================================================
-- MODUL 5: KEUANGAN, ANGGARAN & AKUNTANSI PERTAHANAN (FINANCE & TCO)
-- FILE: 05_finance_budget_accounting/tables.sql
-- =============================================================================

CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

-- -----------------------------------------------------------------------------
-- ENUM TYPES: MODUL 5
-- -----------------------------------------------------------------------------
DO $$ BEGIN
    CREATE TYPE coa_account_type AS ENUM (
        'ASSET',
        'LIABILITY',
        'EQUITY',
        'REVENUE',
        'EXPENSE'
    );
EXCEPTION WHEN duplicate_object THEN null; END $$;

DO $$ BEGIN
    CREATE TYPE budget_program_status_type AS ENUM (
        'DRAFT',
        'APPROVED',
        'ACTIVE',
        'CLOSED'
    );
EXCEPTION WHEN duplicate_object THEN null; END $$;

DO $$ BEGIN
    CREATE TYPE budget_commitment_status_type AS ENUM (
        'COMMITTED',
        'REALISED',
        'CANCELLED'
    );
EXCEPTION WHEN duplicate_object THEN null; END $$;

DO $$ BEGIN
    CREATE TYPE invoice_verification_status_type AS ENUM (
        'PENDING_VERIF',
        'VERIFIED',
        'REJECTED'
    );
EXCEPTION WHEN duplicate_object THEN null; END $$;

DO $$ BEGIN
    CREATE TYPE invoice_payment_status_type AS ENUM (
        'UNPAID',
        'PARTIAL',
        'PAID'
    );
EXCEPTION WHEN duplicate_object THEN null; END $$;

DO $$ BEGIN
    CREATE TYPE payment_method_type AS ENUM (
        'KPPN_TREASURY',
        'BANK_TRANSFER',
        'CASH'
    );
EXCEPTION WHEN duplicate_object THEN null; END $$;

DO $$ BEGIN
    CREATE TYPE journal_source_type AS ENUM (
        'PROCUREMENT',
        'INVENTORY',
        'MRO',
        'ASSET_CAPITALIZATION',
        'PAYROLL'
    );
EXCEPTION WHEN duplicate_object THEN null; END $$;

-- -----------------------------------------------------------------------------
-- TABEL: fin_chart_of_accounts
-- -----------------------------------------------------------------------------
CREATE TABLE IF NOT EXISTS fin_chart_of_accounts (
    account_id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    account_code VARCHAR(50) UNIQUE NOT NULL,
    account_name VARCHAR(150) NOT NULL,
    account_type coa_account_type NOT NULL,
    parent_account_id UUID REFERENCES fin_chart_of_accounts(account_id),
    is_active BOOLEAN DEFAULT TRUE,
    created_by UUID,
    updated_by UUID,
    deleted_by UUID,
    created_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMPTZ,
    deleted_at TIMESTAMPTZ
);

-- -----------------------------------------------------------------------------
-- TABEL: fin_budget_programs
-- -----------------------------------------------------------------------------
CREATE TABLE IF NOT EXISTS fin_budget_programs (
    program_id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    fiscal_year INT NOT NULL,
    dipa_number VARCHAR(100) UNIQUE NOT NULL,
    program_code VARCHAR(50) NOT NULL,
    program_name VARCHAR(200) NOT NULL,
    total_budget NUMERIC(18, 2) NOT NULL,
    responsible_unit_id UUID NOT NULL REFERENCES org_units(unit_id),
    status budget_program_status_type DEFAULT 'ACTIVE',
    created_by UUID,
    updated_by UUID,
    deleted_by UUID,
    created_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMPTZ,
    deleted_at TIMESTAMPTZ
);

-- -----------------------------------------------------------------------------
-- TABEL: fin_budget_allocations
-- -----------------------------------------------------------------------------
CREATE TABLE IF NOT EXISTS fin_budget_allocations (
    allocation_id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    program_id UUID NOT NULL REFERENCES fin_budget_programs(program_id) ON DELETE CASCADE,
    activity_code VARCHAR(50) NOT NULL,
    activity_name VARCHAR(200) NOT NULL,
    target_unit_id UUID NOT NULL REFERENCES org_units(unit_id),
    ship_id UUID REFERENCES mro_ships(ship_id),
    account_id UUID NOT NULL REFERENCES fin_chart_of_accounts(account_id),
    allocated_amount NUMERIC(18, 2) NOT NULL,
    absorbed_amount NUMERIC(18, 2) DEFAULT 0,
    remaining_amount NUMERIC(18, 2) GENERATED ALWAYS AS (allocated_amount - absorbed_amount) STORED,
    created_by UUID,
    updated_by UUID,
    deleted_by UUID,
    created_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMPTZ,
    deleted_at TIMESTAMPTZ
);

-- -----------------------------------------------------------------------------
-- TABEL: fin_budget_commitments
-- -----------------------------------------------------------------------------
CREATE TABLE IF NOT EXISTS fin_budget_commitments (
    commitment_id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    commitment_number VARCHAR(50) UNIQUE NOT NULL,
    allocation_id UUID NOT NULL REFERENCES fin_budget_allocations(allocation_id),
    contract_id UUID REFERENCES proc_contracts(contract_id),
    po_id UUID REFERENCES proc_purchase_orders(po_id),
    work_order_id UUID REFERENCES mro_work_orders(work_order_id),
    committed_amount NUMERIC(18, 2) NOT NULL,
    commitment_date DATE NOT NULL,
    status budget_commitment_status_type DEFAULT 'COMMITTED',
    created_by UUID,
    updated_by UUID,
    deleted_by UUID,
    created_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMPTZ,
    deleted_at TIMESTAMPTZ
);

-- -----------------------------------------------------------------------------
-- TABEL: fin_invoices
-- -----------------------------------------------------------------------------
CREATE TABLE IF NOT EXISTS fin_invoices (
    invoice_id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    invoice_number VARCHAR(50) UNIQUE NOT NULL,
    vendor_id UUID NOT NULL REFERENCES proc_vendors(vendor_id),
    contract_id UUID REFERENCES proc_contracts(contract_id),
    po_id UUID REFERENCES proc_purchase_orders(po_id),
    invoice_date DATE NOT NULL,
    due_date DATE NOT NULL,
    tax_invoice_number VARCHAR(50),
    subtotal NUMERIC(18, 2) NOT NULL,
    tax_amount NUMERIC(18, 2) DEFAULT 0,
    total_amount NUMERIC(18, 2) GENERATED ALWAYS AS (subtotal + tax_amount) STORED,
    verification_status invoice_verification_status_type DEFAULT 'PENDING_VERIF',
    verified_by_user_id UUID REFERENCES sys_users(user_id),
    payment_status invoice_payment_status_type DEFAULT 'UNPAID',
    created_by UUID,
    updated_by UUID,
    deleted_by UUID,
    created_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMPTZ,
    deleted_at TIMESTAMPTZ
);

-- -----------------------------------------------------------------------------
-- TABEL: fin_payments
-- -----------------------------------------------------------------------------
CREATE TABLE IF NOT EXISTS fin_payments (
    payment_id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    payment_reference_no VARCHAR(50) UNIQUE NOT NULL, -- No SP2D / Bukti Bayar
    spp_number VARCHAR(50),
    spm_number VARCHAR(50),
    invoice_id UUID NOT NULL REFERENCES fin_invoices(invoice_id),
    payment_date DATE NOT NULL,
    amount_paid NUMERIC(18, 2) NOT NULL,
    payment_method payment_method_type DEFAULT 'KPPN_TREASURY',
    bank_source_account VARCHAR(100),
    authorised_by_user_id UUID REFERENCES sys_users(user_id),
    created_by UUID,
    updated_by UUID,
    deleted_by UUID,
    created_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMPTZ,
    deleted_at TIMESTAMPTZ
);

-- -----------------------------------------------------------------------------
-- TABEL: fin_journal_entries
-- -----------------------------------------------------------------------------
CREATE TABLE IF NOT EXISTS fin_journal_entries (
    journal_id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    entry_number VARCHAR(50) UNIQUE NOT NULL,
    entry_date DATE NOT NULL,
    description TEXT NOT NULL,
    source_module journal_source_type NOT NULL,
    source_reference_id UUID,
    is_posted BOOLEAN DEFAULT TRUE,
    created_by UUID,
    updated_by UUID,
    deleted_by UUID,
    created_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMPTZ,
    deleted_at TIMESTAMPTZ
);

CREATE TABLE IF NOT EXISTS fin_journal_lines (
    line_id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    journal_id UUID NOT NULL REFERENCES fin_journal_entries(journal_id) ON DELETE CASCADE,
    account_id UUID NOT NULL REFERENCES fin_chart_of_accounts(account_id),
    debit NUMERIC(18, 2) DEFAULT 0,
    credit NUMERIC(18, 2) DEFAULT 0,
    memo VARCHAR(255),
    created_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP
);

-- -----------------------------------------------------------------------------
-- TABEL: fin_platform_tco_summaries
-- -----------------------------------------------------------------------------
CREATE TABLE IF NOT EXISTS fin_platform_tco_summaries (
    tco_id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    ship_id UUID NOT NULL REFERENCES mro_ships(ship_id),
    fiscal_year INT NOT NULL,
    acquisition_amortization NUMERIC(18, 2) DEFAULT 0,
    fuel_lube_cost NUMERIC(18, 2) DEFAULT 0,
    mro_spareparts_cost NUMERIC(18, 2) DEFAULT 0,
    docking_services_cost NUMERIC(18, 2) DEFAULT 0,
    crew_payroll_allowances NUMERIC(18, 2) DEFAULT 0,
    modernization_upgrades_cost NUMERIC(18, 2) DEFAULT 0,
    total_annual_operating_cost NUMERIC(18, 2) GENERATED ALWAYS AS (
        fuel_lube_cost + mro_spareparts_cost + docking_services_cost + crew_payroll_allowances + modernization_upgrades_cost
    ) STORED,
    operating_hours_sea NUMERIC(10, 2) DEFAULT 0,
    cost_per_operating_hour NUMERIC(18, 2) DEFAULT 0,
    remarks TEXT,
    created_by UUID,
    updated_by UUID,
    deleted_by UUID,
    created_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMPTZ,
    deleted_at TIMESTAMPTZ,
    CONSTRAINT uq_ship_fiscal_year UNIQUE (ship_id, fiscal_year)
);

-- Indexes
CREATE INDEX IF NOT EXISTS idx_fin_coa_code ON fin_chart_of_accounts(account_code) WHERE deleted_at IS NULL;
CREATE INDEX IF NOT EXISTS idx_fin_allocations_program ON fin_budget_allocations(program_id) WHERE deleted_at IS NULL;
CREATE INDEX IF NOT EXISTS idx_fin_allocations_ship ON fin_budget_allocations(ship_id) WHERE deleted_at IS NULL;
CREATE INDEX IF NOT EXISTS idx_fin_invoices_vendor ON fin_invoices(vendor_id) WHERE deleted_at IS NULL;
CREATE INDEX IF NOT EXISTS idx_fin_payments_invoice ON fin_payments(invoice_id) WHERE deleted_at IS NULL;
CREATE INDEX IF NOT EXISTS idx_fin_journal_lines_acc ON fin_journal_lines(account_id);
CREATE INDEX IF NOT EXISTS idx_fin_tco_ship ON fin_platform_tco_summaries(ship_id);

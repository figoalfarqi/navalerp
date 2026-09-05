-- =============================================================================
-- MODUL 9: MANAJEMEN DOKUMEN ELEKTRONIK & DIGITAL THREAD (EDRMS)
-- FILE: 09_edrms_documents/tables.sql
-- =============================================================================

CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

-- -----------------------------------------------------------------------------
-- ENUM TYPES: MODUL 9
-- -----------------------------------------------------------------------------
DO $$ BEGIN
    CREATE TYPE document_confidentiality_level_type AS ENUM (
        'SANGAT_RAHASIA',
        'RAHASIA_NEGARA',
        'RAHASIA',
        'TERBATAS',
        'BIASA'
    );
EXCEPTION WHEN duplicate_object THEN null; END $$;

DO $$ BEGIN
    CREATE TYPE document_status_type AS ENUM (
        'DRAFT',
        'REVIEW',
        'APPROVED',
        'ACTIVE',
        'ARCHIVED',
        'SUPERSEDED'
    );
EXCEPTION WHEN duplicate_object THEN null; END $$;

DO $$ BEGIN
    CREATE TYPE document_linked_entity_type AS ENUM (
        'MRO_SHIP',
        'MRO_EQUIPMENT',
        'INV_MATERIAL',
        'PROC_CONTRACT',
        'MRO_WORK_ORDER'
    );
EXCEPTION WHEN duplicate_object THEN null; END $$;

DO $$ BEGIN
    CREATE TYPE document_link_purpose_type AS ENUM (
        'OPERATING_MANUAL',
        'WIRING_DIAGRAM',
        'CERT_KELAIKAN',
        'CONTRACT_AGREEMENT'
    );
EXCEPTION WHEN duplicate_object THEN null; END $$;

-- -----------------------------------------------------------------------------
-- TABEL: doc_categories
-- -----------------------------------------------------------------------------
CREATE TABLE IF NOT EXISTS doc_categories (
    category_id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    category_code VARCHAR(50) UNIQUE NOT NULL,
    category_name VARCHAR(150) NOT NULL,
    retention_years INT DEFAULT 10,
    confidentiality_level document_confidentiality_level_type DEFAULT 'RAHASIA',
    description TEXT,
    created_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP
);

-- -----------------------------------------------------------------------------
-- TABEL: doc_documents
-- -----------------------------------------------------------------------------
CREATE TABLE IF NOT EXISTS doc_documents (
    document_id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    document_number VARCHAR(100) UNIQUE NOT NULL,
    title VARCHAR(255) NOT NULL,
    category_id UUID NOT NULL REFERENCES doc_categories(category_id),
    originating_unit_id UUID REFERENCES org_units(unit_id),
    classification_level document_confidentiality_level_type DEFAULT 'RAHASIA',
    effective_date DATE NOT NULL,
    expiry_date DATE,
    status document_status_type DEFAULT 'ACTIVE',
    approved_by_user_id UUID REFERENCES sys_users(user_id),
    created_by UUID,
    updated_by UUID,
    deleted_by UUID,
    created_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMPTZ,
    deleted_at TIMESTAMPTZ
);

-- -----------------------------------------------------------------------------
-- TABEL: doc_document_versions
-- -----------------------------------------------------------------------------
CREATE TABLE IF NOT EXISTS doc_document_versions (
    version_id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    document_id UUID NOT NULL REFERENCES doc_documents(document_id) ON DELETE CASCADE,
    version_number VARCHAR(20) NOT NULL,
    file_name VARCHAR(255) NOT NULL,
    file_path VARCHAR(500) NOT NULL,
    file_size_bytes BIGINT NOT NULL,
    file_hash_sha256 VARCHAR(64) NOT NULL,
    mime_type VARCHAR(100) NOT NULL,
    change_summary TEXT,
    uploaded_by_user_id UUID REFERENCES sys_users(user_id),
    created_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
    CONSTRAINT uq_doc_version UNIQUE (document_id, version_number)
);

-- -----------------------------------------------------------------------------
-- TABEL: doc_document_links
-- -----------------------------------------------------------------------------
CREATE TABLE IF NOT EXISTS doc_document_links (
    link_id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    document_id UUID NOT NULL REFERENCES doc_documents(document_id) ON DELETE CASCADE,
    entity_type document_linked_entity_type NOT NULL,
    entity_id UUID NOT NULL,
    link_purpose document_link_purpose_type NOT NULL,
    created_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP
);

-- Indexes
CREATE INDEX IF NOT EXISTS idx_doc_docs_number ON doc_documents(document_number) WHERE deleted_at IS NULL;
CREATE INDEX IF NOT EXISTS idx_doc_docs_category ON doc_documents(category_id) WHERE deleted_at IS NULL;
CREATE INDEX IF NOT EXISTS idx_doc_links_entity ON doc_document_links(entity_type, entity_id);

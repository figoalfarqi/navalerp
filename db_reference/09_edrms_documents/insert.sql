-- =============================================================================
-- SEED DATA: MODUL 9 - MANAJEMEN DOKUMEN ELEKTRONIK & DIGITAL THREAD (EDRMS)
-- FILE: 09_edrms_documents/insert.sql
-- =============================================================================

-- 1. Kategori Dokumen Pertahanan
INSERT INTO doc_categories (
    category_id, category_code, category_name, retention_years,
    confidentiality_level, description
) VALUES
(
    '89000000-0000-0000-0000-000000000001',
    'TECH_MANUAL',
    'Buku Petunjuk Teknis & Operasional (OEM Technical Manual)',
    25,
    'TERBATAS',
    'Buku panduan pengoperasian, perbaikan, dan spesifikasi komponen alutsista'
),
(
    '89000000-0000-0000-0000-000000000002',
    'CERT_KELAIKAN',
    'Sertifikat Kelaikan Militer (Seaworthiness & Readiness)',
    10,
    'RAHASIA',
    'Sertifikat kelaikan operasi kapal perang yang diterbitkan Dislaikmatal'
),
(
    '89000000-0000-0000-0000-000000000003',
    'BLUEPRINT_CAD',
    'Gambar Rancang Bangun & Skema Kelistrikan Kapal',
    30,
    'RAHASIA',
    'Gambar teknik galangan kapal, piping diagram, dan single line wiring'
),
(
    '89000000-0000-0000-0000-000000000004',
    'CONTRACT_LEGAL',
    'Dokumen Kontrak Pengadaan & Klausul ToT Rahasia Negara',
    20,
    'RAHASIA_NEGARA',
    'Perjanjian hukum pengadaan alutsista dengan klausul pertahanan strategis'
)
ON CONFLICT (category_id) DO NOTHING;

-- 2. Master Dokumen
INSERT INTO doc_documents (
    document_id, document_number, title, category_id,
    originating_unit_id, classification_level, effective_date, expiry_date,
    status, approved_by_user_id
) VALUES
(
    '90000000-0000-0000-0000-000000000001',
    'MNL-MTU-20V4000-M53B',
    'Technical & Maintenance Manual MTU 20V 4000 M53B Marine Diesel Engine',
    '89000000-0000-0000-0000-000000000001', -- TECH_MANUAL
    '10000000-0000-0000-0000-000000000005', -- Fasharkan Sby
    'TERBATAS',
    '2024-01-01',
    '2034-12-31',
    'APPROVED',
    '20000000-0000-0000-0000-000000000003'  -- Aslog
),
(
    '90000000-0000-0000-0000-000000000002',
    'CERT-LAIK-REM331-2026',
    'Sertifikat Kelaikan Operasi Laut KRI Raden Eddy Martadinata-331 TA 2026',
    '89000000-0000-0000-0000-000000000002', -- CERT_KELAIKAN
    '10000000-0000-0000-0000-000000000001', -- Mabesal (Dislaikmatal)
    'RAHASIA',
    '2026-01-10',
    '2027-01-10',
    'APPROVED',
    '20000000-0000-0000-0000-000000000002'  -- Pangkoarmada II
),
(
    '90000000-0000-0000-0000-000000000003',
    'DWG-SIGMA-CMS-004',
    'Interconnection Wiring Schematic TACTICOS CMS to SMART-S Mk2 Radar',
    '89000000-0000-0000-0000-000000000003', -- BLUEPRINT_CAD
    '10000000-0000-0000-0000-000000000004', -- KRI REM-331
    'RAHASIA',
    '2023-05-15',
    NULL,
    'APPROVED',
    '20000000-0000-0000-0000-000000000004'  -- Dan KRI
)
ON CONFLICT (document_id) DO NOTHING;

-- 3. Versi & Berkas Dokumen
INSERT INTO doc_document_versions (
    version_id, document_id, version_number, file_name, file_path,
    file_size_bytes, file_hash_sha256, mime_type, change_summary, uploaded_by_user_id
) VALUES
(
    '90500000-0000-0000-0000-000000000001',
    '90000000-0000-0000-0000-000000000001',
    'v1.0',
    'mtu_20v4000_m53b_maintenance_manual.pdf',
    '/secure_storage/naval_edrms/tech_manuals/mtu_20v4000_m53b.pdf',
    18450200,
    'e3b0c44298fc1c149afbf4c8996fb92427ae41e4649b934ca495991b7852b855',
    'application/pdf',
    'Dokumen asli pabrikan OEM Rolls-Royce Solutions edisi bahasa Inggris',
    '20000000-0000-0000-0000-000000000005'
),
(
    '90500000-0000-0000-0000-000000000002',
    '90000000-0000-0000-0000-000000000002',
    'v1.0',
    'sertifikat_laik_laut_kri_rem_331_2026.pdf',
    '/secure_storage/naval_edrms/certificates/laik_rem_331_2026.pdf',
    2150000,
    'dca148408a287964b4458f4679720478051ec7495029e2f4705cbab29a6745ef',
    'application/pdf',
    'Sertifikat Kelaikan Penuh Hasil Uji Petik & Uji Laut Dislaikmatal',
    '20000000-0000-0000-0000-000000000003'
),
(
    '90500000-0000-0000-0000-000000000003',
    '90000000-0000-0000-0000-000000000003',
    'v2.1',
    'sigma_tacticos_smarts_wiring_rev2.dwg',
    '/secure_storage/naval_edrms/blueprints/sigma_tacticos_smarts_wiring.dwg',
    45600000,
    '9f86d081884c7d659a2feaa0c55ad015a3bf4f1b2b0b822cd15d6c15b0f00a08',
    'application/acad',
    'Revisi integrasi protokol data bus mil-std-1553B',
    '20000000-0000-0000-0000-000000000004'
)
ON CONFLICT (version_id) DO NOTHING;

-- 4. Benang Merah Digital (Digital Thread Links)
INSERT INTO doc_document_links (
    link_id, document_id, entity_type, entity_id, link_purpose
) VALUES
(
    '90800000-0000-0000-0000-000000000001',
    '90000000-0000-0000-0000-000000000001',
    'MRO_EQUIPMENT',
    '43000000-0000-0000-0000-000000000001', -- Mesin Pokok MTU KRI REM-331
    'OPERATING_MANUAL'
),
(
    '90800000-0000-0000-0000-000000000002',
    '90000000-0000-0000-0000-000000000002',
    'MRO_SHIP',
    '41000000-0000-0000-0000-000000000001', -- Kapal KRI REM-331
    'CERT_KELAIKAN'
),
(
    '90800000-0000-0000-0000-000000000003',
    '90000000-0000-0000-0000-000000000003',
    'MRO_SHIP',
    '41000000-0000-0000-0000-000000000001',
    'WIRING_DIAGRAM'
)
ON CONFLICT (link_id) DO NOTHING;


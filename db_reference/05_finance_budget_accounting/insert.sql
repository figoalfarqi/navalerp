-- =============================================================================
-- SEED DATA: MODUL 5 - KEUANGAN, ANGGARAN & AKUNTANSI PERTAHANAN (FINANCE & TCO)
-- FILE: 05_finance_budget_accounting/insert.sql
-- =============================================================================

-- 1. Bagan Akun Standar (Chart of Accounts)
INSERT INTO fin_chart_of_accounts (
    account_id, account_code, account_name, account_type, is_active
) VALUES
(
    '70000000-0000-0000-0000-000000000001',
    '111110',
    'Kas dan Setara Kas Bendahara Pengeluaran TNI AL',
    'ASSET',
    TRUE
),
(
    '70000000-0000-0000-0000-000000000002',
    '117111',
    'Persediaan Amunisi & Senjata Strategis',
    'ASSET',
    TRUE
),
(
    '70000000-0000-0000-0000-000000000003',
    '117112',
    'Persediaan Suku Cadang Alutsista Matra Laut',
    'ASSET',
    TRUE
),
(
    '70000000-0000-0000-0000-000000000004',
    '135111',
    'Aset Tetap Alutsista Kapal Perang Republik Indonesia (KRI)',
    'ASSET',
    TRUE
),
(
    '70000000-0000-0000-0000-000000000005',
    '212111',
    'Utang Belanja Pengadaan Barang / Jasa Rekanan Pertahanan',
    'LIABILITY',
    TRUE
),
(
    '70000000-0000-0000-0000-000000000006',
    '521811',
    'Belanja Barang Persediaan Pemeliharaan Alutsista Laut',
    'EXPENSE',
    TRUE
),
(
    '70000000-0000-0000-0000-000000000007',
    '523111',
    'Belanja Pemeliharaan & Jasa Docking KRI',
    'EXPENSE',
    TRUE
)
ON CONFLICT (account_id) DO UPDATE SET
    account_name = EXCLUDED.account_name,
    account_type = EXCLUDED.account_type;

-- 2. Program Kerja & Anggaran Pertahanan (DIPA)
INSERT INTO fin_budget_programs (
    program_id, fiscal_year, dipa_number, program_code, program_name,
    total_budget, responsible_unit_id, status
) VALUES
(
    '71000000-0000-0000-0000-000000000001',
    2026,
    'DIPA-012.01.1.412001/2026',
    '012.01.WA',
    'Program Modernisasi Alutsista & Dukungan Kesiapan Operasi Laut Koarmada II',
    450000000000.00,
    '10000000-0000-0000-0000-000000000003', -- Koarmada II
    'ACTIVE'
)
ON CONFLICT (program_id) DO NOTHING;

-- 3. Alokasi Anggaran Kegiatan / Output Platform KRI
INSERT INTO fin_budget_allocations (
    allocation_id, program_id, activity_code, activity_name,
    target_unit_id, ship_id, account_id, allocated_amount, absorbed_amount
) VALUES
(
    '71500000-0000-0000-0000-000000000001',
    '71000000-0000-0000-0000-000000000001',
    'WA.5231.001',
    'Pemeliharaan Terencana dan Perbaikan KRI REM-331 TA 2026',
    '10000000-0000-0000-0000-000000000007', -- Satkor Koarmada II
    '41000000-0000-0000-0000-000000000001', -- KRI REM-331
    '70000000-0000-0000-0000-000000000007', -- Belanja Pemeliharaan
    25000000000.00,
    4850000000.00
),
(
    '71500000-0000-0000-0000-000000000002',
    '71000000-0000-0000-0000-000000000001',
    'WA.5218.002',
    'Pengadaan Suku Cadang Mesin Pendorong Pokok MTU Kelas Martadinata',
    '10000000-0000-0000-0000-000000000003', -- Koarmada II
    '41000000-0000-0000-0000-000000000001',
    '70000000-0000-0000-0000-000000000006', -- Belanja Persediaan
    15000000000.00,
    346000000.00
),
(
    '71500000-0000-0000-0000-000000000003',
    '71000000-0000-0000-0000-000000000001',
    'WA.5231.003',
    'Docking Rutin & Pelapisan Antifouling KRI DPO-365',
    '10000000-0000-0000-0000-000000000009', -- Fasharkan Surabaya
    '41000000-0000-0000-0000-000000000003', -- KRI DPO-365
    '70000000-0000-0000-0000-000000000007',
    12500000000.00,
    0.00
)
ON CONFLICT (allocation_id) DO NOTHING;

-- 4. Komitmen Anggaran
INSERT INTO fin_budget_commitments (
    commitment_id, commitment_number, allocation_id, contract_id, po_id,
    committed_amount, commitment_date, status
) VALUES
(
    '72000000-0000-0000-0000-000000000001',
    'KMT-2026-0034',
    '71500000-0000-0000-0000-000000000001',
    '62000000-0000-0000-0000-000000000001', -- Kontrak MTU
    '63000000-0000-0000-0000-000000000001', -- PO MTU
    4850000000.00,
    '2026-02-12',
    'COMMITTED'
)
ON CONFLICT (commitment_id) DO NOTHING;

-- 5. Tagihan Rekanan (Invoices)
INSERT INTO fin_invoices (
    invoice_id, invoice_number, vendor_id, contract_id, po_id,
    invoice_date, due_date, tax_invoice_number, subtotal, tax_amount,
    verification_status, verified_by_user_id, payment_status
) VALUES
(
    '73000000-0000-0000-0000-000000000001',
    'INV-MTU-2026-0012',
    '60000000-0000-0000-0000-000000000004', -- MTU
    '62000000-0000-0000-0000-000000000001',
    '63000000-0000-0000-0000-000000000001',
    '2026-03-05',
    '2026-04-05',
    '010.000-26.90182741',
    346000000.00,
    38060000.00,
    'VERIFIED',
    '20000000-0000-0000-0000-000000000006', -- Letkol Deni Mulyadi (Perwira Keuangan)
    'PAID'
)
ON CONFLICT (invoice_id) DO NOTHING;

-- 6. Pembayaran Perbendaharaan Militer (SP2D)
INSERT INTO fin_payments (
    payment_id, payment_reference_no, spp_number, spm_number, invoice_id,
    payment_date, amount_paid, payment_method, bank_source_account,
    authorised_by_user_id
) VALUES
(
    '74000000-0000-0000-0000-000000000001',
    'SP2D-2026-KPPN-00991',
    'SPP-0012/KOARMADA2/2026',
    'SPM-0012/KOARMADA2/2026',
    '73000000-0000-0000-0000-000000000001',
    '2026-03-12',
    384060000.00,
    'KPPN_TREASURY',
    'KPPN Jakarta II - Rekening Kas Negara 000.12345.1',
    '20000000-0000-0000-0000-000000000002' -- Panglima
)
ON CONFLICT (payment_id) DO NOTHING;

-- 7. Jurnal Akuntansi & Buku Besar
INSERT INTO fin_journal_entries (
    journal_id, entry_number, entry_date, description,
    source_module, source_reference_id, is_posted
) VALUES
(
    '75000000-0000-0000-0000-000000000001',
    'JRN-2026-03-001',
    '2026-03-05',
    'Penerimaan suku cadang mesin MTU 20V 4000 untuk persediaan alutsista KRI REM-331',
    'PROCUREMENT',
    '73000000-0000-0000-0000-000000000001',
    TRUE
)
ON CONFLICT (journal_id) DO NOTHING;

INSERT INTO fin_journal_lines (
    line_id, journal_id, account_id, debit, credit, memo
) VALUES
(
    '75100000-0000-0000-0000-000000000001',
    '75000000-0000-0000-0000-000000000001',
    '70000000-0000-0000-0000-000000000003', -- Persediaan Suku Cadang
    346000000.00,
    0.00,
    'Pencatatan persediaan suku cadang BAPHP-2026-0045'
),
(
    '75100000-0000-0000-0000-000000000002',
    '75000000-0000-0000-0000-000000000001',
    '70000000-0000-0000-0000-000000000005', -- Utang Belanja Pengadaan
    0.00,
    346000000.00,
    'Kewajiban pembayaran termin vendor MTU'
)
ON CONFLICT (line_id) DO NOTHING;

-- 8. Platform Total Cost of Ownership (TCO)
INSERT INTO fin_platform_tco_summaries (
    tco_id, ship_id, fiscal_year, acquisition_amortization,
    fuel_lube_cost, mro_spareparts_cost, docking_services_cost,
    crew_payroll_allowances, modernization_upgrades_cost,
    operating_hours_sea, cost_per_operating_hour, remarks
) VALUES
(
    '76000000-0000-0000-0000-000000000001',
    '41000000-0000-0000-0000-000000000001', -- KRI REM-331
    2025,
    150000000000.00,
    18500000000.00,
    8400000000.00,
    6500000000.00,
    14200000000.00,
    2100000000.00,
    2150.00,
    23116279.07,
    'TCO Realisasi TA 2025 KRI REM-331 - Operasi Siaga Samudera dan Latihan Bersama Kakadu'
),
(
    '76000000-0000-0000-0000-000000000002',
    '41000000-0000-0000-0000-000000000001', -- KRI REM-331
    2026,
    150000000000.00,
    20000000000.00,
    9500000000.00,
    8000000000.00,
    15000000000.00,
    3500000000.00,
    2400.00,
    23333333.33,
    'TCO Proyeksi Pagu TA 2026 KRI REM-331 - Kesiapan Tempur Penuh Natuna'
)
ON CONFLICT (tco_id) DO NOTHING;

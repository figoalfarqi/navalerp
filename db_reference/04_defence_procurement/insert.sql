-- =============================================================================
-- SEED DATA: MODUL 4 - PENGADAAN PERTAHANAN (DEFENCE PROCUREMENT) (NAVALERP)
-- FILE: 04_defence_procurement/insert.sql
-- =============================================================================

-- 1. Master Vendor / Rekanan Industri Pertahanan
INSERT INTO proc_vendors (
    vendor_id, vendor_code, vendor_name, tax_number, security_clearance_level,
    defence_industry_license_no, country, contact_person, email, phone,
    bank_account_name, bank_account_no, bank_name, performance_rating, is_approved
) VALUES
(
    '60000000-0000-0000-0000-000000000001',
    'VND-PAL-001',
    'PT PAL INDONESIA (PERSERO)',
    '01.000.123.4-051.000',
    'RAHASIA_NEGARA',
    'KKIP/DEF-IND/2021/001',
    'Indonesia',
    'Ir. Bambang Soedjarwo',
    'procurement@pal.co.id',
    '+62-31-3292275',
    'PT PAL INDONESIA',
    '1420009876543',
    'Bank Mandiri',
    4.85,
    TRUE
),
(
    '60000000-0000-0000-0000-000000000002',
    'VND-PINDAD-002',
    'PT PINDAD (PERSERO)',
    '01.000.456.7-042.000',
    'RAHASIA_NEGARA',
    'KKIP/DEF-IND/2020/004',
    'Indonesia',
    'Drs. Joko Prasetyo',
    'sales_defence@pindad.com',
    '+62-22-7312012',
    'PT PINDAD PERSERO',
    '1310001234567',
    'Bank BNI',
    4.70,
    TRUE
),
(
    '60000000-0000-0000-0000-000000000003',
    'VND-LEN-003',
    'PT LEN INDUSTRI (PERSERO) / DEFEND ID',
    '01.000.789.1-041.000',
    'RAHASIA_NEGARA',
    'KKIP/DEF-IND/2022/012',
    'Indonesia',
    'M. Raditya, M.T.',
    'naval_combat@len.co.id',
    '+62-22-5202682',
    'PT LEN INDUSTRI',
    '0700003344556',
    'Bank Mandiri',
    4.65,
    TRUE
),
(
    '60000000-0000-0000-0000-000000000004',
    'VND-MTU-004',
    'MTU FRIEDRICHSHAFEN GMBH / ROLLS-ROYCE SOLUTIONS',
    'DE145892341',
    'NATO_SECRET',
    'BMVg-GE-MRO-88912',
    'Germany',
    'Hans-Juergen Becker',
    'government_sales@rolls-royce.com',
    '+49-7541-90-0',
    'MTU Friedrichshafen GmbH',
    'DE89370400440532013000',
    'Deutsche Bank AG',
    4.90,
    TRUE
),
(
    '60000000-0000-0000-0000-000000000005',
    'VND-THALES-005',
    'THALES NEDERLAND B.V.',
    'NL001239842B01',
    'NATO_SECRET',
    'NLD-MOD-NAV-3310',
    'Netherlands',
    'Pieter van den Berg',
    'naval_systems@nl.thalesgroup.com',
    '+31-74-248-1111',
    'Thales Nederland BV',
    'NL91ABNA0417164300',
    'ABN AMRO',
    4.88,
    TRUE
)
ON CONFLICT (vendor_id) DO NOTHING;

-- 2. Penilaian Kinerja Vendor Pertahanan
INSERT INTO proc_vendor_ratings (
    rating_id, vendor_id, evaluation_date, evaluator_user_id,
    quality_score, delivery_time_score, service_score, price_score, overall_score, remarks
) VALUES
(
    '60100000-0000-0000-0000-000000000001',
    '60000000-0000-0000-0000-000000000001',
    '2025-12-15',
    '20000000-0000-0000-0000-000000000003', -- Kolonel Aslog
    92.50, 88.00, 95.00, 85.00, 90.12,
    'Sangat memuaskan pada pekerjaan docking KRI REM-331 dan kepatuhan mil-spec galangan nasional.'
),
(
    '60100000-0000-0000-0000-000000000002',
    '60000000-0000-0000-0000-000000000004',
    '2025-11-20',
    '20000000-0000-0000-0000-000000000003',
    98.00, 94.00, 96.00, 82.00, 92.50,
    'Kualitas spare parts OEM MTU 20V 4000 sangat tinggi, sertifikat Certificate of Conformity lengkap.'
)
ON CONFLICT (rating_id) DO NOTHING;

-- 3. Purchase Requisitions (Kebutuhan Pengadaan)
INSERT INTO proc_requisitions (
    requisition_id, requisition_number, origin_unit_id, work_order_id, priority,
    requested_date, required_by_date, approval_status, approved_by_user_id, approved_at,
    total_estimated_cost, justification
) VALUES
(
    '61000000-0000-0000-0000-000000000001',
    'REQ-2026-001',
    '10000000-0000-0000-0000-000000000004', -- KRI REM-331
    '45000000-0000-0000-0000-000000000001', -- WO Corrective Repair
    'HIGH',
    '2026-01-10',
    '2026-02-15',
    'APPROVED',
    '20000000-0000-0000-0000-000000000003', -- Approved by Aslog
    '2026-01-12 10:30:00+07',
    346000000.00,
    'Pengadaan suku cadang kritis MTU 20V 4000 M53B untuk kesiapan Operasi Siaga Tempur Laut Natuna'
)
ON CONFLICT (requisition_id) DO NOTHING;

INSERT INTO proc_requisition_items (
    req_item_id, requisition_id, material_id, quantity, estimated_unit_price, notes
) VALUES
(
    '61100000-0000-0000-0000-000000000001',
    '61000000-0000-0000-0000-000000000001',
    '51000000-0000-0000-0000-000000000001', -- Oil Filter MTU
    8.00,
    3500000.00,
    'Penggantian berkala Main Engine Port & Stbd'
),
(
    '61100000-0000-0000-0000-000000000002',
    '61000000-0000-0000-0000-000000000001',
    '51000000-0000-0000-0000-000000000002', -- Fuel Filter
    12.00,
    2750000.00,
    'Filter separator BBM B35 standar TNI AL'
),
(
    '61100000-0000-0000-0000-000000000003',
    '61000000-0000-0000-0000-000000000001',
    '51000000-0000-0000-0000-000000000003', -- Injector Nozzle
    10.00,
    28500000.00,
    'Injector nozzle common rail MTU 20V 4000'
)
ON CONFLICT (req_item_id) DO NOTHING;

-- 4. Tender Pengadaan Alutsista & MRO
INSERT INTO proc_tenders (
    tender_id, tender_number, title, procurement_category, estimated_budget,
    procurement_method, start_date, closing_date, status, winner_vendor_id
) VALUES
(
    '61500000-0000-0000-0000-000000000001',
    'TND-2026-MRO-001',
    'Pengadaan Suku Cadang & Dukungan Teknis Mesin Pokok MTU Kelas Martadinata TA 2026',
    'SPARE_PARTS',
    5000000000.00,
    'LIMITED_TENDER',
    '2026-01-15',
    '2026-02-05',
    'AWARDED',
    '60000000-0000-0000-0000-000000000004' -- MTU Friedrichshafen
)
ON CONFLICT (tender_id) DO NOTHING;

INSERT INTO proc_tender_bids (
    bid_id, tender_id, vendor_id, bid_amount, submission_date,
    technical_score, commercial_score, is_winner, remarks
) VALUES
(
    '61600000-0000-0000-0000-000000000001',
    '61500000-0000-0000-0000-000000000001',
    '60000000-0000-0000-0000-000000000004', -- MTU
    4850000000.00,
    '2026-02-01 14:00:00+07',
    96.50,
    92.00,
    TRUE,
    'Pemenang lelang: Principal OEM dengan lisensi militer resmi'
),
(
    '61600000-0000-0000-0000-000000000002',
    '61500000-0000-0000-0000-000000000001',
    '60000000-0000-0000-0000-000000000001', -- PT PAL
    4950000000.00,
    '2026-02-02 11:30:00+07',
    91.00,
    89.00,
    FALSE,
    'Penawaran mitra lokal galangan terakreditasi'
)
ON CONFLICT (bid_id) DO NOTHING;

-- 5. Kontrak Pengadaan
INSERT INTO proc_contracts (
    contract_id, tender_id, contract_number, vendor_id, contract_title,
    contract_value, currency, start_date, end_date, procurement_method,
    warranty_period_months, tot_clause_summary, status
) VALUES
(
    '62000000-0000-0000-0000-000000000001',
    '61500000-0000-0000-0000-000000000001',
    'CTR-ALUT-2026-014',
    '60000000-0000-0000-0000-000000000004', -- MTU
    'Kontrak Pengadaan Dukungan Suku Cadang Mesin Pendorong Pokok KRI REM-331',
    4850000000.00,
    'IDR',
    '2026-02-10',
    '2026-12-31',
    'LIMITED_TENDER',
    24,
    'Transfer teknologi overhaul intermediate level ke Fasharkan Surabaya & pelatihan teknisi Dislambair/Dislaikmatal.',
    'ACTIVE'
)
ON CONFLICT (contract_id) DO NOTHING;

INSERT INTO proc_contract_amendments (
    amendment_id, contract_id, amendment_number, amendment_date, description,
    additional_value, extended_end_date
) VALUES
(
    '62100000-0000-0000-0000-000000000001',
    '62000000-0000-0000-0000-000000000001',
    'AMD-CTR-2026-014-01',
    '2026-06-01',
    'Penambahan cakupan kalibrasi ECU MTU ADEC pada uji laut (sea trial)',
    150000000.00,
    '2027-02-28'
)
ON CONFLICT (amendment_id) DO NOTHING;

-- 6. Purchase Orders
INSERT INTO proc_purchase_orders (
    po_id, po_number, contract_id, vendor_id, issuing_unit_id,
    order_date, delivery_deadline, destination_warehouse_id,
    total_amount, tax_amount, status
) VALUES
(
    '63000000-0000-0000-0000-000000000001',
    'PO-2026-0089',
    '62000000-0000-0000-0000-000000000001',
    '60000000-0000-0000-0000-000000000004', -- MTU
    '10000000-0000-0000-0000-000000000002', -- Koarmada II
    '2026-02-15',
    '2026-03-15',
    '50000000-0000-0000-0000-000000000001', -- Gudang Disbekal
    346000000.00,
    38060000.00, -- PPN 11%
    'COMPLETED'
)
ON CONFLICT (po_id) DO NOTHING;

INSERT INTO proc_purchase_order_items (
    po_item_id, po_id, material_id, quantity, unit_price, notes
) VALUES
(
    '63100000-0000-0000-0000-000000000001',
    '63000000-0000-0000-0000-000000000001',
    '51000000-0000-0000-0000-000000000001',
    8.00,
    3500000.00,
    'Oil Filter Element 20V 4000'
),
(
    '63100000-0000-0000-0000-000000000002',
    '63000000-0000-0000-0000-000000000001',
    '51000000-0000-0000-0000-000000000002',
    12.00,
    2750000.00,
    'Fuel Filter Water Separator'
),
(
    '63100000-0000-0000-0000-000000000003',
    '63000000-0000-0000-0000-000000000001',
    '51000000-0000-0000-0000-000000000003',
    10.00,
    28500000.00,
    'Main Engine High-Pressure Common Rail Injector'
)
ON CONFLICT (po_item_id) DO NOTHING;

-- 7. Goods Receipt (BAPHP: Berita Acara Penerimaan Hasil Pekerjaan)
INSERT INTO proc_goods_receipts (
    receipt_id, receipt_number, po_id, warehouse_id,
    received_date, delivery_order_number, inspected_by_user_id,
    inspection_passed, remarks
) VALUES
(
    '64000000-0000-0000-0000-000000000001',
    'BAPHP-2026-0045',
    '63000000-0000-0000-0000-000000000001',
    '50000000-0000-0000-0000-000000000001', -- Gudang Disbekal
    '2026-03-02 09:30:00+07',
    'DO-MTU-SGP-99120',
    '20000000-0000-0000-0000-000000000005', -- Diperiksa oleh Mayor Tri Wibowo (Maint Officer)
    TRUE,
    'Barang diterima lengkap sesuai spesifikasi militer OEM MTU, segel utuh dan lulus uji fungsi awal.'
)
ON CONFLICT (receipt_id) DO NOTHING;

INSERT INTO proc_goods_receipt_items (
    receipt_item_id, receipt_id, po_item_id, material_id,
    quantity_received, quantity_accepted, quantity_rejected, rejection_reason
) VALUES
(
    '64100000-0000-0000-0000-000000000001',
    '64000000-0000-0000-0000-000000000001',
    '63100000-0000-0000-0000-000000000001',
    '51000000-0000-0000-0000-000000000001',
    8.00,
    8.00,
    0.00,
    NULL
),
(
    '64100000-0000-0000-0000-000000000002',
    '64000000-0000-0000-0000-000000000001',
    '63100000-0000-0000-0000-000000000002',
    '51000000-0000-0000-0000-000000000002',
    12.00,
    12.00,
    0.00,
    NULL
),
(
    '64100000-0000-0000-0000-000000000003',
    '64000000-0000-0000-0000-000000000001',
    '63100000-0000-0000-0000-000000000003',
    '51000000-0000-0000-0000-000000000003',
    10.00,
    10.00,
    0.00,
    NULL
)
ON CONFLICT (receipt_item_id) DO NOTHING;

-- =============================================================================
-- SEED DATA: MODUL 7 - INFRASTRUKTUR PANGKALAN & LABUH (BASE & PORT INFRASTRUCTURE)
-- FILE: 07_base_infrastructure/insert.sql
-- =============================================================================

-- 1. Master Fasilitas Pangkalan Militer
INSERT INTO infra_facilities (
    facility_id, base_unit_id, facility_code, facility_name,
    facility_type, length_meters, draft_depth_meters, max_displacement_tonnage,
    has_shore_power, has_fresh_water, has_fuel_bunker_line, status
) VALUES
(
    '85000000-0000-0000-0000-000000000001',
    '10000000-0000-0000-0000-000000000003', -- Koarmada II
    'FAC-DERM-MADURA',
    'Dermaga Madura Koarmada II Ujung',
    'BERTH_JETTY',
    350.00,
    11.50,
    12000.00,
    TRUE,
    TRUE,
    TRUE,
    'OPERATIONAL'
),
(
    '85000000-0000-0000-0000-000000000002',
    '10000000-0000-0000-0000-000000000003', -- Koarmada II
    'FAC-DERM-SEMARANG',
    'Dermaga Semarang Koarmada II Ujung',
    'BERTH_JETTY',
    220.00,
    8.50,
    4500.00,
    TRUE,
    TRUE,
    FALSE,
    'OPERATIONAL'
),
(
    '85000000-0000-0000-0000-000000000003',
    '10000000-0000-0000-0000-000000000009', -- Fasharkan Surabaya
    'FAC-DOCK-GRAVING-1',
    'Graving Dock I Fasharkan Surabaya',
    'GRAVING_DOCK',
    150.00,
    7.50,
    6000.00,
    TRUE,
    TRUE,
    FALSE,
    'OPERATIONAL'
),
(
    '85000000-0000-0000-0000-000000000004',
    '10000000-0000-0000-0000-000000000006', -- Lantamal V
    'FAC-BUNKER-UJUNG',
    'Instalasi Tangki Timbun BBM Disbekal Lantamal V',
    'FUEL_STORAGE',
    NULL,
    NULL,
    NULL,
    FALSE,
    FALSE,
    TRUE,
    'OPERATIONAL'
)
ON CONFLICT (facility_id) DO UPDATE SET
    facility_name = EXCLUDED.facility_name,
    status = EXCLUDED.status;

-- 2. Penjadwalan & Riwayat Sandar Kapal
INSERT INTO infra_berth_bookings (
    booking_id, facility_id, ship_id, booking_purpose, eta, etd,
    actual_berth_time, actual_unberth_time, shore_power_kwh_used,
    fresh_water_ton_used, status, approved_by_user_id, remarks
) VALUES
(
    '85500000-0000-0000-0000-000000000001',
    '85000000-0000-0000-0000-000000000001', -- Dermaga Madura
    '41000000-0000-0000-0000-000000000001', -- KRI REM-331
    'LOGISTIC_REPLENISHMENT',
    '2026-02-20 08:00:00+07',
    '2026-02-28 16:00:00+07',
    '2026-02-20 08:15:00+07',
    '2026-02-28 15:45:00+07',
    14500.00,
    85.00,
    'COMPLETED',
    '20000000-0000-0000-0000-000000000004', -- Perwira Logistik
    'Sandar untuk bekal ulang amunisi, bahan bakar, dan air tawar persiapan Operasi Siaga Tempur Laut Natuna.'
),
(
    '85500000-0000-0000-0000-000000000002',
    '85000000-0000-0000-0000-000000000003', -- Graving Dock I
    '41000000-0000-0000-0000-000000000003', -- KRI DPO-365
    'MRO_REPAIR',
    '2026-03-01 07:00:00+07',
    '2026-03-25 17:00:00+07',
    '2026-03-01 07:30:00+07',
    NULL,
    28000.00,
    40.00,
    'BERTHED',
    '20000000-0000-0000-0000-000000000004',
    'Docking berkala, sandblasting lambung dan pembersihan sea chest di Dok Fasharkan.'
)
ON CONFLICT (booking_id) DO NOTHING;

-- 3. Bunkering BBM Militer
INSERT INTO infra_fuel_bunker_records (
    bunker_id, ship_id, facility_id, fuel_type, quantity_liters,
    density_15c, flow_rate_lph, bunkering_start_time, bunkering_end_time,
    receipt_voucher_no, authorised_by_user_id
) VALUES
(
    '86000000-0000-0000-0000-000000000001',
    '41000000-0000-0000-0000-000000000001', -- KRI REM-331
    '85000000-0000-0000-0000-000000000001', -- Dermaga Madura
    'HSD_MILSPEC',
    120000.00,
    0.8420,
    35000.00,
    '2026-02-26 09:00:00+07',
    '2026-02-26 13:00:00+07',
    'BPM/KOARMADA2/2026/0289',
    '20000000-0000-0000-0000-000000000005' -- Diotorisasi Kadepsin Arif Wijaya
)
ON CONFLICT (bunker_id) DO NOTHING;

-- 4. Pemeliharaan Fasilitas Pangkalan
INSERT INTO infra_facility_maintenances (
    maint_id, facility_id, maintenance_type, start_date, end_date,
    cost, performed_by, status, remarks
) VALUES
(
    '86500000-0000-0000-0000-000000000001',
    '85000000-0000-0000-0000-000000000001',
    'DREDGING_ALUR',
    '2025-11-01',
    '2025-12-10',
    4200000000.00,
    'Dinas Fasilitas Pangkalan TNI AL (Disfaslanal)',
    'COMPLETED',
    'Pengerukan kolam pelabuhan Dermaga Madura untuk memastikan kedalaman aman -11.5m LWS bagi Frigat SIGMA.'
)
ON CONFLICT (maint_id) DO NOTHING;

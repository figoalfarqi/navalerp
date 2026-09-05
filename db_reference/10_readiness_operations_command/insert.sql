-- =============================================================================
-- SEED DATA: MODUL 10 - KESIAPAN OPERASI & KOMANDO TEMPUR (READINESS & OPERATIONS)
-- FILE: 10_readiness_operations_command/insert.sql
-- =============================================================================

-- 1. Wilayah Teater Operasi Laut
INSERT INTO ops_theaters (
    theater_id, theater_code, theater_name, responsible_command_unit_id,
    threat_level, description
) VALUES
(
    '91000000-0000-0000-0000-000000000001',
    'THEATER-NATUNA',
    'Teater Operasi Laut Natuna Utara & Selat Karimata',
    '10000000-0000-0000-0000-000000000002', -- Koarmada II (BKO Kogabwilhan I)
    'DEFCON_2',
    'Wilayah pengamanan ZEE perbatasan laut strategis dengan intensitas patroli tinggi'
),
(
    '91000000-0000-0000-0000-000000000002',
    'THEATER-AMBALAT',
    'Teater Operasi Laut Sulawesi & Blok Ambalat',
    '10000000-0000-0000-0000-000000000002', -- Koarmada II
    'DEFCON_3',
    'Wilayah kedaulatan maritim blok migas perbatasan Indonesia - Malaysia'
)
ON CONFLICT (theater_id) DO NOTHING;

-- 2. Operasi & Misi Tempur Militer
INSERT INTO ops_missions (
    mission_id, theater_id, mission_code, mission_name, mission_type,
    start_date, end_date, commanding_officer_user_id, mission_status
) VALUES
(
    '91500000-0000-0000-0000-000000000001',
    '91000000-0000-0000-0000-000000000001', -- Natuna
    'OPS-GARDA-SAMUDERA-26',
    'Operasi Siaga Tempur Laut Garda Samudera 26 (Natuna)',
    'COMBAT_PATROL',
    '2026-03-01',
    '2026-06-30',
    '20000000-0000-0000-0000-000000000002', -- Pangkoarmada II
    'ACTIVE'
)
ON CONFLICT (mission_id) DO NOTHING;

-- 3. Penugasan Kapal Perang dalam Gugus Tugas
INSERT INTO ops_mission_ship_assignments (
    assignment_id, mission_id, ship_id, tactical_callsign,
    role_in_task_force, joined_date, status
) VALUES
(
    '91800000-0000-0000-0000-000000000001',
    '91500000-0000-0000-0000-000000000001',
    '41000000-0000-0000-0000-000000000001', -- KRI REM-331
    'BRAVO-ONE',
    'FLAGSHIP',
    '2026-03-01',
    'ACTIVE'
)
ON CONFLICT (assignment_id) DO NOTHING;

-- 4. Buku Jurnal Harian Kapal Berlayar (Daily Navigational Log)
INSERT INTO ops_daily_logs (
    log_id, ship_id, log_date, latitude, longitude,
    heading_degrees, speed_knots, sea_state, weather_condition,
    fuel_remaining_liters, fresh_water_remaining_tons,
    tactical_summary, logged_by_user_id
) VALUES
(
    '92000000-0000-0000-0000-000000000001',
    '41000000-0000-0000-0000-000000000001', -- KRI REM-331
    '2026-03-04',
    4.251400,
    108.384200,
    45,
    16.5,
    3,
    'Cerah Berawan, Angin Timur Laut 12 Knots',
    112500.00,
    78.00,
    'Patroli sektor Alpha Laut Natuna Utara. Kontak radar permukaan terpantau normal. Seluruh sensor CMS TACTICOS dan mesin MTU beroperasi optimal.',
    '20000000-0000-0000-0000-000000000004' -- Letkol Ahmad Dahlan (Dan KRI)
)
ON CONFLICT (log_id) DO NOTHING;

-- 5. Indeks Kesiapan Komposit Alutsista KRI (Composite Readiness Index)
INSERT INTO ops_ship_readiness_snapshots (
    snapshot_id, ship_id, snapshot_timestamp, readiness_category,
    mro_readiness_score, personnel_manning_score, logistics_supply_score, remarks
) VALUES
(
    '92500000-0000-0000-0000-000000000001',
    '41000000-0000-0000-0000-000000000001', -- KRI REM-331
    '2026-03-04 08:00:00+07',
    'C-1', -- Siap Tempur Penuh
    96.50,
    98.00,
    94.00,
    'Kesiapan tempur KRI REM-331 level tertinggi (C-1 Siap Tempur Penuh) untuk misi patroli kedaulatan maritim laut terluar.'
),
(
    '92500000-0000-0000-0000-000000000002',
    '41000000-0000-0000-0000-000000000002', -- KRI DPO-365
    '2026-03-04 08:00:00+07',
    'C-4', -- Docking / Non-Operasional Sementara
    60.00,
    85.00,
    70.00,
    'KRI DPO-365 berada dalam tahap docking pemeliharaan rutin di Fasharkan Surabaya. Dijadwalkan kembali C-1 pada akhir bulan.'
)
ON CONFLICT (snapshot_id) DO NOTHING;

-- 6. Peringatan Dini Kesiapan (Readiness Alert)
INSERT INTO ops_readiness_alerts (
    alert_id, ship_id, equipment_id, severity, alert_type,
    alert_message, is_acknowledged, acknowledged_by_user_id, acknowledged_at
) VALUES
(
    '93000000-0000-0000-0000-000000000001',
    '41000000-0000-0000-0000-000000000001', -- KRI REM-331
    '43000000-0000-0000-0000-000000000001', -- Mesin MTU
    'MEDIUM',
    'CASREP_DEFECT',
    'Jam kerja mesin MTU Port telah mencapai 4.250 jam. Persiapkan jadwal perawatan berkala 500 jam berikutnya (PMS Level 2).',
    TRUE,
    '20000000-0000-0000-0000-000000000005', -- Kadepsin Tri Wibowo
    '2026-03-04 10:00:00+07'
)
ON CONFLICT (alert_id) DO NOTHING;


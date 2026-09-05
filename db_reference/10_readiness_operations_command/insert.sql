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
    '10000000-0000-0000-0000-000000000003', -- Koarmada II (BKO Kogabwilhan I)
    'DEFCON_2',
    'Wilayah pengamanan ZEE perbatasan laut strategis dengan intensitas patroli tinggi'
),
(
    '91000000-0000-0000-0000-000000000002',
    'THEATER-AMBALAT',
    'Teater Operasi Laut Sulawesi & Blok Ambalat',
    '10000000-0000-0000-0000-000000000003', -- Koarmada II
    'DEFCON_3',
    'Wilayah kedaulatan maritim blok migas perbatasan Indonesia - Malaysia'
),
(
    '91000000-0000-0000-0000-000000000003',
    'THEATER-MALAKA',
    'Teater Operasi Selat Malaka & Selat Singapura',
    '10000000-0000-0000-0000-000000000002', -- Koarmada I
    'DEFCON_4',
    'Jalur Sea Lines of Communication (SLOC) perdagangan internasional tersibuk di dunia'
)
ON CONFLICT (theater_id) DO UPDATE SET
    theater_name = EXCLUDED.theater_name,
    threat_level = EXCLUDED.threat_level;

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
    '20000000-0000-0000-0000-000000000002', -- Panglima
    'ACTIVE'
),
(
    '91500000-0000-0000-0000-000000000002',
    '91000000-0000-0000-0000-000000000003', -- Selat Malaka
    'OPS-PATKOR-MALINDO-26',
    'Patroli Terkoordinasi Malindo Wilayah Selat Malaka',
    'COMBAT_PATROL',
    '2026-02-15',
    '2026-05-15',
    '20000000-0000-0000-0000-000000000002',
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
),
(
    '91800000-0000-0000-0000-000000000002',
    '91500000-0000-0000-0000-000000000001',
    '41000000-0000-0000-0000-000000000002', -- KRI INR-332
    'BRAVO-TWO',
    'ESCORT_SURFACE',
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
    '20000000-0000-0000-0000-000000000003' -- Kolonel Hendra Kurniawan (Dan KRI REM-331)
),
(
    '92000000-0000-0000-0000-000000000002',
    '41000000-0000-0000-0000-000000000002', -- KRI INR-332
    '2026-03-04',
    4.421000,
    108.512000,
    48,
    17.0,
    3,
    'Cerah Berawan, Jarak Pandang 10 NM',
    135000.00,
    82.00,
    'Manuver taktis formasi Line of Bearing mendampingi KRI REM-331.',
    '20000000-0000-0000-0000-000000000008' -- Kolonel Faisal Anwar (Dan KRI INR-332)
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
    '41000000-0000-0000-0000-000000000003', -- KRI DPO-365
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
    '20000000-0000-0000-0000-000000000005', -- Kadepsin Arif Wijaya
    '2026-03-04 10:00:00+07'
)
ON CONFLICT (alert_id) DO NOTHING;

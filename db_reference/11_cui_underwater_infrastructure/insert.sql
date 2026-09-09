-- =============================================================================
-- MODUL 11: SEED DATA INFRASTRUKTUR BAWAH LAUT KRITIS (CUI)
-- FILE: 11_cui_underwater_infrastructure/insert.sql
-- =============================================================================

-- Seed Assets
INSERT INTO cui_assets (
    cui_asset_id, asset_code, asset_name, asset_type, operator_name,
    theater_id, depth_meters, length_km, latitude, longitude,
    start_coordinates, end_coordinates, status, health_score, protection_priority,
    last_inspected_at, next_inspection_due, notes, created_by
) VALUES
(
    '9c000000-0000-0000-0000-000000000001',
    'CUI-CBL-001',
    'Kabel SKKL Palapa Ring Barat - Segmen Natuna',
    'SUBMARINE_CABLE',
    'PT Telkom Indonesia Tbk',
    'c0000000-0000-0000-0000-000000000001',
    85.00,
    345.50,
    3.916700,
    108.383300,
    '3.9167, 108.3833 (Ranai)',
    '1.1304, 104.0530 (Batam)',
    'ACTIVE_MONITORED',
    96,
    'CRITICAL_TIER_1',
    NOW() - INTERVAL '20 days',
    NOW() + INTERVAL '70 days',
    'Tulang punggung telekomunikasi perbatasan utara NKRI dan pangkalan KRI Ranai',
    'SYSTEM'
),
(
    '9c000000-0000-0000-0000-000000000002',
    'CUI-PIP-001',
    'Pipa Gas Bawah Laut Natuna - Batam (WNAT)',
    'SUBSEA_PIPELINE',
    'PT Perusahaan Gas Negara (PGN) / ConocoPhillips',
    'c0000000-0000-0000-0000-000000000001',
    72.50,
    470.00,
    2.800000,
    106.500000,
    '4.5000, 107.0000 (Blok B Natuna)',
    '1.1200, 104.0200 (Grissik/Batam)',
    'ACTIVE_MONITORED',
    92,
    'CRITICAL_TIER_1',
    NOW() - INTERVAL '45 days',
    NOW() + INTERVAL '45 days',
    'Jalur pasokan energi gas strategis lintas batas dan industri domestik Batam-Singapura',
    'SYSTEM'
),
(
    '9c000000-0000-0000-0000-000000000003',
    'CUI-LND-001',
    'Cable Landing Station (CLS) Batam Centre',
    'LANDING_STATION',
    'PT Telekomunikasi Indonesia International (Telin)',
    'c0000000-0000-0000-0000-000000000001',
    0.00,
    0.00,
    1.130100,
    104.053200,
    '1.1301, 104.0532',
    '1.1301, 104.0532',
    'ACTIVE_MONITORED',
    98,
    'CRITICAL_TIER_1',
    NOW() - INTERVAL '15 days',
    NOW() + INTERVAL '75 days',
    'Stasiun pendaratan kabel internasional SEA-ME-WE-5 dan Palapa Ring',
    'SYSTEM'
),
(
    '9c000000-0000-0000-0000-000000000004',
    'CUI-OFF-001',
    'Offshore FSO Gagak Rimang & Anjungan Banyu Urip',
    'OFFSHORE_ENERGY',
    'ExxonMobil Cepu Limited / Pertamina',
    'c0000000-0000-0000-0000-000000000002',
    42.00,
    25.00,
    -6.750000,
    111.850000,
    '-6.8500, 111.8000',
    '-6.7500, 111.8500',
    'ACTIVE_MONITORED',
    94,
    'CRITICAL_TIER_1',
    NOW() - INTERVAL '30 days',
    NOW() + INTERVAL '60 days',
    'Floating Storage and Offloading lifting minyak mentah nasional Laut Jawa',
    'SYSTEM'
),
(
    '9c000000-0000-0000-0000-000000000005',
    'CUI-CBL-002',
    'Kabel Interkoneksi Listrik Bawah Laut Selat Sunda',
    'SUBMARINE_CABLE',
    'PT PLN (Persero)',
    'c0000000-0000-0000-0000-000000000002',
    95.00,
    38.00,
    -5.900000,
    105.850000,
    '-5.8500, 105.7500 (Bakauheni)',
    '-5.9500, 105.9500 (Merak)',
    'ALERT_ANOMALY',
    78,
    'CRITICAL_TIER_1',
    NOW() - INTERVAL '10 days',
    NOW() + INTERVAL '10 days',
    'Peringatan getaran dan proximity kapal kargo terdeteksi di koordinat km 14 Selat Sunda',
    'SYSTEM'
),
(
    '9c000000-0000-0000-0000-000000000006',
    'CUI-MON-001',
    'Sensor Array Sonar Pasif Pemantau ALKI II - Selat Lombok',
    'MONITORING_SYSTEM',
    'Pusat Komando Operasi AL (Puskodal) Mabesal',
    'c0000000-0000-0000-0000-000000000003',
    350.00,
    18.50,
    -8.450000,
    115.720000,
    '-8.4000, 115.7000',
    '-8.5000, 115.7400',
    'ACTIVE_MONITORED',
    95,
    'CRITICAL_TIER_1',
    NOW() - INTERVAL '5 days',
    NOW() + INTERVAL '85 days',
    'Sistem deteksi akustik bawah air perlintasan kapal selam asing ALKI II',
    'SYSTEM'
),
(
    '9c000000-0000-0000-0000-000000000007',
    'CUI-CBL-003',
    'Kabel Laut SKKL Morowali - Kendari (Palapa Ring Tengah)',
    'SUBMARINE_CABLE',
    'PT Telkom Indonesia Tbk',
    'c0000000-0000-0000-0000-000000000003',
    120.00,
    180.00,
    -3.200000,
    122.500000,
    '-2.8000, 122.1500 (Morowali)',
    '-3.9800, 122.6000 (Kendari)',
    'INSPECTION_REQUIRED',
    84,
    'HIGH_TIER_2',
    NOW() - INTERVAL '120 days',
    NOW() - INTERVAL '5 days',
    'Jadwal inspeksi ROV jatuh tempo untuk survei ketebalan sedimen pelindung',
    'SYSTEM'
)
ON CONFLICT (cui_asset_id) DO UPDATE SET
    asset_name = EXCLUDED.asset_name,
    status = EXCLUDED.status,
    health_score = EXCLUDED.health_score;

-- Seed Monitoring Logs
INSERT INTO cui_monitoring_logs (
    log_id, cui_asset_id, sensor_code, sensor_type, log_time,
    metric_value, metric_unit, status, vessel_proximity_mmsi, anomaly_score, description, created_by
) VALUES
(
    '9c100000-0000-0000-0000-000000000001',
    '9c000000-0000-0000-0000-000000000005',
    'DTS-SS-01',
    'FIBER_OPTIC_DTS',
    NOW() - INTERVAL '45 minutes',
    14.20,
    'bar',
    'WARNING',
    '525119822',
    0.785,
    'Tekanan hidrostatik berfluktuasi disertai kapal kargo curah menurunkan jangkar di koridor kabel',
    'SYSTEM'
),
(
    '9c100000-0000-0000-0000-000000000002',
    '9c000000-0000-0000-0000-000000000001',
    'HYDRO-NAT-01',
    'ACOUSTIC_SONAR',
    NOW() - INTERVAL '15 minutes',
    32.50,
    'dB',
    'NORMAL',
    NULL,
    0.045,
    'Pola frekuensi akustik normal, tidak ada indikasi aktivitas pukat harimau atau trawl ilegal',
    'SYSTEM'
),
(
    '9c100000-0000-0000-0000-000000000003',
    '9c000000-0000-0000-0000-000000000006',
    'SONAR-LMBK-03',
    'ACOUSTIC_SONAR',
    NOW() - INTERVAL '2 hours',
    48.10,
    'dB',
    'NORMAL',
    '525001234',
    0.120,
    'Kontak sonar terverifikasi: kapal kontainer komersial melintas jalur TSS ALKI II dengan kecepatan 14 knot',
    'SYSTEM'
)
ON CONFLICT (log_id) DO NOTHING;

-- Seed CUI Alerts
INSERT INTO cui_alerts (
    alert_id, alert_code, cui_asset_id, alert_type, severity, detected_at,
    assigned_ship_id, status, ai_confidence, recommended_action, resolution_notes, created_by
) VALUES
(
    '9c200000-0000-0000-0000-000000000001',
    'ALT-CUI-2026-001',
    '9c000000-0000-0000-0000-000000000005',
    'VESSEL_ANCHOR_DRAG_RISK',
    'CRITICAL',
    NOW() - INTERVAL '40 minutes',
    '31000000-0000-0000-0000-000000000001',
    'DISPATCHED',
    0.945,
    'Segera kontak radio VHF Ch.16 ke kapal MV Ocean Trader (MMSI 525119822) untuk angkat jangkar dan gerakkan KRI terdekat guna pengamanan jalur kabel Selat Sunda.',
    'KRI Raden Eddy Martadinata-331 diarahkan merapat ke lokasi kontak.',
    'AI_ENGINE'
),
(
    '9c200000-0000-0000-0000-000000000002',
    'ALT-CUI-2026-002',
    '9c000000-0000-0000-0000-000000000002',
    'PRESSURE_DROP',
    'HIGH',
    NOW() - INTERVAL '3 hours',
    NULL,
    'INVESTIGATING',
    0.880,
    'Verifikasi telemetri SCADA PGN dan jadwalkan inspeksi visual drone bawah air (ROV) pada pipa gas Natuna.',
    NULL,
    'AI_ENGINE'
)
ON CONFLICT (alert_id) DO NOTHING;

-- Seed CUI Inspections
INSERT INTO cui_inspections (
    inspection_id, inspection_number, cui_asset_id, ship_id, inspection_date,
    inspector_officer_id, method, condition_rating, findings, remedial_action_required,
    next_inspection_date, created_by
) VALUES
(
    '9c300000-0000-0000-0000-000000000001',
    'INSP-CUI-2026-001',
    '9c000000-0000-0000-0000-000000000001',
    '31000000-0000-0000-0000-000000000001',
    CURRENT_DATE - INTERVAL '20 days',
    '81000000-0000-0000-0000-000000000003',
    'ROV_SUBMERSIBLE',
    'EXCELLENT',
    'Kabel tertanam baik di bawah sedimen pasir laut kedalaman 85 meter. Pelindung beton artikulasi utuh.',
    FALSE,
    CURRENT_DATE + INTERVAL '70 days',
    'SYSTEM'
)
ON CONFLICT (inspection_id) DO NOTHING;


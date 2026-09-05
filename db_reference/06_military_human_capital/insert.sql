-- =============================================================================
-- SEED DATA: MODUL 6 - MANAJEMEN PERSONEL & AWAK KAPAL (MILITARY HUMAN CAPITAL)
-- FILE: 06_military_human_capital/insert.sql
-- =============================================================================

-- 1. Master Kepangkatan Militer TNI AL
INSERT INTO hcm_ranks (rank_id, rank_code, rank_name, rank_category, nato_rank_code, seniority_order) VALUES
('80000000-0000-0000-0000-000000000001', 'LAKSDYA', 'Laksamana Madya TNI', 'PATI', 'OF-8', 2),
('80000000-0000-0000-0000-000000000002', 'LAKSDA',  'Laksamana Muda TNI',   'PATI', 'OF-7', 3),
('80000000-0000-0000-0000-000000000003', 'LAKSMA',  'Laksamana Pertama TNI','PATI', 'OF-6', 4),
('80000000-0000-0000-0000-000000000004', 'KOLONEL', 'Kolonel Laut',         'PAMEN', 'OF-5', 5),
('80000000-0000-0000-0000-000000000005', 'LETKOL',  'Letnan Kolonel Laut',  'PAMEN', 'OF-4', 6),
('80000000-0000-0000-0000-000000000006', 'MAYOR',   'Mayor Laut',           'PAMEN', 'OF-3', 7),
('80000000-0000-0000-0000-000000000007', 'KAPTEN',  'Kapten Laut',          'PAMA',  'OF-2', 8),
('80000000-0000-0000-0000-000000000008', 'LETTU',   'Letnan Satu Laut',     'PAMA',  'OF-1', 9),
('80000000-0000-0000-0000-000000000009', 'SERKA',   'Sersan Kepala',        'BINTARA', 'OR-6', 15),
('80000000-0000-0000-0000-000000000010', 'KLK',     'Kelasi Kepala',        'TAMTAMA', 'OR-4', 21)
ON CONFLICT (rank_id) DO NOTHING;

-- 2. Master Korps TNI AL
INSERT INTO hcm_corps (corps_id, corps_code, corps_name, description) VALUES
('80500000-0000-0000-0000-000000000001', 'P',  'Korps Pelaut', 'Navigasi, komando kapal, dan operasi tempur laut'),
('80500000-0000-0000-0000-000000000002', 'T',  'Korps Teknik', 'Pemeliharaan mesin pendorong, struktur lambung, dan kelistrikan kapal'),
('80500000-0000-0000-0000-000000000003', 'E',  'Korps Elektronika', 'Sistem sensor, radar, senjata terpadu CMS, dan komunikasi tempur'),
('80500000-0000-0000-0000-000000000004', 'S',  'Korps Suplai', 'Logistik perbekalan, pergudangan matbek, dan administrasi keuangan'),
('80500000-0000-0000-0000-000000000005', 'M',  'Korps Marinir', 'Pasukan pendarat amfibi dan pertahanan pangkalan')
ON CONFLICT (corps_id) DO NOTHING;

-- 3. Master Personel Militer
INSERT INTO hcm_personnel (
    personnel_id, nrp, full_name, rank_id, corps_id, current_unit_id,
    current_position, birth_place, birth_date, gender, blood_type, religion,
    education_level, service_entry_date, user_id, status
) VALUES
(
    '81000000-0000-0000-0000-000000000001',
    '11223/P',
    'Surya Pratama, S.E., M.M.',
    '80000000-0000-0000-0000-000000000002', -- Laksda
    '80500000-0000-0000-0000-000000000001', -- Korps Pelaut
    '10000000-0000-0000-0000-000000000002', -- Koarmada II
    'Panglima Komando Armada II',
    'Surabaya', '1970-08-17', 'MALE', 'O', 'ISLAM', 'AAL 1993', '1989-08-01',
    '20000000-0000-0000-0000-000000000002', -- linked user
    'ACTIVE'
),
(
    '81000000-0000-0000-0000-000000000002',
    '13450/T',
    'Budi Santoso, S.T., M.Tr.Opsla',
    '80000000-0000-0000-0000-000000000004', -- Kolonel Laut (T)
    '80500000-0000-0000-0000-000000000002', -- Korps Teknik
    '10000000-0000-0000-0000-000000000002', -- Koarmada II
    'Asisten Logistik Pangkoarmada II',
    'Semarang', '1976-03-24', 'MALE', 'A', 'ISLAM', 'AAL 1998', '1994-08-01',
    '20000000-0000-0000-0000-000000000003', -- linked user
    'ACTIVE'
),
(
    '81000000-0000-0000-0000-000000000003',
    '15200/P',
    'Ahmad Dahlan, M.Tr.Hanla',
    '80000000-0000-0000-0000-000000000005', -- Letkol Laut (P)
    '80500000-0000-0000-0000-000000000001', -- Korps Pelaut
    '10000000-0000-0000-0000-000000000004', -- KRI REM-331
    'Komandan KRI Raden Eddy Martadinata-331',
    'Jakarta', '1981-11-10', 'MALE', 'B', 'ISLAM', 'AAL 2003', '1999-08-01',
    '20000000-0000-0000-0000-000000000004', -- linked user
    'ACTIVE'
),
(
    '81000000-0000-0000-0000-000000000004',
    '16780/T',
    'Tri Wibowo, S.T.',
    '80000000-0000-0000-0000-000000000006', -- Mayor Laut (T)
    '80500000-0000-0000-0000-000000000002', -- Korps Teknik
    '10000000-0000-0000-0000-000000000004', -- KRI REM-331
    'Kepala Departemen Mesin (Kadepsin) KRI REM-331',
    'Yogyakarta', '1986-05-15', 'MALE', 'O', 'ISLAM', 'AAL 2008', '2004-08-01',
    '20000000-0000-0000-0000-000000000005', -- linked user
    'ACTIVE'
),
(
    '81000000-0000-0000-0000-000000000005',
    '17890/S',
    'Hendra Kusuma, S.Sos.',
    '80000000-0000-0000-0000-000000000007', -- Kapten Laut (S)
    '80500000-0000-0000-0000-000000000004', -- Korps Suplai
    '10000000-0000-0000-0000-000000000002', -- Koarmada II
    'Perwira Administrasi Personel (Spers Koarmada II)',
    'Bandung', '1990-02-28', 'MALE', 'AB', 'ISLAM', 'AAL 2012', '2008-08-01',
    '20000000-0000-0000-0000-000000000006', -- linked user
    'ACTIVE'
),
(
    '81000000-0000-0000-0000-000000000006',
    '102340',
    'Didik Supriyadi',
    '80000000-0000-0000-0000-000000000009', -- Serka
    '80500000-0000-0000-0000-000000000002', -- Korps Teknik
    '10000000-0000-0000-0000-000000000004', -- KRI REM-331
    'Bintara Utama Mesin Pokok (Bama Divisi Mesin)',
    'Malang', '1992-07-12', 'MALE', 'O', 'ISLAM', 'Secaba 2013', '2013-03-01',
    NULL,
    'ACTIVE'
)
ON CONFLICT (personnel_id) DO NOTHING;

-- 4. Riwayat Penugasan Militer
INSERT INTO hcm_service_records (
    record_id, personnel_id, order_letter_number, assignment_type,
    from_unit_id, to_unit_id, position_title, start_date, end_date, remarks
) VALUES
(
    '81500000-0000-0000-0000-000000000001',
    '81000000-0000-0000-0000-000000000003', -- Letkol Ahmad Dahlan
    'Kep/120/V/2024',
    'PROMOTION',
    '10000000-0000-0000-0000-000000000002',
    '10000000-0000-0000-0000-000000000004',
    'Komandan KRI Raden Eddy Martadinata-331',
    '2024-05-15',
    NULL,
    'Pengangkatan dalam jabatan Komandan Frigat SIGMA 10514'
)
ON CONFLICT (record_id) DO NOTHING;

-- 5. Master Kualifikasi & Brevet
INSERT INTO hcm_qualifications (
    qualification_id, qualification_code, qualification_name, qualification_category,
    issuing_institution, validity_years, description
) VALUES
(
    '82000000-0000-0000-0000-000000000001',
    'BRV-COMMAND-SEA',
    'Brevet Komando Kapal Perang Republik Indonesia',
    'BREVET',
    'MABESAL',
    10,
    'Kualifikasi kepemimpinan komando tempur permukaan laut'
),
(
    '82000000-0000-0000-0000-000000000002',
    'CERT-MTU-ENG-ADV',
    'Sertifikasi Teknisi Ahli Mesin Diesel MTU 20V 4000 M53B',
    'CERTIFICATION',
    'MTU Training Centre Friedrichshafen',
    3,
    'Kualifikasi pemeliharaan depot level overhaul mesin pendorong pokok MTU'
),
(
    '82000000-0000-0000-0000-000000000003',
    'CERT-TACTICOS-OPER',
    'Sertifikasi Operator TACTICOS Combat Management System',
    'SPECIALIZATION',
    'Thales Naval Academy Netherlands',
    5,
    'Kualifikasi perwira pengarah taktis penembakan sensor senjata terpadu'
)
ON CONFLICT (qualification_id) DO NOTHING;

-- 6. Kualifikasi Prajurit
INSERT INTO hcm_personnel_qualifications (
    personnel_qual_id, personnel_id, qualification_id, certificate_number,
    obtained_date, valid_until, is_active
) VALUES
(
    '82500000-0000-0000-0000-000000000001',
    '81000000-0000-0000-0000-000000000003', -- Dan KRI
    '82000000-0000-0000-0000-000000000001', -- Brevet Komando
    'SK-BRV/2024/041',
    '2024-04-10',
    '2034-04-10',
    TRUE
),
(
    '82500000-0000-0000-0000-000000000002',
    '81000000-0000-0000-0000-000000000004', -- Kadepsin
    '82000000-0000-0000-0000-000000000002', -- MTU Cert
    'MTU-DE-2023-9912',
    '2023-09-15',
    '2026-09-15',
    TRUE
)
ON CONFLICT (personnel_qual_id) DO NOTHING;

-- 7. Kesiapan Medis (URIKKES)
INSERT INTO hcm_medical_readiness (
    medical_id, personnel_id, examination_date, stakes_category,
    physical_fitness_score, vision_status, dental_status, cardio_status,
    general_health_status, doctor_remarks, valid_until
) VALUES
(
    '83000000-0000-0000-0000-000000000001',
    '81000000-0000-0000-0000-000000000003', -- Dan KRI
    '2026-01-05',
    'STAKES_I',
    88.50,
    'NORMAL',
    'FIT',
    'FIT',
    'Sangat Prima (Siap Tempur Operasi Penuh)',
    'Kondisi fisik dan mental siap berlayar dan memimpin operasi tempur laut.',
    '2027-01-05'
),
(
    '83000000-0000-0000-0000-000000000004',
    '81000000-0000-0000-0000-000000000004', -- Kadepsin
    '2026-01-06',
    'STAKES_I',
    85.00,
    'NORMAL',
    'FIT',
    'FIT',
    'Sangat Prima',
    'Siap tugas dinas jaga mesin laut lepas.',
    '2027-01-06'
)
ON CONFLICT (medical_id) DO NOTHING;

-- 8. Manifest Awak Kapal Perang (Crew Assignment KRI REM-331)
INSERT INTO hcm_crew_assignments (
    assignment_id, ship_id, personnel_id, crew_role, department,
    watch_bill_duty, assigned_date, is_active
) VALUES
(
    '83500000-0000-0000-0000-000000000001',
    '41000000-0000-0000-0000-000000000001', -- KRI REM-331
    '81000000-0000-0000-0000-000000000003', -- Letkol Ahmad Dahlan
    'KOMANDAN KRI',
    'DEPOPS',
    'COMMAND_POST',
    '2024-05-15',
    TRUE
),
(
    '83500000-0000-0000-0000-000000000002',
    '41000000-0000-0000-0000-000000000001', -- KRI REM-331
    '81000000-0000-0000-0000-000000000004', -- Mayor Tri Wibowo
    'KEPALA DEPARTEMEN MESIN (KADEPSIN)',
    'DEPSIN',
    'ENGINEERING_CONTROL_ROOM',
    '2024-06-01',
    TRUE
),
(
    '83500000-0000-0000-0000-000000000003',
    '41000000-0000-0000-0000-000000000001', -- KRI REM-331
    '81000000-0000-0000-0000-000000000006', -- Serka Didik Supriyadi
    'BINTARA MESIN POKOK',
    'DEPSIN',
    'VIGOUR_A',
    '2024-06-01',
    TRUE
)
ON CONFLICT (assignment_id) DO NOTHING;

-- 9. Uang Layar / Tunjangan Operasi
INSERT INTO hcm_sea_duty_allowances (
    allowance_id, personnel_id, ship_id, mission_name, start_date,
    end_date, days_at_sea, daily_allowance_rate, payment_status, payment_reference_no
) VALUES
(
    '84000000-0000-0000-0000-000000000001',
    '81000000-0000-0000-0000-000000000003', -- Dan KRI
    '41000000-0000-0000-0000-000000000001',
    'Operasi Siaga Tempur Perbatasan Laut Natuna Utara 2026',
    '2026-03-01',
    '2026-03-15',
    15,
    350000.00,
    'PAID',
    'TR-LAYAR-2026-0041'
),
(
    '84000000-0000-0000-0000-000000000002',
    '81000000-0000-0000-0000-000000000004', -- Kadepsin
    '41000000-0000-0000-0000-000000000001',
    'Operasi Siaga Tempur Perbatasan Laut Natuna Utara 2026',
    '2026-03-01',
    '2026-03-15',
    15,
    300000.00,
    'PAID',
    'TR-LAYAR-2026-0041'
)
ON CONFLICT (allowance_id) DO NOTHING;


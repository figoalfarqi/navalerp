-- =============================================================================
-- MODUL 1: ORGANISASI, PENGGUNA & COMMON MASTER DATA (NAVALERP)
-- FILE: 01_organization_user/insert.sql
-- =============================================================================

-- 1. Insert Satuan Organisasi TNI AL
INSERT INTO org_units (unit_id, parent_unit_id, unit_code, unit_name, unit_type, description, command_level, latitude, longitude, address, phone, is_active, created_at)
VALUES
    ('10000000-0000-0000-0000-000000000001', NULL, 'MABESAL', 'Markas Besar TNI Angkatan Laut', 'HEADQUARTERS', 'Pusat Komando Strategis TNI AL Cilangkap', 1, -6.3195000, 106.9032000, 'Jl. Raya Hankam Cilangkap, Jakarta Timur', '+62218720100', TRUE, CURRENT_TIMESTAMP),
    ('10000000-0000-0000-0000-000000000002', '10000000-0000-0000-0000-000000000001', 'KOARMADA1', 'Komando Armada I (Tanjung Pinang)', 'FLEET', 'Komando Operasional Wilayah Laut Barat', 2, 0.9167000, 104.4500000, 'Jl. Yos Sudarso No. 1, Tanjung Pinang, Kepri', '+6277121234', TRUE, CURRENT_TIMESTAMP),
    ('10000000-0000-0000-0000-000000000003', '10000000-0000-0000-0000-000000000001', 'KOARMADA2', 'Komando Armada II (Surabaya)', 'FLEET', 'Komando Operasional Wilayah Laut Tengah & Pangkalan Utama KRI', 2, -7.2023000, 112.7410000, 'Dermaga Ujung, Semampir, Surabaya', '+62313291001', TRUE, CURRENT_TIMESTAMP),
    ('10000000-0000-0000-0000-000000000004', '10000000-0000-0000-0000-000000000001', 'KOARMADA3', 'Komando Armada III (Sorong)', 'FLEET', 'Komando Operasional Wilayah Laut Timur', 2, -0.8950000, 131.2550000, 'Klamono Km 16, Sorong, Papua Barat Daya', '+62951321000', TRUE, CURRENT_TIMESTAMP),
    ('10000000-0000-0000-0000-000000000005', '10000000-0000-0000-0000-000000000001', 'KOLINLAMIL', 'Komando Lintas Laut Militer (Jakarta)', 'FLEET', 'Komando Angkutan Laut Militer & Proyeksi Kekuatan Amfibi', 2, -6.1165000, 106.8833000, 'Pelabuhan Tanjung Priok, Jakarta Utara', '+62214301001', TRUE, CURRENT_TIMESTAMP),
    ('10000000-0000-0000-0000-000000000006', '10000000-0000-0000-0000-000000000003', 'LANTAMAL5', 'Pangkalan Utama TNI AL V (Surabaya)', 'LANTAMAL', 'Penyedia Dukungan Logistik Pangkalan Koarmada II', 3, -7.2150000, 112.7350000, 'Jl. Laksda M. Nasir No. 56, Surabaya', '+62313293005', TRUE, CURRENT_TIMESTAMP),
    ('10000000-0000-0000-0000-000000000007', '10000000-0000-0000-0000-000000000003', 'SATKOR_ARMADA2', 'Satuan Kapal Eskorta Koarmada II', 'SQUADRON', 'Satuan Pembina Kapal Frigate dan Corvette', 4, -7.2050000, 112.7420000, 'Dermaga Madura Ujung, Surabaya', '+62313292002', TRUE, CURRENT_TIMESTAMP),
    ('10000000-0000-0000-0000-000000000008', '10000000-0000-0000-0000-000000000003', 'SATKAT_ARMADA2', 'Satuan Kapal Cepat Koarmada II', 'SQUADRON', 'Satuan Pembina Kapal Cepat Rudal (KCR 60M)', 4, -7.2040000, 112.7430000, 'Dermaga Semampir, Surabaya', '+62313292004', TRUE, CURRENT_TIMESTAMP),
    ('10000000-0000-0000-0000-000000000009', '10000000-0000-0000-0000-000000000006', 'FASHARKAN_SBY', 'Fasilitas Pemeliharaan & Perbaikan Fasharkan Surabaya', 'FASHARKAN', 'Unit Pelaksana Teknis Pemeliharaan & Docking KRI', 4, -7.2080000, 112.7440000, 'Kawasan Dok Fasharkan Ujung, Surabaya', '+62313294001', TRUE, CURRENT_TIMESTAMP),
    ('10000000-0000-0000-0000-000000000010', '10000000-0000-0000-0000-000000000002', 'LANTAMAL4', 'Pangkalan Utama TNI AL IV (Batam)', 'LANTAMAL', 'Dukungan Operasi Wilayah Selat Malaka & Natuna', 3, 1.1300000, 104.0500000, 'Tanjung Sengkuang, Batu Ampar, Batam', '+62778412345', TRUE, CURRENT_TIMESTAMP)
ON CONFLICT (unit_id) DO UPDATE SET
    unit_name = EXCLUDED.unit_name,
    unit_code = EXCLUDED.unit_code,
    description = EXCLUDED.description,
    is_active = TRUE;

-- 2. Insert Pengguna Sistem dengan Role, Rank, dan Department
-- Password untuk SELURUH pengguna adalah: Password123!
-- Hash Bcrypt: $2a$10$7xMvtLLBCyffehFn5sKAmetDVxApVcMnudS9eLJgi4/8xour8YVFe
INSERT INTO sys_users (user_id, unit_id, username, password_hash, full_name, email, phone, military_id, rank_title, department, role, is_active, created_at)
VALUES
    ('20000000-0000-0000-0000-000000000001', '10000000-0000-0000-0000-000000000001', 'admin', '$2a$10$7xMvtLLBCyffehFn5sKAmetDVxApVcMnudS9eLJgi4/8xour8YVFe', 'Administrator Sistem NavalERP', 'admin@navalerp.tni.mil.id', '+62811000001', 'NRP-SYS-001', 'MAYOR_LAUT', 'SRENA', 'SUPER_ADMIN', TRUE, CURRENT_TIMESTAMP),
    ('20000000-0000-0000-0000-000000000002', '10000000-0000-0000-0000-000000000001', 'panglima', '$2a$10$7xMvtLLBCyffehFn5sKAmetDVxApVcMnudS9eLJgi4/8xour8YVFe', 'Laksamana Pertama TNI Judijanto', 'judijanto@navalerp.tni.mil.id', '+62811000002', 'NRP-987654', 'LAKSAMANA_PERTAMA_TNI', 'KOMANDO', 'COMMAND_OFFICER', TRUE, CURRENT_TIMESTAMP),
    ('20000000-0000-0000-0000-000000000003', '10000000-0000-0000-0000-000000000007', 'komandan.rem331', '$2a$10$7xMvtLLBCyffehFn5sKAmetDVxApVcMnudS9eLJgi4/8xour8YVFe', 'Kolonel Laut (P) Hendra Kurniawan', 'komandan.rem331@navalerp.tni.mil.id', '+62811000003', 'NRP-876543', 'KOLONEL_LAUT', 'DEPOPS', 'KRI_COMMANDER', TRUE, CURRENT_TIMESTAMP),
    ('20000000-0000-0000-0000-000000000004', '10000000-0000-0000-0000-000000000006', 'perwira.logistik', '$2a$10$7xMvtLLBCyffehFn5sKAmetDVxApVcMnudS9eLJgi4/8xour8YVFe', 'Letkol Laut (S) Bambang Prasetyo', 'logistik.lantamal5@navalerp.tni.mil.id', '+62811000004', 'NRP-765432', 'LETKOL_LAUT', 'SLOG', 'LOGISTICS_OFFICER', TRUE, CURRENT_TIMESTAMP),
    ('20000000-0000-0000-0000-000000000005', '10000000-0000-0000-0000-000000000007', 'kadepsin.rem331', '$2a$10$7xMvtLLBCyffehFn5sKAmetDVxApVcMnudS9eLJgi4/8xour8YVFe', 'Mayor Laut (T) Arif Wijaya', 'kadepsin.rem331@navalerp.tni.mil.id', '+62811000005', 'NRP-654321', 'MAYOR_LAUT', 'DEPSIN', 'MAINTENANCE_OFFICER', TRUE, CURRENT_TIMESTAMP),
    ('20000000-0000-0000-0000-000000000006', '10000000-0000-0000-0000-000000000001', 'perwira.keuangan', '$2a$10$7xMvtLLBCyffehFn5sKAmetDVxApVcMnudS9eLJgi4/8xour8YVFe', 'Letkol Laut (S) Deni Mulyadi', 'keuangan.mabesal@navalerp.tni.mil.id', '+62811000006', 'NRP-754312', 'LETKOL_LAUT', 'SRENA', 'FINANCE_OFFICER', TRUE, CURRENT_TIMESTAMP),
    ('20000000-0000-0000-0000-000000000007', '10000000-0000-0000-0000-000000000001', 'perwira.personel', '$2a$10$7xMvtLLBCyffehFn5sKAmetDVxApVcMnudS9eLJgi4/8xour8YVFe', 'Kolonel Laut (E) Agus Santoso', 'spers.mabesal@navalerp.tni.mil.id', '+62811000007', 'NRP-843219', 'KOLONEL_LAUT', 'SPERS', 'PERSONNEL_OFFICER', TRUE, CURRENT_TIMESTAMP),
    ('20000000-0000-0000-0000-000000000008', '10000000-0000-0000-0000-000000000007', 'komandan.inr332', '$2a$10$7xMvtLLBCyffehFn5sKAmetDVxApVcMnudS9eLJgi4/8xour8YVFe', 'Kolonel Laut (P) Faisal Anwar', 'komandan.inr332@navalerp.tni.mil.id', '+62811000008', 'NRP-890123', 'KOLONEL_LAUT', 'DEPOPS', 'KRI_COMMANDER', TRUE, CURRENT_TIMESTAMP),
    ('20000000-0000-0000-0000-000000000009', '10000000-0000-0000-0000-000000000007', 'operator.radar', '$2a$10$7xMvtLLBCyffehFn5sKAmetDVxApVcMnudS9eLJgi4/8xour8YVFe', 'Serka Nav Dwi Prasetyo', 'operator.radar@navalerp.tni.mil.id', '+62811000009', 'NRP-102938', 'SERKA', 'DEPOPS', 'OPERATOR', TRUE, CURRENT_TIMESTAMP),
    ('20000000-0000-0000-0000-000000000010', '10000000-0000-0000-0000-000000000006', 'operator.gudang', '$2a$10$7xMvtLLBCyffehFn5sKAmetDVxApVcMnudS9eLJgi4/8xour8YVFe', 'Sertu Bek Rahmat Hidayat', 'operator.gudang@navalerp.tni.mil.id', '+62811000010', 'NRP-104829', 'SERTU', 'DISBEKAL', 'OPERATOR', TRUE, CURRENT_TIMESTAMP)
ON CONFLICT (user_id) DO UPDATE SET
    password_hash = EXCLUDED.password_hash,
    full_name = EXCLUDED.full_name,
    email = EXCLUDED.email,
    phone = EXCLUDED.phone,
    military_id = EXCLUDED.military_id,
    rank_title = EXCLUDED.rank_title,
    department = EXCLUDED.department,
    role = EXCLUDED.role,
    is_active = TRUE;

-- Update auth_version dan pastikan seluruh user memiliki hash Password123!
UPDATE sys_users SET 
    password_hash = '$2a$10$7xMvtLLBCyffehFn5sKAmetDVxApVcMnudS9eLJgi4/8xour8YVFe',
    is_active = TRUE;

-- 3. Audit Log Awal
INSERT INTO sys_audit_logs (log_id, user_id, action, entity_table, entity_id, new_values, ip_address, created_at)
VALUES
    ('21000000-0000-0000-0000-000000000001', '20000000-0000-0000-0000-000000000001', 'INITIALIZE', 'org_units', '10000000-0000-0000-0000-000000000001', '{"status": "System Initialized", "version": "2.0", "security": "Bcrypt Password123!"}', '127.0.0.1', CURRENT_TIMESTAMP)
ON CONFLICT (log_id) DO NOTHING;

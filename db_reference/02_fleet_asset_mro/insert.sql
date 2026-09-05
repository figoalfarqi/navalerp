-- =============================================================================
-- MODUL 2: FLEET, ASSET & HIERARKI PERALATAN KAPAL (MRO & ASSET) (NAVALERP)
-- FILE: 02_fleet_asset_mro/insert.sql
-- =============================================================================

-- 1. Insert Master Kelas Kapal
INSERT INTO mro_ship_classes (class_id, class_code, class_name, category, specifications, builder, total_built, created_at)
VALUES
    ('40000000-0000-0000-0000-000000000001', 'SIGMA_10514', 'Martadinata Class Guided-Missile Frigate', 'FRIGATE', '{"length_m": 105.14, "beam_m": 14.02, "draft_m": 3.75, "displacement_tons": 2365, "speed_knots": 28, "range_nm": 5000, "endurance_days": 20}', 'Damen Schelde Naval Shipbuilding / PT PAL Indonesia', 2, CURRENT_TIMESTAMP),
    ('40000000-0000-0000-0000-000000000002', 'SIGMA_9113', 'Diponegoro Class Guided-Missile Corvette', 'CORVETTE', '{"length_m": 90.71, "beam_m": 13.02, "draft_m": 3.60, "displacement_tons": 1692, "speed_knots": 28, "range_nm": 4000, "endurance_days": 14}', 'Damen Schelde Naval Shipbuilding', 4, CURRENT_TIMESTAMP),
    ('40000000-0000-0000-0000-000000000003', 'MAKASSAR_LPD', 'Makassar Class Landing Platform Dock', 'LPD', '{"length_m": 122.00, "beam_m": 22.00, "draft_m": 4.90, "displacement_tons": 7300, "speed_knots": 16, "range_nm": 10000, "endurance_days": 30}', 'Daesun Shipbuilding / PT PAL Indonesia', 5, CURRENT_TIMESTAMP)
ON CONFLICT (class_id) DO NOTHING;

-- 2. Insert Master Kapal (KRI)
INSERT INTO mro_ships (ship_id, class_id, assigned_unit_id, hull_number, ship_name, call_sign, commission_date, home_port, length_m, beam_m, draft_m, displacement_tons, max_speed_knots, cruise_range_nm, crew_capacity, fuel_capacity_liters, fresh_water_capacity_liters, status, current_readiness_status, created_at)
VALUES
    ('41000000-0000-0000-0000-000000000001', '40000000-0000-0000-0000-000000000001', '10000000-0000-0000-0000-000000000007', '331', 'KRI Raden Eddy Martadinata', '7EAA', '2017-04-07', 'Pangkalan Surabaya (Koarmada II)', 105.14, 14.02, 3.75, 2365.00, 28.00, 5000.00, 122, 280000.00, 45000.00, 'ACTIVE', 'FULLY_MISSION_CAPABLE', CURRENT_TIMESTAMP),
    ('41000000-0000-0000-0000-000000000002', '40000000-0000-0000-0000-000000000001', '10000000-0000-0000-0000-000000000007', '332', 'KRI I Gusti Ngurah Rai', '7EAB', '2018-01-10', 'Pangkalan Surabaya (Koarmada II)', 105.14, 14.02, 3.75, 2365.00, 28.00, 5000.00, 122, 280000.00, 45000.00, 'ACTIVE', 'FULLY_MISSION_CAPABLE', CURRENT_TIMESTAMP),
    ('41000000-0000-0000-0000-000000000003', '40000000-0000-0000-0000-000000000002', '10000000-0000-0000-0000-000000000007', '365', 'KRI Diponegoro', '7EAC', '2007-07-02', 'Pangkalan Surabaya (Koarmada II)', 90.71, 13.02, 3.60, 1692.00, 28.00, 4000.00, 80, 190000.00, 30000.00, 'ACTIVE', 'FULLY_MISSION_CAPABLE', CURRENT_TIMESTAMP),
    ('41000000-0000-0000-0000-000000000004', '40000000-0000-0000-0000-000000000003', '10000000-0000-0000-0000-000000000005', '590', 'KRI Makassar', '7EAD', '2007-04-29', 'Pangkalan Jakarta (Kolinlamil)', 122.00, 22.00, 4.90, 7300.00, 16.00, 10000.00, 518, 650000.00, 120000.00, 'ACTIVE', 'FULLY_MISSION_CAPABLE', CURRENT_TIMESTAMP)
ON CONFLICT (ship_id) DO NOTHING;

-- 3. Insert Hierarki Sistem Kapal (KRI REM-331)
INSERT INTO mro_systems (system_id, ship_id, parent_system_id, system_code, system_name, system_category, system_level, description, created_at)
VALUES
    ('42000000-0000-0000-0000-000000000001', '41000000-0000-0000-0000-000000000001', NULL, 'PROPULSION', 'Sistem Propulsi Gabungan Diesel & Elektrik (CODOE)', 'PROPULSION', 'SYSTEM', '2x MTU 20V 4000 M93L Diesel + 2x Electric Motors', CURRENT_TIMESTAMP),
    ('42000000-0000-0000-0000-000000000002', '41000000-0000-0000-0000-000000000001', NULL, 'ELECTRICAL', 'Sistem Pembangkit & Distribusi Daya Listrik', 'ELECTRICAL', 'SYSTEM', '4x Caterpillar 3412C Diesel Generators', CURRENT_TIMESTAMP),
    ('42000000-0000-0000-0000-000000000003', '41000000-0000-0000-0000-000000000001', NULL, 'RADAR_SENSOR', 'Sistem Sensor, Radar Pengintai & Sonar', 'SENSOR_RADAR', 'SYSTEM', 'Thales SMART-S Mk2, STIR 1.2 EO Mk2, Kingklip Sonar', CURRENT_TIMESTAMP),
    ('42000000-0000-0000-0000-000000000004', '41000000-0000-0000-0000-000000000001', NULL, 'WEAPON_SYSTEM', 'Sistem Artileri & Peluncur Rudal Pertahanan Tempur', 'WEAPON', 'SYSTEM', 'Oto Melara 76mm, VL MICA SAM, Exocet MM40 Block 3', CURRENT_TIMESTAMP)
ON CONFLICT (system_id) DO NOTHING;

-- 4. Insert Equipment Terpasang
INSERT INTO mro_equipments (equipment_id, system_id, serial_number, equipment_tag, equipment_name, manufacturer, model_number, country_of_origin, installation_date, total_operating_hours, design_life_hours, criticality_level, health_status, created_at)
VALUES
    ('43000000-0000-0000-0000-000000000001', '42000000-0000-0000-0000-000000000001', 'MTU-20V-4000-01', 'ME-STBD', 'Main Engine Diesel Kanan (Starboard)', 'MTU Friedrichshafen', '20V 4000 M93L', 'Germany', '2016-08-15', 3450.50, 30000.00, 'CRITICAL_SAFETY', 'OPERATIONAL', CURRENT_TIMESTAMP),
    ('43000000-0000-0000-0000-000000000002', '42000000-0000-0000-0000-000000000002', 'CAT-3412-GEN-01', 'DG-01', 'Diesel Generator Utama No 1', 'Caterpillar Marine', 'CAT 3412C', 'USA', '2016-08-15', 5210.00, 40000.00, 'CRITICAL_SAFETY', 'OPERATIONAL', CURRENT_TIMESTAMP),
    ('43000000-0000-0000-0000-000000000003', '42000000-0000-0000-0000-000000000003', 'THALES-SMARTS-01', 'RADAR-3D', '3D Multi-Beam Air & Surface Surveillance Radar', 'Thales Nederland', 'SMART-S Mk2', 'Netherlands', '2016-10-20', 2890.00, 25000.00, 'MISSION_ESSENTIAL', 'OPERATIONAL', CURRENT_TIMESTAMP),
    ('43000000-0000-0000-0000-000000000004', '42000000-0000-0000-0000-000000000004', 'OTO-76SR-01', 'GUN-MAIN-76', 'Meriam Utama 76mm Super Rapid Gun', 'Leonardo / Oto Melara', '76/62 SR', 'Italy', '2016-11-05', 420.00, 15000.00, 'MISSION_ESSENTIAL', 'OPERATIONAL', CURRENT_TIMESTAMP)
ON CONFLICT (equipment_id) DO NOTHING;

-- 5. Insert Telemetri Kondisi Mesin
INSERT INTO mro_equipment_parameters (param_id, equipment_id, recorded_at, rpm, temperature_celsius, pressure_bar, vibration_level, oil_pressure_bar, running_hours_snapshot, status_flag, recorded_by_user_id)
VALUES
    ('43100000-0000-0000-0000-000000000001', '43000000-0000-0000-0000-000000000001', CURRENT_TIMESTAMP, 1800.00, 84.50, 4.80, 0.045, 5.20, 3450.50, 'NORMAL', '20000000-0000-0000-0000-000000000005')
ON CONFLICT (param_id) DO NOTHING;

-- 6. Insert Jadwal Pemeliharaan Preventif (PMS)
INSERT INTO mro_pm_schedules (pm_id, equipment_id, pm_code, pm_title, interval_hours, interval_days, last_performed_at, next_due_at, task_instructions, estimated_duration_hours, is_active, created_at)
VALUES
    ('43200000-0000-0000-0000-000000000001', '43000000-0000-0000-0000-000000000001', 'PMS-MTU-500H', 'Inspeksi Berkala & Penggantian Filter Pelumas 500 Jam', 500, 90, '2025-11-10', '2026-03-15', 'Ganti filter oli, periksa tekanan bahan bakar, cek celah katup silinder', 6.00, TRUE, CURRENT_TIMESTAMP)
ON CONFLICT (pm_id) DO NOTHING;

-- 7. Insert Laporan Kerusakan (Failure Report)
INSERT INTO mro_failure_reports (report_id, equipment_id, reported_by_user_id, report_number, incident_date, severity, failure_mode, description, operational_impact, immediate_action_taken, status, created_at)
VALUES
    ('44000000-0000-0000-0000-000000000001', '43000000-0000-0000-0000-000000000001', '20000000-0000-0000-0000-000000000005', 'FR-REM331-2026-001', '2026-02-10 08:30:00+07', 'CAT2', 'Fuel Injector Leakage', 'Fluktuasi tekanan bahan bakar silinder 4 pada Main Engine Starboard MTU 20V 4000', 'Kecepatan maksimum kapal terbatasi 18 knot', 'Pengurangan RPM mesin kanan dan isolasi saluran injektor silinder 4', 'WORK_ORDER_CREATED', CURRENT_TIMESTAMP)
ON CONFLICT (report_id) DO NOTHING;

-- 8. Insert Perintah Kerja Pemeliharaan (Work Order)
INSERT INTO mro_work_orders (work_order_id, failure_report_id, pm_schedule_id, equipment_id, work_order_number, work_order_type, priority, scheduled_start_date, scheduled_end_date, actual_start_date, actual_end_date, lead_engineer_user_id, assigned_facility, status, total_labor_hours, estimated_cost, actual_cost, completion_notes, created_at)
VALUES
    ('45000000-0000-0000-0000-000000000001', '44000000-0000-0000-0000-000000000001', NULL, '43000000-0000-0000-0000-000000000001', 'WO-REM331-2026-008', 'CORRECTIVE', 'URGENT', '2026-02-12', '2026-02-15', '2026-02-12', '2026-02-14', '20000000-0000-0000-0000-000000000005', 'Dermaga Madura Koarmada II Surabaya', 'COMPLETED', 24.50, 50000000.00, 45000000.00, 'Penggantian fuel injector silinder 4 selesai. Uji harbour acceptance trial dan sea acceptance trial 24 knot normal.', CURRENT_TIMESTAMP)
ON CONFLICT (work_order_id) DO NOTHING;

-- 9. Insert Rincian Langkah Pemeliharaan
INSERT INTO mro_work_order_tasks (task_id, work_order_id, step_number, task_description, estimated_minutes, actual_minutes, is_completed, completed_by_user_id, notes, created_at)
VALUES
    ('45100000-0000-0000-0000-000000000001', '45000000-0000-0000-0000-000000000001', 1, 'Shut-down dan isolasi sistem bahan bakar Main Engine Starboard', 60, 45, TRUE, '20000000-0000-0000-0000-000000000005', 'Selesai sesuai SOP K3 Fasharkan', CURRENT_TIMESTAMP),
    ('45100000-0000-0000-0000-000000000002', '45000000-0000-0000-0000-000000000001', 2, 'Pelepasan injector assembly silinder 4 dan kalibrasi tekanan pembukaan nozzle', 180, 160, TRUE, '20000000-0000-0000-0000-000000000005', 'Injektor lama aus, diganti kit baru', CURRENT_TIMESTAMP),
    ('45100000-0000-0000-0000-000000000003', '45000000-0000-0000-0000-000000000001', 3, 'Pemasangan kit baru, uji kebocoran tekanan tinggi dan pengetesan beban', 120, 110, TRUE, '20000000-0000-0000-0000-000000000005', 'Hasil pengujian tekanan 1600 bar stabil', CURRENT_TIMESTAMP)
ON CONFLICT (task_id) DO NOTHING;

-- 10. Insert Konsumsi Suku Cadang WO
INSERT INTO mro_work_order_items (wo_item_id, work_order_id, material_id, quantity_required, quantity_issued, unit_cost, total_cost, is_critical_spare, created_at)
VALUES
    ('46000000-0000-0000-0000-000000000001', '45000000-0000-0000-0000-000000000001', '51000000-0000-0000-0000-000000000001', 2.00, 2.00, 22500000.00, 45000000.00, TRUE, CURRENT_TIMESTAMP)
ON CONFLICT (wo_item_id) DO NOTHING;

-- 11. Insert Catatan Docking Galangan
INSERT INTO mro_docking_records (docking_id, ship_id, shipyard_name, docking_type, entry_date, scheduled_exit_date, actual_exit_date, sea_trial_passed, classification_surveyor, certificate_number, total_docking_cost, docking_summary, created_at)
VALUES
    ('47000000-0000-0000-0000-000000000001', '41000000-0000-0000-0000-000000000001', 'PT PAL Indonesia (Persero) Surabaya', 'ANNUAL_DOCKING', '2025-06-01', '2025-07-15', '2025-07-12', TRUE, 'Biro Klasifikasi Indonesia (BKI)', 'BKI-NAV-2025-081', 12500000000.00, 'Docking tahunan, pembersihan lambung bawah air, penggantian sacrificial anode zink, inspeksi shaft propeller dan sea valve.', CURRENT_TIMESTAMP)
ON CONFLICT (docking_id) DO NOTHING;


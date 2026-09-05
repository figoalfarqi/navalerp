-- Seed data for apps/api/create_tables.sql.
-- Run this file only after create_tables.sql on an empty database.

SET TIMEZONE = 'Asia/Jakarta';

CREATE EXTENSION IF NOT EXISTS pgcrypto;

-- app_role_type_id
-- 1 driver
-- 2 checker
-- 3 management and system administration

-- 1. Roles, banks, and users
INSERT INTO app_role (
    app_role_id, app_role_type_id, app_role_name, app_role_description,
    is_active, created_by, updated_by
) VALUES
    (1, 1, 'driver',
        'Melihat rute, perjalanan, berat, dan volume miliknya secara read-only.',
        1, 1, 1),
    (2, 2, 'checker',
        'Melaporkan status perjalanan, berat atau volume, dan bukti truck untuk project yang ditugaskan.',
        1, 1, 1),
    (3, 3, 'owner',
        'Memiliki akses penuh ke seluruh fitur admin, data operasional, laporan, dan manajemen pengguna.',
        1, 1, 1),
    (4, 3, 'itdev',
        'Memiliki akses penuh ke seluruh fitur admin, data, konfigurasi, dan manajemen pengguna.',
        1, 1, 1),
    (5, 3, 'superadmin',
        'Memiliki akses penuh ke seluruh fitur, data, konfigurasi, dan manajemen pengguna.',
        1, 1, 1),
    (6, 3, 'admin',
        'Mengelola data operasional harian, master data, dan penugasan project.',
        1, 1, 1);

INSERT INTO bank_merk (
    bank_merk_id, bank_merk_name, bank_merk_description,
    is_active, created_by, updated_by
) VALUES
    (1, 'BCA', 'Bank Central Asia', 1, 1, 1),
    (2, 'Mandiri', 'Bank Mandiri', 1, 1, 1);

-- created_by and updated_by are NULL for the first user because this user
-- bootstraps the audit-user hierarchy.
INSERT INTO app_user (
    app_user_id, app_role_id, bank_merk_id, username, password,
    app_user_status_id, app_user_name, app_user_preferred_name,
    app_user_phone, city_id, app_user_address, bank_account_number,
    bank_account_name, salary_percentage, created_by, updated_by
) VALUES
    (1, 5, 1, 'superadmin',
        crypt('Password123!', gen_salt('bf', 10)),
        1, 'Super Administrator', 'Super Admin', '081100000001', NULL,
        'Kantor Pusat', '000111222333', 'Super Administrator', NULL, NULL, NULL);

-- System settings are maintained separately from project transactions.
INSERT INTO app_setting (
    app_setting_id, app_setting_key, app_setting_value,
    app_setting_description, is_active, created_by, updated_by
) VALUES
    (1, 'operations.timezone', '{"iana":"Asia/Jakarta"}'::jsonb,
        'Zona waktu operasional dan pelaporan.', 1, 1, 1),
    (2, 'checker.location_cooldown_minutes', '{"minutes":5}'::jsonb,
        'Batas perubahan lokasi checker pada aplikasi mobile.', 1, 1, 1),
    (3, 'checker.truck_cooldown_minutes', '{"minutes":10}'::jsonb,
        'Batas input truck yang sama oleh checker pada aplikasi mobile.', 1, 1, 1),
    (4, 'reporting.currency', '{"code":"IDR","locale":"id-ID"}'::jsonb,
        'Mata uang standar untuk dashboard dan laporan.', 1, 1, 1);

-- 2. Geographic and operational master data
INSERT INTO province (
    province_id, province_name, province_real_name, is_active, created_by, updated_by
) VALUES
    (1, 'Kalimantan Selatan', 'Kalimantan Selatan', 1, 1, 1),
    (2, 'DKI Jakarta', 'Daerah Khusus Ibukota Jakarta', 1, 1, 1),
    (3, 'Jawa Barat', 'Jawa Barat', 1, 1, 1);

INSERT INTO city (
    city_id, province_id, city_name, is_active, created_by, updated_by
) VALUES
    (1, 1, 'Banjarmasin', 1, 1, 1),
    (2, 2, 'Jakarta Utara', 1, 1, 1),
    (3, 3, 'Bekasi', 1, 1, 1);

UPDATE app_user
SET city_id = 1, updated_by = 1
WHERE app_user_id = 1;

INSERT INTO app_user (
    app_user_id, app_role_id, bank_merk_id, username, password,
    app_user_status_id, app_user_name, app_user_preferred_name,
    app_user_phone, city_id, app_user_address, bank_account_number,
    bank_account_name, salary_percentage, created_by, updated_by
) VALUES
    (2, 6, 1, 'admin.operasional',
        crypt('Password123!', gen_salt('bf', 10)),
        1, 'Rina Administrator Operasional', 'Rina', '081100000002', 1,
        'Banjarmasin', '000111222334', 'Rina Administrator', NULL, 1, 1),
    (3, 2, 1, 'checker.a',
        crypt('Password123!', gen_salt('bf', 10)),
        1, 'Andi Checker', 'Checker A', '081100000003', 1,
        'Banjarmasin', '000111222335', 'Andi Checker', NULL, 1, 1),
    (4, 2, 2, 'checker.b',
        crypt('Password123!', gen_salt('bf', 10)),
        1, 'Budi Checker', 'Checker B', '081100000004', 2,
        'Jakarta Utara', '000111222336', 'Budi Checker', NULL, 1, 1),
    (5, 1, 1, 'driver.andi',
        crypt('Password123!', gen_salt('bf', 10)),
        1, 'Andi Pengemudi', 'Andi', '081100000005', 1,
        'Banjarmasin', '000111222337', 'Andi Pengemudi', 10.00, 1, 1),
    (6, 1, 1, 'driver.citra',
        crypt('Password123!', gen_salt('bf', 10)),
        1, 'Citra Pengemudi', 'Citra', '081100000006', 1,
        'Banjarmasin', '000111222338', 'Citra Pengemudi', 10.00, 1, 1),
    (7, 3, 2, 'owner',
        crypt('Password123!', gen_salt('bf', 10)),
        1, 'Owner Perusahaan', 'Owner', '081100000007', 2,
        'Jakarta Utara', '000111222339', 'Owner Perusahaan', NULL, 1, 1),
    (8, 1, 2, 'driver.dedi',
        crypt('Password123!', gen_salt('bf', 10)),
        1, 'Dedi Pengemudi', 'Dedi', '081100000008', 3,
        'Bekasi', '000111222340', 'Dedi Pengemudi', 10.00, 1, 1),
    (9, 4, 2, 'itdev',
        crypt('Password123!', gen_salt('bf', 10)),
        1, 'Developer Aplikasi', 'IT Dev', '081100000009', 2,
        'Jakarta Utara', '000111222341', 'Developer Aplikasi', NULL, 1, 1);

INSERT INTO vendor_type (
    vendor_type_id, vendor_type_name, vendor_type_description,
    is_active, created_by, updated_by
) VALUES
    (1, 'Transporter', 'Vendor jasa angkutan darat', 1, 1, 1),
    (2, 'Supplier Material', 'Pemasok material tambang', 1, 1, 1),
    (3, 'Shipping', 'Vendor kapal dan pengiriman laut', 1, 1, 1);

INSERT INTO vendor (
    vendor_id, vendor_type_id, bank_merk_id, vendor_name, vendor_email,
    vendor_phone, vendor_tin, city_id, vendor_address,
    bank_account_number, bank_account_name, is_active, created_by, updated_by
) VALUES
    (1, 1, 1, 'PT Angkut Nusantara', 'operasional@angkut.test',
        '0511000001', '01.111.222.3-444.000', 1, 'Banjarmasin',
        '1234567890', 'PT Angkut Nusantara', 1, 1, 1),
    (2, 2, 1, 'CV Tambang Sejahtera', 'sales@tambang.test',
        '0511000002', '01.111.222.3-445.000', 1, 'Banjar',
        '1234567891', 'CV Tambang Sejahtera', 1, 1, 1),
    (3, 3, 2, 'PT Samudra Logistik', 'charter@samudra.test',
        '0210000003', '01.111.222.3-446.000', 2, 'Jakarta Utara',
        '1234567892', 'PT Samudra Logistik', 1, 1, 1),
    (4, 1, 2, 'PT Mandiri Trucking', 'dispatch@mandiri.test',
        '0210000004', '01.111.222.3-447.000', 3, 'Bekasi',
        '1234567893', 'PT Mandiri Trucking', 1, 1, 1);

INSERT INTO truck_type (
    truck_type_id, truck_type_name, truck_type_description,
    truck_box_length, truck_box_width, truck_box_height, truck_capacity,
    is_active, created_by, updated_by
) VALUES
    (1, 'Dump Truck 6 Roda', 'Dump truck untuk material curah',
        6.00, 2.30, 1.80, 20.00, 1, 1, 1),
    (2, 'Trailer', 'Trailer untuk muatan besar',
        12.00, 2.50, 2.00, 30.00, 1, 1, 1);

INSERT INTO truck_merk (
    truck_merk_id, truck_merk_name, is_active, created_by, updated_by
) VALUES
    (1, 'Hino', 1, 1, 1),
    (2, 'Mitsubishi Fuso', 1, 1, 1);

INSERT INTO truck (
    truck_id, truck_type_id, truck_merk_id, driver_id, vendor_id,
    license_plate, ownership_status_id, production_year, number_of_tires,
    is_active, created_by, updated_by
) VALUES
    (1, 1, 1, 5, 1, 'DA 8101 AN', 2, 2022, 6, 1, 1, 1),
    (2, 1, 2, 6, 1, 'DA 8102 AN', 2, 2021, 6, 1, 1, 1),
    (3, 2, 1, 8, 4, 'B 9103 MD', 2, 2023, 10, 1, 1, 1);

-- 3. Clients, materials, sources, and stockpiles
INSERT INTO client (
    client_id, client_name, client_email, client_tin, number_of_day_until_due,
    city_id, client_address, operating_hours, is_active, created_by, updated_by
) VALUES
    (1, 'PT Beton Prima', 'procurement@betonprima.test',
        '02.111.222.3-444.000', 30, 3, 'Bekasi',
        '{"mon_fri":"08:00-17:00"}'::jsonb, 1, 1, 1),
    (2, 'PT Konstruksi Jakarta', 'procurement@konstruksi.test',
        '02.111.222.3-445.000', 21, 2, 'Jakarta Utara',
        '{"mon_sat":"08:00-17:00"}'::jsonb, 1, 1, 1);

INSERT INTO client_destination (
    client_destination_id, client_id, client_destination_name, city_id,
    client_destination_address, client_destination_latitude,
    client_destination_longitude, client_destination_map_url,
    operating_hours, is_active, created_by, updated_by
) VALUES
    (1, 1, 'Batching Plant Bekasi', 3, 'Kawasan Industri Bekasi',
        -6.238300, 107.001600, 'https://maps.example.test/bekasi',
        '{"mon_sat":"08:00-17:00"}'::jsonb, 1, 1, 1),
    (2, 2, 'Proyek Jakarta Utara', 2, 'Jakarta Utara',
        -6.121400, 106.893600, 'https://maps.example.test/jakarta-utara',
        '{"mon_sat":"08:00-17:00"}'::jsonb, 1, 1, 1);

INSERT INTO cargo_type (
    cargo_type_id, cargo_type_name, cargo_type_description, cargo_type_grade,
    is_active, created_by, updated_by
) VALUES
    (1, 'Batu Split', 'Material batu pecah untuk beton', 'A', 1, 1, 1),
    (2, 'Pasir Beton', 'Pasir untuk campuran beton', 'A', 1, 1, 1);

INSERT INTO mine (
    mine_id, mine_name, city_id, mine_address, mine_latitude,
    mine_longitude, mine_map_url, is_active, created_by, updated_by
) VALUES
    (1, 'Quarry Banjar', 1, 'Kabupaten Banjar', -3.320000, 114.600000,
        'https://maps.example.test/quarry-banjar', 1, 1, 1),
    (2, 'Quarry Martapura', 1, 'Martapura', -3.410000, 114.850000,
        'https://maps.example.test/quarry-martapura', 1, 1, 1);

INSERT INTO port (
    port_id, port_name, city_id, port_address, port_latitude,
    port_longitude, port_map_url, is_active, created_by, updated_by
) VALUES
    (1, 'Pelabuhan Trisakti', 1, 'Banjarmasin', -3.330000, 114.570000,
        'https://maps.example.test/trisakti', 1, 1, 1),
    (2, 'Pelabuhan Tanjung Priok', 2, 'Jakarta Utara', -6.104000, 106.880000,
        'https://maps.example.test/tanjung-priok', 1, 1, 1);

INSERT INTO vessel (
    vessel_id, vendor_id, vessel_name, imo_number, registration_number,
    is_active, created_by, updated_by
) VALUES
    (1, 3, 'KM Nusantara', 'IMO9000001', 'ID-KM-001', 1, 1, 1),
    (2, 3, 'MV Laut Jaya', 'IMO9000002', 'ID-MV-002', 1, 1, 1);

INSERT INTO vessel_cargo (
    vessel_cargo_id, vessel_id, port_id, cargo_type_id, voyage_number,
    bill_of_lading_number, arrival_at, unloading_started_at,
    unloading_completed_at, manifest_volume_cubic, manifest_weight_ton,
    is_active, created_by, updated_by
) VALUES
    (1, 1, 1, 1, 'VYG-NS-001', 'BL-NS-001',
        CURRENT_TIMESTAMP - INTERVAL '12 days',
        CURRENT_TIMESTAMP - INTERVAL '11 days',
        CURRENT_TIMESTAMP - INTERVAL '10 days', 100.000, 150.000, 1, 1, 1),
    (2, 2, 2, 2, 'VYG-LJ-002', 'BL-LJ-002',
        CURRENT_TIMESTAMP - INTERVAL '8 days',
        CURRENT_TIMESTAMP - INTERVAL '7 days',
        CURRENT_TIMESTAMP - INTERVAL '6 days', 80.000, 120.000, 1, 1, 1);

INSERT INTO stockpile (
    stockpile_id, stockpile_name, city_id, stockpile_address,
    stockpile_latitude, stockpile_longitude, stockpile_map_url,
    is_active, created_by, updated_by
) VALUES
    (1, 'Stockpile Banjarmasin', 1, 'Banjarmasin', -3.340000, 114.580000,
        'https://maps.example.test/stockpile-bjm', 1, 1, 1),
    (2, 'Stockpile Cakung', 2, 'Cakung, Jakarta Timur', -6.185000, 106.940000,
        'https://maps.example.test/stockpile-cakung', 1, 1, 1);

INSERT INTO stockpile_cargo (
    stockpile_cargo_id, stockpile_id, cargo_type_id,
    capacity_volume_cubic, capacity_weight_ton,
    current_volume_cubic, current_weight_ton,
    is_active, created_by, updated_by
) VALUES
    (1, 1, 1, 1000.000, 1500.000, 0.334, 0.500, 1, 1, 1),
    (2, 1, 2, 800.000, 1200.000, 0.000, 0.000, 1, 1, 1),
    (3, 2, 2, 600.000, 900.000, 0.500, 0.750, 1, 1, 1);

-- 4. Projects: one example for every supported route pattern
INSERT INTO project (
    project_id, project_code, project_name, route_type,
    mine_id, vessel_cargo_id, stockpile_cargo_id,
    client_destination_id, cargo_type_id, project_status,
    start_date, end_date, planned_volume_cubic, planned_weight_ton,
    volume_to_weight_conversion,
    material_purchase_unit, material_buy_price_per_cubic, material_buy_price_per_ton,
    material_sale_unit, material_sell_price_per_cubic, material_sell_price_per_ton,
    fixed_other_income, fixed_other_expense, project_note,
    is_active, created_by, updated_by
) VALUES
    (1, 'P-MC-001', 'Quarry Banjar ke Batching Plant Bekasi', 'MINE_CLIENT',
        1, NULL, NULL, 1, 1, 'ACTIVE',
        CURRENT_DATE - 30, NULL, 333.333, 500.000, 1.500000,
        'TON', NULL, 145000.00, 'TON', NULL, 210000.00,
        1000000.00, 500000.00, 'Pengiriman langsung dari tambang', 1, 2, 2),
    (2, 'P-MS-002', 'Quarry Banjar melalui Stockpile Banjarmasin', 'MINE_STOCKPILE_CLIENT',
        1, NULL, 1, 1, 1, 'ACTIVE',
        CURRENT_DATE - 25, NULL, 400.000, 600.000, 1.500000,
        'TON', NULL, 140000.00, 'TON', NULL, 225000.00,
        0.00, 250000.00, 'Dua ruas: tambang ke stockpile lalu client', 1, 2, 2),
    (3, 'P-VC-003', 'KM Nusantara langsung ke Proyek Jakarta Utara', 'VESSEL_CLIENT',
        NULL, 1, NULL, 2, 1, 'ACTIVE',
        CURRENT_DATE - 12, NULL, 100.000, 150.000, 1.500000,
        'M3', 90000.00, NULL, 'M3', 145000.00, NULL,
        0.00, 0.00, 'Material dari kapal langsung dikirim ke client', 1, 2, 2),
    (4, 'P-VS-004', 'MV Laut Jaya melalui Stockpile Cakung', 'VESSEL_STOCKPILE_CLIENT',
        NULL, 2, 3, 2, 2, 'ACTIVE',
        CURRENT_DATE - 8, NULL, 80.000, 120.000, 1.500000,
        'M3', 75000.00, NULL, 'M3', 130000.00, NULL,
        0.00, 150000.00, 'Dua ruas: kapal ke stockpile lalu client', 1, 2, 2);

-- Checker A can see projects 1 and 2; Checker B can see projects 3 and 4.
INSERT INTO project_checker_assignment (
    project_checker_assignment_id, project_id, checker_id,
    access_started_at, access_ended_at, assignment_note,
    is_active, is_default, created_by, updated_by
) VALUES
    (1, 1, 3, CURRENT_TIMESTAMP - INTERVAL '30 days', NULL,
        'Checker A - project direct mine to client', 1, 1, 2, 2),
    (2, 2, 3, CURRENT_TIMESTAMP - INTERVAL '25 days', NULL,
        'Checker A - project through stockpile', 1, 0, 2, 2),
    (3, 3, 4, CURRENT_TIMESTAMP - INTERVAL '12 days', NULL,
        'Checker B - vessel direct to client', 1, 1, 2, 2),
    (4, 4, 4, CURRENT_TIMESTAMP - INTERVAL '8 days', NULL,
        'Checker B - vessel through stockpile', 1, 0, 2, 2);

-- Trucks are allocated to active projects before transport records are made.
INSERT INTO project_truck_assignment (
    project_truck_assignment_id, project_id, truck_id,
    assignment_started_at, assignment_ended_at, assignment_note,
    is_active, created_by, updated_by
) VALUES
    (1, 1, 1, CURRENT_TIMESTAMP - INTERVAL '30 days', NULL,
        'Dump truck 1 untuk pengiriman langsung project P-MC-001', 1, 2, 2),
    (2, 1, 2, CURRENT_TIMESTAMP - INTERVAL '30 days', NULL,
        'Dump truck 2 untuk pengiriman langsung project P-MC-001', 1, 2, 2),
    (3, 2, 1, CURRENT_TIMESTAMP - INTERVAL '25 days', NULL,
        'Dump truck 1 untuk dua ruas project P-MS-002', 1, 2, 2),
    (4, 2, 2, CURRENT_TIMESTAMP - INTERVAL '25 days', NULL,
        'Dump truck 2 untuk dua ruas project P-MS-002', 1, 2, 2),
    (5, 3, 3, CURRENT_TIMESTAMP - INTERVAL '12 days', NULL,
        'Trailer untuk project kapal langsung P-VC-003', 1, 2, 2),
    (6, 4, 3, CURRENT_TIMESTAMP - INTERVAL '8 days', NULL,
        'Trailer untuk dua ruas project P-VS-004', 1, 2, 2);

INSERT INTO project_route (
    project_route_id, project_id, project_pattern, route_sequence, route_type,
    route_name, distance_km,
    transport_service_unit, transport_service_price_per_cubic,
    transport_service_price_per_ton, transport_service_price_per_transport,
    transport_cost_unit, transport_cost_per_cubic,
    transport_cost_per_ton, transport_cost_per_transport,
    road_money_per_transport, loading_cost_per_transport,
    unloading_cost_per_transport, fuel_cost_per_transport,
    toll_cost_per_transport, other_income_per_transport,
    other_expense_per_transport, route_note,
    is_active, created_by, updated_by
) VALUES
    (1, 1, 'MINE_CLIENT', 1, 'SOURCE_TO_CLIENT',
        'Quarry Banjar ke Batching Plant Bekasi', 120.000,
        'TRANSPORT', NULL, NULL, 400000.00,
        'TON', NULL, 45000.00, NULL,
        200000.00, 50000.00, 50000.00, 100000.00, 25000.00, 0.00, 0.00,
        'Rute langsung', 1, 2, 2),
    (2, 2, 'MINE_STOCKPILE_CLIENT', 1, 'SOURCE_TO_STOCKPILE',
        'Quarry Banjar ke Stockpile Banjarmasin', 35.000,
        'NONE', NULL, NULL, NULL,
        'TON', NULL, 30000.00, NULL,
        150000.00, 30000.00, 30000.00, 60000.00, 0.00, 0.00, 0.00,
        'Ruas penerimaan stockpile', 1, 2, 2),
    (3, 2, 'MINE_STOCKPILE_CLIENT', 2, 'STOCKPILE_TO_CLIENT',
        'Stockpile Banjarmasin ke Batching Plant Bekasi', 105.000,
        'TRANSPORT', NULL, NULL, 350000.00,
        'TON', NULL, 50000.00, NULL,
        250000.00, 30000.00, 30000.00, 100000.00, 25000.00, 0.00, 0.00,
        'Ruas pengiriman ke client', 1, 2, 2),
    (4, 3, 'VESSEL_CLIENT', 1, 'SOURCE_TO_CLIENT',
        'Pelabuhan Trisakti ke Proyek Jakarta Utara', 20.000,
        'TON', NULL, 25000.00, NULL,
        'TON', NULL, 18000.00, NULL,
        220000.00, 60000.00, 60000.00, 80000.00, 10000.00, 0.00, 0.00,
        'Rute langsung dari kapal', 1, 2, 2),
    (5, 4, 'VESSEL_STOCKPILE_CLIENT', 1, 'SOURCE_TO_STOCKPILE',
        'Pelabuhan Tanjung Priok ke Stockpile Cakung', 15.000,
        'NONE', NULL, NULL, NULL,
        'M3', 14000.00, NULL, NULL,
        120000.00, 40000.00, 40000.00, 50000.00, 10000.00, 0.00, 0.00,
        'Ruas kapal ke stockpile', 1, 2, 2),
    (6, 4, 'VESSEL_STOCKPILE_CLIENT', 2, 'STOCKPILE_TO_CLIENT',
        'Stockpile Cakung ke Proyek Jakarta Utara', 28.000,
        'M3', 30000.00, NULL, NULL,
        'M3', 18000.00, NULL, NULL,
        220000.00, 40000.00, 40000.00, 70000.00, 20000.00, 0.00, 0.00,
        'Ruas stockpile ke client', 1, 2, 2);

-- 5. Actual transports. Rates are snapshots copied from the project and route.
INSERT INTO project_transport (
    project_transport_id, project_id, project_route_id,
    transport_number, delivery_note_number, transported_at,
    truck_id, project_truck_assignment_id, driver_id, transport_vendor_id,
    loaded_volume_cubic, loaded_weight_ton,
    delivered_volume_cubic, delivered_weight_ton,
    purchase_volume_cubic, purchase_weight_ton,
    sale_volume_cubic, sale_weight_ton,
    transport_service_volume_cubic, transport_service_weight_ton,
    transport_cost_volume_cubic, transport_cost_weight_ton,
    volume_to_weight_conversion,
    material_purchase_unit, material_buy_price_per_cubic, material_buy_price_per_ton,
    material_sale_unit, material_sell_price_per_cubic, material_sell_price_per_ton,
    transport_service_unit, transport_service_price_per_cubic,
    transport_service_price_per_ton, transport_service_price_per_transport,
    transport_cost_unit, transport_cost_per_cubic,
    transport_cost_per_ton, transport_cost_per_transport,
    road_money_amount, loading_cost_amount, unloading_cost_amount,
    fuel_cost_amount, toll_cost_amount, other_income_amount,
    other_expense_amount, is_completed, transport_note,
    created_by, updated_by
) VALUES
    (1, 1, 1, 'MC-001-001', 'SJ-MC-001-001', CURRENT_TIMESTAMP - INTERVAL '10 days',
        1, 1, 5, 1, 13.333, 20.000, 13.333, 20.000,
        NULL, 20.000, NULL, 20.000, NULL, NULL, NULL, 20.000, 1.500000,
        'TON', NULL, 145000.00, 'TON', NULL, 210000.00,
        'TRANSPORT', NULL, NULL, 400000.00,
        'TON', NULL, 45000.00, NULL,
        200000.00, 50000.00, 50000.00, 100000.00, 25000.00, 0.00, 0.00,
        1, 'Pengiriman langsung pertama selesai', 2, 2),
    (2, 1, 1, 'MC-001-002', 'SJ-MC-001-002', CURRENT_TIMESTAMP - INTERVAL '9 days',
        2, 2, 6, 1, 12.000, 18.000, 12.000, 18.000,
        NULL, 18.000, NULL, 18.000, NULL, NULL, NULL, 18.000, 1.500000,
        'TON', NULL, 145000.00, 'TON', NULL, 210000.00,
        'TRANSPORT', NULL, NULL, 400000.00,
        'TON', NULL, 45000.00, NULL,
        200000.00, 50000.00, 50000.00, 100000.00, 25000.00, 0.00, 0.00,
        1, 'Pengiriman langsung kedua selesai', 2, 2),
    (3, 2, 2, 'MS-002-001', 'SJ-MS-002-001', CURRENT_TIMESTAMP - INTERVAL '7 days',
        1, 3, 5, 1, 12.667, 19.000, 12.667, 19.000,
        NULL, 19.000, NULL, NULL, NULL, NULL, NULL, 19.000, 1.500000,
        'TON', NULL, 140000.00, 'NONE', NULL, NULL,
        'NONE', NULL, NULL, NULL,
        'TON', NULL, 30000.00, NULL,
        150000.00, 30000.00, 30000.00, 60000.00, 0.00, 0.00, 0.00,
        1, 'Penerimaan dari tambang ke stockpile', 2, 2),
    (4, 2, 3, 'MS-002-002', 'SJ-MS-002-002', CURRENT_TIMESTAMP - INTERVAL '6 days',
        2, 4, 6, 1, 12.333, 18.500, 12.333, 18.500,
        NULL, NULL, NULL, 18.500, NULL, NULL, NULL, 18.500, 1.500000,
        'NONE', NULL, NULL, 'TON', NULL, 225000.00,
        'TRANSPORT', NULL, NULL, 350000.00,
        'TON', NULL, 50000.00, NULL,
        250000.00, 30000.00, 30000.00, 100000.00, 25000.00, 0.00, 0.00,
        1, 'Pengiriman dari stockpile ke client', 2, 2),
    (5, 3, 4, 'VC-003-001', 'SJ-VC-003-001', CURRENT_TIMESTAMP - INTERVAL '5 days',
        3, 5, 8, 4, 10.000, 15.000, 10.000, 15.000,
        10.000, NULL, 10.000, NULL, NULL, 15.000, NULL, 15.000, 1.500000,
        'M3', 90000.00, NULL, 'M3', 145000.00, NULL,
        'TON', NULL, 25000.00, NULL,
        'TON', NULL, 18000.00, NULL,
        220000.00, 60000.00, 60000.00, 80000.00, 10000.00, 0.00, 0.00,
        1, 'Pengiriman langsung material kapal', 2, 2),
    (6, 4, 5, 'VS-004-001', 'SJ-VS-004-001', CURRENT_TIMESTAMP - INTERVAL '4 days',
        3, 6, 8, 4, 12.000, 18.000, 12.000, 18.000,
        12.000, NULL, NULL, NULL, NULL, NULL, 12.000, NULL, 1.500000,
        'M3', 75000.00, NULL, 'NONE', NULL, NULL,
        'NONE', NULL, NULL, NULL,
        'M3', 14000.00, NULL, NULL,
        120000.00, 40000.00, 40000.00, 50000.00, 10000.00, 0.00, 0.00,
        1, 'Penerimaan kapal ke stockpile Cakung', 2, 2),
    (7, 4, 6, 'VS-004-002', 'SJ-VS-004-002', CURRENT_TIMESTAMP - INTERVAL '3 days',
        3, 6, 8, 4, 11.500, 17.250, 11.500, 17.250,
        NULL, NULL, 11.500, NULL, 11.500, NULL, 11.500, NULL, 1.500000,
        'NONE', NULL, NULL, 'M3', 130000.00, NULL,
        'M3', 30000.00, NULL, NULL,
        'M3', 18000.00, NULL, NULL,
        220000.00, 40000.00, 40000.00, 70000.00, 20000.00, 0.00, 0.00,
        1, 'Pengiriman stockpile Cakung ke client', 2, 2);

INSERT INTO project_transport_status (
    project_transport_status_id, project_transport_id,
    project_transport_status_type_id, status_time,
    cargo_box_length, cargo_box_width, cargo_box_height,
    cargo_volume_cubic, cargo_weight_ton,
    project_transport_status_note, is_fraud, is_active,
    created_by, updated_by
) VALUES
    (1, 1, 8, CURRENT_TIMESTAMP - INTERVAL '10 days' + INTERVAL '8 hours',
        6.00, 2.30, 1.80, 13.333, 20.000, 'Selesai di client', 0, 1, 3, 3),
    (2, 2, 8, CURRENT_TIMESTAMP - INTERVAL '9 days' + INTERVAL '8 hours',
        6.00, 2.30, 1.80, 12.000, 18.000, 'Selesai di client', 0, 1, 3, 3),
    (3, 3, 8, CURRENT_TIMESTAMP - INTERVAL '7 days' + INTERVAL '3 hours',
        6.00, 2.30, 1.80, 12.667, 19.000, 'Diterima di stockpile', 0, 1, 3, 3),
    (4, 4, 8, CURRENT_TIMESTAMP - INTERVAL '6 days' + INTERVAL '8 hours',
        6.00, 2.30, 1.80, 12.333, 18.500, 'Diterima client', 0, 1, 3, 3),
    (5, 5, 8, CURRENT_TIMESTAMP - INTERVAL '5 days' + INTERVAL '4 hours',
        12.00, 2.50, 2.00, 10.000, 15.000, 'Diterima client', 0, 1, 4, 4),
    (6, 6, 8, CURRENT_TIMESTAMP - INTERVAL '4 days' + INTERVAL '3 hours',
        12.00, 2.50, 2.00, 12.000, 18.000, 'Diterima stockpile', 0, 1, 4, 4),
    (7, 7, 8, CURRENT_TIMESTAMP - INTERVAL '3 days' + INTERVAL '5 hours',
        12.00, 2.50, 2.00, 11.500, 17.250, 'Diterima client', 0, 1, 4, 4);

INSERT INTO project_transport_photo (
    project_transport_photo_id, project_transport_status_id,
    photo_url, photo_type_id, photo_description, created_by, updated_by
) VALUES
    (1, 1, 'https://storage.example.test/project/1/transport/1/delivery-note.jpg',
        1, 'Surat jalan pengiriman project P-MC-001', 3, 3),
    (2, 3, 'https://storage.example.test/project/2/transport/3/cargo-box.jpg',
        2, 'Muatan saat tiba di stockpile', 3, 3),
    (3, 5, 'https://storage.example.test/project/3/transport/5/delivery-note.jpg',
        1, 'Surat jalan pengiriman dari kapal', 4, 4),
    (4, 7, 'https://storage.example.test/project/4/transport/7/delivery-note.jpg',
        1, 'Surat jalan pengiriman dari stockpile', 4, 4);

-- 6. Finance: receivables/payables are separate from actual receipts/payments.
INSERT INTO project_financial_transaction (
    project_financial_transaction_id, project_id, project_transport_id,
    transaction_number, transaction_kind, transaction_date, due_date, paid_at,
    client_id, vendor_id, reference_number,
    material_sale_income, transport_service_income, other_income,
    material_purchase_expense, transport_expense, road_money_expense,
    loading_expense, unloading_expense, fuel_expense, toll_expense, other_expense,
    transaction_note, is_posted, created_by, updated_by
) VALUES
    (1, 1, NULL, 'AR-MC-001', 'RECEIVABLE', CURRENT_DATE - 9, CURRENT_DATE + 21, NULL,
        1, NULL, 'INV-MC-001',
        7980000.00, 800000.00, 0.00,
        0.00, 0.00, 0.00, 0.00, 0.00, 0.00, 0.00, 0.00,
        'Tagihan material dan jasa angkut project P-MC-001', 1, 7, 7),
    (2, 1, NULL, 'AP-MC-001', 'PAYABLE', CURRENT_DATE - 9, CURRENT_DATE + 14, NULL,
        NULL, 1, 'BILL-MC-001',
        0.00, 0.00, 0.00,
        5510000.00, 1710000.00, 400000.00, 100000.00, 100000.00, 200000.00, 50000.00, 0.00,
        'Utang material dan biaya operasi project P-MC-001', 1, 7, 7),
    (3, 1, NULL, 'RC-MC-001', 'RECEIPT', CURRENT_DATE - 2, NULL, CURRENT_TIMESTAMP - INTERVAL '2 days',
        1, NULL, 'RCPT-MC-001',
        4000000.00, 400000.00, 0.00,
        0.00, 0.00, 0.00, 0.00, 0.00, 0.00, 0.00, 0.00,
        'Penerimaan sebagian tagihan project P-MC-001', 1, 7, 7),
    (4, 2, NULL, 'AR-MS-002', 'RECEIVABLE', CURRENT_DATE - 6, CURRENT_DATE + 24, NULL,
        1, NULL, 'INV-MS-002',
        4162500.00, 350000.00, 0.00,
        0.00, 0.00, 0.00, 0.00, 0.00, 0.00, 0.00, 0.00,
        'Tagihan material dan pengiriman project P-MS-002', 1, 7, 7),
    (5, 2, NULL, 'AP-MS-002', 'PAYABLE', CURRENT_DATE - 6, CURRENT_DATE + 14, NULL,
        NULL, 1, 'BILL-MS-002',
        0.00, 0.00, 0.00,
        2660000.00, 1495000.00, 400000.00, 60000.00, 60000.00, 160000.00, 25000.00, 0.00,
        'Utang material dan dua ruas transport project P-MS-002', 1, 7, 7),
    (6, 3, 5, 'AR-VC-003', 'RECEIVABLE', CURRENT_DATE - 5, CURRENT_DATE + 16, NULL,
        2, NULL, 'INV-VC-003',
        1450000.00, 375000.00, 0.00,
        0.00, 0.00, 0.00, 0.00, 0.00, 0.00, 0.00, 0.00,
        'Tagihan kapal langsung ke client', 1, 7, 7),
    (7, 3, 5, 'AP-VC-003', 'PAYABLE', CURRENT_DATE - 5, CURRENT_DATE + 10, NULL,
        NULL, 3, 'BILL-VC-003',
        0.00, 0.00, 0.00,
        900000.00, 270000.00, 220000.00, 60000.00, 60000.00, 80000.00, 10000.00, 0.00,
        'Utang pembelian dan pengiriman material kapal', 1, 7, 7),
    (8, 4, NULL, 'AR-VS-004', 'RECEIVABLE', CURRENT_DATE - 3, CURRENT_DATE + 18, NULL,
        2, NULL, 'INV-VS-004',
        1495000.00, 345000.00, 0.00,
        0.00, 0.00, 0.00, 0.00, 0.00, 0.00, 0.00, 0.00,
        'Tagihan material dari stockpile ke client', 1, 7, 7),
    (9, 4, NULL, 'AP-VS-004', 'PAYABLE', CURRENT_DATE - 3, CURRENT_DATE + 12, NULL,
        NULL, 3, 'BILL-VS-004',
        0.00, 0.00, 0.00,
        900000.00, 375000.00, 340000.00, 80000.00, 80000.00, 120000.00, 30000.00, 0.00,
        'Utang pembelian dan dua ruas pengiriman kapal', 1, 7, 7);

-- 7. Stockpile movements are allocated by project, not just by material.
INSERT INTO stockpile_adjustment (
    stockpile_adjustment_id, project_id, stockpile_cargo_id,
    adjustment_date, project_transport_id, reference_type_id,
    amount_volume_cubic, amount_weight_ton, reason, created_by, updated_by
) VALUES
    (1, 2, 1, CURRENT_DATE - 7, 3, 1, 12.667, 19.000,
        'Penerimaan dari Quarry Banjar', 2, 2),
    (2, 2, 1, CURRENT_DATE - 6, 4, 2, -12.333, -18.500,
        'Pengiriman dari Stockpile Banjarmasin ke client', 2, 2),
    (3, 4, 3, CURRENT_DATE - 4, 6, 1, 12.000, 18.000,
        'Penerimaan material dari kapal', 2, 2),
    (4, 4, 3, CURRENT_DATE - 3, 7, 2, -11.500, -17.250,
        'Pengiriman dari Stockpile Cakung ke client', 2, 2);

INSERT INTO stockpile_ledger (
    stockpile_ledger_id, project_id, stockpile_cargo_id,
    project_transport_id, adjustment_id, reference_type_id,
    amount_volume_cubic, balance_volume_cubic,
    amount_weight_ton, balance_weight_ton, created_by, updated_by
) VALUES
    (1, 2, 1, 3, 1, 1, 12.667, 12.667, 19.000, 19.000, 2, 2),
    (2, 2, 1, 4, 2, 2, -12.333, 0.334, -18.500, 0.500, 2, 2),
    (3, 4, 3, 6, 3, 1, 12.000, 12.000, 18.000, 18.000, 2, 2),
    (4, 4, 3, 7, 4, 2, -11.500, 0.500, -17.250, 0.750, 2, 2);

-- Keep sequences aligned with the explicit IDs above.
SELECT setval(pg_get_serial_sequence('app_role', 'app_role_id'), 6, true);
SELECT setval(pg_get_serial_sequence('bank_merk', 'bank_merk_id'), 2, true);
SELECT setval(pg_get_serial_sequence('app_user', 'app_user_id'), 9, true);
SELECT setval(pg_get_serial_sequence('app_setting', 'app_setting_id'), 4, true);
SELECT setval(pg_get_serial_sequence('province', 'province_id'), 3, true);
SELECT setval(pg_get_serial_sequence('city', 'city_id'), 3, true);
SELECT setval(pg_get_serial_sequence('vendor_type', 'vendor_type_id'), 3, true);
SELECT setval(pg_get_serial_sequence('vendor', 'vendor_id'), 4, true);
SELECT setval(pg_get_serial_sequence('truck_type', 'truck_type_id'), 2, true);
SELECT setval(pg_get_serial_sequence('truck_merk', 'truck_merk_id'), 2, true);
SELECT setval(pg_get_serial_sequence('truck', 'truck_id'), 3, true);
SELECT setval(pg_get_serial_sequence('client', 'client_id'), 2, true);
SELECT setval(pg_get_serial_sequence('client_destination', 'client_destination_id'), 2, true);
SELECT setval(pg_get_serial_sequence('cargo_type', 'cargo_type_id'), 2, true);
SELECT setval(pg_get_serial_sequence('mine', 'mine_id'), 2, true);
SELECT setval(pg_get_serial_sequence('port', 'port_id'), 2, true);
SELECT setval(pg_get_serial_sequence('vessel', 'vessel_id'), 2, true);
SELECT setval(pg_get_serial_sequence('vessel_cargo', 'vessel_cargo_id'), 2, true);
SELECT setval(pg_get_serial_sequence('stockpile', 'stockpile_id'), 2, true);
SELECT setval(pg_get_serial_sequence('stockpile_cargo', 'stockpile_cargo_id'), 3, true);
SELECT setval(pg_get_serial_sequence('project', 'project_id'), 4, true);
SELECT setval(pg_get_serial_sequence('project_checker_assignment', 'project_checker_assignment_id'), 4, true);
SELECT setval(pg_get_serial_sequence('project_truck_assignment', 'project_truck_assignment_id'), 6, true);
SELECT setval(pg_get_serial_sequence('project_route', 'project_route_id'), 6, true);
SELECT setval(pg_get_serial_sequence('project_transport', 'project_transport_id'), 7, true);
SELECT setval(pg_get_serial_sequence('project_transport_status', 'project_transport_status_id'), 7, true);
SELECT setval(pg_get_serial_sequence('project_transport_photo', 'project_transport_photo_id'), 4, true);
SELECT setval(pg_get_serial_sequence('project_financial_transaction', 'project_financial_transaction_id'), 9, true);
SELECT setval(pg_get_serial_sequence('stockpile_adjustment', 'stockpile_adjustment_id'), 4, true);
SELECT setval(pg_get_serial_sequence('stockpile_ledger', 'stockpile_ledger_id'), 4, true);

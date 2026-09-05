-- Seeder data operasional PML.
-- Jalankan file ini setelah create_tables.sql pada database yang masih kosong.
-- Password awal seluruh user di file ini adalah: Password123!
--
-- Catatan harga:
-- Struktur project hanya mendukung harga pembelian material per M3 atau TON.
-- Karena harga beli yang diminta adalah Rp500.000 per truk/sekali angkut,
-- nilainya disimpan sebagai transport_cost_per_transport pada project_route.
-- Harga jual disimpan per M3 pada project, sedangkan jasa transportasi
-- disimpan per sekali angkut pada project_route.
--
-- Seeder ini sengaja tidak mengisi project_transport beserta tabel status,
-- foto, transaksi keuangan, dan pergerakan stockpile turunannya.

BEGIN;

SET LOCAL TIMEZONE = 'Asia/Jakarta';

CREATE EXTENSION IF NOT EXISTS pgcrypto;

-- 1. Role dan user bootstrap
INSERT INTO app_role (
    app_role_id,
    app_role_type_id,
    app_role_name,
    app_role_description,
    is_active,
    created_by,
    updated_by
) VALUES
    (
        1, 1, 'driver',
        'Melihat penugasan dan perjalanan miliknya.',
        1, 1, 1
    ),
    (
        2, 2, 'checker',
        'Mengelola pemeriksaan perjalanan pada project yang ditugaskan.',
        1, 1, 1
    ),
    (
        3, 3, 'owner',
        'Memiliki akses penuh ke seluruh fitur, data, dan konfigurasi.',
        1, 1, 1
    ),
    (
        4, 3, 'itdev',
        'Memiliki akses penuh ke seluruh fitur, data, dan konfigurasi.',
        1, 1, 1
    ),
    (
        5, 3, 'superadmin',
        'Memiliki akses penuh ke seluruh fitur, data, dan konfigurasi.',
        1, 1, 1
    ),
    (
        6, 3, 'admin',
        'Mengelola data master dan operasional harian.',
        1, 1, 1
    );

-- itdev dibuat pertama untuk menjadi audit user bagi seluruh seed berikutnya.
INSERT INTO app_user (
    app_user_id,
    app_role_id,
    bank_merk_id,
    username,
    password,
    app_user_status_id,
    app_user_name,
    app_user_preferred_name,
    app_user_phone,
    city_id,
    app_user_address,
    salary_percentage,
    created_by,
    updated_by
) VALUES (
    1,
    4,
    NULL,
    'itdev',
    crypt('Password123!', gen_salt('bf', 10)),
    1,
    'IT Developer',
    'IT Dev',
    '080000000001',
    NULL,
    'Kantor PML',
    NULL,
    NULL,
    NULL
);

INSERT INTO app_setting (
    app_setting_id,
    app_setting_key,
    app_setting_value,
    app_setting_description,
    is_active,
    created_by,
    updated_by
) VALUES
    (
        1,
        'operations.timezone',
        '{"iana":"Asia/Jakarta"}'::jsonb,
        'Zona waktu operasional dan pelaporan.',
        1, 1, 1
    ),
    (
        2,
        'checker.location_cooldown_minutes',
        '{"minutes":5}'::jsonb,
        'Batas perubahan lokasi checker pada aplikasi mobile.',
        1, 1, 1
    ),
    (
        3,
        'checker.truck_cooldown_minutes',
        '{"minutes":10}'::jsonb,
        'Batas input truk yang sama oleh checker pada aplikasi mobile.',
        1, 1, 1
    ),
    (
        4,
        'reporting.currency',
        '{"code":"IDR","locale":"id-ID"}'::jsonb,
        'Mata uang standar untuk dashboard dan laporan.',
        1, 1, 1
    );

-- 2. Wilayah
INSERT INTO province (
    province_id,
    province_name,
    province_real_name,
    is_active,
    created_by,
    updated_by
) VALUES
    (1, 'Jawa Barat', 'Jawa Barat', 1, 1, 1),
    (2, 'Jawa Timur', 'Jawa Timur', 1, 1, 1);

INSERT INTO city (
    city_id,
    province_id,
    city_name,
    is_active,
    created_by,
    updated_by
) VALUES
    (1, 1, 'Subang', 1, 1, 1),
    (2, 2, 'Gresik', 1, 1, 1);

UPDATE app_user
SET
    city_id = 1,
    updated_by = 1
WHERE app_user_id = 1;

-- Owner, checker, dan driver
INSERT INTO app_user (
    app_user_id,
    app_role_id,
    bank_merk_id,
    username,
    password,
    app_user_status_id,
    app_user_name,
    app_user_preferred_name,
    app_user_phone,
    city_id,
    app_user_address,
    salary_percentage,
    created_by,
    updated_by
) VALUES
    (
        2, 3, NULL, 'judijanto',
        crypt('Password123!', gen_salt('bf', 10)),
        1, 'judijanto', 'judijanto', '080000000002', 1,
        'Subang', NULL, 1, 1
    ),
    (
        3, 2, NULL, 'checker1',
        crypt('Password123!', gen_salt('bf', 10)),
        1, 'Checker 1', 'Checker 1', '080000000003', 1,
        'Subang', NULL, 1, 1
    ),
    (
        4, 2, NULL, 'checker2',
        crypt('Password123!', gen_salt('bf', 10)),
        1, 'Checker 2', 'Checker 2', '080000000004', 1,
        'Subang', NULL, 1, 1
    ),
    (
        5, 2, NULL, 'checker3',
        crypt('Password123!', gen_salt('bf', 10)),
        1, 'Checker 3', 'Checker 3', '080000000005', 1,
        'Subang', NULL, 1, 1
    ),
    (
        6, 2, NULL, 'checker4',
        crypt('Password123!', gen_salt('bf', 10)),
        1, 'Checker 4', 'Checker 4', '080000000006', 1,
        'Subang', NULL, 1, 1
    ),
    (
        7, 2, NULL, 'checker5',
        crypt('Password123!', gen_salt('bf', 10)),
        1, 'Checker 5', 'Checker 5', '080000000007', 2,
        'Gresik', NULL, 1, 1
    ),
    (
        8, 2, NULL, 'checker6',
        crypt('Password123!', gen_salt('bf', 10)),
        1, 'Checker 6', 'Checker 6', '080000000008', 2,
        'Gresik', NULL, 1, 1
    ),
    (
        9, 1, NULL, 'driver1',
        crypt('Password123!', gen_salt('bf', 10)),
        1, 'Driver 1', 'Driver 1', '080000000009', 1,
        'Subang', NULL, 1, 1
    ),
    (
        10, 1, NULL, 'driver2',
        crypt('Password123!', gen_salt('bf', 10)),
        1, 'Driver 2', 'Driver 2', '080000000010', 1,
        'Subang', NULL, 1, 1
    ),
    (
        11, 1, NULL, 'driver3',
        crypt('Password123!', gen_salt('bf', 10)),
        1, 'Driver 3', 'Driver 3', '080000000011', 1,
        'Subang', NULL, 1, 1
    ),
    (
        12, 1, NULL, 'driver4',
        crypt('Password123!', gen_salt('bf', 10)),
        1, 'Driver 4', 'Driver 4', '080000000012', 2,
        'Gresik', NULL, 1, 1
    ),
    (
        13, 1, NULL, 'driver5',
        crypt('Password123!', gen_salt('bf', 10)),
        1, 'Driver 5', 'Driver 5', '080000000013', 2,
        'Gresik', NULL, 1, 1
    );

-- 3. Vendor dan armada
INSERT INTO vendor_type (
    vendor_type_id,
    vendor_type_name,
    vendor_type_description,
    is_active,
    created_by,
    updated_by
) VALUES (
    1,
    'Transporter',
    'Vendor jasa transportasi darat.',
    1, 1, 1
);

INSERT INTO vendor (
    vendor_id,
    vendor_type_id,
    bank_merk_id,
    vendor_name,
    vendor_email,
    vendor_phone,
    vendor_tin,
    city_id,
    vendor_address,
    bank_account_number,
    bank_account_name,
    is_active,
    created_by,
    updated_by
) VALUES (
    1,
    1,
    NULL,
    'PT Angkut angkut mania',
    NULL,
    NULL,
    NULL,
    1,
    'Subang',
    NULL,
    NULL,
    1,
    1,
    1
);

INSERT INTO truck_type (
    truck_type_id,
    truck_type_name,
    truck_type_description,
    truck_box_length,
    truck_box_width,
    truck_box_height,
    truck_capacity,
    is_active,
    created_by,
    updated_by
) VALUES (
    1,
    'Index 28',
    'Bak panjang 6 m, lebar 2 m, tinggi 2,15 m.',
    6.00,
    2.00,
    2.15,
    28.00,
    1,
    1,
    1
);

INSERT INTO truck_merk (
    truck_merk_id,
    truck_merk_name,
    is_active,
    created_by,
    updated_by
) VALUES (
    1,
    'Hino',
    1,
    1,
    1
);

INSERT INTO truck (
    truck_id,
    truck_type_id,
    truck_merk_id,
    driver_id,
    vendor_id,
    license_plate,
    ownership_status_id,
    production_year,
    number_of_tires,
    is_active,
    created_by,
    updated_by
) VALUES
    (1, 1, 1, 9, 1, 'D 9757 VC', 2, NULL, NULL, 1, 1, 1),
    (2, 1, 1, 10, 1, 'B 9794 EC', 2, NULL, NULL, 1, 1, 1),
    (3, 1, 1, 11, 1, 'D 9711 VC', 2, NULL, NULL, 1, 1, 1),
    (4, 1, 1, 12, 1, 'D 8581 HL', 2, NULL, NULL, 1, 1, 1),
    (5, 1, 1, 13, 1, 'B 9074 CU', 2, NULL, NULL, 1, 1, 1);

-- 4. Client dan lokasi project
INSERT INTO client (
    client_id,
    client_name,
    client_email,
    client_tin,
    number_of_day_until_due,
    city_id,
    client_address,
    operating_hours,
    is_active,
    created_by,
    updated_by
) VALUES
    (
        1, 'Client Tol Patimban Paket 1', NULL, NULL, NULL,
        1, 'Patimban, Subang', NULL, 1, 1, 1
    ),
    (
        2, 'Client Pertamina Patimban - Subang', NULL, NULL, NULL,
        1, 'Patimban, Subang', NULL, 1, 1, 1
    ),
    (
        3, 'Client AKR - JIIPE Gresik', NULL, NULL, NULL,
        2, 'JIIPE, Gresik', NULL, 1, 1, 1
    );

INSERT INTO client_destination (
    client_destination_id,
    client_id,
    client_destination_name,
    city_id,
    client_destination_address,
    client_destination_latitude,
    client_destination_longitude,
    client_destination_map_url,
    operating_hours,
    is_active,
    created_by,
    updated_by
) VALUES
    (
        1,
        1,
        'Penimbunan Tol Patimban Paket 1',
        1,
        'Lokasi Penimbunan Tol Patimban Paket 1, Patimban, Subang',
        NULL,
        NULL,
        'https://www.google.com/maps/search/?api=1&query=Tol+Patimban+Subang',
        NULL,
        1, 1, 1
    ),
    (
        2,
        2,
        'Penimbunan Pertamina Patimban - Subang',
        1,
        'Lokasi Penimbunan Pertamina Patimban, Subang',
        NULL,
        NULL,
        'https://www.google.com/maps/search/?api=1&query=Pertamina+Patimban+Subang',
        NULL,
        1, 1, 1
    ),
    (
        3,
        3,
        'Penimbunan AKR - JIIPE Gresik',
        2,
        'Lokasi Penimbunan AKR, JIIPE, Gresik',
        NULL,
        NULL,
        'https://www.google.com/maps/search/?api=1&query=AKR+JIIPE+Gresik',
        NULL,
        1, 1, 1
    );

INSERT INTO cargo_type (
    cargo_type_id,
    cargo_type_name,
    cargo_type_description,
    cargo_type_grade,
    is_active,
    created_by,
    updated_by
) VALUES (
    1,
    'Material Timbunan',
    'Material untuk pekerjaan penimbunan.',
    'Standar',
    1,
    1,
    1
);

-- Project 1 dan 2 menggunakan Tambang 1 dari CV. Bumi Kahuripan Abadi.
-- Project 3 menggunakan Tambang 2.
INSERT INTO mine (
    mine_id,
    mine_name,
    city_id,
    mine_address,
    mine_latitude,
    mine_longitude,
    mine_map_url,
    is_active,
    created_by,
    updated_by
) VALUES
    (
        1,
        'Tambang 1',
        1,
        'Tambang 1 - CV. Bumi Kahuripan Abadi, Subang',
        NULL,
        NULL,
        'https://www.google.com/maps/search/?api=1&query=CV+Bumi+Kahuripan+Abadi+Subang',
        1,
        1,
        1
    ),
    (
        2,
        'Tambang 2',
        2,
        'Tambang 2, Gresik',
        NULL,
        NULL,
        'https://www.google.com/maps/search/?api=1&query=Tambang+Gresik',
        1,
        1,
        1
    );

-- 5. Project: semuanya berpola langsung dari tambang ke client.
INSERT INTO project (
    project_id,
    project_code,
    project_name,
    route_type,
    mine_id,
    vessel_cargo_id,
    stockpile_cargo_id,
    client_destination_id,
    cargo_type_id,
    project_status,
    start_date,
    end_date,
    planned_volume_cubic,
    planned_weight_ton,
    volume_to_weight_conversion,
    material_purchase_unit,
    material_buy_price_per_cubic,
    material_buy_price_per_ton,
    material_sale_unit,
    material_sell_price_per_cubic,
    material_sell_price_per_ton,
    fixed_other_income,
    fixed_other_expense,
    project_note,
    is_active,
    created_by,
    updated_by
) VALUES
    (
        1,
        'PRJ-PATIMBAN-01',
        'Penimbunan Tol Patimban Paket 1',
        'MINE_CLIENT',
        1,
        NULL,
        NULL,
        1,
        1,
        'ACTIVE',
        CURRENT_DATE,
        NULL,
        NULL,
        NULL,
        NULL,
        'NONE',
        NULL,
        NULL,
        'M3',
        52000.00,
        NULL,
        0,
        0,
        'Sumber material: Tambang 1 - CV. Bumi Kahuripan Abadi.',
        1,
        1,
        1
    ),
    (
        2,
        'PRJ-PERTAMINA-02',
        'Penimbunan Pertamina Patimban - Subang',
        'MINE_CLIENT',
        1,
        NULL,
        NULL,
        2,
        1,
        'ACTIVE',
        CURRENT_DATE,
        NULL,
        NULL,
        NULL,
        NULL,
        'NONE',
        NULL,
        NULL,
        'M3',
        52000.00,
        NULL,
        0,
        0,
        'Sumber material: Tambang 1 - CV. Bumi Kahuripan Abadi.',
        1,
        1,
        1
    ),
    (
        3,
        'PRJ-AKR-JIIPE-03',
        'Penimbunan AKR - JIIPE Gresik',
        'MINE_CLIENT',
        2,
        NULL,
        NULL,
        3,
        1,
        'ACTIVE',
        CURRENT_DATE,
        NULL,
        NULL,
        NULL,
        NULL,
        'NONE',
        NULL,
        NULL,
        'M3',
        52000.00,
        NULL,
        0,
        0,
        'Sumber material: Tambang 2.',
        1,
        1,
        1
    );

INSERT INTO project_route (
    project_route_id,
    project_id,
    project_pattern,
    route_sequence,
    route_type,
    route_name,
    distance_km,
    transport_service_unit,
    transport_service_price_per_cubic,
    transport_service_price_per_ton,
    transport_service_price_per_transport,
    transport_cost_unit,
    transport_cost_per_cubic,
    transport_cost_per_ton,
    transport_cost_per_transport,
    road_money_per_transport,
    loading_cost_per_transport,
    unloading_cost_per_transport,
    fuel_cost_per_transport,
    toll_cost_per_transport,
    other_income_per_transport,
    other_expense_per_transport,
    route_note,
    is_active,
    created_by,
    updated_by
) VALUES
    (
        1, 1, 'MINE_CLIENT', 1, 'SOURCE_TO_CLIENT',
        'Tambang 1 - Penimbunan Tol Patimban Paket 1',
        NULL,
        'TRANSPORT', NULL, NULL, 650000.00,
        'TRANSPORT', NULL, NULL, 500000.00,
        0, 0, 0, 0, 0, 0, 0,
        'Rute langsung Tambang 1 ke lokasi project.',
        1, 1, 1
    ),
    (
        2, 2, 'MINE_CLIENT', 1, 'SOURCE_TO_CLIENT',
        'Tambang 1 - Penimbunan Pertamina Patimban - Subang',
        NULL,
        'TRANSPORT', NULL, NULL, 650000.00,
        'TRANSPORT', NULL, NULL, 500000.00,
        0, 0, 0, 0, 0, 0, 0,
        'Rute langsung Tambang 1 ke lokasi project.',
        1, 1, 1
    ),
    (
        3, 3, 'MINE_CLIENT', 1, 'SOURCE_TO_CLIENT',
        'Tambang 2 - Penimbunan AKR - JIIPE Gresik',
        NULL,
        'TRANSPORT', NULL, NULL, 650000.00,
        'TRANSPORT', NULL, NULL, 500000.00,
        0, 0, 0, 0, 0, 0, 0,
        'Rute langsung Tambang 2 ke lokasi project.',
        1, 1, 1
    );

-- Setiap project mendapatkan dua checker.
INSERT INTO project_checker_assignment (
    project_checker_assignment_id,
    project_id,
    checker_id,
    access_started_at,
    access_ended_at,
    assignment_note,
    is_active,
    is_default,
    created_by,
    updated_by
) VALUES
    (
        1, 1, 3, CURRENT_TIMESTAMP, NULL,
        'Checker 1 untuk Project 1.', 1, 1, 1, 1
    ),
    (
        2, 1, 4, CURRENT_TIMESTAMP, NULL,
        'Checker 2 untuk Project 1.', 1, 1, 1, 1
    ),
    (
        3, 2, 5, CURRENT_TIMESTAMP, NULL,
        'Checker 3 untuk Project 2.', 1, 1, 1, 1
    ),
    (
        4, 2, 6, CURRENT_TIMESTAMP, NULL,
        'Checker 4 untuk Project 2.', 1, 1, 1, 1
    ),
    (
        5, 3, 7, CURRENT_TIMESTAMP, NULL,
        'Checker 5 untuk Project 3.', 1, 1, 1, 1
    ),
    (
        6, 3, 8, CURRENT_TIMESTAMP, NULL,
        'Checker 6 untuk Project 3.', 1, 1, 1, 1
    );

-- Armada vendor dipakai bersama, sehingga seluruh truk tersedia pada pilihan
-- transport untuk ketiga project.
INSERT INTO project_truck_assignment (
    project_truck_assignment_id,
    project_id,
    truck_id,
    assignment_started_at,
    assignment_ended_at,
    assignment_note,
    is_active,
    created_by,
    updated_by
) VALUES
    (1, 1, 1, CURRENT_TIMESTAMP, NULL, 'Armada bersama Project 1.', 1, 1, 1),
    (2, 1, 2, CURRENT_TIMESTAMP, NULL, 'Armada bersama Project 1.', 1, 1, 1),
    (3, 1, 3, CURRENT_TIMESTAMP, NULL, 'Armada bersama Project 1.', 1, 1, 1),
    (4, 1, 4, CURRENT_TIMESTAMP, NULL, 'Armada bersama Project 1.', 1, 1, 1),
    (5, 1, 5, CURRENT_TIMESTAMP, NULL, 'Armada bersama Project 1.', 1, 1, 1),
    (6, 2, 1, CURRENT_TIMESTAMP, NULL, 'Armada bersama Project 2.', 1, 1, 1),
    (7, 2, 2, CURRENT_TIMESTAMP, NULL, 'Armada bersama Project 2.', 1, 1, 1),
    (8, 2, 3, CURRENT_TIMESTAMP, NULL, 'Armada bersama Project 2.', 1, 1, 1),
    (9, 2, 4, CURRENT_TIMESTAMP, NULL, 'Armada bersama Project 2.', 1, 1, 1),
    (10, 2, 5, CURRENT_TIMESTAMP, NULL, 'Armada bersama Project 2.', 1, 1, 1),
    (11, 3, 1, CURRENT_TIMESTAMP, NULL, 'Armada bersama Project 3.', 1, 1, 1),
    (12, 3, 2, CURRENT_TIMESTAMP, NULL, 'Armada bersama Project 3.', 1, 1, 1),
    (13, 3, 3, CURRENT_TIMESTAMP, NULL, 'Armada bersama Project 3.', 1, 1, 1),
    (14, 3, 4, CURRENT_TIMESTAMP, NULL, 'Armada bersama Project 3.', 1, 1, 1),
    (15, 3, 5, CURRENT_TIMESTAMP, NULL, 'Armada bersama Project 3.', 1, 1, 1);

-- Sinkronisasi sequence setelah insert menggunakan ID eksplisit.
SELECT setval(pg_get_serial_sequence('app_role', 'app_role_id'), 6, true);
SELECT setval(pg_get_serial_sequence('app_user', 'app_user_id'), 13, true);
SELECT setval(pg_get_serial_sequence('app_setting', 'app_setting_id'), 4, true);
SELECT setval(pg_get_serial_sequence('province', 'province_id'), 2, true);
SELECT setval(pg_get_serial_sequence('city', 'city_id'), 2, true);
SELECT setval(pg_get_serial_sequence('vendor_type', 'vendor_type_id'), 1, true);
SELECT setval(pg_get_serial_sequence('vendor', 'vendor_id'), 1, true);
SELECT setval(pg_get_serial_sequence('truck_type', 'truck_type_id'), 1, true);
SELECT setval(pg_get_serial_sequence('truck_merk', 'truck_merk_id'), 1, true);
SELECT setval(pg_get_serial_sequence('truck', 'truck_id'), 5, true);
SELECT setval(pg_get_serial_sequence('client', 'client_id'), 3, true);
SELECT setval(
    pg_get_serial_sequence('client_destination', 'client_destination_id'),
    3,
    true
);
SELECT setval(pg_get_serial_sequence('cargo_type', 'cargo_type_id'), 1, true);
SELECT setval(pg_get_serial_sequence('mine', 'mine_id'), 2, true);
SELECT setval(pg_get_serial_sequence('project', 'project_id'), 3, true);
SELECT setval(
    pg_get_serial_sequence(
        'project_checker_assignment',
        'project_checker_assignment_id'
    ),
    6,
    true
);
SELECT setval(
    pg_get_serial_sequence(
        'project_truck_assignment',
        'project_truck_assignment_id'
    ),
    15,
    true
);
SELECT setval(
    pg_get_serial_sequence('project_route', 'project_route_id'),
    3,
    true
);

-- Validasi data inti sebelum transaksi seed disimpan.
DO $$
DECLARE
    project_count INT;
    distinct_client_count INT;
    truck_count INT;
    checker_count INT;
    driver_count INT;
    truck_assignment_count INT;
BEGIN
    SELECT COUNT(*) INTO project_count FROM project;
    SELECT COUNT(DISTINCT destination.client_id)
    INTO distinct_client_count
    FROM project
    JOIN client_destination AS destination
        ON destination.client_destination_id = project.client_destination_id;
    SELECT COUNT(*) INTO truck_count FROM truck;
    SELECT COUNT(*) INTO checker_count
    FROM app_user
    WHERE app_role_id = 2;
    SELECT COUNT(*) INTO driver_count
    FROM app_user
    WHERE app_role_id = 1;
    SELECT COUNT(*) INTO truck_assignment_count
    FROM project_truck_assignment
    WHERE is_active = 1 AND deleted_at IS NULL;

    IF project_count <> 3 THEN
        RAISE EXCEPTION 'Seeder harus menghasilkan tepat 3 project';
    END IF;

    IF distinct_client_count <> 3 THEN
        RAISE EXCEPTION 'Setiap project harus memiliki client yang berbeda';
    END IF;

    IF truck_count <> 5 OR driver_count <> 5 OR checker_count <> 6 THEN
        RAISE EXCEPTION
            'Seeder harus menghasilkan 5 truk, 5 driver, dan 6 checker';
    END IF;

    IF truck_assignment_count <> 15 THEN
        RAISE EXCEPTION
            'Seluruh 5 truk harus aktif pada masing-masing dari 3 project';
    END IF;

    IF EXISTS (
        SELECT 1
        FROM project
        WHERE route_type <> 'MINE_CLIENT'
            OR material_sale_unit <> 'M3'
            OR material_sell_price_per_cubic <> 52000.00
    ) THEN
        RAISE EXCEPTION
            'Seluruh project harus MINE_CLIENT dengan harga jual Rp52.000/M3';
    END IF;

    IF EXISTS (
        SELECT 1
        FROM project_route
        WHERE transport_cost_unit <> 'TRANSPORT'
            OR transport_cost_per_transport <> 500000.00
            OR transport_service_unit <> 'TRANSPORT'
            OR transport_service_price_per_transport <> 650000.00
    ) THEN
        RAISE EXCEPTION
            'Harga rute harus Rp500.000 biaya dan Rp650.000 jasa per angkut';
    END IF;

    IF EXISTS (SELECT 1 FROM project_transport) THEN
        RAISE EXCEPTION 'Data transport harus tetap kosong';
    END IF;
END
$$;

COMMIT;

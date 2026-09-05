CREATE TABLE app_role (
    app_role_id SERIAL PRIMARY KEY,
    app_role_type_id SMALLINT NOT NULL,
    app_role_name VARCHAR(50) UNIQUE NOT NULL,
    app_role_description TEXT,
    is_active SMALLINT NOT NULL DEFAULT 1,
    created_by INT NOT NULL,
    updated_by INT NOT NULL,
    deleted_by INT,
    created_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMPTZ
);

CREATE TABLE bank_merk (
    bank_merk_id SERIAL PRIMARY KEY,
    bank_merk_name VARCHAR(50) UNIQUE NOT NULL,
    bank_merk_description TEXT,
    is_active SMALLINT NOT NULL DEFAULT 1,
    created_by INT NOT NULL,
    updated_by INT NOT NULL,
    deleted_by INT,
    created_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMPTZ
);



CREATE TABLE app_user (
    app_user_id SERIAL PRIMARY KEY,
    app_role_id INT NOT NULL REFERENCES app_role(app_role_id),
    bank_merk_id INT REFERENCES bank_merk(bank_merk_id),
    client_id INT,
    username VARCHAR(200) UNIQUE,
    password TEXT NOT NULL,
    app_user_status_id SMALLINT NOT NULL,
    app_user_name VARCHAR(200) NOT NULL,
    app_user_preferred_name VARCHAR(200),
    app_user_phone VARCHAR(20) NOT NULL,
    city_id INT NULL,
    app_user_address TEXT,
    app_user_photo_url TEXT,

    -- Additional Personal Data
    id_card_photo_url TEXT,
    id_card_number VARCHAR(50),
    family_card_photo_url TEXT,
    family_card_number VARCHAR(50),
    driver_license_b_photo_url TEXT,
    driver_license_b_number VARCHAR(50),
    driver_license_b_expiry DATE,
    tax_id_photo_url TEXT,
    tax_id_number VARCHAR(50),
    bpjs_photo_url TEXT,
    bpjs_number VARCHAR(50),

    bank_account_number VARCHAR(100),
    bank_account_name VARCHAR(100),
    salary_percentage NUMERIC(5,2),
    created_by INT REFERENCES app_user(app_user_id),
    updated_by INT REFERENCES app_user(app_user_id),
    deleted_by INT REFERENCES app_user(app_user_id),
    created_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMPTZ
);

CREATE TABLE province (
    province_id SERIAL PRIMARY KEY,
    province_name VARCHAR(100) NOT NULL,
    province_real_name VARCHAR(100),
    is_active SMALLINT NOT NULL DEFAULT 1,
    created_by INT NOT NULL REFERENCES app_user(app_user_id),
    updated_by INT NOT NULL REFERENCES app_user(app_user_id),
    deleted_by INT REFERENCES app_user(app_user_id),
    created_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMPTZ
);

CREATE TABLE city (
    city_id SERIAL PRIMARY KEY,
    province_id INT REFERENCES province(province_id),
    city_name VARCHAR(100) NOT NULL,
    is_active SMALLINT NOT NULL DEFAULT 1,
    created_by INT NOT NULL REFERENCES app_user(app_user_id),
    updated_by INT NOT NULL REFERENCES app_user(app_user_id),
    deleted_by INT REFERENCES app_user(app_user_id),
    created_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMPTZ
);


CREATE TABLE vendor_type (
    vendor_type_id SERIAL PRIMARY KEY,
    vendor_type_name VARCHAR(100) NOT NULL,
    vendor_type_description TEXT,
    is_active SMALLINT NOT NULL DEFAULT 1,
    created_by INT NOT NULL REFERENCES app_user(app_user_id),
    updated_by INT NOT NULL REFERENCES app_user(app_user_id),
    deleted_by INT REFERENCES app_user(app_user_id),
    created_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMPTZ
);

CREATE TABLE vendor (
    vendor_id SERIAL PRIMARY KEY,
    vendor_type_id INT NULL REFERENCES vendor_type(vendor_type_id),
    bank_merk_id INT NULL REFERENCES bank_merk(bank_merk_id),
    vendor_name VARCHAR(100) NOT NULL,
    vendor_email VARCHAR(100),
    vendor_phone VARCHAR(20),
    vendor_tin VARCHAR(100),
    city_id INT NULL REFERENCES city(city_id),
    vendor_address TEXT,
    bank_account_number VARCHAR(100),
    bank_account_name VARCHAR(100),
    is_active SMALLINT NOT NULL DEFAULT 1,
    created_by INT NOT NULL REFERENCES app_user(app_user_id),
    updated_by INT NOT NULL REFERENCES app_user(app_user_id),
    deleted_by INT REFERENCES app_user(app_user_id),
    created_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMPTZ
);

CREATE TABLE truck_type (
    truck_type_id SERIAL PRIMARY KEY,
    truck_type_name VARCHAR(100) NOT NULL,
    truck_type_description TEXT,
    truck_box_length NUMERIC(10,2) NOT NULL,
    truck_box_width NUMERIC(10,2) NOT NULL,
    truck_box_height NUMERIC(10,2) NOT NULL,
    truck_capacity NUMERIC(10,2) NOT NULL,
    is_active SMALLINT NOT NULL DEFAULT 1,
    created_by INT NOT NULL REFERENCES app_user(app_user_id),
    updated_by INT NOT NULL REFERENCES app_user(app_user_id),
    deleted_by INT REFERENCES app_user(app_user_id),
    created_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMPTZ
);

CREATE TABLE truck_merk (
    truck_merk_id SERIAL PRIMARY KEY,
    truck_merk_name VARCHAR(100) NOT NULL,
    is_active SMALLINT NOT NULL DEFAULT 1,
    created_by INT NOT NULL REFERENCES app_user(app_user_id),
    updated_by INT NOT NULL REFERENCES app_user(app_user_id),
    deleted_by INT REFERENCES app_user(app_user_id),
    created_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMPTZ
);

CREATE TABLE truck (
    truck_id SERIAL PRIMARY KEY,
    truck_type_id INT REFERENCES truck_type(truck_type_id),
    truck_merk_id INT REFERENCES truck_merk(truck_merk_id),
    driver_id INT REFERENCES app_user(app_user_id),
    vendor_id INT REFERENCES vendor(vendor_id),
    license_plate VARCHAR(20) NOT NULL,
    ownership_status_id SMALLINT NOT NULL DEFAULT 1,
    production_year SMALLINT,
    number_of_tires SMALLINT,
    is_active SMALLINT NOT NULL DEFAULT 1,
    created_by INT NOT NULL REFERENCES app_user(app_user_id),
    updated_by INT NOT NULL REFERENCES app_user(app_user_id),
    deleted_by INT REFERENCES app_user(app_user_id),
    created_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMPTZ
);


CREATE TABLE client (
    client_id SERIAL PRIMARY KEY,
    client_name VARCHAR(100) NOT NULL,
    client_email VARCHAR(100) UNIQUE,
    client_tin VARCHAR(100),
    number_of_day_until_due INT,
    city_id INT NULL REFERENCES city(city_id),
    client_address TEXT,
    operating_hours JSONB,
    is_active SMALLINT NOT NULL DEFAULT 1,
    created_by INT NOT NULL REFERENCES app_user(app_user_id),
    updated_by INT NOT NULL REFERENCES app_user(app_user_id),
    deleted_by INT REFERENCES app_user(app_user_id),
    created_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMPTZ
);

CREATE TABLE client_destination (
    client_destination_id SERIAL PRIMARY KEY,
    client_id INT REFERENCES client(client_id),
    client_destination_name VARCHAR(100) NOT NULL,
    city_id INT NULL REFERENCES city(city_id),
    client_destination_address TEXT,
    client_destination_latitude NUMERIC(9,6),
    client_destination_longitude NUMERIC(9,6),
    client_destination_map_url TEXT,
    operating_hours JSONB,
    is_active SMALLINT NOT NULL DEFAULT 1,
    created_by INT NOT NULL REFERENCES app_user(app_user_id),
    updated_by INT NOT NULL REFERENCES app_user(app_user_id),
    deleted_by INT REFERENCES app_user(app_user_id),
    created_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMPTZ
);

CREATE TABLE cargo_type (
    cargo_type_id SERIAL PRIMARY KEY,
    cargo_type_name VARCHAR(100) NOT NULL,
    cargo_type_description TEXT,
    cargo_type_grade VARCHAR(100) NOT NULL,
    is_active SMALLINT NOT NULL DEFAULT 1,
    created_by INT NOT NULL REFERENCES app_user(app_user_id),
    updated_by INT NOT NULL REFERENCES app_user(app_user_id),
    deleted_by INT REFERENCES app_user(app_user_id),
    created_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMPTZ
);

CREATE TABLE mine (
    mine_id SERIAL PRIMARY KEY,
    mine_name VARCHAR(200) NOT NULL,
    city_id INT REFERENCES city(city_id),
    mine_address TEXT,
    mine_latitude NUMERIC(9,6),
    mine_longitude NUMERIC(9,6),
    mine_map_url TEXT,
    is_active SMALLINT NOT NULL DEFAULT 1,
    created_by INT NOT NULL REFERENCES app_user(app_user_id),
    updated_by INT NOT NULL REFERENCES app_user(app_user_id),
    deleted_by INT REFERENCES app_user(app_user_id),
    created_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMPTZ
);


CREATE TABLE stockpile (
    stockpile_id SERIAL PRIMARY KEY,
    stockpile_name VARCHAR(200) NOT NULL,
    city_id INT REFERENCES city(city_id),
    stockpile_address TEXT,
    stockpile_latitude NUMERIC(9,6),
    stockpile_longitude NUMERIC(9,6),
    stockpile_map_url TEXT,
    is_active SMALLINT NOT NULL DEFAULT 1,
    created_by INT NOT NULL REFERENCES app_user(app_user_id),
    updated_by INT NOT NULL REFERENCES app_user(app_user_id),
    deleted_by INT REFERENCES app_user(app_user_id),
    created_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMPTZ
);


CREATE TABLE stockpile_cargo (
    stockpile_cargo_id SERIAL PRIMARY KEY,
    stockpile_id INT REFERENCES stockpile(stockpile_id),
    cargo_type_id INT REFERENCES cargo_type(cargo_type_id),
    capacity NUMERIC(15,2), -- kapasitas maksimum
    current_volume NUMERIC(15,2) DEFAULT 0, -- volume saat ini
    is_active SMALLINT NOT NULL DEFAULT 1,
    created_by INT NOT NULL REFERENCES app_user(app_user_id),
    updated_by INT NOT NULL REFERENCES app_user(app_user_id),
    deleted_by INT REFERENCES app_user(app_user_id),
    created_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMPTZ
);


CREATE TABLE stockpile_project (
    stockpile_project_id SERIAL PRIMARY KEY,
    mine_id INT REFERENCES mine(mine_id),
    stockpile_cargo_id INT REFERENCES stockpile_cargo(stockpile_cargo_id),
    stockpile_project_name VARCHAR(100) NOT NULL,
    volume_to_weight_convertion NUMERIC(10,2),
    price_per_cubic NUMERIC(15,2),
    price_per_ton NUMERIC(15,2),
    price_per_transport NUMERIC(15,2),
    is_active SMALLINT NOT NULL DEFAULT 1,
    created_by INT NOT NULL REFERENCES app_user(app_user_id),
    updated_by INT NOT NULL REFERENCES app_user(app_user_id),
    deleted_by INT REFERENCES app_user(app_user_id),
    created_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMPTZ
);

CREATE TABLE client_project (
    client_project_id SERIAL PRIMARY KEY,
    stockpile_cargo_id INT REFERENCES stockpile_cargo(stockpile_cargo_id),
    client_destination_id INT REFERENCES client_destination(client_destination_id),
    client_project_name VARCHAR(100) NOT NULL,
    volume_to_weight_convertion NUMERIC(10,2),
    price_per_cubic NUMERIC(15,2),
    price_per_ton NUMERIC(15,2),
    price_per_transport NUMERIC(15,2),
    is_active SMALLINT NOT NULL DEFAULT 1,
    created_by INT NOT NULL REFERENCES app_user(app_user_id),
    updated_by INT NOT NULL REFERENCES app_user(app_user_id),
    deleted_by INT REFERENCES app_user(app_user_id),
    created_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMPTZ
);


CREATE TABLE stockpile_transport (
    stockpile_transport_id SERIAL PRIMARY KEY,
    stockpile_project_id INT REFERENCES stockpile_project(stockpile_project_id),
    truck_id INT REFERENCES truck(truck_id),
    driver_id INT REFERENCES app_user(app_user_id),
    volume_to_weight_convertion NUMERIC(10,2),
    price_per_cubic NUMERIC(15,2),
    price_per_ton NUMERIC(15,2),
    price_per_transport NUMERIC(15,2),
    created_by INT NOT NULL REFERENCES app_user(app_user_id),
    updated_by INT NOT NULL REFERENCES app_user(app_user_id),
    deleted_by INT REFERENCES app_user(app_user_id),
    created_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMPTZ
);

CREATE TABLE stockpile_transport_status (
    stockpile_transport_status_id SERIAL PRIMARY KEY,
    stockpile_transport_id INT REFERENCES stockpile_transport(stockpile_transport_id),
    stockpile_transport_status_type_id SMALLINT NOT NULL,
    status_time TIMESTAMPTZ NOT NULL,
    cargo_box_length NUMERIC(10,2),
    cargo_box_width NUMERIC(10,2),
    cargo_box_height NUMERIC(10,2),
    cargo_weight NUMERIC(10,2),
    stockpile_transport_status_note TEXT,
    is_fraud SMALLINT NOT NULL DEFAULT 0,
    is_active SMALLINT NOT NULL DEFAULT 1,
    created_by INT NOT NULL REFERENCES app_user(app_user_id),
    updated_by INT NOT NULL REFERENCES app_user(app_user_id),
    deleted_by INT REFERENCES app_user(app_user_id),
    created_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMPTZ
);

CREATE TABLE stockpile_transport_photo (
    stockpile_transport_photo_id SERIAL PRIMARY KEY,
    stockpile_transport_status_id INT REFERENCES stockpile_transport_status(stockpile_transport_status_id),
    photo_url TEXT NOT NULL,
    photo_type_id SMALLINT NOT NULL,
    photo_description TEXT,
    created_by INT NOT NULL REFERENCES app_user(app_user_id),
    updated_by INT NOT NULL REFERENCES app_user(app_user_id),
    deleted_by INT REFERENCES app_user(app_user_id),
    created_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMPTZ
);


CREATE TABLE client_transport (
    client_transport_id SERIAL PRIMARY KEY,
    client_project_id INT REFERENCES client_project(client_project_id),
    truck_id INT REFERENCES truck(truck_id),
    driver_id INT REFERENCES app_user(app_user_id),
    volume_to_weight_convertion NUMERIC(10,2),
    price_per_cubic NUMERIC(15,2),
    price_per_ton NUMERIC(15,2),
    price_per_transport NUMERIC(15,2),
    created_by INT NOT NULL REFERENCES app_user(app_user_id),
    updated_by INT NOT NULL REFERENCES app_user(app_user_id),
    deleted_by INT REFERENCES app_user(app_user_id),
    created_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMPTZ
);

CREATE TABLE client_transport_status (
    client_transport_status_id SERIAL PRIMARY KEY,
    client_transport_id INT REFERENCES client_transport(client_transport_id),
    client_transport_status_type_id SMALLINT NOT NULL,
    status_time TIMESTAMPTZ NOT NULL,
    cargo_box_length NUMERIC(10,2),
    cargo_box_width NUMERIC(10,2),
    cargo_box_height NUMERIC(10,2),
    cargo_weight NUMERIC(10,2),
    client_transport_status_note TEXT,
    is_fraud SMALLINT NOT NULL DEFAULT 0,
    is_active SMALLINT NOT NULL DEFAULT 1,
    created_by INT NOT NULL REFERENCES app_user(app_user_id),
    updated_by INT NOT NULL REFERENCES app_user(app_user_id),
    deleted_by INT REFERENCES app_user(app_user_id),
    created_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMPTZ
);

CREATE TABLE client_transport_photo (
    client_transport_photo_id SERIAL PRIMARY KEY,
    client_transport_status_id INT REFERENCES client_transport_status(client_transport_status_id),
    photo_url TEXT NOT NULL,
    photo_type_id SMALLINT NOT NULL,
    photo_description TEXT,
    created_by INT NOT NULL REFERENCES app_user(app_user_id),
    updated_by INT NOT NULL REFERENCES app_user(app_user_id),
    deleted_by INT REFERENCES app_user(app_user_id),
    created_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMPTZ
);


CREATE TABLE stockpile_adjustment (
    stockpile_adjustment_id SERIAL PRIMARY KEY,
    stockpile_cargo_id INT REFERENCES stockpile_cargo(stockpile_cargo_id),
    adjustment_date DATE NOT NULL,
    stockpile_transport_id INT NULL REFERENCES stockpile_transport,
    client_transport_id INT NULL REFERENCES client_transport,
    reference_type_id SMALLINT NOT NULL,
    amount_volume NUMERIC(15,2) NOT NULL,
    reason TEXT,
    created_by INT NOT NULL REFERENCES app_user(app_user_id),
    updated_by INT NOT NULL REFERENCES app_user(app_user_id),
    deleted_by INT REFERENCES app_user(app_user_id),
    created_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMPTZ
);

CREATE TABLE stockpile_ledger (
    stockpile_ledger_id SERIAL PRIMARY KEY,
    stockpile_cargo_id INT REFERENCES stockpile_cargo(stockpile_cargo_id),
    stockpile_transport_id INT NULL REFERENCES stockpile_transport,
    client_transport_id INT NULL REFERENCES client_transport,
    adjustment_id INT NULL REFERENCES stockpile_adjustment,
    reference_type_id SMALLINT NOT NULL,
    amount_volume NUMERIC(15,2) DEFAULT 0,
    balance_volume NUMERIC(15,2) DEFAULT 0,
    created_by INT NOT NULL REFERENCES app_user(app_user_id),
    updated_by INT NOT NULL REFERENCES app_user(app_user_id),
    deleted_by INT REFERENCES app_user(app_user_id),
    created_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMPTZ
);

app_role : 3 (Admin, driver, Client)
bank_merk : 5
app_user : 20 (5 admin, 15 driver, 0 client)
province : 5 di jawa saja
city : 10 di jawa saja
vendor_type : 3
vendor : 5
truck_type : 2 (index 8, index 24)
truck_merk : 3
truck : 15 (15 truck untuk 15 driver, 12 index 8 (vendor 1), 3 index 24 (vendor 2))
client : 5
client_destination : 10
cargo_type : 5
mine : 5
stockpile : 5
stockpile_cargo : 10
stockpile_project : 5
client_project : 5
stockpile_transport : 5
stockpile_transport_status : 5
stockpile_transport_photo : 5
client_transport : 5
client_transport_status : 5
client_transport_photo : 5
stockpile_adjustment : 5
stockpile_ledger : 5




-- ==================================================
-- SEEDER FOR DATABASE
-- ==================================================

-- ==================================================
-- 1. APP ROLE (3 rows)
-- ==================================================
INSERT INTO app_role (app_role_id, app_role_type_id, app_role_name, app_role_description, is_active, created_by, updated_by) VALUES
(1, 1, 'Driver', 'Driver who operates trucks', 1, 1, 1),
(2, 2, 'Client', 'Client client', 1, 1, 1),
(3, 3, 'Owner', 'Owner with full access', 1, 1, 1),
(4, 3, 'Super Admin', 'Super Administrator with full access', 1, 1, 1),
(5, 3, 'System Admin', 'System Administrator', 1, 1, 1),
(6, 4, 'Checker', 'Checker', 1, 1, 1);

-- ==================================================
-- 2. BANK MERK (5 rows)
-- ==================================================
INSERT INTO bank_merk (bank_merk_id, bank_merk_name, bank_merk_description, is_active, created_by, updated_by) VALUES
(1, 'BCA', 'Bank Central Asia', 1, 1, 1),
(2, 'Mandiri', 'Bank Mandiri', 1, 1, 1),
(3, 'BRI', 'Bank Rakyat Indonesia', 1, 1, 1),
(4, 'BNI', 'Bank Negara Indonesia', 1, 1, 1),
(5, 'CIMB Niaga', 'Bank CIMB Niaga', 1, 1, 1);

-- ==================================================
-- 3. APP USER (20 rows)
-- 5 Admin, 15 Driver, 0 Client
-- ==================================================
-- 5 Admin (app_role_id = 3,4,5)
INSERT INTO app_user (app_user_id, app_role_id, username, password, app_user_status_id, app_user_name, 
    app_user_phone, created_by, updated_by) VALUES
(1, 3, 'itdev', '$2a$04$ROizWuC5oz08LfI0WrAOFe2bTiBF9LncdZku.TSuZb0ee//GJVl62', 1, 'IT Developer', '081234567801', 1, 1),
(2, 3, 'judijanto', '$2a$04$bmHo3UXlT7gjQCaFBO2KQ.O4QtS8qRKQJ/mUm1OrqAp9KLJ7VeljK', 1, 'Judijanto', '081234567802', 1, 1),
(3, 4, 'admin3', '$2a$04$jL6MyIcj7505HLjKCsZVRup2g1yHouGMjite0AlE6kEBOSkSEFr5C', 1, 'Admin Tiga', '081234567803', 1, 1),
(4, 4, 'admin4', '$2a$04$jL6MyIcj7505HLjKCsZVRup2g1yHouGMjite0AlE6kEBOSkSEFr5C', 1, 'Admin Empat', '081234567804', 1, 1),
(5, 4, 'admin5', '$2a$04$jL6MyIcj7505HLjKCsZVRup2g1yHouGMjite0AlE6kEBOSkSEFr5C', 1, 'Admin Lima', '081234567805', 1, 1);

-- 15 Driver (app_role_id = 1)
INSERT INTO app_user (app_user_id, app_role_id, username, password, app_user_status_id, app_user_name, 
    app_user_phone, bank_merk_id, bank_account_number, bank_account_name,
    salary_percentage, created_by, updated_by) VALUES
(6, 1, 'driver1', '$2a$04$bmHo3UXlT7gjQCaFBO2KQ.O4QtS8qRKQJ/mUm1OrqAp9KLJ7VeljK', 1, 'Driver Satu', '081234567101', 1, '1234567890', 'Driver Satu', 10.5, 1, 1),
(7, 1, 'driver2', '$2a$04$jL6MyIcj7505HLjKCsZVRup2g1yHouGMjite0AlE6kEBOSkSEFr5C', 1, 'Driver Dua', '081234567102', 1, '1234567891', 'Driver Dua', 10.5, 1, 1),
(8, 1, 'driver3', '$2a$04$jL6MyIcj7505HLjKCsZVRup2g1yHouGMjite0AlE6kEBOSkSEFr5C', 1, 'Driver Tiga', '081234567103', 2, '1234567892', 'Driver Tiga', 11.0, 1, 1),
(9, 1, 'driver4', '$2a$04$jL6MyIcj7505HLjKCsZVRup2g1yHouGMjite0AlE6kEBOSkSEFr5C', 1, 'Driver Empat', '081234567104', 2, '1234567893', 'Driver Empat', 11.0, 1, 1),
(10, 1, 'driver5', '$2a$04$jL6MyIcj7505HLjKCsZVRup2g1yHouGMjite0AlE6kEBOSkSEFr5C', 1, 'Driver Lima', '081234567105', 3, '1234567894', 'Driver Lima', 9.5, 1, 1),
(11, 1, 'driver6', '$2a$04$jL6MyIcj7505HLjKCsZVRup2g1yHouGMjite0AlE6kEBOSkSEFr5C', 1, 'Driver Enam', '081234567106', 3, '1234567895', 'Driver Enam', 9.5, 1, 1),
(12, 1, 'driver7', '$2a$04$jL6MyIcj7505HLjKCsZVRup2g1yHouGMjite0AlE6kEBOSkSEFr5C', 1, 'Driver Tujuh', '081234567107', 4, '1234567896', 'Driver Tujuh', 10.0, 1, 1),
(13, 1, 'driver8', '$2a$04$jL6MyIcj7505HLjKCsZVRup2g1yHouGMjite0AlE6kEBOSkSEFr5C', 1, 'Driver Delapan', '081234567108', 4, '1234567897', 'Driver Delapan', 10.0, 1, 1),
(14, 1, 'driver9', '$2a$04$jL6MyIcj7505HLjKCsZVRup2g1yHouGMjite0AlE6kEBOSkSEFr5C', 1, 'Driver Sembilan', '081234567109', 5, '1234567898', 'Driver Sembilan', 10.5, 1, 1),
(15, 1, 'driver10', '$2a$04$jL6MyIcj7505HLjKCsZVRup2g1yHouGMjite0AlE6kEBOSkSEFr5C', 1, 'Driver Sepuluh', '081234567110', 5, '1234567899', 'Driver Sepuluh', 10.5, 1, 1),
(16, 1, 'driver11', '$2a$04$jL6MyIcj7505HLjKCsZVRup2g1yHouGMjite0AlE6kEBOSkSEFr5C', 1, 'Driver Sebelas', '081234567111', 1, '1234567800', 'Driver Sebelas', 11.5, 1, 1),
(17, 1, 'driver12', '$2a$04$jL6MyIcj7505HLjKCsZVRup2g1yHouGMjite0AlE6kEBOSkSEFr5C', 1, 'Driver Dua Belas', '081234567112', 1, '1234567801', 'Driver Dua Belas', 11.5, 1, 1),
(18, 1, 'driver13', '$2a$04$jL6MyIcj7505HLjKCsZVRup2g1yHouGMjite0AlE6kEBOSkSEFr5C', 1, 'Driver Tiga Belas', '081234567113', 2, '1234567802', 'Driver Tiga Belas', 10.0, 1, 1),
(19, 1, 'driver14', '$2a$04$jL6MyIcj7505HLjKCsZVRup2g1yHouGMjite0AlE6kEBOSkSEFr5C', 1, 'Driver Empat Belas', '081234567114', 2, '1234567803', 'Driver Empat Belas', 10.0, 1, 1),
(20, 1, 'driver15', '$2a$04$jL6MyIcj7505HLjKCsZVRup2g1yHouGMjite0AlE6kEBOSkSEFr5C', 1, 'Driver Lima Belas', '081234567115', 3, '1234567804', 'Driver Lima Belas', 9.5, 1, 1);

-- 5 Admin (app_role_id = 6)
INSERT INTO app_user (app_user_id, app_role_id, username, password, app_user_status_id, app_user_name, 
    app_user_phone, created_by, updated_by) VALUES
(21, 6, 'itchecker', '$2a$04$ROizWuC5oz08LfI0WrAOFe2bTiBF9LncdZku.TSuZb0ee//GJVl62', 1, 'IT Checker', '081234567801', 1, 1),
(22, 6, 'checker', '$2a$04$bmHo3UXlT7gjQCaFBO2KQ.O4QtS8qRKQJ/mUm1OrqAp9KLJ7VeljK', 1, 'Checker', '081234567801', 1, 1);


-- ==================================================
-- 4. PROVINCE (5 rows - Jawa only)
-- ==================================================
INSERT INTO province (province_id, province_name, province_real_name, is_active, created_by, updated_by) VALUES
(1, 'Jawa Barat', 'West Java', 1, 1, 1),
(2, 'Jawa Tengah', 'Central Java', 1, 1, 1),
(3, 'Jawa Timur', 'East Java', 1, 1, 1),
(4, 'DKI Jakarta', 'Jakarta', 1, 1, 1),
(5, 'Banten', 'Banten', 1, 1, 1);

-- ==================================================
-- 5. CITY (10 rows - Jawa only)
-- ==================================================
INSERT INTO city (city_id, province_id, city_name, is_active, created_by, updated_by) VALUES
(1, 1, 'Bandung', 1, 1, 1),
(2, 1, 'Bekasi', 1, 1, 1),
(3, 1, 'Bogor', 1, 1, 1),
(4, 2, 'Semarang', 1, 1, 1),
(5, 2, 'Solo', 1, 1, 1),
(6, 3, 'Surabaya', 1, 1, 1),
(7, 3, 'Malang', 1, 1, 1),
(8, 4, 'Jakarta Pusat', 1, 1, 1),
(9, 4, 'Jakarta Selatan', 1, 1, 1),
(10, 5, 'Tangerang', 1, 1, 1);

-- ==================================================
-- 6. VENDOR TYPE (3 rows)
-- ==================================================
INSERT INTO vendor_type (vendor_type_id, vendor_type_name, vendor_type_description, is_active, created_by, updated_by) VALUES
(1, 'Transport Provider', 'Menyediakan jasa transportasi truk', 1, 1, 1),
(2, 'Maintenance Provider', 'Menyediakan jasa perawatan truk', 1, 1, 1),
(3, 'Fuel Provider', 'Menyediakan bahan bakar', 1, 1, 1);

-- ==================================================
-- 7. VENDOR (5 rows)
-- ==================================================
INSERT INTO vendor (vendor_id, vendor_type_id, bank_merk_id, vendor_name, vendor_email, vendor_phone, 
    vendor_tin, city_id, vendor_address, bank_account_number, bank_account_name, 
    is_active, created_by, updated_by) VALUES
(1, 1, 1, 'PT Transport Sejahtera', 'transport@sejahtera.com', '0215555001', '01.123.456.7-001.000', 
 1, 'Jl. Asia Afrika No. 1, Bandung', '1234567001', 'PT Transport Sejahtera', 1, 1, 1),
(2, 1, 2, 'CV Logistik Nusantara', 'logistik@nusantara.com', '0215555002', '01.123.456.7-002.000',
 8, 'Jl. Thamrin No. 2, Jakarta', '1234567002', 'CV Logistik Nusantara', 1, 1, 1),
(3, 1, 3, 'PT Angkutan Cepat', 'angkutan@cepat.com', '0215555003', '01.123.456.7-003.000',
 6, 'Jl. Tunjungan No. 3, Surabaya', '1234567003', 'PT Angkutan Cepat', 1, 1, 1),
(4, 2, 4, 'PT Bengkel Utama', 'bengkel@utama.com', '0215555004', '01.123.456.7-004.000',
 4, 'Jl. Pandanaran No. 4, Semarang', '1234567004', 'PT Bengkel Utama', 1, 1, 1),
(5, 3, 5, 'PT Bahan Bakar Jaya', 'bbm@jaya.com', '0215555005', '01.123.456.7-005.000',
 10, 'Jl. Sudirman No. 5, Tangerang', '1234567005', 'PT Bahan Bakar Jaya', 1, 1, 1);

-- ==================================================
-- 8. TRUCK TYPE (2 rows)
-- ==================================================
INSERT INTO truck_type (truck_type_id, truck_type_name, truck_type_description, 
    truck_box_length, truck_box_width, truck_box_height, truck_capacity,
    is_active, created_by, updated_by) VALUES
(1, 'Index 8', 'Truk ringan 4 roda, kapasitas 8 ton', 6.0, 2.2, 2.0, 8.0, 1, 1, 1),
(2, 'Index 24', 'Truk sedang 6 roda, kapasitas 24 ton', 8.0, 2.4, 2.2, 24.0, 1, 1, 1);

-- ==================================================
-- 9. TRUCK MERK (3 rows)
-- ==================================================
INSERT INTO truck_merk (truck_merk_id, truck_merk_name, is_active, created_by, updated_by) VALUES
(1, 'Hino', 1, 1, 1),
(2, 'Mitsubishi', 1, 1, 1),
(3, 'Isuzu', 1, 1, 1);

-- ==================================================
-- 10. TRUCK (15 rows)
-- 15 truck untuk 15 driver
-- 12 truck index 8 (vendor 1), 3 truck index 24 (vendor 2)
-- ==================================================
-- 12 Truk Index 8 (vendor_id = 1)
INSERT INTO truck (truck_id, truck_type_id, truck_merk_id, driver_id, vendor_id, license_plate,
    ownership_status_id, production_year, number_of_tires, is_active,
    created_by, updated_by) VALUES
(1, 1, 1, 6, 1, 'B 1234 ABC', 1, 2020, 4, 1, 1, 1),
(2, 1, 1, 7, 1, 'B 1235 ABC', 1, 2020, 4, 1, 1, 1),
(3, 1, 2, 8, 1, 'B 1236 ABC', 1, 2021, 4, 1, 1, 1),
(4, 1, 2, 9, 1, 'B 1237 ABC', 1, 2021, 4, 1, 1, 1),
(5, 1, 3, 10, 1, 'B 1238 ABC', 1, 2022, 4, 1, 1, 1),
(6, 1, 3, 11, 1, 'B 1239 ABC', 1, 2022, 4, 1, 1, 1),
(7, 1, 1, 12, 1, 'B 1240 ABC', 1, 2020, 4, 1, 1, 1),
(8, 1, 1, 13, 1, 'B 1241 ABC', 1, 2020, 4, 1, 1, 1),
(9, 1, 2, 14, 1, 'B 1242 ABC', 1, 2021, 4, 1, 1, 1),
(10, 1, 2, 15, 1, 'B 1243 ABC', 1, 2021, 4, 1, 1, 1),
(11, 1, 3, 16, 1, 'B 1244 ABC', 1, 2022, 4, 1, 1, 1),
(12, 1, 3, 17, 1, 'B 1245 ABC', 1, 2022, 4, 1, 1, 1);

-- 3 Truk Index 24 (vendor_id = 2)
INSERT INTO truck (truck_id, truck_type_id, truck_merk_id, driver_id, vendor_id, license_plate,
    ownership_status_id, production_year, number_of_tires, is_active,
    created_by, updated_by) VALUES
(13, 2, 1, 18, 2, 'B 5678 XYZ', 1, 2019, 6, 1, 1, 1),
(14, 2, 2, 19, 2, 'B 5679 XYZ', 1, 2019, 6, 1, 1, 1),
(15, 2, 3, 20, 2, 'B 5680 XYZ', 1, 2020, 6, 1, 1, 1);

-- ==================================================
-- 11. CLIENT (5 rows)
-- ==================================================
INSERT INTO client (client_id, client_name, client_email, client_tin, number_of_day_until_due,
    city_id, client_address, operating_hours, is_active, created_by, updated_by) VALUES
(1, 'PT Maju Bersama', 'maju@bersama.com', '02.123.456.7-001.000', 30, 
 1, 'Jl. Diponegoro No. 10, Bandung', '{"monday":"08:00-17:00","friday":"08:00-16:30"}', 1, 1, 1),
(2, 'CV Sukses Abadi', 'sukses@abadi.com', '02.123.456.7-002.000', 30,
 8, 'Jl. Gatot Subroto No. 20, Jakarta', '{"monday":"09:00-18:00","saturday":"09:00-14:00"}', 1, 1, 1),
(3, 'PT Karya Mandiri', 'karya@mandiri.com', '02.123.456.7-003.000', 45,
 6, 'Jl. Darmo No. 30, Surabaya', '{"monday":"08:00-17:00","thursday":"08:00-17:00"}', 1, 1, 1),
(4, 'UD Jaya Makmur', 'jaya@makmur.com', '02.123.456.7-004.000', 30,
 4, 'Jl. Pahlawan No. 40, Semarang', '{"monday":"08:00-16:00","wednesday":"08:00-16:00"}', 1, 1, 1),
(5, 'PT Indah Karya', 'indah@karya.com', '02.123.456.7-005.000', 60,
 10, 'Jl. Veteran No. 50, Tangerang', '{"monday":"08:30-17:30","tuesday":"08:30-17:30"}', 1, 1, 1);

-- ==================================================
-- 12. CLIENT DESTINATION (10 rows)
-- ==================================================
INSERT INTO client_destination (client_destination_id, client_id, client_destination_name, city_id, client_destination_address,
    client_destination_latitude, client_destination_longitude, client_destination_map_url,
    operating_hours, is_active, created_by, updated_by) VALUES
(1, 1, 'Gudang Bandung 1', 1, 'Jl. Cibaduyut No. 1, Bandung', -6.921700, 107.607100, 'https://maps.google.com/?q=-6.921700,107.607100',
 '{"monday":"08:00-17:00","saturday":"08:00-12:00"}', 1, 1, 1),
(2, 1, 'Gudang Bandung 2', 1, 'Jl. Cimahi No. 2, Bandung', -6.931700, 107.617100, 'https://maps.google.com/?q=-6.931700,107.617100',
 '{"monday":"08:00-17:00","saturday":"08:00-12:00"}', 1, 1, 1),
(3, 2, 'Gudang Jakarta 1', 8, 'Jl. Cakung No. 1, Jakarta', -6.200000, 106.900000, 'https://maps.google.com/?q=-6.200000,106.900000',
 '{"monday":"09:00-18:00","saturday":"09:00-14:00"}', 1, 1, 1),
(4, 2, 'Gudang Jakarta 2', 9, 'Jl. Kebayoran No. 2, Jakarta', -6.250000, 106.800000, 'https://maps.google.com/?q=-6.250000,106.800000',
 '{"monday":"09:00-18:00","saturday":"09:00-14:00"}', 1, 1, 1),
(5, 3, 'Gudang Surabaya 1', 6, 'Jl. Rungkut No. 1, Surabaya', -7.250000, 112.750000, 'https://maps.google.com/?q=-7.250000,112.750000',
 '{"monday":"08:00-17:00","thursday":"08:00-17:00"}', 1, 1, 1),
(6, 3, 'Gudang Surabaya 2', 7, 'Jl. Malang No. 2, Malang', -7.950000, 112.600000, 'https://maps.google.com/?q=-7.950000,112.600000',
 '{"monday":"08:00-17:00","thursday":"08:00-17:00"}', 1, 1, 1),
(7, 4, 'Gudang Semarang 1', 4, 'Jl. Tlogosari No. 1, Semarang', -7.000000, 110.400000, 'https://maps.google.com/?q=-7.000000,110.400000',
 '{"monday":"08:00-16:00","wednesday":"08:00-16:00"}', 1, 1, 1),
(8, 4, 'Gudang Semarang 2', 5, 'Jl. Solo No. 2, Solo', -7.550000, 110.800000, 'https://maps.google.com/?q=-7.550000,110.800000',
 '{"monday":"08:00-16:00","wednesday":"08:00-16:00"}', 1, 1, 1),
(9, 5, 'Gudang Tangerang 1', 10, 'Jl. Cikupa No. 1, Tangerang', -6.200000, 106.500000, 'https://maps.google.com/?q=-6.200000,106.500000',
 '{"monday":"08:30-17:30","tuesday":"08:30-17:30"}', 1, 1, 1),
(10, 5, 'Gudang Tangerang 2', 10, 'Jl. Balaraja No. 2, Tangerang', -6.250000, 106.450000, 'https://maps.google.com/?q=-6.250000,106.450000',
 '{"monday":"08:30-17:30","tuesday":"08:30-17:30"}', 1, 1, 1);

-- ==================================================
-- 13. CARGO TYPE (5 rows)
-- ==================================================
INSERT INTO cargo_type (cargo_type_id, cargo_type_name, cargo_type_description, cargo_type_grade, is_active, created_by, updated_by) VALUES
(1, 'Batu Bara', 'Batubara kalori tinggi', 'A', 1, 1, 1),
(2, 'Batu Bara', 'Batubara kalori sedang', 'B', 1, 1, 1),
(3, 'Batu Bara', 'Batubara kalori rendah', 'C', 1, 1, 1),
(4, 'Bijih Besi', 'Bijih besi kadar tinggi', 'Premium', 1, 1, 1),
(5, 'Nikel', 'Bijih nikel kadar tinggi', 'Grade 1', 1, 1, 1);

-- ==================================================
-- 14. MINE (5 rows)
-- ==================================================
INSERT INTO mine (mine_id, mine_name, city_id, mine_address, mine_latitude, mine_longitude, mine_map_url,
    is_active, created_by, updated_by) VALUES
(1, 'Tambang Batubara Kalimantan 1', 1, 'Jl. Pertambangan No. 1, Kalimantan', -1.250000, 116.000000, 'https://maps.google.com/?q=-1.250000,116.000000', 1, 1, 1),
(2, 'Tambang Batubara Kalimantan 2', 1, 'Jl. Pertambangan No. 2, Kalimantan', -1.350000, 116.100000, 'https://maps.google.com/?q=-1.350000,116.100000', 1, 1, 1),
(3, 'Tambang Bijih Besi Sulawesi', 6, 'Jl. Pertambangan No. 3, Sulawesi', -3.500000, 121.000000, 'https://maps.google.com/?q=-3.500000,121.000000', 1, 1, 1),
(4, 'Tambang Nikel Sulawesi', 6, 'Jl. Pertambangan No. 4, Sulawesi', -3.600000, 121.100000, 'https://maps.google.com/?q=-3.600000,121.100000', 1, 1, 1),
(5, 'Tambang Batubara Sumatra', 8, 'Jl. Pertambangan No. 5, Sumatra', -2.500000, 104.000000, 'https://maps.google.com/?q=-2.500000,104.000000', 1, 1, 1);

-- ==================================================
-- 15. STOCKPILE (5 rows)
-- ==================================================
INSERT INTO stockpile (stockpile_id, stockpile_name, city_id, stockpile_address, stockpile_latitude, stockpile_longitude, stockpile_map_url,
    is_active, created_by, updated_by) VALUES
(1, 'Stockpile Batubara Banjarmasin', 1, 'Jl. Pelabuhan No. 1, Banjarmasin', -3.300000, 114.600000, 'https://maps.google.com/?q=-3.300000,114.600000', 1, 1, 1),
(2, 'Stockpile Batubara Balikpapan', 1, 'Jl. Pelabuhan No. 2, Balikpapan', -1.200000, 116.800000, 'https://maps.google.com/?q=-1.200000,116.800000', 1, 1, 1),
(3, 'Stockpile Bijih Besi Makassar', 6, 'Jl. Pelabuhan No. 3, Makassar', -5.100000, 119.400000, 'https://maps.google.com/?q=-5.100000,119.400000', 1, 1, 1),
(4, 'Stockpile Nikel Kendari', 6, 'Jl. Pelabuhan No. 4, Kendari', -3.900000, 122.500000, 'https://maps.google.com/?q=-3.900000,122.500000', 1, 1, 1),
(5, 'Stockpile Batubara Palembang', 8, 'Jl. Pelabuhan No. 5, Palembang', -2.900000, 104.700000, 'https://maps.google.com/?q=-2.900000,104.700000', 1, 1, 1);

-- ==================================================
-- 16. STOCKPILE CARGO (10 rows)
-- ==================================================
INSERT INTO stockpile_cargo (stockpile_cargo_id, stockpile_id, cargo_type_id, capacity, current_volume, is_active, created_by, updated_by) VALUES
(1, 1, 1, 100000.00, 45000.00, 1, 1, 1),
(2, 1, 2, 80000.00, 35000.00, 1, 1, 1),
(3, 2, 1, 120000.00, 60000.00, 1, 1, 1),
(4, 2, 2, 90000.00, 40000.00, 1, 1, 1),
(5, 2, 3, 70000.00, 25000.00, 1, 1, 1),
(6, 3, 4, 50000.00, 20000.00, 1, 1, 1),
(7, 4, 5, 40000.00, 15000.00, 1, 1, 1),
(8, 5, 1, 80000.00, 30000.00, 1, 1, 1),
(9, 5, 2, 60000.00, 20000.00, 1, 1, 1),
(10, 5, 3, 50000.00, 15000.00, 1, 1, 1);

-- ==================================================
-- 17. STOCKPILE PROJECT (5 rows)
-- ==================================================
INSERT INTO stockpile_project (stockpile_project_id, mine_id, stockpile_cargo_id, stockpile_project_name, volume_to_weight_convertion, price_per_cubic, price_per_ton, price_per_transport,
    is_active, created_by, updated_by) VALUES
(1, 1, 1, 'Proyek Batubara Kalimantan - Grade A', 1.25, 250000, 0, 0, 1, 1, 1),
(2, 1, 2, 'Proyek Batubara Kalimantan - Grade B', 1.20, 250000, 0, 0, 1, 1, 1),
(3, 2, 3, 'Proyek Batubara Balikpapan - Grade A', 1.25, 250000, 0, 0, 1, 1, 1),
(4, 3, 6, 'Proyek Bijih Besi Sulawesi', 2.50, 250000, 0, 0, 1, 1, 1),
(5, 4, 7, 'Proyek Nikel Sulawesi', 2.75, 250000, 0, 0, 1, 1, 1);

-- ==================================================
-- 18. CLIENT PROJECT (5 rows)
-- ==================================================
INSERT INTO client_project (client_project_id, stockpile_cargo_id, client_destination_id, client_project_name, volume_to_weight_convertion, price_per_cubic, price_per_ton, price_per_transport,
    is_active, created_by, updated_by) VALUES
(1, 1, 1, 'Pengiriman Batubara ke Bandung - Maju Bersama', 2.00, 0, 260000, 2500000, 1, 1, 1),
(2, 3, 3, 'Pengiriman Batubara ke Jakarta - Sukses Abadi', 2.00, 0, 260000, 2500000, 1, 1, 1),
(3, 6, 5, 'Pengiriman Bijih Besi ke Surabaya - Karya Mandiri', 2.00, 0, 260000, 2500000, 1, 1, 1),
(4, 7, 7, 'Pengiriman Nikel ke Semarang - Jaya Makmur', 2.00, 0, 260000, 2500000, 1, 1, 1),
(5, 8, 9, 'Pengiriman Batubara ke Tangerang - Indah Karya', 2.00, 0, 260000, 2500000, 1, 1, 1);


-- ==================================================
-- SEEDER FOR STOCKPILE TRANSPORT, STATUS, PHOTOS
-- CLIENT TRANSPORT, STATUS, PHOTOS
-- STOCKPILE ADJUSTMENT, LEDGER
-- ==================================================

-- ==================================================
-- 19. STOCKPILE TRANSPORT (5 rows)
-- Transportasi dari tambang ke stockpile
-- ==================================================
INSERT INTO stockpile_transport (
    stockpile_transport_id, 
    stockpile_project_id, 
    truck_id, 
    driver_id, 
    volume_to_weight_convertion, 
    price_per_cubic, 
    price_per_ton, 
    price_per_transport,
    created_by, 
    updated_by,
    created_at, 
    updated_at
) VALUES
-- Proyek 1 (Batubara Grade A) - Truck Index 8
(1, 1, 1, 6, 1.25, 250000, NULL, 2000000, 1, 1, NOW() - INTERVAL '5 days', NOW() - INTERVAL '5 days'),
-- Proyek 1 (Batubara Grade A) - Truck Index 8 lainnya
(2, 1, 2, 7, 1.25, 250000, NULL, 2000000, 1, 1, NOW() - INTERVAL '4 days', NOW() - INTERVAL '4 days'),
-- Proyek 2 (Batubara Grade B) - Truck Index 8
(3, 2, 3, 8, 1.20, 250000, NULL, 2000000, 1, 1, NOW() - INTERVAL '3 days', NOW() - INTERVAL '3 days'),
-- Proyek 3 (Batubara Balikpapan) - Truck Index 24
(4, 3, 13, 18, 1.25, 250000, NULL, 3000000, 1, 1, NOW() - INTERVAL '2 days', NOW() - INTERVAL '2 days'),
-- Proyek 4 (Bijih Besi) - Truck Index 24
(5, 4, 14, 19, 2.50, 250000, NULL, 3500000, 1, 1, NOW() - INTERVAL '1 day', NOW() - INTERVAL '1 day');

-- ==================================================
-- 20. STOCKPILE TRANSPORT STATUS (5 rows)
-- Status perjalanan untuk setiap transport
-- ==================================================
INSERT INTO stockpile_transport_status (
    stockpile_transport_status_id,
    stockpile_transport_id,
    stockpile_transport_status_type_id,
    status_time,
    cargo_box_length,
    cargo_box_width,
    cargo_box_height,
    cargo_weight,
    stockpile_transport_status_note,
    is_fraud,
    is_active,
    created_by,
    updated_by,
    created_at,
    updated_at
) VALUES
-- Status untuk transport 1: Berangkat dari tambang
(1, 1, 1,
 NOW() - INTERVAL '5 days' + INTERVAL '1 hour',
 6.0, 2.2, 2.0, 15.0,
 'Berangkat dari tambang dengan muatan batubara Grade A',
 0, 1, 1, 1,
 NOW() - INTERVAL '5 days' + INTERVAL '1 hour',
 NOW() - INTERVAL '5 days' + INTERVAL '1 hour'),

-- Status untuk transport 1: Tiba di stockpile
(2, 1, 2,
 NOW() - INTERVAL '5 days' + INTERVAL '5 hours',
 6.0, 2.2, 2.0, 15.0,
 'Tiba di stockpile, muatan sesuai',
 0, 1, 1, 1,
 NOW() - INTERVAL '5 days' + INTERVAL '5 hours',
 NOW() - INTERVAL '5 days' + INTERVAL '5 hours'),

-- Status untuk transport 2: Berangkat dari tambang
(3, 2, 1,
 NOW() - INTERVAL '4 days' + INTERVAL '2 hours',
 6.0, 2.2, 2.0, 14.5,
 'Berangkat dengan muatan batubara Grade A',
 0, 1, 1, 1,
 NOW() - INTERVAL '4 days' + INTERVAL '2 hours',
 NOW() - INTERVAL '4 days' + INTERVAL '2 hours'),

-- Status untuk transport 4: Berangkat dengan truk besar
(4, 4, 1,
 NOW() - INTERVAL '2 days' + INTERVAL '3 hours',
 8.0, 2.4, 2.2, 20.0,
 'Berangkat dengan truk Index 24, muatan penuh',
 0, 1, 1, 1,
 NOW() - INTERVAL '2 days' + INTERVAL '3 hours',
 NOW() - INTERVAL '2 days' + INTERVAL '3 hours'),

-- Status untuk transport 5: Bermasalah di jalan
(5, 5, 3,
 NOW() - INTERVAL '1 day' + INTERVAL '4 hours',
 8.0, 2.4, 2.2, 18.0,
 'Terjadi delay karena hujan deras',
 0, 1, 1, 1,
 NOW() - INTERVAL '1 day' + INTERVAL '4 hours',
 NOW() - INTERVAL '1 day' + INTERVAL '4 hours');
-- ==================================================
-- 21. STOCKPILE TRANSPORT PHOTO (5 rows)
-- Foto dokumentasi untuk setiap status
-- ==================================================
INSERT INTO stockpile_transport_photo (
    stockpile_transport_photo_id,
    stockpile_transport_status_id,
    photo_url,
    photo_type_id,
    photo_description,
    created_by,
    updated_by,
    created_at,
    updated_at
) VALUES
-- Foto untuk status 1 (Berangkat)
(1, 1, 'https://storage.example.com/photos/stockpile/transport/1/departure_1.jpg', 1, 'Foto truk saat berangkat dari tambang', 1, 1, NOW() - INTERVAL '5 days' + INTERVAL '1 hour', NOW() - INTERVAL '5 days' + INTERVAL '1 hour'),
(2, 1, 'https://storage.example.com/photos/stockpile/transport/1/departure_2.jpg', 2, 'Foto muatan batubara', 1, 1, NOW() - INTERVAL '5 days' + INTERVAL '1 hour', NOW() - INTERVAL '5 days' + INTERVAL '1 hour'),
-- Foto untuk status 2 (Tiba)
(3, 2, 'https://storage.example.com/photos/stockpile/transport/1/arrival_1.jpg', 3, 'Foto truk saat tiba di stockpile', 1, 1, NOW() - INTERVAL '5 days' + INTERVAL '5 hours', NOW() - INTERVAL '5 days' + INTERVAL '5 hours'),
-- Foto untuk status 4 (Berangkat truk besar)
(4, 4, 'https://storage.example.com/photos/stockpile/transport/4/departure_1.jpg', 1, 'Truk Index 24 siap berangkat', 1, 1, NOW() - INTERVAL '2 days' + INTERVAL '3 hours', NOW() - INTERVAL '2 days' + INTERVAL '3 hours'),
-- Foto untuk status 5 (Masalah)
(5, 5, 'https://storage.example.com/photos/stockpile/transport/5/issue_1.jpg', 5, 'Kondisi jalan banjir', 1, 1, NOW() - INTERVAL '1 day' + INTERVAL '4 hours', NOW() - INTERVAL '1 day' + INTERVAL '4 hours');

-- ==================================================
-- 22. CLIENT TRANSPORT (5 rows)
-- Transportasi dari stockpile ke client
-- ==================================================
INSERT INTO client_transport (
    client_transport_id,
    client_project_id,
    truck_id,
    driver_id,
    volume_to_weight_convertion,
    price_per_cubic,
    price_per_ton,
    price_per_transport,
    created_by,
    updated_by,
    created_at,
    updated_at
) VALUES
-- Client Project 1 (ke Bandung) - Truck Index 8
(1, 1, 4, 9, 2.00, NULL, 260000, 2500000, 1, 1, NOW() - INTERVAL '4 days', NOW() - INTERVAL '4 days'),
-- Client Project 1 (ke Bandung) - Truck Index 8 lainnya
(2, 1, 5, 10, 2.00, NULL, 260000, 2500000, 1, 1, NOW() - INTERVAL '3 days', NOW() - INTERVAL '3 days'),
-- Client Project 2 (ke Jakarta) - Truck Index 8
(3, 2, 7, 12, 2.00, NULL, 260000, 2500000, 1, 1, NOW() - INTERVAL '3 days', NOW() - INTERVAL '3 days'),
-- Client Project 3 (ke Surabaya) - Truck Index 24
(4, 3, 15, 20, 2.00, NULL, 260000, 3500000, 1, 1, NOW() - INTERVAL '2 days', NOW() - INTERVAL '2 days'),
-- Client Project 5 (ke Tangerang) - Truck Index 8
(5, 5, 11, 16, 2.00, NULL, 260000, 2500000, 1, 1, NOW() - INTERVAL '1 day', NOW() - INTERVAL '1 day');

-- ==================================================
-- 23. CLIENT TRANSPORT STATUS (5 rows)
-- Status perjalanan untuk client transport
-- ==================================================
INSERT INTO client_transport_status (
    client_transport_status_id,
    client_transport_id,
    client_transport_status_type_id,
    status_time,
    cargo_box_length,
    cargo_box_width,
    cargo_box_height,
    cargo_weight,
    client_transport_status_note,
    is_fraud,
    is_active,
    created_by,
    updated_by,
    created_at,
    updated_at
) VALUES
-- Status untuk client transport 1: Ambil dari stockpile
(1, 1, 1,
 NOW() - INTERVAL '4 days' + INTERVAL '2 hours',
 6.0, 2.2, 2.0, 15.5,
 'Mengambil muatan dari stockpile Batubara Banjarmasin',
 0, 1, 1, 1,
 NOW() - INTERVAL '4 days' + INTERVAL '2 hours',
 NOW() - INTERVAL '4 days' + INTERVAL '2 hours'),

-- Status untuk client transport 1: Dalam perjalanan
(2, 1, 2,
 NOW() - INTERVAL '4 days' + INTERVAL '6 hours',
 6.0, 2.2, 2.0, 15.5,
 'Dalam perjalanan menuju Bandung',
 0, 1, 1, 1,
 NOW() - INTERVAL '4 days' + INTERVAL '6 hours',
 NOW() - INTERVAL '4 days' + INTERVAL '6 hours'),

-- Status untuk client transport 2: Ambil dari stockpile
(3, 2, 1,
 NOW() - INTERVAL '3 days' + INTERVAL '3 hours',
 6.0, 2.2, 2.0, 14.0,
 'Mengambil muatan dari stockpile',
 0, 1, 1, 1,
 NOW() - INTERVAL '3 days' + INTERVAL '3 hours',
 NOW() - INTERVAL '3 days' + INTERVAL '3 hours'),

-- Status untuk client transport 4: Ambil dengan truk besar
(4, 4, 1,
 NOW() - INTERVAL '2 days' + INTERVAL '4 hours',
 8.0, 2.4, 2.2, 20.0,
 'Mengambil muatan bijih besi dengan truk besar',
 0, 1, 1, 1,
 NOW() - INTERVAL '2 days' + INTERVAL '4 hours',
 NOW() - INTERVAL '2 days' + INTERVAL '4 hours'),

-- Status untuk client transport 5: Tiba di tujuan
(5, 5, 3,
 NOW() - INTERVAL '1 day' + INTERVAL '8 hours',
 6.0, 2.2, 2.0, 15.5,
 'Tiba di gudang client Tangerang, muatan diterima',
 0, 1, 1, 1,
 NOW() - INTERVAL '1 day' + INTERVAL '8 hours',
 NOW() - INTERVAL '1 day' + INTERVAL '8 hours');
-- ==================================================
-- 24. CLIENT TRANSPORT PHOTO (5 rows)
-- Foto dokumentasi client transport
-- ==================================================
INSERT INTO client_transport_photo (
    client_transport_photo_id,
    client_transport_status_id,
    photo_url,
    photo_type_id,
    photo_description,
    created_by,
    updated_by,
    created_at,
    updated_at
) VALUES
-- Foto untuk client status 1
(1, 1, 'https://storage.example.com/photos/client/transport/1/loading_1.jpg', 1, 'Proses loading di stockpile', 1, 1, NOW() - INTERVAL '4 days' + INTERVAL '2 hours', NOW() - INTERVAL '4 days' + INTERVAL '2 hours'),
-- Foto untuk client status 2
(2, 2, 'https://storage.example.com/photos/client/transport/1/on_road_1.jpg', 2, 'Truk dalam perjalanan', 1, 1, NOW() - INTERVAL '4 days' + INTERVAL '6 hours', NOW() - INTERVAL '4 days' + INTERVAL '6 hours'),
-- Foto untuk client status 3
(3, 3, 'https://storage.example.com/photos/client/transport/2/loading_1.jpg', 1, 'Loading batubara', 1, 1, NOW() - INTERVAL '3 days' + INTERVAL '3 hours', NOW() - INTERVAL '3 days' + INTERVAL '3 hours'),
-- Foto untuk client status 4
(4, 4, 'https://storage.example.com/photos/client/transport/4/loading_1.jpg', 1, 'Loading bijih besi dengan truk besar', 1, 1, NOW() - INTERVAL '2 days' + INTERVAL '4 hours', NOW() - INTERVAL '2 days' + INTERVAL '4 hours'),
-- Foto untuk client status 5 (tiba)
(5, 5, 'https://storage.example.com/photos/client/transport/5/arrival_1.jpg', 3, 'Tiba dan unloading di gudang client', 1, 1, NOW() - INTERVAL '1 day' + INTERVAL '8 hours', NOW() - INTERVAL '1 day' + INTERVAL '8 hours');

-- ==================================================
-- 25. STOCKPILE ADJUSTMENT (5 rows)
-- Penyesuaian stok di stockpile
-- ==================================================
INSERT INTO stockpile_adjustment (
    stockpile_adjustment_id,
    stockpile_cargo_id,
    adjustment_date,
    stockpile_transport_id,
    client_transport_id,
    reference_type_id,
    amount_volume,
    reason,
    created_by,
    updated_by,
    created_at,
    updated_at
) VALUES
-- Adjustment dari mine transport (incoming)
(1, 1, CURRENT_DATE - INTERVAL '4 days', 1, NULL, 1, 15.00, 'Penerimaan dari tambang - Stockpile Project 1', 1, 1, NOW() - INTERVAL '4 days', NOW() - INTERVAL '4 days'),
(2, 3, CURRENT_DATE - INTERVAL '3 days', 3, NULL, 1, 18.00, 'Penerimaan dari tambang - Stockpile Project 2', 1, 1, NOW() - INTERVAL '3 days', NOW() - INTERVAL '3 days'),
-- Adjustment dari client transport (outgoing)
(3, 1, CURRENT_DATE - INTERVAL '3 days', NULL, 1, 2, -12.00, 'Pengiriman ke client - Client Project 1', 1, 1, NOW() - INTERVAL '3 days', NOW() - INTERVAL '3 days'),
(4, 6, CURRENT_DATE - INTERVAL '2 days', NULL, 4, 2, -20.00, 'Pengiriman bijih besi ke client Surabaya', 1, 1, NOW() - INTERVAL '2 days', NOW() - INTERVAL '2 days'),
-- Adjustment manual (stock opname)
(5, 2, CURRENT_DATE, NULL, NULL, 3, -2.50, 'Stock opname - penyesuaian karena penyusutan', 1, 1, NOW(), NOW());

-- ==================================================
-- 26. STOCKPILE LEDGER (5 rows)
-- Buku besar untuk tracking stok
-- ==================================================
-- Untuk stockpile_cargo_id = 1 (Batubara Grade A di Stockpile 1)
-- Saldo awal: 45000 (dari seeder stockpile_cargo)
INSERT INTO stockpile_ledger (
    stockpile_ledger_id,
    stockpile_cargo_id,
    stockpile_transport_id,
    client_transport_id,
    adjustment_id,
    reference_type_id,
    amount_volume,
    balance_volume,
    created_by,
    updated_by,
    created_at,
    updated_at
) VALUES
-- Entry 1: Penerimaan dari tambang (stockpile_transport_id = 1)
(1, 1, 1, NULL, 1, 1, 15.00, 45015.00, 1, 1, NOW() - INTERVAL '4 days', NOW() - INTERVAL '4 days'),
-- Entry 2: Pengiriman ke client (client_transport_id = 1)
(2, 1, NULL, 1, 3, 2, -12.00, 45003.00, 1, 1, NOW() - INTERVAL '3 days', NOW() - INTERVAL '3 days'),

-- Untuk stockpile_cargo_id = 3 (Batubara Grade A di Stockpile 2)
-- Saldo awal: 60000
(3, 3, 3, NULL, 2, 1, 18.00, 60018.00, 1, 1, NOW() - INTERVAL '3 days', NOW() - INTERVAL '3 days'),

-- Untuk stockpile_cargo_id = 6 (Bijih Besi di Stockpile 3)
-- Saldo awal: 20000
(4, 6, NULL, 4, 4, 2, -20.00, 19980.00, 1, 1, NOW() - INTERVAL '2 days', NOW() - INTERVAL '2 days'),

-- Untuk stockpile_cargo_id = 2 (Batubara Grade B di Stockpile 1)
-- Saldo awal: 35000
(5, 2, NULL, NULL, 5, 3, -2.50, 34997.50, 1, 1, NOW(), NOW());


INSERT INTO mine_client_project (
    mine_client_project_id, mine_id, cargo_type_id, client_destination_id,
    mine_client_project_name,
    volume_to_weight_convertion,
    buy_price_per_cubic, buy_price_per_ton, buy_price_per_transport,
    sell_price_per_cubic, sell_price_per_ton, sell_price_per_transport,
    is_active, created_by, updated_by, created_at, updated_at
) VALUES
(1, 1, 1, 1, 'Project A', 1.50, 10000, 15000, 200000, 12000, 17000, 220000, 1, 1, 1, NOW(), NOW()),
(2, 2, 2, 2, 'Project B', 1.60, 11000, 16000, 210000, 13000, 18000, 230000, 1, 1, 1, NOW(), NOW()),
(3, 3, 3, 3, 'Project C', 1.70, 12000, 17000, 220000, 14000, 19000, 240000, 1, 1, 1, NOW(), NOW()),
(4, 4, 4, 4, 'Project D', 1.80, 13000, 18000, 230000, 15000, 20000, 250000, 1, 1, 1, NOW(), NOW()),
(5, 5, 5, 5, 'Project E', 1.90, 14000, 19000, 240000, 16000, 21000, 260000, 1, 1, 1, NOW(), NOW());

INSERT INTO mine_client_transport (
    mine_client_transport_id, mine_client_project_id,
    truck_id, driver_id,
    volume_to_weight_convertion,
    buy_price_per_cubic, buy_price_per_ton, buy_price_per_transport,
    sell_price_per_cubic, sell_price_per_ton, sell_price_per_transport,
    created_by, updated_by, created_at, updated_at
) VALUES
(1, 1, 1, 1, 1.50, 10000, 15000, 200000, 12000, 17000, 220000, 1, 1, NOW(), NOW()),
(2, 2, 2, 2, 1.60, 11000, 16000, 210000, 13000, 18000, 230000, 1, 1, NOW(), NOW()),
(3, 3, 3, 3, 1.70, 12000, 17000, 220000, 14000, 19000, 240000, 1, 1, NOW(), NOW()),
(4, 4, 4, 4, 1.80, 13000, 18000, 230000, 15000, 20000, 250000, 1, 1, NOW(), NOW()),
(5, 5, 5, 5, 1.90, 14000, 19000, 240000, 16000, 21000, 260000, 1, 1, NOW(), NOW());


INSERT INTO mine_client_transport_status (
    mine_client_transport_status_id,
    mine_client_transport_id,
    mine_client_transport_status_type_id,
    status_time,
    cargo_box_length, cargo_box_width, cargo_box_height,
    cargo_weight,
    mine_client_transport_status_note,
    is_fraud, is_active,
    created_by, updated_by,
    created_at, updated_at
) VALUES
(1, 1, 1, NOW(), 2.5, 2.0, 1.5, 10.0, 'Loaded', 0, 1, 1, 1, NOW(), NOW()),
(2, 2, 2, NOW(), 2.6, 2.1, 1.6, 11.0, 'On Delivery', 0, 1, 1, 1, NOW(), NOW()),
(3, 3, 3, NOW(), 2.7, 2.2, 1.7, 12.0, 'Arrived', 0, 1, 1, 1, NOW(), NOW()),
(4, 4, 1, NOW(), 2.8, 2.3, 1.8, 13.0, 'Loaded Again', 0, 1, 1, 1, NOW(), NOW()),
(5, 5, 2, NOW(), 2.9, 2.4, 1.9, 14.0, 'In Transit', 0, 1, 1, 1, NOW(), NOW());


INSERT INTO mine_client_transport_photo (
    mine_client_transport_photo_id,
    mine_client_transport_status_id,
    photo_url,
    photo_type_id,
    photo_description,
    created_by, updated_by,
    created_at, updated_at
) VALUES
(1, 1, 'https://example.com/photo1.jpg', 1, 'Loading Photo', 1, 1, NOW(), NOW()),
(2, 2, 'https://example.com/photo2.jpg', 1, 'Delivery Photo', 1, 1, NOW(), NOW()),
(3, 3, 'https://example.com/photo3.jpg', 2, 'Arrival Photo', 1, 1, NOW(), NOW()),
(4, 4, 'https://example.com/photo4.jpg', 1, 'Reload Photo', 1, 1, NOW(), NOW()),
(5, 5, 'https://example.com/photo5.jpg', 2, 'Transit Photo', 1, 1, NOW(), NOW());



DO $$
DECLARE
    r RECORD;
BEGIN
    FOR r IN
        SELECT
            c.relname AS sequence_name,
            t.relname AS table_name,
            a.attname AS column_name
        FROM pg_class c
        JOIN pg_depend d ON d.objid = c.oid
        JOIN pg_class t ON d.refobjid = t.oid
        JOIN pg_attribute a ON a.attrelid = t.oid AND a.attnum = d.refobjsubid
        WHERE c.relkind = 'S'
    LOOP
        EXECUTE format(
            'SELECT setval(''%I'', COALESCE((SELECT MAX(%I) FROM %I), 1));',
            r.sequence_name,
            r.column_name,
            r.table_name
        );
    END LOOP;
END $$;
app_user

province

city

bank_merk

vendor_type

vendor

truck_type

truck

client

client_destination

cargo_type

stockpile

stockpile_cargo

client_project

stockpile_project

stockpile_transport

stockpile_transport_status

stockpile_transport_photo

client_transport

client_transport_status

client_transport_photo

stockpile_adjustment

stockpile_ledger


ALTER TABLE app_user
ADD CONSTRAINT fk_app_user_city
FOREIGN KEY (city_id) REFERENCES city(city_id);

ALTER TABLE app_user
ADD CONSTRAINT fk_app_user_client
FOREIGN KEY (client_id) REFERENCES client(client_id);

DROP SCHEMA public CASCADE;
CREATE SCHEMA public;

SHOW TIMEZONE;
SET TIMEZONE = 'Asia/Jakarta';

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
    mine_id INT NOT NULL REFERENCES mine(mine_id),
    stockpile_cargo_id INT NOT NULL REFERENCES stockpile_cargo(stockpile_cargo_id),
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
    stockpile_cargo_id INT NOT NULL REFERENCES stockpile_cargo(stockpile_cargo_id),
    client_destination_id INT NOT NULL REFERENCES client_destination(client_destination_id),
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
    stockpile_project_id INT NOT NULL REFERENCES stockpile_project(stockpile_project_id),
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

CREATE TABLE mine_client_project (
    mine_client_project_id SERIAL PRIMARY KEY,
    mine_id INT NOT NULL REFERENCES mine(mine_id),
    cargo_type_id INT NOT NULL REFERENCES cargo_type(cargo_type_id),
    client_destination_id INT NOT NULL REFERENCES client_destination(client_destination_id),
    mine_client_project_name VARCHAR(100) NOT NULL,
    volume_to_weight_convertion NUMERIC(10,2),
    buy_price_per_cubic NUMERIC(15,2),
    buy_price_per_ton NUMERIC(15,2),
    buy_price_per_transport NUMERIC(15,2),
    sell_price_per_cubic NUMERIC(15,2),
    sell_price_per_ton NUMERIC(15,2),
    sell_price_per_transport NUMERIC(15,2),
    is_active SMALLINT NOT NULL DEFAULT 1,
    created_by INT NOT NULL REFERENCES app_user(app_user_id),
    updated_by INT NOT NULL REFERENCES app_user(app_user_id),
    deleted_by INT REFERENCES app_user(app_user_id),
    created_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMPTZ
);


CREATE TABLE mine_client_transport (
    mine_client_transport_id SERIAL PRIMARY KEY,
    mine_client_project_id INT NOT NULL REFERENCES mine_client_project(mine_client_project_id),
    truck_id INT REFERENCES truck(truck_id),
    driver_id INT REFERENCES app_user(app_user_id),
    volume_to_weight_convertion NUMERIC(10,2),
    buy_price_per_cubic NUMERIC(15,2),
    buy_price_per_ton NUMERIC(15,2),
    buy_price_per_transport NUMERIC(15,2),
    sell_price_per_cubic NUMERIC(15,2),
    sell_price_per_ton NUMERIC(15,2),
    sell_price_per_transport NUMERIC(15,2),
    created_by INT NOT NULL REFERENCES app_user(app_user_id),
    updated_by INT NOT NULL REFERENCES app_user(app_user_id),
    deleted_by INT REFERENCES app_user(app_user_id),
    created_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMPTZ
);

CREATE TABLE mine_client_transport_status (
    mine_client_transport_status_id SERIAL PRIMARY KEY,
    mine_client_transport_id INT REFERENCES mine_client_transport(mine_client_transport_id),
    mine_client_transport_status_type_id SMALLINT NOT NULL,
    status_time TIMESTAMPTZ NOT NULL,
    cargo_box_length NUMERIC(10,2),
    cargo_box_width NUMERIC(10,2),
    cargo_box_height NUMERIC(10,2),
    cargo_weight NUMERIC(10,2),
    mine_client_transport_status_note TEXT,
    is_fraud SMALLINT NOT NULL DEFAULT 0,
    is_active SMALLINT NOT NULL DEFAULT 1,
    created_by INT NOT NULL REFERENCES app_user(app_user_id),
    updated_by INT NOT NULL REFERENCES app_user(app_user_id),
    deleted_by INT REFERENCES app_user(app_user_id),
    created_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMPTZ
);

CREATE TABLE mine_client_transport_photo (
    mine_client_transport_photo_id SERIAL PRIMARY KEY,
    mine_client_transport_status_id INT REFERENCES mine_client_transport_status(mine_client_transport_status_id),
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


photo_type_id
1 delivery note
2 cargo box

stockpile_transport_status_type_id
1 arrived at mine
2 loading at mine
3 loaded at mine
4 going to stockpile
5 arrived at stockpile
6 unloading at stockpile
7 unloaded at stockpile
8 completed

client_transport_status_type_id
1 arrived at stockpile
2 loading at stockpile
3 loaded at stockpile
4 going to client
5 arrived at client
6 unloading at client
7 unloaded at client
8 completed

mine_client_transport_status_type_id
1 arrived at mine
2 loading at mine
3 loaded at mine
4 going to client
5 arrived at client
6 unloading at client
7 unloaded at client
8 completed


reference_type_id
1 stockpile_transport
2 client_transport
3 stockpile_adjustment
4 manual
5 correction
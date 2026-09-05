ALTER TABLE app_user
ADD CONSTRAINT fk_app_user_city
FOREIGN KEY (city_id) REFERENCES city(city_id);

ALTER TABLE app_user
ADD CONSTRAINT fk_app_user_district
FOREIGN KEY (district_id) REFERENCES district(district_id);

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
    district_id INT NULL,
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
    created_by INT NOT NULL REFERENCES app_user(app_user_id),
    updated_by INT NOT NULL REFERENCES app_user(app_user_id),
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

CREATE TABLE district (
    district_id SERIAL PRIMARY KEY,
    city_id INT REFERENCES city(city_id),
    district_name VARCHAR(100) NOT NULL,
    is_active SMALLINT NOT NULL DEFAULT 1,
    created_by INT NOT NULL REFERENCES app_user(app_user_id),
    updated_by INT NOT NULL REFERENCES app_user(app_user_id),
    deleted_by INT REFERENCES app_user(app_user_id),
    created_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMPTZ
);



CREATE TABLE fleet_type (
    fleet_type_id SERIAL PRIMARY KEY,
    fleet_type_name VARCHAR(100) NOT NULL,
    cargo_box_length NUMERIC(10,2) NOT NULL,
    cargo_box_width NUMERIC(10,2) NOT NULL,
    cargo_box_height NUMERIC(10,2) NOT NULL,
    cargo_capacity NUMERIC(10,2) NOT NULL,
    is_active SMALLINT NOT NULL DEFAULT 1,
    created_by INT NOT NULL REFERENCES app_user(app_user_id),
    updated_by INT NOT NULL REFERENCES app_user(app_user_id),
    deleted_by INT REFERENCES app_user(app_user_id),
    created_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMPTZ
);

CREATE TABLE fleet_merk (
    fleet_merk_id SERIAL PRIMARY KEY,
    fleet_merk_name VARCHAR(100) NOT NULL,
    is_active SMALLINT NOT NULL DEFAULT 1,
    created_by INT NOT NULL REFERENCES app_user(app_user_id),
    updated_by INT NOT NULL REFERENCES app_user(app_user_id),
    deleted_by INT REFERENCES app_user(app_user_id),
    created_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMPTZ
);

CREATE TABLE fleet (
    fleet_id SERIAL PRIMARY KEY,
    fleet_type_id INT REFERENCES fleet_type(fleet_type_id),
    fleet_merk_id INT REFERENCES fleet_merk(fleet_merk_id),
    driver_id INT REFERENCES app_user(app_user_id),
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

CREATE TABLE fleet_document (
    fleet_document_id SERIAL PRIMARY KEY,
    fleet_id INT NOT NULL REFERENCES fleet(fleet_id),
    fleet_document_expiration_date DATE,
    fleet_document_photo_url TEXT,
    fleet_document_type_id SMALLINT NOT NULL DEFAULT 1,
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
    district_id INT NULL REFERENCES district(district_id),
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


CREATE TABLE client_product (
    client_product_id SERIAL PRIMARY KEY,
    client_id INT REFERENCES client(client_id),
    client_product_name VARCHAR(100) NOT NULL,
    unit_of_measure_id INT NOT NULL,
    is_active SMALLINT NOT NULL DEFAULT 1,
    created_by INT NOT NULL REFERENCES app_user(app_user_id),
    updated_by INT NOT NULL REFERENCES app_user(app_user_id),
    deleted_by INT REFERENCES app_user(app_user_id),
    created_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMPTZ
);

CREATE TABLE destination (
    destination_id SERIAL PRIMARY KEY,
    client_id INT REFERENCES client(client_id),
    destination_name VARCHAR(100) NOT NULL,
    city_id INT NULL REFERENCES city(city_id),
    district_id INT NULL REFERENCES district(district_id),
    destination_address TEXT,
    destination_latitude NUMERIC(9,6),
    destination_longitude NUMERIC(9,6),
    destination_map_url TEXT,
    operating_hours JSONB,
    is_active SMALLINT NOT NULL DEFAULT 1,
    created_by INT NOT NULL REFERENCES app_user(app_user_id),
    updated_by INT NOT NULL REFERENCES app_user(app_user_id),
    deleted_by INT REFERENCES app_user(app_user_id),
    created_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMPTZ
);

-- ALTER TABLE destination
-- ADD COLUMN operating_hours JSONB;

-- SELECT
--   operating_hours -> 'monday' ->> 'open'  AS open_time,
--   operating_hours -> 'monday' ->> 'close' AS close_time
-- FROM destination
-- WHERE destination_id = 1;

-- SELECT
--   operating_hours -> 'sunday' ->> 'closed' = 'true' AS is_closed
-- FROM destination
-- WHERE destination_id = 1;

-- {
--   "monday": {
--     "open": "08:00",
--     "close": "17:00"
--   },
--   "tuesday": {
--     "open": "08:00",
--     "close": "17:00"
--   },
--   "wednesday": {
--     "open": "08:00",
--     "close": "17:00"
--   },
--   "thursday": {
--     "open": "08:00",
--     "close": "17:00"
--   },
--   "friday": {
--     "open": "08:00",
--     "close": "16:00"
--   },
--   "saturday": {
--     "open": "09:00",
--     "close": "14:00"
--   },
--   "sunday": {
--     "closed": true
--   }
-- }



CREATE TABLE project (
    project_id SERIAL PRIMARY KEY,
    client_id INT NOT NULL REFERENCES client(client_id),
    project_name VARCHAR(200) NOT NULL,
    per_overnight_rate NUMERIC(15,2) DEFAULT 0,
    per_multy_rate NUMERIC(15,2) DEFAULT 0,
    per_multy_city_rate NUMERIC(15,2) DEFAULT 0,
    waresix BOOLEAN DEFAULT FALSE,
    is_active SMALLINT NOT NULL DEFAULT 1,
    created_by INT NOT NULL REFERENCES app_user(app_user_id),
    updated_by INT NOT NULL REFERENCES app_user(app_user_id),
    deleted_by INT REFERENCES app_user(app_user_id),
    created_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMPTZ
);

CREATE TABLE project_detail (
    project_detail_id SERIAL PRIMARY KEY,
    project_id INT NOT NULL REFERENCES project(project_id),
    basic_rate NUMERIC(15,2) NOT NULL DEFAULT 0,
    basic_rate_type_id SMALLINT NOT NULL DEFAULT 1,
    driver_fee NUMERIC(15,2) NOT NULL DEFAULT 0,
    project_detail_description TEXT,
    is_active SMALLINT NOT NULL DEFAULT 1,
    created_by INT NOT NULL REFERENCES app_user(app_user_id),
    updated_by INT NOT NULL REFERENCES app_user(app_user_id),
    deleted_by INT REFERENCES app_user(app_user_id),
    created_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMPTZ
);

CREATE TABLE travel_allowance (
    travel_allowance_id SERIAL PRIMARY KEY,
    project_detail_id INT NOT NULL REFERENCES project_detail(project_detail_id),
    fleet_type_id INT NOT NULL REFERENCES fleet_type(fleet_type_id),
    travel_allowance_amount_1 NUMERIC(15,2) NOT NULL,
    travel_allowance_amount_2 NUMERIC(15,2) NOT NULL,
    created_by INT NOT NULL REFERENCES app_user(app_user_id),
    updated_by INT NOT NULL REFERENCES app_user(app_user_id),
    deleted_by INT REFERENCES app_user(app_user_id),
    created_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMPTZ
);

CREATE TABLE project_detail_destination (
    project_detail_destination_id SERIAL PRIMARY KEY,
    project_detail_id INT NOT NULL REFERENCES project_detail(project_detail_id),
    project_detail_destination_type_id SMALLINT NOT NULL,
    project_detail_destination_sequence SMALLINT NOT NULL,
    destination_id INT NOT NULL REFERENCES destination(destination_id),
    to_next_duration_minutes INT NOT NULL DEFAULT 60,
    is_active SMALLINT NOT NULL DEFAULT 1,
    created_by INT NOT NULL REFERENCES app_user(app_user_id),
    updated_by INT NOT NULL REFERENCES app_user(app_user_id),
    deleted_by INT REFERENCES app_user(app_user_id),
    created_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMPTZ
);

CREATE TABLE delivery_order (
    delivery_order_id SERIAL PRIMARY KEY,
    client_id INT REFERENCES client(client_id),
    project_detail_id INT REFERENCES project_detail(project_detail_id),
    fleet_type_id INT REFERENCES fleet_type(fleet_type_id),
    fleet_id INT REFERENCES fleet(fleet_id),
    driver_id INT REFERENCES app_user(app_user_id),
    delivery_order_number VARCHAR(50) UNIQUE NOT NULL,
    invoice_number VARCHAR(50) UNIQUE,
    basic_rate NUMERIC(15,2) DEFAULT 0,
    per_overnight_rate NUMERIC(15,2),
    per_multy_rate NUMERIC(15,2),
    per_multy_city_rate NUMERIC(15,2),
    additional_rate NUMERIC(15,2) DEFAULT 0,
    discount_rate NUMERIC(15,2) DEFAULT 0,
    completed_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
    travel_allowance_amount_1 NUMERIC(15,2) DEFAULT 0,
    travel_allowance_amount_2 NUMERIC(15,2) DEFAULT 0,
    driver_fee NUMERIC(15,2) DEFAULT 0,
    created_by INT NOT NULL REFERENCES app_user(app_user_id),
    updated_by INT NOT NULL REFERENCES app_user(app_user_id),
    deleted_by INT REFERENCES app_user(app_user_id),
    created_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMPTZ
);

CREATE TABLE delivery_order_destination (
    delivery_order_destination_id SERIAL PRIMARY KEY,
    delivery_order_id INT REFERENCES delivery_order(delivery_order_id),
    destination_id INT REFERENCES destination(destination_id),
    delivery_order_destination_type_id SMALLINT NOT NULL, 
    status_time TIMESTAMPTZ,
    delivery_order_destination_sequence SMALLINT NOT NULL,
    created_by INT NOT NULL REFERENCES app_user(app_user_id),
    updated_by INT NOT NULL REFERENCES app_user(app_user_id),
    deleted_by INT REFERENCES app_user(app_user_id),
    created_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMPTZ
);

CREATE TABLE delivery_order_status (
    delivery_order_status_id SERIAL PRIMARY KEY,
    delivery_order_destination_id INT REFERENCES delivery_order_destination(delivery_order_destination_id),
    city_id INT REFERENCES city(city_id),
    district_id INT REFERENCES district(district_id),
    delivery_order_status_type_id INT NOT NULL,
    status_time TIMESTAMPTZ NOT NULL,
    delivery_order_status_latitude NUMERIC(9,6),
    delivery_order_status_longitude NUMERIC(9,6),
    created_by INT NOT NULL REFERENCES app_user(app_user_id),
    updated_by INT NOT NULL REFERENCES app_user(app_user_id),
    deleted_by INT REFERENCES app_user(app_user_id),
    created_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMPTZ
);

CREATE TABLE delivery_order_product (
    delivery_order_product_id SERIAL PRIMARY KEY,
    delivery_order_status_id INT REFERENCES delivery_order_status(delivery_order_status_id),
    client_product_id INT REFERENCES client_product(client_product_id),
    client_product_quantity NUMERIC(15,2) DEFAULT 0,
    created_by INT NOT NULL REFERENCES app_user(app_user_id),
    updated_by INT NOT NULL REFERENCES app_user(app_user_id),
    deleted_by INT REFERENCES app_user(app_user_id),
    created_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMPTZ
);

CREATE TABLE delivery_order_note_photo (
    delivery_order_note_photo_id SERIAL PRIMARY KEY,
    delivery_order_status_id INT REFERENCES delivery_order_status(delivery_order_status_id),
    delivery_order_note_photo_url TEXT,
    created_by INT NOT NULL REFERENCES app_user(app_user_id),
    updated_by INT NOT NULL REFERENCES app_user(app_user_id),
    deleted_by INT REFERENCES app_user(app_user_id),
    created_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMPTZ
);

CREATE TABLE delivery_order_note_photo_approval (
    delivery_order_note_photo_approval_id SERIAL PRIMARY KEY,
    delivery_order_note_photo_id INT NOT NULL 
        REFERENCES delivery_order_note_photo(delivery_order_note_photo_id)
        ON DELETE CASCADE,
    approval_step_id SMALLINT NOT NULL DEFAULT 1,   -- level approval
    approval_status_id SMALLINT NOT NULL,           -- 0=pending, 1=approved, 2=revised, 3=rejected
    approval_notes TEXT,                         -- catatan tiap step
    created_by INT NOT NULL REFERENCES app_user(app_user_id),
    updated_by INT NOT NULL REFERENCES app_user(app_user_id),
    deleted_by INT REFERENCES app_user(app_user_id),
    created_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMPTZ
);

CREATE TABLE delivery_order_cargo_photo (
    delivery_order_cargo_photo_id SERIAL PRIMARY KEY,
    delivery_order_status_id INT REFERENCES delivery_order_status(delivery_order_status_id),
    delivery_order_cargo_photo_url TEXT,
    created_by INT NOT NULL REFERENCES app_user(app_user_id),
    updated_by INT NOT NULL REFERENCES app_user(app_user_id),
    deleted_by INT REFERENCES app_user(app_user_id),
    created_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMPTZ
);

CREATE TABLE delivery_order_cargo_photo_approval (
    delivery_order_cargo_photo_approval_id SERIAL PRIMARY KEY,
    delivery_order_cargo_photo_id INT NOT NULL 
        REFERENCES delivery_order_cargo_photo(delivery_order_cargo_photo_id)
        ON DELETE CASCADE,
    approval_step_id SMALLINT NOT NULL DEFAULT 1,   -- level approval
    approval_status_id SMALLINT NOT NULL,           -- 0=pending, 1=approved, 2=revised, 3=rejected
    approval_notes TEXT,                         -- catatan tiap step
    created_by INT NOT NULL REFERENCES app_user(app_user_id),
    updated_by INT NOT NULL REFERENCES app_user(app_user_id),
    deleted_by INT REFERENCES app_user(app_user_id),
    created_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMPTZ
);

CREATE TABLE delivery_order_travel_allowance (
    delivery_order_travel_allowance_id SERIAL PRIMARY KEY,
    delivery_order_id INT NOT NULL REFERENCES delivery_order(delivery_order_id)
        ON DELETE CASCADE,
    driver_id INT NOT NULL REFERENCES app_user(app_user_id),
    bank_merk_id INT REFERENCES bank_merk(bank_merk_id),
    bank_account_number VARCHAR(100),
    bank_account_name VARCHAR(100),
    delivery_sequence SMALLINT NOT NULL,
    expected_amount NUMERIC(15,2) NOT NULL,
    delivered_amount NUMERIC(15,2) DEFAULT 0,
    created_by INT NOT NULL REFERENCES app_user(app_user_id),
    updated_by INT NOT NULL REFERENCES app_user(app_user_id),
    deleted_by INT REFERENCES app_user(app_user_id),
    created_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMPTZ
);

CREATE TABLE delivery_order_travel_allowance_approval (
    delivery_order_travel_allowance_approval_id SERIAL PRIMARY KEY,
    delivery_order_travel_allowance_id INT NOT NULL 
        REFERENCES delivery_order_travel_allowance(delivery_order_travel_allowance_id)
        ON DELETE CASCADE,
    approval_step_id SMALLINT NOT NULL DEFAULT 1,   -- level approval
    approval_status_id SMALLINT NOT NULL,           -- 0=pending, 1=approved, 2=revised, 3=rejected
    approval_notes TEXT,                         -- catatan tiap step
    created_by INT NOT NULL REFERENCES app_user(app_user_id),
    updated_by INT NOT NULL REFERENCES app_user(app_user_id),
    deleted_by INT REFERENCES app_user(app_user_id),
    created_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMPTZ
);

CREATE TABLE delivery_order_incident (
    delivery_order_incident_id SERIAL PRIMARY KEY,
    delivery_order_destination_id INT NOT NULL REFERENCES delivery_order_destination(delivery_order_destination_id),
    reported_by INT NOT NULL REFERENCES app_user(app_user_id),
    client_product_id INT REFERENCES client_product(client_product_id),
    client_product_quantity NUMERIC(15,2) DEFAULT 0,
    incident_type_id INT NOT NULL,
    incident_description TEXT, 
    compensation_amount NUMERIC(15,2) DEFAULT 0,
    created_by INT NOT NULL REFERENCES app_user(app_user_id),
    updated_by INT NOT NULL REFERENCES app_user(app_user_id),
    deleted_by INT REFERENCES app_user(app_user_id),
    created_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMPTZ
);

CREATE TABLE delivery_order_incident_photo (
    delivery_order_incident_photo_id SERIAL PRIMARY KEY,
    delivery_order_incident_id INT REFERENCES delivery_order_incident(delivery_order_incident_id),
    delivery_order_incident_photo_url TEXT,
    created_by INT NOT NULL REFERENCES app_user(app_user_id),
    updated_by INT NOT NULL REFERENCES app_user(app_user_id),
    deleted_by INT REFERENCES app_user(app_user_id),
    created_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMPTZ
);

CREATE TABLE delivery_order_incident_approval (
    delivery_order_incident_approval_id SERIAL PRIMARY KEY,
    delivery_order_incident_id INT NOT NULL 
        REFERENCES delivery_order_incident(delivery_order_incident_id)
        ON DELETE CASCADE,
    approval_step_id SMALLINT NOT NULL DEFAULT 1,   -- level approval
    approval_status_id SMALLINT NOT NULL,           -- 0=pending, 1=approved, 2=revised, 3=rejected
    approval_notes TEXT,                         -- catatan tiap step
    created_by INT NOT NULL REFERENCES app_user(app_user_id),
    updated_by INT NOT NULL REFERENCES app_user(app_user_id),
    deleted_by INT REFERENCES app_user(app_user_id),
    created_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMPTZ
);


CREATE TABLE workshop (
    workshop_id SERIAL PRIMARY KEY,
    workshop_name VARCHAR(100) NOT NULL,
    workshop_phone VARCHAR(20) NOT NULL,
    city_id INT NULL,
    district_id INT NULL,
    workshop_photo_url TEXT,
    workshop_address TEXT,
    workshop_latitude NUMERIC(9,6),
    workshop_longitude NUMERIC(9,6),
    workshop_map_url TEXT,
    is_active SMALLINT NOT NULL DEFAULT 1,
    created_by INT NOT NULL REFERENCES app_user(app_user_id),
    updated_by INT NOT NULL REFERENCES app_user(app_user_id),
    deleted_by INT REFERENCES app_user(app_user_id),
    created_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMPTZ
);

CREATE TABLE irregular_cost_type (
    irregular_cost_type_id SERIAL PRIMARY KEY,
    irregular_cost_type_name VARCHAR(100) NOT NULL,
    irregular_cost_type_description TEXT,
    is_active SMALLINT NOT NULL DEFAULT 1,
    created_by INT NOT NULL REFERENCES app_user(app_user_id),
    updated_by INT NOT NULL REFERENCES app_user(app_user_id),
    deleted_by INT REFERENCES app_user(app_user_id),
    created_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMPTZ
);

CREATE TABLE irregular_cost (
    irregular_cost_id SERIAL PRIMARY KEY,
    fleet_id INT REFERENCES fleet(fleet_id),
    reported_by INT REFERENCES app_user(app_user_id),
    workshop_id INT REFERENCES workshop(workshop_id),
    delivery_order_id INT REFERENCES delivery_order(delivery_order_id),
    irregular_cost_type_id INT REFERENCES irregular_cost_type(irregular_cost_type_id),
    irregular_cost_description TEXT,
    irregular_cost_amount NUMERIC(15,2) DEFAULT 0,
    irregular_cost_latitude NUMERIC(9,6),
    irregular_cost_longitude NUMERIC(9,6),
    created_by INT NOT NULL REFERENCES app_user(app_user_id),
    updated_by INT NOT NULL REFERENCES app_user(app_user_id),
    deleted_by INT REFERENCES app_user(app_user_id),
    created_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMPTZ
);

CREATE TABLE irregular_cost_photo (
    irregular_cost_photo_id SERIAL PRIMARY KEY,
    irregular_cost_id INT REFERENCES irregular_cost(irregular_cost_id),
    irregular_cost_photo_url TEXT,
    created_by INT NOT NULL REFERENCES app_user(app_user_id),
    updated_by INT NOT NULL REFERENCES app_user(app_user_id),
    deleted_by INT REFERENCES app_user(app_user_id),
    created_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMPTZ
);

CREATE TABLE irregular_cost_approval (
    irregular_cost_approval_id SERIAL PRIMARY KEY,
    irregular_cost_id INT NOT NULL 
        REFERENCES irregular_cost(irregular_cost_id)
        ON DELETE CASCADE,
    approval_step_id SMALLINT NOT NULL DEFAULT 1,   -- level approval
    approval_status_id SMALLINT NOT NULL,           -- 0=pending, 1=approved, 2=revised, 3=rejected
    approval_notes TEXT,                         -- catatan tiap step
    created_by INT NOT NULL REFERENCES app_user(app_user_id),
    updated_by INT NOT NULL REFERENCES app_user(app_user_id),
    deleted_by INT REFERENCES app_user(app_user_id),
    created_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMPTZ
);


CREATE TABLE leave_request_type (
    leave_request_type_id SERIAL PRIMARY KEY,
    leave_request_type_name VARCHAR(100) NOT NULL,
    leave_request_type_description TEXT,
    is_active SMALLINT NOT NULL DEFAULT 1,
    created_by INT NOT NULL REFERENCES app_user(app_user_id),
    updated_by INT NOT NULL REFERENCES app_user(app_user_id),
    deleted_by INT REFERENCES app_user(app_user_id),
    created_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMPTZ
);

CREATE TABLE leave_request (
    leave_request_id SERIAL PRIMARY KEY,
    requested_by INT NOT NULL REFERENCES app_user(app_user_id),
    leave_request_type_id INT NOT NULL REFERENCES leave_request_type(leave_request_type_id),
    leave_request_reason TEXT,
    leave_request_start_at TIMESTAMPTZ NOT NULL,
    leave_request_end_at TIMESTAMPTZ,
    -- CHECK (leave_request_end_at >= leave_request_start_at + INTERVAL '12 hours'),
    created_by INT NOT NULL REFERENCES app_user(app_user_id),
    updated_by INT NOT NULL REFERENCES app_user(app_user_id),
    deleted_by INT REFERENCES app_user(app_user_id),
    created_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMPTZ
);

CREATE TABLE leave_request_photo (
    leave_request_photo_id SERIAL PRIMARY KEY,
    leave_request_id INT NOT NULL REFERENCES leave_request(leave_request_id),
    leave_request_photo_url TEXT,
    created_by INT NOT NULL REFERENCES app_user(app_user_id),
    updated_by INT NOT NULL REFERENCES app_user(app_user_id),
    deleted_by INT REFERENCES app_user(app_user_id),
    created_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMPTZ
);

CREATE TABLE leave_request_approval (
    leave_request_approval_id SERIAL PRIMARY KEY,
    leave_request_id INT NOT NULL 
        REFERENCES leave_request(leave_request_id)
        ON DELETE CASCADE,
    approval_step_id SMALLINT NOT NULL DEFAULT 1,   -- level approval
    approval_status_id SMALLINT NOT NULL,           -- 0=pending, 1=approved, 2=revised, 3=rejected
    approval_notes TEXT,                         -- catatan tiap step
    created_by INT NOT NULL REFERENCES app_user(app_user_id),
    updated_by INT NOT NULL REFERENCES app_user(app_user_id),
    deleted_by INT REFERENCES app_user(app_user_id),
    created_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMPTZ
);


CREATE TABLE app_user_loan_type (
    app_user_loan_type_id SERIAL PRIMARY KEY,
    app_user_loan_type_name VARCHAR(100) NOT NULL,
    app_user_loan_type_description TEXT,
    is_active SMALLINT NOT NULL DEFAULT 1,
    created_by INT NOT NULL REFERENCES app_user(app_user_id),
    updated_by INT NOT NULL REFERENCES app_user(app_user_id),
    deleted_by INT REFERENCES app_user(app_user_id),
    created_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMPTZ
);

CREATE TABLE app_user_loan (
    app_user_loan_id SERIAL PRIMARY KEY,
    requested_by INT NOT NULL REFERENCES app_user(app_user_id),
    app_user_loan_type_id INT NOT NULL REFERENCES app_user_loan_type(app_user_loan_type_id),
    app_user_loan_reason TEXT,
    app_user_loan_amount NUMERIC(15,2) DEFAULT 0,
    loan_duration_months INT NOT NULL DEFAULT 1,
    first_payment_date TIMESTAMPTZ,
    deduction_method_id SMALLINT NOT NULL DEFAULT 1,
    created_by INT NOT NULL REFERENCES app_user(app_user_id),
    updated_by INT NOT NULL REFERENCES app_user(app_user_id),
    deleted_by INT REFERENCES app_user(app_user_id),
    created_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMPTZ
);

CREATE TABLE app_user_loan_approval (
    app_user_loan_approval_id SERIAL PRIMARY KEY,
    app_user_loan_id INT NOT NULL 
        REFERENCES app_user_loan(app_user_loan_id)
        ON DELETE CASCADE,
    approval_step_id SMALLINT NOT NULL DEFAULT 1,   -- level approval
    approval_status_id SMALLINT NOT NULL,           -- 0=pending, 1=approved, 2=revised, 3=rejected
    approval_notes TEXT,                         -- catatan tiap step
    created_by INT NOT NULL REFERENCES app_user(app_user_id),
    updated_by INT NOT NULL REFERENCES app_user(app_user_id),
    deleted_by INT REFERENCES app_user(app_user_id),
    created_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMPTZ
);

CREATE INDEX idx_app_user_loan_approval_loan_id 
ON app_user_loan_approval(app_user_loan_id);

CREATE INDEX idx_app_user_loan_approval_status_id 
ON app_user_loan_approval(approval_status_id);

CREATE INDEX idx_app_user_loan_approval_step_id 
ON app_user_loan_approval(approval_step_id);

CREATE TABLE app_user_loan_photo (
    app_user_loan_photo_id SERIAL PRIMARY KEY,
    app_user_loan_id INT NOT NULL REFERENCES app_user_loan(app_user_loan_id),
    app_user_loan_photo_url TEXT,
    created_by INT NOT NULL REFERENCES app_user(app_user_id),
    updated_by INT NOT NULL REFERENCES app_user(app_user_id),
    deleted_by INT REFERENCES app_user(app_user_id),
    created_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMPTZ
);


CREATE TABLE app_user_loan_installment (
    loan_installment_id SERIAL PRIMARY KEY,

    app_user_loan_id INT NOT NULL
        REFERENCES app_user_loan(app_user_loan_id)
        ON DELETE CASCADE,
    installment_number INT NOT NULL,        -- bulan ke berapa (1,2,3,...)
    due_date DATE NOT NULL,                 -- jatuh tempo bulan itu
    amount NUMERIC(15,2) NOT NULL,          -- total kewajiban cicilan bulan itu
    amount_paid NUMERIC(15,2) DEFAULT 0,    -- total yang sudah dibayar

    is_paid BOOLEAN NOT NULL DEFAULT FALSE, -- TRUE jika amount_paid >= amount

    created_by INT NOT NULL REFERENCES app_user(app_user_id),
    updated_by INT NOT NULL REFERENCES app_user(app_user_id),
    deleted_by INT REFERENCES app_user(app_user_id),

    created_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMPTZ
);

CREATE TABLE push_subscriptions (
    push_subscription_id SERIAL PRIMARY KEY,

    app_user_id INT NOT NULL REFERENCES app_user(app_user_id),
    app_role_id INT NOT NULL REFERENCES app_role(app_role_id),

    topic VARCHAR(100) NOT NULL,
    -- contoh:
    -- driver:delivery_order
    -- admin:order
    -- client:invoice

    endpoint TEXT NOT NULL,
    p256dh TEXT NOT NULL,
    auth TEXT NOT NULL,

    user_agent TEXT,
    ip_address INET,
    is_active SMALLINT NOT NULL DEFAULT 1,

    created_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMPTZ
);

-- CREATE UNIQUE INDEX uniq_push_subscription_active
-- ON push_subscriptions (endpoint, app_user_id)
-- WHERE deleted_at IS NULL;

ALTER TABLE push_subscriptions
ADD CONSTRAINT uniq_push_subscription_active
UNIQUE (endpoint, app_user_id, is_active);



CREATE INDEX idx_push_user_topic
ON push_subscriptions (app_user_id, topic);

CREATE INDEX idx_push_role_topic
ON push_subscriptions (app_role_id, topic);


-- CREATE TABLE city_route (
--     city_route_id SERIAL PRIMARY KEY,
--     city_a_id INT REFERENCES city(city_id),
--     city_b_id INT REFERENCES city(city_id),
--     route_distance NUMERIC(10,2) DEFAULT 0,
--     created_by INT NOT NULL REFERENCES app_user(app_user_id),
--     updated_by INT NOT NULL REFERENCES app_user(app_user_id),
--     deleted_by INT REFERENCES app_user(app_user_id),
--     created_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,
--     updated_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,
--     deleted_at TIMESTAMPTZ,
--     CONSTRAINT city_route_unique_pair UNIQUE (city_a_id, city_b_id)
-- );

-- CREATE TABLE city_route_rate (
--     city_route_rate_id SERIAL PRIMARY KEY,
--     city_route_id INT REFERENCES city_route(city_route_id),
--     fleet_type_id INT REFERENCES fleet_type(fleet_type_id),
--     route_rate NUMERIC(15,2) DEFAULT 0,
--     created_by INT NOT NULL REFERENCES app_user(app_user_id),
--     updated_by INT NOT NULL REFERENCES app_user(app_user_id),
--     deleted_by INT REFERENCES app_user(app_user_id),
--     created_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,
--     updated_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,
--     deleted_at TIMESTAMPTZ
-- );




-- CREATE TABLE delivery_order_destination_type (
--     delivery_order_destination_type_id SERIAL PRIMARY KEY,
--     delivery_order_destination_type_name VARCHAR(50) NOT NULL,
--     created_by INT NOT NULL REFERENCES app_user(app_user_id),
--     updated_by INT NOT NULL REFERENCES app_user(app_user_id),
--     deleted_by INT REFERENCES app_user(app_user_id),
--     created_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,
--     updated_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,
--     deleted_at TIMESTAMPTZ
-- );

-- CREATE TABLE delivery_order_status_type (
--     delivery_order_status_type_id SERIAL PRIMARY KEY,
--     delivery_order_status_name VARCHAR(50) NOT NULL,
--     is_active SMALLINT NOT NULL DEFAULT 1,
--     created_by INT NOT NULL REFERENCES app_user(app_user_id),
--     updated_by INT NOT NULL REFERENCES app_user(app_user_id),
--     deleted_by INT REFERENCES app_user(app_user_id),
--     created_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,
--     updated_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,
--     deleted_at TIMESTAMPTZ
-- );



-- CREATE TABLE client_pic (
--     client_pic_id SERIAL PRIMARY KEY,
--     client_id INT REFERENCES client(client_id),
--     pic_name VARCHAR(100) NOT NULL,
--     pic_phone VARCHAR(20),
--     city_id INT NULL REFERENCES city(city_id),
--     district_id INT NULL REFERENCES district(district_id),
--     pic_address TEXT,
--     is_active SMALLINT NOT NULL DEFAULT 1,
--     created_by INT NOT NULL REFERENCES app_user(app_user_id),
--     updated_by INT NOT NULL REFERENCES app_user(app_user_id),
--     deleted_by INT REFERENCES app_user(app_user_id),
--     created_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,
--     updated_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,
--     deleted_at TIMESTAMPTZ
-- );


-- CREATE TABLE client_pic (
--     client_pic_id SERIAL PRIMARY KEY,
--     client_id INT REFERENCES client(client_id),
--     pic_name VARCHAR(100) NOT NULL,
--     pic_phone VARCHAR(20),
--     city_id INT NULL REFERENCES city(city_id),
--     district_id INT NULL REFERENCES district(district_id),
--     pic_address TEXT,
--     is_active SMALLINT NOT NULL DEFAULT 1,
--     created_by INT NOT NULL REFERENCES app_user(app_user_id),
--     updated_by INT NOT NULL REFERENCES app_user(app_user_id),
--     deleted_by INT REFERENCES app_user(app_user_id),
--     created_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,
--     updated_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,
--     deleted_at TIMESTAMPTZ
-- );


CONST

app_role_type_id
1 driver
2 client_pic
3 admin

app_user_status_id
0 pending
1 aktif
2 non aktif

is_active default 1
0 tidak aktif
1 aktif

approval_status_id default 0
0 pending
1 approved
2 rejected
3 canceled


basic_rate_type_id
1 trip
2 kg
3 box
4 m3

unit_of_measure_id 
1 pcs
2 kg
3 box
4 m3




delivery_order_destination_type_id

1 pickup
2 drop


ownership_status_id
1 pribadi
2 sewa

fleet_document_type_id
1 annual_stnk
2 5year_stnk
3 kir1
4 kir2
5 bpkb

delivery_order_status_type_id
0 pending
1 assigned on delivery (driver sudah menerima DO, tapi masih mengerjakan delivery lain)
2 assigned ready (driver siap berangkat, tidak ada delivery lain)
2 going to pick up
3 arrived at pick up (peringatan untuk foto)
4 loading (foto bukti awal muat)
5 loaded (foto bukti muat dan surat jalan) (pilihan drop atau pickup, hanya jika tidak ada pick selanjutnya)
6 going to delivery
7 arrived at delivery (peringatan untuk foto)
8 unloading (foto buktu awal dibuka)
9 unloaded  (foto bukti selesai bongkar dan surat jalan) (pilihan drop atau selesai, hanya jika tidak ada drop selanjutnya)
10 completed
11 canceled



next development 

?expand untuk menambahkan join table
bulk delete, bulk activate, bulk deactivate
filter data yang dihapus dan endpoint restore dan bulk restore data


driver izin
pencatatan driver menolak delivery_order, dengan cara merubah ke izin sakit
rating for driver

delivery_order_status kurang approval_status_id dan approval_notes

destination kurang 
    destination_latitude NUMERIC(9,6),
    destination_longitude NUMERIC(9,6),
    google map

workshop kurang 
    workshop_latitude NUMERIC(9,6),
    workshop_longitude NUMERIC(9,6),
    google map


client kurang npwp dan tanggal jatuh tempo

client dan client pic ketika create dijadikan 1 saja 


app user atau pegawai

kurang nama panggilan, gaji(sopir %), foto ktp, nomor ktp, foto kk, nomor kk, foto sim, 
nomor sim, masa berlaku sim B, foto npwp, nomor npwp, foto bpjs, bpjs, tipe bank, nomor rekening

ada pengaturan feature admin tapi hardcode saja


driver bisa mendapatkan 2 order tapi statusnya pending

truck kurang tentang foto stnk, nomor stnk, masa berlaku stnk, foto bpkb, status kepemilikan, merk, (buat table snediri aja), 
data kir (dibuatkan tabel tersendiri),
foto kir1, foto kir2,

project 
project adalah penghubung antara client dan order
jadi ketika membuat order harus ada projectnya
project adalah bisa dijadikan satu berdasarkan kota muat
project detail adalah rute
ada 2 tabel 

project nama project, client id, harga per over night, harga per multy drop, harga per multy city, waresix (boolean)

detail project 
kota muat, alamat muat, kota bongkat, alamat bongkar, 
satuan berat, harga muat, harga sewa, keterangan,
gaji supir, sangu supir (berdasarkan tipe truck), 


project 
client id, 
harga per over night, 
harga per multy drop,
harga per multy city, 
waresix (boolean)



project detail 
project detail id, project id,
harga muat, harga sewa, keterangan,
gaji supir,

travel_allowance
travel_allowance_id
project detail id
fleet_id
travel_allowance_amount

project detail destination
id_project_detail_destination
project_detail_id
project_detail_type_id -- 1 muat/ 2 bongkar
destination_id


vendor turck (next development) karena tidak dapat ssan projectnya



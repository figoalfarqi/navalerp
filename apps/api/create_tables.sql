-- Initial database schema.
-- Project, route, transport, and finance tables are intentionally unified so
-- every operational and financial record can be reported by one project_id.

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
    username VARCHAR(200) UNIQUE NOT NULL,
    password TEXT NOT NULL,
    auth_version INT NOT NULL DEFAULT 0,
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
    deleted_at TIMESTAMPTZ,
    CONSTRAINT chk_app_user_status CHECK (
        app_user_status_id IN (0, 1, 2)
    ),
    CONSTRAINT chk_app_user_salary_percentage CHECK (
        salary_percentage IS NULL
        OR salary_percentage BETWEEN 0 AND 100
    )
);

-- Application-wide settings are deliberately stored as keyed JSON values
-- rather than mixed into project data.  Only authorised administration roles
-- may update these records through the dedicated settings CRUD.
CREATE TABLE app_setting (
    app_setting_id SERIAL PRIMARY KEY,
    app_setting_key VARCHAR(100) UNIQUE NOT NULL,
    app_setting_value JSONB NOT NULL DEFAULT '{}'::jsonb,
    app_setting_description TEXT,
    is_active SMALLINT NOT NULL DEFAULT 1,
    created_by INT NOT NULL REFERENCES app_user(app_user_id),
    updated_by INT NOT NULL REFERENCES app_user(app_user_id),
    deleted_by INT REFERENCES app_user(app_user_id),
    created_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMPTZ,
    CONSTRAINT chk_app_setting_is_active CHECK (is_active IN (0, 1)),
    CONSTRAINT chk_app_setting_key CHECK (
        app_setting_key ~ '^[a-z][a-z0-9_.-]{1,99}$'
    )
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

CREATE TABLE port (
    port_id SERIAL PRIMARY KEY,
    port_name VARCHAR(200) NOT NULL,
    city_id INT REFERENCES city(city_id),
    port_address TEXT,
    port_latitude NUMERIC(9,6),
    port_longitude NUMERIC(9,6),
    port_map_url TEXT,
    is_active SMALLINT NOT NULL DEFAULT 1,
    created_by INT NOT NULL REFERENCES app_user(app_user_id),
    updated_by INT NOT NULL REFERENCES app_user(app_user_id),
    deleted_by INT REFERENCES app_user(app_user_id),
    created_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMPTZ
);

CREATE TABLE vessel (
    vessel_id SERIAL PRIMARY KEY,
    vendor_id INT REFERENCES vendor(vendor_id),
    vessel_name VARCHAR(200) NOT NULL,
    imo_number VARCHAR(20),
    registration_number VARCHAR(100),
    is_active SMALLINT NOT NULL DEFAULT 1,
    created_by INT NOT NULL REFERENCES app_user(app_user_id),
    updated_by INT NOT NULL REFERENCES app_user(app_user_id),
    deleted_by INT REFERENCES app_user(app_user_id),
    created_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMPTZ,
    CONSTRAINT uq_vessel_imo_number UNIQUE (imo_number)
);

-- One record represents one cargo/voyage unloaded from a vessel at a port.
-- This is the source reference for VESSEL_* projects.
CREATE TABLE vessel_cargo (
    vessel_cargo_id SERIAL PRIMARY KEY,
    vessel_id INT NOT NULL REFERENCES vessel(vessel_id),
    port_id INT NOT NULL REFERENCES port(port_id),
    cargo_type_id INT NOT NULL REFERENCES cargo_type(cargo_type_id),
    voyage_number VARCHAR(100),
    bill_of_lading_number VARCHAR(100),
    arrival_at TIMESTAMPTZ,
    unloading_started_at TIMESTAMPTZ,
    unloading_completed_at TIMESTAMPTZ,
    manifest_volume_cubic NUMERIC(15,3),
    manifest_weight_ton NUMERIC(15,3),
    is_active SMALLINT NOT NULL DEFAULT 1,
    created_by INT NOT NULL REFERENCES app_user(app_user_id),
    updated_by INT NOT NULL REFERENCES app_user(app_user_id),
    deleted_by INT REFERENCES app_user(app_user_id),
    created_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMPTZ,
    CONSTRAINT uq_vessel_cargo_id_cargo_type UNIQUE (
        vessel_cargo_id,
        cargo_type_id
    ),
    CONSTRAINT chk_vessel_cargo_quantity_nonnegative CHECK (
        COALESCE(manifest_volume_cubic, 0) >= 0
        AND COALESCE(manifest_weight_ton, 0) >= 0
    )
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
    stockpile_id INT NOT NULL REFERENCES stockpile(stockpile_id),
    cargo_type_id INT NOT NULL REFERENCES cargo_type(cargo_type_id),
    capacity_volume_cubic NUMERIC(15,3),
    capacity_weight_ton NUMERIC(15,3),
    current_volume_cubic NUMERIC(15,3) NOT NULL DEFAULT 0,
    current_weight_ton NUMERIC(15,3) NOT NULL DEFAULT 0,
    is_active SMALLINT NOT NULL DEFAULT 1,
    created_by INT NOT NULL REFERENCES app_user(app_user_id),
    updated_by INT NOT NULL REFERENCES app_user(app_user_id),
    deleted_by INT REFERENCES app_user(app_user_id),
    created_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMPTZ,
    CONSTRAINT uq_stockpile_cargo UNIQUE (stockpile_id, cargo_type_id),
    CONSTRAINT uq_stockpile_cargo_id_cargo_type UNIQUE (
        stockpile_cargo_id,
        cargo_type_id
    ),
    CONSTRAINT chk_stockpile_cargo_quantity_nonnegative CHECK (
        COALESCE(capacity_volume_cubic, 0) >= 0
        AND COALESCE(capacity_weight_ton, 0) >= 0
        AND current_volume_cubic >= 0
        AND current_weight_ton >= 0
    )
);


-- A project is one commercial contract. route_type determines the source and
-- whether the cargo transits through a stockpile:
-- MINE_CLIENT, MINE_STOCKPILE_CLIENT, VESSEL_CLIENT,
-- VESSEL_STOCKPILE_CLIENT.
CREATE TABLE project (
    project_id SERIAL PRIMARY KEY,
    project_code VARCHAR(50) UNIQUE NOT NULL,
    project_name VARCHAR(200) NOT NULL,
    route_type VARCHAR(30) NOT NULL,
    mine_id INT REFERENCES mine(mine_id),
    vessel_cargo_id INT REFERENCES vessel_cargo(vessel_cargo_id),
    stockpile_cargo_id INT REFERENCES stockpile_cargo(stockpile_cargo_id),
    client_destination_id INT NOT NULL REFERENCES client_destination(client_destination_id),
    cargo_type_id INT NOT NULL REFERENCES cargo_type(cargo_type_id),
    project_status VARCHAR(20) NOT NULL DEFAULT 'DRAFT',
    start_date DATE,
    end_date DATE,
    planned_volume_cubic NUMERIC(15,3),
    planned_weight_ton NUMERIC(15,3),
    volume_to_weight_conversion NUMERIC(15,6),

    -- Material purchase is an expense. Only the price selected by
    -- material_purchase_unit is used in financial calculations.
    material_purchase_unit VARCHAR(10) NOT NULL DEFAULT 'NONE',
    material_buy_price_per_cubic NUMERIC(18,2),
    material_buy_price_per_ton NUMERIC(18,2),

    -- Material sale is income. Only the price selected by
    -- material_sale_unit is used in financial calculations.
    material_sale_unit VARCHAR(10) NOT NULL DEFAULT 'NONE',
    material_sell_price_per_cubic NUMERIC(18,2),
    material_sell_price_per_ton NUMERIC(18,2),

    -- Project-level amounts that are not calculated per transport.
    fixed_other_income NUMERIC(18,2) NOT NULL DEFAULT 0,
    fixed_other_expense NUMERIC(18,2) NOT NULL DEFAULT 0,
    project_note TEXT,
    is_active SMALLINT NOT NULL DEFAULT 1,
    created_by INT NOT NULL REFERENCES app_user(app_user_id),
    updated_by INT NOT NULL REFERENCES app_user(app_user_id),
    deleted_by INT REFERENCES app_user(app_user_id),
    created_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMPTZ,
    CONSTRAINT uq_project_id_route_type UNIQUE (project_id, route_type),
    CONSTRAINT fk_project_vessel_cargo_type FOREIGN KEY (
        vessel_cargo_id,
        cargo_type_id
    ) REFERENCES vessel_cargo(vessel_cargo_id, cargo_type_id),
    CONSTRAINT fk_project_stockpile_cargo_type FOREIGN KEY (
        stockpile_cargo_id,
        cargo_type_id
    ) REFERENCES stockpile_cargo(stockpile_cargo_id, cargo_type_id),
    CONSTRAINT chk_project_route_type CHECK (
        (route_type = 'MINE_CLIENT'
            AND mine_id IS NOT NULL
            AND vessel_cargo_id IS NULL
            AND stockpile_cargo_id IS NULL)
        OR
        (route_type = 'MINE_STOCKPILE_CLIENT'
            AND mine_id IS NOT NULL
            AND vessel_cargo_id IS NULL
            AND stockpile_cargo_id IS NOT NULL)
        OR
        (route_type = 'VESSEL_CLIENT'
            AND mine_id IS NULL
            AND vessel_cargo_id IS NOT NULL
            AND stockpile_cargo_id IS NULL)
        OR
        (route_type = 'VESSEL_STOCKPILE_CLIENT'
            AND mine_id IS NULL
            AND vessel_cargo_id IS NOT NULL
            AND stockpile_cargo_id IS NOT NULL)
    ),
    CONSTRAINT chk_project_status CHECK (
        project_status IN ('DRAFT', 'ACTIVE', 'COMPLETED', 'CANCELLED')
    ),
    CONSTRAINT chk_project_dates CHECK (
        end_date IS NULL OR start_date IS NULL OR end_date >= start_date
    ),
    CONSTRAINT chk_project_quantity_nonnegative CHECK (
        COALESCE(planned_volume_cubic, 0) >= 0
        AND COALESCE(planned_weight_ton, 0) >= 0
        AND COALESCE(volume_to_weight_conversion, 0) >= 0
    ),
    CONSTRAINT chk_project_material_purchase_price CHECK (
        (material_purchase_unit = 'NONE'
            AND material_buy_price_per_cubic IS NULL
            AND material_buy_price_per_ton IS NULL)
        OR
        (material_purchase_unit = 'M3'
            AND material_buy_price_per_cubic IS NOT NULL
            AND material_buy_price_per_cubic >= 0
            AND material_buy_price_per_ton IS NULL)
        OR
        (material_purchase_unit = 'TON'
            AND material_buy_price_per_ton IS NOT NULL
            AND material_buy_price_per_ton >= 0
            AND material_buy_price_per_cubic IS NULL)
    ),
    CONSTRAINT chk_project_material_sale_price CHECK (
        (material_sale_unit = 'NONE'
            AND material_sell_price_per_cubic IS NULL
            AND material_sell_price_per_ton IS NULL)
        OR
        (material_sale_unit = 'M3'
            AND material_sell_price_per_cubic IS NOT NULL
            AND material_sell_price_per_cubic >= 0
            AND material_sell_price_per_ton IS NULL)
        OR
        (material_sale_unit = 'TON'
            AND material_sell_price_per_ton IS NOT NULL
            AND material_sell_price_per_ton >= 0
            AND material_sell_price_per_cubic IS NULL)
    ),
    CONSTRAINT chk_project_fixed_amount_nonnegative CHECK (
        fixed_other_income >= 0 AND fixed_other_expense >= 0
    )
);

-- Controls which projects can be accessed by each checker.
-- The application must only return projects joined through an active,
-- non-deleted assignment for the authenticated checker.
CREATE TABLE project_checker_assignment (
    project_checker_assignment_id SERIAL PRIMARY KEY,
    project_id INT NOT NULL REFERENCES project(project_id),
    checker_id INT NOT NULL REFERENCES app_user(app_user_id),
    access_started_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
    access_ended_at TIMESTAMPTZ,
    assignment_note TEXT,
    is_active SMALLINT NOT NULL DEFAULT 1,
    -- The checker application preselects this project after login.  A
    -- checker can have many assignments but only one active default.
    is_default SMALLINT NOT NULL DEFAULT 0,
    created_by INT NOT NULL REFERENCES app_user(app_user_id),
    updated_by INT NOT NULL REFERENCES app_user(app_user_id),
    deleted_by INT REFERENCES app_user(app_user_id),
    created_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMPTZ,
    CONSTRAINT chk_project_checker_access_period CHECK (
        access_ended_at IS NULL OR access_ended_at >= access_started_at
    ),
    CONSTRAINT chk_project_checker_is_active CHECK (
        is_active IN (0, 1)
    ),
    CONSTRAINT chk_project_checker_is_default CHECK (
        is_default IN (0, 1)
        AND (is_default = 0 OR is_active = 1)
    )
);

-- Assigns a truck to a project. A truck may have assignments to more than
-- one project, while the application selects only active assignments for new
-- transports. Historical transports retain their assignment reference.
CREATE TABLE project_truck_assignment (
    project_truck_assignment_id SERIAL PRIMARY KEY,
    project_id INT NOT NULL REFERENCES project(project_id),
    truck_id INT NOT NULL REFERENCES truck(truck_id),
    assignment_started_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
    assignment_ended_at TIMESTAMPTZ,
    assignment_note TEXT,
    is_active SMALLINT NOT NULL DEFAULT 1,
    created_by INT NOT NULL REFERENCES app_user(app_user_id),
    updated_by INT NOT NULL REFERENCES app_user(app_user_id),
    deleted_by INT REFERENCES app_user(app_user_id),
    created_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMPTZ,
    CONSTRAINT uq_project_truck_assignment_id_project_truck UNIQUE (
        project_truck_assignment_id,
        project_id,
        truck_id
    ),
    CONSTRAINT chk_project_truck_assignment_period CHECK (
        assignment_ended_at IS NULL
        OR assignment_ended_at >= assignment_started_at
    ),
    CONSTRAINT chk_project_truck_assignment_is_active CHECK (
        is_active IN (0, 1)
    )
);

-- One project has one direct route or two routes when using a stockpile.
-- Pricing here is the default for one route/leg and can be snapshotted or
-- overridden in project_transport without changing historical transports.
CREATE TABLE project_route (
    project_route_id SERIAL PRIMARY KEY,
    project_id INT NOT NULL,
    project_pattern VARCHAR(30) NOT NULL,
    route_sequence SMALLINT NOT NULL,
    route_type VARCHAR(30) NOT NULL,
    route_name VARCHAR(200),
    distance_km NUMERIC(12,3),

    -- Income from transport service billed to the client.
    transport_service_unit VARCHAR(10) NOT NULL DEFAULT 'NONE',
    transport_service_price_per_cubic NUMERIC(18,2),
    transport_service_price_per_ton NUMERIC(18,2),
    transport_service_price_per_transport NUMERIC(18,2),

    -- Expense paid to an external/internal transport provider.
    transport_cost_unit VARCHAR(10) NOT NULL DEFAULT 'NONE',
    transport_cost_per_cubic NUMERIC(18,2),
    transport_cost_per_ton NUMERIC(18,2),
    transport_cost_per_transport NUMERIC(18,2),

    -- Default operational expenses for each transport on this route.
    road_money_per_transport NUMERIC(18,2) NOT NULL DEFAULT 0,
    loading_cost_per_transport NUMERIC(18,2) NOT NULL DEFAULT 0,
    unloading_cost_per_transport NUMERIC(18,2) NOT NULL DEFAULT 0,
    fuel_cost_per_transport NUMERIC(18,2) NOT NULL DEFAULT 0,
    toll_cost_per_transport NUMERIC(18,2) NOT NULL DEFAULT 0,
    other_income_per_transport NUMERIC(18,2) NOT NULL DEFAULT 0,
    other_expense_per_transport NUMERIC(18,2) NOT NULL DEFAULT 0,
    route_note TEXT,
    is_active SMALLINT NOT NULL DEFAULT 1,
    created_by INT NOT NULL REFERENCES app_user(app_user_id),
    updated_by INT NOT NULL REFERENCES app_user(app_user_id),
    deleted_by INT REFERENCES app_user(app_user_id),
    created_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMPTZ,
    CONSTRAINT fk_project_route_project_pattern FOREIGN KEY (
        project_id,
        project_pattern
    ) REFERENCES project(project_id, route_type),
    CONSTRAINT uq_project_route_id_project UNIQUE (
        project_route_id,
        project_id
    ),
    CONSTRAINT uq_project_route_sequence UNIQUE (project_id, route_sequence),
    CONSTRAINT uq_project_route_type UNIQUE (project_id, route_type),
    CONSTRAINT chk_project_route_pattern CHECK (
        (
            project_pattern IN ('MINE_CLIENT', 'VESSEL_CLIENT')
            AND route_sequence = 1
            AND route_type = 'SOURCE_TO_CLIENT'
        )
        OR
        (
            project_pattern IN (
                'MINE_STOCKPILE_CLIENT',
                'VESSEL_STOCKPILE_CLIENT'
            )
            AND (
                (
                    route_sequence = 1
                    AND route_type = 'SOURCE_TO_STOCKPILE'
                )
                OR
                (
                    route_sequence = 2
                    AND route_type = 'STOCKPILE_TO_CLIENT'
                )
            )
        )
    ),
    CONSTRAINT chk_project_route_distance CHECK (
        COALESCE(distance_km, 0) >= 0
    ),
    CONSTRAINT chk_project_route_service_price CHECK (
        (transport_service_unit = 'NONE'
            AND transport_service_price_per_cubic IS NULL
            AND transport_service_price_per_ton IS NULL
            AND transport_service_price_per_transport IS NULL)
        OR
        (transport_service_unit = 'M3'
            AND transport_service_price_per_cubic IS NOT NULL
            AND transport_service_price_per_cubic >= 0
            AND transport_service_price_per_ton IS NULL
            AND transport_service_price_per_transport IS NULL)
        OR
        (transport_service_unit = 'TON'
            AND transport_service_price_per_ton IS NOT NULL
            AND transport_service_price_per_ton >= 0
            AND transport_service_price_per_cubic IS NULL
            AND transport_service_price_per_transport IS NULL)
        OR
        (transport_service_unit = 'TRANSPORT'
            AND transport_service_price_per_transport IS NOT NULL
            AND transport_service_price_per_transport >= 0
            AND transport_service_price_per_cubic IS NULL
            AND transport_service_price_per_ton IS NULL)
    ),
    CONSTRAINT chk_project_route_transport_cost CHECK (
        (transport_cost_unit = 'NONE'
            AND transport_cost_per_cubic IS NULL
            AND transport_cost_per_ton IS NULL
            AND transport_cost_per_transport IS NULL)
        OR
        (transport_cost_unit = 'M3'
            AND transport_cost_per_cubic IS NOT NULL
            AND transport_cost_per_cubic >= 0
            AND transport_cost_per_ton IS NULL
            AND transport_cost_per_transport IS NULL)
        OR
        (transport_cost_unit = 'TON'
            AND transport_cost_per_ton IS NOT NULL
            AND transport_cost_per_ton >= 0
            AND transport_cost_per_cubic IS NULL
            AND transport_cost_per_transport IS NULL)
        OR
        (transport_cost_unit = 'TRANSPORT'
            AND transport_cost_per_transport IS NOT NULL
            AND transport_cost_per_transport >= 0
            AND transport_cost_per_cubic IS NULL
            AND transport_cost_per_ton IS NULL)
    ),
    CONSTRAINT chk_project_route_default_amount_nonnegative CHECK (
        road_money_per_transport >= 0
        AND loading_cost_per_transport >= 0
        AND unloading_cost_per_transport >= 0
        AND fuel_cost_per_transport >= 0
        AND toll_cost_per_transport >= 0
        AND other_income_per_transport >= 0
        AND other_expense_per_transport >= 0
    )
);

-- One row is one actual truck transport on one project route. All price fields
-- are snapshots so later project/route price changes do not rewrite history.
-- Purchase quantity is filled only at the material purchase point; sale
-- quantity is filled only at the client delivery point to prevent double count
-- on projects that have two routes.
CREATE TABLE project_transport (
    project_transport_id SERIAL PRIMARY KEY,
    project_id INT NOT NULL REFERENCES project(project_id),
    project_route_id INT NOT NULL,
    transport_number VARCHAR(100) NOT NULL,
    delivery_note_number VARCHAR(100),
    transported_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
    truck_id INT REFERENCES truck(truck_id),
    project_truck_assignment_id INT,
    driver_id INT REFERENCES app_user(app_user_id),
    transport_vendor_id INT REFERENCES vendor(vendor_id),

    loaded_volume_cubic NUMERIC(15,3),
    loaded_weight_ton NUMERIC(15,3),
    delivered_volume_cubic NUMERIC(15,3),
    delivered_weight_ton NUMERIC(15,3),
    purchase_volume_cubic NUMERIC(15,3),
    purchase_weight_ton NUMERIC(15,3),
    sale_volume_cubic NUMERIC(15,3),
    sale_weight_ton NUMERIC(15,3),
    transport_service_volume_cubic NUMERIC(15,3),
    transport_service_weight_ton NUMERIC(15,3),
    transport_cost_volume_cubic NUMERIC(15,3),
    transport_cost_weight_ton NUMERIC(15,3),
    volume_to_weight_conversion NUMERIC(15,6),

    material_purchase_unit VARCHAR(10) NOT NULL DEFAULT 'NONE',
    material_buy_price_per_cubic NUMERIC(18,2),
    material_buy_price_per_ton NUMERIC(18,2),
    material_purchase_amount NUMERIC(18,2) GENERATED ALWAYS AS (
        CASE material_purchase_unit
            WHEN 'M3' THEN COALESCE(purchase_volume_cubic, 0)
                * COALESCE(material_buy_price_per_cubic, 0)
            WHEN 'TON' THEN COALESCE(purchase_weight_ton, 0)
                * COALESCE(material_buy_price_per_ton, 0)
            ELSE 0
        END
    ) STORED,

    material_sale_unit VARCHAR(10) NOT NULL DEFAULT 'NONE',
    material_sell_price_per_cubic NUMERIC(18,2),
    material_sell_price_per_ton NUMERIC(18,2),
    material_sale_amount NUMERIC(18,2) GENERATED ALWAYS AS (
        CASE material_sale_unit
            WHEN 'M3' THEN COALESCE(sale_volume_cubic, 0)
                * COALESCE(material_sell_price_per_cubic, 0)
            WHEN 'TON' THEN COALESCE(sale_weight_ton, 0)
                * COALESCE(material_sell_price_per_ton, 0)
            ELSE 0
        END
    ) STORED,

    transport_service_unit VARCHAR(10) NOT NULL DEFAULT 'NONE',
    transport_service_price_per_cubic NUMERIC(18,2),
    transport_service_price_per_ton NUMERIC(18,2),
    transport_service_price_per_transport NUMERIC(18,2),
    transport_service_income_amount NUMERIC(18,2) GENERATED ALWAYS AS (
        CASE transport_service_unit
            WHEN 'M3' THEN COALESCE(transport_service_volume_cubic, 0)
                * COALESCE(transport_service_price_per_cubic, 0)
            WHEN 'TON' THEN COALESCE(transport_service_weight_ton, 0)
                * COALESCE(transport_service_price_per_ton, 0)
            WHEN 'TRANSPORT' THEN COALESCE(transport_service_price_per_transport, 0)
            ELSE 0
        END
    ) STORED,

    transport_cost_unit VARCHAR(10) NOT NULL DEFAULT 'NONE',
    transport_cost_per_cubic NUMERIC(18,2),
    transport_cost_per_ton NUMERIC(18,2),
    transport_cost_per_transport NUMERIC(18,2),
    transport_expense_amount NUMERIC(18,2) GENERATED ALWAYS AS (
        CASE transport_cost_unit
            WHEN 'M3' THEN COALESCE(transport_cost_volume_cubic, 0)
                * COALESCE(transport_cost_per_cubic, 0)
            WHEN 'TON' THEN COALESCE(transport_cost_weight_ton, 0)
                * COALESCE(transport_cost_per_ton, 0)
            WHEN 'TRANSPORT' THEN COALESCE(transport_cost_per_transport, 0)
            ELSE 0
        END
    ) STORED,

    -- Actual per-transport operational amounts.
    road_money_amount NUMERIC(18,2) NOT NULL DEFAULT 0,
    loading_cost_amount NUMERIC(18,2) NOT NULL DEFAULT 0,
    unloading_cost_amount NUMERIC(18,2) NOT NULL DEFAULT 0,
    fuel_cost_amount NUMERIC(18,2) NOT NULL DEFAULT 0,
    toll_cost_amount NUMERIC(18,2) NOT NULL DEFAULT 0,
    other_income_amount NUMERIC(18,2) NOT NULL DEFAULT 0,
    other_expense_amount NUMERIC(18,2) NOT NULL DEFAULT 0,
    is_completed SMALLINT NOT NULL DEFAULT 0,
    transport_note TEXT,
    created_by INT NOT NULL REFERENCES app_user(app_user_id),
    updated_by INT NOT NULL REFERENCES app_user(app_user_id),
    deleted_by INT REFERENCES app_user(app_user_id),
    created_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMPTZ,
    CONSTRAINT fk_project_transport_route_project FOREIGN KEY (
        project_route_id,
        project_id
    ) REFERENCES project_route(project_route_id, project_id),
    CONSTRAINT fk_project_transport_truck_assignment FOREIGN KEY (
        project_truck_assignment_id,
        project_id,
        truck_id
    ) REFERENCES project_truck_assignment(
        project_truck_assignment_id,
        project_id,
        truck_id
    ),
    CONSTRAINT chk_project_transport_truck_assignment_presence CHECK (
        (truck_id IS NULL AND project_truck_assignment_id IS NULL)
        OR
        (truck_id IS NOT NULL AND project_truck_assignment_id IS NOT NULL)
    ),
    CONSTRAINT uq_project_transport_id_project UNIQUE (
        project_transport_id,
        project_id
    ),
    CONSTRAINT uq_project_transport_number UNIQUE (
        project_id,
        transport_number
    ),
    CONSTRAINT chk_project_transport_quantity_nonnegative CHECK (
        COALESCE(loaded_volume_cubic, 0) >= 0
        AND COALESCE(loaded_weight_ton, 0) >= 0
        AND COALESCE(delivered_volume_cubic, 0) >= 0
        AND COALESCE(delivered_weight_ton, 0) >= 0
        AND COALESCE(purchase_volume_cubic, 0) >= 0
        AND COALESCE(purchase_weight_ton, 0) >= 0
        AND COALESCE(sale_volume_cubic, 0) >= 0
        AND COALESCE(sale_weight_ton, 0) >= 0
        AND COALESCE(transport_service_volume_cubic, 0) >= 0
        AND COALESCE(transport_service_weight_ton, 0) >= 0
        AND COALESCE(transport_cost_volume_cubic, 0) >= 0
        AND COALESCE(transport_cost_weight_ton, 0) >= 0
        AND COALESCE(volume_to_weight_conversion, 0) >= 0
    ),
    CONSTRAINT chk_project_transport_material_purchase_price CHECK (
        (material_purchase_unit = 'NONE'
            AND material_buy_price_per_cubic IS NULL
            AND material_buy_price_per_ton IS NULL)
        OR
        (material_purchase_unit = 'M3'
            AND material_buy_price_per_cubic IS NOT NULL
            AND material_buy_price_per_cubic >= 0
            AND material_buy_price_per_ton IS NULL)
        OR
        (material_purchase_unit = 'TON'
            AND material_buy_price_per_ton IS NOT NULL
            AND material_buy_price_per_ton >= 0
            AND material_buy_price_per_cubic IS NULL)
    ),
    CONSTRAINT chk_project_transport_material_sale_price CHECK (
        (material_sale_unit = 'NONE'
            AND material_sell_price_per_cubic IS NULL
            AND material_sell_price_per_ton IS NULL)
        OR
        (material_sale_unit = 'M3'
            AND material_sell_price_per_cubic IS NOT NULL
            AND material_sell_price_per_cubic >= 0
            AND material_sell_price_per_ton IS NULL)
        OR
        (material_sale_unit = 'TON'
            AND material_sell_price_per_ton IS NOT NULL
            AND material_sell_price_per_ton >= 0
            AND material_sell_price_per_cubic IS NULL)
    ),
    CONSTRAINT chk_project_transport_service_price CHECK (
        (transport_service_unit = 'NONE'
            AND transport_service_price_per_cubic IS NULL
            AND transport_service_price_per_ton IS NULL
            AND transport_service_price_per_transport IS NULL)
        OR
        (transport_service_unit = 'M3'
            AND transport_service_price_per_cubic IS NOT NULL
            AND transport_service_price_per_cubic >= 0
            AND transport_service_price_per_ton IS NULL
            AND transport_service_price_per_transport IS NULL)
        OR
        (transport_service_unit = 'TON'
            AND transport_service_price_per_ton IS NOT NULL
            AND transport_service_price_per_ton >= 0
            AND transport_service_price_per_cubic IS NULL
            AND transport_service_price_per_transport IS NULL)
        OR
        (transport_service_unit = 'TRANSPORT'
            AND transport_service_price_per_transport IS NOT NULL
            AND transport_service_price_per_transport >= 0
            AND transport_service_price_per_cubic IS NULL
            AND transport_service_price_per_ton IS NULL)
    ),
    CONSTRAINT chk_project_transport_cost_price CHECK (
        (transport_cost_unit = 'NONE'
            AND transport_cost_per_cubic IS NULL
            AND transport_cost_per_ton IS NULL
            AND transport_cost_per_transport IS NULL)
        OR
        (transport_cost_unit = 'M3'
            AND transport_cost_per_cubic IS NOT NULL
            AND transport_cost_per_cubic >= 0
            AND transport_cost_per_ton IS NULL
            AND transport_cost_per_transport IS NULL)
        OR
        (transport_cost_unit = 'TON'
            AND transport_cost_per_ton IS NOT NULL
            AND transport_cost_per_ton >= 0
            AND transport_cost_per_cubic IS NULL
            AND transport_cost_per_transport IS NULL)
        OR
        (transport_cost_unit = 'TRANSPORT'
            AND transport_cost_per_transport IS NOT NULL
            AND transport_cost_per_transport >= 0
            AND transport_cost_per_cubic IS NULL
            AND transport_cost_per_ton IS NULL)
    ),
    CONSTRAINT chk_project_transport_amount_nonnegative CHECK (
        road_money_amount >= 0
        AND loading_cost_amount >= 0
        AND unloading_cost_amount >= 0
        AND fuel_cost_amount >= 0
        AND toll_cost_amount >= 0
        AND other_income_amount >= 0
        AND other_expense_amount >= 0
    ),
    CONSTRAINT chk_project_transport_is_completed CHECK (
        is_completed IN (0, 1)
    )
);

CREATE TABLE project_transport_status (
    project_transport_status_id SERIAL PRIMARY KEY,
    project_transport_id INT NOT NULL REFERENCES project_transport(project_transport_id),
    project_transport_status_type_id SMALLINT NOT NULL,
    status_time TIMESTAMPTZ NOT NULL,
    cargo_box_length NUMERIC(10,2),
    cargo_box_width NUMERIC(10,2),
    cargo_box_height NUMERIC(10,2),
    cargo_volume_cubic NUMERIC(15,3),
    cargo_weight_ton NUMERIC(15,3),
    project_transport_status_note TEXT,
    is_fraud SMALLINT NOT NULL DEFAULT 0,
    is_active SMALLINT NOT NULL DEFAULT 1,
    created_by INT NOT NULL REFERENCES app_user(app_user_id),
    updated_by INT NOT NULL REFERENCES app_user(app_user_id),
    deleted_by INT REFERENCES app_user(app_user_id),
    created_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMPTZ,
    CONSTRAINT chk_project_transport_status_type CHECK (
        project_transport_status_type_id BETWEEN 1 AND 8
    ),
    CONSTRAINT chk_project_transport_status_fraud CHECK (
        is_fraud IN (0, 1)
        AND (
            is_fraud = 0
            OR BTRIM(COALESCE(project_transport_status_note, '')) <> ''
        )
    ),
    CONSTRAINT chk_project_transport_status_measurements CHECK (
        COALESCE(cargo_box_length, 0) >= 0
        AND COALESCE(cargo_box_width, 0) >= 0
        AND COALESCE(cargo_box_height, 0) >= 0
        AND COALESCE(cargo_volume_cubic, 0) >= 0
        AND COALESCE(cargo_weight_ton, 0) >= 0
    )
);

CREATE TABLE project_transport_photo (
    project_transport_photo_id SERIAL PRIMARY KEY,
    project_transport_status_id INT NOT NULL
        REFERENCES project_transport_status(project_transport_status_id),
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

-- Financial documents and cash movements. Categories use explicit columns
-- (not EAV) so a project report can SUM each income/expense component.
-- Operational/accrual values remain available from project_transport; this
-- table records receivables, receipts, payables, payments, and adjustments.
CREATE TABLE project_financial_transaction (
    project_financial_transaction_id SERIAL PRIMARY KEY,
    project_id INT NOT NULL REFERENCES project(project_id),
    project_transport_id INT,
    transaction_number VARCHAR(100) UNIQUE NOT NULL,
    transaction_kind VARCHAR(20) NOT NULL,
    transaction_date DATE NOT NULL,
    due_date DATE,
    paid_at TIMESTAMPTZ,
    client_id INT REFERENCES client(client_id),
    vendor_id INT REFERENCES vendor(vendor_id),
    reference_number VARCHAR(100),

    material_sale_income NUMERIC(18,2) NOT NULL DEFAULT 0,
    transport_service_income NUMERIC(18,2) NOT NULL DEFAULT 0,
    other_income NUMERIC(18,2) NOT NULL DEFAULT 0,
    material_purchase_expense NUMERIC(18,2) NOT NULL DEFAULT 0,
    transport_expense NUMERIC(18,2) NOT NULL DEFAULT 0,
    road_money_expense NUMERIC(18,2) NOT NULL DEFAULT 0,
    loading_expense NUMERIC(18,2) NOT NULL DEFAULT 0,
    unloading_expense NUMERIC(18,2) NOT NULL DEFAULT 0,
    fuel_expense NUMERIC(18,2) NOT NULL DEFAULT 0,
    toll_expense NUMERIC(18,2) NOT NULL DEFAULT 0,
    other_expense NUMERIC(18,2) NOT NULL DEFAULT 0,

    transaction_note TEXT,
    is_posted SMALLINT NOT NULL DEFAULT 0,
    created_by INT NOT NULL REFERENCES app_user(app_user_id),
    updated_by INT NOT NULL REFERENCES app_user(app_user_id),
    deleted_by INT REFERENCES app_user(app_user_id),
    created_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMPTZ,
    CONSTRAINT fk_project_financial_transport_project FOREIGN KEY (
        project_transport_id,
        project_id
    ) REFERENCES project_transport(project_transport_id, project_id),
    CONSTRAINT chk_project_financial_transaction_kind CHECK (
        transaction_kind IN (
            'RECEIVABLE',
            'RECEIPT',
            'PAYABLE',
            'PAYMENT',
            'ADJUSTMENT'
        )
    ),
    CONSTRAINT chk_project_financial_due_date CHECK (
        due_date IS NULL OR due_date >= transaction_date
    ),
    CONSTRAINT chk_project_financial_amount_nonnegative CHECK (
        material_sale_income >= 0
        AND transport_service_income >= 0
        AND other_income >= 0
        AND material_purchase_expense >= 0
        AND transport_expense >= 0
        AND road_money_expense >= 0
        AND loading_expense >= 0
        AND unloading_expense >= 0
        AND fuel_expense >= 0
        AND toll_expense >= 0
        AND other_expense >= 0
    ),
    CONSTRAINT chk_project_financial_kind_direction CHECK (
        (
            transaction_kind IN ('RECEIVABLE', 'RECEIPT')
            AND material_purchase_expense
                + transport_expense
                + road_money_expense
                + loading_expense
                + unloading_expense
                + fuel_expense
                + toll_expense
                + other_expense = 0
        )
        OR
        (
            transaction_kind IN ('PAYABLE', 'PAYMENT')
            AND material_sale_income
                + transport_service_income
                + other_income = 0
        )
        OR transaction_kind = 'ADJUSTMENT'
    ),
    CONSTRAINT chk_project_financial_has_amount CHECK (
        material_sale_income
        + transport_service_income
        + other_income
        + material_purchase_expense
        + transport_expense
        + road_money_expense
        + loading_expense
        + unloading_expense
        + fuel_expense
        + toll_expense
        + other_expense > 0
    )
);


CREATE TABLE stockpile_adjustment (
    stockpile_adjustment_id SERIAL PRIMARY KEY,
    project_id INT NOT NULL REFERENCES project(project_id),
    stockpile_cargo_id INT NOT NULL REFERENCES stockpile_cargo(stockpile_cargo_id),
    adjustment_date DATE NOT NULL,
    project_transport_id INT,
    reference_type_id SMALLINT NOT NULL,
    amount_volume_cubic NUMERIC(15,3) NOT NULL DEFAULT 0,
    amount_weight_ton NUMERIC(15,3) NOT NULL DEFAULT 0,
    reason TEXT,
    created_by INT NOT NULL REFERENCES app_user(app_user_id),
    updated_by INT NOT NULL REFERENCES app_user(app_user_id),
    deleted_by INT REFERENCES app_user(app_user_id),
    created_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMPTZ,
    CONSTRAINT fk_stockpile_adjustment_transport_project FOREIGN KEY (
        project_transport_id,
        project_id
    ) REFERENCES project_transport(project_transport_id, project_id),
    CONSTRAINT chk_stockpile_adjustment_reference_type CHECK (
        reference_type_id BETWEEN 1 AND 5
    ),
    CONSTRAINT chk_stockpile_adjustment_has_quantity CHECK (
        amount_volume_cubic <> 0 OR amount_weight_ton <> 0
    )
);

CREATE TABLE stockpile_ledger (
    stockpile_ledger_id SERIAL PRIMARY KEY,
    project_id INT NOT NULL REFERENCES project(project_id),
    stockpile_cargo_id INT NOT NULL REFERENCES stockpile_cargo(stockpile_cargo_id),
    project_transport_id INT,
    adjustment_id INT REFERENCES stockpile_adjustment(stockpile_adjustment_id),
    reference_type_id SMALLINT NOT NULL,
    amount_volume_cubic NUMERIC(15,3) NOT NULL DEFAULT 0,
    balance_volume_cubic NUMERIC(15,3) NOT NULL DEFAULT 0,
    amount_weight_ton NUMERIC(15,3) NOT NULL DEFAULT 0,
    balance_weight_ton NUMERIC(15,3) NOT NULL DEFAULT 0,
    created_by INT NOT NULL REFERENCES app_user(app_user_id),
    updated_by INT NOT NULL REFERENCES app_user(app_user_id),
    deleted_by INT REFERENCES app_user(app_user_id),
    created_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMPTZ,
    CONSTRAINT fk_stockpile_ledger_transport_project FOREIGN KEY (
        project_transport_id,
        project_id
    ) REFERENCES project_transport(project_transport_id, project_id),
    CONSTRAINT chk_stockpile_ledger_reference_type CHECK (
        reference_type_id BETWEEN 1 AND 5
    )
);

-- app_user is created before city/client to break their circular audit
-- dependency, so these two foreign keys are added after all related tables.
ALTER TABLE app_user
ADD CONSTRAINT fk_app_user_city
FOREIGN KEY (city_id) REFERENCES city(city_id);

ALTER TABLE app_user
ADD CONSTRAINT fk_app_user_client
FOREIGN KEY (client_id) REFERENCES client(client_id);

CREATE INDEX idx_project_route_project_id
    ON project_route(project_id);
CREATE UNIQUE INDEX uq_project_checker_active_assignment
    ON project_checker_assignment(project_id, checker_id)
    WHERE deleted_at IS NULL AND is_active = 1;
CREATE UNIQUE INDEX uq_project_checker_one_default
    ON project_checker_assignment(checker_id)
    WHERE deleted_at IS NULL AND is_active = 1 AND is_default = 1;
CREATE INDEX idx_project_checker_checker_access
    ON project_checker_assignment(checker_id, is_active, access_started_at);
CREATE INDEX idx_project_checker_project_id
    ON project_checker_assignment(project_id);
CREATE UNIQUE INDEX uq_project_truck_active_assignment
    ON project_truck_assignment(project_id, truck_id)
    WHERE deleted_at IS NULL AND is_active = 1;
CREATE INDEX idx_project_truck_truck_assignment
    ON project_truck_assignment(truck_id, is_active, assignment_started_at);
CREATE INDEX idx_project_truck_project_id
    ON project_truck_assignment(project_id);
CREATE INDEX idx_project_transport_route_id
    ON project_transport(project_route_id);
CREATE INDEX idx_project_transport_project_date
    ON project_transport(project_id, transported_at DESC)
    WHERE deleted_at IS NULL;
CREATE INDEX idx_project_transport_driver_date
    ON project_transport(driver_id, transported_at DESC)
    WHERE deleted_at IS NULL;
CREATE INDEX idx_project_transport_transported_at
    ON project_transport(transported_at);
CREATE INDEX idx_project_transport_status_transport_id
    ON project_transport_status(project_transport_id);
CREATE UNIQUE INDEX uq_project_transport_status_once
    ON project_transport_status(
        project_transport_id,
        project_transport_status_type_id
    )
    WHERE deleted_at IS NULL;
CREATE INDEX idx_project_transport_status_latest
    ON project_transport_status(
        project_transport_id,
        status_time DESC,
        project_transport_status_id DESC
    )
    WHERE deleted_at IS NULL;
CREATE INDEX idx_project_financial_project_date
    ON project_financial_transaction(project_id, transaction_date);
CREATE INDEX idx_project_financial_transaction_date
    ON project_financial_transaction(transaction_date DESC)
    WHERE deleted_at IS NULL;
CREATE INDEX idx_stockpile_ledger_project_id
    ON stockpile_ledger(project_id);

-- Lookup values used by the application:
--
-- photo_type_id
-- 1 delivery note
-- 2 cargo box
--
-- project_transport_status_type_id
-- 1 arrived at origin
-- 2 loading at origin
-- 3 loaded at origin
-- 4 going to destination
-- 5 arrived at destination
-- 6 unloading at destination
-- 7 unloaded at destination
-- 8 completed
--
-- stockpile reference_type_id
-- 1 project_transport_in
-- 2 project_transport_out
-- 3 stockpile_adjustment
-- 4 manual
-- 5 correction

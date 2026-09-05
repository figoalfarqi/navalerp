-- =============================================================================
-- SEED DATA: MODUL 8 - LOGISTIK & TRANSPORTASI MILITER (LOGISTICS & TRANSPORTATION)
-- FILE: 08_logistics_transportation/insert.sql
-- =============================================================================

-- 1. Armada Angkut Militer (Transport Units)
INSERT INTO log_transport_units (
    transport_unit_id, unit_code, unit_name, transport_type,
    cargo_capacity_tons, fuel_capacity_liters, operating_unit_id, status
) VALUES
(
    '87000000-0000-0000-0000-000000000001',
    'KRI-TRK-905',
    'KRI Tarakan-905 (Kapal Bantu Cair Minyak / BCM)',
    'NAVAL_AUXILIARY_VESSEL',
    5500.00,
    1150000.00,
    '10000000-0000-0000-0000-000000000002', -- Koarmada II
    'AVAILABLE'
),
(
    '87000000-0000-0000-0000-000000000002',
    'KRI-TBN-520',
    'KRI Teluk Bintuni-520 (Kapal Angkut Tank / LST)',
    'NAVAL_AUXILIARY_VESSEL',
    4300.00,
    450000.00,
    '10000000-0000-0000-0000-000000000002',
    'AVAILABLE'
),
(
    '87000000-0000-0000-0000-000000000003',
    'TRK-DISBEKAL-08',
    'Truk Tronton Angkutan Berat Disbekal Surabaya',
    'LAND_TRUCK',
    25.00,
    400.00,
    '10000000-0000-0000-0000-000000000003', -- Lantamal V
    'AVAILABLE'
)
ON CONFLICT (transport_unit_id) DO NOTHING;

-- 2. Rute Distribusi Laut Militer
INSERT INTO log_routes (
    route_id, route_code, route_name, origin_facility_id,
    destination_facility_id, distance_nautical_miles, estimated_transit_hours, risk_level
) VALUES
(
    '87500000-0000-0000-0000-000000000001',
    'RTE-SBY-NATUNA',
    'Jalur Pangkalan Utama Surabaya - Pangkalan Aju Ranai Natuna',
    '85000000-0000-0000-0000-000000000001', -- Dermaga Madura Sby
    '85000000-0000-0000-0000-000000000001', -- (Ref Dermaga Madura as endpoint dummy)
    720.00,
    48.00,
    'HIGH_SEA'
),
(
    '87500000-0000-0000-0000-000000000002',
    'RTE-SBY-TARAKAN',
    'Jalur Pangkalan Utama Surabaya - Lantamal XIII Tarakan',
    '85000000-0000-0000-0000-000000000001',
    '85000000-0000-0000-0000-000000000002',
    850.00,
    56.00,
    'NORMAL'
)
ON CONFLICT (route_id) DO NOTHING;

-- 3. Pengiriman Bebekal / Logistik (Shipments)
INSERT INTO log_shipments (
    shipment_id, manifest_number, route_id, transport_unit_id,
    origin_warehouse_id, destination_warehouse_id, departure_date, arrival_date,
    escort_security_level, status, authorized_by_user_id, remarks
) VALUES
(
    '88000000-0000-0000-0000-000000000001',
    'MAN-LOG-2026-0033',
    '87500000-0000-0000-0000-000000000001',
    '87000000-0000-0000-0000-000000000001', -- KRI Tarakan-905
    '50000000-0000-0000-0000-000000000001', -- Gudang Disbekal Surabaya
    '50000000-0000-0000-0000-000000000004', -- Gudang Kapal KRI REM-331 (titik aju)
    '2026-02-27 10:00:00+07',
    '2026-03-01 14:00:00+07',
    'WARSHIP_ESCORT',
    'DELIVERED',
    '20000000-0000-0000-0000-000000000003', -- Kolonel Aslog
    'Dukungan logistik garis depan (sea replenishment) pelumas dan filter cadangan siaga tempur.'
)
ON CONFLICT (shipment_id) DO NOTHING;

-- 4. Manifest Rincian Muatan Logistik
INSERT INTO log_shipment_items (
    shipment_item_id, shipment_id, material_id, quantity_dispatched,
    quantity_received, packaging_type, weight_kg, notes
) VALUES
(
    '88500000-0000-0000-0000-000000000001',
    '88000000-0000-0000-0000-000000000001',
    '51000000-0000-0000-0000-000000000004', -- Oli Mil-L-9000
    10.00,
    10.00,
    'DRUM',
    2100.00,
    'Drum kedap oli mesin pendorong'
),
(
    '88500000-0000-0000-0000-000000000002',
    '88000000-0000-0000-0000-000000000001',
    '51000000-0000-0000-0000-000000000001', -- Oil Filter MTU
    4.00,
    4.00,
    'CRATE',
    60.00,
    'Peti kayu mil-spec tahan air laut'
)
ON CONFLICT (shipment_item_id) DO NOTHING;


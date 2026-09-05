import { Vendor } from "./vendor.type";

export interface TruckType {
  truck_type_id: number;
  truck_type_name: string;
  truck_box_length: number;
  truck_box_width: number;
  truck_box_height: number;
  truck_capacity: number;
  is_active: number;
  created_at: string;
  updated_at: string;
}

export interface TruckMerk {
  truck_merk_id: number;
  truck_merk_name: string;
  is_active: number;
  created_at: string;
  updated_at: string;
}

export interface Truck {
  truck_id: number;
  truck_type_id: number;
  truck_merk_id: number;
  driver_id: number;
  vendor_id: number;
  license_plate: string;
  ownership_status_id: number;
  production_year: number;
  number_of_tires: number;
  is_active: number;
  created_at: string;
  updated_at: string;

  truck_type: TruckType;
  truck_merk: TruckMerk;
  vendor: Vendor;
}
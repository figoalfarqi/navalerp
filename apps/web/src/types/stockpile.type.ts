import { CargoType } from "./cargoType.type";
import { City } from "./location.type";

export interface Stockpile {
  stockpile_id: number;
  stockpile_name: string;
  city_id: number;
  stockpile_address: string;
  is_active: number;
  created_at: string;
  updated_at: string;
  city: City;
}

export interface StockpileCargo {
  stockpile_cargo_id: number;
  stockpile_id: number;
  cargo_type_id: number;
  current_volume: number;
  is_active: number;
  created_at: string;
  updated_at: string;

  stockpile: Stockpile;
  cargo_type: CargoType;
}
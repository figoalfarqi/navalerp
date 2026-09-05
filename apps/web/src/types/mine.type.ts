import { City } from "./location.type";

export interface Mine {
  mine_id: number;
  mine_name: string;
  city_id: number;
  mine_address: string;
  mine_latitude: number;
  mine_longitude: number;
  mine_map_url: string;
  is_active: number;
  created_at: string;
  updated_at: string;
  city: City;
}
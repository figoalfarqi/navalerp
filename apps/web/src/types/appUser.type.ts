import { City } from "./location.type";

export interface AppUser {
  app_user_id: number;
  username: string;
  app_user_status_id: number;
  app_user_name: string;
  app_user_phone: string;
  city_id: number;
  created_by: number;
  updated_by: number;
  created_at: string;
  updated_at: string;
  city: City;
}
export interface Province {
  province_id: number;
  province_name: string;
  province_real_name: string;
  is_active: number;
  created_at: string;
  updated_at: string;
}

export interface City {
  city_id: number;
  province_id: number;
  city_name: string;
  is_active: number;
  created_at: string;
  updated_at: string;
  province: Province;
}
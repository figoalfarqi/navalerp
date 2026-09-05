export type ProjectRouteType =
  | "MINE_CLIENT"
  | "MINE_STOCKPILE_CLIENT"
  | "VESSEL_CLIENT"
  | "VESSEL_STOCKPILE_CLIENT";

export type ProjectStatus = "DRAFT" | "ACTIVE" | "COMPLETED" | "CANCELLED";

export interface Project {
  project_id: number;
  project_code: string;
  project_name: string;
  route_type: ProjectRouteType;
  mine_id?: number | null;
  vessel_cargo_id?: number | null;
  stockpile_cargo_id?: number | null;
  client_destination_id: number;
  cargo_type_id: number;
  project_status: ProjectStatus;
  start_date?: string | null;
  end_date?: string | null;
  planned_volume_cubic?: number | null;
  planned_weight_ton?: number | null;
  volume_to_weight_conversion?: number | null;
  project_note?: string | null;
  is_active: number;
  created_at?: string;
  updated_at?: string;

  is_default?: number;
  mine?: {
    mine_id: number;
    mine_name: string;
  } | null;
  vessel_cargo?: {
    vessel_cargo_id: number;
    voyage_number?: string | null;
    vessel?: {
      vessel_id: number;
      vessel_name: string;
    } | null;
    port?: {
      port_id: number;
      port_name: string;
    } | null;
  } | null;
  stockpile_cargo?: {
    stockpile_cargo_id: number;
    stockpile?: {
      stockpile_id: number;
      stockpile_name: string;
    } | null;
  } | null;
  client_destination?: {
    client_destination_id: number;
    client_destination_name: string;
    client?: {
      client_id: number;
      client_name: string;
    } | null;
  } | null;
  cargo_type?: {
    cargo_type_id: number;
    cargo_type_name: string;
    cargo_type_grade?: string;
  } | null;
}

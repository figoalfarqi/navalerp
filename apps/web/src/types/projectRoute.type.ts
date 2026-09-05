export type ProjectRouteLegType =
  | "SOURCE_TO_CLIENT"
  | "SOURCE_TO_STOCKPILE"
  | "STOCKPILE_TO_CLIENT";

export interface ProjectRoute {
  project_route_id: number;
  project_id: number;
  project_pattern: string;
  route_sequence: number;
  route_type: ProjectRouteLegType;
  route_name?: string | null;
  distance_km?: number | null;
  route_note?: string | null;
  is_active: number;
  created_at?: string;
  updated_at?: string;
}

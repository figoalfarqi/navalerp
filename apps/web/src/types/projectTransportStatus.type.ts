import { AppUser } from "./appUser.type";

export type ProjectTransportStatusTypeId = 1 | 2 | 3 | 4 | 5 | 6 | 7 | 8;

export interface ProjectTransportStatus {
  project_transport_status_id: number;
  project_transport_id: number;
  project_transport_status_type_id: ProjectTransportStatusTypeId;
  status_time: string;
  cargo_box_length?: number | null;
  cargo_box_width?: number | null;
  cargo_box_height?: number | null;
  cargo_volume_cubic?: number | null;
  cargo_weight_ton?: number | null;
  project_transport_status_note?: string | null;
  is_fraud: number;
  is_active: number;
  created_by?: number;
  updated_by?: number;
  created_at?: string;
  updated_at?: string;
  created_by_user?: AppUser;
}

export interface CreateProjectTransportStatusPayload {
  project_transport_id: number;
  project_transport_status_type_id: ProjectTransportStatusTypeId;
  status_time: string;
  cargo_volume_cubic?: number;
  cargo_weight_ton?: number;
  project_transport_status_note?: string;
  is_fraud: 0 | 1;
}

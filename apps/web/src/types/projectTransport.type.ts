import { AppUser } from "./appUser.type";
import { Project } from "./project.type";
import { ProjectRoute } from "./projectRoute.type";
import { ProjectTransportStatus } from "./projectTransportStatus.type";
import { ProjectTruckAssignment } from "./projectTruckAssignment.type";
import { Truck } from "./truck.type";
import { Vendor } from "./vendor.type";

export interface ProjectTransport {
  project_transport_id: number;
  project_id: number;
  project_route_id: number;
  transport_number: string;
  delivery_note_number?: string | null;
  transported_at: string;
  truck_id?: number | null;
  project_truck_assignment_id?: number | null;
  driver_id?: number | null;
  transport_vendor_id?: number | null;
  loaded_volume_cubic?: number | null;
  loaded_weight_ton?: number | null;
  delivered_volume_cubic?: number | null;
  delivered_weight_ton?: number | null;
  volume_to_weight_conversion?: number | null;
  is_completed: number;
  transport_note?: string | null;
  created_at?: string;
  updated_at?: string;

  // Read-only checker/driver DTOs intentionally expose labels as flat fields
  // instead of returning the full admin object graph.
  project_code?: string;
  project_name?: string;
  route_name?: string | null;
  route_type?: string;
  license_plate?: string | null;
  origin_name?: string | null;
  destination_name?: string | null;
  distance_km?: number | null;
  origin_location_name?: string | null;
  destination_location_name?: string | null;

  project?: Project;
  project_route?: ProjectRoute;
  truck?: Truck;
  project_truck_assignment?: ProjectTruckAssignment;
  driver?: AppUser;
  transport_vendor?: Vendor;
  project_transport_statuses?: ProjectTransportStatus[];
  statuses?: ProjectTransportStatus[];
}

export interface CreateProjectTransportPayload {
  project_id: number;
  project_route_id: number;
  transport_number: string;
  transported_at: string;
  truck_id: number;
  project_truck_assignment_id: number;
  driver_id?: number;
  transport_vendor_id?: number;
}

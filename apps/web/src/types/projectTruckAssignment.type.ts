import { Project } from "./project.type";
import { Truck } from "./truck.type";

export interface ProjectTruckAssignment {
  project_truck_assignment_id: number;
  project_id: number;
  truck_id: number;
  assignment_started_at: string;
  assignment_ended_at?: string | null;
  assignment_note?: string | null;
  is_active: number;
  license_plate?: string;
  driver_id?: number | null;
  vendor_id?: number | null;
  project?: Project;
  truck?: Truck;
}

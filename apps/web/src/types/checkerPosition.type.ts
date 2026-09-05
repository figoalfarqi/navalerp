import { ProjectRouteType } from "./project.type";

export type CheckerLocation = "mine" | "vessel" | "stockpile" | "client";

export interface CheckerPosition {
  project_id: number;
  project_code: string;
  project_name: string;
  project_route_type: ProjectRouteType;
  location: CheckerLocation;
  saved_at: number;
  location_locked_until: number;
}

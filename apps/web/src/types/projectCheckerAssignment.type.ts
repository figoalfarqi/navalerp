import { Project } from "./project.type";

export interface ProjectCheckerAssignment {
  project_checker_assignment_id: number;
  project_id: number;
  checker_id: number;
  access_started_at: string;
  access_ended_at?: string | null;
  assignment_note?: string | null;
  is_active: number;
  is_default: number;
  project?: Project;
}

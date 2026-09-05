export interface ProjectTransportPhoto {
  project_transport_photo_id: number;
  project_transport_status_id: number;
  photo_url: string;
  photo_type_id: number;
  photo_description?: string | null;
  created_at?: string;
  updated_at?: string;
}

export interface CreateProjectTransportPhotoPayload {
  project_transport_status_id: number;
  photo_url: string;
  photo_type_id: number;
  photo_description?: string;
}

import { Client } from "./client.type";

export interface ClientDestination {
  client_destination_id: number;
  client_destination_name: string;
  client_id: string;
  is_active: number;
  created_at: string;
  updated_at: string;
  client: Client;
}

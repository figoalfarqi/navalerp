import { ChainItem } from "@/components/form/ChainSelectField";

export const LOCATION_CHAINS: ChainItem[] = [
  {
    name: "city_id",
    label: "City",
    url: `${process.env.NEXT_PUBLIC_API_BASE_URL}/admin/city?limit=399`,
    labelKey: "city_name",
    valueKey: "city_id",
  },
  {
    name: "district_id",
    label: "District",
    url: `${process.env.NEXT_PUBLIC_API_BASE_URL}/admin/district?limit=99&query=city_id%3D{city_id}`,
    labelKey: "district_name",
    valueKey: "district_id",
  },
] as const;

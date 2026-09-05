import {
  pickAdminPayload,
  type AdminCrudMode,
} from "@/components/admin/crud/adminCrud";
import type {
  FormDataObject,
  FormField,
} from "@/components/formCrud/FormCrud";
import type { FilterField } from "@/components/table/FilterFormTable";
import type { ColumnField } from "@/components/table/Table";
import { formatDateTime } from "@/utils/dateTime";
import { resolveFileUrl } from "@/utils/globalUtils";
import {
  getTransportPhotoTypeLabel,
  getTransportStatusLabel,
  transportPhotoTypeOptions as photoTypeOptions,
} from "../project_transport/_labels";

export const projectTransportPhotoEndpoint = `${process.env.NEXT_PUBLIC_API_BASE_URL}/admin/project_transport_photo`;
export const projectTransportPhotoTableName = "project_transport_photo";
export const projectTransportPhotoPrimaryKey = "project_transport_photo_id";

export const projectTransportPhotoColumns: ColumnField[] = [
  {
    key: "transport_number",
    label: "Transport",
    render: (item) =>
      item.project_transport_status?.project_transport?.transport_number ?? "-",
    columnLength: 180,
  },
  {
    key: "project_name",
    label: "Project",
    render: (item) =>
      item.project_transport_status?.project_transport?.project?.project_name ??
      "-",
    columnLength: 220,
  },
  {
    key: "license_plate",
    label: "Truck",
    render: (item) =>
      item.project_transport_status?.project_transport?.truck?.license_plate ??
      "-",
    columnLength: 140,
  },
  {
    key: "project_transport_status_type_id",
    label: "Tahap Status",
    render: (item) =>
      getTransportStatusLabel(
        item.project_transport_status?.project_transport_status_type_id,
      ),
    columnLength: 200,
  },
  {
    key: "photo_type_id",
    label: "Tipe Foto",
    render: (item) => getTransportPhotoTypeLabel(item.photo_type_id),
    sortable: "number",
    columnLength: 160,
  },
  {
    key: "photo_url",
    label: "Foto",
    render: (item) => {
      const photoURL =
        typeof item.photo_url === "string"
          ? resolveFileUrl(item.photo_url)
          : "";
      return photoURL ? (
        <a
          href={photoURL}
          target="_blank"
          rel="noreferrer"
          className="font-medium text-blue-600 hover:underline"
          onClick={(event) => event.stopPropagation()}
        >
          Lihat foto
        </a>
      ) : (
        "-"
      );
    },
    columnLength: 130,
  },
  {
    key: "photo_description",
    label: "Deskripsi",
    sortable: "string",
    columnLength: 260,
  },
  {
    key: "created_at",
    label: "Dibuat",
    render: (item) => formatDateTime(item.created_at),
    sortable: "date",
    columnLength: 180,
  },
];

export const projectTransportPhotoFilters: FilterField[] = [
  {
    name: "project_transport_status_id",
    label: "Status Transport",
    fieldType: "select",
    options: {
      url: `${process.env.NEXT_PUBLIC_API_BASE_URL}/admin/project_transport_status?limit=999`,
      labelKey:
        "{project_transport.transport_number} · Status {project_transport_status_type_id}",
      valueKey: "project_transport_status_id",
    },
    col: "left",
  },
  {
    name: "photo_type_id",
    label: "Tipe Foto",
    fieldType: "select",
    options: photoTypeOptions,
    col: "right",
  },
  {
    name: "created_at",
    label: "Tanggal Foto",
    fieldType: "dateAfterBefore",
    col: "left",
  },
];

export function projectTransportPhotoFields(
  mode: AdminCrudMode,
): FormField[] {
  const disabled = mode === "view";
  return [
    {
      name: "project_transport_status_id",
      label: "Status Transport",
      fieldType: "select",
      options: {
        url: `${process.env.NEXT_PUBLIC_API_BASE_URL}/admin/project_transport_status?limit=999`,
        labelKey:
          "{project_transport.transport_number} · Status {project_transport_status_type_id}",
        valueKey: "project_transport_status_id",
      },
      required: true,
      disabled,
      col: "left",
    },
    {
      name: "photo_type_id",
      label: "Tipe Foto",
      fieldType: "select",
      options: photoTypeOptions,
      required: true,
      disabled,
      col: "right",
    },
    {
      name: "photo_url",
      label: "Foto",
      fieldType: "uploadimage",
      fileFolder: "/project-transport",
      required: true,
      disabled,
      col: "left",
    },
    {
      name: "photo_description",
      label: "Deskripsi Foto",
      fieldType: "textarea",
      disabled,
      col: "right",
    },
  ];
}

const payloadFields = [
  "project_transport_status_id",
  "photo_url",
  "photo_type_id",
  "photo_description",
] as const;

export const buildProjectTransportPhotoPayload = (data: FormDataObject) =>
  pickAdminPayload(data, payloadFields);

"use client";
import { FilterField } from "@/components/table/FilterFormTable";
import Table, { ColumnField } from "@/components/table/Table";
import { formatDateTime } from "@/utils/dateTime";

export default function Page() {
  const columns: ColumnField[] = [
    {
      key: "truck_type_name",
      label: "Truck Type Name",
      columnLength: 200,
      sortable: "string",
    },
    {
      key: "truck_type_description",
      label: "Truck Type Description",
      columnLength: 200,
      sortable: "string",
    },
    {
      key: "truck_box_length",
      label: "Truck Box Length (m)",
      columnLength: 200,
      sortable: "number",
    },
    {
      key: "truck_box_width",
      label: "Truck Box Width (m)",
      columnLength: 200,
      sortable: "number",
    },
    {
      key: "truck_box_height",
      label: "Truck Box Height (m)",
      columnLength: 200,
      sortable: "number",
    },
    {
      key: "truck_capacity",
      label: "Truck Capacity (ton)",
      columnLength: 200,
      sortable: "number",
    },
    {
      key: "is_active",
      label: "Is Active",
      render: (val) =>
        val.is_active == "1" ? (
          <div className="text-green-500">Aktif</div>
        ) : (
          <div className="text-red-500">Tidak Aktif</div>
        ),
      columnLength: 200,
      sortable: "boolean",
    },
    {
      key: "created_at",
      label: "Created At",
      render: (val) => formatDateTime(val.created_at),
      columnLength: 200,
      sortable: "date",
    },
    {
      key: "updated_at",
      label: "Updated At",
      render: (val) => formatDateTime(val.updated_at),
      columnLength: 200,
      sortable: "date",
    },
  ];

  const filters: FilterField[] = [
    {
      name: "truck_type_name",
      label: "Truck Type Name",
      fieldType: "text",
      col: "left",
    },
    {
      name: "truck_box_length",
      label: "Truck Box Length (m)",
      fieldType: "number",
      col: "left",
    },
    {
      name: "truck_box_width",
      label: "Truck Box Width (m)",
      fieldType: "number",
      col: "left",
    },
    {
      name: "truck_box_height",
      label: "Truck Box Height (m)",
      fieldType: "number",
      col: "left",
    },
    {
      name: "truck_capacity",
      label: "Truck Capacity (ton)",
      fieldType: "number",
      col: "left",
    },
    {
      name: "updated_at",
      label: "Updated",
      fieldType: "dateAfterBefore",
      col: "left",
    },
    {
      name: "created_at",
      label: "Created",
      fieldType: "dateAfterBefore",
      col: "right",
    },
    {
      name: "is_active",
      label: "Is Active",
      fieldType: "radio",
      col: "right",
      options: [
        { label: "Aktif", value: "1", variant: "green-outline" },
        { label: "Tidak Aktif", value: "0", variant: "red-outline" },
        { label: "Semua", value: "2", variant: "purple-outline" },
      ],
    },
  ];
  return (
    <Table
      url={`${process.env.NEXT_PUBLIC_API_BASE_URL}/admin/truck_type`}
      title="Truck Type List"
      table_name="truck_type"
      columns={columns}
      filters={filters}
    />
  );
}

/* eslint-disable @typescript-eslint/no-explicit-any */
"use client";
import { FilterField } from "@/components/table/FilterFormTable";
import Table, { ColumnField } from "@/components/table/Table";
import { useFetchAPI } from "@/hooks/useFetchAPI";
import { formatDateTime } from "@/utils/dateTime";

export default function Page() {
  const { getAPI } = useFetchAPI();
  const columns: ColumnField[] = [
    {
      key: "province_id",
      label: "Province Name",
      render: (val) => val.province?.province_name ?? "-",
      columnLength: 200,
      sortable:"table_key",
    },
    {
      key: "city_name",
      label: "City Name",
      columnLength: 200,
      sortable:"string",
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
      sortable:"date",
    },
    {
      key: "updated_at",
      label: "Updated At",
      render: (val) => formatDateTime(val.updated_at),
      columnLength: 200,
      sortable:"date",
    },
  ];

  const filters: FilterField[] = [
    {
      name: "province_id",
      label: "Province",
      fieldType: "select",
      optionsFetcher: async () => {
        const res = await getAPI<any>(
          `${process.env.NEXT_PUBLIC_API_BASE_URL}/admin/province?limit=99`,
          {
            authToken: "admin",
          }
        );
        return res.data.items.map((item: any) => ({
          label: item.province_name,
          value: item.province_id,
        }));
      },
      col: "left",
    },
    {
      name: "city_name",
      label: "City",
      fieldType: "text",
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
      url={`${process.env.NEXT_PUBLIC_API_BASE_URL}/admin/city`}
      title="City List"
      table_name="city"
      columns={columns}
      filters={filters}
    />
  );
}

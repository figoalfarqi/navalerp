"use client";
import Table, { ColumnField } from "@/components/table/Table";
import { ApprovalStatusMap } from "@/consta/ApprovalStatusMap";
import { formatDateTime } from "@/utils/dateTime";
import { toTitleCase } from "@/utils/globalUtils";

export type ApprovalBaseTableNameType = "delivery_order_note_photo" | "delivery_order_cargo_photo" | "delivery_order_travel_allowance"


const ApprovalTable = ({
  approvalData,
  approvalBaseTableName,
}: {
  approvalData: Record<string, any>[];
  approvalBaseTableName: ApprovalBaseTableNameType;
}) => {
  //    {
  //     "delivery_order_note_photo_approval_id": 1,
  //     "delivery_order_note_photo_id": 3,
  //     "approval_step_id": 1,
  //     "approval_status_id": 1,
  //     "approval_notes": "asdasd asd asdasd s",
  //     "created_by": 1,
  //     "updated_by": 1,
  //     "created_at": "2026-01-29T19:37:33.39796+07:00",
  //     "updated_at": "2026-01-29T19:37:33.39796+07:00",
  //     "created_by_app_user": {
  //         "app_user_id": 1,
  //         "username": "itdev",
  //         "app_user_status_id": 1,
  //         "app_user_name": "IT DEVELOPER",
  //         "app_user_preferred_name": "IT DEV",
  //         "created_by": 0,
  //         "updated_by": 0,
  //         "created_at": "2026-01-21T17:27:12.801372+07:00",
  //         "updated_at": "2026-01-21T17:27:12.801372+07:00"
  //     }
  // }
  const approvalDataColumns: ColumnField[] = [
    {
      key: "approval_step_id",
      render: (item: any) => item.approval_step_id ===0 ? "Request" :`Tahap-${item.approval_step_id}`,
      label: "Step",
      columnLength: 120,
    },
    {
      key: "approval_status_id",
      render: (item: any) => (
        <div
          className="w-fit px-2"
          style={{
            color: ApprovalStatusMap[item.approval_status_id].color,
            backgroundColor: ApprovalStatusMap[item.approval_status_id].bgColor,
          }}
        >
          {ApprovalStatusMap[item.approval_status_id].value}
        </div>
      ),
      label: "Status",
      columnLength: 120,
    },
    {
      key: "approval_notes",
      label: "Approval Notes",
      columnLength: 200,
    },
    {
      key: "app_user_name",
      render: (item: any) => item.created_by_app_user?.app_user_name || "-",
      label: "Approved By",
      columnLength: 200,
    },
    {
      key: "created_at",
      render: (item: any) => formatDateTime(item.created_at),
      label: "Approved At",
      columnLength: 200,
    },
  ];
  return (
    <div>
      <Table
        url={``}
        table_data={approvalData}
        title={`${toTitleCase(approvalBaseTableName.replaceAll("_", " "))} Approval List`}
        tableFor="detail"
        table_name={`${approvalBaseTableName}_approval`}
        columns={approvalDataColumns}
        filters={[]}
        default_filter_values={{
          delivery_order_for_page: "active",
        }}
      />
    </div>
  );
};

export default ApprovalTable;

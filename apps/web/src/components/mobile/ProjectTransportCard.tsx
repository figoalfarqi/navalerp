import ItemCard from "@/components/mobile/ItemCard";
import { ProjectTransport } from "@/types/projectTransport.type";
import {
  ProjectTransportStatus,
  ProjectTransportStatusTypeId,
} from "@/types/projectTransportStatus.type";
import { formatDateTime } from "@/utils/dateTime";

const STATUS_STYLES: Record<
  ProjectTransportStatusTypeId,
  { label: string; className: string }
> = {
  1: { label: "Tiba di asal", className: "bg-blue-50 text-blue-700" },
  2: { label: "Muat di asal", className: "bg-amber-50 text-amber-700" },
  3: { label: "Selesai muat", className: "bg-emerald-50 text-emerald-700" },
  4: { label: "Menuju tujuan", className: "bg-violet-50 text-violet-700" },
  5: { label: "Tiba di tujuan", className: "bg-cyan-50 text-cyan-700" },
  6: { label: "Bongkar di tujuan", className: "bg-orange-50 text-orange-700" },
  7: { label: "Selesai bongkar", className: "bg-green-50 text-green-700" },
  8: { label: "Selesai", className: "bg-slate-100 text-slate-700" },
};

export const getProjectTransportStatusLabel = (
  statusTypeId: ProjectTransportStatusTypeId,
) => STATUS_STYLES[statusTypeId].label;

export const getTransportStatuses = (transport: ProjectTransport) =>
  transport.project_transport_statuses ?? transport.statuses ?? [];

export const getLatestTransportStatus = (
  transport: ProjectTransport,
): ProjectTransportStatus | undefined =>
  [...getTransportStatuses(transport)].sort((a, b) => {
    const timeDifference =
      new Date(b.status_time).getTime() - new Date(a.status_time).getTime();
    if (timeDifference !== 0) return timeDifference;
    return b.project_transport_status_id - a.project_transport_status_id;
  })[0];

const formatMeasurement = (
  value: number | null | undefined,
  unit: string,
) => (value === null || value === undefined ? "-" : `${value} ${unit}`);

export default function ProjectTransportCard({
  transport,
  onClick,
  showProject = true,
}: {
  transport: ProjectTransport;
  onClick?: () => void;
  showProject?: boolean;
}) {
  const latestStatus = getLatestTransportStatus(transport);
  const fraudStatus = getTransportStatuses(transport).find(
    (status) => status.is_fraud === 1,
  );
  const statusStyle = latestStatus
    ? STATUS_STYLES[latestStatus.project_transport_status_type_id]
    : undefined;
  const weight =
    transport.delivered_weight_ton ?? transport.loaded_weight_ton ?? null;
  const volume =
    transport.delivered_volume_cubic ?? transport.loaded_volume_cubic ?? null;

  return (
    <ItemCard onClick={onClick}>
      <div className="space-y-3">
        <div className="flex items-start justify-between gap-3">
          <div className="min-w-0">
            <p className="truncate text-base font-semibold text-slate-900">
              {transport.truck?.license_plate ||
                transport.license_plate ||
                transport.transport_number}
            </p>
            {showProject && (
              <p className="truncate text-xs text-slate-500">
                {transport.project?.project_name ||
                  transport.project_name ||
                  transport.project_code ||
                  `Project #${transport.project_id}`}
              </p>
            )}
          </div>
          {statusStyle && (
            <span
              className={`shrink-0 rounded-full px-2.5 py-1 text-[11px] font-semibold ${statusStyle.className}`}
            >
              {statusStyle.label}
            </span>
          )}
        </div>

        <div className="rounded-lg bg-slate-50 px-3 py-2">
          <p className="text-[11px] font-medium uppercase tracking-wide text-slate-400">
            Rute
          </p>
          <p className="mt-0.5 font-medium text-slate-700">
            {transport.project_route?.route_name ||
              transport.route_name ||
              transport.project_route?.route_type?.replaceAll("_", " ") ||
              transport.route_type?.replaceAll("_", " ") ||
              `Rute #${transport.project_route_id}`}
          </p>
        </div>

        <div className="grid grid-cols-2 gap-2 text-xs">
          <div>
            <p className="text-slate-400">Berat</p>
            <p className="font-semibold text-slate-700">
              {formatMeasurement(weight, "ton")}
            </p>
          </div>
          <div>
            <p className="text-slate-400">Volume</p>
            <p className="font-semibold text-slate-700">
              {formatMeasurement(volume, "m³")}
            </p>
          </div>
        </div>

        <div className="flex items-center justify-between border-t border-slate-100 pt-2 text-[11px] text-slate-400">
          <span>{formatDateTime(transport.transported_at)}</span>
          {fraudStatus && (
            <span className="font-semibold text-red-600">Terindikasi fraud</span>
          )}
        </div>
      </div>
    </ItemCard>
  );
}

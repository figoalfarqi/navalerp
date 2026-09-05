import { formatNumberID } from "@/utils/currencyFormater";
import { formatDateTime } from "@/utils/dateTime";
import { resolveFileUrl } from "@/utils/globalUtils";
import {
  getTransportPhotoTypeLabel,
  getTransportStatusLabel,
} from "./_labels";

type UnknownRecord = Record<string, unknown>;

function isRecord(value: unknown): value is UnknownRecord {
  return Boolean(value) && typeof value === "object" && !Array.isArray(value);
}

function optionalText(value: unknown): string | null {
  if (typeof value !== "string") return null;
  const trimmed = value.trim();
  return trimmed || null;
}

function optionalNumber(value: unknown): number | null {
  if (value === null || value === undefined || value === "") return null;
  const numberValue = Number(value);
  return Number.isFinite(numberValue) ? numberValue : null;
}

function dateValue(value: unknown): string | Date {
  return typeof value === "string" || value instanceof Date ? value : "";
}

function statusPhotos(status: UnknownRecord): UnknownRecord[] {
  return Array.isArray(status.photos) ? status.photos.filter(isRecord) : [];
}

export default function TransportActivityDetail({
  record,
  embedded = false,
}: {
  record: UnknownRecord;
  embedded?: boolean;
}) {
  const statuses = Array.isArray(record.statuses)
    ? record.statuses.filter(isRecord)
    : [];

  return (
    <section
      className={
        embedded
          ? "overflow-hidden rounded-xl border border-slate-200 bg-white shadow-sm"
          : "mx-3 mb-6 mt-4 overflow-hidden rounded-2xl border border-slate-200 bg-white shadow-sm"
      }
    >
      <header className="border-b border-slate-200 px-5 py-4 sm:px-6">
        <p className="text-xs font-semibold uppercase tracking-[0.16em] text-blue-600">
          Aktivitas Transport
        </p>
        <div className="mt-1 flex flex-col justify-between gap-2 sm:flex-row sm:items-end">
          <div>
            <h2 className="text-xl font-semibold text-slate-900">
              Status dan foto transport
            </h2>
            <p className="mt-1 text-sm text-slate-500">
              Riwayat ditampilkan berurutan dari aktivitas paling awal.
            </p>
          </div>
          <p className="text-sm font-medium text-slate-600">
            {statuses.length} status · {Number(record.photo_count ?? 0)} foto
          </p>
        </div>
      </header>

      {statuses.length === 0 ? (
        <div className="px-6 py-12 text-center text-sm text-slate-500">
          Belum ada status atau foto untuk transport ini.
        </div>
      ) : (
        <ol className="divide-y divide-slate-100">
          {statuses.map((status, index) => {
            const photos = statusPhotos(status);
            const volume = optionalNumber(status.cargo_volume_cubic);
            const weight = optionalNumber(status.cargo_weight_ton);
            const note = optionalText(status.project_transport_status_note);
            const fraud = Number(status.is_fraud) === 1;

            return (
              <li
                key={String(
                  status.project_transport_status_id ?? `status-${index}`,
                )}
                className="relative px-5 py-5 sm:px-6"
              >
                <div className="flex gap-4">
                  <div className="flex w-8 shrink-0 flex-col items-center">
                    <span
                      className={`grid h-8 w-8 place-items-center rounded-full text-xs font-bold ${
                        fraud
                          ? "bg-red-100 text-red-700"
                          : "bg-blue-100 text-blue-700"
                      }`}
                    >
                      {index + 1}
                    </span>
                    {index < statuses.length - 1 && (
                      <span className="mt-2 h-full w-px bg-slate-200" />
                    )}
                  </div>

                  <div className="min-w-0 flex-1">
                    <div className="flex flex-col justify-between gap-2 sm:flex-row sm:items-start">
                      <div>
                        <h3 className="font-semibold text-slate-900">
                          {getTransportStatusLabel(
                            status.project_transport_status_type_id,
                          )}
                        </h3>
                        <p className="mt-1 text-xs text-slate-500">
                          {formatDateTime(dateValue(status.status_time))}
                        </p>
                      </div>
                      <div className="flex flex-wrap gap-2">
                        {fraud && (
                          <span className="rounded-full bg-red-100 px-2.5 py-1 text-xs font-semibold text-red-700">
                            Terindikasi fraud
                          </span>
                        )}
                        {Number(status.is_active) === 0 && (
                          <span className="rounded-full bg-slate-100 px-2.5 py-1 text-xs font-medium text-slate-600">
                            Tidak aktif
                          </span>
                        )}
                      </div>
                    </div>

                    {(volume !== null || weight !== null) && (
                      <dl className="mt-4 grid gap-3 sm:grid-cols-2">
                        {volume !== null && (
                          <div className="rounded-xl bg-slate-50 px-3 py-2">
                            <dt className="text-xs text-slate-500">
                              Volume cargo
                            </dt>
                            <dd className="mt-0.5 font-medium text-slate-800">
                              {formatNumberID(volume)} m³
                            </dd>
                          </div>
                        )}
                        {weight !== null && (
                          <div className="rounded-xl bg-slate-50 px-3 py-2">
                            <dt className="text-xs text-slate-500">
                              Berat cargo
                            </dt>
                            <dd className="mt-0.5 font-medium text-slate-800">
                              {formatNumberID(weight)} ton
                            </dd>
                          </div>
                        )}
                      </dl>
                    )}

                    {note && (
                      <p className="mt-3 rounded-xl border border-amber-100 bg-amber-50 px-3 py-2 text-sm text-amber-900">
                        {note}
                      </p>
                    )}

                    {photos.length > 0 && (
                      <div className="mt-4 grid gap-3 sm:grid-cols-2 lg:grid-cols-3 xl:grid-cols-4">
                        {photos.map((photo, photoIndex) => {
                          const rawURL = optionalText(photo.photo_url);
                          const photoURL = rawURL
                            ? resolveFileUrl(rawURL)
                            : "";
                          return (
                            <a
                              key={String(
                                photo.project_transport_photo_id ??
                                  `photo-${photoIndex}`,
                              )}
                              href={photoURL || undefined}
                              target="_blank"
                              rel="noreferrer"
                              className={`overflow-hidden rounded-xl border border-slate-200 bg-slate-50 ${
                                photoURL
                                  ? "transition hover:-translate-y-0.5 hover:border-blue-300 hover:shadow-md"
                                  : "pointer-events-none"
                              }`}
                            >
                              <div
                                className="aspect-[4/3] bg-slate-100 bg-cover bg-center"
                                style={
                                  photoURL
                                    ? {
                                        backgroundImage: `url(${JSON.stringify(
                                          photoURL,
                                        )})`,
                                      }
                                    : undefined
                                }
                              />
                              <div className="px-3 py-2">
                                <p className="text-xs font-semibold text-slate-700">
                                  {getTransportPhotoTypeLabel(
                                    photo.photo_type_id,
                                  )}
                                </p>
                                <p className="mt-0.5 line-clamp-2 text-xs text-slate-500">
                                  {optionalText(photo.photo_description) ??
                                    "Tanpa deskripsi"}
                                </p>
                              </div>
                            </a>
                          );
                        })}
                      </div>
                    )}
                  </div>
                </div>
              </li>
            );
          })}
        </ol>
      )}
    </section>
  );
}

import Link from "next/link";
import { BiHistory } from "react-icons/bi";
import { GiPositionMarker } from "react-icons/gi";

export default function MorePage() {
  const menuItems = [
    {
      name: "Riwayat laporan",
      description: "Lihat laporan transport yang telah dicatat",
      icon: <BiHistory size={28} />,
      path: "/checker/transport_history",
    },
    {
      name: "Posisi kerja",
      description: "Pilih proyek dan lokasi checker saat ini",
      icon: <GiPositionMarker size={28} />,
      path: "/checker/position",
    },
  ];

  return (
    <main className="mx-auto min-h-dvh max-w-md bg-slate-50 px-4 pb-28 pt-5">
      <p className="text-xs font-semibold uppercase tracking-[0.15em] text-blue-600">
        Checker
      </p>
      <h1 className="mt-1 text-2xl font-bold text-slate-900">Menu lainnya</h1>
      <div className="mt-5 space-y-3">
        {menuItems.map((item) => (
          <Link
            key={item.path}
            href={item.path}
            className="flex min-h-20 items-center gap-4 rounded-2xl border border-slate-200 bg-white p-4 shadow-sm transition active:bg-slate-50"
          >
            <span className="flex h-12 w-12 shrink-0 items-center justify-center rounded-xl bg-blue-50 text-blue-600">
              {item.icon}
            </span>
            <span>
              <span className="block text-sm font-bold text-slate-900">
                {item.name}
              </span>
              <span className="mt-1 block text-xs leading-5 text-slate-500">
                {item.description}
              </span>
            </span>
          </Link>
        ))}
      </div>
    </main>
  );
}

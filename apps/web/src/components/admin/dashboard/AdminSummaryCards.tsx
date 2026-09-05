import {
  FaBoxesStacked,
  FaCircleCheck,
  FaMoneyBillTrendUp,
  FaRoute,
  FaScaleBalanced,
  FaTruck,
} from "react-icons/fa6";
import type { AdminDashboardSummary } from "@/types/adminDashboard";
import {
  formatCurrencyIDR,
  formatNumberID,
} from "@/utils/currencyFormater";

interface AdminSummaryCardsProps {
  summary: AdminDashboardSummary;
}

export default function AdminSummaryCards({
  summary,
}: AdminSummaryCardsProps) {
  const cards = [
    {
      label: "Project",
      value: formatNumberID(summary.project_count),
      helper: "project pada filter",
      icon: FaRoute,
      color: "bg-blue-50 text-blue-700",
    },
    {
      label: "Transport",
      value: formatNumberID(summary.transport_count),
      helper: `${formatNumberID(summary.completed_transport_count)} selesai`,
      icon: FaTruck,
      color: "bg-cyan-50 text-cyan-700",
    },
    {
      label: "Volume",
      value: `${formatNumberID(summary.volume_cubic)} m³`,
      helper: "realisasi volume",
      icon: FaBoxesStacked,
      color: "bg-violet-50 text-violet-700",
    },
    {
      label: "Berat",
      value: `${formatNumberID(summary.weight_ton)} ton`,
      helper: "realisasi berat",
      icon: FaScaleBalanced,
      color: "bg-amber-50 text-amber-700",
    },
    {
      label: "Transport selesai",
      value: formatNumberID(summary.completed_transport_count),
      helper:
        summary.transport_count > 0
          ? `${Math.round((summary.completed_transport_count / summary.transport_count) * 100)}% completion`
          : "0% completion",
      icon: FaCircleCheck,
      color: "bg-emerald-50 text-emerald-700",
    },
    {
      label: "Laba bersih",
      value: formatCurrencyIDR(summary.net_profit),
      helper: `${formatCurrencyIDR(summary.total_income)} pendapatan`,
      icon: FaMoneyBillTrendUp,
      color:
        summary.net_profit < 0
          ? "bg-red-50 text-red-700"
          : "bg-emerald-50 text-emerald-700",
    },
  ];

  return (
    <section className="grid gap-3 sm:grid-cols-2 xl:grid-cols-3 2xl:grid-cols-6">
      {cards.map((card) => {
        const Icon = card.icon;
        return (
          <article
            key={card.label}
            className="rounded-2xl border border-slate-200 bg-white p-4 shadow-sm"
          >
            <div
              className={`mb-4 grid h-10 w-10 place-items-center rounded-xl ${card.color}`}
            >
              <Icon />
            </div>
            <p className="text-xs font-medium uppercase tracking-wide text-slate-400">
              {card.label}
            </p>
            <p className="mt-1 truncate text-xl font-semibold text-slate-900">
              {card.value}
            </p>
            <p className="mt-1 truncate text-xs text-slate-500">
              {card.helper}
            </p>
          </article>
        );
      })}
    </section>
  );
}

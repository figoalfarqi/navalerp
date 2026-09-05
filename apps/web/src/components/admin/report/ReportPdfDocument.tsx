import {
  Document,
  Page,
  Text,
  View,
  StyleSheet,
} from "@react-pdf/renderer";
import type {
  AdminDashboardSummary,
  AdminReportRow,
} from "@/types/adminDashboard";
import {
  formatCurrencyIDR,
  formatNumberID,
} from "@/utils/currencyFormater";

const styles = StyleSheet.create({
  page: {
    padding: 24,
    paddingBottom: 32,
    fontSize: 8,
    fontFamily: "Helvetica",
    color: "#1e293b",
  },
  header: {
    marginBottom: 10,
    borderBottomWidth: 1,
    borderBottomColor: "#cbd5e1",
    paddingBottom: 8,
  },
  title: {
    fontSize: 15,
    fontFamily: "Helvetica-Bold",
    color: "#0f172a",
  },
  meta: {
    fontSize: 8.5,
    color: "#475569",
    marginTop: 3,
  },
  summaryBar: {
    fontSize: 8,
    color: "#334155",
    marginTop: 4,
    backgroundColor: "#f8fafc",
    paddingVertical: 4,
    paddingHorizontal: 8,
    borderRadius: 4,
    borderWidth: 0.5,
    borderColor: "#e2e8f0",
  },
  table: {
    width: "100%",
    marginTop: 6,
  },
  tableHeader: {
    flexDirection: "row",
    backgroundColor: "#155eaa",
    color: "#ffffff",
    fontFamily: "Helvetica-Bold",
    fontSize: 7.5,
    paddingVertical: 5,
    paddingHorizontal: 4,
    borderRadius: 2,
  },
  tableRow: {
    flexDirection: "row",
    borderBottomWidth: 0.5,
    borderBottomColor: "#e2e8f0",
    paddingVertical: 4,
    paddingHorizontal: 4,
    fontSize: 7,
  },
  rowEven: {
    backgroundColor: "#ffffff",
  },
  rowOdd: {
    backgroundColor: "#f5f8fc",
  },
  colProject: { width: "18%" },
  colPeriod: { width: "10%" },
  colNum: { width: "8%", textAlign: "right" },
  colVolume: { width: "9%", textAlign: "right" },
  colWeight: { width: "9%", textAlign: "right" },
  colCurrency: { width: "12%", textAlign: "right" },
  footer: {
    position: "absolute",
    bottom: 14,
    left: 24,
    right: 24,
    flexDirection: "row",
    justifyContent: "space-between",
    fontSize: 7,
    color: "#94a3b8",
    borderTopWidth: 0.5,
    borderTopColor: "#e2e8f0",
    paddingTop: 4,
  },
});

export interface ReportPdfDocumentProps {
  projectLabel: string;
  periodLabel: string;
  date: string;
  summary: AdminDashboardSummary;
  rows: AdminReportRow[];
}

export default function ReportPdfDocument({
  projectLabel,
  periodLabel,
  date,
  summary,
  rows,
}: ReportPdfDocumentProps) {
  return (
    <Document title={`Laporan-${periodLabel}-${date}`}>
      <Page size="A4" orientation="landscape" style={styles.page}>
        <View style={styles.header}>
          <Text style={styles.title}>Laporan Operasional & Keuangan NavalERP</Text>
          <Text style={styles.meta}>
            {`${projectLabel} | ${periodLabel} | Acuan ${date}`}
          </Text>
          <Text style={styles.summaryBar}>
            {`Transport ${formatNumberID(summary.transport_count)} (${formatNumberID(summary.completed_transport_count)} Selesai)  |  Volume ${formatNumberID(summary.volume_cubic)} m3  |  Berat ${formatNumberID(summary.weight_ton)} ton  |  Pendapatan ${formatCurrencyIDR(summary.total_income)}  |  Pengeluaran ${formatCurrencyIDR(summary.total_expense)}  |  Laba ${formatCurrencyIDR(summary.net_profit)}`}
          </Text>
        </View>

        <View style={styles.table}>
          <View style={styles.tableHeader} fixed>
            <Text style={styles.colProject}>Project</Text>
            <Text style={styles.colPeriod}>Periode</Text>
            <Text style={styles.colNum}>Transport</Text>
            <Text style={styles.colNum}>Selesai</Text>
            <Text style={styles.colVolume}>Volume (m3)</Text>
            <Text style={styles.colWeight}>Berat (ton)</Text>
            <Text style={styles.colCurrency}>Pendapatan</Text>
            <Text style={styles.colCurrency}>Pengeluaran</Text>
            <Text style={styles.colCurrency}>Laba Bersih</Text>
          </View>

          {rows.map((row, index) => (
            <View
              key={`${row.project_id ?? index}-${row.period_label}-${index}`}
              style={[styles.tableRow, index % 2 === 0 ? styles.rowEven : styles.rowOdd]}
              wrap={false}
            >
              <Text style={styles.colProject}>
                {row.project_code ? `${row.project_code} - ${row.project_name}` : row.project_name}
              </Text>
              <Text style={styles.colPeriod}>{row.period_label}</Text>
              <Text style={styles.colNum}>{formatNumberID(row.transport_count)}</Text>
              <Text style={styles.colNum}>{formatNumberID(row.completed_transport_count)}</Text>
              <Text style={styles.colVolume}>{formatNumberID(row.volume_cubic)}</Text>
              <Text style={styles.colWeight}>{formatNumberID(row.weight_ton)}</Text>
              <Text style={styles.colCurrency}>{formatCurrencyIDR(row.total_income)}</Text>
              <Text style={styles.colCurrency}>{formatCurrencyIDR(row.total_expense)}</Text>
              <Text style={styles.colCurrency}>{formatCurrencyIDR(row.net_profit)}</Text>
            </View>
          ))}
        </View>

        <View style={styles.footer} fixed>
          <Text>NavalERP - Dokumen ini digenerate secara otomatis</Text>
          <Text
            render={({ pageNumber, totalPages }) =>
              `Halaman ${pageNumber} dari ${totalPages}`
            }
          />
        </View>
      </Page>
    </Document>
  );
}


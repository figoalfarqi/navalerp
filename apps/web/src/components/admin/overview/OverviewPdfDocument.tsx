import {
  Document,
  Page,
  Text,
  View,
  StyleSheet,
} from "@react-pdf/renderer";
import { formatNumberID } from "@/utils/currencyFormater";

const styles = StyleSheet.create({
  page: {
    paddingTop: 28,
    paddingBottom: 36,
    paddingHorizontal: 28,
    fontSize: 8.5,
    fontFamily: "Helvetica",
    color: "#1e293b",
  },
  // Official Kop Surat
  kopSurat: {
    borderBottomWidth: 2,
    borderBottomColor: "#0f2744",
    paddingBottom: 8,
    marginBottom: 10,
    alignItems: "center",
  },
  kopInstansi: {
    fontSize: 9,
    fontFamily: "Helvetica-Bold",
    color: "#0f2744",
    letterSpacing: 1.2,
    textTransform: "uppercase",
  },
  kopMarkas: {
    fontSize: 12,
    fontFamily: "Helvetica-Bold",
    color: "#0f2744",
    marginTop: 2,
    letterSpacing: 0.5,
  },
  kopSub: {
    fontSize: 7.5,
    color: "#64748b",
    marginTop: 2,
  },
  kopDividerThin: {
    width: "100%",
    height: 0.5,
    backgroundColor: "#0891b2",
    marginTop: 4,
  },

  // Document Title & Meta
  titleContainer: {
    marginTop: 4,
    marginBottom: 12,
    flexDirection: "row",
    justifyContent: "space-between",
    alignItems: "flex-end",
  },
  docTitle: {
    fontSize: 13,
    fontFamily: "Helvetica-Bold",
    color: "#0f172a",
    textTransform: "uppercase",
  },
  docSubtitle: {
    fontSize: 8,
    color: "#475569",
    marginTop: 2,
  },
  metaBox: {
    backgroundColor: "#f1f5f9",
    borderWidth: 0.5,
    borderColor: "#cbd5e1",
    borderRadius: 3,
    paddingVertical: 3,
    paddingHorizontal: 6,
    alignItems: "flex-end",
  },
  metaText: {
    fontSize: 7,
    color: "#334155",
  },
  badgeRestricted: {
    fontSize: 7,
    fontFamily: "Helvetica-Bold",
    color: "#b91c1c",
    marginTop: 1,
  },

  // Section Header
  sectionHeader: {
    backgroundColor: "#0f2744",
    color: "#ffffff",
    fontFamily: "Helvetica-Bold",
    fontSize: 8.5,
    paddingVertical: 4,
    paddingHorizontal: 8,
    borderRadius: 2,
    marginTop: 8,
    marginBottom: 6,
    textTransform: "uppercase",
    letterSpacing: 0.5,
  },

  // KPI Grid / Cards
  kpiContainer: {
    flexDirection: "row",
    flexWrap: "wrap",
    justifyContent: "space-between",
    marginBottom: 10,
  },
  kpiCard: {
    width: "32%",
    backgroundColor: "#f8fafc",
    borderWidth: 0.8,
    borderColor: "#cbd5e1",
    borderRadius: 4,
    padding: 7,
    marginBottom: 6,
  },
  kpiLabel: {
    fontSize: 7.5,
    fontFamily: "Helvetica-Bold",
    color: "#475569",
    textTransform: "uppercase",
  },
  kpiValue: {
    fontSize: 15,
    fontFamily: "Helvetica-Bold",
    color: "#0f2744",
    marginTop: 2,
    marginBottom: 1,
  },
  kpiDesc: {
    fontSize: 6.5,
    color: "#64748b",
  },

  // Table
  table: {
    width: "100%",
    marginTop: 2,
    marginBottom: 10,
  },
  tableHeader: {
    flexDirection: "row",
    backgroundColor: "#1e3a5f",
    color: "#ffffff",
    fontFamily: "Helvetica-Bold",
    fontSize: 7.5,
    paddingVertical: 5,
    paddingHorizontal: 6,
    borderRadius: 2,
  },
  tableRow: {
    flexDirection: "row",
    borderBottomWidth: 0.5,
    borderBottomColor: "#e2e8f0",
    paddingVertical: 4,
    paddingHorizontal: 6,
    fontSize: 7.5,
    alignItems: "center",
  },
  rowEven: {
    backgroundColor: "#ffffff",
  },
  rowOdd: {
    backgroundColor: "#f8fafc",
  },
  colCode: { width: "12%", fontFamily: "Helvetica-Bold", color: "#0f2744" },
  colName: { width: "32%", fontFamily: "Helvetica-Bold", color: "#1e293b" },
  colDesc: { width: "42%", color: "#475569", fontSize: 7 },
  colBadge: { width: "14%", textAlign: "right", color: "#0891b2", fontFamily: "Helvetica-Bold" },

  // Summary Information Banner
  infoBox: {
    backgroundColor: "#e0f2fe",
    borderWidth: 0.8,
    borderColor: "#7dd3fc",
    borderRadius: 4,
    padding: 7,
    marginBottom: 12,
  },
  infoTitle: {
    fontSize: 8,
    fontFamily: "Helvetica-Bold",
    color: "#0369a1",
    marginBottom: 2,
  },
  infoText: {
    fontSize: 7,
    color: "#0c4a6e",
    lineHeight: 1.3,
  },

  // Signatures
  signatureSection: {
    marginTop: 14,
    flexDirection: "row",
    justifyContent: "space-between",
    paddingHorizontal: 20,
  },
  signatureBox: {
    alignItems: "center",
    width: "40%",
  },
  signaturePlaceDate: {
    fontSize: 7.5,
    color: "#334155",
    marginBottom: 2,
  },
  signatureTitle: {
    fontSize: 7.5,
    fontFamily: "Helvetica-Bold",
    color: "#0f172a",
    marginBottom: 38,
    textAlign: "center",
  },
  signatureName: {
    fontSize: 8,
    fontFamily: "Helvetica-Bold",
    color: "#0f172a",
    textDecoration: "underline",
  },
  signatureRole: {
    fontSize: 7,
    color: "#475569",
    marginTop: 1,
  },

  // Footer
  footer: {
    position: "absolute",
    bottom: 16,
    left: 28,
    right: 28,
    flexDirection: "row",
    justifyContent: "space-between",
    fontSize: 6.5,
    color: "#94a3b8",
    borderTopWidth: 0.5,
    borderTopColor: "#e2e8f0",
    paddingTop: 4,
  },
});

export interface OverviewPdfDocumentProps {
  metrics: {
    ships: number;
    missions: number;
    workOrders: number;
    materials: number;
    personnel: number;
    readinessReports: number;
  };
  modules: Array<{
    name: string;
    code: string;
    desc: string;
    badge: string;
  }>;
  userName: string;
  userRole: string;
  date: string;
  time: string;
}

export default function OverviewPdfDocument({
  metrics,
  modules,
  userName,
  userRole,
  date,
  time,
}: OverviewPdfDocumentProps) {
  return (
    <Document title={`Laporan-Overview-NavalERP-${date}`}>
      <Page size="A4" orientation="portrait" style={styles.page}>
        {/* Kop Surat Resmi Militer */}
        <View style={styles.kopSurat}>
          <Text style={styles.kopInstansi}>Tentara Nasional Indonesia Angkatan Laut</Text>
          <Text style={styles.kopMarkas}>Pusat Komando Sistem Operasional Naval ERP</Text>
          <Text style={styles.kopSub}>
            Sistem Informasi Manajemen Terpadu Kesiapan Tempur, Pemeliharaan Alutsista & Rantai Pasok Maritim
          </Text>
          <View style={styles.kopDividerThin} />
        </View>

        {/* Title & Metadata */}
        <View style={styles.titleContainer}>
          <View>
            <Text style={styles.docTitle}>Ringkasan Eksekutif Operasional</Text>
            <Text style={styles.docSubtitle}>
              Laporan Status Komando Alutsista, Personel & Kesiapan Matra Laut
            </Text>
          </View>
          <View style={styles.metaBox}>
            <Text style={styles.metaText}>{`Tanggal: ${date} ${time}`}</Text>
            <Text style={styles.metaText}>{`Operator: ${userName} (${userRole})`}</Text>
            <Text style={styles.badgeRestricted}>KLASIFIKASI: TERBATAS</Text>
          </View>
        </View>

        {/* Indikator Utama Alutsista & Operasi */}
        <View style={styles.sectionHeader}>
          <Text>1. Indikator Kunci Operasional & Kesiapan Alutsista</Text>
        </View>

        <View style={styles.kpiContainer}>
          <View style={styles.kpiCard}>
            <Text style={styles.kpiLabel}>Armada Kapal KRI</Text>
            <Text style={styles.kpiValue}>{formatNumberID(metrics.ships)}</Text>
            <Text style={styles.kpiDesc}>Kapal perang KRI terdaftar</Text>
          </View>
          <View style={styles.kpiCard}>
            <Text style={styles.kpiLabel}>Operasi & Misi Pelayaran</Text>
            <Text style={styles.kpiValue}>{formatNumberID(metrics.missions)}</Text>
            <Text style={styles.kpiDesc}>Misi pelayaran terintegrasi</Text>
          </View>
          <View style={styles.kpiCard}>
            <Text style={styles.kpiLabel}>Pemeliharaan (MRO)</Text>
            <Text style={styles.kpiValue}>{formatNumberID(metrics.workOrders)}</Text>
            <Text style={styles.kpiDesc}>Work order perbaikan & docking</Text>
          </View>
          <View style={styles.kpiCard}>
            <Text style={styles.kpiLabel}>Logistik & Suku Cadang</Text>
            <Text style={styles.kpiValue}>{formatNumberID(metrics.materials)}</Text>
            <Text style={styles.kpiDesc}>Katalog material pertahanan</Text>
          </View>
          <View style={styles.kpiCard}>
            <Text style={styles.kpiLabel}>Personel Prajurit</Text>
            <Text style={styles.kpiValue}>{formatNumberID(metrics.personnel)}</Text>
            <Text style={styles.kpiDesc}>Prajurit & perwira TNI AL</Text>
          </View>
          <View style={styles.kpiCard}>
            <Text style={styles.kpiLabel}>Kesiapan Tempur</Text>
            <Text style={styles.kpiValue}>{formatNumberID(metrics.readinessReports)}</Text>
            <Text style={styles.kpiDesc}>Laporan kesiapan alutsista</Text>
          </View>
        </View>

        {/* Ringkasan 10 Modul Terintegrasi */}
        <View style={styles.sectionHeader}>
          <Text>2. Status 10 Modul Subsistem Naval ERP (72 Tabel)</Text>
        </View>

        <View style={styles.table}>
          <View style={styles.tableHeader} fixed>
            <Text style={styles.colCode}>Kode</Text>
            <Text style={styles.colName}>Nama Modul Subsistem</Text>
            <Text style={styles.colDesc}>Cakupan Operasional</Text>
            <Text style={styles.colBadge}>Jumlah Tabel</Text>
          </View>

          {modules.map((mod, index) => (
            <View
              key={mod.code}
              style={[styles.tableRow, index % 2 === 0 ? styles.rowEven : styles.rowOdd]}
              wrap={false}
            >
              <Text style={styles.colCode}>{mod.code}</Text>
              <Text style={styles.colName}>{mod.name}</Text>
              <Text style={styles.colDesc}>{mod.desc}</Text>
              <Text style={styles.colBadge}>{mod.badge}</Text>
            </View>
          ))}
        </View>

        {/* Catatan Status Sistem */}
        <View style={styles.infoBox}>
          <Text style={styles.infoTitle}>CATATAN KELAIKAN OPERASIONAL SISTEM:</Text>
          <Text style={styles.infoText}>
            Seluruh data operasional di atas dihasilkan secara sinkron dari database terintegrasi Naval ERP.
            Database transactions aktif memastikan konsistensi relasi data induk dan rincian item logistik,
            dokumen kontrak, serta rekam kesiapan tempur unsur KRI di seluruh jajaran Komando Armada.
          </Text>
        </View>

        {/* Lembar Tanda Tangan */}
        <View style={styles.signatureSection} wrap={false}>
          <View style={styles.signatureBox}>
            <Text style={styles.signaturePlaceDate}>Jakarta, {date}</Text>
            <Text style={styles.signatureTitle}>Perwira Operator Sistem</Text>
            <Text style={styles.signatureName}>{userName}</Text>
            <Text style={styles.signatureRole}>{userRole}</Text>
          </View>
          <View style={styles.signatureBox}>
            <Text style={styles.signaturePlaceDate}>Mengetahui,</Text>
            <Text style={styles.signatureTitle}>Komandan / Pimpinan Operasi</Text>
            <Text style={styles.signatureName}>Laksamana Muda TNI</Text>
            <Text style={styles.signatureRole}>Panglima / Kepala Staf Terkait</Text>
          </View>
        </View>

        {/* Footer */}
        <View style={styles.footer} fixed>
          <Text>NAVAL ERP - PUSAT KOMANDO OPERASIONAL MARITIM TNI AL (DOKUMEN RESMI TERBATAS)</Text>
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


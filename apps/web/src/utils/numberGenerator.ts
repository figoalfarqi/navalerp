/* eslint-disable @typescript-eslint/no-explicit-any */

export interface AutoNumberConfig {
  prefix: string;
  entity: string;
  table?: string;
  column?: string;
}

export const AUTO_NUMBER_CONFIGS: Record<string, AutoNumberConfig> = {
  requisition_number: { prefix: "PR", entity: "requisition", table: "proc_requisitions", column: "requisition_number" },
  po_number: { prefix: "PO", entity: "purchase_order", table: "proc_purchase_orders", column: "po_number" },
  receipt_number: { prefix: "BAPHP", entity: "goods_receipt", table: "proc_goods_receipts", column: "receipt_number" },
  transfer_number: { prefix: "TRF", entity: "stock_transfer", table: "inv_stock_transfers", column: "transfer_number" },
  adjustment_number: { prefix: "ADJ", entity: "stock_adjustment", table: "inv_stock_adjustments", column: "adjustment_number" },
  work_order_number: { prefix: "WO", entity: "work_order", table: "mro_work_orders", column: "work_order_number" },
  entry_number: { prefix: "JRN", entity: "journal_entry", table: "fin_journal_entries", column: "entry_number" },
  manifest_number: { prefix: "MNF", entity: "shipment", table: "log_shipments", column: "manifest_number" },
  tender_number: { prefix: "TND", entity: "tender", table: "proc_tenders", column: "tender_number" },
  contract_number: { prefix: "KTR", entity: "contract", table: "proc_contracts", column: "contract_number" },
  invoice_number: { prefix: "INV", entity: "invoice", table: "fin_invoices", column: "invoice_number" },
  payment_reference_no: { prefix: "SP2D", entity: "payment", table: "fin_payments", column: "payment_reference_no" },
  commitment_number: { prefix: "KOM", entity: "budget_commitment", table: "fin_budget_commitments", column: "commitment_number" },
  report_number: { prefix: "REP", entity: "failure_report", table: "mro_failure_reports", column: "report_number" },
  mission_code: { prefix: "MSN", entity: "mission", table: "ops_missions", column: "mission_code" },
  document_number: { prefix: "DOC", entity: "document", table: "doc_documents", column: "document_number" },
  alert_code: { prefix: "ALT", entity: "readiness_alert", table: "ops_readiness_alerts", column: "alert_code" },
};

/**
 * Generates an immediate neat fallback document number formatted as PREFIX-YYYY-0001
 */
export function generateFallbackNumber(prefix: string): string {
  const year = new Date().getFullYear();
  return `${prefix}-${year}-0001`;
}

/**
 * Resolves the configuration for a field name or explicit entity
 */
export function getAutoNumberConfig(fieldName: string, explicitPrefix?: string, explicitEntity?: string): AutoNumberConfig | null {
  if (AUTO_NUMBER_CONFIGS[fieldName]) {
    const base = AUTO_NUMBER_CONFIGS[fieldName];
    return {
      prefix: explicitPrefix || base.prefix,
      entity: explicitEntity || base.entity,
      table: base.table,
      column: base.column,
    };
  }

  if (explicitPrefix) {
    return {
      prefix: explicitPrefix,
      entity: explicitEntity || fieldName.replace(/_number$|_code$/, ""),
    };
  }

  return null;
}

/**
 * Fetches the next neat sequential number from backend API with fallback
 */
export async function fetchAutoNumber(
  fieldName: string,
  explicitPrefix?: string,
  explicitEntity?: string,
  getAPI?: (url: string, options?: any) => Promise<any>
): Promise<string> {
  const cfg = getAutoNumberConfig(fieldName, explicitPrefix, explicitEntity);
  const prefix = cfg?.prefix || "DOC";

  if (!getAPI) {
    return generateFallbackNumber(prefix);
  }

  try {
    const params = new URLSearchParams();
    if (cfg?.entity) params.set("entity", cfg.entity);
    if (cfg?.prefix) params.set("prefix", cfg.prefix);
    if (cfg?.table) params.set("table", cfg.table);
    if (cfg?.column) params.set("column", cfg.column);

    const res: any = await getAPI(`/admin/generate_number?${params.toString()}`, { authToken: "admin" });
    const generated = res?.data?.number || res?.data?.data?.number;
    if (generated && typeof generated === "string") {
      return generated;
    }
  } catch (err) {
    console.warn(`[AutoNumber] Could not fetch auto-number for ${fieldName}, using fallback`, err);
  }

  return generateFallbackNumber(prefix);
}

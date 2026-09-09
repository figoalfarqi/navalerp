/* eslint-disable @typescript-eslint/no-explicit-any */
"use client";

import React, { useState, useEffect, useMemo } from "react";
import { useFetchAPI } from "@/hooks/useFetchAPI";
import { formatCurrencyIDR, formatNumberID } from "@/utils/currencyFormater";
import { PlusIcon, EditIcon, TrashIcon, CloseIcon, CheckIcon, BoxesIcon } from "@/components/icons";
import Button from "../form/Button";
import SelectField, { SelectOption } from "../form/SelectField";
import TextField from "../form/TextField";

export interface DetailItemColumn {
  key: string;
  label: string;
  type: "text" | "number" | "currency" | "select" | "readonly";
  required?: boolean;
  placeholder?: string;
  width?: string;
  options?:
    | Array<{ label: string; value: any }>
    | {
        url: string;
        labelKey: string;
        valueKey: string;
        extraLabelKey?: string;
        subLabelKey?: string;
        priceKey?: string;
      };
  render?: (value: any, item: any, index: number) => React.ReactNode;
  defaultValue?: any;
}

export interface DetailItemsConfig {
  tableName: string; // e.g. "items" or "lines"
  title: string; // e.g. "Rincian Item Material"
  itemName?: string; // e.g. "Item Material"
  columns: DetailItemColumn[];
  qtyKey?: string; // default "quantity"
  priceKey?: string; // default "unit_price"
  totalKey?: string; // default "total_price"
  syncHeaderTotalKey?: string; // e.g. "total_amount"
}

export interface DetailItemsTableProps {
  config: DetailItemsConfig;
  items: any[];
  onChange: (newItems: any[]) => void;
  mode: "add" | "edit" | "view" | "copy";
  onSyncHeaderTotal?: (total: number) => void;
}

export default function DetailItemsTable({
  config,
  items = [],
  onChange,
  mode,
  onSyncHeaderTotal,
}: DetailItemsTableProps) {
  const { getAPI } = useFetchAPI();
  const isView = mode === "view";
  const qtyKey = config.qtyKey || "quantity";
  const priceKey = config.priceKey || "unit_price";
  const totalKey = config.totalKey || "total_price";
  const itemName = config.itemName || "Item";

  // Sanitize title to avoid showing raw database table names (e.g. inv_stock_transfer_items)
  const displayTitle = useMemo(() => {
    return (config.title || "").replace(/\s*\([a-z0-9_]+\)\s*/gi, " ").trim();
  }, [config.title]);

  // State for Modal
  const [isModalOpen, setIsModalOpen] = useState(false);
  const [editingIndex, setEditingIndex] = useState<number | null>(null);
  const [currentItem, setCurrentItem] = useState<Record<string, any>>({});
  const [errors, setErrors] = useState<Record<string, string>>({});

  // Dynamic Options Cache: column.key -> Array<{ label, value, row }>
  const [optionsMap, setOptionsMap] = useState<Record<string, Array<{ label: string; value: any; row?: any }>>>({});
  const [loadingOptions, setLoadingOptions] = useState<Record<string, boolean>>({});

  // Fetch dynamic options for select columns
  useEffect(() => {
    config.columns.forEach(async (col) => {
      if (col.type === "select" && col.options && !Array.isArray(col.options)) {
        const dyn = col.options;
        setLoadingOptions((prev) => ({ ...prev, [col.key]: true }));
        try {
          const res = await getAPI<any>(dyn.url, { authToken: "admin" });
          const rawItems = res?.data?.items || res?.data?.data?.items || res?.data || [];
          if (Array.isArray(rawItems)) {
            const mapped = rawItems.map((r: any) => {
              const labelMain = r[dyn.labelKey] ?? "";
              const labelExtra = dyn.extraLabelKey && r[dyn.extraLabelKey] ? ` (${r[dyn.extraLabelKey]})` : "";
              const labelSub = dyn.subLabelKey && r[dyn.subLabelKey] ? ` - ${r[dyn.subLabelKey]}` : "";
              return {
                label: `${labelMain}${labelExtra}${labelSub}`,
                value: r[dyn.valueKey],
                row: r,
              };
            });
            setOptionsMap((prev) => ({ ...prev, [col.key]: mapped }));
          }
        } catch (err) {
          console.error(`Failed to load options for ${col.key}:`, err);
        } finally {
          setLoadingOptions((prev) => ({ ...prev, [col.key]: false }));
        }
      }
    });
  }, [config, getAPI]);

  // Sync header total when items change
  useEffect(() => {
    if (config.syncHeaderTotalKey && onSyncHeaderTotal) {
      const total = items.reduce((acc, it) => {
        const lineTotal = Number(it[totalKey]) || (Number(it[qtyKey]) || 0) * (Number(it[priceKey]) || 0);
        return acc + (isNaN(lineTotal) ? 0 : lineTotal);
      }, 0);
      onSyncHeaderTotal(total);
    }
  }, [items, config.syncHeaderTotalKey, totalKey, qtyKey, priceKey, onSyncHeaderTotal]);

  // Calculate Summary Totals
  const summaryTotals = useMemo(() => {
    let totalQty = 0;
    let totalAmount = 0;
    items.forEach((it) => {
      const q = Number(it[qtyKey]) || 0;
      const p = Number(it[priceKey]) || 0;
      const lineTotal = Number(it[totalKey]) || q * p;
      totalQty += q;
      totalAmount += lineTotal;
    });
    return { totalQty, totalAmount, count: items.length };
  }, [items, qtyKey, priceKey, totalKey]);

  // Open modal for adding
  const handleOpenAdd = () => {
    const initial: Record<string, any> = {};
    config.columns.forEach((col) => {
      initial[col.key] = col.defaultValue !== undefined ? col.defaultValue : col.type === "number" || col.type === "currency" ? 0 : "";
    });
    setCurrentItem(initial);
    setEditingIndex(null);
    setErrors({});
    setIsModalOpen(true);
  };

  // Open modal for editing
  const handleOpenEdit = (index: number) => {
    setCurrentItem({ ...items[index] });
    setEditingIndex(index);
    setErrors({});
    setIsModalOpen(true);
  };

  // Delete row
  const handleDelete = (index: number) => {
    const updated = items.filter((_, i) => i !== index);
    onChange(updated);
  };

  // Handle input change in modal
  const handleFieldChange = (key: string, value: any, col: DetailItemColumn, selectedRow?: any) => {
    const updated = { ...currentItem, [key]: value };

    // Auto populate details when select changes (e.g. material selection)
    if (col.type === "select" && col.options && !Array.isArray(col.options)) {
      const dyn = col.options;
      const opts = optionsMap[key] || [];
      const foundRow = selectedRow || opts.find((o) => String(o.value) === String(value))?.row;
      if (foundRow) {
        if (dyn.labelKey) {
          updated[dyn.labelKey] = foundRow[dyn.labelKey];
        }
        if (dyn.extraLabelKey) {
          updated[dyn.extraLabelKey] = foundRow[dyn.extraLabelKey];
        }
        // Auto fill unit price from standard cost if current price is 0
        if (dyn.priceKey && (!updated[priceKey] || Number(updated[priceKey]) === 0)) {
          updated[priceKey] = Number(foundRow[dyn.priceKey]) || 0;
        }
      }
    }

    // Auto calculate line total if qty or price changes
    if (key === qtyKey || key === priceKey) {
      const q = key === qtyKey ? Number(value) || 0 : Number(updated[qtyKey]) || 0;
      const p = key === priceKey ? Number(value) || 0 : Number(updated[priceKey]) || 0;
      updated[totalKey] = q * p;
    }

    setCurrentItem(updated);
    if (errors[key]) {
      setErrors((prev) => {
        const cp = { ...prev };
        delete cp[key];
        return cp;
      });
    }
  };

  // Validate and save modal item
  const handleSaveModal = (e: React.FormEvent) => {
    e.preventDefault();
    const newErrors: Record<string, string> = {};

    config.columns.forEach((col) => {
      const val = currentItem[col.key];
      if (col.required && (val === undefined || val === null || val === "")) {
        newErrors[col.key] = `${col.label} wajib diisi`;
      }
      if (col.type === "number" || col.type === "currency") {
        const num = Number(val);
        if (col.required && (isNaN(num) || num <= 0)) {
          newErrors[col.key] = `${col.label} harus lebih dari 0`;
        }
      }
    });

    if (Object.keys(newErrors).length > 0) {
      setErrors(newErrors);
      return;
    }

    // Calculate final line total
    const q = Number(currentItem[qtyKey]) || 0;
    const p = Number(currentItem[priceKey]) || 0;
    const itemToSave = {
      ...currentItem,
      [qtyKey]: q,
      [priceKey]: p,
      [totalKey]: currentItem[totalKey] !== undefined ? Number(currentItem[totalKey]) || q * p : q * p,
    };

    let updated: any[];
    if (editingIndex !== null) {
      updated = [...items];
      updated[editingIndex] = itemToSave;
    } else {
      updated = [...items, itemToSave];
    }

    onChange(updated);
    setIsModalOpen(false);
  };

  // Helper to format cell display
  const renderCellContent = (col: DetailItemColumn, row: any, idx: number) => {
    if (col.render) {
      return col.render(row[col.key], row, idx);
    }
    const val = row[col.key];

    if (col.type === "currency") {
      return (
        <span className="font-semibold text-slate-800">
          {formatCurrencyIDR(val ?? 0)}
        </span>
      );
    }

    if (col.type === "number") {
      return (
        <span className="font-medium text-slate-700">
          {formatNumberID(Number(val) || 0)}
        </span>
      );
    }

    if (col.type === "select") {
      const opts = Array.isArray(col.options)
        ? col.options
        : optionsMap[col.key] || [];
      const found = opts.find((o) => String(o.value) === String(val));
      if (found) return <span className="font-medium text-slate-900">{found.label}</span>;

      if (col.key === "material_id") {
        const matName = row.material_name || row.material_code;
        if (matName) {
          return (
            <div>
              <div className="font-semibold text-slate-900">{row.material_name || val}</div>
              {row.material_code && (
                <span className="text-xs text-slate-400 font-mono">Kode: {row.material_code}</span>
              )}
            </div>
          );
        }
      }
      return <span>{val || "-"}</span>;
    }

    return <span>{val !== undefined && val !== null && val !== "" ? String(val) : "-"}</span>;
  };

  const hasPricing = config.columns.some((c) => c.key === priceKey);

  return (
    <div className="col-span-1 md:col-span-2 mt-4 bg-white border border-slate-200 rounded-xl shadow-sm overflow-hidden">
      {/* Header Bar */}
      <div className="flex flex-col sm:flex-row sm:items-center justify-between px-4 sm:px-5 py-3 sm:py-3.5 bg-slate-50 border-b border-slate-200 gap-3">
        <div className="flex items-center gap-2.5">
          <div className="w-8 h-8 rounded-lg bg-blue-50 text-[#0a2540] flex items-center justify-center font-bold shrink-0">
            <BoxesIcon size={18} />
          </div>
          <div>
            <h3 className="text-base font-bold text-slate-900 leading-tight">
              {displayTitle}
            </h3>
            <p className="text-xs text-slate-500">
              Total {items.length} {itemName.toLowerCase()} terdaftar
            </p>
          </div>
        </div>

        {!isView && (
          <Button
            id="btn-add-detail-item"
            type="button"
            size="sm"
            variant="blue-solid"
            onClick={handleOpenAdd}
            className="flex items-center justify-center gap-1.5 shadow-sm w-full sm:w-auto cursor-pointer"
          >
            <PlusIcon size={16} />
            <span>Tambah {itemName}</span>
          </Button>
        )}
      </div>

      {/* Mobile Card List View for Detail Items */}
      <div className="block md:hidden p-3 space-y-3">
        {items.length === 0 ? (
          <div className="py-8 text-center text-slate-400">
            <div className="flex flex-col items-center justify-center">
              <BoxesIcon size={32} className="text-slate-300 mb-2" />
              <p className="font-medium text-slate-600 text-sm">
                Belum ada {itemName.toLowerCase()} yang ditambahkan
              </p>
              {!isView && (
                <p className="text-xs text-slate-400 mt-1">
                  Klik tombol &quot;Tambah {itemName}&quot; di atas untuk mulai memasukkan data
                </p>
              )}
            </div>
          </div>
        ) : (
          items.map((row, idx) => {
            const q = Number(row[qtyKey]) || 0;
            const p = Number(row[priceKey]) || 0;
            const lineTotal = Number(row[totalKey]) || q * p;
            const primaryCol = config.columns[0];
            const otherCols = config.columns.slice(1);

            return (
              <div
                key={idx}
                className="rounded-xl border border-slate-200 bg-white p-3.5 shadow-xs space-y-2.5"
              >
                {/* Header: # + Primary label */}
                <div className="flex items-start justify-between gap-2">
                  <div className="flex items-start gap-2 min-w-0">
                    <span className="inline-flex items-center justify-center h-5 min-w-5 px-1 rounded bg-blue-50 text-[#004f7f] text-xs font-bold shrink-0 mt-0.5">
                      #{idx + 1}
                    </span>
                    <div className="min-w-0">
                      <span className="block text-sm font-bold text-slate-900 leading-snug break-words">
                        {renderCellContent(primaryCol, row, idx)}
                      </span>
                      {primaryCol && (
                        <span className="block text-[10px] text-slate-400 leading-tight">
                          {primaryCol.label}
                        </span>
                      )}
                    </div>
                  </div>

                  {/* Subtotal badge if pricing */}
                  {hasPricing && (
                    <span className="shrink-0 text-xs font-extrabold text-[#004f7f] bg-blue-50 px-2 py-0.5 rounded">
                      {formatCurrencyIDR(lineTotal)}
                    </span>
                  )}
                </div>

                {/* Body fields */}
                {otherCols.length > 0 && (
                  <div className="grid grid-cols-2 gap-2 pt-2 border-t border-slate-100 text-xs">
                    {otherCols.map((col) => (
                      <div key={col.key} className="min-w-0">
                        <span className="block text-[11px] font-medium text-slate-400 truncate">
                          {col.label}
                        </span>
                        <span className="block text-xs font-semibold text-slate-800 break-words mt-0.5">
                          {renderCellContent(col, row, idx)}
                        </span>
                      </div>
                    ))}
                  </div>
                )}

                {/* Card Actions */}
                {!isView && (
                  <div className="flex items-center gap-2 pt-2 border-t border-slate-100">
                    <Button
                      id={`btn-mobile-edit-${idx}`}
                      type="button"
                      size="2xs"
                      variant="green-solid"
                      onClick={() => handleOpenEdit(idx)}
                      className="flex-1 justify-center !py-1.5 text-xs font-medium cursor-pointer shadow-xs"
                    >
                      <EditIcon size={13} className="mr-1" />
                      <span>Ubah</span>
                    </Button>
                    <Button
                      id={`btn-mobile-delete-${idx}`}
                      type="button"
                      size="2xs"
                      variant="red-solid"
                      onClick={() => handleDelete(idx)}
                      className="flex-1 justify-center !py-1.5 text-xs font-medium cursor-pointer shadow-xs"
                    >
                      <TrashIcon size={13} className="mr-1" />
                      <span>Hapus</span>
                    </Button>
                  </div>
                )}
              </div>
            );
          })
        )}

        {/* Mobile Summary Card */}
        {items.length > 0 && (
          <div className="rounded-xl border border-blue-100 bg-blue-50/50 p-3 flex flex-wrap items-center justify-between gap-2 text-xs font-semibold">
            <span className="text-slate-600">
              Total {items.length} item ({summaryTotals.totalQty} kuantitas)
            </span>
            {hasPricing && (
              <span className="text-sm font-extrabold text-[#004f7f]">
                {formatCurrencyIDR(summaryTotals.totalAmount)}
              </span>
            )}
          </div>
        )}
      </div>

      {/* Table Container (Desktop) */}
      <div className="hidden md:block overflow-x-auto">
        <table className="w-full text-left text-sm text-slate-600">
          <thead className="bg-slate-100/75 text-xs uppercase font-semibold text-slate-600 border-b border-slate-200">
            <tr>
              <th className="py-3 px-4 w-12 text-center">#</th>
              {config.columns.map((col) => (
                <th
                  key={col.key}
                  className={`py-3 px-4 ${
                    col.type === "number" || col.type === "currency"
                      ? "text-right"
                      : "text-left"
                  }`}
                  style={{ width: col.width }}
                >
                  {col.label}
                </th>
              ))}
              {hasPricing && (
                <th className="py-3 px-4 text-right w-36">Subtotal</th>
              )}
              {!isView && <th className="py-3 px-4 text-center w-28">Aksi</th>}
            </tr>
          </thead>
          <tbody className="divide-y divide-slate-100">
            {items.length === 0 ? (
              <tr>
                <td
                  colSpan={
                    config.columns.length + 1 + (hasPricing ? 1 : 0) + (!isView ? 1 : 0)
                  }
                  className="py-10 text-center text-slate-400"
                >
                  <div className="flex flex-col items-center justify-center">
                    <BoxesIcon size={36} className="text-slate-300 mb-2" />
                    <p className="font-medium text-slate-600">
                      Belum ada {itemName.toLowerCase()} yang ditambahkan
                    </p>
                    {!isView && (
                      <p className="text-xs text-slate-400 mt-1">
                        Klik tombol &quot;Tambah {itemName}&quot; di atas untuk mulai memasukkan data
                      </p>
                    )}
                  </div>
                </td>
              </tr>
            ) : (
              items.map((row, idx) => {
                const q = Number(row[qtyKey]) || 0;
                const p = Number(row[priceKey]) || 0;
                const lineTotal = Number(row[totalKey]) || q * p;

                return (
                  <tr
                    key={idx}
                    className="hover:bg-blue-50/40 transition-colors border-b border-slate-100"
                  >
                    <td className="py-3 px-4 text-center text-xs font-semibold text-slate-400">
                      {idx + 1}
                    </td>
                    {config.columns.map((col) => (
                      <td
                        key={col.key}
                        className={`py-3 px-4 ${
                          col.type === "number" || col.type === "currency"
                            ? "text-right"
                            : "text-left"
                        }`}
                      >
                        {renderCellContent(col, row, idx)}
                      </td>
                    ))}
                    {hasPricing && (
                      <td className="py-3 px-4 text-right font-semibold text-slate-900">
                        {formatCurrencyIDR(lineTotal)}
                      </td>
                    )}
                    {!isView && (
                      <td className="py-3 px-4 text-center">
                        <div className="flex items-center justify-center gap-1.5">
                          <Button
                            id={`btn-edit-item-${idx}`}
                            type="button"
                            size="2xs"
                            variant="blue-solid"
                            onClick={() => handleOpenEdit(idx)}
                            title="Ubah Item"
                            className="!p-1.5 !px-2 cursor-pointer shadow-xs"
                          >
                            <EditIcon size={14} />
                          </Button>
                          <Button
                            id={`btn-delete-item-${idx}`}
                            type="button"
                            size="2xs"
                            variant="red-solid"
                            onClick={() => handleDelete(idx)}
                            title="Hapus Item"
                            className="!p-1.5 !px-2 cursor-pointer shadow-xs"
                          >
                            <TrashIcon size={14} />
                          </Button>
                        </div>
                      </td>
                    )}
                  </tr>
                );
              })
            )}
          </tbody>

          {/* Table Footer with Summaries */}
          {items.length > 0 && (
            <tfoot className="bg-slate-50 font-bold border-t-2 border-slate-200 text-slate-800">
              <tr>
                <td colSpan={2} className="py-3 px-4 text-right uppercase text-xs tracking-wider text-slate-500">
                  Total Ringkasan:
                </td>
                {config.columns.slice(1).map((col) => {
                  if (col.key === qtyKey) {
                    return (
                      <td key={col.key} className="py-3 px-4 text-right text-slate-900 font-extrabold">
                        {formatNumberID(summaryTotals.totalQty)}
                      </td>
                    );
                  }
                  return <td key={col.key} className="py-3 px-4" />;
                })}
                {hasPricing && (
                  <td className="py-3 px-4 text-right text-[#004f7f] text-base font-extrabold">
                    {formatCurrencyIDR(summaryTotals.totalAmount)}
                  </td>
                )}
                {!isView && <td className="py-3 px-4" />}
              </tr>
            </tfoot>
          )}
        </table>
      </div>

      {/* Modal Dialog for Add / Edit Item */}
      {isModalOpen && (
        <div className="fixed inset-0 z-50 flex items-end sm:items-center justify-center p-0 sm:p-4 bg-slate-900/40 backdrop-blur-xs animate-fadeIn">
          <div className="bg-white rounded-t-2xl sm:rounded-2xl shadow-2xl border border-slate-200 max-w-xl w-full max-h-[92vh] sm:max-h-[85vh] overflow-hidden flex flex-col">
            {/* Modal Header */}
            <div className="flex items-center justify-between px-4 sm:px-6 py-3.5 sm:py-4 bg-slate-50 border-b border-slate-200 shrink-0">
              <div className="flex items-center gap-2">
                <div className="p-1.5 sm:p-2 bg-blue-100/70 text-[#004f7f] rounded-lg">
                  <BoxesIcon size={18} />
                </div>
                <h4 className="text-base sm:text-lg font-bold text-slate-900">
                  {editingIndex !== null ? `Ubah ${itemName}` : `Tambah ${itemName}`}
                </h4>
              </div>
              <Button
                id="btn-close-detail-modal"
                type="button"
                variant="gray-ghost"
                size="2xs"
                onClick={() => setIsModalOpen(false)}
                className="!p-1 text-slate-400 hover:text-slate-600 hover:bg-slate-100 rounded-lg cursor-pointer"
                title="Tutup"
              >
                <CloseIcon size={20} />
              </Button>
            </div>

            {/* Modal Body */}
            <form onSubmit={handleSaveModal} className="p-4 sm:p-6 space-y-4 overflow-y-auto flex-1 custom-scrollbar">
              {config.columns.map((col) => {
                const val = currentItem[col.key] ?? "";
                const err = errors[col.key];

                if (col.type === "select") {
                  const opts: SelectOption[] = Array.isArray(col.options)
                    ? col.options
                    : optionsMap[col.key] || [];
                  const isLoading = loadingOptions[col.key];

                  return (
                    <SelectField
                      key={col.key}
                      id={`detail-col-${col.key}`}
                      label={col.label}
                      value={val ?? ""}
                      options={opts}
                      disabled={isLoading}
                      required={col.required}
                      uncloseable={col.required}
                      className="w-full"
                      wrapperClassName="w-full"
                      placeholder={
                        isLoading
                          ? "Memuat opsi..."
                          : col.placeholder || `Pilih ${col.label}...`
                      }
                      error={err}
                      onChange={(selectedVal, _selectedLabel, selectedRow) => {
                        handleFieldChange(
                          col.key,
                          selectedVal ?? "",
                          col,
                          selectedRow
                        );
                      }}
                    />
                  );
                }

                if (col.type === "number" || col.type === "currency") {
                  return (
                    <TextField
                      key={col.key}
                      id={`detail-col-${col.key}`}
                      label={col.label}
                      type="number"
                      value={val === 0 ? "" : val}
                      placeholder={col.placeholder || "0"}
                      required={col.required}
                      error={err}
                      className="w-full"
                      wrapperClassName="w-full"
                      onChange={(newVal) =>
                        handleFieldChange(
                          col.key,
                          newVal === "" ? 0 : Number(newVal),
                          col
                        )
                      }
                    />
                  );
                }

                if (col.type === "readonly") {
                  return (
                    <TextField
                      key={col.key}
                      id={`detail-col-${col.key}`}
                      label={col.label}
                      type="text"
                      value={val ?? ""}
                      readOnly
                      disabled
                      required={col.required}
                      error={err}
                      className="w-full"
                      wrapperClassName="w-full"
                      onChange={() => {}}
                    />
                  );
                }

                // Default: text
                return (
                  <TextField
                    key={col.key}
                    id={`detail-col-${col.key}`}
                    label={col.label}
                    type="text"
                    value={val ?? ""}
                    placeholder={
                      col.placeholder || `Masukkan ${col.label.toLowerCase()}`
                    }
                    required={col.required}
                    error={err}
                    className="w-full"
                    wrapperClassName="w-full"
                    onChange={(newVal) =>
                      handleFieldChange(col.key, newVal, col)
                    }
                  />
                );
              })}

              {/* Subtotal Preview Card */}
              {hasPricing && (
                <div className="mt-4 p-3.5 bg-blue-50/70 border border-blue-100 rounded-xl flex items-center justify-between">
                  <span className="text-xs font-semibold uppercase tracking-wider text-slate-600">
                    Perkiraan Subtotal:
                  </span>
                  <span className="text-base font-extrabold text-[#0a2540]">
                    {formatCurrencyIDR(
                      (Number(currentItem[qtyKey]) || 0) * (Number(currentItem[priceKey]) || 0)
                    )}
                  </span>
                </div>
              )}

              {/* Modal Actions */}
              <div className="flex items-center justify-end gap-2.5 pt-4 border-t border-slate-100">
                <Button
                  id="btn-cancel-detail-modal"
                  type="button"
                  variant="gray-outline"
                  size="md"
                  onClick={() => setIsModalOpen(false)}
                  className="flex-1 sm:flex-none justify-center cursor-pointer"
                >
                  Batal
                </Button>
                <Button
                  id="btn-save-detail-modal"
                  type="submit"
                  variant="blue-solid"
                  size="md"
                  className="flex-1 sm:flex-none justify-center flex items-center gap-1.5 cursor-pointer"
                >
                  <CheckIcon size={16} />
                  <span>Simpan Item</span>
                </Button>
              </div>
            </form>
          </div>
        </div>
      )}
    </div>
  );
}

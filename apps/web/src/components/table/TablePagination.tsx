"use client";

import React from "react";
import { FaChevronLeft, FaChevronRight } from "@/components/icons";

interface TablePaginationProps {
  page: number;
  total: number;
  limit: number;
  onPageChange: (page: number) => void;
  isLoading?: boolean;
}

export default function TablePagination({
  page,
  total,
  limit,
  onPageChange,
  isLoading = false,
}: TablePaginationProps) {
  const totalPages = Math.max(1, Math.ceil(total / (limit || 10)));
  const startItem = total === 0 ? 0 : (page - 1) * limit + 1;
  const endItem = Math.min(page * limit, total);

  // Generate pagination buttons
  const getPageNumbers = () => {
    const pages: (number | string)[] = [];
    if (totalPages <= 7) {
      for (let i = 1; i <= totalPages; i++) {
        pages.push(i);
      }
    } else {
      pages.push(1);
      if (page > 3) {
        pages.push("...");
      }
      const start = Math.max(2, page - 1);
      const end = Math.min(totalPages - 1, page + 1);
      for (let i = start; i <= end; i++) {
        pages.push(i);
      }
      if (page < totalPages - 2) {
        pages.push("...");
      }
      pages.push(totalPages);
    }
    return pages;
  };

  const pageNumbers = getPageNumbers();

  return (
    <div className="flex flex-col sm:flex-row justify-between items-center gap-3 pt-3 pb-1 px-2 border-t border-slate-200 text-sm">
      <div className="text-xs text-slate-500 font-medium">
        Menampilkan{" "}
        <span className="font-semibold text-slate-700">
          {startItem} - {endItem}
        </span>{" "}
        dari{" "}
        <span className="font-semibold text-slate-700">{total}</span> data
      </div>

      <div className="flex items-center gap-1.5 sm:gap-2">
        {/* Previous button */}
        <button
          id="prev-page-button"
          type="button"
          disabled={page <= 1 || isLoading}
          onClick={() => onPageChange(page - 1)}
          className={`flex items-center gap-1.5 px-3 py-1.5 rounded-md text-xs font-semibold transition border select-none ${
            page <= 1 || isLoading
              ? "bg-slate-100 text-slate-400 border-slate-200 cursor-not-allowed opacity-60"
              : "bg-white text-slate-700 border-slate-300 hover:bg-slate-50 hover:text-navy-900 active:bg-slate-100 shadow-sm cursor-pointer"
          }`}
          title="Halaman Sebelumnya"
        >
          <FaChevronLeft className="text-[11px]" />
          <span>Previous</span>
        </button>

        {/* Page buttons */}
        <div className="hidden sm:flex items-center gap-1">
          {pageNumbers.map((p, idx) =>
            p === "..." ? (
              <span
                key={`ellipsis-${idx}`}
                className="px-2 py-1 text-xs text-slate-400"
              >
                ...
              </span>
            ) : (
              <button
                key={p}
                type="button"
                disabled={isLoading || p === page}
                onClick={() => onPageChange(p as number)}
                className={`min-w-[28px] h-7 px-2 text-xs font-semibold rounded transition border select-none ${
                  p === page
                    ? "bg-[#0a2540] text-white border-[#0a2540] shadow-sm"
                    : "bg-white text-slate-600 border-slate-300 hover:bg-slate-50 cursor-pointer"
                }`}
              >
                {p}
              </button>
            ),
          )}
        </div>

        {/* Mobile current page indicator */}
        <span className="sm:hidden px-2.5 py-1 text-xs font-semibold text-slate-600 bg-slate-100 rounded border border-slate-200">
          {page} / {totalPages}
        </span>

        {/* Next button */}
        <button
          id="next-page-button"
          type="button"
          disabled={page >= totalPages || isLoading}
          onClick={() => onPageChange(page + 1)}
          className={`flex items-center gap-1.5 px-3 py-1.5 rounded-md text-xs font-semibold transition border select-none ${
            page >= totalPages || isLoading
              ? "bg-slate-100 text-slate-400 border-slate-200 cursor-not-allowed opacity-60"
              : "bg-white text-slate-700 border-slate-300 hover:bg-slate-50 hover:text-navy-900 active:bg-slate-100 shadow-sm cursor-pointer"
          }`}
          title="Halaman Berikutnya"
        >
          <span>Next</span>
          <FaChevronRight className="text-[11px]" />
        </button>
      </div>
    </div>
  );
}


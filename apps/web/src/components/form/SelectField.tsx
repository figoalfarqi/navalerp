/* eslint-disable @typescript-eslint/no-explicit-any */
"use client";

import { useState, useEffect, useRef, useMemo } from "react";
import { FaChevronDown, FaSearch } from "react-icons/fa";
import { FaX } from "react-icons/fa6";
import Table, { TableProps } from "../table/Table";
import Modal from "../Modal";
import { PiTableFill } from "react-icons/pi";
import { usePopoverPosition } from "@/hooks/usePopoverPosition";
import { createPortal } from "react-dom";
import { useCloseOnScrollDistance } from "@/hooks/useCloseOnScrollDistance";

export interface SelectOption {
  label: string;
  value: string | number;
  row?: Record<string, any>;
}

interface SelectFieldProps {
  id: string;
  label: string;
  value: string | number;
  options: SelectOption[];
  valueKey?: string;
  labelKey?: string;
  tableProps?: TableProps;
  className?: string;
  wrapperClassName?: string;
  labelClassName?: string;
  columnLength?: number;
  onChange: (
    value: string | number | null,
    label?: string | number | null,
    row?: Record<string, any> | null,
  ) => void;
  required?: boolean;
  disabled?: boolean;
  uncloseable?: boolean;
  placeholder?: string;
  error?: string;
}

export default function SelectField({
  id,
  label,
  value,
  options,
  valueKey,
  labelKey,
  tableProps,
  className = "",
  wrapperClassName = "",
  labelClassName = "",
  columnLength,
  onChange,
  required = false,
  disabled = false,
  uncloseable = false,
  placeholder = "Pilih...",
  error = "",
}: SelectFieldProps) {
  const [showPopover, setShowPopover] = useState(false);
  const [isOpenModal, setIsOpenModal] = useState(false);
  const [searchTerm, setSearchTerm] = useState("");
  // console.log("options inside selectField", options);
  const filteredOptions = useMemo(
    () =>
      options.filter((opt) =>
        String(opt.label).toLowerCase().includes(searchTerm.toLowerCase()),
      ),
    [searchTerm, options],
  );

  type NavSource = "keyboard" | "open" | "mouse" | null;
  const navSourceRef = useRef<NavSource>(null);

  const [activeIndex, setActiveIndex] = useState<number>(-1);

  const inputRef = useRef<HTMLButtonElement>(null);
  const popoverRef = useRef<HTMLDivElement>(null);

  const optionRefs = useRef<Record<string | number, HTMLDivElement | null>>({});
  const openedRef = useRef(false);
  const searchInputRef = useRef<HTMLInputElement>(null);

  const position = usePopoverPosition({
    inputRef,
    popoverRef,
    visible: showPopover,
    popoverHeight: 240,
  });

  useEffect(() => {
    if (!showPopover) return;

    const index = filteredOptions.findIndex((opt) => opt.value === value);
    navSourceRef.current = "open";
    setActiveIndex(index >= 0 ? index : 0);
  }, [showPopover]);

  useEffect(() => {
    if (!showPopover) {
      openedRef.current = false;
      return;
    }

    if (openedRef.current) return; // ⛔ jangan reset lagi

    const index = filteredOptions.findIndex((opt) => opt.value === value);

    setActiveIndex(index >= 0 ? index : 0);
    openedRef.current = true;
  }, [showPopover, value, filteredOptions]);

  useEffect(() => {
    const handleClickOutside = (e: MouseEvent) => {
      const target = e.target as Node;
      if (
        inputRef.current &&
        popoverRef.current &&
        !inputRef.current.contains(target) &&
        !popoverRef.current.contains(target)
      ) {
        setShowPopover(false);
      }
    };
    document.addEventListener("mousedown", handleClickOutside);
    return () => document.removeEventListener("mousedown", handleClickOutside);
  }, []);

  // Tutup saat scroll
  useCloseOnScrollDistance({
    triggerRef: inputRef,
    active: showPopover,
    onClose: () => setShowPopover(false),
    distance: 20,
  });

  useEffect(() => {
    if (!showPopover) return;

    // tunggu portal render
    requestAnimationFrame(() => {
      searchInputRef.current?.focus();
    });
  }, [showPopover]);

  useEffect(() => {
    if (activeIndex < 0) return;

    const source = navSourceRef.current;

    if (source !== "keyboard" && source !== "open") return;

    const opt = filteredOptions[activeIndex];
    const el = opt && optionRefs.current[opt.value];

    requestAnimationFrame(() => {
      if (el) {
        el.scrollIntoView({
          block: "center",
          behavior: "smooth",
        });
      }
    });

    navSourceRef.current = null; // reset
  }, [activeIndex, filteredOptions, showPopover]);

  useEffect(() => {
    if (!showPopover) return;

    const handleKeyDown = (e: KeyboardEvent) => {
      if (filteredOptions.length === 0) return;

      switch (e.key) {
        case "ArrowDown":
          e.preventDefault();
          navSourceRef.current = "keyboard";
          setActiveIndex((prev) =>
            prev < filteredOptions.length - 1 ? prev + 1 : 0,
          );
          break;

        case "ArrowUp":
          e.preventDefault();
          navSourceRef.current = "keyboard";
          setActiveIndex((prev) =>
            prev > 0 ? prev - 1 : filteredOptions.length - 1,
          );
          break;

        case "Enter":
          e.preventDefault();
          const selected = filteredOptions[activeIndex];
          if (selected) {
            inputRef.current?.focus();
            onChange(selected.value, selected.label, selected.row);
            setShowPopover(false);
            setSearchTerm("");
          }
          break;

        case "Escape":
          e.preventDefault();
          inputRef.current?.focus();
          setShowPopover(false);
          break;
      }
    };

    window.addEventListener("keydown", handleKeyDown);
    return () => window.removeEventListener("keydown", handleKeyDown);
  }, [showPopover, activeIndex, filteredOptions, onChange]);

  const baseWrapperClass = "flex flex-col";
  const baseLabelClass = "mb-1 font-medium text-gray-700";
  const baseButtonClass =
    "flex items-center justify-between px-3 py-2 border rounded-sm focus:outline-none focus:border-blue-500 focus:shadow-[0_0_6px_rgba(59,130,246,0.3)] cursor-pointer bg-white disabled:bg-gray-100 disabled:text-gray-400";
  const baseDropdownClass =
    "app-scrollbar absolute z-50 bg-white border rounded-sm shadow-lg min-w-60 h-60 overflow-auto";
  const baseSearchClass =
    "flex items-center gap-2 px-2 py-2 border-b bg-gray-50 sticky top-0 z-10";

  const baseOptionClass = "px-3 py-2 cursor-pointer hover:bg-blue-100";

  return (
    <div className={`${baseWrapperClass} ${wrapperClassName}`}>
      {label && (
        <label htmlFor={id} className={`${baseLabelClass} ${labelClassName}`}>
          {label} {required && <span className="text-red-500">*</span>}
        </label>
      )}
      <div className="relative" style={{ width: columnLength }}>
        <button
          type="button"
          onClick={() => !disabled && setShowPopover(true)}
          onFocus={() => !disabled && setShowPopover(true)}
          onBlur={(e) => {
            const nextFocused = e.relatedTarget as Node | null;
            if (
              popoverRef.current &&
              nextFocused &&
              popoverRef.current.contains(nextFocused)
            ) {
              return;
            }
            setShowPopover(false);
          }}
          className={`${baseButtonClass} ${className}`}
          disabled={disabled}
          ref={inputRef}
        >
          <div className="flex gap-1 items-center justify-center">
            {value && !uncloseable && (
              <div
                onMouseDown={(e) => {
                  e.preventDefault(); // ⛔ cegah focus button
                  e.stopPropagation();
                }}
                onClick={(e) => {
                  e.stopPropagation();
                  onChange(null, null, null);
                }}
                className={`${
                  disabled
                    ? "text-red-300"
                    : "text-red-500 hover:bg-red-100 hover:text-red-600 "
                } rounded-xs p-0.5`}
              >
                <FaX size={10} />
              </div>
            )}
            {value ? (
              options.find((opt) => String(opt.value) === String(value))?.label || (
                <div className="text-gray-400">{placeholder}</div>
              )
            ) : (
              <div className="text-gray-400">{placeholder}</div>
            )}
          </div>
          <div className="flex gap-2 items-center">
            {valueKey && tableProps && (
              <PiTableFill
                strokeWidth={3}
                onMouseDown={(e) => {
                  e.preventDefault(); // ⛔ cegah focus button
                  e.stopPropagation();
                }}
                onClick={(e) => {
                  e.stopPropagation();
                  setIsOpenModal(true);
                  setShowPopover(false);
                }}
                className="hover:h-6 hover:w-6 hover:mr-0 text-xl h-5.5 w-5.5 mr-px transition-all duration-300"
              />
            )}
            <FaChevronDown
              className={`${showPopover && "rotate-180"} ${
                disabled ? " text-gray-300" : " text-gray-500"
              } transition-transform duration-300 `}
            />
          </div>
        </button>
        {showPopover &&
          createPortal(
            <div
              className={baseDropdownClass}
              style={{
                position: "absolute",
                top: position.top,
                left: position.left,
              }}
              ref={popoverRef}
              tabIndex={-1}
            >
              <div className={baseSearchClass}>
                <FaSearch className="text-gray-500" />
                <input
                  ref={searchInputRef}
                  type="text"
                  value={searchTerm}
                  onChange={(e) => setSearchTerm(e.target.value)}
                  placeholder={`Cari ${label}...`}
                  className="flex-1 bg-transparent outline-none text-sm"
                />
              </div>
              {filteredOptions.length > 0 ? (
                // filteredOptions.map((opt, idx) => (
                //   <div
                //     key={idx}
                //     className={baseOptionClass}
                //     onClick={() => {
                //       // console.log("asds", filteredOptions);
                //       onChange(opt.value, opt.label, opt.row);
                //       setShowPopover(false);
                //       setSearchTerm("");
                //     }}
                //   >
                //     {opt.label}
                //   </div>
                // ))
                filteredOptions.map((opt, idx) => (
                  <div
                    key={idx}
                    ref={(el) => {
                      optionRefs.current[opt.value] = el;
                    }}
                    className={`${baseOptionClass} ${
                      idx === activeIndex ? "bg-blue-100 font-medium" : ""
                    }`}
                    onClick={() => {
                      inputRef.current?.focus();
                      onChange(opt.value, opt.label, opt.row);
                      setShowPopover(false);
                      setSearchTerm("");
                    }}
                  >
                    {opt.label}
                  </div>
                ))
              ) : (
                <div className="px-3 py-2 text-gray-500 text-sm">
                  Tidak ada hasil
                </div>
              )}
            </div>,
            document.body,
          )}
      </div>
      {error && <p className="text-red-500 text-sm mt-1">{error}</p>}
      {/* Modal */}
      {tableProps && valueKey && (
        <Modal
          size="7xl"
          isOpen={isOpenModal}
          title={`Pilih ${label}`}
          cancelText="Tutup"
          confirmText=""
          onCancel={() => setIsOpenModal(false)}
        >
          <div className="h-[60vh]">
            <Table
              {...tableProps}
              tableFor="select"
              onCellClick={(row) => {
                onChange(
                  row[valueKey],
                  labelKey ? row[labelKey] : undefined,
                  row,
                );
                setIsOpenModal(false);
              }}
            />
          </div>
        </Modal>
      )}
    </div>
  );
}

"use client";
import { ReactNode } from "react";
import { AiOutlineClose } from "@/components/icons";
import LoaderDots from "./form/LoaderDots";
import Button from "./form/Button";
import { createPortal } from "react-dom";
import { ButtonVariant } from "@/consta/VariantClassesButton";

interface ModalProps {
  isOpen: boolean;
  title?: string;
  children?: ReactNode;
  confirmText?: string;
  cancelText?: string;
  confirmVariant?: ButtonVariant;
  isHideClose?: boolean;
  size?:
    | "sm"
    | "md"
    | "lg"
    | "xl"
    | "2xl"
    | "3xl"
    | "4xl"
    | "5xl"
    | "6xl"
    | "7xl";
  onConfirm?: () => void;
  onCancel?: () => void;
  loading?: boolean;
}

export default function Modal({
  isOpen,
  title,
  children,
  confirmText = "Confirm",
  cancelText = "Cancel",
  confirmVariant = "blue-solid",
  isHideClose = false,
  size = "md",
  onConfirm,
  onCancel,
  loading = false,
}: ModalProps) {
  if (!isOpen) return null;

  const sizeMap = {
    sm: "max-w-sm",
    md: "max-w-md",
    lg: "max-w-lg",
    xl: "max-w-xl",
    "2xl": "max-w-2xl",
    "3xl": "max-w-3xl",
    "4xl": "max-w-4xl",
    "5xl": "max-w-5xl",
    "6xl": "max-w-6xl",
    "7xl": "max-w-7xl",
  };

  return createPortal(
    <div
      className="fixed inset-0 z-40 flex items-center justify-center bg-black/20 backdrop-blur-sm"
      onClick={onCancel} // close when clicking outside modal
    >
      <div
        className={`bg-white rounded-lg shadow-lg py-6 px-3 w-[calc(100%-16px)] mx-auto relative ${sizeMap[size]}`}
        onClick={(e) => e.stopPropagation()} // prevent close when clicking inside
      >
        {/* Close Icon */}
        {!isHideClose && (
          <Button
            variant="red-ghost"
            size="sm"
            wrapperClassName="!absolute top-3 right-3 text-gray-500 hover:text-gray-700"
            onClick={onCancel}
            id={"close-button"}
          >
            <AiOutlineClose size={20} />
          </Button>
        )}

        {/* Title */}
        {title && <h2 className="text-lg font-semibold mb-4 ml-3">{title}</h2>}

        {/* Content */}
        <div className="app-scrollbar mb-4 overflow-auto px-3 max-h-[70vh]">
          {children}
        </div>

        {/* Actions */}
        <div className="flex justify-end gap-3 mr-3">
          {cancelText && (
            <Button
              id="cancel-button"
              onClick={onCancel}
              variant="gray-outline"
              size="md"
              disabled={loading}
            >
              {cancelText}
            </Button>
          )}
          {confirmText && (
            <Button
              id="confirm-button"
              onClick={onConfirm}
              variant={confirmVariant}
              size="md"
              disabled={loading}
            >
              {loading ? <LoaderDots /> : confirmText}
            </Button>
          )}
        </div>
      </div>
    </div>,
    document.body,
  );
}

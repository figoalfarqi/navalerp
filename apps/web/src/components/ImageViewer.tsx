"use client";

import { resolveFileUrl } from "@/utils/globalUtils";
import Image from "next/image";
import { useCallback, useEffect, useMemo, useRef, useState } from "react";
import {
  LuCheck,
  LuChevronLeft,
  LuChevronRight,
  LuDownload,
  LuTable,
  LuX,
} from "@/components/icons";
import Modal from "./Modal";
import TextAreaField from "./form/TextAreaField";
import ApprovalTable, {
  ApprovalBaseTableNameType,
} from "./approval/ApprovalTable";
import { useAuth } from "@/context/AuthContext";
import ApprovalMarkIcon from "./approval/ApprovalMarkIcon";
import { isApproveableData } from "@/utils/approvement";
interface ImageViewerProps {
  images: string[];
  activeIndex?: number;
  itemInImageViewer?: any;
  onClose: () => void;
  imageIDs?: number[];
  addedImageName?: string;
  approvalBaseTableName?: ApprovalBaseTableNameType;
  approvalDataAll?: Record<string, any>[][];
  onApprove?: (payload: Record<string, any>) => Promise<void>;
  onReject?: (payload: Record<string, any>) => Promise<void>;
}

export default function ImageViewer({
  images,
  activeIndex = 0,
  itemInImageViewer,
  onClose,
  imageIDs,
  addedImageName,
  approvalBaseTableName,
  approvalDataAll,
  onApprove,
  onReject,
}: ImageViewerProps) {
  const [current, setCurrent] = useState(activeIndex);

  // Zoom states
  const [isZoomed, setIsZoomed] = useState(false);
  const [zoomPos, setZoomPos] = useState({ x: 50, y: 50 });

  const [showModalApproval, setShowModalApproval] = useState<
    "table" | "approve" | "reject" | null
  >();

  const [approvalNote, setApprovalNote] = useState("");

  const approvalNoteRef = useRef<HTMLTextAreaElement>(null);

  const [isloadingOnConfirm, setIsloadingOnConfirm] = useState(false);

  const { adminPayload } = useAuth();
  const { isApproveable, isApproveable1 } = useMemo(() => {
    if (onApprove) {
      return isApproveableData({
        approvalData: approvalDataAll?.[current],
        adminPayload,
        firstStepAllowed: [3, 4, 5, 6],
      });
    } else {
      return {
        isApproveable: false,
        isApproveable1: false,
        isApproveable2: false,
        isApproveable3: false,
      };
    }
  }, [approvalDataAll, current, adminPayload]);

  // if (!images || images.length === 0) {
  //   return <p className="text-gray-500">No images available</p>;
  // }

  const prevImage = () => {
    setCurrent((prev) => (prev === 0 ? images.length - 1 : prev - 1));
    setIsZoomed(false);
  };

  const nextImage = () => {
    setCurrent((prev) => (prev === images.length - 1 ? 0 : prev + 1));
    setIsZoomed(false);
  };

  const toggleZoom = () => {
    setIsZoomed((z) => !z);
  };

  // Follow cursor when zooming
  const handleMouseMove = (e: React.MouseEvent) => {
    if (!isZoomed) return;

    const rect = e.currentTarget.getBoundingClientRect();
    const x = ((e.clientX - rect.left) / rect.width) * 100;
    const y = ((e.clientY - rect.top) / rect.height) * 100;
    setZoomPos({ x, y });
  };

  const handleApprovePhoto = useCallback(async () => {
    setIsloadingOnConfirm(true);
    if (imageIDs) {
      const payload = {
        table_id: imageIDs[current],
        approval_step_id: isApproveable1 ? 1 : 2,
        approval_status_id:
          showModalApproval === "approve"
            ? 1
            : showModalApproval === "reject"
              ? 3
              : 2,
        approval_notes: approvalNote,
      };
      if (showModalApproval === "approve") {
        await onApprove?.(payload);
      } else {
        await onReject?.(payload);
      }
      setShowModalApproval(null);
      setIsloadingOnConfirm(false);
    }
  }, [isApproveable1, imageIDs, current, showModalApproval, approvalNote]);

  useEffect(() => {
    const handleKeyDown = (e: KeyboardEvent) => {
      switch (e.key) {
        case "ArrowRight":
          e.preventDefault();
          setCurrent((prev) => (prev < images.length - 1 ? prev + 1 : 0));
          break;

        case "ArrowLeft":
          e.preventDefault();
          setCurrent((prev) => (prev > 0 ? prev - 1 : images.length - 1));
          break;

        case "Enter":
          e.preventDefault();
          if (
            showModalApproval &&
            ["approve", "reject"].includes(showModalApproval)
          ) {
            approvalNoteRef.current?.focus();
            handleApprovePhoto();
          }
          break;

        case "Escape":
          e.preventDefault();
          if (
            showModalApproval &&
            ["approve", "reject"].includes(showModalApproval)
          ) {
            setShowModalApproval(null);
          } else {
            onClose();
          }
          break;
      }
    };

    window.addEventListener("keydown", handleKeyDown);
    return () => window.removeEventListener("keydown", handleKeyDown);
  }, [showModalApproval]);

  useEffect(() => {
    if (showModalApproval === "approve" || showModalApproval === "reject") {
      // tunggu modal render
      setTimeout(() => {
        approvalNoteRef.current?.focus();
      }, 150);
    }
  }, [showModalApproval]);

  const handleDownload = useCallback(async () => {
    try {
      const imageUrl = resolveFileUrl(images[current]);

      // ambil extension dari url
      const ext = imageUrl.split(".").pop()?.split("?")[0] || "jpg";

      // timestamp aman untuk filename
      // const timestamp = new Date().toISOString().replace(/[:.]/g, "-");

      const addedFileNames = [];

      if (itemInImageViewer?.customer?.customer_name) {
        addedFileNames.push(itemInImageViewer.customer.customer_name);
      }

      if (itemInImageViewer?.delivery_order_number) {
        addedFileNames.push(itemInImageViewer.delivery_order_number);
      }

      if (itemInImageViewer?.project_name) {
        addedFileNames.push(itemInImageViewer?.project_name);
      }

      if (addedImageName) {
        addedFileNames.push(addedImageName);
      }

      if (approvalBaseTableName === "delivery_order_note_photo") {
        addedFileNames.push("Surat-Jalan");
      } else if (approvalBaseTableName === "delivery_order_cargo_photo") {
        addedFileNames.push("cargo");
      }

      if (imageIDs && imageIDs[current]) {
        addedFileNames.push(imageIDs[current]);
      }
      const fileName = `${addedFileNames.join("_")}.${ext}`;

      const res = await fetch(imageUrl);
      const blob = await res.blob();

      const blobUrl = URL.createObjectURL(blob);

      const link = document.createElement("a");
      link.href = blobUrl;
      link.download = fileName;
      document.body.appendChild(link);
      link.click();

      document.body.removeChild(link);
      URL.revokeObjectURL(blobUrl);
    } catch (err) {
      console.error("Download failed:", err);
    }
  }, [current]);

  return (
    <div
      className="fixed inset-0 z-40 flex flex-col items-center justify-center bg-black/60 backdrop-blur-sm"
      onMouseMove={handleMouseMove}
    >
      {!images?.length ? (
        <>
          <button
            onClick={onClose}
            className="absolute top-4 right-4 text-white bg-black/50 p-2 rounded-full hover:bg-black/70 cursor-pointer z-30"
          >
            <LuX size={24} />
          </button>
          <p className="text-gray-800">No images available</p>
        </>
      ) : (
        <>
          {/* Tombol Close */}
          {approvalDataAll && (
            <button
              onClick={() => {
                setShowModalApproval("table");
              }}
              className={`${isApproveable ? "right-52" : "right-28"} absolute top-4  text-white bg-black/50 p-2 rounded-full hover:bg-black/70 cursor-pointer z-30`}
            >
              <LuTable size={24} />
            </button>
          )}

          {onReject && isApproveable && (
            <button
              onClick={() => {
                setShowModalApproval("reject");
              }}
              className="absolute top-4 right-40 text-white bg-red-500/50 p-2 rounded-full hover:bg-red-500/70 cursor-pointer z-30"
            >
              <LuX size={24} />
            </button>
          )}
          {onApprove && isApproveable && (
            <button
              onClick={() => {
                setShowModalApproval("approve");
              }}
              className="absolute top-4 right-28 text-white bg-green-500/50 p-2 rounded-full hover:bg-green-500/70 cursor-pointer z-30"
            >
              <LuCheck size={24} />
            </button>
          )}

          <button
            onClick={() => handleDownload()}
            className="absolute top-4 right-16 text-white bg-black/50 p-2 rounded-full hover:bg-black/70 cursor-pointer z-30"
          >
            <LuDownload size={24} />
          </button>

          <button
            onClick={onClose}
            className="absolute top-4 right-4 text-white bg-black/50 p-2 rounded-full hover:bg-black/70 cursor-pointer z-30"
          >
            <LuX size={24} />
          </button>

          {/* Main image with zoom */}
          <div className="relative w-full flex items-center justify-center px-2">
            <Image
              src={`${resolveFileUrl(images[current])}`}
              alt={`image-${current}`}
              width={1200}
              height={800}
              onClick={toggleZoom}
              className={`w-auto max-h-[80vh] transition-transform duration-300 cursor-zoom-in
            ${isZoomed ? "cursor-zoom-out" : ""}
          `}
              style={{
                transform: isZoomed ? `scale(2)` : "scale(1)",
                transformOrigin: `${zoomPos.x}% ${zoomPos.y}%`,
              }}
              priority={current === 0}
              unoptimized
            />

            {/* Prev/Next */}
            {!isZoomed && (
              <>
                <button
                  type="button"
                  onClick={prevImage}
                  className="absolute left-4 top-1/2 -translate-y-1/2 bg-white/30 p-3 rounded-full shadow hover:bg-white/60 cursor-pointer"
                >
                  <LuChevronLeft size={24} opacity={0.3} />
                </button>
                <button
                  type="button"
                  onClick={nextImage}
                  className="absolute right-4 top-1/2 -translate-y-1/2 bg-white/30 p-3 rounded-full shadow hover:bg-white/60 cursor-pointer"
                >
                  <LuChevronRight size={24} opacity={0.3} />
                </button>
              </>
            )}
          </div>

          {/* Thumbnails */}
          <div className="mt-6 flex space-x-3 overflow-x-auto max-w-full px-4">
            {images.map((img, idx) => (
              <div
                key={idx}
                className="relative cursor-pointer"
                onClick={() => {
                  setCurrent(idx);
                  setIsZoomed(false);
                }}
              >
                <Image
                  src={`${resolveFileUrl(img)}`}
                  width={64}
                  height={64}
                  className={`w-16 h-16 object-cover rounded-lg border-2 ${
                    idx === current ? "border-blue-500" : "border-transparent"
                  }`}
                  alt={`thumbnail-${idx}`}
                  unoptimized
                />
                {approvalBaseTableName && (
                  <div className="absolute right-1 bottom-1">
                    <ApprovalMarkIcon approvalData={approvalDataAll?.[idx]} />
                  </div>
                )}
              </div>
            ))}
          </div>
          <Modal
            isOpen={
              !!showModalApproval &&
              ["approve", "reject"].includes(showModalApproval)
            }
            onCancel={() => {
              setShowModalApproval(null);
            }}
            confirmText={showModalApproval === "approve" ? "Approve" : "Reject"}
            confirmVariant={
              showModalApproval === "approve" ? "green-solid" : "red-solid"
            }
            loading={isloadingOnConfirm}
            onConfirm={() => {
              handleApprovePhoto();
            }}
          >
            <div>
              <div>
                Apakah Anda yakin ingin{" "}
                <span
                  className={`${showModalApproval === "approve" ? "text-green-500" : "text-red-500"}`}
                >
                  {showModalApproval}
                </span>{" "}
                item ini?
              </div>
              <div>Tindakan ini tidak dapat dibatalkan.</div>
              <Image
                src={`${resolveFileUrl(images[current])}`}
                alt={`image-${current}`}
                width={1200}
                height={800}
                className={`w-auto max-h-[80vh] transition-transform duration-300
            
          `}
                priority={current === 0}
                unoptimized
              />
              <div className="h-[130px]">
                <TextAreaField
                  ref={approvalNoteRef}
                  value={approvalNote}
                  className="w-full"
                  id="approval-note"
                  label="Catatan"
                  onChange={(val) => setApprovalNote(val as string)}
                />
              </div>
            </div>
          </Modal>

          <Modal
            size="5xl"
            isOpen={
              !!showModalApproval && ["table"].includes(showModalApproval)
            }
            onCancel={() => {
              setShowModalApproval(null);
            }}
            onConfirm={() => {
              setShowModalApproval(null);
            }}
            confirmText={"Ok"}
          >
            {approvalBaseTableName && (
              <ApprovalTable
                approvalData={approvalDataAll?.[current] ?? []}
                approvalBaseTableName={approvalBaseTableName}
              />
            )}
          </Modal>
        </>
      )}
    </div>
  );
}

"use client";
import {
  FaArrowRight,
  FaCopy,
  FaEdit,
  FaEye,
  FaTrash,
  FaTrashAlt,
  FaTrashRestore,
  FaTrashRestoreAlt,
} from "@/components/icons";
import Button from "../form/Button";
import { FaMapLocation } from "@/components/icons";
import { MdAssignmentInd, MdAttachMoney } from "@/components/icons";

interface TableToolbarProps {
  selectedCount: number;
  onCopy?: () => void;
  onEdit?: () => void;
  onView?: () => void;
  onDelete?: () => void;
  onRestore?: () => void;
  onActivate?: () => void;
  onDeactivate?: () => void;
  onApproveTA?: () => void;
  onApproval?: () => void;
  onOpen?: () => void;
  onAssign?: () => void;
  assignIcon?: React.ReactNode;
  isOnOpenDisabled?: boolean;
}

export default function TableToolbar({
  selectedCount,
  onCopy,
  onEdit,
  onView,
  onDelete,
  onRestore,
  onActivate,
  onDeactivate,
  onOpen,
  onApproveTA,
  onApproval,
  onAssign,
  assignIcon,
  isOnOpenDisabled,
}: TableToolbarProps) {
  const iconNonActive = () => {
    return selectedCount > 1 ? (
      <div className="flex flex-col justify-between">
        <div className="border border-white w-1 h-1"></div>
        <div className="border border-white w-1 h-1"></div>
        <div className="border border-white w-1 h-1"></div>
      </div>
    ) : (
      <div className="border border-white w-1 h-4"></div>
    );
  };

  const iconActive = () => {
    return selectedCount > 1 ? (
      <div className="flex flex-col justify-between">
        <div className="bg-white w-1 h-1"></div>
        <div className="bg-white w-1 h-1"></div>
        <div className="bg-white w-1 h-1"></div>
      </div>
    ) : (
      <div className="bg-white w-1 h-4"></div>
    );
  };
  return (
    <div
      className={`flex items-center gap-2 mb-2 rounded transition-all duration-300 overflow-hidden`}
      style={{
        maxHeight: selectedCount > 0 ? `${32}px` : "0px",
      }}
    >
      {/* {selectedCount === 1 && ( */}
      <>
        <Button
          id="copy-button"
          disabled={selectedCount > 1}
          onClick={onCopy}
          variant="blue-solid"
        >
          <FaCopy />
        </Button>
        <Button
          id="edit-button"
          disabled={selectedCount > 1}
          onClick={onEdit}
          variant="green-solid"
        >
          <FaEdit />
        </Button>
        <Button
          id="view-button"
          disabled={selectedCount > 1}
          onClick={onView}
          variant="gray-solid"
        >
          <FaEye />
        </Button>
        {onOpen && (
          <Button
            id="open-button"
            disabled={selectedCount > 1 || isOnOpenDisabled}
            onClick={onOpen}
            variant="purple-solid"
          >
            <FaMapLocation />
          </Button>
        )}
        {onApproveTA && (
          <Button
            id="approveta-button"
            disabled={selectedCount > 1}
            onClick={onApproveTA}
            variant="amber-solid"
            
          >
            <div className="h-4 w-4">
              <MdAttachMoney className="h-5 w-5 -m-0.5"/>
            </div>
          </Button>
        )}
        {onApproval && (
          <Button
            id="approval-button"
            disabled={selectedCount > 1}
            onClick={onApproval}
            variant="amber-solid"
            className="flex items-center gap-1 px-3 py-1 font-semibold text-xs"
            title="Persetujuan & Otorisasi Digital"
          >
            <span className="font-bold text-xs">✓ Otorisasi</span>
          </Button>
        )}
        {onActivate && (
          <Button id="open-button" onClick={onActivate} variant="yellow-solid">
            <div className="h-4 w-4 flex gap-[1px]">
              {iconNonActive()}
              <div className="flex items-center">
                <FaArrowRight size={5} />
              </div>
              {iconActive()}
            </div>
          </Button>
        )}
        {onDeactivate && (
          <Button
            id="open-button"
            onClick={onDeactivate}
            variant="yellow-solid"
          >
            <div className="h-4 w-4 flex gap-[1px]">
              {iconActive()}
              <div className="flex items-center">
                <FaArrowRight size={5} />
              </div>
              {iconNonActive()}
            </div>
          </Button>
        )}
      </>
      {/* )} */}
      {onDelete && (
        <Button id="delete-button" onClick={onDelete} variant="red-solid">
          {selectedCount > 1 ? <FaTrashAlt /> : <FaTrash />}
        </Button>
      )}

      {onRestore && (
        <Button id="delete-button" onClick={onRestore} variant="yellow-solid">
          {selectedCount > 1 ? <FaTrashRestore /> : <FaTrashRestoreAlt />}
        </Button>
      )}
      {onAssign && (
        <Button
          id="assign-button"
          onClick={onAssign}
          variant="yellow-solid"
          className="py-[5px]! px-[13px]!"
          disabled={selectedCount > 1}
        >
          {assignIcon || <MdAssignmentInd size={22} />}
        </Button>
      )}
    </div>
  );
}

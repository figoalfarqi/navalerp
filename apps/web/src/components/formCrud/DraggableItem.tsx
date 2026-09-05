import { Reorder, useDragControls } from "framer-motion";
import { LuEllipsis } from "react-icons/lu";
import Button from "../form/Button";

interface DraggableItemProps {
  id: string | number;
  dragEnabled: boolean;
  children: React.ReactNode;
  onDelete: () => void;
}

const DraggableItem = ({
  id,
  dragEnabled,
  children,
  onDelete,
}: DraggableItemProps) => {
  const dragControls = useDragControls(); // ✅ legal

  return (
    <Reorder.Item
      value={id}
      drag={dragEnabled}
      dragListener={false}
      dragControls={dragControls}
      initial={{ opacity: 0, height: 0 }}
      animate={{ opacity: 1, height: "auto" }}
      exit={{ opacity: 0, height: 0 }}
      transition={{ duration: 0.25, ease: "easeInOut" }}
      className="relative group"
    >
      <div className="relative group/ellipsis">
        <LuEllipsis
          className={`absolute top-0 right-0 ${
            dragEnabled ? "cursor-grab active:cursor-grabbing" : ""
          } text-transparent group-hover:text-gray-400`}
          size={24}
          onPointerDown={(e) => {
            e.preventDefault();
            e.stopPropagation();
            dragControls.start(e);
          }}
        />
        <div className="absolute right-0 top-6 hidden group-hover/ellipsis:flex group-hover/ellipsis:flex-col bg-white shadow rounded z-10">
          <Button
            id="button-del"
            variant="red-ghost"
            onClick={() => onDelete()}
          >
            Delete
          </Button>
        </div>
      </div>
      {children}
    </Reorder.Item>
  );
};

export default DraggableItem;

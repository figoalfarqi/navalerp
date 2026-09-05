import TextFieldSkeleton from "../form/TextFieldSkeleton";

const ItemCardSkeleton = () => {
  return (
    <div className="border rounded-lg p-[14px] text-sm shadow-sm">
      <div className="flex flex-col space-y-2">
        <TextFieldSkeleton className="h-4! w-3/4" />
        <TextFieldSkeleton className="h-4! w-1/2" />
        <TextFieldSkeleton className="h-4! w-full" />
      </div>
    </div>
  );
};

export default ItemCardSkeleton;

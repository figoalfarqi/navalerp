import { ReactNode } from "react";

const ItemCard = ({
  onClick,
  children,
}: {
  onClick?: () => void;
  children: ReactNode;
}) => {
  const className =
    "w-full border border-gray-200 bg-white rounded-xl p-4 text-sm text-left shadow-sm transition";

  if (onClick) {
    return (
      <button
        type="button"
        className={`${className} cursor-pointer active:scale-[0.99] hover:shadow-md focus:outline-none focus:ring-2 focus:ring-blue-500/40`}
        onClick={onClick}
      >
        {children}
      </button>
    );
  }

  return (
    <div className={className}>{children}</div>
  );
};

export default ItemCard;

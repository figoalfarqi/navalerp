export const InfoRow = ({
  label,
  children,
}: {
  label: string;
  children: React.ReactNode;
}) => (
  <div className="flex gap-4">
    <div className="text-base text-gray-500 w-1/3">{label}</div>
    <div className="text-base text-gray-500">:</div>
    <div className="text-base font-medium text-gray-800">
      {children}
    </div>
  </div>
);

const ApprovalCountIcon = ({
  className,
  count,
}: {
  className: string;
  count: number;
}) => {
  return (
    <div
      className={`${className} min-w-5 h-5 px-1 flex items-center justify-center rounded-full text-white text-sm`}
    >
      {count}
    </div>
  );
};

export default ApprovalCountIcon;

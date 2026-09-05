"use client"
const LoadingRow = ({
  rowCount,
  columnCount,
}: {
  rowCount: number;
  columnCount: number;
}) => {
  return (
    <>
      {Array.from({ length: rowCount }).map((_, idx) => (
        <tr key={idx}>
          {Array.from({ length: columnCount }).map((_, idx2) => (
            <td
              key={idx2}
              className={`text-center p-1 ${idx==0 ? "border-x":"border"}`}
            >
             <div className="bg-gray-500 animate-pulse rounded-sm">{" "} &nbsp;</div>
            </td>
          ))}
        </tr>
      ))}
    </>
  );
};

export default LoadingRow;

// "use client";

// import Button from "../form/Button";

// interface TablePaginationProps {
//   page: number;
//   total: number;
//   limit: number;
//   onPageChange: (page: number) => void;
// }

// export default function TablePagination({
//   page,
//   total,
//   limit,
//   onPageChange,
// }: TablePaginationProps) {
//   const totalPages = Math.ceil(total / limit);

//   return (
//     <div className="flex justify-between items-center mt-4">
//       <span>
//         Page {page} of {totalPages}
//       </span>
//       <div className="flex gap-2">
//         <Button
//           id="prev-button"
//           disabled={page <= 1}
//           onClick={() => onPageChange(page - 1)}
//           className="btn btn-sm bg-gray-200"
//           variant="gray-solid"
//         >
//           Prev
//         </Button>
//         <Button
//           id="next-button"
//           disabled={page >= totalPages}
//           onClick={() => onPageChange(page + 1)}
//           className="btn btn-sm bg-gray-200"
//           variant="gray-solid"
//         >
//           Next
//         </Button>
//       </div>
//     </div>
//   );
// }

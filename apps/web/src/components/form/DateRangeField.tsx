// "use client";

// import { ChangeEvent } from "react";

// interface DateRangeValue {
//   start: string; // ISO "YYYY-MM-DD"
//   end: string; // ISO "YYYY-MM-DD"
// }

// interface DateRangeFieldProps {
//   id: string;
//   label: string;
//   value: DateRangeValue;
//   className?: string;
//   wrapperClassName?: string;
//   labelClassName?: string;
//   onChange: (value: DateRangeValue) => void;
//   required?: boolean;
//   disabled?: boolean;
//   placeholder?: { start?: string; end?: string };
//   error?: string;
// }

// export default function DateRangeField({
//   id,
//   label,
//   value,
//   className = "",
//   wrapperClassName = "",
//   labelClassName = "",
//   onChange,
//   required = false,
//   disabled = false,
//   placeholder = {},
//   error = "",
// }: DateRangeFieldProps) {
//   const baseWrapperClass = "flex flex-col mb-4";
//   const baseLabelClass = "mb-1 font-medium text-gray-700";
//   const baseInputClass =
//     "px-3 py-2 border ounded-sm focus:outline-none focus:border-blue-500 focus:shadow-[0_0_6px_rgba(59,130,246,0.3)] disabled:bg-gray-100 disabled:text-gray-400";

//   const handleChange = (
//     e: ChangeEvent<HTMLInputElement>,
//     field: "start" | "end"
//   ) => {
//     onChange({ ...value, [field]: e.target.value });
//   };

//   return (
//     <div className={`${baseWrapperClass} ${wrapperClassName}`}>
//       <label htmlFor={id} className={`${baseLabelClass} ${labelClassName}`}>
//         {label} {required && <span className="text-red-500">*</span>}
//       </label>
//       <div className="flex gap-2">
//         <input
//           id={`${id}-start`}
//           type="date"
//           value={value.start}
//           onChange={(e) => handleChange(e, "start")}
//           placeholder={placeholder.start}
//           disabled={disabled}
//           className={`${baseInputClass} ${className} w-full`}
//         />
//         <input
//           id={`${id}-end`}
//           type="date"
//           value={value.end}
//           onChange={(e) => handleChange(e, "end")}
//           placeholder={placeholder.end}
//           disabled={disabled}
//           className={`${baseInputClass} ${className} w-full`}
//         />
//       </div>
//       {error && <p className="text-red-500 text-sm mt-1">{error}</p>}
//     </div>
//   );
// }

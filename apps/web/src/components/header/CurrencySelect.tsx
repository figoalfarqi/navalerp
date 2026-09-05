// "use client";

// import { useEffect, useRef, useState } from "react";
// import ReactCountryFlag from "react-country-flag";
// import { currencyFlags } from "@/utils/currencyFlags";

// export default function CurrencySelect() {
//   const [open, setOpen] = useState(false);
//   const dropdownRef = useRef<HTMLDivElement>(null);

//   // Tutup dropdown kalau klik di luar
//   useEffect(() => {
//     function handleClickOutside(event: MouseEvent) {
//       if (
//         dropdownRef.current &&
//         !dropdownRef.current.contains(event.target as Node)
//       ) {
//         setOpen(false);
//       }
//     }

//     document.addEventListener("mousedown", handleClickOutside);
//     return () => document.removeEventListener("mousedown", handleClickOutside);
//   }, []);


//   useEffect(() => {
//     fetchRates();

//     // load target dari localStorage
//     const savedTarget = localStorage.getItem("currency_target");
//     if (savedTarget) {
//       setTarget(savedTarget as "IDR" | "JPY" | "USD");
//     }
//   }, [fetchRates, setTarget]);

//   return (
//     <div ref={dropdownRef} className="relative w-22">
//       {/* Button */}
//       <button
//         onClick={() => setOpen(!open)}
//         className="flex w-full items-center justify-between bg-white px-4 py-3 text-left cursor-pointer"
//       >
//         <span className="flex items-center gap-2">
//           <ReactCountryFlag
//             className="drop-shadow-[0_0_6px_rgba(0,0,0,0.3)]"
//             countryCode={currencyFlags[target].code}
//             svg
//             style={{ fontSize: "1.25rem" }}
//           />
//           {target}
//         </span>
//       </button>

//       {/* Dropdown Options */}
//       {open && (
//         <ul className="absolute z-10 mt-1 w-full bg-white shadow-lg">
//           {Object.entries(currencyFlags).map(([currency, { code }]) => (
//             <li
//               key={currency}
//               onClick={() => {
//                 setTarget(currency as "IDR" | "JPY" | "USD");
//                 setOpen(false);
//               }}
//               className="flex cursor-pointer items-center gap-2 px-4 py-2 hover:bg-gray-100"
//             >
//               <ReactCountryFlag
//                 className="drop-shadow-[0_0_6px_rgba(0,0,0,0.3)] w-fit"
//                 countryCode={code}
//                 svg
//                 style={{ fontSize: "1.25rem" }}
//               />
//               <span>{currency}</span>
//             </li>
//           ))}
//         </ul>
//       )}
//     </div>
//   );
// }

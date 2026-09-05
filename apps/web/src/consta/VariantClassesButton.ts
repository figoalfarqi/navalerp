export const VariantClassesButtonMap = {
  // 🔹 Solid
  "amber-solid": "bg-amber-500 text-white hover:bg-amber-600",
  "blue-solid": "bg-blue-500 text-white hover:bg-blue-600",
  "red-solid": "bg-red-500 text-white hover:bg-red-600",
  "yellow-solid": "bg-yellow-500 text-white hover:bg-yellow-600",
  "green-solid": "bg-green-500 text-white hover:bg-green-600",
  "purple-solid": "bg-purple-500 text-white hover:bg-purple-600",
  "gray-solid": "bg-gray-500 text-white hover:bg-gray-600",
  "blue-dkl": "bg-[#004f7f] text-white hover:bg-[#02304d]",

  // 🔹 Outline
  "amber-outline": "border border-amber-500 text-amber-500 hover:bg-amber-50",
  "blue-outline": "border border-blue-500 text-blue-500 hover:bg-blue-50",
  "red-outline": "border border-red-500 text-red-500 hover:bg-red-50",
  "yellow-outline":
    "border border-yellow-500 text-yellow-500 hover:bg-yellow-50",
  "green-outline": "border border-green-500 text-green-500 hover:bg-green-50",
  "purple-outline":
    "border border-purple-500 text-purple-500 hover:bg-purple-50",
  "gray-outline": "border border-gray-500 text-gray-500 hover:bg-gray-50",

  // 🔹 Ghost
  "amber-ghost": "bg-transparent text-amber-500 hover:bg-amber-50",
  "blue-ghost": "bg-transparent text-blue-500 hover:bg-blue-50",
  "red-ghost": "bg-transparent text-red-500 hover:bg-red-50",
  "yellow-ghost": "bg-transparent text-yellow-500 hover:bg-yellow-50",
  "green-ghost": "bg-transparent text-green-500 hover:bg-green-50",
  "purple-ghost": "bg-transparent text-purple-500 hover:bg-purple-50",
  "gray-ghost": "bg-transparent text-gray-500 hover:bg-gray-50",
};

export type ButtonVariant = keyof typeof VariantClassesButtonMap;

export const SizeClassButtonMap = {
  "2xs": "px-2 py-0.5 text-xs",
  xs: "px-2.5 py-1 text-xs",
  sm: "px-3 py-1.5 text-sm",
  md: "px-4 py-2 text-base",
  lg: "px-5 py-2.5 text-lg",
  xl: "px-6 py-3 text-xl",
  "2xl": "px-7 py-3.5 text-2xl",
};

export type ButtonSize = keyof typeof SizeClassButtonMap;


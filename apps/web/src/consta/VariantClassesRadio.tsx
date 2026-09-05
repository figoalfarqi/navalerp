export const VariantClassesRadioMap: Record<string, string> = {
  "blue-outline": "border border-blue-500 text-blue-500 hover:bg-blue-50",
  "red-outline": "border border-red-500 text-red-500 hover:bg-red-50",
  "yellow-outline":
    "border border-yellow-500 text-yellow-500 hover:bg-yellow-50",
  "green-outline": "border border-green-500 text-green-500 hover:bg-green-50",
  "purple-outline":
    "border border-purple-500 text-purple-500 hover:bg-purple-50",
  "gray-outline": "border border-gray-500 text-gray-500 hover:bg-gray-50",
};

export const VariantSelectedBgRadioMap: Record<string, string> = {
  "blue-outline": "shadow-[0_0_6px_rgba(59,130,246,0.3)]", // blue-500
  "red-outline": "shadow-[0_0_6px_rgba(239,68,68,0.3)]", // red-500
  "yellow-outline": "shadow-[0_0_6px_rgba(234,179,8,0.3)]", // yellow-500
  "green-outline": "shadow-[0_0_6px_rgba(34,197,94,0.3)]", // green-500
  "purple-outline": "shadow-[0_0_6px_rgba(168,85,247,0.3)]", // purple-500
  "gray-outline": "shadow-[0_0_6px_rgba(107,114,128,0.3)]", // gray-500
};

export type RadioVariant = keyof typeof VariantClassesRadioMap;

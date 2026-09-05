"use client";

import BarcodeScanner from "@/components/form/BarcodeScanner";
import { useState } from "react";

export default function Page() {
  const [result, setResult] = useState("");

  return (
    <div className="p-4 flex flex-col gap-4">
      <h1 className="text-lg font-semibold">Scanner Barcode</h1>

      <BarcodeScanner
        onScan={(value) => {
          setResult(value);
        }}
      />

      {result && (
        <div className="p-3 bg-green-100 rounded-md">
          <div className="text-sm text-gray-600">Hasil:</div>
          <div className="font-bold">{result}</div>
        </div>
      )}
    </div>
  );
}

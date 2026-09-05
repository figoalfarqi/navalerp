"use client";
import { useState } from "react";

export default function TableTabs() {
  const [activeTab, setActiveTab] = useState<"detail" | "status" | "cost">(
    "detail"
  );

  return (
    <div>
      {/* Tab Header */}
      <div className="flex border-b mb-4">
        <button
          type="button"
          className={`cursor-pointer px-4 py-2 ${
            activeTab === "detail"
              ? "border-b-2 border-blue-500 text-blue-600 font-semibold"
              : "text-gray-500"
          }`}
          onClick={() => setActiveTab("detail")}
        >
          Delivery Order Detail
        </button>
        <button
          type="button"
          className={`cursor-pointer px-4 py-2 ${
            activeTab === "status"
              ? "border-b-2 border-blue-500 text-blue-600 font-semibold"
              : "text-gray-500"
          }`}
          onClick={() => setActiveTab("status")}
        >
          Delivery Order Status
        </button>
        <button
          type="button"
          className={`cursor-pointer px-4 py-2 ${
            activeTab === "cost"
              ? "border-b-2 border-blue-500 text-blue-600 font-semibold"
              : "text-gray-500"
          }`}
          onClick={() => setActiveTab("cost")}
        >
          Irregular Cost
        </button>
      </div>

      {/* Tab Content */}
      <div className="p-4 border rounded-md bg-white">
        {activeTab === "detail" && <div>delivery_order detail Tabel</div>}
        {activeTab === "status" && <div>delivery_order status Tabel</div>}
        {activeTab === "cost" && <div>Irregular cost Tabel</div>}
      </div>
    </div>
  );
}

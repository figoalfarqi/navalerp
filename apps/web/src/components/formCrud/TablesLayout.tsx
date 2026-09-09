import { useEffect, useState } from "react";
import FormCrud, { FormDataObject, tablesProps } from "./FormCrud";
import Table from "../table/Table";

const TablesLayout = ({
  tables,
  parentFormData,
  resetKey,
}: {
  tables: tablesProps[];
  parentFormData?: FormDataObject;
  resetKey?: string;
}) => {
  const [activeTab, setActiveTab] = useState(0);

  return (
    <div className="col-span-1 md:col-span-2">
      {/* TAB HEADERS */}
      <div className="app-scrollbar flex flex-nowrap mb-4 overflow-x-auto">
        {tables?.map((table, idx) => {
          const isActive = activeTab === idx;
          return (
            <div
              key={idx}
              onClick={() => setActiveTab(idx)}
              className={`px-4 py-2 border-b-4 whitespace-nowrap cursor-pointer font-bold text-lg transition-all duration-300 ${
                isActive
                  ? "border-blue-500 text-blue-500"
                  : "border-gray-500 text-gray-500 hover:text-gray-700 hover:border-gray-700"
              }`}
            >
              {table.title}
            </div>
          );
        })}
      </div>
      <div className="overflow-hidden">
        <div
          className="flex transition-transform duration-500"
          style={{
            width: `${tables.length * 100}%`,
            transform: `translateX(-${activeTab * (100 / tables.length)}%)`,
          }}
        >
          {tables?.map((table, index) => {
            return (
              <div
                key={index}
                className="px-2"
                style={{ width: `${100 / tables.length}%` }}
              >
                <div className="flex flex-col gap-4">
                  <FormCrud
                    {...table.formCrudProps}
                    resetKey={resetKey}
                    parentFormData={parentFormData}
                  />
                  <Table {...table.tableProps} />
                </div>
              </div>
            );
          })}
        </div>
      </div>
    </div>
  );
};

export default TablesLayout;

/* eslint-disable @typescript-eslint/no-explicit-any */
function getNestedValue(obj: any, path: string): any {
    return path.split(".").reduce((acc, key) => acc?.[key], obj);
}

export function sortData<T extends Record<string, any>>(
    data: T[],
    sortBy: string,
    sortOrder: "asc" | "desc"
): T[] {
    if (!sortBy) return data;

    return [...data].sort((a, b) => {

        const valA = getNestedValue(a, sortBy);
        const valB = getNestedValue(b, sortBy);
        // handle null/undefined
        if (valA == null) return 1;
        if (valB == null) return -1;

        // angka
        if (typeof valA === "number" && typeof valB === "number") {
            return sortOrder === "asc" ? valA - valB : valB - valA;
        }

        // tanggal
        if (valA instanceof Date && valB instanceof Date) {
            return sortOrder === "asc"
                ? valA.getTime() - valB.getTime()
                : valB.getTime() - valA.getTime();
        }

        // fallback string
        return sortOrder === "asc"
            ? String(valA).localeCompare(String(valB))
            : String(valB).localeCompare(String(valA));
    });
}

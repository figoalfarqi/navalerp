export function formatCurrencyIDR(value: number | string): string {
    const num = Number(value);

    if (isNaN(num)) return "Rp0";


    const formatted = new Intl.NumberFormat("id-ID", {
        style: "currency",
        currency: "IDR",
        minimumFractionDigits: num % 1 === 0 ? 0 : 2,
        maximumFractionDigits: 2,
    }).format(num);

    return formatted.replace(/\u00A0/g, "").replace("Rp", "Rp");
}

export function formatCurrencyIDRFixed(value: number | string): string {
    const num = Number(value);

    if (isNaN(num)) return "Rp0,00";

    const formatted = new Intl.NumberFormat("id-ID", {
        style: "currency",
        currency: "IDR",
        minimumFractionDigits: 2,
        maximumFractionDigits: 2,
    }).format(num);

    return formatted.replace(/\u00A0/g, "").replace("Rp", "Rp");
}


export function formatNumberID(value: number): string {
    const num = Number(value);

    if (isNaN(num)) return "0";

    return new Intl.NumberFormat("id-ID", {
        minimumFractionDigits: num % 1 === 0 ? 0 : 1,
        maximumFractionDigits: 20,
    }).format(num);
}
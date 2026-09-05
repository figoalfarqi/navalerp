/* eslint-disable @typescript-eslint/no-explicit-any */
import { useState } from "react";

interface Address {
  provinsi?: string;
  kota?: string;
  kecamatan?: string;
  desa?: string;
  alamat?: string;
}

interface UseGeocodingResult {
  address: Address;
  loading: boolean;
  error: string | null;
  fetchAddress: (lat: number, lng: number) => Promise<void>;
}

export function useReverseGeocoding(): UseGeocodingResult {
  const [address, setAddress] = useState<Address>({});
  const [loading, setLoading] = useState(false);
  const [error, setError] = useState<string | null>(null);

  const fetchAddress = async (lat: number, lng: number) => {
    setLoading(true);
    setError(null);

    // 1️⃣ Coba BigDataCloud dulu
    try {
      const apiKey = process.env.NEXT_PUBLIC_BIGDATACLOUD_API_KEY;
      const res = await fetch(
        `https://api.bigdatacloud.net/data/reverse-geocode?latitude=${lat}&longitude=${lng}&localityLanguage=id&key=${apiKey}`
      );
      if (!res.ok) throw new Error("BigDataCloud API error");
      const data = await res.json();

      setAddress({
        provinsi: data.principalSubdivision || "",
        kota: data.city || data.locality || "",
        kecamatan: data.locality || "",
        desa: data.locality || "",
        alamat: data.locality || "",
      });
      setLoading(false);
      return;
    } catch (err) {
      console.warn("BigDataCloud failed, fallback to Nominatim", err);
    }

    // 2️⃣ Fallback ke Nominatim
    try {
      const res = await fetch(
        `https://nominatim.openstreetmap.org/reverse?format=jsonv2&lat=${lat}&lon=${lng}`
      );
      if (!res.ok) throw new Error("Nominatim API error");
      const data = await res.json();
      setAddress({
        provinsi: data.address?.state || "",
        kota: data.address?.city || data.address?.town || data.address?.county || "",
        kecamatan: data.address?.suburb || data.address?.neighbourhood || "",
        desa: data.address?.village || data.address?.neighbourhood || "",
        alamat: data.address?.road || data.address?.house_number || "",
      });
    } catch (err: any) {
      setError(err.message || "Error fetching address");
    } finally {
      setLoading(false);
    }
  };

  return { address, loading, error, fetchAddress };
}

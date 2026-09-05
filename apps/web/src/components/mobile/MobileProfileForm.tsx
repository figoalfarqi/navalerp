"use client";

import Button from "@/components/form/Button";
import LoaderDots from "@/components/form/LoaderDots";
import SelectField, { SelectOption } from "@/components/form/SelectField";
import TextField from "@/components/form/TextField";
import BackButton from "@/components/mobile/BackButton";
import MobilePageLoader from "@/components/mobile/MobilePageLoader";
import { useToast } from "@/components/ToastContext";
import { authTokenType, useFetchAPI } from "@/hooks/useFetchAPI";
import { useEffect, useRef, useState } from "react";

type MobileRole = Extract<authTokenType, "checker" | "driver">;

type ProfileFormData = {
  app_user_name: string;
  app_user_phone: string;
  app_user_address: string;
  city_id?: number;
};

type CityRow = {
  city_id: number;
  city_name: string;
};

const readItems = <T,>(value: unknown): T[] => {
  if (Array.isArray(value)) return value as T[];
  if (!value || typeof value !== "object") return [];
  const record = value as Record<string, unknown>;
  if (Array.isArray(record.items)) return record.items as T[];
  if (record.data && typeof record.data === "object") {
    const nested = record.data as Record<string, unknown>;
    if (Array.isArray(nested.items)) return nested.items as T[];
  }
  return [];
};

export default function MobileProfileForm({
  role,
  userId,
}: {
  role: MobileRole;
  userId: number;
}) {
  const { showToast } = useToast();
  const { getAPI, putAPI } = useFetchAPI();
  const showToastRef = useRef(showToast);
  useEffect(() => {
    showToastRef.current = showToast;
  }, [showToast]);

  const [formData, setFormData] = useState<ProfileFormData>({
    app_user_name: "",
    app_user_phone: "",
    app_user_address: "",
  });
  const [cityOptions, setCityOptions] = useState<SelectOption[]>([]);
  const [isLoading, setIsLoading] = useState(true);
  const [isSubmitting, setIsSubmitting] = useState(false);

  useEffect(() => {
    if (!userId) return;
    let cancelled = false;

    const loadProfile = async () => {
      setIsLoading(true);
      const baseUrl = process.env.NEXT_PUBLIC_API_BASE_URL;
      const [profileResponse, cityResponse] = await Promise.all([
        getAPI<unknown>(`${baseUrl}/${role}/${role}/${userId}`, {
          authToken: role,
        }),
        getAPI<unknown>(`${baseUrl}/${role}/city?limit=399`, {
          authToken: role,
        }),
      ]);

      if (cancelled) return;
      if (profileResponse.code === 200 && profileResponse.data) {
        const profile = profileResponse.data as Partial<ProfileFormData>;
        setFormData({
          app_user_name: profile.app_user_name ?? "",
          app_user_phone: profile.app_user_phone ?? "",
          app_user_address: profile.app_user_address ?? "",
          city_id: profile.city_id,
        });
      } else {
        showToastRef.current(
          3000,
          "error",
          profileResponse.message || "Gagal memuat profil",
        );
      }

      setCityOptions(
        cityResponse.code === 200 && cityResponse.data
          ? readItems<CityRow>(cityResponse.data).map((city) => ({
              value: city.city_id,
              label: city.city_name,
            }))
          : [],
      );
      setIsLoading(false);
    };

    void loadProfile();
    return () => {
      cancelled = true;
    };
  }, [getAPI, role, userId]);

  const handleSubmit = async (event: React.FormEvent) => {
    event.preventDefault();
    setIsSubmitting(true);
    const response = await putAPI(
      `${process.env.NEXT_PUBLIC_API_BASE_URL}/${role}/${role}/${userId}`,
      { authToken: role },
      formData,
    );
    setIsSubmitting(false);

    if ([200, 201].includes(response.code)) {
      showToast(2500, "success", "Profil berhasil diperbarui");
    } else {
      showToast(3000, "error", response.message || "Gagal memperbarui profil");
    }
  };

  if (isLoading) return <MobilePageLoader />;

  return (
    <main className="mx-auto min-h-dvh max-w-md bg-slate-50 px-4 pb-24 pt-5">
      <BackButton url={`/${role}/account`} />
      <div className="mb-5 mt-4">
        <p className="text-xs font-semibold uppercase tracking-[0.15em] text-blue-600">
          Akun
        </p>
        <h1 className="mt-1 text-2xl font-bold text-slate-900">
          Profil {role === "checker" ? "checker" : "driver"}
        </h1>
      </div>

      <form
        onSubmit={handleSubmit}
        className="space-y-4 rounded-2xl border border-slate-200 bg-white p-5 shadow-sm"
      >
        <TextField
          id={`${role}_profile_name`}
          label="Nama"
          required
          value={formData.app_user_name}
          onChange={(value) =>
            setFormData((current) => ({
              ...current,
              app_user_name: String(value),
            }))
          }
          className="min-h-12 w-full rounded-lg"
        />
        <TextField
          id={`${role}_profile_phone`}
          label="Nomor telepon"
          required
          value={formData.app_user_phone}
          onChange={(value) =>
            setFormData((current) => ({
              ...current,
              app_user_phone: String(value),
            }))
          }
          className="min-h-12 w-full rounded-lg"
        />
        <TextField
          id={`${role}_profile_address`}
          label="Alamat"
          value={formData.app_user_address}
          onChange={(value) =>
            setFormData((current) => ({
              ...current,
              app_user_address: String(value),
            }))
          }
          className="min-h-12 w-full rounded-lg"
        />
        <SelectField
          id={`${role}_profile_city`}
          label="Kota"
          value={formData.city_id ?? ""}
          options={cityOptions}
          onChange={(value) =>
            setFormData((current) => ({
              ...current,
              city_id: value ? Number(value) : undefined,
            }))
          }
          className="min-h-12 w-full rounded-lg"
          placeholder="Pilih kota"
        />

        <Button
          id={`${role}_save_profile`}
          type="submit"
          disabled={isSubmitting}
          wrapperClassName="w-full pt-2"
          className="min-h-12 w-full rounded-xl"
        >
          {isSubmitting ? <LoaderDots /> : "Simpan perubahan"}
        </Button>
      </form>
    </main>
  );
}

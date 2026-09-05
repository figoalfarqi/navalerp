"use client";

import Button from "@/components/form/Button";
import LoaderDots from "@/components/form/LoaderDots";
import TextField from "@/components/form/TextField";
import BackButton from "@/components/mobile/BackButton";
import { useToast } from "@/components/ToastContext";
import { authTokenType, useFetchAPI } from "@/hooks/useFetchAPI";
import { useState } from "react";

type MobileRole = Extract<authTokenType, "checker" | "driver">;

type PasswordFormData = {
  old_password: string;
  new_password: string;
  retype_new_password: string;
};

const initialForm: PasswordFormData = {
  old_password: "",
  new_password: "",
  retype_new_password: "",
};

export default function MobileChangePasswordForm({
  role,
}: {
  role: MobileRole;
}) {
  const { showToast } = useToast();
  const { putAPI } = useFetchAPI();
  const [formData, setFormData] = useState(initialForm);
  const [isSubmitting, setIsSubmitting] = useState(false);

  const handleSubmit = async (event: React.FormEvent) => {
    event.preventDefault();
    if (formData.new_password !== formData.retype_new_password) {
      showToast(3000, "error", "Konfirmasi password baru tidak sama");
      return;
    }

    setIsSubmitting(true);
    const response = await putAPI(
      `${process.env.NEXT_PUBLIC_API_BASE_URL}/${role}/change_password`,
      { authToken: role },
      formData,
    );
    setIsSubmitting(false);

    if ([200, 201].includes(response.code)) {
      setFormData(initialForm);
      showToast(2500, "success", "Password berhasil diubah");
    } else {
      showToast(3000, "error", response.message || "Password gagal diubah");
    }
  };

  return (
    <main className="mx-auto min-h-dvh max-w-md bg-slate-50 px-4 pb-24 pt-5">
      <BackButton url={`/${role}/account`} />
      <div className="mb-5 mt-4">
        <p className="text-xs font-semibold uppercase tracking-[0.15em] text-blue-600">
          Keamanan
        </p>
        <h1 className="mt-1 text-2xl font-bold text-slate-900">
          Ubah password
        </h1>
      </div>

      <form
        onSubmit={handleSubmit}
        className="space-y-4 rounded-2xl border border-slate-200 bg-white p-5 shadow-sm"
      >
        <TextField
          id={`${role}_old_password`}
          label="Password lama"
          type="password"
          required
          value={formData.old_password}
          onChange={(value) =>
            setFormData((current) => ({
              ...current,
              old_password: String(value),
            }))
          }
          isPasswordShowable
          className="min-h-12 w-full rounded-lg"
        />
        <TextField
          id={`${role}_new_password`}
          label="Password baru"
          type="password"
          required
          value={formData.new_password}
          onChange={(value) =>
            setFormData((current) => ({
              ...current,
              new_password: String(value),
            }))
          }
          isPasswordShowable
          className="min-h-12 w-full rounded-lg"
        />
        <TextField
          id={`${role}_confirm_password`}
          label="Ulangi password baru"
          type="password"
          required
          value={formData.retype_new_password}
          onChange={(value) =>
            setFormData((current) => ({
              ...current,
              retype_new_password: String(value),
            }))
          }
          isPasswordShowable
          className="min-h-12 w-full rounded-lg"
        />

        <Button
          id={`${role}_change_password`}
          type="submit"
          disabled={isSubmitting}
          wrapperClassName="w-full pt-2"
          className="min-h-12 w-full rounded-xl"
        >
          {isSubmitting ? <LoaderDots /> : "Simpan password"}
        </Button>
      </form>
    </main>
  );
}

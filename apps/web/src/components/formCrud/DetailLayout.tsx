import { Dispatch, SetStateAction, useEffect } from "react";
import { useFetchAPI } from "@/hooks/useFetchAPI";
import { ResolveUrl } from "@/utils/globalUtils";
import DetailForm from "./DetailForm";
import { DetailFormProps, FormDataObject } from "./FormCrud";

interface DetailLayoutProps {
  detailForms: DetailFormProps[];
  parentFormData?: FormDataObject;
  formData: FormDataObject;
  setFormData: Dispatch<SetStateAction<FormDataObject>>;
  errorForm: Record<string, string>;
}

interface DetailFormSectionProps extends Omit<DetailLayoutProps, "detailForms"> {
  detailForm: DetailFormProps;
}

function DetailFormSection({
  detailForm,
  parentFormData,
  formData,
  setFormData,
  errorForm,
}: DetailFormSectionProps) {
  const { getAPI } = useFetchAPI();
  const dependsKeys = detailForm.detailFormDependsOn;
  const dependsValues = dependsKeys?.map((key) => formData[key]) ?? [];
  const dependencySignature = JSON.stringify(dependsValues);
  const isVisible =
    dependsKeys?.every(
      (key) => formData[key] !== undefined && formData[key] !== null,
    ) ?? true;
  const resolvedUrl = detailForm.detailFormGetUrl
    ? ResolveUrl(
        detailForm.detailFormGetUrl,
        parentFormData ? { ...parentFormData, ...formData } : formData,
      )
    : "";

  useEffect(() => {
    if (!resolvedUrl || !isVisible) return;

    let cancelled = false;
    const fetchDetailData = async () => {
      const response = await getAPI<unknown>(resolvedUrl, {
        authToken: "admin",
      });
      if (cancelled || response.code !== 200 || !response.data) return;

      const outer = response.data as Record<string, unknown>;
      const payload =
        outer.data && typeof outer.data === "object"
          ? (outer.data as Record<string, unknown>)
          : outer;
      const detailRows = payload.project_detail_destinations;
      setFormData((previous) => ({
        ...previous,
        [detailForm.table_name]: Array.isArray(detailRows) ? detailRows : [],
      }));
    };

    void fetchDetailData();
    return () => {
      cancelled = true;
    };
  }, [
    dependencySignature,
    detailForm.table_name,
    getAPI,
    isVisible,
    resolvedUrl,
    setFormData,
  ]);

  if (!isVisible) return null;

  return (
    <div className="flex flex-col gap-4">
      <DetailForm
        title={detailForm.title}
        table_key={detailForm.table_key}
        table_name={detailForm.table_name}
        fields={detailForm.fields}
        mode={detailForm.mode}
        maxItems={detailForm.maxItems}
        parentFormData={parentFormData}
        formData={formData}
        setFormData={setFormData}
        errorForm={errorForm}
      />
    </div>
  );
}

export default function DetailLayout({
  detailForms,
  parentFormData,
  formData,
  setFormData,
  errorForm,
}: DetailLayoutProps) {
  return (
    <div className="col-span-1 md:col-span-2 space-y-4">
      {detailForms.map((detailForm) => (
        <DetailFormSection
          key={detailForm.table_name}
          detailForm={detailForm}
          parentFormData={parentFormData}
          formData={formData}
          setFormData={setFormData}
          errorForm={errorForm}
        />
      ))}
    </div>
  );
}

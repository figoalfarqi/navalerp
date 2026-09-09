import { useRouter } from "next/navigation";
import { ReactNode } from "react";
import { BiChevronLeft } from "@/components/icons";

export default function BackButton({
  url,
  icon = <BiChevronLeft size={25} />,
  warna = "text-red-500",
  label = "back",
  className = "", }: {
    url: string;
    warna?: string;
    icon?: ReactNode;
    label?: string;
    className?: string;
  }) {
  const router = useRouter();
  const handleBack = () => {
    if (window.history.length > 1) {
      router.back();
    } else {
      router.push(url);
    }
  };
  return (
    <div
      className={`flex flex-row ${warna} ${className} cursor-pointer`}
      onClick={handleBack}
    >
      {icon} <div>{label}</div>{" "}
    </div>
  );
}

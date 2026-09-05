"use client";

import { useRouter } from "next/navigation";
import Image from "next/image";
import { useIsMobile } from "@/hooks/useIsMobile";

export default function Home() {
  const router = useRouter();

  const { isMobile } = useIsMobile();

  // useEffect(() => {
  //   router.push("/driver/login");
  // }, [router]);

  const goToLogins = [
    {
      title: "I am Driver",
      url: "driver/login",
      image: "/icons/icon_driver.png",
    },
    {
      title: "I am Checker",
      url: "checker/login",
      image: "/icons/icon_checker.png",
    },
    {
      title: "I am Admin",
      url: "admin/login",
      image: "/icons/icon_admin.png",
    },
  ];
  return (
    <main
      className={`flex min-h-screen flex-col items-center justify-center ${
        isMobile ? "px-3" : "px-24"
      }`}
    >
      <h1 className={`${isMobile ? "text-2xl" : "text-4xl"} font-bold`}>
        Welcome to the PML
      </h1>

      <Image
        src="/logo-pml.png"
        alt="App Logo"
        width={isMobile ? 200 : 250}
        height={isMobile ? 200 : 250}
        className="mt-8"
      />
      <div
        className={`${
          isMobile ? "text-2xl max-w-xs" : "text-3xl max-w-sm"
        } flex flex-col gap-3 w-full items-center mt-8`}
      >
        {goToLogins.map((data) => (
          <div
            key={data.title}
            className="flex flex-row w-full px-3 py-1 border border-gray-400 rounded-xl hover:bg-blue-100 hover:border-blue-400 transition-colors duration-300 cursor-pointer"
            onClick={() => {
              router.push(data.url);
            }}
          >
            <Image
              src={data.image}
              alt={data.title}
              width={isMobile ? 60 : 80}
              height={isMobile ? 60 : 80}
            />
            <div className="w-full flex items-center justify-center text-center">
              {data.title}
            </div>
          </div>
        ))}
      </div>
    </main>
  );
}

import type { Metadata } from "next";
import "./globals.css";
import { AuthProvider } from "@/context/AuthContext";
import { ToastProvider } from "@/components/ToastContext";

export const metadata: Metadata = {
  title: "Naval ERP - TNI Angkatan Laut",
  description: "Sistem Informasi & Manajemen Operasi Alutsista Matra Laut",
  icons: {
    icon: "/logo.png",
  },
};

export default function RootLayout({
  children,
}: Readonly<{
  children: React.ReactNode;
}>) {
  return (
    <html lang="en">
      <head>
        {/* <link rel="manifest" href="/manifest.json" /> */}
        <meta name="theme-color" content="#0f172a" />
        <link rel="apple-touch-icon" href="/icons/icon-192x192.png" />
      </head>
      <body>
        <AuthProvider>
          <ToastProvider>
            {/* <ClientSync /> */}
            {children}
          </ToastProvider>
        </AuthProvider>
      </body>
    </html>
  );
}

import type { NextConfig } from "next";

const requiredPublicEnv = [
  "NEXT_PUBLIC_API_BASE_URL",
  "NEXT_PUBLIC_API_FILESERVICE_URL",
] as const;

if (process.env.NODE_ENV === "production") {
  const missingEnv = requiredPublicEnv.filter(
    (name) => !process.env[name]?.trim(),
  );

  if (missingEnv.length > 0) {
    throw new Error(
      `Missing required frontend environment variables: ${missingEnv.join(", ")}`,
    );
  }
}

const nextConfig: NextConfig = {
  reactStrictMode: false,
};

export default nextConfig;

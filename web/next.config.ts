import type { NextConfig } from "next";

const nextConfig: NextConfig = {
  output: 'standalone', // Tells Next.js to trace dependencies and create a minimal production bundle
};

export default nextConfig;

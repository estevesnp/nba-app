/** @type {import('next').NextConfig} */
const nextConfig = {
  images: {
    remotePatterns: [
      {
        protocol: "https",
        hostname: "cdn.nba.com",
        pathname: "**",
      },
    ],
  },
  /* config options here */
};

module.exports = nextConfig;

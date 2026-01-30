import type { Metadata } from "next";
import { Inter } from "next/font/google";
import "./globals.css";
import Header from "@/components/Header";
import CryptoTicker from "@/components/CryptoTicker";
import NewsTicker from "@/components/NewsTicker";
import Footer from "@/components/Footer";

const inter = Inter({
  subsets: ["latin"],
  display: "swap",
  variable: "--font-inter",
});

export const metadata: Metadata = {
  title: "Crypto News Feed - Latest Cryptocurrency News",
  description: "Stay updated with the latest cryptocurrency news from Coindesk, Cointelegraph, and more. Real-time crypto news aggregator.",
  keywords: ["crypto", "cryptocurrency", "bitcoin", "ethereum", "blockchain", "news"],
  authors: [{ name: "Crypto News Feed" }],
  openGraph: {
    title: "Crypto News Feed",
    description: "Latest cryptocurrency news aggregated from top sources",
    type: "website",
  },
};

export default function RootLayout({
  children,
}: Readonly<{
  children: React.ReactNode;
}>) {
  return (
    <html lang="en" className="dark">
      <body className={`${inter.variable} font-sans antialiased`}>
        <div className="min-h-screen bg-gradient-to-br from-gray-950 via-gray-900 to-gray-950">
          <Header />
          <CryptoTicker />
          <NewsTicker />
          <main>{children}</main>
          <Footer />
        </div>
      </body>
    </html>
  );
}

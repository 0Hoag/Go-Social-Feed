import type { Metadata } from "next";
import { Inter } from "next/font/google";
import "./globals.css";
import Header from "@/components/Header";
import CryptoTicker from "@/components/CryptoTicker";
import NewsTicker from "@/components/NewsTicker";
import Footer from "@/components/Footer";
import { LanguageProvider } from "@/context/LanguageContext";
import { Toaster } from "sonner";

const inter = Inter({
  subsets: ["latin"],
  display: "swap",
  variable: "--font-inter",
});

export const metadata: Metadata = {
  title: "CryptoCheck - Smart Contract Security Analysis",
  description: "Phân tích bảo mật Smart Contract tự động với AI. Bảo vệ bạn khỏi Rug Pull, Honeypot và các lỗ hổng bảo mật. Kết quả trong 30 giây.",
  keywords: ["smart contract", "security", "audit", "blockchain", "rug pull", "honeypot", "crypto", "defi"],
  authors: [{ name: "CryptoCheck" }],
  openGraph: {
    title: "CryptoCheck - Bảo mật thông minh, Đầu tư an tâm",
    description: "Phân tích Smart Contract với AI trong 30 giây. Trust Score 0-100.",
    type: "website",
  },
};

export default function RootLayout({
  children,
}: Readonly<{
  children: React.ReactNode;
}>) {
  return (
    <html lang="vi" className="dark">
      <body className={`${inter.variable} font-sans antialiased`}>
        <div className="min-h-screen bg-[#050505]">
          <LanguageProvider>
            <Header />
            <CryptoTicker />
            <NewsTicker />
            <main>{children}</main>
            <Footer />
          </LanguageProvider>
          <Toaster richColors position="top-right" />
        </div>
      </body>
    </html>
  );
}

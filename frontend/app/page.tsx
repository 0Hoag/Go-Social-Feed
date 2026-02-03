import HeroSection from "@/components/HeroSection";
import TrustScoreDemo from "@/components/TrustScoreDemo";
import FeaturesGrid from "@/components/FeaturesGrid";
import HowItWorks from "@/components/HowItWorks";
import SecurityIssues from "@/components/SecurityIssues";
import { Shield, Zap } from "lucide-react";
import Link from "next/link";

export default function Home() {
  return (
    <main className="min-h-screen bg-[#050505] text-gray-200 font-sans selection:bg-cyan-500/20 selection:text-cyan-200">
      {/* Hero Section */}
      <HeroSection />

      {/* Trust Score Demo */}
      <TrustScoreDemo />

      {/* Features Grid */}
      <FeaturesGrid />

      {/* How It Works */}
      <HowItWorks />

      {/* Security Issues */}
      <SecurityIssues />

      {/* Final CTA Section */}
      <section className="py-20 px-4">
        <div className="max-w-4xl mx-auto text-center">
          <div className="bg-gradient-to-br from-cyan-500/10 to-blue-500/10 backdrop-blur-xl border border-cyan-500/20 rounded-3xl p-12 md:p-16">
            <div className="w-20 h-20 mx-auto mb-6 rounded-2xl bg-gradient-to-br from-cyan-500 to-blue-600 flex items-center justify-center shadow-lg shadow-cyan-500/30">
              <Shield className="w-10 h-10 text-white" strokeWidth={2.5} />
            </div>

            <h2 className="text-3xl md:text-5xl font-black text-white mb-4">
              Sẵn sàng bảo vệ tài sản của bạn?
            </h2>

            <p className="text-gray-400 text-lg mb-8 max-w-2xl mx-auto">
              Tham gia hàng ngàn nhà đầu tư thông minh đang sử dụng ChainGuardian AI để bảo vệ tài sản của họ
            </p>

            <div className="flex flex-col sm:flex-row items-center justify-center gap-4">
              <Link
                href="#demo"
                className="group px-8 py-4 bg-gradient-to-r from-cyan-500 to-blue-600 text-white font-bold rounded-xl hover:shadow-lg hover:shadow-cyan-500/50 transition-all duration-300 flex items-center gap-2 text-lg"
              >
                <Zap className="w-5 h-5" />
                Quét Contract Miễn Phí
              </Link>

              <Link
                href="https://t.me/chainguardian_bot"
                target="_blank"
                className="px-8 py-4 bg-white/5 border border-white/10 text-white font-semibold rounded-xl hover:bg-white/10 transition-all duration-300 text-lg"
              >
                Kết nối Telegram Bot
              </Link>
            </div>

            <p className="text-gray-500 text-sm mt-6">
              Không cần đăng ký • Không cần thẻ tín dụng • Kết quả trong 30 giây
            </p>
          </div>
        </div>
      </section>

      {/* Footer */}
      <footer className="border-t border-white/5 py-12 px-4">
        <div className="max-w-7xl mx-auto">
          <div className="grid grid-cols-1 md:grid-cols-4 gap-8 mb-8">
            {/* Brand */}
            <div className="col-span-1 md:col-span-2">
              <div className="flex items-center gap-2 mb-4">
                <div className="w-10 h-10 rounded-xl bg-gradient-to-br from-cyan-500 to-blue-600 flex items-center justify-center">
                  <Shield className="w-5 h-5 text-white" />
                </div>
                <span className="text-xl font-black text-white">ChainGuardian AI</span>
              </div>
              <p className="text-gray-500 text-sm max-w-md">
                Bảo mật thông minh – Đầu tư an tâm. Phân tích Smart Contract tự động với AI, bảo vệ bạn khỏi Rug Pull và Honeypot.
              </p>
            </div>

            {/* Links */}
            <div>
              <h3 className="text-white font-bold mb-4">Sản phẩm</h3>
              <ul className="space-y-2 text-gray-500 text-sm">
                <li><Link href="#demo" className="hover:text-cyan-400 transition-colors">Scan Contract</Link></li>
                <li><Link href="/analysis" className="hover:text-cyan-400 transition-colors">Market Analysis</Link></li>
                <li><Link href="#how-it-works" className="hover:text-cyan-400 transition-colors">Cách hoạt động</Link></li>
              </ul>
            </div>

            <div>
              <h3 className="text-white font-bold mb-4">Cộng đồng</h3>
              <ul className="space-y-2 text-gray-500 text-sm">
                <li><a href="https://t.me/chainguardian" target="_blank" className="hover:text-cyan-400 transition-colors">Telegram</a></li>
                <li><a href="https://twitter.com/chainguardian" target="_blank" className="hover:text-cyan-400 transition-colors">Twitter</a></li>
                <li><a href="https://github.com/chainguardian" target="_blank" className="hover:text-cyan-400 transition-colors">GitHub</a></li>
              </ul>
            </div>
          </div>

          <div className="border-t border-white/5 pt-8 flex flex-col md:flex-row justify-between items-center gap-4">
            <p className="text-gray-500 text-sm">
              © 2026 ChainGuardian AI. All rights reserved.
            </p>
            <div className="flex gap-6 text-gray-500 text-sm">
              <Link href="/privacy" className="hover:text-cyan-400 transition-colors">Privacy Policy</Link>
              <Link href="/terms" className="hover:text-cyan-400 transition-colors">Terms of Service</Link>
            </div>
          </div>
        </div>
      </footer>
    </main>
  );
}

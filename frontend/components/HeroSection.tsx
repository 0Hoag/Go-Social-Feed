"use client";

import { Zap, ArrowRight } from "lucide-react";
import Link from "next/link";
import { CryptoCheckLogo } from "@/components/CryptoCheckLogo";
import { useLanguage } from "@/context/LanguageContext";

export default function HeroSection() {
    const { t } = useLanguage();

    return (
        <section className="relative min-h-[90vh] flex items-center justify-center overflow-hidden">
            {/* Animated Background */}
            <div className="absolute inset-0 bg-gradient-to-br from-cyan-500/5 via-transparent to-blue-500/5" />
            <div className="absolute inset-0 bg-[radial-gradient(circle_at_50%_50%,rgba(0,217,255,0.1),transparent_50%)]" />

            {/* Grid Pattern */}
            <div className="absolute inset-0 bg-[linear-gradient(rgba(255,255,255,0.02)_1px,transparent_1px),linear-gradient(90deg,rgba(255,255,255,0.02)_1px,transparent_1px)] bg-[size:50px_50px]" />

            <div className="relative z-10 max-w-6xl mx-auto px-4 py-20 text-center">
                {/* Logo & Badge */}
                <div className="flex items-center justify-center gap-3 mb-6">
                    <div className="w-16 h-16 rounded-2xl bg-gradient-to-br from-cyan-500 to-blue-600 flex items-center justify-center shadow-lg shadow-cyan-500/20">
                        <CryptoCheckLogo className="w-10 h-10 text-white" />
                    </div>
                </div>

                <h1 className="text-5xl md:text-7xl font-black text-white mb-4 tracking-tight">
                    <span className="text-transparent bg-clip-text bg-gradient-to-r from-cyan-400 to-blue-500">{t.hero.title_highlight}</span>
                </h1>

                <p className="text-xl md:text-2xl text-cyan-400 font-semibold mb-6">
                    {t.hero.badge}
                </p>

                <p
                    className="text-lg md:text-xl text-gray-400 max-w-3xl mx-auto mb-12 leading-relaxed"
                    dangerouslySetInnerHTML={{ __html: t.hero.subtitle + "<br/>" + t.hero.support_text }}
                />

                {/* CTA Buttons */}
                <div className="flex flex-col sm:flex-row items-center justify-center gap-4 mb-16">
                    <Link
                        href="/scanner"
                        className="group px-8 py-4 bg-gradient-to-r from-cyan-500 to-blue-600 text-white font-bold rounded-xl hover:shadow-lg hover:shadow-cyan-500/50 transition-all duration-300 flex items-center gap-2 text-lg"
                    >
                        <Zap className="w-5 h-5" />
                        {t.hero.btn_scan}
                        <ArrowRight className="w-5 h-5 group-hover:translate-x-1 transition-transform" />
                    </Link>

                    <Link
                        href="#how-it-works"
                        className="px-8 py-4 bg-white/5 border border-white/10 text-white font-semibold rounded-xl hover:bg-white/10 transition-all duration-300 text-lg"
                    >
                        {t.hero.btn_demo}
                    </Link>
                </div>

                {/* Trust Indicators */}
                <div className="grid grid-cols-1 md:grid-cols-3 gap-8 max-w-3xl mx-auto">
                    <div className="flex flex-col items-center">
                        <div className="text-3xl font-black text-cyan-400 mb-2">10,000+</div>
                        <div className="text-sm text-gray-500 uppercase tracking-wider">{t.hero.stats.scanned}</div>
                    </div>
                    <div className="flex flex-col items-center">
                        <div className="text-3xl font-black text-green-400 mb-2">$5M+</div>
                        <div className="text-sm text-gray-500 uppercase tracking-wider">{t.hero.stats.protected}</div>
                    </div>
                    <div className="flex flex-col items-center">
                        <div className="text-3xl font-black text-blue-400 mb-2">95%</div>
                        <div className="text-sm text-gray-500 uppercase tracking-wider">{t.hero.stats.accuracy}</div>
                    </div>
                </div>
            </div>

            {/* Bottom Gradient Fade */}
            <div className="absolute bottom-0 left-0 right-0 h-32 bg-gradient-to-t from-[#050505] to-transparent" />
        </section>
    );
}

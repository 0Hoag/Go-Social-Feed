"use client";

import { Shield, Zap } from "lucide-react";
import Link from "next/link";
import { useLanguage } from "@/context/LanguageContext";

export default function CTASection() {
    const { t } = useLanguage();

    return (
        <section className="py-20 px-4">
            <div className="max-w-4xl mx-auto text-center">
                <div className="bg-gradient-to-br from-cyan-500/10 to-blue-500/10 backdrop-blur-xl border border-cyan-500/20 rounded-3xl p-12 md:p-16">
                    <div className="w-20 h-20 mx-auto mb-6 rounded-2xl bg-gradient-to-br from-cyan-500 to-blue-600 flex items-center justify-center shadow-lg shadow-cyan-500/30">
                        <Shield className="w-10 h-10 text-white" strokeWidth={2.5} />
                    </div>

                    <h2 className="text-3xl md:text-5xl font-black text-white mb-4">
                        {t.cta.title}
                    </h2>

                    <p className="text-gray-400 text-lg mb-8 max-w-2xl mx-auto">
                        {t.cta.description}
                    </p>

                    <div className="flex flex-col sm:flex-row items-center justify-center gap-4">
                        <Link
                            href="/scanner"
                            className="group px-8 py-4 bg-gradient-to-r from-cyan-500 to-blue-600 text-white font-bold rounded-xl hover:shadow-lg hover:shadow-cyan-500/50 transition-all duration-300 flex items-center gap-2 text-lg"
                        >
                            <Zap className="w-5 h-5" />
                            {t.cta.btn_scan}
                        </Link>

                        <Link
                            href="https://t.me/chainguardian_bot"
                            target="_blank"
                            className="px-8 py-4 bg-white/5 border border-white/10 text-white font-semibold rounded-xl hover:bg-white/10 transition-all duration-300 text-lg"
                        >
                            {t.cta.btn_telegram}
                        </Link>
                    </div>

                    <p className="text-gray-500 text-sm mt-6">
                        {t.cta.note}
                    </p>
                </div>
            </div>
        </section>
    );
}

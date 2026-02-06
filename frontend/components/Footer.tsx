"use client";

import Link from "next/link";
import { Send, Twitter, Facebook } from "lucide-react";
import { useLanguage } from "@/context/LanguageContext";

export default function Footer() {
    const { t } = useLanguage();

    return (
        <footer className="bg-[#050505] border-t border-white/5 pt-16 pb-8 text-sm">
            <div className="max-w-[1600px] mx-auto px-4">
                <div className="grid grid-cols-1 lg:grid-cols-12 gap-12 mb-12">
                    {/* Brand Column */}
                    <div className="lg:col-span-3 space-y-6">
                        <Link href="/" className="block">
                            <span className="text-2xl font-bold text-blue-500 tracking-tight">
                                CryptoCheck
                            </span>
                        </Link>
                        <p className="text-gray-400 leading-relaxed text-xs">
                            {t.footer.description}
                        </p>
                        <div className="flex items-center gap-4">
                            <a href="#" className="w-8 h-8 rounded-full bg-white/5 flex items-center justify-center text-gray-400 hover:text-white hover:bg-white/10 transition-colors">
                                <Send className="w-4 h-4" />
                            </a>
                            <a href="#" className="w-8 h-8 rounded-full bg-white/5 flex items-center justify-center text-gray-400 hover:text-white hover:bg-white/10 transition-colors">
                                <Twitter className="w-4 h-4" />
                            </a>
                            <a href="#" className="w-8 h-8 rounded-full bg-white/5 flex items-center justify-center text-gray-400 hover:text-white hover:bg-white/10 transition-colors">
                                <Facebook className="w-4 h-4" />
                            </a>
                        </div>
                    </div>

                    {/* Spacer */}
                    <div className="hidden lg:block lg:col-span-1"></div>

                    {/* Links Column 1 */}
                    <div className="lg:col-span-2 space-y-4">
                        <h4 className="text-gray-500 font-bold text-xs uppercase tracking-wider mb-2">{t.footer.info_title}</h4>
                        <ul className="space-y-3">
                            <li><Link href="#" className="text-gray-400 hover:text-white transition-colors text-xs">{t.footer.about_us}</Link></li>
                            <li><Link href="#" className="text-gray-400 hover:text-white transition-colors text-xs">{t.footer.roadmap}</Link></li>
                            <li><Link href="#" className="text-gray-400 hover:text-white transition-colors text-xs">{t.footer.contact}</Link></li>
                        </ul>
                    </div>

                    {/* Links Column 2 */}
                    <div className="lg:col-span-2 space-y-4">
                        <h4 className="text-gray-500 font-bold text-xs uppercase tracking-wider mb-2">{t.footer.resources_title}</h4>
                        <ul className="space-y-3">
                            <li><Link href="#" className="text-gray-400 hover:text-white transition-colors text-xs">{t.footer.eco_cal}</Link></li>
                            <li><Link href="#" className="text-gray-400 hover:text-white transition-colors text-xs">{t.footer.options}</Link></li>
                            <li><Link href="#" className="text-gray-400 hover:text-white transition-colors text-xs">{t.footer.tech_analysis}</Link></li>
                        </ul>
                    </div>

                    {/* Links Column 3 */}
                    <div className="lg:col-span-2 space-y-4">
                        <h4 className="text-gray-500 font-bold text-xs uppercase tracking-wider mb-2">{t.footer.legal_title}</h4>
                        <ul className="space-y-3">
                            <li><Link href="#" className="text-gray-400 hover:text-white transition-colors text-xs">{t.footer.terms}</Link></li>
                            <li><Link href="#" className="text-gray-400 hover:text-white transition-colors text-xs">{t.footer.privacy}</Link></li>
                            <li><Link href="#" className="text-gray-400 hover:text-white transition-colors text-xs">{t.footer.faq}</Link></li>
                        </ul>
                    </div>
                </div>

                {/* Bottom Bar */}
                <div className="border-t border-white/5 pt-8 mt-8">
                    <p className="text-[10px] text-gray-500 text-center leading-relaxed max-w-4xl mx-auto mb-4">
                        {t.footer.risk_warning}
                    </p>
                    <p className="text-[10px] text-gray-600 text-center">
                        {t.footer.copyright}
                    </p>
                </div>
            </div>
        </footer>
    );
}

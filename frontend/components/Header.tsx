"use client";

import Link from "next/link";
import { Globe } from "lucide-react";
import { CryptoCheckLogo } from "@/components/CryptoCheckLogo";
import { useLanguage } from "@/context/LanguageContext";

export default function Header() {
    const { t, language, setLanguage } = useLanguage();

    const toggleLanguage = () => {
        setLanguage(language === 'vi' ? 'en' : 'vi');
    };

    return (
        <header className="sticky top-0 z-50 w-full border-b border-white/5 bg-[#050505]/90 backdrop-blur-md">
            <div className="max-w-[1600px] mx-auto flex h-16 items-center justify-between px-4">
                {/* Logo Area */}
                <Link href="/" className="flex items-center gap-2 group">
                    <div className="w-9 h-9 rounded-lg bg-gradient-to-br from-cyan-500 to-blue-600 flex items-center justify-center group-hover:scale-105 transition-transform shadow-lg shadow-cyan-500/20">
                        <CryptoCheckLogo className="w-6 h-6 text-white" />
                    </div>
                    <span className="text-xl font-black text-white tracking-tight">
                        <span className="text-transparent bg-clip-text bg-gradient-to-r from-cyan-400 to-blue-500">CryptoCheck</span>
                    </span>
                </Link>

                {/* Main Navigation - Centered */}
                <nav className="hidden md:flex absolute left-1/2 -translate-x-1/2 items-center gap-8">
                    <Link href="/" className="text-sm font-medium text-gray-400 hover:text-white cursor-pointer transition-colors">{t.nav.home}</Link>
                    <Link href="/news" className="text-sm font-medium text-gray-400 hover:text-cyan-400 cursor-pointer transition-colors">{t.nav.news}</Link>
                    <Link href="/scanner" className="text-sm font-medium text-gray-400 hover:text-cyan-400 cursor-pointer transition-colors">{t.nav.scan}</Link>
                    <Link href="/analysis" className="text-sm font-medium text-gray-400 hover:text-cyan-400 cursor-pointer transition-colors">{t.nav.analysis}</Link>
                </nav>

                {/* Right Actions: Language Switcher */}
                <div className="flex items-center gap-4">
                    <button
                        onClick={toggleLanguage}
                        className="flex items-center gap-2 px-3 py-1.5 rounded-full bg-white/5 border border-white/10 hover:bg-white/10 transition-all text-sm font-bold text-gray-300 hover:text-white"
                    >
                        <Globe className="w-4 h-4" />
                        <span>{language === 'vi' ? 'VN' : 'EN'}</span>
                    </button>

                    {/* Placeholder for Login if needed later */}
                    {/* <div className="w-[100px] hidden md:block"></div> */}
                </div>
            </div>
        </header>
    );
}

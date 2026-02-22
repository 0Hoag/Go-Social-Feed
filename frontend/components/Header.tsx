"use client";

import Link from "next/link";
import { useState, useRef, useEffect } from "react";
import { CryptoCheckLogo } from "@/components/CryptoCheckLogo";
import { useLanguage } from "@/context/LanguageContext";
import { useAuthStore } from "@/lib/stores/authStore";
import { useRouter } from "next/navigation";
import { toast } from "sonner";

function AuthSection() {
    const { isAuthenticated, user, logout } = useAuthStore();
    const router = useRouter();
    const [open, setOpen] = useState(false);
    const ref = useRef<HTMLDivElement>(null);

    useEffect(() => {
        const handler = (e: MouseEvent) => {
            if (ref.current && !ref.current.contains(e.target as Node)) setOpen(false);
        };
        document.addEventListener('mousedown', handler);
        return () => document.removeEventListener('mousedown', handler);
    }, []);

    const handleLogout = () => {
        logout();
        setOpen(false);
        toast.success('Đã đăng xuất');
        router.push('/');
    };

    if (!isAuthenticated) {
        return (
            <div className="flex items-center gap-2">
                <Link href="/login" className="px-4 py-1.5 text-sm font-medium text-zinc-300 hover:text-white border border-zinc-700 rounded-full hover:bg-zinc-800 transition-all">
                    Đăng nhập
                </Link>
                <Link href="/register" className="px-4 py-1.5 text-sm font-semibold text-white bg-sky-500 hover:bg-sky-400 rounded-full transition-all shadow-md shadow-sky-500/20">
                    Đăng ký
                </Link>
            </div>
        );
    }

    return (
        <div className="relative" ref={ref}>
            <button
                onClick={() => setOpen(!open)}
                className="flex items-center gap-2 px-3 py-1.5 rounded-full bg-white/5 border border-white/10 hover:bg-white/10 transition-all"
            >
                {user?.avatar_url ? (
                    <img src={user.avatar_url} alt={user.username} className="w-6 h-6 rounded-full object-cover" />
                ) : (
                    <div className="w-6 h-6 rounded-full bg-gradient-to-br from-purple-500 to-cyan-500 flex items-center justify-center text-white text-xs font-bold">
                        {user?.username?.charAt(0).toUpperCase() || 'U'}
                    </div>
                )}
                <span className="text-sm font-medium text-white hidden md:block">{user?.username}</span>
                <svg className={`w-3 h-3 text-slate-400 transition-transform ${open ? 'rotate-180' : ''}`} fill="none" stroke="currentColor" viewBox="0 0 24 24">
                    <path strokeLinecap="round" strokeLinejoin="round" strokeWidth={2} d="M19 9l-7 7-7-7" />
                </svg>
            </button>
            {open && (
                <div className="absolute right-0 top-full mt-2 w-48 bg-slate-900 border border-white/10 rounded-xl shadow-2xl overflow-hidden z-50">
                    <Link href="/profile/me" onClick={() => setOpen(false)} className="flex items-center gap-3 px-4 py-3 text-sm text-slate-300 hover:bg-white/5 hover:text-white transition-colors">
                        <svg className="w-4 h-4" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                            <path strokeLinecap="round" strokeLinejoin="round" strokeWidth={2} d="M16 7a4 4 0 11-8 0 4 4 0 018 0zM12 14a7 7 0 00-7 7h14a7 7 0 00-7-7z" />
                        </svg>
                        Hồ sơ của tôi
                    </Link>
                    <Link href="/feed" onClick={() => setOpen(false)} className="flex items-center gap-3 px-4 py-3 text-sm text-slate-300 hover:bg-white/5 hover:text-white transition-colors">
                        <svg className="w-4 h-4" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                            <path strokeLinecap="round" strokeLinejoin="round" strokeWidth={2} d="M19 20H5a2 2 0 01-2-2V6a2 2 0 012-2h10a2 2 0 012 2v1m2 13a2 2 0 01-2-2V7m2 13a2 2 0 002-2V9a2 2 0 00-2-2h-2m-4-3H9M7 16h6M7 8h6v4H7V8z" />
                        </svg>
                        Bảng tin
                    </Link>
                    <div className="border-t border-white/10" />
                    <button onClick={handleLogout} className="w-full flex items-center gap-3 px-4 py-3 text-sm text-rose-400 hover:bg-rose-500/10 transition-colors">
                        <svg className="w-4 h-4" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                            <path strokeLinecap="round" strokeLinejoin="round" strokeWidth={2} d="M17 16l4-4m0 0l-4-4m4 4H7m6 4v1a3 3 0 01-3 3H6a3 3 0 01-3-3V7a3 3 0 013-3h4a3 3 0 013 3v1" />
                        </svg>
                        Đăng xuất
                    </button>
                </div>
            )}
        </div>
    );
}

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
                    <Link href="/feed" className="text-sm font-medium text-gray-400 hover:text-cyan-400 cursor-pointer transition-colors">{t.nav.feed}</Link>
                    <Link href="/news" className="text-sm font-medium text-gray-400 hover:text-cyan-400 cursor-pointer transition-colors">{t.nav.news}</Link>
                    <Link href="/scanner" className="text-sm font-medium text-gray-400 hover:text-cyan-400 cursor-pointer transition-colors">{t.nav.scan}</Link>
                    <Link href="/analysis" className="text-sm font-medium text-gray-400 hover:text-cyan-400 cursor-pointer transition-colors">{t.nav.analysis}</Link>
                </nav>

                {/* Right Actions */}
                <div className="flex items-center gap-3">
                    <button
                        onClick={toggleLanguage}
                        className="flex items-center gap-2 px-3 py-1.5 rounded-full bg-white/5 border border-white/10 hover:bg-white/10 transition-all text-sm font-bold text-gray-300 hover:text-white"
                    >
                        <svg className="w-4 h-4" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                            <path strokeLinecap="round" strokeLinejoin="round" strokeWidth={2} d="M3 5h12M9 3v2m1.048 9.5A18.022 18.022 0 016.412 9m6.088 9h7M11 21l5-10 5 10M12.751 5C11.783 10.77 8.07 15.61 3 18.129" />
                        </svg>
                        <span>{language === 'vi' ? 'VN' : 'EN'}</span>
                    </button>
                    <AuthSection />
                </div>
            </div>
        </header>
    );
}

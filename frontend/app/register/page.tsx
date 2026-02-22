'use client';

import { useState } from 'react';
import { useRouter } from 'next/navigation';
import Link from 'next/link';
import { register, login } from '@/lib/api/auth';
import { getMyInfo } from '@/lib/api/users';
import { useAuthStore } from '@/lib/stores/authStore';
import { useLanguage } from '@/context/LanguageContext';
import { toast } from 'sonner';

export default function RegisterPage() {
    const router = useRouter();
    const { setToken, setUser } = useAuthStore();
    const { t, language, setLanguage } = useLanguage();
    const a = t.auth;

    const [form, setForm] = useState({ username: '', phone: '', password: '' });
    const [birthday, setBirthday] = useState({ day: '', month: '', year: '' });
    const [showPassword, setShowPassword] = useState(false);
    const [loading, setLoading] = useState(false);

    const handleChange = (e: React.ChangeEvent<HTMLInputElement>) => {
        setForm({ ...form, [e.target.name]: e.target.value });
    };

    const handleSubmit = async (e: React.FormEvent) => {
        e.preventDefault();
        if (!birthday.day || !birthday.month || !birthday.year) {
            toast.error(language === 'vi' ? 'Vui lòng chọn ngày sinh' : 'Please select your date of birth');
            return;
        }
        setLoading(true);
        try {
            const birthdayISO = new Date(`${birthday.year}-${birthday.month.padStart(2, '0')}-${birthday.day.padStart(2, '0')}`).toISOString();
            await register({ username: form.username, phone: form.phone, password: form.password, birthday: birthdayISO });
            const data = await login(form.phone, form.password);
            setToken(data.token);
            const user = await getMyInfo();
            setUser(user);
            toast.success(language === 'vi' ? 'Đăng ký thành công!' : 'Account created successfully!');
            router.push('/feed');
        } catch (err: any) {
            toast.error(err?.response?.data?.message || (language === 'vi' ? 'Đăng ký thất bại' : 'Registration failed'));
        } finally {
            setLoading(false);
        }
    };

    const steps = language === 'vi'
        ? ['Bảo vệ khỏi Rug Pull', 'Phân tích AI trong 30s', 'Cảnh báo real-time', 'Cộng đồng nhà đầu tư']
        : ['Protect against Rug Pulls', 'AI analysis in 30s', 'Real-time alerts', 'Investor community'];

    return (
        <div className="min-h-screen bg-black flex">
            {/* Left decorative panel */}
            <div className="hidden lg:flex lg:w-1/2 xl:w-[55%] bg-zinc-950 items-center justify-center relative overflow-hidden">
                <div className="absolute inset-0 bg-gradient-to-br from-sky-900/20 via-transparent to-transparent" />
                <div className="absolute top-1/4 right-1/4 w-96 h-96 bg-sky-500/10 rounded-full blur-3xl" />
                <div className="absolute bottom-1/4 left-1/4 w-64 h-64 bg-blue-500/10 rounded-full blur-3xl" />

                <div className="relative z-10 max-w-md px-8">
                    <div className="inline-flex items-center gap-3 mb-10">
                        <div className="w-10 h-10 rounded-xl bg-sky-500 flex items-center justify-center shadow-lg shadow-sky-500/30">
                            <svg className="w-6 h-6 text-white" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                                <path strokeLinecap="round" strokeLinejoin="round" strokeWidth={2} d="M9 12l2 2 4-4m5.618-4.016A11.955 11.955 0 0112 2.944a11.955 11.955 0 01-8.618 3.04A12.02 12.02 0 003 9c0 5.591 3.824 10.29 9 11.622 5.176-1.332 9-6.03 9-11.622 0-1.042-.133-2.052-.382-3.016z" />
                            </svg>
                        </div>
                        <span className="text-2xl font-black text-white">CryptoCheck</span>
                    </div>

                    <h2 className="text-4xl font-bold text-white mb-4 leading-tight">
                        {language === 'vi' ? 'Tham gia cộng đồng.' : 'Join the community.'}<br />
                        <span className="text-sky-400">{language === 'vi' ? 'Đầu tư thông minh.' : 'Invest smarter.'}</span>
                    </h2>
                    <p className="text-zinc-500 text-base mb-10 leading-relaxed">
                        {language === 'vi'
                            ? 'Hàng nghìn nhà đầu tư đang sử dụng CryptoCheck để bảo vệ tài sản của họ.'
                            : 'Thousands of investors use CryptoCheck to protect their assets.'}
                    </p>

                    <div className="space-y-3">
                        {steps.map((step, i) => (
                            <div key={i} className="flex items-center gap-3">
                                <div className="w-6 h-6 rounded-full bg-sky-500/20 border border-sky-500/30 flex items-center justify-center flex-shrink-0">
                                    <svg className="w-3 h-3 text-sky-400" fill="currentColor" viewBox="0 0 20 20">
                                        <path fillRule="evenodd" d="M16.707 5.293a1 1 0 010 1.414l-8 8a1 1 0 01-1.414 0l-4-4a1 1 0 011.414-1.414L8 12.586l7.293-7.293a1 1 0 011.414 0z" clipRule="evenodd" />
                                    </svg>
                                </div>
                                <span className="text-zinc-300 text-sm">{step}</span>
                            </div>
                        ))}
                    </div>
                </div>
            </div>

            {/* Right — form */}
            <div className="flex-1 flex items-center justify-center p-6 lg:p-12 relative">
                <div className="w-full max-w-sm">
                    {/* Mobile logo */}
                    <div className="flex items-center gap-2 mb-8 lg:hidden">
                        <div className="w-8 h-8 rounded-xl bg-sky-500 flex items-center justify-center">
                            <svg className="w-5 h-5 text-white" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                                <path strokeLinecap="round" strokeLinejoin="round" strokeWidth={2} d="M9 12l2 2 4-4m5.618-4.016A11.955 11.955 0 0112 2.944a11.955 11.955 0 01-8.618 3.04A12.02 12.02 0 003 9c0 5.591 3.824 10.29 9 11.622 5.176-1.332 9-6.03 9-11.622 0-1.042-.133-2.052-.382-3.016z" />
                            </svg>
                        </div>
                        <span className="text-xl font-bold text-white">CryptoCheck</span>
                    </div>

                    <h1 className="text-3xl font-bold text-white mb-1">{a.register_title}</h1>
                    <p className="text-zinc-500 mb-8">{a.register_subtitle}</p>

                    {/* Google OAuth — placeholder */}
                    <button
                        disabled
                        title={a.google_coming_soon}
                        className="w-full flex items-center justify-center gap-3 px-4 py-3 bg-zinc-900 border border-zinc-800 rounded-xl text-zinc-400 cursor-not-allowed relative group mb-6"
                    >
                        <svg className="w-5 h-5 flex-shrink-0" viewBox="0 0 24 24">
                            <path d="M22.56 12.25c0-.78-.07-1.53-.2-2.25H12v4.26h5.92c-.26 1.37-1.04 2.53-2.21 3.31v2.77h3.57c2.08-1.92 3.28-4.74 3.28-8.09z" fill="#4285F4" />
                            <path d="M12 23c2.97 0 5.46-.98 7.28-2.66l-3.57-2.77c-.98.66-2.23 1.06-3.71 1.06-2.86 0-5.29-1.93-6.16-4.53H2.18v2.84C3.99 20.53 7.7 23 12 23z" fill="#34A853" />
                            <path d="M5.84 14.09c-.22-.66-.35-1.36-.35-2.09s.13-1.43.35-2.09V7.07H2.18C1.43 8.55 1 10.22 1 12s.43 3.45 1.18 4.93l2.85-2.22.81-.62z" fill="#FBBC05" />
                            <path d="M12 5.38c1.62 0 3.06.56 4.21 1.64l3.15-3.15C17.45 2.09 14.97 1 12 1 7.7 1 3.99 3.47 2.18 7.07l3.66 2.84c.87-2.6 3.3-4.53 6.16-4.53z" fill="#EA4335" />
                        </svg>
                        <span className="font-medium">{a.google_btn}</span>
                        <span className="absolute right-3 text-xs px-2 py-0.5 bg-sky-500/10 text-sky-400 border border-sky-500/20 rounded-full opacity-0 group-hover:opacity-100 transition-opacity">
                            {a.google_coming_soon}
                        </span>
                    </button>

                    {/* Divider */}
                    <div className="flex items-center gap-3 mb-6">
                        <div className="flex-1 h-px bg-zinc-800" />
                        <span className="text-xs text-zinc-600">{a.or_continue}</span>
                        <div className="flex-1 h-px bg-zinc-800" />
                    </div>

                    {/* Form */}
                    <form onSubmit={handleSubmit} className="space-y-4">
                        <div>
                            <label className="block text-sm font-medium text-zinc-300 mb-2">{a.username}</label>
                            <input name="username" type="text" value={form.username} onChange={handleChange}
                                placeholder={a.username_placeholder} required minLength={3}
                                className="w-full bg-zinc-900 border border-zinc-800 rounded-xl px-4 py-3 text-white placeholder-zinc-600 focus:outline-none focus:ring-2 focus:ring-sky-500 focus:border-transparent transition-all text-sm" />
                        </div>
                        <div>
                            <label className="block text-sm font-medium text-zinc-300 mb-2">{a.phone}</label>
                            <input name="phone" type="tel" value={form.phone} onChange={handleChange}
                                placeholder={a.phone_placeholder} required
                                className="w-full bg-zinc-900 border border-zinc-800 rounded-xl px-4 py-3 text-white placeholder-zinc-600 focus:outline-none focus:ring-2 focus:ring-sky-500 focus:border-transparent transition-all text-sm" />
                        </div>
                        <div>
                            <label className="block text-sm font-medium text-zinc-300 mb-2">{a.password}</label>
                            <div className="relative">
                                <input name="password" type={showPassword ? 'text' : 'password'} value={form.password} onChange={handleChange}
                                    placeholder={a.password_placeholder} required minLength={6}
                                    className="w-full bg-zinc-900 border border-zinc-800 rounded-xl px-4 py-3 pr-12 text-white placeholder-zinc-600 focus:outline-none focus:ring-2 focus:ring-sky-500 focus:border-transparent transition-all text-sm" />
                                <button type="button" onClick={() => setShowPassword(!showPassword)}
                                    className="absolute right-3 top-1/2 -translate-y-1/2 text-zinc-500 hover:text-zinc-300 transition-colors p-1">
                                    {showPassword ? (
                                        <svg className="w-4 h-4" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                                            <path strokeLinecap="round" strokeLinejoin="round" strokeWidth={2} d="M13.875 18.825A10.05 10.05 0 0112 19c-4.478 0-8.268-2.943-9.543-7a9.97 9.97 0 011.563-3.029m5.858.908a3 3 0 114.243 4.243M9.878 9.878l4.242 4.242M9.88 9.88l-3.29-3.29m7.532 7.532l3.29 3.29M3 3l3.59 3.59m0 0A9.953 9.953 0 0112 5c4.478 0 8.268 2.943 9.543 7a10.025 10.025 0 01-4.132 5.411m0 0L21 21" />
                                        </svg>
                                    ) : (
                                        <svg className="w-4 h-4" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                                            <path strokeLinecap="round" strokeLinejoin="round" strokeWidth={2} d="M15 12a3 3 0 11-6 0 3 3 0 016 0z" />
                                            <path strokeLinecap="round" strokeLinejoin="round" strokeWidth={2} d="M2.458 12C3.732 7.943 7.523 5 12 5c4.478 0 8.268 2.943 9.542 7-1.274 4.057-5.064 7-9.542 7-4.477 0-8.268-2.943-9.542-7z" />
                                        </svg>
                                    )}
                                </button>
                            </div>
                        </div>
                        <div>
                            <label className="block text-sm font-medium text-zinc-300 mb-2">{a.birthday}</label>
                            <div className="grid grid-cols-3 gap-2">
                                {/* Day */}
                                <div className="relative">
                                    <select
                                        value={birthday.day}
                                        onChange={(e) => setBirthday({ ...birthday, day: e.target.value })}
                                        className="w-full appearance-none bg-zinc-900 border border-zinc-800 rounded-xl px-3 py-3 text-sm text-white focus:outline-none focus:ring-2 focus:ring-sky-500 focus:border-transparent transition-all pr-8 cursor-pointer"
                                    >
                                        <option value="" disabled>{language === 'vi' ? 'Ngày' : 'Day'}</option>
                                        {Array.from({ length: 31 }, (_, i) => i + 1).map(d => (
                                            <option key={d} value={String(d)}>{d}</option>
                                        ))}
                                    </select>
                                    <svg className="pointer-events-none absolute right-2.5 top-1/2 -translate-y-1/2 w-4 h-4 text-zinc-500" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                                        <path strokeLinecap="round" strokeLinejoin="round" strokeWidth={2} d="M19 9l-7 7-7-7" />
                                    </svg>
                                </div>
                                {/* Month */}
                                <div className="relative">
                                    <select
                                        value={birthday.month}
                                        onChange={(e) => setBirthday({ ...birthday, month: e.target.value })}
                                        className="w-full appearance-none bg-zinc-900 border border-zinc-800 rounded-xl px-3 py-3 text-sm text-white focus:outline-none focus:ring-2 focus:ring-sky-500 focus:border-transparent transition-all pr-8 cursor-pointer"
                                    >
                                        <option value="" disabled>{language === 'vi' ? 'Tháng' : 'Month'}</option>
                                        {(language === 'vi'
                                            ? ['Th.1', 'Th.2', 'Th.3', 'Th.4', 'Th.5', 'Th.6', 'Th.7', 'Th.8', 'Th.9', 'Th.10', 'Th.11', 'Th.12']
                                            : ['Jan', 'Feb', 'Mar', 'Apr', 'May', 'Jun', 'Jul', 'Aug', 'Sep', 'Oct', 'Nov', 'Dec']
                                        ).map((m, i) => (
                                            <option key={i + 1} value={String(i + 1)}>{m}</option>
                                        ))}
                                    </select>
                                    <svg className="pointer-events-none absolute right-2.5 top-1/2 -translate-y-1/2 w-4 h-4 text-zinc-500" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                                        <path strokeLinecap="round" strokeLinejoin="round" strokeWidth={2} d="M19 9l-7 7-7-7" />
                                    </svg>
                                </div>
                                {/* Year */}
                                <div className="relative">
                                    <select
                                        value={birthday.year}
                                        onChange={(e) => setBirthday({ ...birthday, year: e.target.value })}
                                        className="w-full appearance-none bg-zinc-900 border border-zinc-800 rounded-xl px-3 py-3 text-sm text-white focus:outline-none focus:ring-2 focus:ring-sky-500 focus:border-transparent transition-all pr-8 cursor-pointer"
                                    >
                                        <option value="" disabled>{language === 'vi' ? 'Năm' : 'Year'}</option>
                                        {Array.from({ length: 100 }, (_, i) => new Date().getFullYear() - 10 - i).map(y => (
                                            <option key={y} value={String(y)}>{y}</option>
                                        ))}
                                    </select>
                                    <svg className="pointer-events-none absolute right-2.5 top-1/2 -translate-y-1/2 w-4 h-4 text-zinc-500" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                                        <path strokeLinecap="round" strokeLinejoin="round" strokeWidth={2} d="M19 9l-7 7-7-7" />
                                    </svg>
                                </div>
                            </div>
                        </div>
                        <button type="submit" disabled={loading}
                            className="w-full bg-sky-500 hover:bg-sky-400 disabled:opacity-50 disabled:cursor-not-allowed text-white font-semibold py-3 rounded-xl transition-all duration-200 shadow-lg shadow-sky-500/20 flex items-center justify-center gap-2 text-sm mt-2">
                            {loading ? (
                                <><svg className="animate-spin w-4 h-4" fill="none" viewBox="0 0 24 24"><circle className="opacity-25" cx="12" cy="12" r="10" stroke="currentColor" strokeWidth="4" /><path className="opacity-75" fill="currentColor" d="M4 12a8 8 0 018-8V0C5.373 0 0 5.373 0 12h4z" /></svg>{a.registering}</>
                            ) : a.register_btn}
                        </button>
                    </form>

                    <p className="mt-6 text-center text-sm text-zinc-500">
                        {a.has_account}{' '}
                        <Link href="/login" className="text-sky-400 hover:text-sky-300 font-medium transition-colors">{a.login_link}</Link>
                    </p>
                </div>
            </div>
        </div>
    );
}

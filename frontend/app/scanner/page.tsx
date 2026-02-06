"use client";

import { useState, useEffect } from 'react';
import { ScanResult, scanToken } from '@/lib/scannerApi';
import { Search, ShieldCheck, ShieldAlert, Cpu, Check, Loader2 } from 'lucide-react';
import { useLanguage } from '@/context/LanguageContext';

export default function ScannerPage() {
    const { t } = useLanguage();
    const [query, setQuery] = useState('');
    const [result, setResult] = useState<ScanResult | null>(null);
    const [loading, setLoading] = useState(false);
    const [currentStep, setCurrentStep] = useState(0);
    const [error, setError] = useState('');

    const loadingSteps = t.scanner.steps;

    useEffect(() => {
        let interval: NodeJS.Timeout;
        if (loading) {
            setCurrentStep(0);
            interval = setInterval(() => {
                setCurrentStep((prev) => {
                    if (prev < loadingSteps.length - 1) return prev + 1;
                    return prev;
                });
            }, 800); // Advance step every 800ms
        }
        return () => clearInterval(interval);
    }, [loading]);

    const handleScan = async () => {
        if (!query) return;
        setLoading(true);
        setError('');
        setResult(null);

        // Minimum loading time: 4 steps * 800ms = 3200ms
        const minDurationPromise = new Promise(resolve => setTimeout(resolve, 3200));
        const scanScanPromise = scanToken(query);

        try {
            // Wait for both the minimum animation time and the API result
            const [_, data] = await Promise.all([minDurationPromise, scanScanPromise]);
            setResult(data);
        } catch (err: any) {
            const errorMsg = err.response?.data?.message || err.message || t.scanner_page.error_generic;
            setError(errorMsg);
        } finally {
            setLoading(false);
        }
    };

    return (
        <div className="min-h-screen bg-slate-950 text-white p-8 font-sans">
            <div className="max-w-4xl mx-auto">
                <header className="mb-12 text-center">
                    <h1 className="text-4xl font-bold bg-gradient-to-r from-blue-400 to-purple-500 bg-clip-text text-transparent">
                        {t.scanner_page.title}
                    </h1>
                    <p className="text-slate-400 mt-2">{t.scanner_page.subtitle}</p>
                </header>

                {/* Search Bar */}
                <div className="flex gap-2 mb-10">
                    <input
                        type="text"
                        placeholder={t.scanner_page.placeholder}
                        className="flex-1 bg-slate-900 border border-slate-700 rounded-lg px-4 py-3 focus:outline-none focus:border-blue-500 transition-colors"
                        value={query}
                        onChange={(e) => setQuery(e.target.value)}
                        onKeyDown={(e) => e.key === 'Enter' && handleScan()}
                    />
                    <button
                        onClick={handleScan}
                        disabled={loading}
                        className="bg-blue-600 hover:bg-blue-700 disabled:opacity-50 px-6 py-3 rounded-lg font-medium flex items-center gap-2 transition-colors"
                    >
                        {loading ? <Cpu className="animate-spin" /> : <Search />}
                        {loading ? t.scanner_page.analyzing_btn : t.scanner_page.scan_btn}
                    </button>
                </div>

                {/* Loading Indicator */}
                {loading && (
                    <div className="flex flex-col items-center justify-center p-8 animate-in fade-in duration-300">
                        <div className="relative w-20 h-20 mb-6">
                            <div className="absolute inset-0 border-4 border-blue-500/20 rounded-full"></div>
                            <div className="absolute inset-0 border-4 border-blue-500 rounded-full border-t-transparent animate-spin"></div>
                            <ShieldCheck className="absolute inset-0 m-auto text-blue-400 w-8 h-8 animate-pulse" />
                        </div>

                        <h3 className="text-xl font-bold text-blue-400 mb-6 animate-pulse">{t.scanner.analyzing}</h3>

                        <div className="flex flex-col gap-3 w-full max-w-sm">
                            {loadingSteps.map((step, index) => {
                                const isCompleted = index < currentStep;
                                const isCurrent = index === currentStep;

                                return (
                                    <div
                                        key={index}
                                        className={`flex items-center gap-3 transition-opacity duration-300 ${index > currentStep ? 'opacity-30' : 'opacity-100'
                                            }`}
                                    >
                                        <div className="w-5 h-5 flex items-center justify-center">
                                            {isCompleted ? (
                                                <div className="w-5 h-5 bg-green-500 rounded-full flex items-center justify-center">
                                                    <Check className="w-3 h-3 text-white" />
                                                </div>
                                            ) : isCurrent ? (
                                                <Loader2 className="w-5 h-5 text-blue-400 animate-spin" />
                                            ) : (
                                                <div className="w-2 h-2 bg-slate-700 rounded-full" />
                                            )}
                                        </div>
                                        <span className={`text-sm ${isCompleted ? 'text-green-400' : isCurrent ? 'text-blue-300 font-medium' : 'text-slate-500'
                                            }`}>
                                            {step}
                                        </span>
                                    </div>
                                );
                            })}
                        </div>
                    </div>
                )}

                {/* Error Message */}
                {error && (
                    <div className="bg-red-900/50 border border-red-700 text-red-200 p-4 rounded-lg mb-8 text-center">
                        {error}
                    </div>
                )}

                {/* Report Card */}
                {result && (
                    <div className="bg-slate-900 border border-slate-800 rounded-xl overflow-hidden shadow-2xl animate-in fade-in slide-in-from-bottom-4 duration-500">
                        {/* Header */}
                        <div className="p-6 border-b border-slate-800 flex justify-between items-center bg-slate-900/50">
                            <div>
                                <h2 className="text-2xl font-bold flex items-center gap-2">
                                    {result.name}
                                    <span className="text-xs bg-slate-800 px-2 py-1 rounded text-slate-400 uppercase">{result.network}</span>
                                </h2>
                                <p className="text-slate-500 text-sm font-mono mt-1">{result.address}</p>
                            </div>
                            <div className="text-center">
                                <div className={`text-4xl font-black ${getScoreColor(result.trust_score)}`}>
                                    {result.trust_score}/100
                                </div>
                                <div className="text-xs text-slate-500 uppercase tracking-widest">{t.scanner_page.trust_score}</div>
                            </div>
                        </div>

                        <div className="grid md:grid-cols-2 gap-0">
                            {/* Risk Column */}
                            <div className="p-6 border-r border-slate-800 bg-slate-900/30">
                                <h3 className="text-red-400 font-bold mb-4 flex items-center gap-2">
                                    <ShieldAlert className="w-5 h-5" /> {t.scanner_page.risk_analysis}
                                </h3>
                                {(!result.issues || result.issues.length === 0) ? (
                                    <p className="text-slate-500 italic">{t.scanner_page.no_risks}</p>
                                ) : (
                                    <ul className="space-y-4">
                                        {result.issues.map((issue, idx) => {
                                            // Helper to translate issue name
                                            let translatedName = issue.name;

                                            // Map common English issue names to translation keys
                                            if (issue.name.includes("High Tax")) translatedName = t.security_issues.high_tax || issue.name;
                                            else if (issue.name.includes("Honeypot")) translatedName = t.security_issues.honeypot || issue.name;
                                            else if (issue.name.includes("Hidden Ownership")) translatedName = t.security_issues.hidden_ownership || issue.name;
                                            else if (issue.name.includes("Fee Modification")) translatedName = t.security_issues.fee_modification || issue.name;
                                            else if (issue.name.includes("Rug Pull")) translatedName = t.security_issues.rug_pull || issue.name;
                                            else if (issue.name.includes("Blacklist")) translatedName = t.security_issues.blacklist || issue.name;
                                            else if (issue.name.includes("Hidden Mint")) translatedName = t.security_issues.hidden_mint || issue.name;

                                            // Helper to translate description
                                            let translatedDesc = issue.description;
                                            if (issue.description === "Detected via Pattern Matching (AI Unavailable)") {
                                                translatedDesc = t.security_issues.pattern_matching || issue.description;
                                            }

                                            return (
                                                <li key={idx} className="bg-red-950/30 p-3 rounded border border-red-900/30">
                                                    <div className="flex justify-between items-start">
                                                        <span className="font-semibold text-red-200">{translatedName}</span>
                                                        <span className="text-xs bg-red-900/50 text-red-300 px-1.5 py-0.5 rounded">-{issue.impact}</span>
                                                    </div>
                                                    <p className="text-sm text-red-300/70 mt-1">{translatedDesc}</p>
                                                </li>
                                            );
                                        })}
                                    </ul>
                                )}
                            </div>

                            {/* Safe Column */}
                            <div className="p-6 bg-slate-900/30">
                                <h3 className="text-green-400 font-bold mb-4 flex items-center gap-2">
                                    <ShieldCheck className="w-5 h-5" /> {t.scanner_page.safety_features}
                                </h3>
                                {(!result.safe_features || result.safe_features.length === 0) ? (
                                    <p className="text-slate-500 italic">{t.scanner_page.no_features}</p>
                                ) : (
                                    <ul className="space-y-2">
                                        {result.safe_features.map((feat, idx) => {
                                            let translatedFeat = feat;
                                            if (feat === "Ownable Pattern") translatedFeat = t.security_issues.ownable || feat;
                                            else if (feat === "Standard Interface") translatedFeat = t.security_issues.standard_interface || feat;
                                            else if (feat === "Verified Contract") translatedFeat = t.security_issues.verified_contract || feat;

                                            return (
                                                <li key={idx} className="flex items-center gap-2 text-green-200/80 text-sm">
                                                    <span className="w-1.5 h-1.5 bg-green-500 rounded-full"></span>
                                                    {translatedFeat}
                                                </li>
                                            );
                                        })}
                                    </ul>
                                )}
                            </div>
                        </div>
                    </div>
                )}
            </div>
        </div>
    );
}

function getScoreColor(score: number): string {
    if (score >= 80) return 'text-green-500';
    if (score >= 50) return 'text-yellow-500';
    return 'text-red-500';
}

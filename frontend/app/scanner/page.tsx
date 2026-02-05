"use client";

import { useState } from 'react';
import { ScanResult, scanToken } from '@/lib/scannerApi';
import { Search, ShieldCheck, ShieldAlert, Cpu } from 'lucide-react';

export default function ScannerPage() {
    const [query, setQuery] = useState('');
    const [result, setResult] = useState<ScanResult | null>(null);
    const [loading, setLoading] = useState(false);
    const [error, setError] = useState('');

    const handleScan = async () => {
        if (!query) return;
        setLoading(true);
        setError('');
        setResult(null);

        try {
            const data = await scanToken(query);
            setResult(data);
        } catch (err) {
            setError('Scan failed. Token not found or API error.');
        } finally {
            setLoading(false);
        }
    };

    return (
        <div className="min-h-screen bg-slate-950 text-white p-8 font-sans">
            <div className="max-w-4xl mx-auto">
                <header className="mb-12 text-center">
                    <h1 className="text-4xl font-bold bg-gradient-to-r from-blue-400 to-purple-500 bg-clip-text text-transparent">
                        ChainGuardian Scanner
                    </h1>
                    <p className="text-slate-400 mt-2">AI-Powered Smart Contract Auditor</p>
                </header>

                {/* Search Bar */}
                <div className="flex gap-2 mb-10">
                    <input
                        type="text"
                        placeholder="Enter Token Symbol (e.g. ENA) or Contract Address"
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
                        {loading ? 'Analyzing...' : 'Scan Now'}
                    </button>
                </div>

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
                                <div className="text-xs text-slate-500 uppercase tracking-widest">Trust Score</div>
                            </div>
                        </div>

                        <div className="grid md:grid-cols-2 gap-0">
                            {/* Risk Column */}
                            <div className="p-6 border-r border-slate-800 bg-slate-900/30">
                                <h3 className="text-red-400 font-bold mb-4 flex items-center gap-2">
                                    <ShieldAlert className="w-5 h-5" /> Risk Analysis
                                </h3>
                                {result.issues.length === 0 ? (
                                    <p className="text-slate-500 italic">No critical risks found.</p>
                                ) : (
                                    <ul className="space-y-4">
                                        {result.issues.map((issue, idx) => (
                                            <li key={idx} className="bg-red-950/30 p-3 rounded border border-red-900/30">
                                                <div className="flex justify-between items-start">
                                                    <span className="font-semibold text-red-200">{issue.Name}</span>
                                                    <span className="text-xs bg-red-900/50 text-red-300 px-1.5 py-0.5 rounded">-{issue.Impact}</span>
                                                </div>
                                                <p className="text-sm text-red-300/70 mt-1">{issue.Description}</p>
                                            </li>
                                        ))}
                                    </ul>
                                )}
                            </div>

                            {/* Safe Column */}
                            <div className="p-6 bg-slate-900/30">
                                <h3 className="text-green-400 font-bold mb-4 flex items-center gap-2">
                                    <ShieldCheck className="w-5 h-5" /> Safety Features
                                </h3>
                                {result.safe_features.length === 0 ? (
                                    <p className="text-slate-500 italic">No specific safety features detected.</p>
                                ) : (
                                    <ul className="space-y-2">
                                        {result.safe_features.map((feat, idx) => (
                                            <li key={idx} className="flex items-center gap-2 text-green-200/80 text-sm">
                                                <span className="w-1.5 h-1.5 bg-green-500 rounded-full"></span>
                                                {feat}
                                            </li>
                                        ))}
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

"use client";

import { Search, TrendingUp, Shield } from "lucide-react";
import { useState } from "react";
import Link from "next/link";
import { useLanguage } from "@/context/LanguageContext";

// Mock trending tokens data
const trendingTokens = [
    { name: "Arbitrum", symbol: "ARB", network: "ARB", address: "0x912ce59144191c1204e64559fe8253a0e49e6548", trustScore: 85, searches: 1250 },
    { name: "Polygon", symbol: "MATIC", network: "POLY", address: "0x7d1afa7b718fb893db30a3abc0cfc608aacfebb0", trustScore: 100, searches: 1100 },
    { name: "Tether USD", symbol: "USDT", network: "ETH", address: "0xdac17f958d2ee523a2206206994597c13d831ec7", trustScore: 85, searches: 980 },
    { name: "BNB", symbol: "BNB", network: "BSC", address: "0xbb4cdb9cbd36b01bd1cbaebf2de08d9173bc095c", trustScore: 95, searches: 850 },
    { name: "USD Coin", symbol: "USDC", network: "ETH", address: "0xa0b86991c6218b36c1d19d4a2e9eb0ce3606eb48", trustScore: 95, searches: 820 },
    { name: "Pepe", symbol: "PEPE", network: "ETH", address: "0x6982508145454ce325ddbe47a25d4ec3d2311933", trustScore: 45, searches: 750 },
    { name: "Shiba Inu", symbol: "SHIB", network: "ETH", address: "0x95ad61b0a150d79219dcf64e1e6cc01f0b64c4ce", trustScore: 100, searches: 680 },
    { name: "Chainlink", symbol: "LINK", network: "ETH", address: "0x514910771af9ca656af840dff83e8264ecf986ca", trustScore: 85, searches: 620 },
    { name: "Uniswap", symbol: "UNI", network: "ETH", address: "0x1f9840a85d5af5bf1d1762f925bdaddc4201f984", trustScore: 95, searches: 580 },
    { name: "Optimism", symbol: "OP", network: "OP", address: "0x4200000000000000000000000000000000000042", trustScore: 90, searches: 520 },
];

export default function TokenSearch() {
    const { t } = useLanguage();
    const [searchQuery, setSearchQuery] = useState("");
    const [filteredTokens, setFilteredTokens] = useState(trendingTokens);

    const handleSearch = (query: string) => {
        setSearchQuery(query);
        if (query.trim() === "") {
            setFilteredTokens(trendingTokens);
        } else {
            const filtered = trendingTokens.filter(token =>
                token.name.toLowerCase().includes(query.toLowerCase()) ||
                token.symbol.toLowerCase().includes(query.toLowerCase()) ||
                token.address.toLowerCase().includes(query.toLowerCase())
            );
            setFilteredTokens(filtered);
        }
    };

    const getTrustScoreColor = (score: number) => {
        if (score >= 80) return "text-green-400 bg-green-500/10";
        if (score >= 50) return "text-yellow-400 bg-yellow-500/10";
        return "text-red-400 bg-red-500/10";
    };

    const getTrustScoreLabel = (score: number) => {
        if (score >= 80) return t.search.safe;
        if (score >= 50) return t.search.warning;
        return t.search.risk;
    };

    return (
        <section id="search" className="py-20 px-4">
            <div className="max-w-6xl mx-auto">
                {/* Header */}
                <div className="text-center mb-12">
                    <h2 className="text-4xl md:text-5xl font-black text-white mb-4">
                        {t.search.title}
                    </h2>
                    <p className="text-gray-400 text-lg">
                        {t.search.subtitle} ({t.search.hint})
                    </p>
                </div>

                {/* Search Box */}
                <div className="mb-12">
                    <div className="relative max-w-2xl mx-auto">
                        <Search className="absolute left-4 top-1/2 -translate-y-1/2 w-5 h-5 text-gray-400" />
                        <input
                            type="text"
                            placeholder={t.search.placeholder}
                            value={searchQuery}
                            onChange={(e) => handleSearch(e.target.value)}
                            className="w-full pl-12 pr-4 py-4 bg-white/5 border border-white/10 rounded-xl text-white placeholder-gray-500 focus:outline-none focus:border-cyan-500/50 focus:bg-white/10 transition-all"
                        />
                    </div>
                </div>

                {/* Trending Tokens Table */}
                <div className="bg-gradient-to-br from-white/5 to-white/[0.02] backdrop-blur-xl border border-white/10 rounded-3xl p-6 md:p-8">
                    <div className="flex items-center gap-2 mb-6">
                        <TrendingUp className="w-5 h-5 text-cyan-400" />
                        <h3 className="text-xl font-bold text-white">
                            {searchQuery ? `${t.search.results_title} (${filteredTokens.length})` : t.search.trending_title}
                        </h3>
                    </div>

                    {filteredTokens.length === 0 ? (
                        <div className="text-center py-12">
                            <p className="text-gray-400">{t.search.no_results}</p>
                        </div>
                    ) : (
                        <div className="grid gap-3">
                            {filteredTokens.map((token, index) => (
                                <Link
                                    key={index}
                                    href={`/scanner?token=${token.symbol}`}
                                    className="group flex items-center gap-4 p-4 bg-white/5 rounded-xl border border-white/10 hover:bg-white/10 hover:border-cyan-500/30 transition-all cursor-pointer"
                                >
                                    {/* Rank */}
                                    <div className="w-8 h-8 rounded-lg bg-gradient-to-br from-cyan-500/20 to-blue-500/20 flex items-center justify-center flex-shrink-0">
                                        <span className="text-sm font-bold text-cyan-400">#{index + 1}</span>
                                    </div>

                                    {/* Token Info */}
                                    <div className="flex-1 min-w-0">
                                        <div className="flex items-center gap-2 mb-1">
                                            <span className="font-bold text-white">{token.name}</span>
                                            <span className="text-gray-500">({token.symbol})</span>
                                        </div>
                                        <div className="text-xs text-gray-500 truncate">
                                            {token.network} • {token.address.slice(0, 6)}...{token.address.slice(-4)}
                                        </div>
                                    </div>

                                    {/* Trust Score Badge */}
                                    <div className={`px-3 py-1.5 rounded-lg font-semibold text-sm ${getTrustScoreColor(token.trustScore)}`}>
                                        {token.trustScore} • {getTrustScoreLabel(token.trustScore)}
                                    </div>

                                    {/* Scan Icon */}
                                    <Shield className="w-5 h-5 text-gray-400 group-hover:text-cyan-400 transition-colors flex-shrink-0" />
                                </Link>
                            ))}
                        </div>
                    )}
                </div>

                {/* Info Note */}
                <div className="mt-6 text-center text-sm text-gray-500">
                    {t.search.scan_now_note}
                </div>
            </div>
        </section>
    );
}

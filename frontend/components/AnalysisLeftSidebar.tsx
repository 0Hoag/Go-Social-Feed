"use client";

import { useEffect, useState } from 'react';
import { TrendingUp, TrendingDown, Activity, BarChart3 } from 'lucide-react';
import axios from 'axios';
import { useLanguage } from '@/context/LanguageContext';

interface MarketData {
    symbol: string;
    name: string;
    price: number;
    change: number;
}

export default function AnalysisLeftSidebar() {
    const { t } = useLanguage();
    const [topGainers, setTopGainers] = useState<MarketData[]>([]);
    const [topLosers, setTopLosers] = useState<MarketData[]>([]);
    const [marketCap, setMarketCap] = useState('$0');
    const [volume24h, setVolume24h] = useState('$0');
    const [btcDominance, setBtcDominance] = useState('0%');
    const [activeCoins, setActiveCoins] = useState('0');

    useEffect(() => {
        const fetchTopMovers = async () => {
            try {
                const res = await axios.get('https://api.binance.com/api/v3/ticker/24hr');
                const data = res.data;

                // Filter USDT pairs only
                const usdtPairs = data.filter((t: any) => t.symbol.endsWith('USDT'));

                // Calculate total market metrics
                let totalVolume = 0;
                usdtPairs.forEach((t: any) => {
                    totalVolume += parseFloat(t.quoteVolume);
                });

                // Get BTC data for dominance calculation
                const btcData = data.find((t: any) => t.symbol === 'BTCUSDT');
                if (btcData) {
                    const btcMarketCap = parseFloat(btcData.lastPrice) * 19500000; // Approximate BTC supply
                    const totalMarketCap = btcMarketCap / 0.534; // Assuming ~53.4% dominance
                    setMarketCap(`$${(totalMarketCap / 1e12).toFixed(1)}T`);
                    setBtcDominance(`${((btcMarketCap / totalMarketCap) * 100).toFixed(1)}%`);
                }

                setVolume24h(`$${(totalVolume / 1e9).toFixed(1)}B`);
                setActiveCoins(usdtPairs.length.toLocaleString());

                // Sort by price change percentage
                const sortedByChange = [...usdtPairs].sort((a: any, b: any) =>
                    parseFloat(b.priceChangePercent) - parseFloat(a.priceChangePercent)
                );

                // Top 3 gainers
                const gainers = sortedByChange.slice(0, 3).map((t: any) => ({
                    symbol: t.symbol.replace('USDT', ''),
                    name: t.symbol.replace('USDT', ''),
                    price: parseFloat(t.lastPrice),
                    change: parseFloat(t.priceChangePercent)
                }));

                // Top 3 losers (reverse sort)
                const losers = sortedByChange.slice(-3).reverse().map((t: any) => ({
                    symbol: t.symbol.replace('USDT', ''),
                    name: t.symbol.replace('USDT', ''),
                    price: parseFloat(t.lastPrice),
                    change: parseFloat(t.priceChangePercent)
                }));

                setTopGainers(gainers);
                setTopLosers(losers);
            } catch (error) {
                console.error('Failed to fetch top movers:', error);
            }
        };

        fetchTopMovers();
        const interval = setInterval(fetchTopMovers, 30000); // Update every 30s

        return () => clearInterval(interval);
    }, []);

    return (
        <div className="space-y-2.5">
            {/* Market Overview */}
            <div className="bg-black/40 border border-white/5 rounded-xl p-3 hover:border-blue-500/20 transition-all duration-300">
                <div className="flex items-center gap-2 mb-3">
                    <div className="w-7 h-7 rounded-lg bg-gradient-to-br from-blue-500 to-cyan-600 flex items-center justify-center">
                        <Activity className="w-3.5 h-3.5 text-white" />
                    </div>
                    <h3 className="text-sm font-bold text-white">{t.left_sidebar.market_overview}</h3>
                </div>
                <div className="grid grid-cols-2 gap-2">
                    <div className="bg-white/5 rounded-lg p-2 hover:bg-white/10 transition-colors">
                        <div className="text-[10px] text-gray-400 mb-0.5">{t.left_sidebar.market_cap}</div>
                        <div className="text-sm font-bold text-white">{marketCap}</div>
                    </div>
                    <div className="bg-white/5 rounded-lg p-2 hover:bg-white/10 transition-colors">
                        <div className="text-[10px] text-gray-400 mb-0.5">{t.left_sidebar.vol_24h}</div>
                        <div className="text-sm font-bold text-white">{volume24h}</div>
                    </div>
                    <div className="bg-white/5 rounded-lg p-2 hover:bg-white/10 transition-colors">
                        <div className="text-[10px] text-gray-400 mb-0.5">{t.left_sidebar.btc_dom}</div>
                        <div className="text-sm font-bold text-blue-400">{btcDominance}</div>
                    </div>
                    <div className="bg-white/5 rounded-lg p-2 hover:bg-white/10 transition-colors">
                        <div className="text-[10px] text-gray-400 mb-0.5">{t.left_sidebar.active_coins}</div>
                        <div className="text-sm font-bold text-white">{activeCoins}</div>
                    </div>
                </div>
            </div>

            {/* Top Gainers */}
            <div className="bg-black/40 border border-white/5 rounded-xl p-3 hover:border-green-500/20 transition-all duration-300">
                <div className="flex items-center gap-2 mb-3">
                    <div className="w-7 h-7 rounded-lg bg-gradient-to-br from-green-500 to-emerald-600 flex items-center justify-center">
                        <TrendingUp className="w-4 h-4 text-white" />
                    </div>
                    <h3 className="text-sm font-bold text-white">{t.left_sidebar.top_gainers}</h3>
                </div>
                <div className="space-y-1.5">
                    {topGainers.length > 0 ? topGainers.slice(0, 2).map((coin, idx) => (
                        <div
                            key={coin.symbol}
                            className="flex items-center justify-between p-2 rounded-lg bg-white/5 hover:bg-gradient-to-r hover:from-green-500/10 hover:to-emerald-500/10 border border-transparent hover:border-green-500/20 transition-all duration-200 cursor-pointer group"
                        >
                            <div className="flex items-center gap-2">
                                <div className="relative">
                                    <div className="w-8 h-8 rounded-full bg-gradient-to-br from-green-500 to-emerald-600 flex items-center justify-center text-white text-xs font-bold shadow-lg shadow-green-500/20">
                                        {coin.symbol[0]}
                                    </div>
                                    <div className="absolute -top-0.5 -right-0.5 w-4 h-4 bg-green-500 rounded-full flex items-center justify-center text-[9px] font-bold text-white">
                                        {idx + 1}
                                    </div>
                                </div>
                                <div>
                                    <div className="text-xs font-bold text-white group-hover:text-green-400 transition-colors">{coin.symbol}</div>
                                    <div className="text-[10px] text-gray-500">{coin.name}</div>
                                </div>
                            </div>
                            <div className="text-right">
                                <div className="text-xs font-bold text-white">${coin.price.toLocaleString('en-US', { minimumFractionDigits: 2, maximumFractionDigits: 2 })}</div>
                                <div className="text-[11px] font-bold text-green-400 flex items-center gap-0.5 justify-end">
                                    <TrendingUp className="w-3 h-3" />
                                    +{coin.change.toFixed(2)}%
                                </div>
                            </div>
                        </div>
                    )) : (
                        <div className="text-xs text-gray-500 text-center py-4 flex flex-col items-center gap-2">
                            <div className="w-6 h-6 border-2 border-gray-700 border-t-green-500 rounded-full animate-spin"></div>
                            <span className="text-[10px]">Loading...</span>
                        </div>
                    )}
                </div>
            </div>

            {/* Top Losers */}
            <div className="bg-black/40 border border-white/5 rounded-xl p-3 hover:border-red-500/20 transition-all duration-300">
                <div className="flex items-center gap-2 mb-3">
                    <div className="w-7 h-7 rounded-lg bg-gradient-to-br from-red-500 to-rose-600 flex items-center justify-center">
                        <TrendingDown className="w-4 h-4 text-white" />
                    </div>
                    <h3 className="text-sm font-bold text-white">{t.left_sidebar.top_losers}</h3>
                </div>
                <div className="space-y-1.5">
                    {topLosers.length > 0 ? topLosers.slice(0, 2).map((coin, idx) => (
                        <div
                            key={coin.symbol}
                            className="flex items-center justify-between p-2 rounded-lg bg-white/5 hover:bg-gradient-to-r hover:from-red-500/10 hover:to-rose-500/10 border border-transparent hover:border-red-500/20 transition-all duration-200 cursor-pointer group"
                        >
                            <div className="flex items-center gap-2">
                                <div className="relative">
                                    <div className="w-8 h-8 rounded-full bg-gradient-to-br from-red-500 to-rose-600 flex items-center justify-center text-white text-xs font-bold shadow-lg shadow-red-500/20">
                                        {coin.symbol[0]}
                                    </div>
                                    <div className="absolute -top-0.5 -right-0.5 w-4 h-4 bg-red-500 rounded-full flex items-center justify-center text-[9px] font-bold text-white">
                                        {idx + 1}
                                    </div>
                                </div>
                                <div>
                                    <div className="text-xs font-bold text-white group-hover:text-red-400 transition-colors">{coin.symbol}</div>
                                    <div className="text-[10px] text-gray-500">{coin.name}</div>
                                </div>
                            </div>
                            <div className="text-right">
                                <div className="text-xs font-bold text-white">${coin.price.toLocaleString('en-US', { minimumFractionDigits: 2, maximumFractionDigits: 2 })}</div>
                                <div className="text-[11px] font-bold text-red-400 flex items-center gap-0.5 justify-end">
                                    <TrendingDown className="w-3 h-3" />
                                    {coin.change.toFixed(2)}%
                                </div>
                            </div>
                        </div>
                    )) : (
                        <div className="text-xs text-gray-500 text-center py-4 flex flex-col items-center gap-2">
                            <div className="w-6 h-6 border-2 border-gray-700 border-t-red-500 rounded-full animate-spin"></div>
                            <span className="text-[10px]">Loading...</span>
                        </div>
                    )}
                </div>
            </div>
        </div>
    );
}

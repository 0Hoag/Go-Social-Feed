"use client";

import { useEffect, useState } from 'react';
import { TrendingUp, TrendingDown, Activity, BarChart3 } from 'lucide-react';
import axios from 'axios';

interface MarketData {
    symbol: string;
    name: string;
    price: number;
    change: number;
}

export default function AnalysisLeftSidebar() {
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
        <div className="space-y-4">
            {/* Market Overview */}
            <div className="bg-[#111] border border-white/5 rounded-2xl p-4">
                <h3 className="text-sm font-bold text-white mb-4 flex items-center gap-2">
                    <Activity className="w-4 h-4 text-blue-500" />
                    Market Overview
                </h3>
                <div className="space-y-3">
                    <div className="flex justify-between items-center">
                        <span className="text-xs text-gray-500">Market Cap</span>
                        <span className="text-xs font-semibold text-white">{marketCap}</span>
                    </div>
                    <div className="flex justify-between items-center">
                        <span className="text-xs text-gray-500">24h Volume</span>
                        <span className="text-xs font-semibold text-white">{volume24h}</span>
                    </div>
                    <div className="flex justify-between items-center">
                        <span className="text-xs text-gray-500">BTC Dominance</span>
                        <span className="text-xs font-semibold text-white">{btcDominance}</span>
                    </div>
                    <div className="flex justify-between items-center">
                        <span className="text-xs text-gray-500">Active Coins</span>
                        <span className="text-xs font-semibold text-white">{activeCoins}</span>
                    </div>
                </div>
            </div>

            {/* Top Gainers */}
            <div className="bg-[#111] border border-white/5 rounded-2xl p-4">
                <h3 className="text-sm font-bold text-white mb-4 flex items-center gap-2">
                    <TrendingUp className="w-4 h-4 text-[#26a69a]" />
                    Top Gainers (24h)
                </h3>
                <div className="space-y-3">
                    {topGainers.length > 0 ? topGainers.map((coin) => (
                        <div key={coin.symbol} className="flex items-center justify-between">
                            <div className="flex items-center gap-2">
                                <div className="w-6 h-6 rounded-full bg-gradient-to-br from-green-500 to-emerald-600 flex items-center justify-center text-white text-[8px] font-bold">
                                    {coin.symbol[0]}
                                </div>
                                <div>
                                    <div className="text-xs font-medium text-white">{coin.symbol}</div>
                                    <div className="text-[10px] text-gray-500">{coin.name}</div>
                                </div>
                            </div>
                            <div className="text-right">
                                <div className="text-xs font-semibold text-white">${coin.price.toLocaleString('en-US', { minimumFractionDigits: 2, maximumFractionDigits: 2 })}</div>
                                <div className="text-[10px] font-semibold text-[#26a69a]">+{coin.change.toFixed(2)}%</div>
                            </div>
                        </div>
                    )) : (
                        <div className="text-xs text-gray-500 text-center py-2">Loading...</div>
                    )}
                </div>
            </div>

            {/* Top Losers */}
            <div className="bg-[#111] border border-white/5 rounded-2xl p-4">
                <h3 className="text-sm font-bold text-white mb-4 flex items-center gap-2">
                    <TrendingDown className="w-4 h-4 text-[#ef5350]" />
                    Top Losers (24h)
                </h3>
                <div className="space-y-3">
                    {topLosers.length > 0 ? topLosers.map((coin) => (
                        <div key={coin.symbol} className="flex items-center justify-between">
                            <div className="flex items-center gap-2">
                                <div className="w-6 h-6 rounded-full bg-gradient-to-br from-red-500 to-rose-600 flex items-center justify-center text-white text-[8px] font-bold">
                                    {coin.symbol[0]}
                                </div>
                                <div>
                                    <div className="text-xs font-medium text-white">{coin.symbol}</div>
                                    <div className="text-[10px] text-gray-500">{coin.name}</div>
                                </div>
                            </div>
                            <div className="text-right">
                                <div className="text-xs font-semibold text-white">${coin.price.toLocaleString('en-US', { minimumFractionDigits: 2, maximumFractionDigits: 2 })}</div>
                                <div className="text-[10px] font-semibold text-[#ef5350]">{coin.change.toFixed(2)}%</div>
                            </div>
                        </div>
                    )) : (
                        <div className="text-xs text-gray-500 text-center py-2">Loading...</div>
                    )}
                </div>
            </div>

            {/* Technical Indicators */}
            <div className="bg-[#111] border border-white/5 rounded-2xl p-4">
                <h3 className="text-sm font-bold text-white mb-4 flex items-center gap-2">
                    <BarChart3 className="w-4 h-4 text-purple-500" />
                    Technical Indicators
                </h3>
                <div className="space-y-3">
                    <div className="flex justify-between items-center">
                        <span className="text-xs text-gray-500">RSI (14)</span>
                        <span className="text-xs font-semibold text-yellow-500">58.3</span>
                    </div>
                    <div className="flex justify-between items-center">
                        <span className="text-xs text-gray-500">MACD</span>
                        <span className="text-xs font-semibold text-[#26a69a]">Bullish</span>
                    </div>
                    <div className="flex justify-between items-center">
                        <span className="text-xs text-gray-500">MA (50)</span>
                        <span className="text-xs font-semibold text-white">$63,245</span>
                    </div>
                    <div className="flex justify-between items-center">
                        <span className="text-xs text-gray-500">MA (200)</span>
                        <span className="text-xs font-semibold text-white">$58,890</span>
                    </div>
                </div>
            </div>
        </div>
    );
}

"use client";

import { useEffect, useState } from "react";
import { ArrowUp, ArrowDown, Activity } from "lucide-react";

interface CryptoPrice {
    symbol: string;
    price: string;
    percentChange: number;
}

export default function CryptoTicker() {
    const [prices, setPrices] = useState<CryptoPrice[]>([]);
    const [loading, setLoading] = useState(true);

    // Symbols to track (vs USDT)
    const symbols = ["BTCUSDT", "ETHUSDT", "BNBUSDT", "SOLUSDT", "XRPUSDT", "ADAUSDT", "DOGEUSDT", "TRXUSDT"];

    useEffect(() => {
        const fetchPrices = async () => {
            try {
                // Fetch 24h ticker price change statistics
                const res = await fetch("https://api.binance.com/api/v3/ticker/24hr");
                const data = await res.json();

                const filtered = data
                    .filter((item: any) => symbols.includes(item.symbol))
                    .map((item: any) => ({
                        symbol: item.symbol.replace("USDT", ""),
                        price: parseFloat(item.lastPrice).toLocaleString(undefined, { minimumFractionDigits: 2, maximumFractionDigits: 2 }),
                        percentChange: parseFloat(item.priceChangePercent),
                    }));

                // Sort by specified order
                const sorted = filtered.sort((a: CryptoPrice, b: CryptoPrice) =>
                    symbols.indexOf(a.symbol + "USDT") - symbols.indexOf(b.symbol + "USDT")
                );

                setPrices(sorted);
                setLoading(false);
            } catch (error) {
                console.error("Failed to fetch crypto prices:", error);
                setLoading(false);
            }
        };

        fetchPrices();
        const interval = setInterval(fetchPrices, 10000); // Update every 10s

        return () => clearInterval(interval);
    }, []);

    if (loading) return null;

    return (
        <div className="w-full bg-[#0a0a0a] border-b border-white/5 overflow-hidden py-2 flex items-center relative z-40">
            <div className="flex items-center gap-2 px-4 border-r border-white/10 shrink-0 bg-[#0a0a0a] z-10 text-xs font-bold text-gray-400 uppercase tracking-wider">
                <Activity className="w-3 h-3 text-cyan-500" />
                Live Market
            </div>

            {/* Ticker Container */}
            <div className="overflow-hidden whitespace-nowrap mask-linear-gradient w-full flex">
                <div className="flex items-center gap-8 animate-marquee shrink-0 pr-8">
                    {/* Duplicate list to ensure it covers screen width */}
                    {[...prices, ...prices, ...prices, ...prices].map((coin, index) => (
                        <div key={`orig-${coin.symbol}-${index}`} className="flex items-center gap-2 text-xs">
                            <span className="font-bold text-gray-300">{coin.symbol}</span>
                            <span className="text-gray-400">${coin.price}</span>
                            <span className={`flex items-center ${coin.percentChange >= 0 ? "text-green-500" : "text-red-500"}`}>
                                {coin.percentChange >= 0 ? <ArrowUp className="w-3 h-3" /> : <ArrowDown className="w-3 h-3" />}
                                {Math.abs(coin.percentChange).toFixed(2)}%
                            </span>
                        </div>
                    ))}
                </div>
                {/* Second duplicated container for seamless loop */}
                <div className="flex items-center gap-8 animate-marquee shrink-0 pr-8" aria-hidden="true">
                    {[...prices, ...prices, ...prices, ...prices].map((coin, index) => (
                        <div key={`copy-${coin.symbol}-${index}`} className="flex items-center gap-2 text-xs">
                            <span className="font-bold text-gray-300">{coin.symbol}</span>
                            <span className="text-gray-400">${coin.price}</span>
                            <span className={`flex items-center ${coin.percentChange >= 0 ? "text-green-500" : "text-red-500"}`}>
                                {coin.percentChange >= 0 ? <ArrowUp className="w-3 h-3" /> : <ArrowDown className="w-3 h-3" />}
                                {Math.abs(coin.percentChange).toFixed(2)}%
                            </span>
                        </div>
                    ))}
                </div>
            </div>
        </div>
    );
}

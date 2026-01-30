"use client";

import { useState, useEffect } from "react";
import TradingViewWidget from "@/components/TradingViewWidget";
import CoinList from "@/components/CoinList";
import { ArrowUp, ArrowDown } from "lucide-react";
import axios from "axios";

export default function AnalysisPage() {
    const [selectedSymbol, setSelectedSymbol] = useState("BINANCE:BTCUSDT");
    const [selectedCoinName, setSelectedCoinName] = useState("Bitcoin");
    const [currentPrice, setCurrentPrice] = useState(82695.50);

    // Fetch current price when coin changes
    useEffect(() => {
        const fetchPrice = async () => {
            try {
                // Extract symbol (e.g., "BTCUSDT" from "BINANCE:BTCUSDT")
                const symbol = selectedSymbol.split(":")[1];

                // Only fetch if it's a Binance symbol (not Gold or other assets)
                if (selectedSymbol.startsWith("BINANCE:")) {
                    const res = await axios.get(`https://api.binance.com/api/v3/ticker/price?symbol=${symbol}`);
                    setCurrentPrice(parseFloat(res.data.price));
                } else {
                    // For non-Binance symbols like Gold, use a placeholder
                    setCurrentPrice(2650.00);
                }
            } catch (error) {
                console.error("Failed to fetch current price:", error);
            }
        };

        fetchPrice();
    }, [selectedSymbol]);

    // Generate dynamic Order Book data based on current price
    const asks = Array.from({ length: 15 }, (_, i) => ({
        price: (currentPrice + (i + 1) * (currentPrice * 0.0001)).toFixed(2),
        amount: (Math.random() * 2).toFixed(4),
        total: (Math.random() * 5).toFixed(2),
        fill: Math.floor(Math.random() * 40),
    })).reverse();

    const bids = Array.from({ length: 15 }, (_, i) => ({
        price: (currentPrice - (i + 1) * (currentPrice * 0.0001)).toFixed(2),
        amount: (Math.random() * 2).toFixed(4),
        total: (Math.random() * 5).toFixed(2),
        fill: Math.floor(Math.random() * 40),
    }));

    return (
        <main className="min-h-screen bg-[#050505] text-gray-200 p-4">
            <div className="max-w-[1920px] mx-auto grid grid-cols-1 lg:grid-cols-12 gap-4 h-[calc(100vh-140px)]">

                {/* LEFT: Order Book (Simulated) */}
                <div className="hidden lg:block lg:col-span-2 xl:col-span-2 bg-[#111] border border-white/5 rounded-2xl overflow-hidden flex flex-col">
                    <div className="p-3 border-b border-white/5 flex justify-between items-center">
                        <h3 className="text-xs font-bold text-gray-400 uppercase tracking-wider">Order Book</h3>
                        <span className="text-[10px] text-gray-600">BTC/USDT</span>
                    </div>

                    <div className="flex-1 overflow-y-auto custom-scrollbar text-[10px] font-mono">
                        <div className="sticky top-0 bg-[#111] grid grid-cols-3 px-3 py-2 text-gray-500 font-bold border-b border-white/5 z-10">
                            <span>Price</span>
                            <span className="text-right">Amount</span>
                            <span className="text-right">Total</span>
                        </div>

                        {/* Asks (Sell) */}
                        <div className="py-1">
                            {asks.map((ask, i) => (
                                <div key={i} className="grid grid-cols-3 px-3 py-0.5 hover:bg-white/5 relative group cursor-pointer">
                                    <span className="text-red-400 z-10">{ask.price}</span>
                                    <span className="text-right text-gray-300 z-10">{ask.amount}</span>
                                    <span className="text-right text-gray-500 z-10">{ask.total}</span>
                                    {/* Valid visual depth bar */}
                                    <div className="absolute top-0 right-0 bottom-0 bg-red-500/10 transition-all duration-300" style={{ width: `${ask.fill}%` }}></div>
                                </div>
                            ))}
                        </div>

                        <div className="px-3 py-2 text-center text-lg font-bold text-white border-y border-white/5 bg-white/5 my-1">
                            {currentPrice.toLocaleString('en-US', { minimumFractionDigits: 2, maximumFractionDigits: 2 })}
                            <span className="text-[10px] text-gray-400 ml-2 font-normal">${currentPrice.toLocaleString('en-US', { minimumFractionDigits: 2, maximumFractionDigits: 2 })}</span>
                        </div>

                        {/* Bids (Buy) */}
                        <div className="py-1">
                            {bids.map((bid, i) => (
                                <div key={i} className="grid grid-cols-3 px-3 py-0.5 hover:bg-white/5 relative group cursor-pointer">
                                    <span className="text-green-400 z-10">{bid.price}</span>
                                    <span className="text-right text-gray-300 z-10">{bid.amount}</span>
                                    <span className="text-right text-gray-500 z-10">{bid.total}</span>
                                    <div className="absolute top-0 right-0 bottom-0 bg-green-500/10 transition-all duration-300" style={{ width: `${bid.fill}%` }}></div>
                                </div>
                            ))}
                        </div>
                    </div>
                </div>

                {/* MIDDLE: Chart */}
                <div className="col-span-1 lg:col-span-7 xl:col-span-8 bg-[#111] border border-white/5 rounded-2xl overflow-hidden relative">
                    <div className="absolute top-4 right-4 z-10 bg-black/50 backdrop-blur-sm px-3 py-1.5 rounded-lg border border-white/10">
                        <span className="text-xs text-gray-400 mr-2">Viewing:</span>
                        <span className="text-sm font-bold text-white">{selectedCoinName}</span>
                    </div>
                    <TradingViewWidget symbol={selectedSymbol} />
                </div>

                {/* RIGHT: Coin List (Clickable) */}
                <div className="hidden lg:block lg:col-span-3 xl:col-span-2 h-full overflow-hidden">
                    <CoinList
                        selectedSymbol={selectedSymbol}
                        onCoinSelect={(symbol, name) => {
                            setSelectedSymbol(symbol);
                            setSelectedCoinName(name);
                        }}
                    />
                </div>
            </div>
        </main>
    );
}

"use client";

import { useEffect, useState, useRef, memo } from "react";
import axios from "axios";
import { Search, ChevronLeft, ChevronRight } from "lucide-react";
import { useLanguage } from '@/context/LanguageContext';

interface CoinListProps {
    onCoinSelect: (symbol: string, name: string) => void;
    selectedSymbol: string;
}

interface CoinData {
    id: string;
    symbol: string;
    name: string;
    icon: string;
    price: string;
    change: string;
    changePercent: string;
    tvSymbol?: string;
    isGold?: boolean;
}

interface MiniTicker {
    s: string; // Symbol
    c: string; // Close price
    o: string; // Open price
    h: string; // High price
    l: string; // Low price
    v: string; // Volume
    q: string; // Quote Volume
}

function CoinList({ onCoinSelect, selectedSymbol }: CoinListProps) {
    const { t } = useLanguage();
    const [coins, setCoins] = useState<CoinData[]>([]);
    const [searchTerm, setSearchTerm] = useState("");
    const [currentPage, setCurrentPage] = useState(1);
    const scrollContainerRef = useRef<HTMLDivElement>(null);
    const latestTickersRef = useRef<Map<string, MiniTicker>>(new Map());
    const ITEMS_PER_PAGE = 8;

    useEffect(() => {
        const fetchCoins = async () => {
            try {
                const res = await axios.get("https://api.binance.com/api/v3/ticker/24hr");
                let data = res.data;

                // Filter for USDT pairs and sort by quote volume (liquidity)
                data = data.filter((t: any) => t.symbol.endsWith("USDT"));
                data.sort((a: any, b: any) => parseFloat(b.quoteVolume) - parseFloat(a.quoteVolume));

                // Take ALL pairs (no limit)
                const formattedCoins = data.map((t: any) => {
                    const symbol = t.symbol.replace("USDT", "");
                    return {
                        id: symbol.toLowerCase(),
                        symbol: t.symbol,
                        name: symbol, // We don't have full names without a huge map, using symbol as name
                        price: parseFloat(t.lastPrice).toLocaleString('en-US', { minimumFractionDigits: 2, maximumFractionDigits: 2 }),
                        change: parseFloat(t.priceChange).toFixed(2),
                        changePercent: parseFloat(t.priceChangePercent).toFixed(2),
                        icon: `https://assets.coincap.io/assets/icons/${symbol.toLowerCase()}@2x.png`,
                        tvSymbol: `BINANCE:${t.symbol}`
                    };
                });

                setCoins(formattedCoins);

            } catch (error) {
                console.error("Failed to fetch top coins:", error);
            }
        };

        fetchCoins();
        const interval = setInterval(fetchCoins, 10000); // Poll every 10s as a simple backup or use WS

        // WebSocket for real-time updates (Throttled)
        const ws = new WebSocket('wss://stream.binance.com:9443/ws/!miniTicker@arr');

        ws.onmessage = (event) => {
            const tickers = JSON.parse(event.data);
            tickers.forEach((t: MiniTicker) => {
                latestTickersRef.current.set(t.s, t);
            });
        };

        const throttleInterval = setInterval(() => {
            if (latestTickersRef.current.size === 0) return;

            setCoins(prevCoins => {
                let hasChanges = false;
                const newCoins = prevCoins.map(coin => {
                    const update = latestTickersRef.current.get(coin.symbol);
                    if (update) {
                        hasChanges = true;
                        const currentPrice = parseFloat(update.c);
                        const openPrice = parseFloat(update.o);
                        const change = currentPrice - openPrice;
                        const changePercent = (change / openPrice) * 100;

                        return {
                            ...coin,
                            price: currentPrice.toLocaleString('en-US', { minimumFractionDigits: 2, maximumFractionDigits: 2 }),
                            change: change.toFixed(2),
                            changePercent: changePercent.toFixed(2)
                        };
                    }
                    return coin;
                });

                if (hasChanges) {
                    latestTickersRef.current.clear(); // Clear consumed updates
                    return newCoins;
                }
                return prevCoins;
            });
        }, 2000); // 2 seconds throttle to reduce CPU load

        return () => {
            clearInterval(interval);
            clearInterval(throttleInterval);
            ws.close();
        }
    }, []);

    // Filter coins based on search term
    const filteredCoins = coins.filter(coin =>
        coin.name.toLowerCase().includes(searchTerm.toLowerCase()) ||
        coin.symbol.toLowerCase().includes(searchTerm.toLowerCase())
    );

    // Paginate filtered coins
    const totalPages = Math.ceil(filteredCoins.length / ITEMS_PER_PAGE);
    const startIndex = (currentPage - 1) * ITEMS_PER_PAGE;
    const endIndex = startIndex + ITEMS_PER_PAGE;
    const paginatedCoins = filteredCoins.slice(startIndex, endIndex);

    // Reset to page 1 when search term changes
    useEffect(() => {
        setCurrentPage(1);
    }, [searchTerm]);

    // Scroll to top when page changes
    useEffect(() => {
        if (scrollContainerRef.current) {
            scrollContainerRef.current.scrollTop = 0;
        }
    }, [currentPage]);

    return (
        <div className="h-full w-full bg-[#111] border border-white/5 rounded-2xl overflow-hidden flex flex-col">
            <div className="p-4 border-b border-white/5 space-y-3">
                <h3 className="text-xs font-bold text-gray-400 uppercase tracking-wider">{t.coin_list.market_overview}</h3>

                {/* Search Input */}
                {/* Search Input - Pro Style */}
                <div className="relative group z-20">
                    <div className="absolute inset-y-0 left-0 pl-3 flex items-center pointer-events-none">
                        <Search className="w-4 h-4 text-gray-500 group-focus-within:text-blue-400 transition-colors duration-300" />
                    </div>
                    <input
                        type="text"
                        placeholder={t.coin_list.search_placeholder}
                        value={searchTerm}
                        onChange={(e) => setSearchTerm(e.target.value)}
                        className="block w-full pl-10 pr-4 py-2.5 bg-[#0a0a0a]/60 border border-white/5 rounded-xl text-sm text-gray-200 placeholder-gray-600 focus:outline-none focus:border-blue-500/50 focus:ring-2 focus:ring-blue-500/20 focus:bg-[#0a0a0a]/90 transition-all duration-300 shadow-lg shadow-black/20 backdrop-blur-md"
                    />
                    <div className="absolute inset-0 rounded-xl bg-gradient-to-r from-blue-500/10 to-teal-500/10 opacity-0 group-focus-within:opacity-100 pointer-events-none transition-opacity duration-500" />
                </div>
            </div>

            <div ref={scrollContainerRef} className="flex-1 overflow-y-auto">
                {paginatedCoins.map((coin, index) => {
                    const isSelected = selectedSymbol === coin.symbol;
                    const isPositive = parseFloat(coin.changePercent) >= 0;
                    const isGold = coin.symbol === "XAUUSD";

                    // Show separator after Gold (first item)
                    const showSeparator = isGold && startIndex === 0 && index === 0;

                    return (
                        <div key={coin.symbol}>
                            <div
                                onClick={() => onCoinSelect(coin.symbol, coin.name)}
                                className={`
                                    px-3 py-3 cursor-pointer transition-all border-b border-white/5
                                    ${isSelected
                                        ? 'bg-blue-600/20 border-l-2 border-l-blue-500'
                                        : 'hover:bg-white/5'
                                    }
                                `}
                            >
                                <div className="flex items-center justify-between">
                                    <div className="flex items-center gap-2 flex-1 min-w-0">
                                        <img
                                            src={coin.icon}
                                            alt={coin.name}
                                            className="w-6 h-6 rounded-full flex-shrink-0 object-cover"
                                            onError={(e) => {
                                                e.currentTarget.onerror = null;
                                                e.currentTarget.src = `https://ui-avatars.com/api/?name=${coin.symbol}&background=random&color=fff&size=32`;
                                            }}
                                        />
                                        <div className="flex-1 min-w-0">
                                            <div className="text-xs font-semibold text-white truncate">{coin.name}</div>
                                            <div className="text-[10px] text-gray-500">{coin.symbol.replace('USDT', '')}</div>
                                        </div>
                                    </div>
                                    <div className="text-right ml-2">
                                        <div className="text-xs font-bold text-white">${coin.price}</div>
                                        <div className={`text-[10px] font-semibold ${isPositive ? 'text-green-400' : 'text-red-400'}`}>
                                            {isPositive ? '+' : ''}{coin.changePercent}%
                                        </div>
                                    </div>
                                </div>
                            </div>

                            {/* Separator after Gold */}
                            {showSeparator && (
                                <div className="my-2 mx-3 border-t-2 border-dashed border-gray-700 relative">
                                    <span className="absolute -top-2 left-1/2 transform -translate-x-1/2 bg-[#111] px-2 text-[9px] text-gray-500 uppercase tracking-wider">
                                        {t.coin_list.cryptocurrencies}
                                    </span>
                                </div>
                            )}
                        </div>
                    );
                })}
            </div>

            {/* Pagination Controls - Sticky at bottom */}
            {filteredCoins.length > 0 && (
                <div className="sticky bottom-0 bg-[#0a0a0a] border-t border-white/5 p-2 flex items-center justify-between text-xs">
                    <button
                        onClick={() => setCurrentPage(prev => Math.max(1, prev - 1))}
                        disabled={currentPage === 1}
                        className="flex items-center gap-1 px-2 py-1 rounded-lg bg-gradient-to-r from-blue-600/20 to-blue-500/20 hover:from-blue-600/30 hover:to-blue-500/30 border border-blue-500/30 text-blue-400 font-medium disabled:opacity-30 disabled:cursor-not-allowed transition-all disabled:hover:from-blue-600/20 disabled:hover:to-blue-500/20"
                    >
                        <ChevronLeft className="w-3 h-3" />
                    </button>

                    <div className="flex items-center gap-1.5">
                        <span className="text-gray-500 text-[9px] uppercase tracking-wider">{t.coin_list.page}</span>
                        <span className="px-1.5 py-0.5 bg-blue-500/20 border border-blue-500/30 rounded text-blue-400 font-bold text-[10px]">
                            {currentPage}
                        </span>
                        <span className="text-gray-600 text-[10px]">/</span>
                        <span className="text-gray-500 text-[10px]">{totalPages}</span>
                    </div>

                    <button
                        onClick={() => setCurrentPage(prev => Math.min(totalPages, prev + 1))}
                        disabled={currentPage === totalPages}
                        className="flex items-center gap-1 px-2 py-1 rounded-lg bg-gradient-to-r from-blue-600/20 to-blue-500/20 hover:from-blue-600/30 hover:to-blue-500/30 border border-blue-500/30 text-blue-400 font-medium disabled:opacity-30 disabled:cursor-not-allowed transition-all disabled:hover:from-blue-600/20 disabled:hover:to-blue-500/20"
                    >
                        <ChevronRight className="w-3 h-3" />
                    </button>
                </div>
            )}
        </div>
    );
}

const CoinDataList = memo(CoinList);
export default CoinDataList;

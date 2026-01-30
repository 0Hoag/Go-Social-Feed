import { TrendingUp, TrendingDown, Loader2 } from "lucide-react";
import Image from "next/image";
import { useEffect, useState } from "react";

interface CoinData {
    rank: number;
    symbol: string;
    price: number;
    change: number;
}

export default function CryptoRanking() {
    const [coins, setCoins] = useState<CoinData[]>([]);
    const [loading, setLoading] = useState(true);

    // Initial list to maintain order and selection
    const trackedSymbols = ["BTC", "ETH", "BNB", "XRP", "SOL", "TRX", "DOGE", "ADA", "BCH", "LINK"];

    useEffect(() => {
        const fetchPrices = async () => {
            try {
                const res = await fetch("https://api.binance.com/api/v3/ticker/24hr");
                const data = await res.json();

                const updatedCoins = trackedSymbols.map((sym, index) => {
                    const ticker = data.find((t: any) => t.symbol === `${sym}USDT`);
                    return {
                        rank: index + 1,
                        symbol: sym,
                        price: ticker ? parseFloat(ticker.lastPrice) : 0,
                        change: ticker ? parseFloat(ticker.priceChangePercent) : 0,
                    };
                });

                setCoins(updatedCoins);
                setLoading(false);
            } catch (error) {
                console.error("Failed to fetch crypto ranking:", error);
                setLoading(false);
            }
        };

        fetchPrices();
        // efficient polling every 30s
        const interval = setInterval(fetchPrices, 30000);
        return () => clearInterval(interval);
    }, []);

    if (loading) {
        return (
            <div className="bg-[#111] border border-white/5 rounded-2xl p-5 min-h-[400px] flex items-center justify-center">
                <Loader2 className="w-6 h-6 text-gray-600 animate-spin" />
            </div>
        );
    }

    return (
        <div className="bg-[#111] border border-white/5 rounded-2xl p-5">
            <h3 className="text-gray-400 text-xs font-bold tracking-wider uppercase mb-5 flex items-center gap-2">
                <span className="w-2 h-2 rounded-full bg-gray-600"></span> Top 10 Crypto
            </h3>

            <div className="space-y-4">
                {coins.map((coin) => (
                    <div key={coin.symbol} className="flex items-center justify-between group cursor-pointer">
                        <div className="flex items-center gap-3">
                            <span className="text-gray-600 text-xs font-mono w-4">{coin.rank}</span>
                            <div className="relative w-6 h-6 rounded-full overflow-hidden bg-gray-800">
                                <Image
                                    src={`https://assets.coincap.io/assets/icons/${coin.symbol.toLowerCase()}@2x.png`}
                                    alt={coin.symbol}
                                    fill
                                    className="object-cover"
                                    unoptimized
                                />
                            </div>
                            <span className="text-gray-300 font-bold text-xs group-hover:text-white transition-colors">
                                {coin.symbol}
                            </span>
                        </div>

                        <div className="text-right">
                            <div className="text-white text-xs font-medium">
                                ${coin.price.toLocaleString(undefined, { minimumFractionDigits: 2, maximumFractionDigits: coin.price < 1 ? 4 : 2 })}
                            </div>
                            <div className={`text-[10px] flex items-center justify-end gap-1 ${coin.change >= 0 ? 'text-green-500' : 'text-red-500'}`}>
                                {coin.change >= 0 ? '+' : ''}{coin.change.toFixed(2)}%
                            </div>
                        </div>
                    </div>
                ))}
            </div>
        </div>
    );
}

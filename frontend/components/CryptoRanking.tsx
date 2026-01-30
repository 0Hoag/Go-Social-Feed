import { TrendingUp, TrendingDown } from "lucide-react";

export default function CryptoRanking() {
    const coins = [
        { rank: 1, symbol: "BTC", price: 82473, change: -6.46, icon: "₿" },
        { rank: 2, symbol: "ETH", price: 2718, change: -7.85, icon: "Ξ" },
        { rank: 3, symbol: "BNB", price: 839.12, change: -6.64, icon: "BNB" },
        { rank: 4, symbol: "XRP", price: 1.74, change: -7.74, icon: "✕" },
        { rank: 5, symbol: "SOL", price: 113.54, change: -7.95, icon: "S" },
        { rank: 6, symbol: "TRX", price: 0.2909, change: -1.19, icon: "T" },
        { rank: 7, symbol: "DOGE", price: 0.1127, change: -7.41, icon: "Ð" },
        { rank: 8, symbol: "ADA", price: 0.3213, change: -8.41, icon: "A" },
        { rank: 9, symbol: "BCH", price: 540.49, change: -8.54, icon: "B" },
        { rank: 10, symbol: "LINK", price: 10.70, change: -7.89, icon: "L" },
    ];

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
                            <div className={`w-6 h-6 rounded-full flex items-center justify-center text-[10px] font-bold ${coin.symbol === 'BTC' ? 'bg-orange-500 text-white' :
                                coin.symbol === 'ETH' ? 'bg-indigo-400 text-white' :
                                    'bg-gray-800 text-gray-400'
                                }`}>
                                {coin.icon}
                            </div>
                            <span className="text-gray-300 font-bold text-xs group-hover:text-white transition-colors">
                                {coin.symbol}
                            </span>
                        </div>

                        <div className="text-right">
                            <div className="text-white text-xs font-medium">${coin.price.toLocaleString()}</div>
                            <div className={`text-[10px] flex items-center justify-end gap-1 ${coin.change >= 0 ? 'text-green-500' : 'text-red-500'}`}>
                                {coin.change}%
                            </div>
                        </div>
                    </div>
                ))}
            </div>
        </div>
    );
}

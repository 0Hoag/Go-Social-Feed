import { TrendingUp, TrendingDown, Activity, Calendar, Wallet, Clock, Loader2 } from "lucide-react";
import { useEffect, useState } from "react";

interface FNGData {
    value: string;
    value_classification: string;
    timestamp: string;
    time_until_update: string;
}

export default function MarketWidgets() {
    const [fng, setFng] = useState<FNGData | null>(null);
    const [loadingFng, setLoadingFng] = useState(true);

    useEffect(() => {
        const fetchFNG = async () => {
            try {
                const res = await fetch("https://api.alternative.me/fng/");
                const data = await res.json();
                if (data.data && data.data.length > 0) {
                    setFng(data.data[0]);
                }
                setLoadingFng(false);
            } catch (error) {
                console.error("Failed to fetch Fear & Greed Index:", error);
                setLoadingFng(false);
            }
        };

        fetchFNG();
    }, []);

    // Simulated Calendar Events (Static for demo purposes but formatted to look live)
    const events = [
        { title: "US Core Inflation Rate YoY", time: "1h 15m", importance: "high", forecast: "3.2%", previous: "3.1%" },
        { title: "Fed Interest Rate Decision", time: "Tomorrow", importance: "high", forecast: "5.50%", previous: "5.50%" },
        { title: "Initial Jobless Claims", time: "2h 30m", importance: "medium", forecast: "215K", previous: "212K" },
    ];

    const getFngColor = (value: number) => {
        if (value >= 75) return "text-green-500"; // Extreme Greed
        if (value >= 55) return "text-green-400"; // Greed
        if (value >= 45) return "text-yellow-400"; // Neutral
        if (value >= 25) return "text-orange-400"; // Fear
        return "text-red-500"; // Extreme Fear
    };

    const getFngPosition = (value: number) => {
        // Ensure 0-100 maps correctly to 0%-100% of the bar
        return Math.min(Math.max(value, 0), 100);
    };

    return (
        <div className="space-y-6">
            {/* Sentiment Widget (Fear & Greed) */}
            <div className="bg-[#111] border border-white/5 rounded-2xl p-5">
                <div className="flex items-center justify-between mb-4">
                    <h3 className="text-gray-400 text-xs font-bold tracking-wider uppercase flex items-center gap-2">
                        <Activity className="w-3 h-3" /> Market Sentiment
                    </h3>
                </div>

                {loadingFng ? (
                    <div className="h-32 flex items-center justify-center">
                        <Loader2 className="w-5 h-5 text-gray-600 animate-spin" />
                    </div>
                ) : fng ? (
                    <div className="space-y-4">
                        <div className="flex justify-between items-end">
                            <div>
                                <div className="text-gray-500 text-[10px] uppercase font-bold mb-1">Fear & Greed Index</div>
                                <div className={`text-2xl font-bold ${getFngColor(parseInt(fng.value))}`}>
                                    {fng.value}
                                </div>
                            </div>
                            <div className={`text-xs font-bold px-2 py-1 rounded bg-white/5 ${getFngColor(parseInt(fng.value))}`}>
                                {fng.value_classification}
                            </div>
                        </div>

                        {/* Progress Bar */}
                        <div className="h-3 w-full bg-gray-800 rounded-full overflow-hidden flex relative mt-2">
                            {/* Gradient Background */}
                            <div className="absolute inset-0 bg-gradient-to-r from-red-600 via-yellow-400 to-green-500 opacity-80" />

                            {/* Indicator Dot */}
                            <div
                                className="absolute top-0 bottom-0 w-1 bg-white shadow-[0_0_10px_rgba(255,255,255,0.8)] z-10 transition-all duration-1000 ease-out"
                                style={{ left: `${getFngPosition(parseInt(fng.value))}%` }}
                            />
                        </div>

                        <div className="flex justify-between text-[10px] font-medium text-gray-500 pt-1">
                            <span>Extreme Fear 0</span>
                            <span>100 Extreme Greed</span>
                        </div>
                    </div>
                ) : (
                    <div className="text-gray-500 text-xs">Failed to load data</div>
                )}
            </div>

            {/* Economic Calendar Widget */}
            <div className="bg-[#111] border border-white/5 rounded-2xl p-5">
                <div className="flex items-center justify-between mb-4">
                    <h3 className="text-gray-400 text-xs font-bold tracking-wider uppercase flex items-center gap-2">
                        <Calendar className="w-3 h-3" /> Economic Calendar
                    </h3>
                    <span className="text-[10px] text-blue-400 cursor-pointer hover:underline">VIEW ALL</span>
                </div>

                <div className="space-y-4">
                    {events.map((event, index) => (
                        <div key={index} className="flex gap-3 group">
                            <div className="mt-1">
                                <span className={`block w-2 h-2 rounded-full shadow-sm ${event.importance === 'high' ? 'bg-red-500 shadow-red-500/50' : 'bg-blue-500 shadow-blue-500/50'}`}></span>
                            </div>
                            <div className="flex-1">
                                <div className="flex justify-between items-start">
                                    <h4 className="text-gray-200 text-xs font-medium leading-tight group-hover:text-white transition-colors">
                                        {event.title}
                                    </h4>
                                    <span className="text-gray-500 text-[10px] whitespace-nowrap ml-2 bg-gray-800/50 px-1.5 py-0.5 rounded flex items-center gap-1">
                                        <Clock className="w-2.5 h-2.5" /> {event.time}
                                    </span>
                                </div>
                                <div className="flex gap-3 mt-1.5 text-[10px] text-gray-500">
                                    <span className="text-gray-400">Fcst: {event.forecast}</span>
                                    <span>Prev: {event.previous}</span>
                                </div>
                            </div>
                        </div>
                    ))}
                </div>
            </div>

            {/* On-Chain Flows Widget */}
            <div className="bg-[#111] border border-white/5 rounded-2xl p-5">
                <div className="flex items-center justify-between mb-4">
                    <h3 className="text-gray-400 text-xs font-bold tracking-wider uppercase flex items-center gap-2">
                        <Wallet className="w-3 h-3" /> On-Chain Flows
                    </h3>
                </div>

                <div className="space-y-3">
                    <div className="flex items-center gap-3 p-3 bg-white/5 rounded-lg border border-white/5 hover:bg-white/10 transition-colors cursor-pointer group">
                        <div className="w-8 h-8 rounded-full bg-green-500/10 flex items-center justify-center text-green-500 group-hover:scale-110 transition-transform">
                            <TrendingDown className="w-4 h-4" />
                        </div>
                        <div>
                            <div className="text-white text-xs font-medium">Exchange Outflow</div>
                            <div className="text-green-400 text-[10px] font-bold">-45,200 BTC <span className="text-gray-500 font-normal">(24h)</span></div>
                        </div>
                    </div>
                    <div className="flex items-center gap-3 p-3 bg-white/5 rounded-lg border border-white/5 hover:bg-white/10 transition-colors cursor-pointer group">
                        <div className="w-8 h-8 rounded-full bg-blue-500/10 flex items-center justify-center text-blue-500 group-hover:scale-110 transition-transform">
                            <Activity className="w-4 h-4" />
                        </div>
                        <div>
                            <div className="text-white text-xs font-medium">Whale Alert</div>
                            <div className="text-blue-400 text-[10px] font-bold">12,000 ETH <span className="text-gray-500 font-normal">moved</span></div>
                        </div>
                    </div>
                </div>
            </div>
        </div>
    );
}

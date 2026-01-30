import { TrendingUp, TrendingDown, Activity, Calendar, Wallet } from "lucide-react";

export default function MarketWidgets() {
    return (
        <div className="space-y-6">
            {/* Sentiment Widget */}
            <div className="bg-[#111] border border-white/5 rounded-2xl p-5">
                <div className="flex items-center justify-between mb-4">
                    <h3 className="text-gray-400 text-xs font-bold tracking-wider uppercase">Sentiment</h3>
                </div>

                <div className="space-y-4">
                    <div className="flex justify-between items-center text-sm">
                        <span className="text-white font-medium">CALL/PUT OPTIONS</span>
                    </div>

                    {/* Progress Bar */}
                    <div className="h-2 w-full bg-gray-800 rounded-full overflow-hidden flex relative">
                        <div className="h-full bg-gradient-to-r from-red-500 to-orange-500 w-[30%]" />
                        <div className="h-full bg-gray-700 w-[10%]" />
                        <div className="h-full bg-gradient-to-r from-green-500 to-emerald-400 flex-1" />

                        {/* Indicator Dot */}
                        <div className="absolute top-1/2 left-[65%] -translate-y-1/2 w-3 h-3 bg-yellow-400 rounded-full shadow-[0_0_10px_rgba(250,204,21,0.5)] border-2 border-[#111]" />
                    </div>

                    <div className="flex justify-between text-xs font-medium">
                        <span className="text-red-400 flex items-center gap-1">🐻 Bear</span>
                        <span className="text-green-400 flex items-center gap-1">Bullish 🐂</span>
                    </div>

                    <div className="grid grid-cols-3 gap-2 mt-4 pt-4 border-t border-white/5">
                        <div className="text-center">
                            <div className="text-gray-500 text-[10px] mb-1">P/C Ratio</div>
                            <div className="text-white text-sm font-mono">0.68</div>
                        </div>
                        <div className="text-center border-l border-white/5">
                            <div className="text-gray-500 text-[10px] mb-1">Max Pain</div>
                            <div className="text-white text-sm font-mono">$90K</div>
                        </div>
                        <div className="text-center border-l border-white/5">
                            <div className="text-gray-500 text-[10px] mb-1">Vol 24h</div>
                            <div className="text-white text-sm font-mono">59.6K</div>
                        </div>
                    </div>
                </div>
            </div>

            {/* Economic Calendar Widget */}
            <div className="bg-[#111] border border-white/5 rounded-2xl p-5">
                <div className="flex items-center justify-between mb-4">
                    <h3 className="text-gray-400 text-xs font-bold tracking-wider uppercase">Economic Calendar</h3>
                    <span className="text-[10px] text-blue-400 cursor-pointer hover:underline">VIEW ALL</span>
                </div>

                <div className="space-y-4">
                    <div className="flex gap-3">
                        <div className="mt-1">
                            <span className="block w-2 h-2 rounded-full bg-blue-500 shadow-[0_0_8px_rgba(59,130,246,0.5)]"></span>
                        </div>
                        <div className="flex-1">
                            <div className="flex justify-between items-start">
                                <h4 className="text-white text-xs font-medium leading-tight">French Prelim Private Payrolls q/q</h4>
                                <span className="text-gray-500 text-[10px] whitespace-nowrap ml-2">1h 21m</span>
                            </div>
                            <div className="flex gap-3 mt-1.5 text-[10px] text-gray-500">
                                <span>Actual: --</span>
                                <span>Forecast: 0.1%</span>
                                <span>Prev: -0.1%</span>
                            </div>
                        </div>
                    </div>

                    <div className="flex gap-3">
                        <div className="mt-1">
                            <span className="block w-2 h-2 rounded-full bg-blue-500 shadow-[0_0_8px_rgba(59,130,246,0.5)]"></span>
                        </div>
                        <div className="flex-1">
                            <div className="flex justify-between items-start">
                                <h4 className="text-white text-xs font-medium leading-tight">Italian Prelim GDP q/q</h4>
                                <span className="text-gray-500 text-[10px] whitespace-nowrap ml-2">2h 36m</span>
                            </div>
                            <div className="flex gap-3 mt-1.5 text-[10px] text-gray-500">
                                <span>Actual: --</span>
                                <span>Forecast: 0.2%</span>
                                <span>Prev: 0.1%</span>
                            </div>
                        </div>
                    </div>
                </div>
            </div>

            {/* On-Chain Flows Widget */}
            <div className="bg-[#111] border border-white/5 rounded-2xl p-5">
                <div className="flex items-center justify-between mb-4">
                    <h3 className="text-gray-400 text-xs font-bold tracking-wider uppercase">On-Chain Flows</h3>
                </div>

                <div className="space-y-3">
                    <div className="flex items-center gap-3 p-3 bg-white/5 rounded-lg border border-white/5">
                        <Wallet className="w-4 h-4 text-green-400" />
                        <div>
                            <div className="text-white text-xs font-medium">Exchange Outflow</div>
                            <div className="text-gray-500 text-[10px]">Data coming soon</div>
                        </div>
                    </div>
                    <div className="flex items-center gap-3 p-3 bg-white/5 rounded-lg border border-white/5">
                        <Activity className="w-4 h-4 text-blue-400" />
                        <div>
                            <div className="text-white text-xs font-medium">Whale Alert</div>
                            <div className="text-gray-500 text-[10px]">Data coming soon</div>
                        </div>
                    </div>
                </div>
            </div>
        </div>
    );
}

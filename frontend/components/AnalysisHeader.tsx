import { ArrowUp, ArrowDown, ExternalLink, Share2, Star } from 'lucide-react';
import { TokenMetadata } from '@/lib/tokenConstants';

interface AnalysisHeaderProps {
    symbol: string;
    name: string;
    price: number;
    priceChange: number;
    stats: {
        high: string;
        low: string;
        vol: string;
    };
    metadata: TokenMetadata;
}

export default function AnalysisHeader({ symbol, name, price, priceChange, stats, metadata }: AnalysisHeaderProps) {
    const isPositive = priceChange >= 0;
    const formatPrice = (val: number) => val.toLocaleString('en-US', { minimumFractionDigits: 2, maximumFractionDigits: 2 });

    return (
        <div className="bg-[#111] border-b border-white/5 p-3 lg:p-4">
            <div className="max-w-[1920px] mx-auto">
                <div className="flex items-center justify-between gap-4">
                    <div className="flex items-center gap-3">
                        {/* Placeholder Logo */}
                        <div className="w-7 h-7 rounded-full bg-gradient-to-br from-blue-500 to-purple-600 flex items-center justify-center text-white font-bold text-xs">
                            {symbol[0]}
                        </div>
                        <div>
                            <div className="flex items-center gap-2">
                                <h1 className="text-base font-bold text-white">{name}</h1>
                                <span className="bg-white/10 text-gray-400 text-[9px] px-1 py-0.5 rounded">{symbol}</span>
                                <button className="text-gray-500 hover:text-yellow-400 transition-colors">
                                    <Star className="w-3 h-3" />
                                </button>
                            </div>
                        </div>
                    </div>

                    {/* Price & Change - Compact */}
                    <div className="flex items-center gap-3">
                        <span className="text-xl font-bold text-white">
                            ${formatPrice(price)}
                        </span>
                        <span className={`flex items-center gap-0.5 text-xs font-semibold ${isPositive ? 'text-[#26a69a]' : 'text-[#ef5350]'}`}>
                            {isPositive ? <ArrowUp className="w-3 h-3" /> : <ArrowDown className="w-3 h-3" />}
                            {Math.abs(priceChange).toFixed(2)}% (1d)
                        </span>
                    </div>

                    {/* Stats - Ultra Compact */}
                    <div className="flex items-center gap-3 text-xs">
                        <div className="flex flex-col">
                            <span className="text-gray-500 text-[9px]">Vol (24h)</span>
                            <span className="text-white text-xs font-medium">${parseInt(stats.vol || '0').toLocaleString()}</span>
                        </div>
                        <div className="flex flex-col">
                            <span className="text-gray-500 text-[9px]">Supply</span>
                            <span className="text-white text-xs font-medium">{metadata.circulatingSupply || 'N/A'}</span>
                        </div>
                    </div>


                </div>
            </div>
        </div>
    );
}

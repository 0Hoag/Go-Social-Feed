"use client";

import { useEffect, useState } from "react";
import axios from "axios";
import Image from "next/image";

interface CoinListProps {
    onCoinSelect: (symbol: string, name: string) => void;
    selectedSymbol: string;
}

interface CoinData {
    symbol: string;
    name: string;
    price: string;
    change: string;
    changePercent: string;
    icon: string;
    tvSymbol?: string;
}

export default function CoinList({ onCoinSelect, selectedSymbol }: CoinListProps) {
    const [coins, setCoins] = useState<CoinData[]>([]);

    const coinMapping = [
        // Commodity (separated at top)
        { symbol: "XAUUSD", name: "Gold", id: "gold", logo: "gold", tvSymbol: "OANDA:XAUUSD", isGold: true },

        // Top 25 Cryptocurrencies
        { symbol: "BTCUSDT", name: "Bitcoin", id: "1", logo: "btc", tvSymbol: "BINANCE:BTCUSDT" },
        { symbol: "ETHUSDT", name: "Ethereum", id: "1027", logo: "eth", tvSymbol: "BINANCE:ETHUSDT" },
        { symbol: "SOLUSDT", name: "Solana", id: "5426", logo: "sol", tvSymbol: "BINANCE:SOLUSDT" },
        { symbol: "BNBUSDT", name: "BNB", id: "1839", logo: "bnb", tvSymbol: "BINANCE:BNBUSDT" },
        { symbol: "XRPUSDT", name: "XRP", id: "52", logo: "xrp", tvSymbol: "BINANCE:XRPUSDT" },
        { symbol: "DOGEUSDT", name: "Dogecoin", id: "74", logo: "doge", tvSymbol: "BINANCE:DOGEUSDT" },
        { symbol: "ADAUSDT", name: "Cardano", id: "2010", logo: "ada", tvSymbol: "BINANCE:ADAUSDT" },
        { symbol: "AVAXUSDT", name: "Avalanche", id: "5805", logo: "avax", tvSymbol: "BINANCE:AVAXUSDT" },
        { symbol: "TRXUSDT", name: "Tron", id: "1958", logo: "trx", tvSymbol: "BINANCE:TRXUSDT" },
        { symbol: "LINKUSDT", name: "Chainlink", id: "1975", logo: "link", tvSymbol: "BINANCE:LINKUSDT" },
        { symbol: "DOTUSDT", name: "Polkadot", id: "6636", logo: "dot", tvSymbol: "BINANCE:DOTUSDT" },
        { symbol: "MATICUSDT", name: "Polygon", id: "3890", logo: "matic", tvSymbol: "BINANCE:MATICUSDT" },
        { symbol: "TONUSDT", name: "Toncoin", id: "11419", logo: "ton", tvSymbol: "BINANCE:TONUSDT" },
        { symbol: "LTCUSDT", name: "Litecoin", id: "2", logo: "ltc", tvSymbol: "BINANCE:LTCUSDT" },
        { symbol: "WBTCUSDT", name: "Wrapped Bitcoin", id: "3717", logo: "wbtc", tvSymbol: "BINANCE:WBTCUSDT" },
        { symbol: "NEARUSDT", name: "NEAR Protocol", id: "6535", logo: "near", tvSymbol: "BINANCE:NEARUSDT" },
        { symbol: "UNIUSDT", name: "Uniswap", id: "7083", logo: "uni", tvSymbol: "BINANCE:UNIUSDT" },
        { symbol: "APTUSDT", name: "Aptos", id: "21794", logo: "apt", tvSymbol: "BINANCE:APTUSDT" },
        { symbol: "ATOMUSDT", name: "Cosmos", id: "3794", logo: "atom", tvSymbol: "BINANCE:ATOMUSDT" },
        { symbol: "SUIUSDT", name: "Sui", id: "20947", logo: "sui", tvSymbol: "BINANCE:SUIUSDT" },
        { symbol: "PEPEUSDT", name: "Pepe", id: "24478", logo: "pepe", tvSymbol: "BINANCE:PEPEUSDT" },
        { symbol: "ARBUSDT", name: "Arbitrum", id: "11841", logo: "arb", tvSymbol: "BINANCE:ARBUSDT" },
        { symbol: "ENAUSDT", name: "Ethena", id: "30171", logo: "ena", tvSymbol: "BINANCE:ENAUSDT" },
        { symbol: "OPUSDT", name: "Optimism", id: "11840", logo: "op", tvSymbol: "BINANCE:OPUSDT" },
        { symbol: "SHIBUSDT", name: "Shiba Inu", id: "5994", logo: "shib", tvSymbol: "BINANCE:SHIBUSDT" },
    ];

    useEffect(() => {
        const fetchData = async () => {
            try {
                // Separate crypto and gold
                const cryptoCoins = coinMapping.filter(c => !c.isGold);
                const symbols = cryptoCoins.map(c => c.symbol).join(",");
                const res = await axios.get(`https://api.binance.com/api/v3/ticker/24hr?symbols=["${symbols.split(",").join('","')}"]`);

                const mapped = res.data.map((item: any) => {
                    const config = coinMapping.find(c => c.symbol === item.symbol);
                    return {
                        symbol: item.symbol,
                        name: config?.name || item.symbol,
                        price: parseFloat(item.lastPrice).toLocaleString('en-US', { minimumFractionDigits: 2, maximumFractionDigits: 2 }),
                        change: parseFloat(item.priceChange).toFixed(2),
                        changePercent: parseFloat(item.priceChangePercent).toFixed(2),
                        icon: `https://s2.coinmarketcap.com/static/img/coins/64x64/${config?.id}.png`,
                        tvSymbol: config?.tvSymbol,
                    };
                });

                // Add gold as THE FIRST entry (prepend, not append)
                const goldConfig = coinMapping.find(c => c.isGold);
                if (goldConfig) {
                    mapped.unshift({
                        symbol: goldConfig.symbol,
                        name: goldConfig.name,
                        price: "2,650.00", // Placeholder - could fetch from another API
                        change: "+5.20",
                        changePercent: "+0.20",
                        icon: "https://s2.coinmarketcap.com/static/img/coins/64x64/7226.png", // Gold futures icon from CMC
                        tvSymbol: goldConfig.tvSymbol,
                    });
                }

                setCoins(mapped);
            } catch (error) {
                console.error("Failed to fetch coin data:", error);
            }
        };

        fetchData();
        const interval = setInterval(fetchData, 30000);
        return () => clearInterval(interval);
    }, []);

    return (
        <div className="h-full w-full bg-[#111] border border-white/5 rounded-2xl overflow-hidden flex flex-col">
            <div className="p-4 border-b border-white/5">
                <h3 className="text-xs font-bold text-gray-400 uppercase tracking-wider">Market Overview</h3>
            </div>
            <div className="flex-1 overflow-y-auto custom-scrollbar">
                {coins.map((coin, index) => {
                    const coinTvSymbol = coin.tvSymbol || `BINANCE:${coin.symbol}`;
                    const isSelected = selectedSymbol === coinTvSymbol;
                    const isPositive = parseFloat(coin.changePercent) >= 0;
                    const isGold = coin.symbol === "XAUUSD";

                    // Show separator after Gold (first item)
                    const showDivider = index === 0 && isGold;

                    return (
                        <div key={coin.symbol}>
                            <div
                                onClick={() => onCoinSelect(coinTvSymbol, coin.name)}
                                className={`flex items-center justify-between p-3 border-b border-white/5 cursor-pointer transition-all ${isSelected ? 'bg-blue-500/20 border-l-4 border-l-blue-500' : 'hover:bg-white/5'
                                    }`}
                            >
                                <div className="flex items-center gap-3 flex-1 min-w-0">
                                    <img src={coin.icon} alt={coin.name} className="w-8 h-8 rounded-full" />
                                    <div className="flex-1 min-w-0">
                                        <div className="text-sm font-bold text-white truncate">{coin.name}</div>
                                        <div className="text-[10px] text-gray-500 font-mono">
                                            {isGold ? "XAU/USD" : coin.symbol.replace('USDT', '/USDT')}
                                        </div>
                                    </div>
                                </div>
                                <div className="text-right ml-2">
                                    <div className="text-sm font-bold text-white">${coin.price}</div>
                                    <div className={`text-[10px] font-bold ${isPositive ? 'text-green-400' : 'text-red-400'}`}>
                                        {isPositive ? '+' : ''}{coin.changePercent}%
                                    </div>
                                </div>
                            </div>

                            {/* Divider after Gold */}
                            {showDivider && (
                                <div className="my-2 mx-3 border-t-2 border-dashed border-gray-700 relative">
                                    <span className="absolute -top-2 left-1/2 transform -translate-x-1/2 bg-[#111] px-2 text-[9px] text-gray-500 uppercase tracking-wider">
                                        Cryptocurrencies
                                    </span>
                                </div>
                            )}
                        </div>
                    );
                })}
            </div>
        </div>
    );
}

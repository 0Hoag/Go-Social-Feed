"use client";

import { useEffect, useState } from "react";
import axios from "axios";

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
        // Initial setup from config
        const initialCoins = coinMapping.map(c => ({
            id: c.id,
            symbol: c.symbol,
            name: c.name,
            price: "0.00",
            change: "0.00",
            changePercent: "0.00",
            icon: `https://s2.coinmarketcap.com/static/img/coins/64x64/${c.id}.png`,
            tvSymbol: c.tvSymbol
        }));
        setCoins(initialCoins);

        // WebSocket for Crypto
        const ws = new WebSocket('wss://stream.binance.com:9443/ws/!miniTicker@arr');

        ws.onmessage = (event) => {
            const data = JSON.parse(event.data);

            setCoins(prevCoins => {
                const newCoins = [...prevCoins];
                let updated = false;

                data.forEach((ticker: any) => {
                    const index = newCoins.findIndex(c => c.symbol === ticker.s);
                    if (index !== -1) {
                        newCoins[index] = {
                            ...newCoins[index],
                            price: parseFloat(ticker.c).toLocaleString('en-US', { minimumFractionDigits: 2, maximumFractionDigits: 2 }),
                            change: (parseFloat(ticker.c) - parseFloat(ticker.o)).toFixed(2),
                            changePercent: ((parseFloat(ticker.c) - parseFloat(ticker.o)) / parseFloat(ticker.o) * 100).toFixed(2),
                        };
                        updated = true;
                    }
                });

                // Set Gold manually since it's not in Binance Crypto WS
                // In a real app, this would need a separate API/WS
                const goldIndex = newCoins.findIndex(c => c.symbol === "XAUUSD");
                if (goldIndex !== -1) {
                    newCoins[goldIndex] = {
                        ...newCoins[goldIndex],
                        price: "2,650.00",
                        change: "+5.20",
                        changePercent: "+0.20"
                    };
                }

                return updated ? newCoins : prevCoins;
            });
        };

        return () => ws.close();
    }, []);

    return (
        <div className="h-full w-full bg-[#111] border border-white/5 rounded-2xl overflow-hidden flex flex-col">
            <div className="p-4 border-b border-white/5">
                <h3 className="text-xs font-bold text-gray-400 uppercase tracking-wider">Market Overview</h3>
            </div>
            <div className="flex-1 overflow-y-auto custom-scrollbar">
                {coins.map((coin, index) => {
                    const isSelected = selectedSymbol === coin.symbol;
                    const isPositive = parseFloat(coin.changePercent) >= 0;
                    const isGold = coin.symbol === "XAUUSD";

                    // Show separator after Gold (first item)
                    const showSeparator = isGold && index === 0;

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
                                                e.currentTarget.onerror = null; // Prevent infinite loop if fallback fails
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

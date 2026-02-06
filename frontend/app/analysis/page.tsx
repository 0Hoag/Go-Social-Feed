"use client";

import { useState, useEffect } from "react";
import { LineChart, CandlestickChart, ArrowUp, ArrowDown } from "lucide-react";
import ProfessionalChart from "@/components/ProfessionalChart";
import CoinList from "@/components/CoinList";
import LoadingScreen from "@/components/LoadingScreen";
import AnalysisLeftSidebar from "@/components/AnalysisLeftSidebar";

import axios from "axios";
import { getMetadata } from "@/lib/tokenConstants";

export default function AnalysisPage() {
    const [isLoading, setIsLoading] = useState(true);
    const [isDataReady, setIsDataReady] = useState(false);
    const [selectedSymbol, setSelectedSymbol] = useState("BTCUSDT");
    const [selectedCoinName, setSelectedCoinName] = useState("Bitcoin");
    const [currentPrice, setCurrentPrice] = useState(65758.13);
    const [priceChange, setPriceChange] = useState(-6.78);

    const [stats, setStats] = useState({ high: '0.00', low: '0.00', vol: '0.00' });
    const [chartType, setChartType] = useState<'candle' | 'area'>('candle');
    const [timeframe, setTimeframe] = useState<string>('15m');
    const [coinData, setCoinData] = useState({ ath: 0, athChangePercent: 0, rank: 0, dominance: 0 });

    // Fetch current price and stats via WebSocket
    useEffect(() => {
        // Initial Fetch
        const fetchInitial = async () => {
            try {
                const res = await axios.get(`https://api.binance.com/api/v3/ticker/24hr?symbol=${selectedSymbol}`);
                const data = res.data;
                setCurrentPrice(parseFloat(data.lastPrice));
                setPriceChange(parseFloat(data.priceChangePercent));
                setStats({
                    high: parseFloat(data.highPrice).toFixed(2),
                    low: parseFloat(data.lowPrice).toFixed(2),
                    vol: parseFloat(data.volume).toFixed(2)
                });
                setIsDataReady(true);
            } catch (error) {
                console.error("Failed to fetch ticker:", error);
            }
        };
        fetchInitial();

        // WebSocket for 24h ticker updates (real-time price & stats)
        const tickerWs = new WebSocket(`wss://stream.binance.com:9443/ws/${selectedSymbol.toLowerCase()}@ticker`);

        tickerWs.onmessage = (event) => {
            const ticker = JSON.parse(event.data);
            if (ticker) {
                setCurrentPrice(parseFloat(ticker.c));
                setPriceChange(parseFloat(ticker.P));
                setStats({
                    high: parseFloat(ticker.h).toFixed(2),
                    low: parseFloat(ticker.l).toFixed(2),
                    vol: parseFloat(ticker.v).toFixed(2)
                });
            }
        };

        return () => tickerWs.close();
    }, [selectedSymbol]);

    // Fetch ATH and other metadata from CoinGecko
    useEffect(() => {
        const fetchCoinData = async () => {
            try {
                // Map Binance symbols to CoinGecko IDs (Top 100 coins)
                const coinGeckoMap: { [key: string]: string } = {
                    'BTCUSDT': 'bitcoin',
                    'ETHUSDT': 'ethereum',
                    'USDTUSDT': 'tether',
                    'BNBUSDT': 'binancecoin',
                    'SOLUSDT': 'solana',
                    'USDCUSDT': 'usd-coin',
                    'XRPUSDT': 'ripple',
                    'STETHUSDT': 'staked-ether',
                    'DOGEUSDT': 'dogecoin',
                    'ADAUSDT': 'cardano',
                    'TRXUSDT': 'tron',
                    'AVAXUSDT': 'avalanche-2',
                    'WBTCUSDT': 'wrapped-bitcoin',
                    'LINKUSDT': 'chainlink',
                    'TONUSDT': 'the-open-network',
                    'SHIBUSDT': 'shiba-inu',
                    'DOTUSDT': 'polkadot',
                    'BCHUSDT': 'bitcoin-cash',
                    'SUIUSDT': 'sui',
                    'DAIUSDT': 'dai',
                    'LTCUSDT': 'litecoin',
                    'UNIUSDT': 'uniswap',
                    'NEARUSDT': 'near',
                    'MATICUSDT': 'matic-network',
                    'PEPEUSDT': 'pepe',
                    'ICPUSDT': 'internet-computer',
                    'APTUSDT': 'aptos',
                    'FETUSDT': 'fetch-ai',
                    'ETCUSDT': 'ethereum-classic',
                    'XLMUSDT': 'stellar',
                    'KASUSDT': 'kaspa',
                    'FILUSDT': 'filecoin',
                    'ARBUSDT': 'arbitrum',
                    'RENDERUSDT': 'render-token',
                    'RNDRUSDT': 'render-token',
                    'OKBUSDT': 'okb',
                    'IMXUSDT': 'immutable-x',
                    'INJUSDT': 'injective-protocol',
                    'ATOMUSDT': 'cosmos',
                    'MKRUSDT': 'maker',
                    'VETUSDT': 'vechain',
                    'OPUSDT': 'optimism',
                    'GRTUSDT': 'the-graph',
                    'TAOUSDT': 'bittensor',
                    'WIFUSDT': 'dogwifcoin',
                    'ALGOUSDT': 'algorand',
                    'AAVEUSDT': 'aave',
                    'THETAUSDT': 'theta-token',
                    'HBARUSDT': 'hedera-hashgraph',
                    'LDOUSDT': 'lido-dao',
                    'RUNEUSDT': 'thorchain',
                    'FTMUSDT': 'fantom',
                    'BONKUSDT': 'bonk',
                    'FLOKIUSDT': 'floki',
                    'MANAUSDT': 'decentraland',
                    'SANDUSDT': 'the-sandbox',
                    'AXSUSDT': 'axie-infinity',
                    'GMTUSDT': 'stepn',
                    'SNXUSDT': 'havven',
                    'COMPUSDT': 'compound-governance-token',
                    'CRVUSDT': 'curve-dao-token',
                    '1INCHUSDT': '1inch',
                    'CHZUSDT': 'chiliz',
                    'ZILUSDT': 'zilliqa',
                    'ENSUSDT': 'ethereum-name-service',
                    'QNTUSDT': 'quant-network',
                    'FLOWUSDT': 'flow',
                    'XTZUSDT': 'tezos',
                    'EOSUSDT': 'eos',
                    'ARUSDT': 'arweave',
                    'NEOUSDT': 'neo',
                    'MINAUSDT': 'mina-protocol',
                    'EGLDUSDT': 'elrond-egd-2',
                    'XMRUSDT': 'monero',
                    'BTGUSDT': 'bitcoin-gold',
                    'KSMUSDT': 'kusama',
                    'AXLUSDT': 'axelar',
                    'CFXUSDT': 'conflux-token',
                    'KLAYUSDT': 'klay-token',
                    'GALAUSDT': 'gala',
                    'ENJUSDT': 'enjincoin',
                    'SXPUSDT': 'swipe',
                    'CAKEUSDT': 'pancakeswap-token',
                    'APEUSDT': 'apecoin',
                    'BLURUSDT': 'blur',
                    'LRCUSDT': 'loopring',
                    'BATUSDT': 'basic-attention-token',
                    'ZRXUSDT': '0x',
                    'IOTAUSDT': 'iota',
                    'OMGUSDT': 'omisego',
                    'ICXUSDT': 'icon',
                    'ONTUSDT': 'ontology',
                    'QTUMUSDT': 'qtum',
                    'ZECUSDT': 'zcash',
                    'DASHUSDT': 'dash',
                    'WLDUSDT': 'worldcoin-wld',
                    'JUPUSDT': 'jupiter-exchange-solana',
                    'PYTHUSDT': 'pyth-network',
                    'SEIUSDT': 'sei-network',
                    'TIAUSDT': 'celestia',
                    'DYMUSDT': 'dymension',
                    'STRKUSDT': 'starknet',
                    'PENDLEUSDT': 'pendle',
                    'BEAMXUSDT': 'beam-2'
                };

                const coinId = coinGeckoMap[selectedSymbol] || 'bitcoin';
                const res = await axios.get(`https://api.coingecko.com/api/v3/coins/${coinId}`);
                const data = res.data;

                setCoinData({
                    ath: data.market_data.ath.usd,
                    athChangePercent: data.market_data.ath_change_percentage.usd,
                    rank: data.market_cap_rank || 0,
                    dominance: data.market_data.market_cap_percentage?.btc || 0
                });
            } catch (error) {
                console.error('Failed to fetch CoinGecko data:', error);
            }
        };

        fetchCoinData();
        const coinDataInterval = setInterval(fetchCoinData, 300000) as NodeJS.Timeout; // Update every 5 minutes

        return () => clearInterval(coinDataInterval);
    }, [selectedSymbol]);


    // Handle loading completion
    const handleLoadingComplete = () => {
        setIsLoading(false);
    };

    // Auto-hide loading if data is ready
    // Removed to allow LoadingScreen to control completion via isReady prop
    /* useEffect(() => {
        if (isDataReady) {
            const timer = setTimeout(() => setIsLoading(false), 500);
            return () => clearTimeout(timer);
        }
    }, [isDataReady]); */

    return (
        <>
            {isLoading && <LoadingScreen isReady={isDataReady} onLoadingComplete={handleLoadingComplete} />}

            <main className={`min-h-screen bg-[#050505] text-gray-200 transition-opacity duration-500 ${isLoading ? 'opacity-0' : 'opacity-100'}`}>



                <div className="max-w-[1920px] mx-auto p-4 lg:p-6 grid grid-cols-1 lg:grid-cols-12 gap-4">
                    {/* LEFT: Sidebar */}
                    <div className="lg:col-span-3">
                        <div className="sticky top-4 overflow-y-auto" style={{ maxHeight: 'calc(100vh - 120px)' }}>
                            <AnalysisLeftSidebar />
                        </div>
                    </div>

                    {/* CENTER: Main Content (Chart) */}
                    <div className="lg:col-span-7 flex flex-col gap-6">
                        {/* Chart Container */}
                        <div className="bg-[#111] border border-white/5 rounded-2xl overflow-hidden h-[600px] flex flex-col">
                            {/* Integrated Header + Chart Controls */}
                            <div className="border-b border-white/5 px-3 py-2 flex items-center justify-between">
                                {/* LEFT: Coin Info */}
                                <div className="flex items-center gap-2">
                                    <img
                                        src={`https://assets.coincap.io/assets/icons/${selectedSymbol.replace('USDT', '').toLowerCase()}@2x.png`}
                                        alt={selectedCoinName}
                                        className="w-6 h-6 rounded-full bg-gradient-to-br from-orange-500 to-yellow-600"
                                        onError={(e) => {
                                            const target = e.currentTarget;
                                            // Try alternative source
                                            if (target.src.includes('coincap.io')) {
                                                target.src = `https://cryptologos.cc/logos/${selectedCoinName.toLowerCase().replace(/\s+/g, '-')}-${selectedSymbol.replace('USDT', '').toLowerCase()}-logo.png`;
                                            } else if (target.src.includes('cryptologos.cc')) {
                                                // Final fallback to letter avatar
                                                target.src = `https://ui-avatars.com/api/?name=${selectedSymbol.replace('USDT', '')}&background=f7931a&color=fff&size=24&bold=true`;
                                            }
                                        }}
                                    />
                                    <div className="flex items-center gap-2">
                                        <div className="flex items-baseline gap-1">
                                            <h1 className="text-sm font-bold text-white">{selectedCoinName}</h1>
                                            <span className="bg-white/10 text-gray-400 text-[9px] px-1 py-0.5 rounded">{selectedSymbol.replace('USDT', '')}</span>
                                        </div>
                                        <div className="h-3 w-px bg-white/10"></div>
                                        <span className="text-base font-bold text-white">
                                            ${currentPrice.toLocaleString('en-US', { minimumFractionDigits: 2, maximumFractionDigits: 2 })}
                                        </span>
                                        <span className={`flex items-center gap-0.5 text-xs font-semibold ${priceChange >= 0 ? 'text-[#26a69a]' : 'text-[#ef5350]'}`}>
                                            {priceChange >= 0 ? <ArrowUp className="w-3 h-3" /> : <ArrowDown className="w-3 h-3" />}
                                            {Math.abs(priceChange).toFixed(2)}%
                                        </span>
                                    </div>
                                </div>

                                {/* MIDDLE: 24h Stats */}
                                <div className="flex items-center gap-3 text-[10px]">
                                    <div className="flex flex-col">
                                        <span className="text-gray-500">24h High</span>
                                        <span className="text-white font-semibold">${parseFloat(stats.high || '0').toLocaleString('en-US', { minimumFractionDigits: 2, maximumFractionDigits: 2 })}</span>
                                    </div>
                                    <div className="flex flex-col">
                                        <span className="text-gray-500">24h Low</span>
                                        <span className="text-white font-semibold">${parseFloat(stats.low || '0').toLocaleString('en-US', { minimumFractionDigits: 2, maximumFractionDigits: 2 })}</span>
                                    </div>
                                    <div className="flex flex-col">
                                        <span className="text-gray-500">Vol (24h)</span>
                                        <span className="text-white font-semibold">${parseInt(stats.vol || '0').toLocaleString()}</span>
                                    </div>
                                </div>

                                {/* RIGHT: Chart Controls + Timeframe */}
                                <div className="flex items-center gap-2">
                                    {/* Chart Type Selector */}
                                    <div className="flex bg-white/5 p-0.5 rounded-lg">
                                        <button
                                            onClick={() => setChartType('area')}
                                            className={`p-1 rounded-md transition-all ${chartType === 'area'
                                                ? 'bg-blue-600 text-white shadow-lg'
                                                : 'text-gray-400 hover:text-white hover:bg-white/5'
                                                }`}
                                            title="Area Chart"
                                        >
                                            <LineChart className="w-3.5 h-3.5" />
                                        </button>
                                        <button
                                            onClick={() => setChartType('candle')}
                                            className={`p-1 rounded-md transition-all ${chartType === 'candle'
                                                ? 'bg-blue-600 text-white shadow-lg'
                                                : 'text-gray-400 hover:text-white hover:bg-white/5'
                                                }`}
                                            title="Candlestick Chart"
                                        >
                                            <CandlestickChart className="w-3.5 h-3.5" />
                                        </button>
                                    </div>

                                    {/* Timeframe Selector */}
                                    <div className="flex gap-0.5">
                                        {['5m', '15m', '1h', '4h', '1d', '1w'].map((tf) => (
                                            <button
                                                key={tf}
                                                onClick={() => setTimeframe(tf)}
                                                className={`px-2 py-0.5 rounded text-[10px] font-medium transition-all ${timeframe === tf
                                                    ? 'bg-blue-600 text-white'
                                                    : 'bg-white/5 text-gray-400 hover:bg-white/10 hover:text-white'
                                                    }`}
                                            >
                                                {tf}
                                            </button>
                                        ))}
                                    </div>
                                </div>
                            </div>

                            <div className="flex-1 relative">
                                <div className="flex-1 relative">
                                    <ProfessionalChart symbol={selectedSymbol} chartType={chartType} interval={timeframe} />
                                </div>
                            </div>
                        </div>

                        {/* Description / About Section */}
                        <div className="bg-[#111] border border-white/5 rounded-2xl p-6">
                            <h2 className="text-xl font-bold text-white mb-4">About {selectedCoinName}</h2>
                            <p className="text-gray-400 text-sm leading-relaxed mb-4">
                                {selectedCoinName} is the world's first cryptocurrency, a digital asset that uses public-key cryptography to record, sign and send transactions over the Bitcoin blockchain - all done without the oversight of a central authority.
                            </p>

                            <div className="grid grid-cols-1 md:grid-cols-3 gap-4 mt-6 pt-6 border-t border-white/5">
                                <div className="p-4 bg-white/5 rounded-xl">
                                    <h4 className="text-gray-500 text-xs mb-1">Dominance</h4>
                                    <p className="text-white text-lg font-bold">{coinData.dominance > 0 ? `${coinData.dominance.toFixed(1)}%` : 'N/A'}</p>
                                </div>
                                <div className="p-4 bg-white/5 rounded-xl">
                                    <h4 className="text-gray-500 text-xs mb-1">Rank</h4>
                                    <p className="text-white text-lg font-bold">#{coinData.rank || 'N/A'}</p>
                                </div>
                                <div className="p-4 bg-white/5 rounded-xl">
                                    <h4 className="text-gray-500 text-xs mb-1">All Time High</h4>
                                    <p className="text-white text-lg font-bold">${coinData.ath > 0 ? coinData.ath.toLocaleString('en-US', { minimumFractionDigits: 2, maximumFractionDigits: 2 }) : 'N/A'}</p>
                                    <p className={`text-xs -mt-1 ${coinData.athChangePercent >= 0 ? 'text-green-400' : 'text-red-400'}`}>
                                        {coinData.athChangePercent !== 0 ? `${coinData.athChangePercent.toFixed(1)}%` : 'N/A'}
                                    </p>
                                </div>
                            </div>
                        </div>
                    </div>

                    {/* RIGHT: Coin List */}
                    <div className="lg:col-span-2">
                        <div className="sticky top-4 bg-[#111] border border-white/5 rounded-2xl overflow-hidden" style={{ maxHeight: 'calc(100vh - 120px)' }}>
                            <CoinList
                                selectedSymbol={selectedSymbol}
                                onCoinSelect={(symbol, name) => {
                                    setSelectedSymbol(symbol);
                                    setSelectedCoinName(name);
                                }}
                            />
                        </div>
                    </div>
                </div>
            </main>
        </>
    );
}

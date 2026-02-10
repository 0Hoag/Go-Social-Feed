"use client";

import { useState, useEffect, useCallback } from "react";
import dynamic from "next/dynamic";
import { LineChart, CandlestickChart, ArrowUp, ArrowDown } from "lucide-react";

const ProfessionalChart = dynamic(() => import("@/components/ProfessionalChart"), {
    ssr: false,
    loading: () => <div className="animate-pulse bg-white/5 w-full h-full rounded-lg" />
});

import CoinList from "@/components/CoinList";
import LoadingScreen from "@/components/LoadingScreen";
import AnalysisLeftSidebar from "@/components/AnalysisLeftSidebar";
import AIAnalysisChat from "@/components/AIAnalysisChat";

import axios from "axios";
import { getMetadata } from "@/lib/tokenConstants";

import { useLanguage } from "@/context/LanguageContext";
import { fallbackDescriptions } from "@/utils/coinDescriptions";
import CoinDescription from "@/components/CoinDescription";

export default function AnalysisPage() {
    const { t, language } = useLanguage();
    const [isLoading, setIsLoading] = useState(true);
    const [isDataReady, setIsDataReady] = useState(false);
    const [selectedSymbol, setSelectedSymbol] = useState("BTCUSDT");
    const [selectedCoinName, setSelectedCoinName] = useState("Bitcoin");
    const [currentPrice, setCurrentPrice] = useState(65758.13);
    const [priceChange, setPriceChange] = useState(-6.78);

    const [stats, setStats] = useState({ high: '0.00', low: '0.00', vol: '0.00' });
    const [chartType, setChartType] = useState<'candle' | 'area'>('candle');
    const [timeframe, setTimeframe] = useState<string>('15m');
    const [coinData, setCoinData] = useState({ ath: 0, athChangePercent: 0, rank: 0, dominance: 0, description: { en: '', vi: '' } });
    const [isAIChatOpen, setIsAIChatOpen] = useState(false);

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
                    dominance: data.market_data.market_cap_percentage?.btc || 0,
                    description: {
                        en: (data.description?.en || '').replace(/<[^>]*>?/gm, '').trim(),
                        vi: fallbackDescriptions[selectedSymbol.replace('USDT', '').toLowerCase()] || ((data.description?.vi || '').replace(/<[^>]*>?/gm, '').trim()) || ''
                    }
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

    // Memoized handler for coin selection
    const handleCoinSelect = useCallback((symbol: string, name: string) => {
        setSelectedSymbol(symbol);
        setSelectedCoinName(name);
    }, []);

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
                        <div className="bg-gradient-to-b from-[#111] via-[#050505] to-[#000] border border-white/5 rounded-2xl overflow-hidden h-[600px] flex flex-col shadow-2xl shadow-black/50">
                            {/* Header Section (Unified) */}
                            <div className="flex flex-col border-b border-white/5 bg-white/[0.02] backdrop-blur-md">
                                {/* Top Row: Identity & Price */}
                                <div className="flex items-center justify-between px-6 pt-5 pb-3">
                                    {/* Left: Coin Identity */}
                                    <div className="flex items-center gap-5">
                                        <div className="relative group">
                                            <div className="absolute -inset-2 bg-gradient-to-r from-blue-500/20 to-teal-500/20 rounded-full opacity-0 group-hover:opacity-100 blur-md transition duration-500"></div>
                                            <img
                                                src={`https://assets.coincap.io/assets/icons/${selectedSymbol.replace('USDT', '').toLowerCase()}@2x.png`}
                                                alt={selectedCoinName}
                                                className="relative w-12 h-12 rounded-full shadow-lg ring-1 ring-white/10"
                                                onError={(e) => {
                                                    const target = e.currentTarget;
                                                    if (target.src.includes('coincap.io')) {
                                                        target.src = `https://cryptologos.cc/logos/${selectedCoinName.toLowerCase().replace(/\s+/g, '-')}-${selectedSymbol.replace('USDT', '').toLowerCase()}-logo.png`;
                                                    } else if (target.src.includes('cryptologos.cc')) {
                                                        target.src = `https://ui-avatars.com/api/?name=${selectedSymbol.replace('USDT', '')}&background=f7931a&color=fff&size=32&bold=true`;
                                                    }
                                                }}
                                            />
                                        </div>
                                        <div>
                                            <div className="flex items-baseline gap-2">
                                                <h1 className="text-3xl font-bold text-white tracking-tight drop-shadow-sm">{selectedCoinName}</h1>
                                                <span className="text-xs font-bold text-gray-500 bg-white/5 px-2 py-0.5 rounded-md border border-white/5">{selectedSymbol.replace('USDT', '')}</span>
                                            </div>
                                        </div>
                                    </div>

                                    {/* Right: Price */}
                                    <div className="flex flex-col items-end">
                                        <div className="flex items-center gap-3">
                                            <span className="text-3xl font-bold text-white tracking-tight">
                                                ${currentPrice.toLocaleString('en-US', { minimumFractionDigits: 2, maximumFractionDigits: 2 })}
                                            </span>
                                            <span className={`flex items-center gap-1 px-2.5 py-1 rounded-lg text-sm font-bold border ${priceChange >= 0 ? 'bg-emerald-500/5 text-emerald-400 border-emerald-500/20' : 'bg-rose-500/5 text-rose-400 border-rose-500/20'}`}>
                                                {priceChange >= 0 ? <ArrowUp className="w-3.5 h-3.5" /> : <ArrowDown className="w-3.5 h-3.5" />}
                                                {Math.abs(priceChange).toFixed(2)}%
                                            </span>
                                        </div>
                                    </div>
                                </div>

                                {/* Bottom Row: Stats & Controls (Merged) */}
                                <div className="flex items-center justify-between px-6 pb-4">
                                    {/* Left: AI & Stats */}
                                    <div className="flex items-center gap-6">
                                        <button
                                            onClick={() => setIsAIChatOpen(true)}
                                            className="hidden md:flex items-center gap-2 pl-3 pr-4 py-1.5 bg-gradient-to-r from-indigo-500/10 via-purple-500/10 to-indigo-500/10 border border-indigo-500/20 hover:border-indigo-400/40 rounded-full transition-all hover:shadow-[0_0_15px_rgba(99,102,241,0.15)] group relative overflow-hidden"
                                        >
                                            <div className="absolute inset-0 bg-gradient-to-r from-indigo-500/0 via-white/5 to-indigo-500/0 translate-x-[-100%] group-hover:translate-x-[100%] transition-transform duration-1000"></div>
                                            <div className="w-5 h-5 rounded-full bg-indigo-500/20 flex items-center justify-center group-hover:scale-110 transition-transform">
                                                <svg className="w-3 h-3 text-indigo-300" fill="none" viewBox="0 0 24 24" stroke="currentColor">
                                                    <path strokeLinecap="round" strokeLinejoin="round" strokeWidth={2} d="M13 10V3L4 14h7v7l9-11h-7z" />
                                                </svg>
                                            </div>
                                            <span className="text-xs font-semibold text-indigo-200 group-hover:text-white transition-colors">
                                                {t.analysis_page.ask_ai
                                                    .replace('{coin}', selectedCoinName)
                                                    .replace('{trend}', priceChange < 0 ? (language === 'vi' ? 'giảm' : 'decreasing') : (language === 'vi' ? 'tăng' : 'increasing'))}
                                            </span>
                                        </button>

                                        <div className="h-6 w-px bg-white/10 hidden md:block"></div>

                                        <div className="hidden lg:flex items-center gap-6 text-xs">
                                            <div className="flex flex-col gap-0.5">
                                                <span className="text-gray-500 font-medium text-[10px] uppercase tracking-wider">{t.analysis_page.high_24h}</span>
                                                <span className="text-white font-medium tracking-wide">${parseFloat(stats.high || '0').toLocaleString('en-US', { minimumFractionDigits: 2, maximumFractionDigits: 2 })}</span>
                                            </div>
                                            <div className="flex flex-col gap-0.5">
                                                <span className="text-gray-500 font-medium text-[10px] uppercase tracking-wider">{t.analysis_page.low_24h}</span>
                                                <span className="text-white font-medium tracking-wide">${parseFloat(stats.low || '0').toLocaleString('en-US', { minimumFractionDigits: 2, maximumFractionDigits: 2 })}</span>
                                            </div>
                                            <div className="flex flex-col gap-0.5">
                                                <span className="text-gray-500 font-medium text-[10px] uppercase tracking-wider">{t.analysis_page.vol_24h}</span>
                                                <span className="text-white font-medium tracking-wide">${parseInt(stats.vol || '0').toLocaleString()}</span>
                                            </div>
                                        </div>
                                    </div>

                                    {/* Right: Chart Controls */}
                                    <div className="flex items-center gap-3">
                                        <div className="flex bg-[#000]/50 p-1 rounded-lg border border-white/5">
                                            <button
                                                onClick={() => setChartType('area')}
                                                className={`p-1.5 rounded transition-all ${chartType === 'area' ? 'bg-white/10 text-white shadow-sm' : 'text-gray-500 hover:text-gray-300'}`}
                                            >
                                                <LineChart className="w-4 h-4" />
                                            </button>
                                            <button
                                                onClick={() => setChartType('candle')}
                                                className={`p-1.5 rounded transition-all ${chartType === 'candle' ? 'bg-white/10 text-white shadow-sm' : 'text-gray-500 hover:text-gray-300'}`}
                                            >
                                                <CandlestickChart className="w-4 h-4" />
                                            </button>
                                        </div>

                                        <div className="flex gap-1 p-1 bg-[#000]/50 rounded-lg border border-white/5">
                                            {['5m', '15m', '1h', '4h', '1d', '1w'].map((tf) => (
                                                <button
                                                    key={tf}
                                                    onClick={() => setTimeframe(tf)}
                                                    className={`px-3 py-1 rounded text-[11px] font-bold transition-all ${timeframe === tf
                                                        ? 'bg-gradient-to-r from-blue-600 to-blue-500 text-white shadow-lg shadow-blue-500/20'
                                                        : 'text-gray-500 hover:text-gray-300 hover:bg-white/5'
                                                        }`}
                                                >
                                                    {tf}
                                                </button>
                                            ))}
                                        </div>
                                    </div>
                                </div>
                            </div>

                            <div className="flex-1 relative">
                                <ProfessionalChart symbol={selectedSymbol} chartType={chartType} interval={timeframe} />
                            </div>
                        </div>

                        {/* Description / About Section */}
                        {/* Description / About Section */}
                        <CoinDescription
                            description={coinData.description[language] || coinData.description['en'] || (language === 'vi' ? `Thông tin về ${selectedCoinName} đang được cập nhật.` : `${selectedCoinName} information is being updated.`)}
                            coinName={selectedCoinName}
                        />

                        <div className="grid grid-cols-1 md:grid-cols-3 gap-4 mt-4">
                            <div className="p-4 bg-[#111] border border-white/5 rounded-2xl">
                                <h4 className="text-gray-500 text-xs mb-1">{t.analysis_page.dominance}</h4>
                                <p className="text-white text-lg font-bold">{coinData.dominance > 0 ? `${coinData.dominance.toFixed(1)}%` : 'N/A'}</p>
                            </div>
                            <div className="p-4 bg-[#111] border border-white/5 rounded-2xl">
                                <h4 className="text-gray-500 text-xs mb-1">{t.analysis_page.rank}</h4>
                                <p className="text-white text-lg font-bold">#{coinData.rank || 'N/A'}</p>
                            </div>
                            <div className="p-4 bg-[#111] border border-white/5 rounded-2xl">
                                <h4 className="text-gray-500 text-xs mb-1">{t.analysis_page.ath}</h4>
                                <p className="text-white text-lg font-bold">${coinData.ath > 0 ? coinData.ath.toLocaleString('en-US', { minimumFractionDigits: 2, maximumFractionDigits: 2 }) : 'N/A'}</p>
                                <p className={`text-xs -mt-1 ${coinData.athChangePercent >= 0 ? 'text-green-400' : 'text-red-400'}`}>
                                    {coinData.athChangePercent !== 0 ? `${coinData.athChangePercent.toFixed(1)}%` : 'N/A'}
                                </p>
                            </div>
                        </div>

                    </div>

                    {/* RIGHT: Coin List */}
                    <div className="lg:col-span-2">
                        <div className="sticky top-4 bg-[#111] border border-white/5 rounded-2xl overflow-hidden" style={{ maxHeight: 'calc(100vh - 120px)' }}>
                            <CoinList
                                selectedSymbol={selectedSymbol}
                                onCoinSelect={handleCoinSelect}
                            />
                        </div>
                    </div>
                </div >

                {/* AI Analysis Chat Overlay */}
                < AIAnalysisChat
                    isOpen={isAIChatOpen}
                    onClose={() => setIsAIChatOpen(false)}
                    coinSymbol={selectedSymbol}
                    coinName={selectedCoinName}
                    currentPrice={currentPrice}
                    priceChange={priceChange}
                />
            </main >
        </>
    );
}

"use client";

import { useEffect, useRef, useState } from 'react';
import { createChart, ColorType, LineStyle, ISeriesApi, IChartApi, Time } from 'lightweight-charts';

interface ProfessionalChartProps {
    symbol?: string;
}

interface CandleData {
    time: number;
    open: number;
    high: number;
    low: number;
    close: number;
    volume: number;
}

export default function ProfessionalChart({ symbol = "BTCUSDT" }: ProfessionalChartProps) {
    const chartContainerRef = useRef<HTMLDivElement>(null);
    const chartRef = useRef<IChartApi | null>(null);
    const candlestickSeriesRef = useRef<ISeriesApi<"Candlestick"> | null>(null);
    const volumeSeriesRef = useRef<ISeriesApi<"Histogram"> | null>(null);
    const ema7SeriesRef = useRef<ISeriesApi<"Line"> | null>(null);
    const ema25SeriesRef = useRef<ISeriesApi<"Line"> | null>(null);
    const ema99SeriesRef = useRef<ISeriesApi<"Line"> | null>(null);
    const ma7SeriesRef = useRef<ISeriesApi<"Line"> | null>(null);
    const ma25SeriesRef = useRef<ISeriesApi<"Line"> | null>(null);

    const [interval, setInterval] = useState<string>('15m');
    const currentPriceRef = useRef<string>("0.00");
    const [currentPrice, _setCurrentPrice] = useState("0.00");
    const setCurrentPrice = (price: string) => {
        currentPriceRef.current = price;
        _setCurrentPrice(price);
    };
    const [priceChange, setPriceChange] = useState<number>(0);

    const timeframes = [
        { label: '5m', value: '5m', limit: 1000, total: 5000 },   // Fetch 5000 candles total
        { label: '15m', value: '15m', limit: 1000, total: 3000 }, // Fetch 3000 candles total
        { label: '1h', value: '1h', limit: 1000, total: 2000 },   // Fetch 2000 candles total
        { label: '4h', value: '4h', limit: 1000, total: 1000 },
        { label: '1d', value: '1d', limit: 1000, total: 1000 },
        { label: '1w', value: '1w', limit: 1000, total: 1000 },
    ];

    // Calculate Moving Averages
    const calculateMA = (data: CandleData[], period: number) => {
        return data.map((item, index) => {
            if (index < period - 1) return { time: item.time, value: null };
            const sum = data.slice(index - period + 1, index + 1).reduce((acc, d) => acc + d.close, 0);
            return { time: item.time, value: sum / period };
        });
    };

    // Helper to calculate EMA
    const calculateEMA = (data: CandleData[], count: number) => {
        const k = 2 / (count + 1);
        const emaData = [];
        let ema = data.length > 0 ? data[0].close : 0; // Initialize with first close price or 0 if no data

        for (let i = 0; i < data.length; i++) {
            ema = data[i].close * k + ema * (1 - k);
            emaData.push({ time: data[i].time, value: ema });
        }
        return emaData;
    };

    useEffect(() => {
        if (!chartContainerRef.current) return;

        // Create chart
        const chart = createChart(chartContainerRef.current, {
            width: chartContainerRef.current.clientWidth,
            height: 500,
            layout: {
                background: { color: '#0a0a0a' },
                textColor: '#999',
            },
            grid: {
                vertLines: { color: '#1a1a1a', style: LineStyle.Solid },
                horzLines: { color: '#1a1a1a', style: LineStyle.Solid },
            },
            crosshair: {
                vertLine: {
                    color: '#666',
                    width: 1,
                    style: LineStyle.Solid,
                    labelBackgroundColor: '#333',
                },
                horzLine: {
                    color: '#666',
                    width: 1,
                    style: LineStyle.Solid,
                    labelBackgroundColor: '#333',
                    labelVisible: false, // Disable native label to use custom one
                },
            },
            timeScale: {
                timeVisible: true,
                secondsVisible: false,
                borderColor: '#333',
                lockVisibleTimeRangeOnResize: true,  // Prevent auto-scale on resize
                fixLeftEdge: false,                   // Allow scrolling left
                fixRightEdge: false,                  // Allow scrolling right
                rightOffset: 12, // Add 12 bars of empty space to the right
            },
            rightPriceScale: {
                borderColor: '#333',
                scaleMargins: {
                    top: 0.1,
                    bottom: 0.2, // Leave space for volume
                },
            },
            watermark: {
                visible: false,
            },
            // Disable zoom on scroll, enable pan on scroll
            handleScale: {
                mouseWheel: false,
            },
            handleScroll: {
                mouseWheel: true,
                pressedMouseMove: true,
                horzTouchDrag: true,
                vertTouchDrag: true,
            },
        });

        chartRef.current = chart;

        // Add candlestick series
        const candlestickSeries = chart.addCandlestickSeries({
            upColor: '#26a69a',
            downColor: '#ef5350',
            borderUpColor: '#26a69a',
            borderDownColor: '#ef5350',
            wickUpColor: '#26a69a',
            wickDownColor: '#ef5350',
        });
        candlestickSeriesRef.current = candlestickSeries;

        // Volume Series
        const volumeSeries = chart.addHistogramSeries({
            priceScaleId: 'volume',
            priceFormat: { type: 'volume' },
        });
        volumeSeriesRef.current = volumeSeries;

        chart.priceScale('volume').applyOptions({
            scaleMargins: {
                top: 0.85, // Place volume at the bottom 8% (very small)
                bottom: 0,
            },
        });

        // EMA Series
        const ema7 = chart.addLineSeries({
            color: 'rgba(251, 140, 0, 0.5)',
            lineWidth: 1,
            priceScaleId: 'right',
            lastValueVisible: false,
            priceLineVisible: false,
            crosshairMarkerVisible: false,
        }); // Yellow
        const ema25 = chart.addLineSeries({
            color: 'rgba(41, 98, 255, 0.5)',
            lineWidth: 1,
            priceScaleId: 'right',
            lastValueVisible: false,
            priceLineVisible: false,
            crosshairMarkerVisible: false,
        }); // Blue
        const ema99 = chart.addLineSeries({
            color: 'rgba(224, 64, 251, 0.5)',
            lineWidth: 1,
            priceScaleId: 'right',
            lastValueVisible: false,
            priceLineVisible: false,
            crosshairMarkerVisible: false,
        }); // Purple

        ema7SeriesRef.current = ema7;
        ema25SeriesRef.current = ema25;
        ema99SeriesRef.current = ema99;

        // Handle resize
        const handleResize = () => {
            if (chartContainerRef.current) {
                chart.applyOptions({
                    width: chartContainerRef.current.clientWidth,
                });
            }
        };

        window.addEventListener('resize', handleResize);

        const handleCrosshairMove = (param: any) => {
            // Check if point is valid (remove !param.time check to allow empty space)
            if (
                param.point === undefined ||
                param.point.x < 0 ||
                param.point.x > chartContainerRef.current!.clientWidth ||
                param.point.y < 0 ||
                param.point.y > chartContainerRef.current!.clientHeight
            ) {
                setCursorData(null);
                return;
            }

            const price = candlestickSeriesRef.current!.coordinateToPrice(param.point.y);
            if (price !== null) {
                const currentPriceVal = parseFloat(currentPriceRef.current);
                // Fallback to avoid weird display if currentPrice is 0 (initial state)
                if (isNaN(currentPriceVal) || currentPriceVal === 0) {
                    setCursorData({
                        visible: true,
                        x: param.point.x,
                        y: param.point.y,
                        price: price.toFixed(2),
                        percentDiff: '0,00%'
                    });
                    return;
                }

                const diff = ((price - currentPriceVal) / currentPriceVal) * 100;
                setCursorData({
                    visible: true,
                    x: param.point.x,
                    y: param.point.y,
                    price: price.toFixed(2),
                    percentDiff: (diff > 0 ? '+' : '') + diff.toFixed(2).replace('.', ',') + '%'
                });
            } else {
                setCursorData(null);
            }
        };

        chart.subscribeCrosshairMove(handleCrosshairMove);


        return () => {
            window.removeEventListener('resize', handleResize);
            chart.unsubscribeCrosshairMove(handleCrosshairMove);
            chart.remove();
        };
    }, []); // Only run once on mount (technically depends on nothing)

    const [stats, setStats] = useState({ high: '0.00', low: '0.00', vol: '0.00' });
    const [cursorData, setCursorData] = useState<{ visible: boolean; x: number; y: number; price: string; percentDiff: string } | null>(null);
    // Data Fetching
    useEffect(() => {
        const fetchData = async () => {
            // ... existing kline fetch logic ...
            try {
                // Fetch 24h stats
                const statsRes = await fetch(`https://api.binance.com/api/v3/ticker/24hr?symbol=${symbol}`);
                const statsData = await statsRes.json();
                setStats({
                    high: parseFloat(statsData.highPrice).toFixed(2),
                    low: parseFloat(statsData.lowPrice).toFixed(2),
                    vol: parseFloat(statsData.volume).toFixed(2)
                });

                const timeframe = timeframes.find(tf => tf.value === interval);
                const limit = timeframe?.limit || 1000;
                // const total = timeframe?.total || 1000;
                // const batches = Math.ceil(total / limit);

                // let allData: any[] = [];
                // let endTime: number | undefined = undefined;

                // Fetch multiple batches to get more historical data
                // for (let i = 0; i < batches; i++) {
                const url: string = `https://api.binance.com/api/v3/klines?symbol=${symbol}&interval=${interval}&limit=${limit}`;
                // ? `https://api.binance.com/api/v3/klines?symbol=${symbol}&interval=${interval}&limit=${limit}&endTime=${endTime}`
                // : `https://api.binance.com/api/v3/klines?symbol=${symbol}&interval=${interval}&limit=${limit}`;

                const response: Response = await fetch(url);
                const data: any[] = await response.json();

                // if (data.length === 0) break;

                // allData = [...data, ...allData]; // Prepend older data
                // endTime = data[0][0] - 1; // Set endTime for next batch

                // await new Promise(resolve => setTimeout(resolve, 100)); // Rate limit delay
                // }

                const formattedData: CandleData[] = data.map((item: any) => ({
                    time: Math.floor(item[0] / 1000) + 25200, // Add 7 hours for UTC+7 (Vietnam)
                    open: parseFloat(item[1]),
                    high: parseFloat(item[2]),
                    low: parseFloat(item[3]),
                    close: parseFloat(item[4]),
                    volume: parseFloat(item[5]),
                }));

                const volumeData = data.map((item: any) => ({
                    time: Math.floor(item[0] / 1000) + 25200,
                    value: parseFloat(item[5]),
                    color: parseFloat(item[4]) >= parseFloat(item[1]) ? 'rgba(38, 166, 154, 0.5)' : 'rgba(239, 83, 80, 0.5)',
                }));

                if (candlestickSeriesRef.current && volumeSeriesRef.current) {
                    candlestickSeriesRef.current.setData(formattedData as any);
                    volumeSeriesRef.current.setData(volumeData as any);

                    // Set EMA Data
                    if (formattedData.length > 0) {
                        const ema7Data = calculateEMA(formattedData, 7);
                        const ema25Data = calculateEMA(formattedData, 25);
                        const ema99Data = calculateEMA(formattedData, 99);

                        ema7SeriesRef.current?.setData(ema7Data as any);
                        ema25SeriesRef.current?.setData(ema25Data as any);
                        ema99SeriesRef.current?.setData(ema99Data as any);

                        // Add 10 empty candles for future grid
                        const lastTime = formattedData[formattedData.length - 1].time as number;
                        const intervalSeconds =
                            interval === '5m' ? 300 :
                                interval === '15m' ? 900 :
                                    interval === '1h' ? 3600 :
                                        interval === '4h' ? 14400 :
                                            interval === '1d' ? 86400 : 604800; // 1w

                        for (let i = 1; i <= 10; i++) {
                            candlestickSeriesRef.current.update({ time: (lastTime + (i * intervalSeconds)) as any } as any);
                        }
                    }
                }

                if (formattedData.length > 0) {
                    const latest = formattedData[formattedData.length - 1];
                    const first = formattedData[0];
                    setCurrentPrice(latest.close.toFixed(2));
                    const change = ((latest.close - first.close) / first.close) * 100;
                    setPriceChange(change);
                }

                // Auto-scroll logic...
                if (chartRef.current && formattedData.length > 0) {
                    const visibleCandles = 100;
                    const latestTime = formattedData[formattedData.length - 1].time;
                    const startIndex = Math.max(0, formattedData.length - visibleCandles);
                    const startTime = formattedData[startIndex].time;

                    // 1. Set zoom level
                    chartRef.current.timeScale().setVisibleRange({
                        from: startTime as any,
                        to: latestTime as any,
                    });

                    // 2. Shift view to center
                    chartRef.current.timeScale().scrollToPosition(50, false);
                }
            } catch (error) {
                console.error('Error fetching data:', error);
            }
        };

        fetchData();

        // Setup WebSocket for kline updates
        const wsUrl = `wss://stream.binance.com:9443/ws/${symbol.toLowerCase()}@kline_${interval}`;
        const ws = new WebSocket(wsUrl);

        ws.onmessage = (event) => {
            const message = JSON.parse(event.data);
            const candle = message.k;

            if (candle && candlestickSeriesRef.current) {
                const newCandle = {
                    time: (Math.floor(candle.t / 1000) + 25200) as any,
                    open: parseFloat(candle.o),
                    high: parseFloat(candle.h),
                    low: parseFloat(candle.l),
                    close: parseFloat(candle.c),
                };

                try {
                    candlestickSeriesRef.current.update(newCandle);
                } catch (e) {
                    // Ignore "Cannot update oldest data" errors during race conditions
                    console.warn("Chart update skipped:", e);
                }
                setCurrentPrice(newCandle.close.toFixed(2));

                // Update Volume
                if (volumeSeriesRef.current) {
                    volumeSeriesRef.current.update({
                        time: newCandle.time,
                        value: parseFloat(candle.v),
                        color: newCandle.close >= newCandle.open ? 'rgba(38, 166, 154, 0.5)' : 'rgba(239, 83, 80, 0.5)',
                    });
                }

                // Note: Updating EMAs in real-time accurately requires the full history or storing the last EMA state. 
                // For simplicity/visual smoothness, we can just let it re-fetch on interval change or implement a simplified incremental update if needed.
                // But for now, since we have the full `formattedData` from fetch, we can't easily append without recalculating state.
                // A full production app would manage standard indicators state more consistently. 
                // Given the constraint, we will skip incremental EMA updates for this specific minimal scope change unless requested, 
                // as frequent calculating on every tick might be overkill or require state refactoring.
                // However, let's at least try to update the last point if we tracked the last EMA data.
            }
        };

        // Setup WebSocket for 24h ticker updates
        const tickerWsUrl = `wss://stream.binance.com:9443/ws/${symbol.toLowerCase()}@ticker`;
        const tickerWs = new WebSocket(tickerWsUrl);

        tickerWs.onmessage = (event) => {
            const ticker = JSON.parse(event.data);
            if (ticker) {
                setStats({
                    high: parseFloat(ticker.h).toFixed(2),
                    low: parseFloat(ticker.l).toFixed(2),
                    vol: parseFloat(ticker.v).toFixed(2)
                });
            }
        };

        return () => {
            ws.close();
            tickerWs.close();
        };
    }, [symbol, interval]);

    const isPositive = priceChange >= 0;

    const formatPrice = (value: string | number) => {
        const val = typeof value === 'string' ? parseFloat(value) : value;
        if (isNaN(val)) return '0,00';
        return val.toLocaleString('vi-VN', { minimumFractionDigits: 2, maximumFractionDigits: 2 });
    };

    return (
        <div className="relative w-full h-full bg-gradient-to-b from-[#0a0a0a] to-[#050505] p-6">
            {/* Header */}
            <div className="flex items-center justify-between mb-4">
                <div className="flex items-center gap-8">
                    <div>
                        <div className="flex items-baseline gap-3">
                            <span className="text-4xl font-bold text-white">
                                ${formatPrice(currentPrice)}
                            </span>
                            <span className={`text-lg font-semibold ${isPositive ? 'text-[#26a69a]' : 'text-[#ef5350]'}`}>
                                {isPositive ? '+' : ''}{priceChange.toFixed(2)}%
                            </span>
                        </div>
                        <div className="text-sm text-gray-500 mt-1">
                            {symbol.replace('USDT', '')}/USDT
                        </div>
                    </div>

                    {/* 24h Stats */}
                    <div className="flex gap-8 text-sm">
                        <div className="flex flex-col">
                            <span className="text-gray-500 text-xs mb-1">24h High</span>
                            <span className="text-gray-200 font-medium">{formatPrice(stats.high)}</span>
                        </div>
                        <div className="flex flex-col">
                            <span className="text-gray-500 text-xs mb-1">24h Low</span>
                            <span className="text-gray-200 font-medium">{formatPrice(stats.low)}</span>
                        </div>
                        <div className="flex flex-col">
                            <span className="text-gray-500 text-xs mb-1">24h Vol</span>
                            <span className="text-gray-200 font-medium">{formatPrice(stats.vol)}</span>
                        </div>
                    </div>
                </div>

                {/* Timeframe Selector */}
                <div className="flex gap-2">
                    {timeframes.map((tf) => (
                        <button
                            key={tf.value}
                            onClick={() => setInterval(tf.value)}
                            className={`px-3 py-1 rounded text-xs font-medium transition-all ${interval === tf.value
                                ? 'bg-blue-600 text-white shadow-lg shadow-blue-600/50'
                                : 'bg-[#1a1a1a] text-gray-400 hover:bg-[#222] hover:text-white'
                                }`}
                        >
                            {tf.label}
                        </button>
                    ))}
                </div>
            </div>

            {/* Tooltip removed to prevent blocking candles */}

            {/* Custom Right Scale Label (Axis) */}
            {cursorData && cursorData.visible && (
                <div
                    className="absolute z-40 pointer-events-none bg-[#333] text-white text-[11px] font-mono px-1 flex items-center justify-center border-l-2 border-white/20"
                    style={{
                        right: 0,
                        top: cursorData.y + 95, // Direct centering: y - half_height
                        height: '20px',
                        minWidth: '60px',
                    }}
                >
                    {formatPrice(cursorData.price)} ({cursorData.percentDiff})
                </div>
            )}

            {/* Chart */}
            <style jsx global>{`
                /* Hide TradingView watermark and attribution */
                .tv-lightweight-charts > table > tr:nth-child(2) > td:nth-child(2) > a {
                    display: none !important;
                }
                .tv-lightweight-charts > div:last-child {
                    display: none !important;
                }
                a[href^="https://www.tradingview.com/"] {
                    display: none !important;
                }
                /* Additional backup selectors */
                div[style*="z-index: 3"] > a[href*="tradingview"] {
                    display: none !important;
                }
                .tv-lightweight-charts__watermark {
                    display: none !important;
                }
            `}</style>
            <div ref={chartContainerRef} className="w-full" />

            {/* Legend removed - MA lines disabled */}
        </div>
    );
}

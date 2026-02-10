"use client";

import { useEffect, useRef, useState, memo } from 'react';
import { createChart, ColorType, CrosshairMode, LineStyle, ISeriesApi, IChartApi, Time } from 'lightweight-charts';

interface ProfessionalChartProps {
    symbol?: string;
    chartType?: 'candle' | 'area';
    interval?: string;
}

interface CandleData {
    time: number;
    open: number;
    high: number;
    low: number;
    close: number;
    volume: number;
}

const ProfessionalChart = memo(({ symbol = "BTCUSDT", chartType = 'candle', interval: externalInterval = '15m' }: ProfessionalChartProps) => {
    const chartContainerRef = useRef<HTMLDivElement>(null);
    const chartRef = useRef<IChartApi | null>(null);
    const mainSeriesRef = useRef<ISeriesApi<"Candlestick" | "Area"> | null>(null);

    const ema7SeriesRef = useRef<ISeriesApi<"Line"> | null>(null);
    const ema25SeriesRef = useRef<ISeriesApi<"Line"> | null>(null);
    const ema99SeriesRef = useRef<ISeriesApi<"Line"> | null>(null);

    const interval = externalInterval;
    const currentPriceRef = useRef<string>("0.00");
    const [currentPrice, _setCurrentPrice] = useState("0.00");
    const setCurrentPrice = (price: string) => {
        currentPriceRef.current = price;
        _setCurrentPrice(price);
    };
    const [priceChange, setPriceChange] = useState<number>(0);
    const [cursorData, setCursorData] = useState<{ visible: boolean; x: number; y: number; price: string; percentDiff: string } | null>(null);
    const [isChartReady, setIsChartReady] = useState(false);

    const timeframes = [
        { label: '5m', value: '5m', limit: 1000 },
        { label: '15m', value: '15m', limit: 1000 },
        { label: '1h', value: '1h', limit: 1000 },
        { label: '4h', value: '4h', limit: 1000 },
        { label: '1d', value: '1d', limit: 1000 },
        { label: '1w', value: '1w', limit: 1000 },
    ];

    // Helper to calculate EMA
    const calculateEMA = (data: CandleData[], count: number) => {
        const k = 2 / (count + 1);
        const emaData = [];
        let ema = data.length > 0 ? data[0].close : 0;

        for (let i = 0; i < data.length; i++) {
            ema = data[i].close * k + ema * (1 - k);
            emaData.push({ time: data[i].time, value: ema });
        }
        return emaData;
    };

    // Initialize Chart
    useEffect(() => {
        if (!chartContainerRef.current) return;

        const createChartInstance = () => {
            if (chartRef.current || !chartContainerRef.current) return;

            const chart = createChart(chartContainerRef.current, {
                width: chartContainerRef.current.clientWidth,
                height: chartContainerRef.current.clientHeight,
                layout: {
                    background: { color: 'transparent' }, // Transparent to blend with gradient
                    textColor: '#999',
                },
                grid: {
                    vertLines: { color: 'rgba(42, 46, 57, 0.5)', style: LineStyle.Solid },
                    horzLines: { color: 'rgba(42, 46, 57, 0.5)', style: LineStyle.Solid },
                },
                crosshair: {
                    mode: CrosshairMode.Normal,
                    vertLine: {
                        color: '#666',
                        width: 1,
                        style: LineStyle.Solid,
                        labelBackgroundColor: '#333',
                        labelVisible: true,
                    },
                    horzLine: {
                        color: '#666',
                        width: 1,
                        style: LineStyle.Solid,
                        labelBackgroundColor: '#333',
                        labelVisible: false,
                    },
                },
                timeScale: {
                    timeVisible: true,
                    secondsVisible: false,
                    borderColor: '#333',
                    rightOffset: 12,
                },
                rightPriceScale: {
                    borderColor: '#333',
                },
                handleScale: {
                    mouseWheel: true, // Enable mouse wheel scaling
                },
                handleScroll: {
                    mouseWheel: true,
                    pressedMouseMove: true,
                    horzTouchDrag: true,
                    vertTouchDrag: true,
                },
            });

            chartRef.current = chart;

            // Add main series
            let mainSeries: ISeriesApi<"Candlestick" | "Area">;

            if (chartType === 'area') {
                mainSeries = chart.addAreaSeries({
                    topColor: 'rgba(38, 166, 154, 0.56)',
                    bottomColor: 'rgba(38, 166, 154, 0.04)',
                    lineColor: 'rgba(38, 166, 154, 1)',
                    lineWidth: 2,
                });
            } else {
                mainSeries = chart.addCandlestickSeries({
                    upColor: '#26a69a',
                    downColor: '#ef5350',
                    borderUpColor: '#26a69a',
                    borderDownColor: '#ef5350',
                    wickUpColor: '#26a69a',
                    wickDownColor: '#ef5350',
                });
            }
            mainSeriesRef.current = mainSeries;

            // EMA Series (Candle only)
            if (chartType === 'candle') {
                const emaColors = ['rgba(251, 140, 0, 0.5)', 'rgba(41, 98, 255, 0.5)', 'rgba(224, 64, 251, 0.5)'];
                const emaRefs = [ema7SeriesRef, ema25SeriesRef, ema99SeriesRef];

                emaRefs.forEach((ref, index) => {
                    ref.current = chart.addLineSeries({
                        color: emaColors[index],
                        lineWidth: 1,
                        priceScaleId: 'right',
                        lastValueVisible: false,
                        priceLineVisible: false,
                        crosshairMarkerVisible: false,
                    });
                });
            } else {
                ema7SeriesRef.current = null;
                ema25SeriesRef.current = null;
                ema99SeriesRef.current = null;
            }

            // Crosshair handler
            const handleCrosshairMove = (param: any) => {
                if (!param.point || !param.time || !chartContainerRef.current) {
                    setCursorData(null);
                    return;
                }

                const price = mainSeries.coordinateToPrice(param.point.y);
                if (price !== null) {
                    const currentPriceVal = parseFloat(currentPriceRef.current);
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
                }
            };

            chart.subscribeCrosshairMove(handleCrosshairMove);
            setIsChartReady(true); // Signal that chart is ready
        };

        // Handle resize with ResizeObserver
        const resizeObserver = new ResizeObserver((entries) => {
            if (!chartContainerRef.current || entries.length === 0) return;
            const newRect = entries[0].contentRect;

            if (newRect.width === 0 || newRect.height === 0) return;

            if (chartRef.current) {
                chartRef.current.applyOptions({
                    width: newRect.width,
                    height: newRect.height
                });
            } else {
                // Initialize chart if it doesn't exist and we have dimensions
                createChartInstance();
            }
        });

        resizeObserver.observe(chartContainerRef.current);

        // Immediate check in case ResizeObserver is lazy or dimensions are already present
        if (chartContainerRef.current.clientWidth > 0 && chartContainerRef.current.clientHeight > 0) {
            createChartInstance();
        }

        // cleanup
        return () => {
            resizeObserver.disconnect();
            if (chartRef.current) {
                chartRef.current.remove();
                chartRef.current = null;

                // Explicitly clear refs to avoid stale usage
                mainSeriesRef.current = null;
                ema7SeriesRef.current = null;
                ema25SeriesRef.current = null;
                ema99SeriesRef.current = null;

                setIsChartReady(false);
            }
        };
    }, [chartType]); // Re-init on chartType change

    // Data Fetching
    useEffect(() => {
        if (!isChartReady || !chartRef.current || !mainSeriesRef.current) return;

        const fetchData = async () => {
            try {
                const timeframe = timeframes.find(tf => tf.value === interval);
                const limit = timeframe?.limit || 1000;
                const url = `https://api.binance.com/api/v3/klines?symbol=${symbol}&interval=${interval}&limit=${limit}`;

                const response = await fetch(url);
                const data = await response.json();

                if (!Array.isArray(data)) return;

                const formattedData: CandleData[] = data.map((item: any) => ({
                    time: Math.floor(item[0] / 1000) + 25200, // UTC+7
                    open: parseFloat(item[1]),
                    high: parseFloat(item[2]),
                    low: parseFloat(item[3]),
                    close: parseFloat(item[4]),
                    volume: parseFloat(item[5]),
                }));

                if (chartType === 'area') {
                    const areaData = formattedData.map(d => ({ time: d.time, value: d.close }));
                    mainSeriesRef.current?.setData(areaData as any);
                } else {
                    mainSeriesRef.current?.setData(formattedData as any);
                }

                // EMA Data
                if (chartType === 'candle' && formattedData.length > 0) {
                    ema7SeriesRef.current?.setData(calculateEMA(formattedData, 7) as any);
                    ema25SeriesRef.current?.setData(calculateEMA(formattedData, 25) as any);
                    ema99SeriesRef.current?.setData(calculateEMA(formattedData, 99) as any);
                }

                if (formattedData.length > 0) {
                    const latest = formattedData[formattedData.length - 1];
                    setCurrentPrice(latest.close.toFixed(2));
                    const first = formattedData[0];
                    setPriceChange(((latest.close - first.close) / first.close) * 100);

                    // Force a specific zoom level using Logical Range (bars)
                    // limit to 50 bars visible, centered on the last bar
                    if (chartRef.current) {
                        const visibleCandles = 50;
                        const halfRange = visibleCandles / 2;

                        // Logical index of the last bar is effectively the count (relative to current view)
                        // logic: from (past) to (future)
                        // We want 25 bars empty space to the right.

                        chartRef.current.timeScale().setVisibleLogicalRange({
                            from: formattedData.length - halfRange,
                            to: formattedData.length + halfRange,
                        });
                    }
                }

            } catch (error) {
                console.error('Error fetching data:', error);
            }
        };

        fetchData();

        // WebSocket
        const wsUrl = `wss://stream.binance.com:9443/ws/${symbol.toLowerCase()}@kline_${interval}`;
        const ws = new WebSocket(wsUrl);

        ws.onmessage = (event) => {
            const message = JSON.parse(event.data);
            const candle = message.k;

            if (candle && mainSeriesRef.current) {
                const newCandle = {
                    time: (Math.floor(candle.t / 1000) + 25200) as any,
                    open: parseFloat(candle.o),
                    high: parseFloat(candle.h),
                    low: parseFloat(candle.l),
                    close: parseFloat(candle.c),
                };

                try {
                    if (chartType === 'area') {
                        mainSeriesRef.current.update({ time: newCandle.time, value: newCandle.close } as any);
                    } else {
                        (mainSeriesRef.current as ISeriesApi<"Candlestick">).update(newCandle);
                    }
                    setCurrentPrice(newCandle.close.toFixed(2));
                } catch (e) {
                    // console.warn("Chart update error", e);
                }
            }
        };

        return () => {
            ws.close();
        };

    }, [symbol, interval, chartType, isChartReady]); // Depend on chartRef.current and isChartReady

    const formatPrice = (value: string | number) => {
        const val = typeof value === 'string' ? parseFloat(value) : value;
        if (isNaN(val)) return '0,00';
        return val.toLocaleString('vi-VN', { minimumFractionDigits: 2, maximumFractionDigits: 2 });
    };

    return (
        <div className="relative w-full h-full bg-gradient-to-b from-[#0a0a0a] to-[#050505] p-6">
            {cursorData && cursorData.visible && (
                <div
                    className="absolute z-40 pointer-events-none bg-[#333] text-white text-[11px] font-mono px-1 flex items-center justify-center border-l-2 border-white/20"
                    style={{
                        right: 0,
                        top: cursorData.y + 14,
                        height: '20px',
                        minWidth: '60px',
                    }}
                >
                    {formatPrice(cursorData.price)} ({cursorData.percentDiff})
                </div>
            )}
            <style jsx global>{`
                a[href^="https://www.tradingview.com/"] { display: none !important; }
                .tv-lightweight-charts__watermark { display: none !important; }
            `}</style>
            {/* Added id for easier debugging if needed */}
            <div ref={chartContainerRef} id="tv_chart_container" className="w-full h-full" />
        </div>
    );
});

export default ProfessionalChart;

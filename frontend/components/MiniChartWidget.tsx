"use client";

import { useEffect, useRef, memo } from "react";

interface MiniChartWidgetProps {
    symbol: string;
}

function MiniChartWidget({ symbol }: MiniChartWidgetProps) {
    const container = useRef<HTMLDivElement>(null);

    useEffect(() => {
        if (container.current && !container.current.querySelector("script")) {
            const script = document.createElement("script");
            script.src = "https://s3.tradingview.com/external-embedding/embed-widget-mini-symbol-overview.js";
            script.type = "text/javascript";
            script.async = true;
            script.innerHTML = `
            {
              "symbol": "${symbol}",
              "width": "100%",
              "height": "220",
              "locale": "en",
              "dateRange": "1D",
              "colorTheme": "dark",
              "isTransparent": true,
              "autosize": false,
              "largeChartUrl": ""
            }`;
            container.current.appendChild(script);
        }
    }, [symbol]);

    return (
        <div className="tradingview-widget-container w-full bg-[#111] border border-white/5 rounded-2xl overflow-hidden mb-4 last:mb-0" ref={container}>
            <div className="tradingview-widget-container__widget"></div>
        </div>
    );
}

export default memo(MiniChartWidget);

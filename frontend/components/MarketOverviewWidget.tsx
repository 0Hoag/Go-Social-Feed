"use client";

import { useEffect, useRef, memo } from "react";

function MarketOverviewWidget() {
    const container = useRef<HTMLDivElement>(null);

    useEffect(() => {
        if (container.current && !container.current.querySelector("script")) {
            const script = document.createElement("script");
            script.src = "https://s3.tradingview.com/external-embedding/embed-widget-market-overview.js";
            script.type = "text/javascript";
            script.async = true;
            script.innerHTML = `
            {
              "colorTheme": "dark",
              "dateRange": "12M",
              "showChart": true,
              "locale": "en",
              "largeChartUrl": "",
              "isTransparent": true,
              "showSymbolLogo": true,
              "showFloatingTooltip": false,
              "width": "100%",
              "height": "100%",
              "plotLineColorGrowing": "rgba(41, 98, 255, 1)",
              "plotLineColorFalling": "rgba(41, 98, 255, 1)",
              "gridLineColor": "rgba(42, 46, 57, 0)",
              "scaleFontColor": "rgba(209, 212, 220, 1)",
              "belowLineFillColorGrowing": "rgba(41, 98, 255, 0.12)",
              "belowLineFillColorFalling": "rgba(41, 98, 255, 0.12)",
              "belowLineFillColorGrowingBottom": "rgba(41, 98, 255, 0)",
              "belowLineFillColorFallingBottom": "rgba(41, 98, 255, 0)",
              "symbolActiveColor": "rgba(41, 98, 255, 0.12)",
              "tabs": [
                {
                  "title": "Crypto",
                  "symbols": [
                    {
                      "s": "BINANCE:BTCUSDT",
                      "d": "Bitcoin"
                    },
                    {
                      "s": "BINANCE:ETHUSDT",
                      "d": "Ethereum"
                    },
                    {
                      "s": "BINANCE:SOLUSDT",
                      "d": "Solana"
                    },
                    {
                      "s": "BINANCE:BNBUSDT",
                      "d": "BNB"
                    },
                    {
                      "s": "BINANCE:XRPUSDT",
                      "d": "XRP"
                    },
                    {
                      "s": "BINANCE:DOGEUSDT",
                      "d": "Dogecoin"
                    },
                    {
                      "s": "BINANCE:ADAUSDT",
                      "d": "Cardano"
                    }
                  ],
                  "originalTitle": "Crypto"
                }
              ]
            }`;
            container.current.appendChild(script);
        }
    }, []);

    return (
        <div className="tradingview-widget-container h-full w-full" ref={container}>
            <div className="tradingview-widget-container__widget"></div>
        </div>
    );
}

export default memo(MarketOverviewWidget);

"use client";

import { useEffect, useState } from 'react';

interface LoadingScreenProps {
    onLoadingComplete?: () => void;
}

export default function LoadingScreen({ onLoadingComplete }: LoadingScreenProps) {
    const [progress, setProgress] = useState(0);
    const [messageIndex, setMessageIndex] = useState(0);
    const [isVisible, setIsVisible] = useState(true);

    const messages = [
        "Initializing...",
        "Connecting to exchanges...",
        "Loading market data...",
        "Preparing terminal..."
    ];

    useEffect(() => {
        // Progress animation
        const progressInterval = setInterval(() => {
            setProgress(prev => {
                if (prev >= 100) {
                    clearInterval(progressInterval);
                    return 100;
                }
                return prev + 2;
            });
        }, 30);

        // Message cycling
        const messageInterval = setInterval(() => {
            setMessageIndex(prev => (prev + 1) % messages.length);
        }, 800);

        // Minimum display time
        const minDisplayTimer = setTimeout(() => {
            if (onLoadingComplete) {
                onLoadingComplete();
            }
        }, 1500);

        return () => {
            clearInterval(progressInterval);
            clearInterval(messageInterval);
            clearTimeout(minDisplayTimer);
        };
    }, []);

    return (
        <div
            className={`fixed inset-0 z-50 flex items-center justify-center bg-[#0a0a0a] transition-opacity duration-500 ${isVisible ? 'opacity-100' : 'opacity-0 pointer-events-none'
                }`}
        >
            <div className="flex flex-col items-center gap-8">
                {/* Logo */}
                <div className="flex items-center gap-3">
                    <h1 className="text-5xl font-bold text-[#00d4ff]">
                        Syntax
                    </h1>
                    <span className="px-3 py-1 text-sm font-bold text-white bg-[#ff6b35] rounded">
                        PRO
                    </span>
                </div>

                {/* Subtitle */}
                <p className="text-sm tracking-[0.3em] text-gray-500 uppercase">
                    On-Chain Intelligence Terminal
                </p>

                {/* Progress Bar Container */}
                <div className="w-80 space-y-4">
                    {/* Progress Bar */}
                    <div className="h-1 bg-gray-800 rounded-full overflow-hidden">
                        <div
                            className="h-full bg-gradient-to-r from-[#00d4ff] to-[#0099cc] transition-all duration-300 ease-out"
                            style={{ width: `${progress}%` }}
                        />
                    </div>

                    {/* Loading Message */}
                    <p className="text-center text-sm text-gray-400 font-mono transition-opacity duration-300">
                        {messages[messageIndex]}
                    </p>
                </div>

                {/* Animated dots */}
                <div className="flex gap-2">
                    {[0, 1, 2].map((i) => (
                        <div
                            key={i}
                            className="w-2 h-2 bg-[#00d4ff] rounded-full animate-pulse"
                            style={{
                                animationDelay: `${i * 0.2}s`,
                                animationDuration: '1.4s'
                            }}
                        />
                    ))}
                </div>
            </div>
        </div>
    );
}

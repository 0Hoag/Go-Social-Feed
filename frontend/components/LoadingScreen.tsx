"use client";

import { useEffect, useState } from 'react';

interface LoadingScreenProps {
    onLoadingComplete?: () => void;
    isReady?: boolean;
}

export default function LoadingScreen({ onLoadingComplete, isReady = false }: LoadingScreenProps) {
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
                // If ready, speed up to 100
                if (isReady) {
                    const next = prev + 5;
                    if (next >= 100) {
                        return 100;
                    }
                    return next;
                }

                // If not ready, slow down at 90%
                if (prev >= 90) return prev;
                return prev + 1;
            });
        }, 30);

        // Message cycling
        const messageInterval = setInterval(() => {
            setMessageIndex(prev => (prev + 1) % messages.length);
        }, 800);

        return () => {
            clearInterval(progressInterval);
            clearInterval(messageInterval);
        };
    }, [isReady]);

    // Handle completion when progress hits 100
    useEffect(() => {
        if (progress === 100 && onLoadingComplete) {
            // Small delay to show 100%
            const timer = setTimeout(() => {
                onLoadingComplete();
            }, 500);
            return () => clearTimeout(timer);
        }
    }, [progress, onLoadingComplete]);

    return (
        <div
            className={`fixed inset-0 z-50 flex items-center justify-center bg-[#0a0a0a] transition-opacity duration-500 ${isVisible ? 'opacity-100' : 'opacity-0 pointer-events-none'
                }`}
        >
            <div className="flex flex-col items-center gap-8">
                {/* Logo */}
                <div className="flex items-center gap-3">
                    <h1 className="text-5xl font-bold text-[#00d4ff]">
                        CryptoCheck
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

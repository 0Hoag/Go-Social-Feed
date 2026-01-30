"use client";

import { useEffect, useState } from "react";
import { getPosts } from "@/lib/api"; // Reuse existing API
import { Post } from "@/lib/types";
import { truncateText } from "@/lib/utils";
import { Zap } from "lucide-react";
import Link from "next/link";

export default function NewsTicker() {
    const [headlines, setHeadlines] = useState<Post[]>([]);

    useEffect(() => {
        const fetchHeadlines = async () => {
            try {
                const res = await getPosts({ limit: 10 });
                setHeadlines(res.posts);
            } catch (error) {
                console.error("Failed to fetch headlines:", error);
            }
        };

        fetchHeadlines();
    }, []);

    if (headlines.length === 0) return null;

    return (
        <div className="w-full bg-cyan-950/20 border-b border-cyan-900/30 overflow-hidden py-2 flex items-center relative z-30">
            <div className="flex items-center gap-2 px-4 border-r border-cyan-500/20 shrink-0 bg-[#050510] z-10 text-xs font-bold text-cyan-500 uppercase tracking-wider">
                <Zap className="w-3 h-3 text-cyan-400 fill-cyan-400 animate-pulse" />
                Breaking News
            </div>

            {/* Ticker Container */}
            <div className="overflow-hidden whitespace-nowrap mask-linear-gradient w-full flex hover:pause-animation group">
                <div className="flex items-center gap-12 animate-marquee-slow shrink-0 pr-12 group-hover:[animation-play-state:paused]">
                    {/* Duplicate list to ensure it covers screen width */}
                    {[...headlines, ...headlines, ...headlines].map((post, index) => (
                        <Link
                            key={`orig-${post.id}-${index}`}
                            href={`/posts/${post.id}`}
                            className="text-xs text-gray-300 hover:text-cyan-400 transition-colors flex items-center gap-3"
                        >
                            <span className="w-1.5 h-1.5 rounded-full bg-white/20"></span>
                            {truncateText(post.title, 80)}
                        </Link>
                    ))}
                </div>
                {/* Second duplicated container for seamless loop */}
                <div className="flex items-center gap-12 animate-marquee-slow shrink-0 pr-12 group-hover:[animation-play-state:paused]" aria-hidden="true">
                    {[...headlines, ...headlines, ...headlines].map((post, index) => (
                        <Link
                            key={`copy-${post.id}-${index}`}
                            href={`/posts/${post.id}`}
                            className="text-xs text-gray-300 hover:text-cyan-400 transition-colors flex items-center gap-3"
                        >
                            <span className="w-1.5 h-1.5 rounded-full bg-white/20"></span>
                            {truncateText(post.title, 80)}
                        </Link>
                    ))}
                </div>
            </div>
            <style jsx>{`
                .hover\:pause-animation:hover .animate-marquee-slow {
                    animation-play-state: paused;
                }
            `}</style>
        </div>
    );
}

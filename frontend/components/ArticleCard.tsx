import Link from "next/link";
import Image from "next/image";
import { Post } from "@/lib/types";
import { formatDate, getSourceName, extractImageUrl } from "@/lib/utils";
import { Clock } from "lucide-react";
import { cn } from "@/lib/utils";

interface ArticleCardProps {
    post: Post;
    variant?: "default" | "compact";
    className?: string;
}

export default function ArticleCard({ post, variant = "default", className }: ArticleCardProps) {
    const sourceName = post.source_url ? getSourceName(post.source_url) : "CryptoNews";
    const imageUrl = extractImageUrl(post.content);
    const title = post.title || "Untitled Article";

    return (
        <Link
            href={`/posts/${post.id}`}
            className={cn(
                "group flex flex-col bg-[#111] border border-white/5 rounded-2xl overflow-hidden hover:border-white/10 transition-all duration-300 h-full",
                className
            )}
        >
            {/* Image Section */}
            <div className={cn("relative w-full overflow-hidden bg-gray-900 border-b border-white/5",
                variant === "compact" ? "aspect-[16/9]" : "aspect-[16/9]"
            )}>
                {imageUrl ? (
                    <Image
                        src={imageUrl}
                        alt={title}
                        fill
                        className="object-cover group-hover:scale-105 transition-transform duration-500 opacity-90 group-hover:opacity-100"
                    />
                ) : (
                    <div className="w-full h-full bg-gradient-to-br from-[#1a1a1a] to-[#0a0a0a] flex items-center justify-center p-4">
                        <span className="text-2xl font-black text-white/5 uppercase tracking-tighter">News</span>
                    </div>
                )}
            </div>

            {/* Content Section */}
            <div className="p-4 flex flex-col flex-1 relative">
                <div className="flex items-center justify-between mb-3">
                    <span className="text-[10px] uppercase font-bold text-cyan-500 tracking-wider">
                        {sourceName}
                    </span>
                    {variant === "default" && (
                        <span className="text-[10px] text-gray-500 flex items-center gap-1">
                            <Clock className="w-3 h-3" />
                            {formatDate(post.created_at)}
                        </span>
                    )}
                </div>

                <h3 className={cn(
                    "font-bold text-gray-100 leading-snug group-hover:text-cyan-400 transition-colors",
                    variant === "compact" ? "text-sm line-clamp-3 mb-1" : "text-lg line-clamp-2 mb-3"
                )}>
                    {title}
                </h3>

                {variant === "default" && (
                    <div className="mt-auto pt-4 flex items-center justify-between text-[11px] text-gray-500 border-t border-white/5">
                        <span className="flex items-center gap-1.5 opacity-60">
                            By {sourceName}
                        </span>
                        <span className="group-hover:translate-x-1 transition-transform text-gray-400 group-hover:text-white flex items-center gap-0.5">
                            Read more &rarr;
                        </span>
                    </div>
                )}
            </div>
        </Link>
    );
}

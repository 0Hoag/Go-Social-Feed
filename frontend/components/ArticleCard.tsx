import Link from "next/link";
import Image from "next/image";
import { Post } from "@/lib/types";
import { formatDate, getSourceName, truncateText } from "@/lib/utils";
import { ExternalLink, Clock } from "lucide-react";

interface ArticleCardProps {
    post: Post;
}

export default function ArticleCard({ post }: ArticleCardProps) {
    const sourceName = post.source_url ? getSourceName(post.source_url) : "Unknown";
    const imageUrl = post.file_ids && post.file_ids.length > 0
        ? post.file_ids[0]
        : "/placeholder-crypto.jpg";

    return (
        <article className="group relative overflow-hidden rounded-2xl bg-gradient-to-br from-gray-900/90 to-gray-800/90 border border-white/10 backdrop-blur-sm transition-all duration-300 hover:border-cyan-500/50 hover:shadow-xl hover:shadow-cyan-500/20 hover:-translate-y-1">
            {/* Image Section */}
            <div className="relative h-48 w-full overflow-hidden bg-gray-800">
                {imageUrl !== "/placeholder-crypto.jpg" ? (
                    <Image
                        src={imageUrl}
                        alt={post.content || "Crypto news article"}
                        fill
                        className="object-cover transition-transform duration-500 group-hover:scale-110"
                        sizes="(max-width: 768px) 100vw, (max-width: 1200px) 50vw, 33vw"
                    />
                ) : (
                    <div className="flex h-full items-center justify-center bg-gradient-to-br from-cyan-900/50 to-blue-900/50">
                        <TrendingUp className="h-16 w-16 text-cyan-400/50" />
                    </div>
                )}

                {/* Gradient Overlay */}
                <div className="absolute inset-0 bg-gradient-to-t from-gray-900 via-gray-900/40 to-transparent opacity-60" />

                {/* Source Badge */}
                <div className="absolute top-3 right-3">
                    <span className="inline-flex items-center rounded-full bg-cyan-500/20 backdrop-blur-sm px-3 py-1 text-xs font-semibold text-cyan-300 border border-cyan-400/30">
                        {sourceName}
                    </span>
                </div>
            </div>

            {/* Content Section */}
            <div className="p-5 space-y-3">
                {/* Title */}
                <Link href={`/posts/${post.id}`}>
                    <h3 className="text-lg font-bold text-white line-clamp-2 transition-colors group-hover:text-cyan-400">
                        {post.content || "Untitled Article"}
                    </h3>
                </Link>

                {/* Summary */}
                {post.full_content && (
                    <p className="text-sm text-gray-400 line-clamp-3">
                        {truncateText(post.full_content, 150)}
                    </p>
                )}

                {/* Footer */}
                <div className="flex items-center justify-between pt-3 border-t border-white/10">
                    <div className="flex items-center space-x-2 text-xs text-gray-500">
                        <Clock className="h-3.5 w-3.5" />
                        <time dateTime={post.created_at}>
                            {formatDate(post.created_at)}
                        </time>
                    </div>

                    {post.source_url && (
                        <a
                            href={post.source_url}
                            target="_blank"
                            rel="noopener noreferrer"
                            className="flex items-center space-x-1 text-xs font-medium text-cyan-400 transition-colors hover:text-cyan-300"
                            onClick={(e) => e.stopPropagation()}
                        >
                            <span>Read more</span>
                            <ExternalLink className="h-3.5 w-3.5" />
                        </a>
                    )}
                </div>
            </div>
        </article>
    );
}

// Import for placeholder
import { TrendingUp } from "lucide-react";

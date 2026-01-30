import Link from "next/link";
import Image from "next/image";
import { Post } from "@/lib/types";
import { formatDate, getSourceName, extractImageUrl } from "@/lib/utils";
import { Clock } from "lucide-react";

interface HeroPostProps {
    post: Post;
}

export default function HeroPost({ post }: HeroPostProps) {
    const imageUrl = extractImageUrl(post.content);
    const sourceName = post.source_url ? getSourceName(post.source_url) : "CryptoNews";

    return (
        <Link href={`/posts/${post.id}`} className="group block relative rounded-3xl overflow-hidden bg-[#111] border border-white/5 aspect-[16/9] sm:aspect-[2/1] md:aspect-[21/9] lg:aspect-[2/1] transition-transform hover:scale-[1.01] duration-500">
            {/* Background Image */}
            <div className="absolute inset-0 z-0">
                {imageUrl ? (
                    <Image
                        src={imageUrl}
                        alt={post.title}
                        fill
                        className="object-cover opacity-60 group-hover:opacity-70 transition-opacity duration-500"
                        priority
                    />
                ) : (
                    <div className="w-full h-full bg-gradient-to-br from-blue-900/40 to-purple-900/40" />
                )}
                {/* Gradient Overlays */}
                <div className="absolute inset-0 bg-gradient-to-t from-[#000] via-[#000]/60 to-transparent" />
                <div className="absolute inset-0 bg-gradient-to-r from-[#000] via-[#000]/40 to-transparent" />
            </div>

            {/* Content */}
            <div className="absolute inset-0 z-10 flex flex-col justify-end p-6 md:p-10 max-w-3xl">
                <div className="flex items-center gap-3 mb-4">
                    <span className="bg-cyan-500 text-black text-[10px] font-bold px-2 py-1 rounded-full uppercase tracking-wide">
                        Top Story
                    </span>
                    <span className="text-gray-300 text-xs font-medium flex items-center gap-1.5 bg-black/50 backdrop-blur-sm px-2 py-1 rounded-full border border-white/10">
                        <span className="w-1.5 h-1.5 rounded-full bg-cyan-400 animate-pulse" />
                        {sourceName}
                    </span>
                </div>

                <h1 className="text-2xl md:text-3xl lg:text-4xl font-bold text-white leading-tight mb-4 group-hover:text-cyan-50 transition-colors">
                    {post.title}
                </h1>

                <div className="flex items-center gap-4 text-gray-400 text-xs md:text-sm">
                    <span className="flex items-center gap-1.5">
                        <Clock className="w-3.5 h-3.5" />
                        {formatDate(post.created_at)}
                    </span>
                </div>
            </div>
        </Link>
    );
}

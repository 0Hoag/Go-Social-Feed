import { Post } from "@/lib/types";
import { formatDate } from "@/lib/utils";
import Link from "next/link";
import Image from "next/image";
import { extractImageUrl } from "@/lib/utils";

interface RelatedNewsProps {
    posts: Post[];
}

export default function RelatedNews({ posts }: RelatedNewsProps) {
    if (!posts || posts.length === 0) return null;

    return (
        <div className="space-y-4">
            <div className="flex items-center justify-between mb-2">
                <h3 className="text-gray-400 text-xs font-bold tracking-wider uppercase">Related News</h3>
            </div>

            {posts.slice(0, 5).map((post) => {
                const imageUrl = extractImageUrl(post.content);
                return (
                    <Link href={`/posts/${post.id}`} key={post.id} className="group flex gap-4 bg-[#111] border border-white/5 p-3 rounded-xl hover:border-white/10 transition-colors">
                        <div className="flex-1">
                            <h4 className="text-gray-200 text-xs font-bold leading-snug line-clamp-3 group-hover:text-cyan-400 transition-colors">
                                {post.title}
                            </h4>
                            <span className="text-[10px] text-gray-500 mt-2 block">
                                {formatDate(post.created_at)}
                            </span>
                        </div>
                        {imageUrl && (
                            <div className="relative w-16 h-16 rounded-lg overflow-hidden shrink-0">
                                <Image
                                    src={imageUrl}
                                    alt={post.title}
                                    fill
                                    className="object-cover opacity-80 group-hover:opacity-100 transition-opacity"
                                />
                            </div>
                        )}
                    </Link>
                );
            })}
        </div>
    );
}

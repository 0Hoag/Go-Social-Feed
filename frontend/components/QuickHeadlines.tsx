import { Post } from "@/lib/types";
import { formatDate } from "@/lib/utils";
import { Zap } from "lucide-react";
import Link from "next/link";

interface QuickHeadlinesProps {
    posts: Post[];
}

export default function QuickHeadlines({ posts }: QuickHeadlinesProps) {
    return (
        <div className="bg-[#111] border border-white/5 rounded-2xl p-5 h-fit sticky top-24">
            <div className="flex items-center gap-2 mb-6">
                <Zap className="w-4 h-4 text-yellow-500 fill-yellow-500" />
                <h3 className="text-gray-400 text-xs font-bold tracking-wider uppercase">Quick Headlines</h3>
            </div>

            <div className="space-y-5">
                {posts.slice(0, 8).map((post) => (
                    <Link href={`/posts/${post.id}`} key={post.id} className="group block">
                        <div className="flex gap-3">
                            <div className="mt-1.5 min-w-[6px]">
                                <span className="block w-1.5 h-1.5 rounded-full bg-blue-500/50 group-hover:bg-blue-400 group-hover:shadow-[0_0_8px_rgba(59,130,246,0.5)] transition-all duration-300"></span>
                            </div>
                            <div>
                                <h4 className="text-gray-300 group-hover:text-white text-sm font-medium leading-snug transition-colors line-clamp-3">
                                    {post.title}
                                </h4>
                                <span className="text-[10px] text-gray-500 mt-1 block">
                                    {formatDate(post.created_at)}
                                </span>
                            </div>
                        </div>
                    </Link>
                ))}
            </div>

            <button className="text-blue-400 text-xs font-medium mt-6 hover:text-blue-300 transition-colors flex items-center gap-1 group">
                See all news
                <span className="group-hover:translate-x-1 transition-transform">→</span>
            </button>
        </div>
    );
}

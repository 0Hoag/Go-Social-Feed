"use client";

import { useEffect, useState, use } from "react";
import { getPostById, getPosts } from "@/lib/api";
import { Post } from "@/lib/types";
import { formatDate, extractImageUrl, getSourceName } from "@/lib/utils";
import ReactMarkdown from "react-markdown";
import remarkGfm from "remark-gfm";
import Link from "next/link";
import Image from "next/image";
import { Share2, Clock, ExternalLink, Loader2 } from "lucide-react";
import CryptoRanking from "@/components/CryptoRanking";
import MarketWidgets from "@/components/MarketWidgets";
import RelatedNews from "@/components/RelatedNews";
import rehypeRaw from "rehype-raw";

export default function PostDetail({ params }: { params: Promise<{ id: string }> }) {
    const { id } = use(params);
    const [post, setPost] = useState<Post | null>(null);
    const [relatedPosts, setRelatedPosts] = useState<Post[]>([]);
    const [loading, setLoading] = useState(true);

    useEffect(() => {
        const fetchData = async () => {
            try {
                const [postData, allPostsResponse] = await Promise.all([
                    getPostById(id),
                    getPosts({ limit: 10 })
                ]);

                setPost(postData);

                // Filter out current post from related list
                const related = allPostsResponse.posts.filter(p => p.id !== id);

                // DEMO Logic: Fake ID generation removed to prevent 404s
                // If we need more posts, we should handle that in the UI or fetch more from API

                setRelatedPosts(related);

            } catch (err) {
                console.error("Failed to load post data", err);
            } finally {
                setLoading(false);
            }
        };

        fetchData();
    }, [id]);

    if (loading) {
        return (
            <div className="min-h-screen bg-[#050505] flex items-center justify-center">
                <Loader2 className="w-8 h-8 text-cyan-500 animate-spin" />
            </div>
        );
    }

    if (!post) return <div className="text-white text-center py-20 min-h-screen bg-[#050505]">Post not found</div>;

    const imageUrl = extractImageUrl(post.content);
    // Remove image md syntax from content to avoid duplicate images
    const cleanContent = post.content.replace(/!\[.*?\]\(.*?\)/g, "").trim();
    const sourceName = post.source_url ? getSourceName(post.source_url) : "Source";

    return (
        <main className="min-h-screen bg-[#050505] text-gray-200 font-sans selection:bg-cyan-500/20 selection:text-cyan-200 pb-20">

            {/* Breadcrumb / Nav */}
            <div className="border-b border-white/5 bg-[#050505]">
                <div className="max-w-[1600px] mx-auto px-4 py-3 flex items-center gap-2 text-xs text-gray-500">
                    <Link href="/" className="hover:text-white transition-colors">Home</Link>
                    <span>/</span>
                    <span className="text-gray-300 truncate max-w-[300px]">{post.title}</span>
                </div>
            </div>

            <div className="max-w-[1600px] mx-auto px-4 py-8">
                <div className="grid grid-cols-1 lg:grid-cols-12 gap-8 relative">

                    {/* LEFT COLUMN: Crypto Ranking (3/12) */}
                    <div className="hidden xl:block col-span-3">
                        <div className="sticky top-24">
                            <CryptoRanking />
                        </div>
                    </div>

                    {/* CENTER COLUMN: Main Article (6/12) */}
                    <div className="col-span-1 lg:col-span-8 xl:col-span-6">
                        <article>
                            <h1 className="text-3xl md:text-3xl lg:text-4xl font-bold text-white leading-tight mb-6">
                                {post.title}
                            </h1>

                            <div className="flex items-center gap-4 text-xs text-gray-400 mb-8 border-b border-white/5 pb-6">
                                <span className="flex items-center gap-1.5">
                                    <Clock className="w-4 h-4" />
                                    {formatDate(post.created_at)}
                                </span>
                                <span className="bg-cyan-500/10 text-cyan-400 px-2 py-1 rounded border border-cyan-500/20">
                                    {sourceName}
                                </span>
                            </div>

                            {/* Featured Image */}
                            {imageUrl && (
                                <div className="relative w-full aspect-[16/9] rounded-2xl overflow-hidden mb-8 border border-white/5 bg-gray-900 group">
                                    <Image
                                        src={imageUrl}
                                        alt={post.title}
                                        fill
                                        className="object-cover group-hover:scale-105 transition-transform duration-700"
                                        priority
                                    />
                                    <div className="absolute inset-0 bg-gradient-to-t from-[#050505]/50 to-transparent"></div>
                                </div>
                            )}

                            {/* Main Content */}
                            <div className="prose prose-invert prose-lg max-w-none text-gray-300 leading-relaxed prose-headings:text-white prose-a:text-cyan-400 hover:prose-a:text-cyan-300 prose-img:rounded-xl">
                                <ReactMarkdown remarkPlugins={[remarkGfm]} rehypePlugins={[rehypeRaw]}>
                                    {cleanContent}
                                </ReactMarkdown>
                            </div>

                            {/* Footer Actions */}
                            <div className="mt-12 flex items-center justify-between border-t border-white/10 pt-8">
                                <button className="flex items-center gap-2 text-sm text-gray-400 hover:text-white transition-colors bg-white/5 px-4 py-2 rounded-lg border border-white/5 hover:border-white/10">
                                    <Share2 className="w-4 h-4" /> Share this article
                                </button>

                                {post.source_url && (
                                    <a
                                        href={post.source_url}
                                        target="_blank"
                                        rel="noopener noreferrer"
                                        className="flex items-center gap-2 text-sm text-cyan-400 hover:text-cyan-300 transition-colors font-medium border border-cyan-500/20 px-4 py-2 rounded-lg hover:bg-cyan-500/10"
                                    >
                                        Original Article <ExternalLink className="w-4 h-4" />
                                    </a>
                                )}
                            </div>
                        </article>
                    </div>

                    {/* RIGHT COLUMN: Related & Calendar (3/12) */}
                    <div className="hidden lg:block col-span-4 xl:col-span-3 space-y-8">
                        {/* Related News */}
                        <RelatedNews posts={relatedPosts} />

                        {/* Widgets Wrapper */}
                        <div className="sticky top-24">
                            <MarketWidgets />
                        </div>
                    </div>

                </div>
            </div>
        </main>
    );
}

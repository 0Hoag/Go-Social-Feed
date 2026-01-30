import { notFound } from "next/navigation";
import Image from "next/image";
import Link from "next/link";
import { getPostById } from "@/lib/api";
import { formatDate, getSourceName, extractImageUrl } from "@/lib/utils";
import { ArrowLeft, ExternalLink, Clock, TrendingUp } from "lucide-react";

interface PostPageProps {
    params: Promise<{ id: string }>;
}

export default async function PostPage({ params }: PostPageProps) {
    const { id } = await params;

    let post;
    try {
        post = await getPostById(id);
    } catch (error) {
        console.error("Failed to fetch post:", error);
        notFound();
    }

    if (!post) {
        notFound();
    }

    const sourceName = post.source_url ? getSourceName(post.source_url) : "Unknown";
    const imageUrl = extractImageUrl(post.content);
    const title = post.title || "Untitled Article";

    return (
        <div className="min-h-screen">
            {/* Back Button */}
            <div className="container mx-auto px-4 py-6">
                <Link
                    href="/"
                    className="inline-flex items-center space-x-2 text-gray-400 hover:text-cyan-400 transition-colors"
                >
                    <ArrowLeft className="h-4 w-4" />
                    <span>Back to feed</span>
                </Link>
            </div>

            {/* Article Content */}
            <article className="container mx-auto px-4 pb-16">
                <div className="max-w-4xl mx-auto">
                    {/* Header */}
                    <header className="mb-8">
                        <div className="flex items-center space-x-3 mb-4">
                            <span className="inline-flex items-center rounded-full bg-cyan-500/20 backdrop-blur-sm px-4 py-1.5 text-sm font-semibold text-cyan-300 border border-cyan-400/30">
                                {sourceName}
                            </span>
                            <div className="flex items-center space-x-2 text-sm text-gray-500">
                                <Clock className="h-4 w-4" />
                                <time dateTime={post.created_at}>
                                    {formatDate(post.created_at)}
                                </time>
                            </div>
                        </div>

                        <h1 className="text-4xl md:text-5xl font-bold text-white mb-6 leading-tight">
                            {title}
                        </h1>

                        {post.source_url && (
                            <a
                                href={post.source_url}
                                target="_blank"
                                rel="noopener noreferrer"
                                className="inline-flex items-center space-x-2 text-cyan-400 hover:text-cyan-300 transition-colors font-medium"
                            >
                                <span>Read original article</span>
                                <ExternalLink className="h-4 w-4" />
                            </a>
                        )}
                    </header>

                    {/* Featured Image */}
                    {imageUrl && (
                        <div className="relative w-full h-[400px] md:h-[500px] rounded-2xl overflow-hidden mb-8 border border-white/10">
                            <Image
                                src={imageUrl}
                                alt={post.content || "Article image"}
                                fill
                                className="object-cover"
                                priority
                            />
                        </div>
                    )}

                    {/* Article Body */}
                    <div className="prose prose-invert prose-lg max-w-none">
                        <div className="rounded-2xl bg-gradient-to-br from-gray-900/90 to-gray-800/90 border border-white/10 p-8">
                            {post.content ? (
                                <div className="text-gray-300 leading-relaxed whitespace-pre-wrap">
                                    {/* Remove the markdown image since we display it at the top */}
                                    {post.content.replace(/!\[.*?\]\(.*?\)/g, "").trim()}
                                </div>
                            ) : (
                                <div className="text-center py-12">
                                    <TrendingUp className="h-12 w-12 text-gray-600 mx-auto mb-4" />
                                    <p className="text-gray-500">
                                        Full content not available. Please visit the original source.
                                    </p>
                                </div>
                            )}
                        </div>
                    </div>

                    {/* Footer CTA */}
                    {post.source_url && (
                        <div className="mt-12 p-6 rounded-2xl bg-gradient-to-r from-cyan-500/10 to-blue-500/10 border border-cyan-500/20">
                            <div className="flex items-center justify-between">
                                <div>
                                    <h3 className="text-lg font-semibold text-white mb-1">
                                        Want to read more?
                                    </h3>
                                    <p className="text-gray-400 text-sm">
                                        Visit the original article on {sourceName}
                                    </p>
                                </div>
                                <a
                                    href={post.source_url}
                                    target="_blank"
                                    rel="noopener noreferrer"
                                    className="px-6 py-3 rounded-lg bg-cyan-500 text-white font-medium hover:bg-cyan-600 transition-colors flex items-center space-x-2"
                                >
                                    <span>Read More</span>
                                    <ExternalLink className="h-4 w-4" />
                                </a>
                            </div>
                        </div>
                    )}
                </div>
            </article>
        </div>
    );
}

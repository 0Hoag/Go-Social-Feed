"use client";

import { useEffect, useState } from "react";
import { getPosts } from "@/lib/api";
import { Post } from "@/lib/types";
import ArticleGrid from "@/components/ArticleGrid";
import LoadingSpinner from "@/components/LoadingSpinner";
import { Sparkles } from "lucide-react";

export default function HomePage() {
  const [posts, setPosts] = useState<Post[]>([]);
  const [loading, setLoading] = useState(true);
  const [error, setError] = useState<string | null>(null);

  useEffect(() => {
    const fetchPosts = async () => {
      try {
        setLoading(true);
        setError(null);
        const response = await getPosts({ page: 1, limit: 30 });
        setPosts(response.posts);
      } catch (err) {
        setError("Failed to load articles. Please try again later.");
        console.error("Error fetching posts:", err);
      } finally {
        setLoading(false);
      }
    };

    fetchPosts();
  }, []);

  return (
    <div className="container mx-auto px-4 py-8">
      {/* Hero Section */}
      <section className="mb-12 text-center">
        <div className="inline-flex items-center space-x-2 rounded-full bg-cyan-500/10 border border-cyan-500/20 px-4 py-2 mb-6">
          <Sparkles className="h-4 w-4 text-cyan-400" />
          <span className="text-sm font-medium text-cyan-300">
            Live Crypto News
          </span>
        </div>

        <h1 className="text-5xl md:text-6xl font-bold mb-4 bg-gradient-to-r from-cyan-400 via-blue-500 to-purple-600 bg-clip-text text-transparent">
          Latest Crypto News
        </h1>

        <p className="text-xl text-gray-400 max-w-2xl mx-auto">
          Stay ahead of the market with real-time cryptocurrency news from top sources
        </p>
      </section>

      {/* Articles Section */}
      <section className="animate-fade-in">
        {loading ? (
          <LoadingSpinner size="lg" />
        ) : error ? (
          <div className="text-center py-20">
            <div className="rounded-2xl bg-red-500/10 border border-red-500/20 p-8 max-w-md mx-auto">
              <p className="text-red-400 font-medium">{error}</p>
              <button
                onClick={() => window.location.reload()}
                className="mt-4 px-6 py-2 rounded-lg bg-red-500/20 text-red-300 hover:bg-red-500/30 transition-colors"
              >
                Retry
              </button>
            </div>
          </div>
        ) : (
          <>
            <div className="flex items-center justify-between mb-6">
              <h2 className="text-2xl font-bold text-white">
                Recent Articles
                <span className="ml-3 text-sm font-normal text-gray-500">
                  ({posts.length} articles)
                </span>
              </h2>
            </div>

            <ArticleGrid posts={posts} />
          </>
        )}
      </section>
    </div>
  );
}

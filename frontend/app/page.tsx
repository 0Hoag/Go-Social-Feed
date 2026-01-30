"use client";

import { useEffect, useState } from "react";
import { getPosts } from "@/lib/api";
import { Post } from "@/lib/types";
import HeroPost from "@/components/HeroPost";
import QuickHeadlines from "@/components/QuickHeadlines";
import MarketWidgets from "@/components/MarketWidgets";
import ArticleCard from "@/components/ArticleCard";
import { Loader2 } from "lucide-react";

export default function Home() {
  const [posts, setPosts] = useState<Post[]>([]);
  const [loading, setLoading] = useState(true);
  const [error, setError] = useState("");

  useEffect(() => {
    const fetchPosts = async () => {
      try {
        const response = await getPosts();
        let fetchedPosts = response.posts;

        setPosts(fetchedPosts);
      } catch (err) {
        console.error("Failed to fetch posts:", err);
        setError("Failed to load posts. Please try again later.");
      } finally {
        setLoading(false);
      }
    };

    fetchPosts();
  }, []);

  if (loading) {
    return (
      <div className="min-h-screen bg-[#050505] flex items-center justify-center">
        <Loader2 className="w-8 h-8 text-cyan-500 animate-spin" />
      </div>
    );
  }

  // Data slicing strategy
  const topStory = posts.length > 0 ? posts[0] : null;
  const subStories = posts.length > 1 ? posts.slice(1, 4) : [];
  const otherNews = posts.length > 4 ? posts.slice(4) : [];

  // For Headlines, we use the sliced list. Exclude top story usually.
  const recentHeadlines = posts.length > 1 ? posts.slice(1, 15) : [];

  return (
    <main className="min-h-screen bg-[#050505] text-gray-200 font-sans selection:bg-cyan-500/20 selection:text-cyan-200">
      <div className="max-w-[1600px] mx-auto px-4 py-8">

        {/* 3-Column Layout */}
        <div className="grid grid-cols-1 lg:grid-cols-12 gap-6 relative">

          {/* LEFT COLUMN: Quick Headlines (2/12) - Sticky Sidebar */}
          <div className="hidden xl:block col-span-2">
            {recentHeadlines.length > 0 && <QuickHeadlines posts={recentHeadlines} />}
          </div>

          {/* CENTER COLUMN: Main Content (7/12) */}
          <div className="col-span-1 lg:col-span-8 xl:col-span-7 space-y-8">
            {/* Top Story Hero */}
            {topStory && <HeroPost post={topStory} />}

            {/* Sub Stories Row */}
            {subStories.length > 0 && (
              <div>
                <h2 className="text-white/90 text-sm font-bold mb-4 flex items-center gap-2 uppercase tracking-wider">
                  <span className="text-cyan-500">⚡</span> Must Read
                </h2>
                <div className="grid grid-cols-1 sm:grid-cols-3 gap-4">
                  {subStories.map(post => (
                    <ArticleCard key={post.id} post={post} variant="compact" />
                  ))}
                </div>
              </div>
            )}

            {/* Other News Grid */}
            {otherNews.length > 0 && (
              <div>
                <h2 className="text-white/90 text-sm font-bold mb-4 border-t border-white/5 pt-8 uppercase tracking-wider">
                  Latest Stories
                </h2>
                <div className="grid grid-cols-1 md:grid-cols-2 gap-4">
                  {otherNews.map(post => (
                    <ArticleCard key={post.id} post={post} variant="default" />
                  ))}
                </div>
              </div>
            )}
          </div>

          {/* RIGHT COLUMN: Widgets (3/12) - Sticky Sidebar */}
          <div className="hidden lg:block col-span-4 xl:col-span-3">
            <div className="sticky top-24">
              <MarketWidgets />
            </div>
          </div>
        </div>
      </div>
    </main>
  );
}

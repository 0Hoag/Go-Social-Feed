'use client';

import { useState, useEffect, useCallback } from 'react';
import { getPosts, Post } from '@/lib/api/posts';
import { useAuthStore } from '@/lib/stores/authStore';
import PostCard from '@/components/PostCard';
import CreatePostModal from '@/components/CreatePostModal';
import { toast } from 'sonner';

export default function FeedPage() {
    const { isAuthenticated } = useAuthStore();
    const [posts, setPosts] = useState<Post[]>([]);
    const [loading, setLoading] = useState(true);
    const [page, setPage] = useState(1);
    const [hasMore, setHasMore] = useState(true);
    const [showCreateModal, setShowCreateModal] = useState(false);

    const loadPosts = useCallback(async (reset = false) => {
        setLoading(true);
        try {
            const currentPage = reset ? 1 : page;
            const data = await getPosts({ page: currentPage, limit: 20 });
            // Exclude crawled news posts — they always have a source_url; user posts do not
            const userPosts = (data.items || []).filter(
                (p) => !p.source_url
            );
            if (reset) {
                setPosts(userPosts);
                setPage(1);
            } else {
                setPosts((prev) => [...prev, ...userPosts]);
            }
            setHasMore((data.items || []).length === 20);
        } catch {
            toast.error('Không thể tải bài viết!');
        } finally {
            setLoading(false);
        }
    }, [page]);

    useEffect(() => {
        loadPosts(true);
    }, []);

    const handleLoadMore = () => {
        const next = page + 1;
        setPage(next);
        loadPosts();
    };

    const handlePostCreated = () => {
        loadPosts(true);
    };

    const handlePostDeleted = (id: string) => {
        setPosts((prev) => prev.filter((p) => p.id !== id));
    };

    return (
        <div className="min-h-screen bg-black">
            <div className="max-w-2xl mx-auto px-4 py-8">
                {/* Page header */}
                <div className="flex items-center justify-between mb-6">
                    <div>
                        <h1 className="text-2xl font-bold text-white">Bảng Tin</h1>
                        <p className="text-slate-400 text-sm mt-0.5">Khám phá những bài viết mới nhất</p>
                    </div>
                    {isAuthenticated && (
                        <button
                            onClick={() => setShowCreateModal(true)}
                            className="flex items-center gap-2 bg-sky-500 hover:bg-sky-400 text-white font-semibold px-4 py-2.5 rounded-xl transition-all duration-200 shadow-lg shadow-sky-500/20 text-sm"
                        >
                            <svg className="w-4 h-4" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                                <path strokeLinecap="round" strokeLinejoin="round" strokeWidth={2} d="M12 4v16m8-8H4" />
                            </svg>
                            Đăng bài
                        </button>
                    )}
                </div>

                {/* Not logged in banner */}
                {!isAuthenticated && (
                    <div className="bg-sky-500/10 border border-sky-500/20 rounded-2xl p-4 mb-6 flex items-center justify-between">
                        <div className="flex items-center gap-3">
                            <div className="w-8 h-8 rounded-full bg-sky-500/20 flex items-center justify-center">
                                <svg className="w-4 h-4 text-sky-400" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                                    <path strokeLinecap="round" strokeLinejoin="round" strokeWidth={2} d="M16 7a4 4 0 11-8 0 4 4 0 018 0zM12 14a7 7 0 00-7 7h14a7 7 0 00-7-7z" />
                                </svg>
                            </div>
                            <p className="text-zinc-300 text-sm">Đăng nhập để đăng bài và tương tác!</p>
                        </div>
                        <a href="/login" className="text-sky-400 hover:text-sky-300 text-sm font-medium transition-colors">
                            Đăng nhập →
                        </a>
                    </div>
                )}

                {/* Posts */}
                {loading && posts.length === 0 ? (
                    <div className="space-y-4">
                        {[...Array(3)].map((_, i) => (
                            <div key={i} className="bg-zinc-900 border border-zinc-800 rounded-2xl p-5 animate-pulse">
                                <div className="flex gap-3 mb-3">
                                    <div className="w-10 h-10 bg-zinc-800 rounded-full" />
                                    <div className="flex-1 space-y-2">
                                        <div className="h-3 bg-zinc-800 rounded w-1/4" />
                                        <div className="h-2 bg-zinc-900 rounded w-1/6" />
                                    </div>
                                </div>
                                <div className="space-y-2">
                                    <div className="h-3 bg-zinc-800 rounded" />
                                    <div className="h-3 bg-zinc-800 rounded w-4/5" />
                                    <div className="h-3 bg-zinc-900 rounded w-2/3" />
                                </div>
                            </div>
                        ))}
                    </div>
                ) : posts.length === 0 ? (
                    <div className="text-center py-20">
                        <div className="text-6xl mb-4">📭</div>
                        <p className="text-zinc-500">Chưa có bài viết nào</p>
                        {isAuthenticated && (
                            <button
                                onClick={() => setShowCreateModal(true)}
                                className="mt-4 text-sky-400 hover:text-sky-300 font-medium transition-colors"
                            >
                                Đăng bài đầu tiên →
                            </button>
                        )}
                    </div>
                ) : (
                    <>
                        <div className="space-y-4">
                            {posts.map((post) => (
                                <PostCard
                                    key={post.id}
                                    post={post}
                                    onDeleted={handlePostDeleted}
                                />
                            ))}
                        </div>

                        {hasMore && (
                            <div className="mt-6 text-center">
                                <button
                                    onClick={handleLoadMore}
                                    disabled={loading}
                                    className="px-6 py-3 bg-zinc-900 border border-zinc-800 hover:bg-zinc-800 text-zinc-300 rounded-xl text-sm font-medium transition-all duration-200 disabled:opacity-50"
                                >
                                    {loading ? 'Đang tải...' : 'Tải thêm'}
                                </button>
                            </div>
                        )}
                    </>
                )}
            </div>

            {/* Create Post Modal */}
            {showCreateModal && (
                <CreatePostModal
                    onClose={() => setShowCreateModal(false)}
                    onCreated={handlePostCreated}
                />
            )}
        </div>
    );
}

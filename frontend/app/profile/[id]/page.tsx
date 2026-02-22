'use client';

import { useState, useEffect } from 'react';
import { useParams } from 'next/navigation';
import { getUser, User } from '@/lib/api/users';
import { getPosts, Post } from '@/lib/api/posts';
import PostCard from '@/components/PostCard';
import FollowButton from '@/components/FollowButton';

export default function ProfilePage() {
    const { id } = useParams<{ id: string }>();
    const [user, setUser] = useState<User | null>(null);
    const [posts, setPosts] = useState<Post[]>([]);
    const [loading, setLoading] = useState(true);

    useEffect(() => {
        if (!id) return;
        const load = async () => {
            setLoading(true);
            try {
                const [u, postsData] = await Promise.all([getUser(id), getPosts({ page: 1, limit: 20 })]);
                setUser(u);
                const userPosts = (postsData.items || []).filter((p) => !p.source_url);
                setPosts(userPosts);
            } catch { } finally { setLoading(false); }
        };
        load();
    }, [id]);

    if (loading) return (
        <div className="min-h-screen bg-black flex items-center justify-center">
            <div className="w-10 h-10 border-2 border-sky-500 border-t-transparent rounded-full animate-spin" />
        </div>
    );

    if (!user) return (
        <div className="min-h-screen bg-black flex items-center justify-center">
            <p className="text-zinc-500">Không tìm thấy người dùng</p>
        </div>
    );

    return (
        <div className="min-h-screen bg-black">
            <div className="max-w-2xl mx-auto px-4 py-8">
                {/* Profile Card */}
                <div className="bg-zinc-900 border border-zinc-800 rounded-2xl overflow-hidden mb-6">
                    <div className="h-32 bg-gradient-to-r from-sky-950 via-zinc-900 to-black" />
                    <div className="px-6 pb-6">
                        <div className="flex items-end justify-between -mt-10 mb-4">
                            {user.avatar_url ? (
                                <img src={user.avatar_url} alt={user.username} className="w-20 h-20 rounded-2xl object-cover ring-4 ring-zinc-900 shadow-xl" />
                            ) : (
                                <div className="w-20 h-20 rounded-2xl bg-sky-500 flex items-center justify-center text-white font-bold text-3xl ring-4 ring-zinc-900 shadow-xl">
                                    {user.username.charAt(0).toUpperCase()}
                                </div>
                            )}
                            <FollowButton targetId={user.id} />
                        </div>
                        <h1 className="text-2xl font-bold text-white">{user.username}</h1>
                        <p className="text-zinc-500 text-sm mt-1">Thành viên từ {new Date(user.created_at).toLocaleDateString('vi-VN')}</p>
                    </div>
                </div>

                <h2 className="text-lg font-semibold text-white mb-4">Bài viết</h2>
                {posts.length === 0 ? (
                    <div className="text-center py-12 text-zinc-600">
                        <div className="text-4xl mb-2">📝</div>
                        <p>Chưa có bài viết nào</p>
                    </div>
                ) : (
                    <div className="space-y-4">
                        {posts.map((post) => (
                            <PostCard key={post.id} post={post} authorName={user.username} authorAvatar={user.avatar_url} authorId={user.id} />
                        ))}
                    </div>
                )}
            </div>
        </div>
    );
}

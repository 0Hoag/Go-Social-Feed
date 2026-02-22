'use client';

import { useState, useEffect } from 'react';
import { useRouter } from 'next/navigation';
import { getMyInfo, updateUser } from '@/lib/api/users';
import { getPosts, Post } from '@/lib/api/posts';
import { useAuthStore } from '@/lib/stores/authStore';
import PostCard from '@/components/PostCard';
import { toast } from 'sonner';

export default function MyProfilePage() {
    const router = useRouter();
    const { user, setUser, isAuthenticated, logout } = useAuthStore();
    const [posts, setPosts] = useState<Post[]>([]);
    const [loading, setLoading] = useState(true);
    const [editing, setEditing] = useState(false);
    const [form, setForm] = useState({ username: '', avatar_url: '' });
    const [saving, setSaving] = useState(false);

    useEffect(() => {
        if (!isAuthenticated) { router.push('/login'); return; }
        const load = async () => {
            setLoading(true);
            try {
                const [me, postsData] = await Promise.all([getMyInfo(), getPosts({ page: 1, limit: 20 })]);
                setUser(me);
                setForm({ username: me.username, avatar_url: me.avatar_url || '' });
                const userPosts = (postsData.items || []).filter((p) => !p.source_url);
                setPosts(userPosts);
            } catch { toast.error('Không thể tải thông tin!'); } finally { setLoading(false); }
        };
        load();
    }, [isAuthenticated]);

    const handleSave = async () => {
        if (!user) return;
        setSaving(true);
        try {
            await updateUser(user.id, { username: form.username, avatar_url: form.avatar_url });
            const updated = await getMyInfo();
            setUser(updated); setEditing(false); toast.success('Đã cập nhật hồ sơ!');
        } catch { toast.error('Cập nhật thất bại!'); } finally { setSaving(false); }
    };

    const handleLogout = () => { logout(); router.push('/'); toast.success('Đã đăng xuất'); };

    if (loading) return (
        <div className="min-h-screen bg-black flex items-center justify-center">
            <div className="w-10 h-10 border-2 border-sky-500 border-t-transparent rounded-full animate-spin" />
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
                            {user?.avatar_url ? (
                                <img src={user.avatar_url} alt={user.username} className="w-20 h-20 rounded-2xl object-cover ring-4 ring-zinc-900 shadow-xl" />
                            ) : (
                                <div className="w-20 h-20 rounded-2xl bg-sky-500 flex items-center justify-center text-white font-bold text-3xl ring-4 ring-zinc-900 shadow-xl">
                                    {user?.username?.charAt(0).toUpperCase() || 'U'}
                                </div>
                            )}
                            <div className="flex gap-2">
                                <button onClick={() => setEditing(!editing)}
                                    className="px-4 py-1.5 bg-zinc-800 border border-zinc-700 text-zinc-300 hover:text-white rounded-xl text-sm font-medium transition-all">
                                    {editing ? 'Hủy' : '✏️ Chỉnh sửa'}
                                </button>
                                <button onClick={handleLogout}
                                    className="px-4 py-1.5 bg-red-500/10 border border-red-500/30 text-red-400 hover:bg-red-500/20 rounded-xl text-sm font-medium transition-all">
                                    Đăng xuất
                                </button>
                            </div>
                        </div>

                        {editing ? (
                            <div className="space-y-4">
                                <div>
                                    <label className="text-sm text-zinc-400 block mb-1">Tên hiển thị</label>
                                    <input value={form.username} onChange={(e) => setForm({ ...form, username: e.target.value })}
                                        className="w-full bg-zinc-800 border border-zinc-700 rounded-xl px-4 py-2.5 text-white focus:outline-none focus:ring-2 focus:ring-sky-500 transition-all" />
                                </div>
                                <div>
                                    <label className="text-sm text-zinc-400 block mb-1">URL ảnh đại diện</label>
                                    <input value={form.avatar_url} onChange={(e) => setForm({ ...form, avatar_url: e.target.value })} placeholder="https://..."
                                        className="w-full bg-zinc-800 border border-zinc-700 rounded-xl px-4 py-2.5 text-white placeholder-zinc-500 focus:outline-none focus:ring-2 focus:ring-sky-500 transition-all" />
                                </div>
                                <button onClick={handleSave} disabled={saving}
                                    className="w-full bg-sky-500 hover:bg-sky-400 disabled:opacity-50 text-white font-semibold py-2.5 rounded-xl transition-all">
                                    {saving ? 'Đang lưu...' : 'Lưu thay đổi'}
                                </button>
                            </div>
                        ) : (
                            <>
                                <h1 className="text-2xl font-bold text-white">{user?.username}</h1>
                                <p className="text-zinc-500 text-sm mt-1">{user?.phone}</p>
                            </>
                        )}
                    </div>
                </div>

                <h2 className="text-lg font-semibold text-white mb-4">Bài viết của tôi</h2>
                {posts.length === 0 ? (
                    <div className="text-center py-12 text-zinc-600">
                        <div className="text-4xl mb-2">📝</div>
                        <p>Chưa có bài viết nào</p>
                    </div>
                ) : (
                    <div className="space-y-4">
                        {posts.map((post) => (
                            <PostCard key={post.id} post={post} authorName={user?.username} authorAvatar={user?.avatar_url} authorId={user?.id}
                                onDeleted={(id) => setPosts((prev) => prev.filter((p) => p.id !== id))} />
                        ))}
                    </div>
                )}
            </div>
        </div>
    );
}

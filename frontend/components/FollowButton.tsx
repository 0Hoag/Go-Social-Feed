'use client';

import { useState } from 'react';
import { follow, unfollow } from '@/lib/api/follows';
import { useAuthStore } from '@/lib/stores/authStore';
import { toast } from 'sonner';

interface FollowButtonProps { targetId: string; initialFollowing?: boolean; initialFollowId?: string; }

export default function FollowButton({ targetId, initialFollowing = false, initialFollowId }: FollowButtonProps) {
    const { isAuthenticated, user } = useAuthStore();
    const [following, setFollowing] = useState(initialFollowing);
    const [followId, setFollowId] = useState<string | null>(initialFollowId || null);
    const [loading, setLoading] = useState(false);

    if (user?.id === targetId) return null;

    if (!isAuthenticated) return (
        <a href="/login" className="text-xs text-sky-400 hover:text-sky-300 transition-colors">Đăng nhập để theo dõi</a>
    );

    const handleToggle = async () => {
        if (loading) return;
        setLoading(true);
        try {
            if (following && followId) {
                await unfollow(followId); setFollowing(false); setFollowId(null); toast.success('Đã bỏ theo dõi');
            } else {
                const f = await follow(targetId); setFollowing(true); setFollowId(f.id); toast.success('Đã theo dõi!');
            }
        } catch { toast.error('Có lỗi xảy ra!'); } finally { setLoading(false); }
    };

    return (
        <button onClick={handleToggle} disabled={loading}
            className={`px-4 py-1.5 rounded-xl text-sm font-medium transition-all duration-200 disabled:opacity-50 ${following
                    ? 'bg-zinc-800 text-zinc-300 border border-zinc-700 hover:bg-red-500/10 hover:text-red-400 hover:border-red-500/30'
                    : 'bg-sky-500 hover:bg-sky-400 text-white shadow-md shadow-sky-500/20'
                }`}>
            {loading ? '...' : following ? 'Đang theo dõi' : '+ Theo dõi'}
        </button>
    );
}

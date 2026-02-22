'use client';

import { useState } from 'react';
import { createReaction, deleteReaction } from '@/lib/api/posts';
import { useAuthStore } from '@/lib/stores/authStore';
import { toast } from 'sonner';

interface ReactionBarProps {
    postId: string;
    initialCount?: number;
}

export default function ReactionBar({ postId, initialCount = 0 }: ReactionBarProps) {
    const { isAuthenticated } = useAuthStore();
    const [liked, setLiked] = useState(false);
    const [count, setCount] = useState(initialCount);
    const [reactionId, setReactionId] = useState<string | null>(null);
    const [loading, setLoading] = useState(false);

    const handleToggle = async () => {
        if (!isAuthenticated) { toast.error('Vui lòng đăng nhập để thả cảm xúc!'); return; }
        if (loading) return;
        setLoading(true);
        try {
            if (liked && reactionId) {
                await deleteReaction(reactionId);
                setLiked(false); setCount((c) => c - 1); setReactionId(null);
            } else {
                const r = await createReaction(postId, 'like');
                setLiked(true); setCount((c) => c + 1); setReactionId(r.id);
            }
        } catch { toast.error('Có lỗi xảy ra!'); } finally { setLoading(false); }
    };

    return (
        <button
            onClick={handleToggle}
            disabled={loading}
            className={`flex items-center gap-2 px-3 py-1.5 rounded-xl text-sm font-medium transition-all duration-200 ${liked
                    ? 'bg-sky-500/20 text-sky-400 border border-sky-500/40'
                    : 'bg-zinc-800 text-zinc-400 border border-zinc-700 hover:bg-zinc-700 hover:text-zinc-200'
                }`}
        >
            <svg className={`w-4 h-4 transition-transform ${liked ? 'scale-110' : ''}`}
                fill={liked ? 'currentColor' : 'none'} stroke="currentColor" viewBox="0 0 24 24">
                <path strokeLinecap="round" strokeLinejoin="round" strokeWidth={2}
                    d="M4.318 6.318a4.5 4.5 0 000 6.364L12 20.364l7.682-7.682a4.5 4.5 0 00-6.364-6.364L12 7.636l-1.318-1.318a4.5 4.5 0 00-6.364 0z" />
            </svg>
            <span>{count}</span>
        </button>
    );
}

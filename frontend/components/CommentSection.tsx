'use client';

import { useState, useEffect } from 'react';
import { getComments, createComment, deleteComment, Comment } from '@/lib/api/comments';
import { useAuthStore } from '@/lib/stores/authStore';
import { toast } from 'sonner';
import { formatDistanceToNow } from 'date-fns';
import { vi } from 'date-fns/locale';

interface CommentSectionProps { postId: string; }

export default function CommentSection({ postId }: CommentSectionProps) {
    const { isAuthenticated, user } = useAuthStore();
    const [comments, setComments] = useState<Comment[]>([]);
    const [loading, setLoading] = useState(false);
    const [submitting, setSubmitting] = useState(false);
    const [content, setContent] = useState('');

    useEffect(() => { loadComments(); }, [postId]);

    const loadComments = async () => {
        setLoading(true);
        try { const data = await getComments(postId); setComments(data.items || []); }
        catch { } finally { setLoading(false); }
    };

    const handleSubmit = async (e: React.FormEvent) => {
        e.preventDefault();
        if (!content.trim()) return;
        if (!isAuthenticated) { toast.error('Vui lòng đăng nhập để bình luận!'); return; }
        setSubmitting(true);
        try {
            const comment = await createComment(postId, content.trim());
            setComments((prev) => [comment, ...prev]);
            setContent('');
        } catch { toast.error('Không thể gửi bình luận!'); } finally { setSubmitting(false); }
    };

    const handleDelete = async (id: string) => {
        try { await deleteComment(id); setComments((prev) => prev.filter((c) => c.id !== id)); toast.success('Đã xóa bình luận'); }
        catch { toast.error('Không thể xóa bình luận!'); }
    };

    return (
        <div className="mt-4 pt-4 border-t border-zinc-800">
            {isAuthenticated ? (
                <form onSubmit={handleSubmit} className="flex gap-3 mb-4">
                    <div className="flex-shrink-0 w-8 h-8 rounded-full bg-sky-500 flex items-center justify-center text-white text-xs font-bold">
                        {user?.username?.charAt(0).toUpperCase() || 'U'}
                    </div>
                    <div className="flex-1 flex gap-2">
                        <input value={content} onChange={(e) => setContent(e.target.value)} placeholder="Viết bình luận..."
                            className="flex-1 bg-zinc-800 border border-zinc-700 rounded-xl px-4 py-2 text-sm text-white placeholder-zinc-500 focus:outline-none focus:ring-2 focus:ring-sky-500 focus:border-transparent transition-all" />
                        <button type="submit" disabled={submitting || !content.trim()}
                            className="px-4 py-2 bg-sky-500 hover:bg-sky-400 disabled:opacity-50 text-white text-sm rounded-xl transition-all duration-200 font-medium">
                            {submitting ? '...' : 'Gửi'}
                        </button>
                    </div>
                </form>
            ) : (
                <p className="text-zinc-500 text-sm mb-4">
                    <a href="/login" className="text-sky-400 hover:underline">Đăng nhập</a> để bình luận
                </p>
            )}

            {loading ? (
                <div className="flex justify-center py-4">
                    <div className="w-5 h-5 border-2 border-sky-500 border-t-transparent rounded-full animate-spin" />
                </div>
            ) : comments.length === 0 ? (
                <p className="text-zinc-600 text-sm text-center py-2">Chưa có bình luận nào</p>
            ) : (
                <div className="space-y-3">
                    {comments.map((comment) => (
                        <div key={comment.id} className="flex gap-3 group">
                            <div className="flex-shrink-0 w-8 h-8 rounded-full bg-zinc-700 flex items-center justify-center text-white text-xs font-bold">U</div>
                            <div className="flex-1">
                                <div className="bg-zinc-800 rounded-xl px-3 py-2">
                                    <p className="text-white text-sm">{comment.content}</p>
                                </div>
                                <div className="flex items-center gap-3 mt-1 px-1">
                                    <span className="text-xs text-zinc-600">
                                        {formatDistanceToNow(new Date(comment.created_at), { addSuffix: true, locale: vi })}
                                    </span>
                                    {user && comment.user_id === user.id && (
                                        <button onClick={() => handleDelete(comment.id)}
                                            className="text-xs text-zinc-600 hover:text-red-400 transition-colors opacity-0 group-hover:opacity-100">
                                            Xóa
                                        </button>
                                    )}
                                </div>
                            </div>
                        </div>
                    ))}
                </div>
            )}
        </div>
    );
}

'use client';

import { useState } from 'react';
import { Post, deletePost } from '@/lib/api/posts';
import { useAuthStore } from '@/lib/stores/authStore';
import ReactionBar from './ReactionBar';
import CommentSection from './CommentSection';
import { formatDistanceToNow } from 'date-fns';
import { vi } from 'date-fns/locale';
import Link from 'next/link';
import { toast } from 'sonner';

interface PostCardProps {
    post: Post;
    onDeleted?: (id: string) => void;
    authorName?: string;
    authorAvatar?: string;
    authorId?: string;
}

export default function PostCard({ post, onDeleted, authorName, authorAvatar, authorId }: PostCardProps) {
    const { user } = useAuthStore();
    const [showComments, setShowComments] = useState(false);
    const [deleting, setDeleting] = useState(false);

    const handleDelete = async () => {
        if (!confirm('Xóa bài viết này?')) return;
        setDeleting(true);
        try { await deletePost(post.id); onDeleted?.(post.id); toast.success('Đã xóa bài viết'); }
        catch { toast.error('Không thể xóa bài viết!'); } finally { setDeleting(false); }
    };

    const displayName = authorName || 'Người dùng';

    return (
        <article className="bg-zinc-900 border border-zinc-800 rounded-2xl p-5 hover:border-zinc-700 transition-all duration-200">
            {/* Author row */}
            <div className="flex items-start justify-between mb-3">
                <div className="flex items-center gap-3">
                    {authorAvatar ? (
                        <img src={authorAvatar} alt={displayName} className="w-10 h-10 rounded-full object-cover ring-2 ring-sky-500/30" />
                    ) : (
                        <div className="w-10 h-10 rounded-full bg-sky-500 flex items-center justify-center text-white font-bold text-sm">
                            {displayName.charAt(0).toUpperCase()}
                        </div>
                    )}
                    <div>
                        {authorId ? (
                            <Link href={`/profile/${authorId}`} className="font-semibold text-white hover:text-sky-400 transition-colors text-sm">
                                {displayName}
                            </Link>
                        ) : (
                            <span className="font-semibold text-white text-sm">{displayName}</span>
                        )}
                        <p className="text-xs text-zinc-500">
                            {formatDistanceToNow(new Date(post.created_at), { addSuffix: true, locale: vi })}
                            {post.permission === 'private' && (
                                <span className="ml-2 inline-flex items-center gap-1 text-zinc-600">
                                    <svg className="w-3 h-3" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                                        <path strokeLinecap="round" strokeLinejoin="round" strokeWidth={2} d="M12 15v2m-6 4h12a2 2 0 002-2v-6a2 2 0 00-2-2H6a2 2 0 00-2 2v6a2 2 0 002 2zm10-10V7a4 4 0 00-8 0v4h8z" />
                                    </svg>
                                    Riêng tư
                                </span>
                            )}
                        </p>
                    </div>
                </div>

                {user && authorId === user.id && (
                    <button onClick={handleDelete} disabled={deleting}
                        className="w-8 h-8 flex items-center justify-center rounded-lg hover:bg-red-500/10 text-zinc-600 hover:text-red-400 transition-all">
                        <svg className="w-4 h-4" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                            <path strokeLinecap="round" strokeLinejoin="round" strokeWidth={2} d="M19 7l-.867 12.142A2 2 0 0116.138 21H7.862a2 2 0 01-1.995-1.858L5 7m5 4v6m4-6v6m1-10V4a1 1 0 00-1-1h-4a1 1 0 00-1 1v3M4 7h16" />
                        </svg>
                    </button>
                )}
            </div>

            {/* Content */}
            {post.title && <h3 className="text-white font-semibold mb-2 text-base leading-snug">{post.title}</h3>}
            <p className="text-zinc-300 text-sm leading-relaxed whitespace-pre-wrap break-words">{post.content}</p>

            {post.source_url && (
                <a href={post.source_url} target="_blank" rel="noopener noreferrer"
                    className="mt-2 inline-flex items-center gap-1 text-xs text-sky-400 hover:text-sky-300 transition-colors">
                    <svg className="w-3 h-3" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                        <path strokeLinecap="round" strokeLinejoin="round" strokeWidth={2} d="M10 6H6a2 2 0 00-2 2v10a2 2 0 002 2h10a2 2 0 002-2v-4M14 4h6m0 0v6m0-6L10 14" />
                    </svg>
                    Xem nguồn
                </a>
            )}

            {/* Actions */}
            <div className="flex items-center gap-3 mt-4 pt-3 border-t border-zinc-800">
                <ReactionBar postId={post.id} />
                <button onClick={() => setShowComments(!showComments)}
                    className="flex items-center gap-2 px-3 py-1.5 rounded-xl text-sm font-medium bg-zinc-800 text-zinc-400 border border-zinc-700 hover:bg-zinc-700 hover:text-zinc-200 transition-all duration-200">
                    <svg className="w-4 h-4" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                        <path strokeLinecap="round" strokeLinejoin="round" strokeWidth={2} d="M8 12h.01M12 12h.01M16 12h.01M21 12c0 4.418-4.03 8-9 8a9.863 9.863 0 01-4.255-.949L3 20l1.395-3.72C3.512 15.042 3 13.574 3 12c0-4.418 4.03-8 9-8s9 3.582 9 8z" />
                    </svg>
                    Bình luận
                </button>
            </div>

            {showComments && <CommentSection postId={post.id} />}
        </article>
    );
}

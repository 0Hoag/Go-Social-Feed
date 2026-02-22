'use client';

import { useState } from 'react';
import { createPost } from '@/lib/api/posts';
import { toast } from 'sonner';

interface CreatePostModalProps { onClose: () => void; onCreated: () => void; }

export default function CreatePostModal({ onClose, onCreated }: CreatePostModalProps) {
    const [content, setContent] = useState('');
    const [permission, setPermission] = useState<'public' | 'private'>('public');
    const [loading, setLoading] = useState(false);

    const handleSubmit = async (e: React.FormEvent) => {
        e.preventDefault();
        if (!content.trim()) return;
        setLoading(true);
        try {
            await createPost({ content: content.trim(), permission });
            toast.success('Đã đăng bài viết!');
            onCreated(); onClose();
        } catch (err: any) {
            toast.error(err?.response?.data?.message || 'Không thể đăng bài!');
        } finally { setLoading(false); }
    };

    return (
        <div className="fixed inset-0 bg-black/70 backdrop-blur-sm z-50 flex items-center justify-center p-4" onClick={onClose}>
            <div className="w-full max-w-lg bg-zinc-900 border border-zinc-700 rounded-2xl shadow-2xl overflow-hidden" onClick={(e) => e.stopPropagation()}>
                {/* Header */}
                <div className="flex items-center justify-between px-6 py-4 border-b border-zinc-800">
                    <h2 className="text-lg font-semibold text-white">Tạo bài viết</h2>
                    <button onClick={onClose}
                        className="w-8 h-8 flex items-center justify-center rounded-full hover:bg-zinc-800 text-zinc-400 hover:text-white transition-all">
                        <svg className="w-5 h-5" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                            <path strokeLinecap="round" strokeLinejoin="round" strokeWidth={2} d="M6 18L18 6M6 6l12 12" />
                        </svg>
                    </button>
                </div>

                {/* Body */}
                <form onSubmit={handleSubmit} className="p-6 space-y-4">
                    <textarea value={content} onChange={(e) => setContent(e.target.value)}
                        placeholder="Bạn đang nghĩ gì vậy? 💭" rows={5} autoFocus
                        className="w-full bg-zinc-800 border border-zinc-700 rounded-xl px-4 py-3 text-white placeholder-zinc-500 focus:outline-none focus:ring-2 focus:ring-sky-500 resize-none transition-all" />

                    {/* Permission */}
                    <div className="flex items-center gap-3">
                        <span className="text-sm text-zinc-400">Hiển thị:</span>
                        <div className="flex gap-2">
                            {(['public', 'private'] as const).map((p) => (
                                <button key={p} type="button" onClick={() => setPermission(p)}
                                    className={`flex items-center gap-1.5 px-3 py-1.5 rounded-lg text-sm transition-all ${permission === p ? 'bg-sky-500 text-white' : 'bg-zinc-800 text-zinc-400 hover:bg-zinc-700'
                                        }`}>
                                    {p === 'public' ? (
                                        <><svg className="w-3.5 h-3.5" fill="none" stroke="currentColor" viewBox="0 0 24 24"><path strokeLinecap="round" strokeLinejoin="round" strokeWidth={2} d="M3.055 11H5a2 2 0 012 2v1a2 2 0 002 2 2 2 0 012 2v2.945M8 3.935V5.5A2.5 2.5 0 0010.5 8h.5a2 2 0 012 2 2 2 0 104 0 2 2 0 012-2h1.064M15 20.488V18a2 2 0 012-2h3.064" /></svg>Công khai</>
                                    ) : (
                                        <><svg className="w-3.5 h-3.5" fill="none" stroke="currentColor" viewBox="0 0 24 24"><path strokeLinecap="round" strokeLinejoin="round" strokeWidth={2} d="M12 15v2m-6 4h12a2 2 0 002-2v-6a2 2 0 00-2-2H6a2 2 0 00-2 2v6a2 2 0 002 2zm10-10V7a4 4 0 00-8 0v4h8z" /></svg>Riêng tư</>
                                    )}
                                </button>
                            ))}
                        </div>
                    </div>

                    <button type="submit" disabled={loading || !content.trim()}
                        className="w-full bg-sky-500 hover:bg-sky-400 disabled:opacity-50 disabled:cursor-not-allowed text-white font-semibold py-3 rounded-xl transition-all duration-200 shadow-lg shadow-sky-500/20 flex items-center justify-center gap-2">
                        {loading ? (
                            <><svg className="animate-spin w-4 h-4" fill="none" viewBox="0 0 24 24"><circle className="opacity-25" cx="12" cy="12" r="10" stroke="currentColor" strokeWidth="4" /><path className="opacity-75" fill="currentColor" d="M4 12a8 8 0 018-8V0C5.373 0 0 5.373 0 12h4z" /></svg>Đang đăng...</>
                        ) : '🚀 Đăng bài'}
                    </button>
                </form>
            </div>
        </div>
    );
}

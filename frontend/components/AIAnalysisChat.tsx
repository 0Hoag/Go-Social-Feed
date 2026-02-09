"use client";

import { useState, useEffect } from 'react';
import { X, Sparkles, TrendingDown, TrendingUp, ChevronRight, Send, Loader2 } from 'lucide-react';
import ReactMarkdown from 'react-markdown';

interface AIAnalysisChatProps {
    isOpen: boolean;
    onClose: () => void;
    coinSymbol: string;
    coinName: string;
    currentPrice: number;
    priceChange: number;
}

interface AnalysisData {
    tldr: string;
    mainCause: {
        title: string;
        overview: string;
        implication: string;
        watchFor: string;
    };
    secondaryCause?: {
        title: string;
        overview: string;
        implication: string;
        watchFor: string;
    };
    outlook: {
        title: string;
        overview: string;
        implication: string;
        watchFor: string;
    };
    conclusion: {
        outlook: string;
        keyPoint: string;
    };
}

export default function AIAnalysisChat({ isOpen, onClose, coinSymbol, coinName, currentPrice, priceChange }: AIAnalysisChatProps) {
    const [isLoading, setIsLoading] = useState(false);
    const [analysis, setAnalysis] = useState<AnalysisData | null>(null);
    const [userQuestion, setUserQuestion] = useState('');
    const [chatHistory, setChatHistory] = useState<Array<{ role: 'user' | 'assistant', content: string }>>([]);

    const isNegative = priceChange < 0;
    const initialQuestion = `Tại sao giá ${coinName} ${isNegative ? 'giảm' : 'tăng'}?`;

    const [suggestedQuestions, setSuggestedQuestions] = useState<string[]>([
        `Yếu tố nào có thể ảnh hưởng đến giá của ${coinName} trong tương lai?`,
        `Mọi người đang nói gì về ${coinName}?`,
        `Có tin tức mới nhất nào về ${coinName}?`,
        `${coinName} là gì?`
    ]);

    const fetchAnalysis = async () => {
        setIsLoading(true);
        try {
            const response = await fetch('/api/ai-analysis', {
                method: 'POST',
                headers: { 'Content-Type': 'application/json' },
                body: JSON.stringify({
                    symbol: coinSymbol,
                    name: coinName,
                    price: currentPrice,
                    priceChange: priceChange,
                    timeframe: '24h'
                })
            });

            if (!response.ok) {
                throw new Error('Failed to fetch analysis');
            }

            const text = await response.text();
            if (!text) {
                throw new Error('Empty response');
            }

            const data = JSON.parse(text);
            setAnalysis(data.analysis);
        } catch (error) {
            console.error('Failed to fetch analysis:', error);
        } finally {
            setIsLoading(false);
        }
    };

    const handleSendQuestion = async () => {
        if (!userQuestion.trim()) return;

        const newUserMsg = { role: 'user' as const, content: userQuestion };
        setChatHistory(prev => [...prev, newUserMsg]);
        setUserQuestion('');

        // Add a temporary loading message
        const loadingId = Date.now();
        setChatHistory(prev => [...prev, { role: 'assistant', content: 'THINKING_INDICATOR', id: loadingId } as any]);

        try {
            const response = await fetch('/api/ai-chat', {
                method: 'POST',
                headers: { 'Content-Type': 'application/json' },
                body: JSON.stringify({
                    message: newUserMsg.content,
                    coinSymbol,
                    coinName,
                    currentPrice,
                    priceChange
                })
            });

            const data = await response.json();

            // Parse reply and suggestions
            const fullReply = data.reply;
            const parts = fullReply.split('---QUESTIONS---');
            const mainContent = parts[0].trim();

            if (parts.length > 1) {
                const questions = parts[1]
                    .split('\n')
                    .map((q: string) => q.trim())
                    .filter((q: string) => q.length > 0);
                if (questions.length > 0) {
                    setSuggestedQuestions(questions);
                }
            }

            // Updates message, removing loading
            setChatHistory(prev => prev.map(msg =>
                (msg as any).id === loadingId ? { role: 'assistant', content: mainContent } : msg
            ));
        } catch (error) {
            setChatHistory(prev => prev.map(msg =>
                (msg as any).id === loadingId ? { role: 'assistant', content: 'Xin lỗi, tôi không thể trả lời ngay lúc này.' } : msg
            ));
        }
    };

    // Fetch analysis when chat opens
    useEffect(() => {
        if (isOpen && !analysis && !isLoading) {
            fetchAnalysis();
        }
    }, [isOpen]);

    if (!isOpen) return null;

    return (
        <div className="fixed inset-0 z-50 flex items-center justify-center bg-black/60 backdrop-blur-sm animate-fadeIn">
            <div className="relative w-full max-w-4xl h-[90vh] bg-black/95 rounded-2xl shadow-2xl border border-white/10 flex flex-col animate-slideInRight">
                {/* Header */}
                <div className="flex items-center justify-between px-6 py-4 border-b border-white/10">
                    <div className="flex items-center gap-3">
                        <div className="w-10 h-10 rounded-full bg-gradient-to-br from-blue-500 to-teal-400 flex items-center justify-center">
                            <Sparkles className="w-5 h-5 text-white" />
                        </div>
                        <div>
                            <h2 className="text-lg font-bold text-white">CryptoCheck AI</h2>
                            <p className="text-xs text-gray-400">Phân tích thị trường</p>
                        </div>
                    </div>
                    <button
                        onClick={onClose}
                        className="w-8 h-8 rounded-lg bg-white/5 hover:bg-white/10 flex items-center justify-center transition-colors"
                    >
                        <X className="w-5 h-5 text-gray-400" />
                    </button>
                </div>

                {/* Content */}
                <div className="flex-1 overflow-y-auto px-6 py-4 space-y-6">
                    {/* User Question */}
                    <div className="flex justify-end">
                        <div className="bg-gradient-to-r from-blue-600 to-teal-500 text-white px-4 py-2 rounded-2xl rounded-tr-sm max-w-md shadow-lg shadow-blue-500/20">
                            {initialQuestion}
                        </div>
                    </div>

                    {/* Loading State */}
                    {isLoading && (
                        <div className="flex justify-start">
                            <div className="flex items-center gap-3 p-4 bg-white/5 rounded-xl border border-white/10 animate-pulse w-fit">
                                <Loader2 className="w-5 h-5 text-teal-400 animate-spin" />
                                <span className="text-gray-400 text-sm">CryptoCheck AI đang phân tích dữ liệu thị trường...</span>
                            </div>
                        </div>
                    )}

                    {/* Analysis Content */}
                    {analysis && !isLoading && (
                        <div className="space-y-6">
                            {/* TLDR Section */}
                            <div className="bg-gradient-to-br from-blue-900/40 to-teal-900/40 border border-blue-500/30 rounded-xl p-5">
                                <div className="flex items-center gap-2 mb-3">
                                    <Sparkles className="w-5 h-5 text-teal-400" />
                                    <h3 className="text-lg font-bold text-white">TLDR</h3>
                                </div>
                                <p className="text-gray-300 leading-relaxed">{analysis.tldr}</p>
                            </div>

                            {/* Main Cause */}
                            <div className="space-y-3">
                                <div className="flex items-center gap-2">
                                    <div className="w-6 h-6 rounded-full bg-blue-600 flex items-center justify-center text-white text-sm font-bold">1</div>
                                    <h4 className="text-white font-bold">{analysis.mainCause.title}</h4>
                                </div>
                                <div className="pl-8 space-y-2 text-sm">
                                    <p className="text-gray-300"><span className="font-semibold text-white">Tổng quan:</span> {analysis.mainCause.overview}</p>
                                    <p className="text-gray-300"><span className="font-semibold text-white">Ý nghĩa:</span> {analysis.mainCause.implication}</p>
                                    <p className="text-gray-300"><span className="font-semibold text-white">Cần theo dõi:</span> {analysis.mainCause.watchFor}</p>
                                </div>
                            </div>

                            {/* Secondary Cause */}
                            {analysis.secondaryCause && (
                                <div className="space-y-3">
                                    <div className="flex items-center gap-2">
                                        <div className="w-6 h-6 rounded-full bg-blue-600 flex items-center justify-center text-white text-sm font-bold">2</div>
                                        <h4 className="text-white font-bold">{analysis.secondaryCause.title}</h4>
                                    </div>
                                    <div className="pl-8 space-y-2 text-sm">
                                        <p className="text-gray-300"><span className="font-semibold text-white">Tổng quan:</span> {analysis.secondaryCause.overview}</p>
                                        <p className="text-gray-300"><span className="font-semibold text-white">Ý nghĩa:</span> {analysis.secondaryCause.implication}</p>
                                        <p className="text-gray-300"><span className="font-semibold text-white">Cần theo dõi:</span> {analysis.secondaryCause.watchFor}</p>
                                    </div>
                                </div>
                            )}

                            {/* Outlook */}
                            <div className="space-y-3">
                                <div className="flex items-center gap-2">
                                    <div className="w-6 h-6 rounded-full bg-blue-600 flex items-center justify-center text-white text-sm font-bold">{analysis.secondaryCause ? '3' : '2'}</div>
                                    <h4 className="text-white font-bold">{analysis.outlook.title}</h4>
                                </div>
                                <div className="pl-8 space-y-2 text-sm">
                                    <p className="text-gray-300"><span className="font-semibold text-white">Tổng quan:</span> {analysis.outlook.overview}</p>
                                    <p className="text-gray-300"><span className="font-semibold text-white">Ý nghĩa:</span> {analysis.outlook.implication}</p>
                                    <p className="text-gray-300"><span className="font-semibold text-white">Cần theo dõi:</span> {analysis.outlook.watchFor}</p>
                                </div>
                            </div>

                            {/* Conclusion */}
                            <div className="bg-gradient-to-r from-blue-900/30 to-teal-900/30 border border-teal-500/30 rounded-xl p-5">
                                <h4 className="text-white font-bold mb-3">Kết luận</h4>
                                <p className="text-gray-300 mb-2"><span className="font-semibold text-white">Triển vọng thị trường:</span> {analysis.conclusion.outlook}</p>
                                <p className="text-gray-300"><span className="font-semibold text-white">Điểm cần lưu ý:</span> {analysis.conclusion.keyPoint}</p>
                            </div>



                            {/* Chat History */}
                            {chatHistory.map((msg, idx) => (
                                <div key={idx} className={`flex ${msg.role === 'user' ? 'justify-end' : 'justify-start'}`}>
                                    <div className={`py-2 ${msg.role === 'user'
                                        ? 'bg-gradient-to-r from-blue-600 to-teal-500 text-white rounded-2xl rounded-tr-sm shadow-lg shadow-blue-500/20 px-4 max-w-md'
                                        : 'w-full px-1' // assistant full width, no background
                                        }`}>
                                        {msg.role === 'assistant' ? (
                                            msg.content === 'THINKING_INDICATOR' ? (
                                                <div className="flex items-center gap-3 p-4 bg-white/5 rounded-xl border border-white/10 animate-pulse w-fit">
                                                    <Loader2 className="w-5 h-5 text-teal-400 animate-spin" />
                                                    <span className="text-gray-400 text-sm">CryptoCheck AI đang suy nghĩ...</span>
                                                </div>
                                            ) : (
                                                <div className="text-sm animate-fadeIn" style={{ counterReset: 'chat-section' }}>
                                                    <ReactMarkdown components={{
                                                        h3: ({ node, ...props }) => (
                                                            <div className="flex items-center gap-2 mt-6 mb-3">
                                                                <div
                                                                    className="w-6 h-6 rounded-full bg-blue-600 flex items-center justify-center text-white text-sm font-bold shrink-0 before:content-[counter(chat-section)]"
                                                                    style={{ counterIncrement: 'chat-section' }}
                                                                />
                                                                <h3 className="text-base font-bold text-white" {...props} />
                                                            </div>
                                                        ),
                                                        p: ({ node, ...props }) => <p className="mb-2 last:mb-0 leading-relaxed text-gray-300 pl-8" {...props} />,
                                                        ul: ({ node, ...props }) => <ul className="list-disc pl-12 mb-3 space-y-1 text-gray-300" {...props} />,
                                                        li: ({ node, ...props }) => <li className="marker:text-blue-500" {...props} />,
                                                        strong: ({ node, ...props }) => <strong className="text-white font-semibold" {...props} />,
                                                        blockquote: ({ node, ...props }) => (
                                                            <div className="bg-gradient-to-br from-blue-900/40 to-teal-900/40 border border-blue-500/30 rounded-xl p-4 my-4 flex gap-3">
                                                                <Sparkles className="w-5 h-5 text-teal-400 shrink-0 mt-0.5" />
                                                                <div className="text-gray-300 leading-relaxed">
                                                                    {props.children}
                                                                </div>
                                                            </div>
                                                        )
                                                    }}>
                                                        {msg.content}
                                                    </ReactMarkdown>
                                                </div>
                                            )
                                        ) : (
                                            <span>{msg.content}</span>
                                        )}
                                    </div>
                                </div>
                            ))}

                            {/* Suggested Questions */}
                            <div className="space-y-3 pt-4 border-t border-white/10">
                                <h4 className="text-white font-semibold text-sm">Gợi ý tiếp theo</h4>
                                <div className="space-y-2">
                                    {suggestedQuestions.map((question, idx) => (
                                        <button
                                            key={idx}
                                            className="w-full text-left px-4 py-3 bg-white/5 hover:bg-white/10 border border-white/10 rounded-xl text-sm text-gray-300 flex items-center justify-between group transition-all"
                                            onClick={() => setUserQuestion(question)}
                                        >
                                            <span>{question}</span>
                                            <ChevronRight className="w-4 h-4 text-gray-500 group-hover:text-blue-400 transition-colors" />
                                        </button>
                                    ))}
                                </div>
                            </div>
                        </div>
                    )}
                </div>

                {/* Input Box */}
                <div className="border-t border-white/10 px-6 py-4">
                    <div className="flex items-center gap-3">
                        <input
                            type="text"
                            value={userQuestion}
                            onChange={(e) => setUserQuestion(e.target.value)}
                            onKeyPress={(e) => e.key === 'Enter' && handleSendQuestion()}
                            placeholder="Hỏi CryptoCheck AI..."
                            className="flex-1 bg-white/5 border border-white/10 rounded-xl px-4 py-3 text-white placeholder-gray-500 focus:outline-none focus:border-blue-500/50 transition-colors"
                        />
                        <button
                            onClick={handleSendQuestion}
                            disabled={!userQuestion.trim()}
                            className="w-12 h-12 rounded-xl bg-gradient-to-r from-blue-600 to-teal-500 hover:from-blue-700 hover:to-teal-600 disabled:opacity-30 disabled:cursor-not-allowed flex items-center justify-center transition-all shadow-lg shadow-blue-500/20"
                        >
                            <Send className="w-5 h-5 text-white" />
                        </button>
                    </div>
                </div>
            </div>
        </div>
    );
}

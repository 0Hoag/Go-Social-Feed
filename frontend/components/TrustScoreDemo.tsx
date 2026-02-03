"use client";

import { Shield, AlertTriangle, CheckCircle, XCircle } from "lucide-react";
import { useState } from "react";

export default function TrustScoreDemo() {
    const [selectedExample, setSelectedExample] = useState<"safe" | "risky">("risky");

    const examples = {
        risky: {
            name: "$HOANG_COIN",
            score: 35,
            level: "RỦI RO CAO",
            color: "red",
            issues: [
                { type: "critical", icon: XCircle, text: "Admin có thể khóa ví của bạn", impact: -40 },
                { type: "warning", icon: AlertTriangle, text: "Thanh khoản (Liquidity) chưa được khóa", impact: -25 },
                { type: "success", icon: CheckCircle, text: "Mã nguồn đã được xác minh trên Etherscan", impact: +10 },
            ]
        },
        safe: {
            name: "$SAFE_TOKEN",
            score: 85,
            level: "AN TOÀN",
            color: "green",
            issues: [
                { type: "success", icon: CheckCircle, text: "Không phát hiện hàm nguy hiểm", impact: +30 },
                { type: "success", icon: CheckCircle, text: "Liquidity đã khóa 12 tháng", impact: +25 },
                { type: "success", icon: CheckCircle, text: "Ownership đã renounce", impact: +20 },
                { type: "success", icon: CheckCircle, text: "Mã nguồn đã verify", impact: +10 },
            ]
        }
    };

    const example = examples[selectedExample];
    const scorePercentage = example.score;

    const getScoreColor = (score: number) => {
        if (score >= 70) return "from-green-500 to-emerald-600";
        if (score >= 40) return "from-yellow-500 to-orange-500";
        return "from-red-500 to-orange-600";
    };

    const getLevelColor = (color: string) => {
        if (color === "green") return "text-green-400 bg-green-500/10 border-green-500/30";
        if (color === "yellow") return "text-yellow-400 bg-yellow-500/10 border-yellow-500/30";
        return "text-red-400 bg-red-500/10 border-red-500/30";
    };

    return (
        <section id="demo" className="py-20 px-4">
            <div className="max-w-5xl mx-auto">
                <div className="text-center mb-12">
                    <h2 className="text-4xl md:text-5xl font-black text-white mb-4">
                        Live Demo: Kết quả Quét
                    </h2>
                    <p className="text-gray-400 text-lg">
                        Xem cách ChainGuardian AI phân tích Smart Contract trong thực tế
                    </p>
                </div>

                {/* Example Selector */}
                <div className="flex justify-center gap-4 mb-8">
                    <button
                        onClick={() => setSelectedExample("risky")}
                        className={`px-6 py-3 rounded-lg font-semibold transition-all ${selectedExample === "risky"
                            ? "bg-red-500/20 text-red-400 border-2 border-red-500/50"
                            : "bg-white/5 text-gray-400 border-2 border-white/10 hover:bg-white/10"
                            }`}
                    >
                        Ví dụ: Rủi ro cao
                    </button>
                    <button
                        onClick={() => setSelectedExample("safe")}
                        className={`px-6 py-3 rounded-lg font-semibold transition-all ${selectedExample === "safe"
                            ? "bg-green-500/20 text-green-400 border-2 border-green-500/50"
                            : "bg-white/5 text-gray-400 border-2 border-white/10 hover:bg-white/10"
                            }`}
                    >
                        Ví dụ: An toàn
                    </button>
                </div>

                {/* Demo Card */}
                <div className="bg-gradient-to-br from-white/5 to-white/[0.02] backdrop-blur-xl border border-white/10 rounded-3xl p-8 md:p-12 shadow-2xl">
                    {/* Contract Name */}
                    <div className="flex items-center gap-3 mb-8">
                        <div className="w-12 h-12 rounded-xl bg-gradient-to-br from-cyan-500/20 to-blue-500/20 flex items-center justify-center">
                            <Shield className="w-6 h-6 text-cyan-400" />
                        </div>
                        <div>
                            <div className="text-sm text-gray-500 uppercase tracking-wider">Contract Address</div>
                            <div className="text-xl font-bold text-white">{example.name}</div>
                        </div>
                    </div>

                    {/* Trust Score */}
                    <div className="mb-8">
                        <div className="flex items-end justify-between mb-4">
                            <div>
                                <div className="text-sm text-gray-500 uppercase tracking-wider mb-2">Trust Score</div>
                                <div className="text-6xl font-black text-white">{example.score}<span className="text-3xl text-gray-500">/100</span></div>
                            </div>
                            <div className={`px-4 py-2 rounded-xl border-2 font-bold text-lg ${getLevelColor(example.color)}`}>
                                {example.level}
                            </div>
                        </div>

                        {/* Progress Bar */}
                        <div className="relative h-4 bg-gray-800/50 rounded-full overflow-hidden">
                            <div
                                className={`absolute top-0 left-0 h-full bg-gradient-to-r ${getScoreColor(example.score)} transition-all duration-1000 ease-out`}
                                style={{ width: `${scorePercentage}%` }}
                            />
                        </div>
                    </div>

                    {/* Security Issues */}
                    <div>
                        <h3 className="text-lg font-bold text-white mb-4 flex items-center gap-2">
                            <Shield className="w-5 h-5 text-cyan-400" />
                            Chi tiết phân tích
                        </h3>
                        <div className="space-y-3">
                            {example.issues.map((issue, index) => {
                                const Icon = issue.icon;
                                const iconColor = issue.type === "critical" ? "text-red-400" :
                                    issue.type === "warning" ? "text-orange-400" :
                                        "text-green-400";

                                return (
                                    <div
                                        key={index}
                                        className="flex items-start gap-3 p-4 bg-white/5 rounded-xl border border-white/10 hover:bg-white/10 transition-colors"
                                    >
                                        <Icon className={`w-5 h-5 mt-0.5 flex-shrink-0 ${iconColor}`} />
                                        <div className="flex-1">
                                            <div className="text-white font-medium">{issue.text}</div>
                                        </div>
                                        <div className={`text-sm font-bold ${issue.impact > 0 ? "text-green-400" : "text-red-400"}`}>
                                            {issue.impact > 0 ? "+" : ""}{issue.impact}đ
                                        </div>
                                    </div>
                                );
                            })}
                        </div>
                    </div>

                    {/* CTA */}
                    <div className="mt-8 pt-8 border-t border-white/10">
                        <button className="w-full py-4 bg-gradient-to-r from-cyan-500 to-blue-600 text-white font-bold rounded-xl hover:shadow-lg hover:shadow-cyan-500/50 transition-all duration-300 flex items-center justify-center gap-2">
                            <Shield className="w-5 h-5" />
                            Quét Contract của bạn ngay
                        </button>
                    </div>
                </div>
            </div>
        </section>
    );
}

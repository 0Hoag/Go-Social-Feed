"use client";

import { AlertTriangle, Lock, TrendingDown, Ban, Unlock, DollarSign } from "lucide-react";
import { useLanguage } from "@/context/LanguageContext";

export default function SecurityIssues() {
    const { t } = useLanguage();

    const issues = [
        {
            icon: Lock,
            title: t.security_issues.honeypot,
            description: t.security_issues.honeypot_desc,
            severity: "critical",
            color: "red"
        },
        {
            icon: TrendingDown,
            title: t.security_issues.rug_pull,
            description: t.security_issues.rug_pull_desc,
            severity: "critical",
            color: "red"
        },
        {
            icon: Ban,
            title: t.security_issues.blacklist,
            description: t.security_issues.blacklist_desc,
            severity: "warning",
            color: "orange"
        },
        {
            icon: AlertTriangle,
            title: t.security_issues.backdoor,
            description: t.security_issues.backdoor_desc,
            severity: "warning",
            color: "orange"
        },
        {
            icon: Unlock,
            title: t.security_issues.unlocked_liquidity,
            description: t.security_issues.unlocked_liquidity_desc,
            severity: "medium",
            color: "yellow"
        },
        {
            icon: DollarSign,
            title: t.security_issues.high_tax,
            description: t.security_issues.high_tax_desc,
            severity: "medium",
            color: "yellow"
        }
    ];

    const getSeverityColor = (color: string) => {
        if (color === "red") return "from-red-500 to-orange-600";
        if (color === "orange") return "from-orange-500 to-yellow-500";
        return "from-yellow-500 to-amber-500";
    };

    const getSeverityBadge = (severity: string) => {
        if (severity === "critical") return "bg-red-500/10 text-red-400 border-red-500/30";
        if (severity === "warning") return "bg-orange-500/10 text-orange-400 border-orange-500/30";
        return "bg-yellow-500/10 text-yellow-400 border-yellow-500/30";
    };

    return (
        <section className="py-20 px-4 bg-gradient-to-b from-white/[0.02] to-transparent">
            <div className="max-w-7xl mx-auto">
                <div className="text-center mb-16">
                    <h2 className="text-4xl md:text-5xl font-black text-white mb-4">
                        {t.security_issues.title_prefix} <span className="text-transparent bg-clip-text bg-gradient-to-r from-red-400 to-orange-500">{t.security_issues.title_highlight}</span>
                    </h2>
                    <p className="text-gray-400 text-lg max-w-2xl mx-auto">
                        {t.security_issues.subtitle}
                    </p>
                </div>

                <div className="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-6">
                    {issues.map((issue, index) => {
                        const Icon = issue.icon;
                        return (
                            <div
                                key={index}
                                className="group relative bg-gradient-to-br from-white/5 to-white/[0.02] backdrop-blur-sm border border-white/10 rounded-2xl p-6 hover:border-red-500/30 transition-all duration-300 hover:shadow-xl hover:shadow-red-500/10"
                            >
                                {/* Severity Badge */}
                                <div className={`absolute top-4 right-4 px-3 py-1 rounded-lg border text-xs font-bold uppercase ${getSeverityBadge(issue.severity)}`}>
                                    {issue.severity}
                                </div>

                                {/* Icon */}
                                <div className={`w-14 h-14 rounded-xl bg-gradient-to-br ${getSeverityColor(issue.color)} bg-opacity-10 flex items-center justify-center mb-4 group-hover:scale-110 transition-transform duration-300`}>
                                    <Icon className="w-7 h-7 text-white" strokeWidth={2} />
                                </div>

                                {/* Content */}
                                <h3 className="text-xl font-bold text-white mb-2">
                                    {issue.title}
                                </h3>
                                <p className="text-gray-400 leading-relaxed">
                                    {issue.description}
                                </p>

                                {/* Hover Gradient */}
                                <div className={`absolute inset-0 rounded-2xl bg-gradient-to-br ${getSeverityColor(issue.color)} opacity-0 group-hover:opacity-5 transition-opacity duration-300`} />
                            </div>
                        );
                    })}
                </div>

                {/* Bottom Note */}
                <div className="mt-12 text-center">
                    <p className="text-gray-500 text-sm">
                        {t.security_issues.note}
                    </p>
                </div>
            </div>
        </section>
    );
}

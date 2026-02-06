"use client";

import { Zap, Bot, Target, Bell, DollarSign, Globe } from "lucide-react";
import { useLanguage } from "@/context/LanguageContext";

export default function FeaturesGrid() {
    const { t } = useLanguage();

    const features = [
        {
            icon: Zap,
            title: t.features.f1_title,
            description: t.features.f1_desc,
            color: "from-yellow-400 to-orange-500"
        },
        {
            icon: Bot,
            title: t.features.f2_title,
            description: t.features.f2_desc,
            color: "from-cyan-400 to-blue-500"
        },
        {
            icon: Target,
            title: t.features.f3_title,
            description: t.features.f3_desc,
            color: "from-green-400 to-emerald-500"
        },
        {
            icon: Bell,
            title: t.features.f4_title,
            description: t.features.f4_desc,
            color: "from-purple-400 to-pink-500"
        },
        {
            icon: DollarSign,
            title: t.features.f5_title,
            description: t.features.f5_desc,
            color: "from-green-400 to-teal-500"
        },
        {
            icon: Globe,
            title: t.features.f6_title,
            description: t.features.f6_desc,
            color: "from-blue-400 to-indigo-500"
        }
    ];

    return (
        <section className="py-20 px-4 bg-gradient-to-b from-transparent to-white/[0.02]">
            <div className="max-w-7xl mx-auto">
                <div className="text-center mb-16">
                    <h2 className="text-4xl md:text-5xl font-black text-white mb-4">
                        {t.features.title_prefix} <span className="text-transparent bg-clip-text bg-gradient-to-r from-cyan-400 to-blue-500">{t.features.title_highlight}</span>{t.features.title_suffix}
                    </h2>
                    <p className="text-gray-400 text-lg max-w-2xl mx-auto">
                        {t.features.subtitle}
                    </p>
                </div>

                <div className="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-6">
                    {features.map((feature, index) => {
                        const Icon = feature.icon;
                        return (
                            <div
                                key={index}
                                className="group relative bg-gradient-to-br from-white/5 to-white/[0.02] backdrop-blur-sm border border-white/10 rounded-2xl p-8 hover:border-cyan-500/50 transition-all duration-300 hover:shadow-xl hover:shadow-cyan-500/10 hover:-translate-y-1"
                            >
                                {/* Icon */}
                                <div className={`w-14 h-14 rounded-xl bg-gradient-to-br ${feature.color} bg-opacity-10 flex items-center justify-center mb-6 group-hover:scale-110 transition-transform duration-300`}>
                                    <Icon className="w-7 h-7 text-white" strokeWidth={2} />
                                </div>

                                {/* Content */}
                                <h3 className="text-xl font-bold text-white mb-3">
                                    {feature.title}
                                </h3>
                                <p className="text-gray-400 leading-relaxed">
                                    {feature.description}
                                </p>

                                {/* Hover Gradient */}
                                <div className={`absolute inset-0 rounded-2xl bg-gradient-to-br ${feature.color} opacity-0 group-hover:opacity-5 transition-opacity duration-300`} />
                            </div>
                        );
                    })}
                </div>
            </div>
        </section>
    );
}

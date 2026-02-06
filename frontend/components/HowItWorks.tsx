"use client";

import { Search, Cpu, ShieldCheck } from "lucide-react";
import { useLanguage } from "@/context/LanguageContext";

export default function HowItWorks() {
    const { t } = useLanguage();

    const steps = [
        {
            number: 1,
            icon: Search,
            title: t.how_it_works.step1_title,
            description: t.how_it_works.step1_desc,
            color: "from-cyan-400 to-blue-500"
        },
        {
            number: 2,
            icon: Cpu,
            title: t.how_it_works.step2_title,
            description: t.how_it_works.step2_desc,
            color: "from-blue-400 to-indigo-500"
        },
        {
            number: 3,
            icon: ShieldCheck,
            title: t.how_it_works.step3_title,
            description: t.how_it_works.step3_desc,
            color: "from-green-400 to-emerald-500"
        }
    ];

    return (
        <section id="how-it-works" className="py-20 px-4">
            <div className="max-w-7xl mx-auto">
                <div className="text-center mb-16">
                    <h2 className="text-4xl md:text-5xl font-black text-white mb-4">
                        {t.how_it_works.title}
                    </h2>
                    <p className="text-gray-400 text-lg max-w-2xl mx-auto">
                        {t.how_it_works.subtitle}
                    </p>
                </div>

                <div className="grid grid-cols-1 md:grid-cols-3 gap-8">
                    {steps.map((step, index) => {
                        const Icon = step.icon;
                        return (
                            <div
                                key={index}
                                className="relative group h-full"
                            >
                                {/* Connection Arrow (Except last) */}
                                {index < steps.length - 1 && (
                                    <div className="hidden md:block absolute top-1/2 -right-4 w-8 h-8 z-20 text-gray-700 transform -translate-y-1/2 translate-x-1/2">
                                        →
                                    </div>
                                )}

                                {/* Step Card */}
                                <div className="bg-gradient-to-br from-white/5 to-white/[0.02] backdrop-blur-sm border border-white/10 rounded-2xl p-6 hover:border-cyan-500/50 transition-all duration-300 hover:shadow-xl hover:shadow-cyan-500/10 h-full flex flex-col items-center text-center">
                                    {/* Step Number */}
                                    <div className={`absolute -top-4 -left-4 w-10 h-10 rounded-xl bg-gradient-to-br ${step.color} flex items-center justify-center font-bold text-white text-lg shadow-lg`}>
                                        {step.number}
                                    </div>

                                    {/* Icon */}
                                    <div className={`w-16 h-16 rounded-2xl bg-gradient-to-br ${step.color} bg-opacity-10 flex items-center justify-center mb-6 mt-2 group-hover:scale-110 transition-transform duration-300`}>
                                        <Icon className="w-8 h-8 text-white" strokeWidth={2} />
                                    </div>

                                    {/* Content */}
                                    <h3 className="text-xl font-bold text-white mb-3">
                                        {step.title}
                                    </h3>
                                    <p className="text-gray-400 leading-relaxed">
                                        {step.description}
                                    </p>
                                </div>
                            </div>
                        );
                    })}
                </div>
            </div>
        </section>
    );
}

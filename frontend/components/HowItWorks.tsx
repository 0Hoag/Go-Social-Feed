import { Settings, Link2, Bot, Calculator, MessageSquare, LayoutDashboard } from "lucide-react";

export default function HowItWorks() {
    const steps = [
        {
            number: 1,
            icon: Settings,
            title: "Setup môi trường",
            description: "Cài đặt Go, Docker, Slither trên hệ thống",
            color: "from-cyan-400 to-blue-500"
        },
        {
            number: 2,
            icon: Link2,
            title: "Blockchain Connection",
            description: "Lấy Source Code Solidity từ Etherscan API",
            color: "from-blue-400 to-indigo-500"
        },
        {
            number: 3,
            icon: Bot,
            title: "Tích hợp AI",
            description: "Kết nối với OpenAI API để phân tích code",
            color: "from-purple-400 to-pink-500"
        },
        {
            number: 4,
            icon: Calculator,
            title: "Thuật toán Trust Score",
            description: "Tính điểm dựa trên mức độ nghiêm trọng",
            color: "from-pink-400 to-red-500"
        },
        {
            number: 5,
            icon: MessageSquare,
            title: "Telegram Bot",
            description: "Giao diện Bot: Gửi địa chỉ → Nhận kết quả",
            color: "from-green-400 to-emerald-500"
        },
        {
            number: 6,
            icon: LayoutDashboard,
            title: "Frontend Dashboard",
            description: "Hiển thị danh sách coin mới an toàn",
            color: "from-orange-400 to-yellow-500"
        }
    ];

    return (
        <section id="how-it-works" className="py-20 px-4">
            <div className="max-w-7xl mx-auto">
                <div className="text-center mb-16">
                    <h2 className="text-4xl md:text-5xl font-black text-white mb-4">
                        Cách hoạt động
                    </h2>
                    <p className="text-gray-400 text-lg max-w-2xl mx-auto">
                        6 bước xây dựng hệ thống phân tích bảo mật Smart Contract
                    </p>
                </div>

                <div className="relative">
                    {/* Connection Line */}
                    <div className="hidden lg:block absolute top-1/2 left-0 right-0 h-0.5 bg-gradient-to-r from-cyan-500/20 via-purple-500/20 to-orange-500/20 -translate-y-1/2" />

                    <div className="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-8">
                        {steps.map((step, index) => {
                            const Icon = step.icon;
                            return (
                                <div
                                    key={index}
                                    className="relative group"
                                >
                                    {/* Step Card */}
                                    <div className="bg-gradient-to-br from-white/5 to-white/[0.02] backdrop-blur-sm border border-white/10 rounded-2xl p-6 hover:border-cyan-500/50 transition-all duration-300 hover:shadow-xl hover:shadow-cyan-500/10 h-full">
                                        {/* Step Number */}
                                        <div className={`absolute -top-4 -left-4 w-12 h-12 rounded-xl bg-gradient-to-br ${step.color} flex items-center justify-center font-black text-white text-xl shadow-lg`}>
                                            {step.number}
                                        </div>

                                        {/* Icon */}
                                        <div className={`w-14 h-14 rounded-xl bg-gradient-to-br ${step.color} bg-opacity-10 flex items-center justify-center mb-4 mt-4 group-hover:scale-110 transition-transform duration-300`}>
                                            <Icon className="w-7 h-7 text-white" strokeWidth={2} />
                                        </div>

                                        {/* Content */}
                                        <h3 className="text-lg font-bold text-white mb-2">
                                            {step.title}
                                        </h3>
                                        <p className="text-gray-400 text-sm leading-relaxed">
                                            {step.description}
                                        </p>
                                    </div>
                                </div>
                            );
                        })}
                    </div>
                </div>
            </div>
        </section>
    );
}

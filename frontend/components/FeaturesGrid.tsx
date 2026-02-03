import { Zap, Bot, Target, Bell, DollarSign, Globe } from "lucide-react";

export default function FeaturesGrid() {
    const features = [
        {
            icon: Zap,
            title: "Tốc độ thực thời",
            description: "Kết quả phân tích chỉ trong 30-60 giây. Không cần chờ đợi hàng giờ.",
            color: "from-yellow-400 to-orange-500"
        },
        {
            icon: Bot,
            title: "AI Logic Auditor",
            description: "Sử dụng LLM để phát hiện lỗ hổng logic tinh vi mà máy quét thông thường bỏ qua.",
            color: "from-cyan-400 to-blue-500"
        },
        {
            icon: Target,
            title: "Dễ hiểu cho mọi người",
            description: "Trust Score 0-100 đơn giản. Không cần hiểu code Solidity phức tạp.",
            color: "from-green-400 to-emerald-500"
        },
        {
            icon: Bell,
            title: "Telegram Bot",
            description: "Cảnh báo real-time khi phát hiện hợp đồng mới nguy hiểm trên mạng lưới.",
            color: "from-purple-400 to-pink-500"
        },
        {
            icon: DollarSign,
            title: "Miễn phí cho cá nhân",
            description: "Bình dân hóa bảo mật. Chi phí cực thấp hoặc miễn phí cho nhà đầu tư cá nhân.",
            color: "from-green-400 to-teal-500"
        },
        {
            icon: Globe,
            title: "Đa mạng lưới",
            description: "Hỗ trợ Ethereum, BSC, Polygon và nhiều blockchain khác.",
            color: "from-blue-400 to-indigo-500"
        }
    ];

    return (
        <section className="py-20 px-4 bg-gradient-to-b from-transparent to-white/[0.02]">
            <div className="max-w-7xl mx-auto">
                <div className="text-center mb-16">
                    <h2 className="text-4xl md:text-5xl font-black text-white mb-4">
                        Tại sao chọn <span className="text-transparent bg-clip-text bg-gradient-to-r from-cyan-400 to-blue-500">ChainGuardian AI</span>?
                    </h2>
                    <p className="text-gray-400 text-lg max-w-2xl mx-auto">
                        Công nghệ tiên tiến kết hợp với trải nghiệm người dùng đơn giản
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

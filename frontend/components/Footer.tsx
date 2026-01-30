import Link from "next/link";
import { Send, Twitter, Facebook } from "lucide-react";

export default function Footer() {
    return (
        <footer className="bg-[#050505] border-t border-white/5 pt-16 pb-8 text-sm">
            <div className="max-w-[1600px] mx-auto px-4">
                <div className="grid grid-cols-1 lg:grid-cols-12 gap-12 mb-12">
                    {/* Brand Column */}
                    <div className="lg:col-span-3 space-y-6">
                        <Link href="/" className="block">
                            <span className="text-2xl font-bold text-blue-500 tracking-tight">
                                Syntax
                            </span>
                        </Link>
                        <p className="text-gray-400 leading-relaxed text-xs">
                            Nền tảng cung cấp thông tin dữ liệu và phân tích thị trường tiền điện tử chuyên sâu, cung cấp góc nhìn đa chiều, nhanh chóng và chính xác.
                        </p>
                        <div className="flex items-center gap-4">
                            <a href="#" className="w-8 h-8 rounded-full bg-white/5 flex items-center justify-center text-gray-400 hover:text-white hover:bg-white/10 transition-colors">
                                <Send className="w-4 h-4" />
                            </a>
                            <a href="#" className="w-8 h-8 rounded-full bg-white/5 flex items-center justify-center text-gray-400 hover:text-white hover:bg-white/10 transition-colors">
                                <Twitter className="w-4 h-4" />
                            </a>
                            <a href="#" className="w-8 h-8 rounded-full bg-white/5 flex items-center justify-center text-gray-400 hover:text-white hover:bg-white/10 transition-colors">
                                <Facebook className="w-4 h-4" />
                            </a>
                        </div>
                    </div>

                    {/* Spacer */}
                    <div className="hidden lg:block lg:col-span-1"></div>

                    {/* Links Column 1 */}
                    <div className="lg:col-span-2 space-y-4">
                        <h4 className="text-gray-500 font-bold text-xs uppercase tracking-wider mb-2">Thông tin</h4>
                        <ul className="space-y-3">
                            <li><Link href="#" className="text-gray-400 hover:text-white transition-colors text-xs">Về chúng tôi</Link></li>
                            <li><Link href="#" className="text-gray-400 hover:text-white transition-colors text-xs">Lộ trình phát triển</Link></li>
                            <li><Link href="#" className="text-gray-400 hover:text-white transition-colors text-xs">Liên hệ quảng cáo</Link></li>
                        </ul>
                    </div>

                    {/* Links Column 2 */}
                    <div className="lg:col-span-2 space-y-4">
                        <h4 className="text-gray-500 font-bold text-xs uppercase tracking-wider mb-2">Tài nguyên</h4>
                        <ul className="space-y-3">
                            <li><Link href="#" className="text-gray-400 hover:text-white transition-colors text-xs">Lịch kinh tế</Link></li>
                            <li><Link href="#" className="text-gray-400 hover:text-white transition-colors text-xs">Quyền chọn Deribit</Link></li>
                            <li><Link href="#" className="text-gray-400 hover:text-white transition-colors text-xs">Phân tích kỹ thuật</Link></li>
                        </ul>
                    </div>

                    {/* Links Column 3 */}
                    <div className="lg:col-span-2 space-y-4">
                        <h4 className="text-gray-500 font-bold text-xs uppercase tracking-wider mb-2">Pháp lý & Hỗ trợ</h4>
                        <ul className="space-y-3">
                            <li><Link href="#" className="text-gray-400 hover:text-white transition-colors text-xs">Điều khoản sử dụng</Link></li>
                            <li><Link href="#" className="text-gray-400 hover:text-white transition-colors text-xs">Chính sách bảo mật</Link></li>
                            <li><Link href="#" className="text-gray-400 hover:text-white transition-colors text-xs">Câu hỏi thường gặp</Link></li>
                        </ul>
                    </div>
                </div>

                {/* Bottom Bar */}
                <div className="border-t border-white/5 pt-8 mt-8">
                    <p className="text-[10px] text-gray-500 text-center leading-relaxed max-w-4xl mx-auto mb-4">
                        Cảnh báo rủi ro: Giao dịch tiền điện tử tiềm ẩn rủi ro cao. Thông tin trên Syntax chỉ mang tính chất tham khảo, không phải lời khuyên đầu tư. Bạn chịu trách nhiệm hoàn toàn về các quyết định giao dịch của mình.
                    </p>
                    <p className="text-[10px] text-gray-600 text-center">
                        © 2026 Syntax. Bản quyền nền tảng và phân tích. Tin tức thị trường được tổng hợp từ các nguồn uy tín.
                    </p>
                </div>
            </div>
        </footer>
    );
}

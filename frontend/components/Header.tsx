import Link from "next/link";
import { Shield } from "lucide-react";

export default function Header() {
    return (
        <header className="sticky top-0 z-50 w-full border-b border-white/5 bg-[#050505]/80 backdrop-blur-xl supports-[backdrop-filter]:bg-[#050505]/60">
            <div className="max-w-[1600px] mx-auto flex h-16 items-center justify-between px-4">
                {/* Logo Area */}
                <Link href="/" className="flex items-center gap-2 group">
                    <div className="w-9 h-9 rounded-lg bg-gradient-to-br from-cyan-500 to-blue-600 flex items-center justify-center group-hover:scale-105 transition-transform">
                        <Shield className="w-5 h-5 text-white" strokeWidth={2.5} />
                    </div>
                    <span className="text-xl font-black text-white tracking-tight">
                        Chain<span className="text-transparent bg-clip-text bg-gradient-to-r from-cyan-400 to-blue-500">Guardian</span> AI
                    </span>
                </Link>

                {/* Main Navigation - Centered */}
                <nav className="absolute left-1/2 -translate-x-1/2 flex items-center gap-8">
                    <Link href="/" className="text-sm font-medium text-gray-400 hover:text-white cursor-pointer transition-colors">Home</Link>
                    <Link href="/news" className="text-sm font-medium text-gray-400 hover:text-cyan-400 cursor-pointer transition-colors">News</Link>
                    <Link href="#demo" className="text-sm font-medium text-gray-400 hover:text-cyan-400 cursor-pointer transition-colors">Scan</Link>
                    <Link href="/analysis" className="text-sm font-medium text-gray-400 hover:text-cyan-400 cursor-pointer transition-colors">Analysis</Link>
                </nav>

                {/* Empty right space for balance */}
                <div className="w-[100px]"></div>
            </div>
        </header>
    );
}

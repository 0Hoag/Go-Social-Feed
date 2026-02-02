import Link from "next/link";
import { TrendingUp } from "lucide-react";

export default function Header() {
    return (
        <header className="sticky top-0 z-50 w-full border-b border-white/5 bg-[#050505]/80 backdrop-blur-xl supports-[backdrop-filter]:bg-[#050505]/60">
            <div className="max-w-[1600px] mx-auto flex h-16 items-center justify-between px-4">
                {/* Logo Area */}
                <Link href="/" className="flex items-center gap-2 group">
                    <span className="text-2xl font-bold text-blue-500 tracking-tight group-hover:text-blue-400 transition-colors">
                        Syntax
                    </span>
                </Link>

                {/* Main Navigation - Centered */}
                <nav className="absolute left-1/2 -translate-x-1/2 flex items-center gap-8">
                    <Link href="/" className="text-sm font-medium text-gray-400 hover:text-white cursor-pointer transition-colors">News</Link>
                    <Link href="/analysis" className="text-sm font-medium text-white cursor-pointer transition-colors">On-chain</Link>
                </nav>

                {/* Empty right space for balance */}
                <div className="w-[100px]"></div>
            </div>
        </header>
    );
}

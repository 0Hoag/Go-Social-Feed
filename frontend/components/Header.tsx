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

                {/* Main Navigation */}
                <nav className="flex items-center gap-8">
                    <Link href="/" className="text-sm font-medium text-gray-400 hover:text-white cursor-pointer transition-colors">News</Link>
                    <Link href="/analysis" className="text-sm font-medium text-white cursor-pointer transition-colors">Analysis</Link>
                    <span className="text-sm font-medium text-gray-400 hover:text-cyan-400 cursor-pointer transition-colors">Calendar</span>
                    <span className="text-sm font-medium text-gray-400 hover:text-cyan-400 cursor-pointer transition-colors">Options</span>
                    <span className="text-sm font-medium text-gray-400 hover:text-cyan-400 cursor-pointer transition-colors">On-chain</span>
                </nav>

                {/* Right Area (Search/Profile Placeholder) */}
                <div className="flex items-center gap-4">
                    <button className="text-gray-400 hover:text-white transition-colors">
                        <svg xmlns="http://www.w3.org/2000/svg" width="20" height="20" viewBox="0 0 24 24" fill="none" stroke="currentColor" strokeWidth="2" strokeLinecap="round" strokeLinejoin="round"><circle cx="11" cy="11" r="8" /><path d="m21 21-4.3-4.3" /></svg>
                    </button>
                    <div className="w-8 h-8 rounded-full bg-gray-800 border border-white/10"></div>
                </div>
            </div>
        </header>
    );
}

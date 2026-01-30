import Link from "next/link";
import { TrendingUp } from "lucide-react";

export default function Header() {
    return (
        <header className="sticky top-0 z-50 w-full border-b border-white/10 bg-gray-950/80 backdrop-blur-xl supports-[backdrop-filter]:bg-gray-950/60">
            <div className="container mx-auto flex h-16 items-center justify-between px-4">
                <Link href="/" className="flex items-center space-x-2 group">
                    <div className="rounded-lg bg-gradient-to-br from-cyan-500 to-blue-600 p-2 transition-transform group-hover:scale-110">
                        <TrendingUp className="h-5 w-5 text-white" />
                    </div>
                    <span className="text-xl font-bold bg-gradient-to-r from-cyan-400 to-blue-500 bg-clip-text text-transparent">
                        Crypto News Feed
                    </span>
                </Link>

                <nav className="flex items-center space-x-6">
                    <Link
                        href="/"
                        className="text-sm font-medium text-gray-300 transition-colors hover:text-cyan-400"
                    >
                        Home
                    </Link>
                    <a
                        href="https://github.com"
                        target="_blank"
                        rel="noopener noreferrer"
                        className="text-sm font-medium text-gray-300 transition-colors hover:text-cyan-400"
                    >
                        GitHub
                    </a>
                </nav>
            </div>
        </header>
    );
}

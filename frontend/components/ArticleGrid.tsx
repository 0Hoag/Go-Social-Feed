import { Post } from "@/lib/types";
import ArticleCard from "./ArticleCard";

interface ArticleGridProps {
    posts: Post[];
}

export default function ArticleGrid({ posts }: ArticleGridProps) {
    if (posts.length === 0) {
        return (
            <div className="flex flex-col items-center justify-center py-20 text-center">
                <div className="rounded-full bg-gray-800/50 p-6 mb-4">
                    <svg
                        className="h-12 w-12 text-gray-600"
                        fill="none"
                        viewBox="0 0 24 24"
                        stroke="currentColor"
                    >
                        <path
                            strokeLinecap="round"
                            strokeLinejoin="round"
                            strokeWidth={2}
                            d="M19 20H5a2 2 0 01-2-2V6a2 2 0 012-2h10a2 2 0 012 2v1m2 13a2 2 0 01-2-2V7m2 13a2 2 0 002-2V9a2 2 0 00-2-2h-2m-4-3H9M7 16h6M7 8h6v4H7V8z"
                        />
                    </svg>
                </div>
                <h3 className="text-xl font-semibold text-gray-300 mb-2">No articles found</h3>
                <p className="text-gray-500">Check back later for the latest crypto news!</p>
            </div>
        );
    }

    return (
        <div className="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-6">
            {posts.map((post) => (
                <ArticleCard key={post.id} post={post} />
            ))}
        </div>
    );
}

import { clsx, type ClassValue } from "clsx";
import { twMerge } from "tailwind-merge";
import { formatDistanceToNow, format } from "date-fns";

export function cn(...inputs: ClassValue[]) {
    return twMerge(clsx(inputs));
}

export function formatDate(dateString: string): string {
    const date = new Date(dateString);
    const now = new Date();
    const diffInHours = (now.getTime() - date.getTime()) / (1000 * 60 * 60);

    // If less than 24 hours, show relative time
    if (diffInHours < 24) {
        return formatDistanceToNow(date, { addSuffix: true });
    }

    // Otherwise show formatted date
    return format(date, "MMM d, yyyy");
}

export function extractDomain(url: string): string {
    try {
        const domain = new URL(url).hostname;
        // Remove 'www.' prefix if present
        return domain.replace(/^www\./, "");
    } catch {
        return url;
    }
}

export function getSourceName(url: string): string {
    const domain = extractDomain(url);

    if (domain.includes("coindesk")) return "Coindesk";
    if (domain.includes("cointelegraph")) return "Cointelegraph";

    return domain;
}

export function truncateText(text: string, maxLength: number): string {
    if (text.length <= maxLength) return text;
    return text.slice(0, maxLength).trim() + "...";
}

/**
 * Extract first image URL from markdown content
 * Matches: ![alt](url) or ![](url)
 */
export function extractImageUrl(content: string): string | null {
    if (!content) return null;
    const match = content.match(/!\[.*?\]\((https?:\/\/[^\)]+)\)/);
    return match ? match[1] : null;
}



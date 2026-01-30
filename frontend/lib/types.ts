export interface Post {
    id: string;
    pin: boolean;
    title: string;
    content: string;
    file_ids?: string[];
    tagged_target?: string[];
    permission: 'public' | 'justme';
    author_id: string;
    source_url: string;
    created_at: string;
    updated_at: string;
    deleted_at?: string | null;
}

export interface PaginationParams {
    page?: number;
    limit?: number;
    sort?: string;
}

export interface ApiResponse<T> {
    data: T;
    message?: string;
    error?: string;
}

export interface PostsResponse {
    posts: Post[];
    pagination?: {
        total: number;
        count: number;
        per_page: number;
        current_page: number;
        total_pages: number;
    };
}

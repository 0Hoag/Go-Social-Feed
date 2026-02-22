import { apiClient } from '../api';

export interface Comment {
    id: string;
    post_id: string;
    user_id: string;
    content: string;
    created_at: string;
    updated_at: string;
}

export interface CommentsResponse {
    items: Comment[];
    meta: {
        total: number;
        page: number;
        limit: number;
        total_pages: number;
    };
}

export async function getComments(postId: string, page: number = 1): Promise<CommentsResponse> {
    const response = await apiClient.get<{ data: CommentsResponse }>('/api/v1/news-feed/posts/comment', {
        params: { post_id: postId, page, limit: 20, sort: '-created_at' },
    });
    return response.data.data;
}

export async function createComment(postId: string, content: string): Promise<Comment> {
    const response = await apiClient.post<{ data: Comment }>('/api/v1/news-feed/posts/comment', {
        post_id: postId,
        content,
    });
    return response.data.data;
}

export async function updateComment(id: string, content: string): Promise<Comment> {
    const response = await apiClient.put<{ data: Comment }>('/api/v1/news-feed/posts/comment', {
        id,
        content,
    });
    return response.data.data;
}

export async function deleteComment(id: string): Promise<void> {
    await apiClient.delete(`/api/v1/news-feed/posts/comment/${id}`);
}

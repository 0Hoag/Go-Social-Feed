import { apiClient } from '../api';

export interface Post {
    id: string;
    title: string;
    title_en: string;
    content: string;
    full_content: string;
    full_content_en: string;
    file_ids: string[];
    tagged_target: string[];
    pin: boolean;
    author_id: string;
    source_url: string;
    permission: string;
    created_at: string;
    updated_at: string;
}

export interface CreatePostInput {
    content: string;
    permission: 'public' | 'private';
    pin?: boolean;
    file_ids?: string[];
    tagged_target?: string[];
}

export interface Reaction {
    id: string;
    post_id: string;
    user_id: string;
    type: string;
    created_at: string;
}

export interface PostsResponse {
    items: Post[];
    meta: {
        total: number;
        page: number;
        limit: number;
        total_pages: number;
    };
}

export async function getPosts(params?: { page?: number; limit?: number }): Promise<PostsResponse> {
    const response = await apiClient.get<{ data: PostsResponse }>('/api/v1/news-feed/posts', {
        params: {
            page: params?.page || 1,
            limit: params?.limit || 20,
            sort: '-created_at',
        },
    });
    return response.data.data;
}

export async function getPostById(id: string): Promise<Post> {
    const response = await apiClient.get<{ data: Post }>(`/api/v1/news-feed/posts/${id}`);
    return response.data.data;
}

export async function createPost(input: CreatePostInput): Promise<Post> {
    const response = await apiClient.post<{ data: Post }>('/api/v1/news-feed/posts', input);
    return response.data.data;
}

export async function updatePost(id: string, input: Partial<CreatePostInput>): Promise<void> {
    await apiClient.put(`/api/v1/news-feed/posts`, { id, ...input });
}

export async function deletePost(id: string): Promise<void> {
    await apiClient.delete(`/api/v1/news-feed/posts/${id}`);
}

// Reactions
export async function createReaction(postId: string, type: string = 'like'): Promise<Reaction> {
    const response = await apiClient.post<{ data: Reaction }>('/api/v1/news-feed/posts/reaction', {
        post_id: postId,
        type,
    });
    return response.data.data;
}

export async function getReactions(postId: string): Promise<{ items: Reaction[]; meta: any }> {
    const response = await apiClient.get<{ data: { items: Reaction[]; meta: any } }>('/api/v1/news-feed/posts/reaction', {
        params: { post_id: postId },
    });
    return response.data.data;
}

export async function deleteReaction(id: string): Promise<void> {
    await apiClient.delete(`/api/v1/news-feed/posts/reaction/${id}`);
}

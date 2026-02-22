import { apiClient } from '../api';

export interface Follow {
    id: string;
    follower_id: string;
    target_id: string;
    created_at: string;
}

export interface FollowsResponse {
    items: Follow[];
    meta: {
        total: number;
        page: number;
        limit: number;
        total_pages: number;
    };
}

export async function follow(targetId: string): Promise<Follow> {
    const response = await apiClient.post<{ data: Follow }>('/api/v1/news-feed/follows', {
        target_id: targetId,
    });
    return response.data.data;
}

export async function unfollow(id: string): Promise<void> {
    await apiClient.delete(`/api/v1/news-feed/follows/${id}`);
}

export async function getFollows(params?: {
    follower_id?: string;
    target_id?: string;
    page?: number;
    limit?: number;
}): Promise<FollowsResponse> {
    const response = await apiClient.get<{ data: FollowsResponse }>('/api/v1/news-feed/follows', {
        params: {
            ...params,
            page: params?.page || 1,
            limit: params?.limit || 50,
        },
    });
    return response.data.data;
}

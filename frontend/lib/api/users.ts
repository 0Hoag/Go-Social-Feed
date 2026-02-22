import { apiClient } from '../api';

export interface User {
    id: string;
    username: string;
    phone: string;
    avatar_url: string;
    created_at: string;
    updated_at: string;
}

export interface UpdateUserInput {
    username?: string;
    avatar_url?: string;
}

export async function getMyInfo(): Promise<User> {
    const response = await apiClient.get<{ data: User }>('/api/v1/news-feed/users/myinfo');
    return response.data.data;
}

export async function getUser(id: string): Promise<User> {
    const response = await apiClient.get<{ data: User }>(`/api/v1/news-feed/users/${id}`);
    return response.data.data;
}

export async function getUsers(params?: { username?: string; page?: number; limit?: number }): Promise<{ items: User[]; meta: any }> {
    const response = await apiClient.get<{ data: { items: User[]; meta: any } }>('/api/v1/news-feed/users', {
        params: {
            username: params?.username,
            page: params?.page || 1,
            limit: params?.limit || 20,
        },
    });
    return response.data.data;
}

export async function updateUser(id: string, input: UpdateUserInput): Promise<void> {
    await apiClient.put(`/api/v1/news-feed/users`, { id, ...input });
}

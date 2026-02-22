import { apiClient } from '../api';

export interface LoginResponse {
    token: string;
}

export interface RegisterInput {
    username: string;
    phone: string;
    password: string;
    birthday: string; // ISO date string
}

export async function login(phone: string, password: string): Promise<LoginResponse> {
    const response = await apiClient.post<{ data: LoginResponse }>('/api/v1/news-feed/auth/login', {
        phone,
        password,
    });
    return response.data.data;
}

export async function register(input: RegisterInput): Promise<void> {
    await apiClient.post('/api/v1/news-feed/users', {
        username: input.username,
        phone: input.phone,
        password: input.password,
        birthday: input.birthday,
    });
}

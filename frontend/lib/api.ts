import axios from "axios";
import { Post, PaginationParams, PostsResponse } from "./types";

const API_BASE_URL = process.env.NEXT_PUBLIC_API_URL || "http://localhost:8080";

const apiClient = axios.create({
    baseURL: API_BASE_URL,
    headers: {
        "Content-Type": "application/json",
    },
});

// Add response interceptor for error handling
apiClient.interceptors.response.use(
    (response) => response,
    (error) => {
        console.error("API Error:", error.response?.data || error.message);
        return Promise.reject(error);
    }
);

export async function getPosts(params?: PaginationParams): Promise<PostsResponse> {
    try {
        const response = await apiClient.get<{ data: { items: Post[], meta: any } }>("/api/v1/news-feed/posts", {
            params: {
                page: params?.page || 1,
                limit: params?.limit || 30,
                sort: params?.sort || "-created_at",
            },
        });
        // Backend returns {data: {items: [], meta: {}}}
        return {
            posts: response.data.data.items,
            pagination: response.data.data.meta
        };
    } catch (error) {
        console.error("Failed to fetch posts:", error);
        throw error;
    }
}

export async function getPostById(id: string): Promise<Post> {
    try {
        const response = await apiClient.get<{ data: Post }>(`/api/v1/news-feed/posts/${id}`);
        return response.data.data;
    } catch (error) {
        console.error(`Failed to fetch post ${id}:`, error);
        throw error;
    }
}

export { apiClient };

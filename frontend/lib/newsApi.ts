import axios from 'axios';

const API_BASE_URL = process.env.NEXT_PUBLIC_API_URL || 'http://localhost:8080/api/v1';

export const fetchTrendingNews = async () => {
    try {
        const response = await axios.get(`${API_BASE_URL}/news-feed/posts`, {
            params: {
                limit: 5,
                status: 'published' // Assuming standard filter, can adjust based on backend
            }
        });
        return response.data.data; // Assuming standard response structure
    } catch (error) {
        console.error("Failed to fetch news:", error);
        return [];
    }
};

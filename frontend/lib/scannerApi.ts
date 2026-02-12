import axios from 'axios';

const API_BASE_URL = process.env.NEXT_PUBLIC_API_URL || 'http://localhost:8080/api/v1';

export interface ScanResult {
    network: string;
    name: string;
    address: string;
    trust_score: number;
    issues: Issue[];
    safe_features: string[];
}

export interface Issue {
    name: string;
    description: string;
    impact: number;
    type: string;
}

export const scanToken = async (query: string, language: string = 'en'): Promise<ScanResult> => {
    try {
        const response = await axios.get(`${API_BASE_URL}/news-feed/scanner`, {
            params: {
                token: query,
                lang: language  // Send language to Backend
            }
        });
        return response.data.data;
    } catch (error) {
        console.error("Scan error:", error);
        throw error;
    }
};

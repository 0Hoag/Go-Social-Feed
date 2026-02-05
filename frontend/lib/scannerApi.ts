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
    Name: string;
    Description: string;
    Impact: number;
    Type: string;
}

export const scanToken = async (query: string): Promise<ScanResult> => {
    try {
        const response = await axios.get(`${API_BASE_URL}/scan`, {
            params: { q: query }
        });
        return response.data;
    } catch (error) {
        console.error("Scan error:", error);
        throw error;
    }
};

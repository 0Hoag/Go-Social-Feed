import { NextRequest, NextResponse } from 'next/server';
import { GoogleGenerativeAI } from '@google/generative-ai';

const genAI = new GoogleGenerativeAI(process.env.GEMINI_API_KEY || '');

export async function POST(request: NextRequest) {
    try {
        const { message, coinSymbol, coinName, currentPrice, priceChange } = await request.json();

        const prompt = `
        Bạn là chuyên gia phân tích thị trường crypto (CryptoCheck AI).
        Hiện tại người dùng đang hỏi về đồng: ${coinName} (${coinSymbol}).
        Giá hiện tại: $${currentPrice} (Biến động 24h: ${priceChange}%).

        Câu hỏi của người dùng: "${message}"

        Hãy trả lời ngắn gọn, súc tích (dưới 200 từ), tập trung vào dữ liệu và xu hướng.
        **YÊU CẦU ĐỊNH DẠNG (BẮT BUỘC):**
        - Dùng tiêu đề \`###\` cho các ý lớn (Ví dụ: ### Tình hình chung). **KHÔNG ĐÁNH SỐ**.
        - Dùng \`>\` (blockquote) cho các câu tóm tắt hoặc điểm nhấn quan trọng.
        - **ƯU TIÊN dùng đoạn văn (paragraph)** để diễn giải chi tiết. 
        - Chỉ dùng danh sách \`-\` khi liệt kê các ý ngắn gọn.
        - **In đậm** các từ khóa quan trọng.
        - Tách đoạn rõ ràng để dễ đọc.

        Cuối cùng, hãy đề xuất 3 câu hỏi ngắn gọn (dưới 10 từ) mà người dùng có thể muốn hỏi tiếp theo liên quan đến câu trả lời của bạn.
        Định dạng bắt buộc:
        ---QUESTIONS---
        Câu hỏi 1
        Câu hỏi 2
        Câu hỏi 3

        Nếu câu hỏi không liên quan đến crypto, hãy lái về chủ đề thị trường một cách khéo léo.
        Phong cách: Chuyên nghiệp, khách quan, hữu ích.
        `;

        const model = genAI.getGenerativeModel({ model: 'gemini-2.5-flash' });
        const result = await model.generateContent(prompt);
        const response = await result.response;
        const text = response.text();

        return NextResponse.json({ reply: text });

    } catch (error) {
        console.error('Chat API Error:', error);
        return NextResponse.json({
            reply: `Xin lỗi, có lỗi xảy ra: ${error instanceof Error ? error.message : String(error)}`
        });
    }
}

import { NextRequest, NextResponse } from 'next/server';
import { GoogleGenerativeAI } from '@google/generative-ai';

const genAI = new GoogleGenerativeAI(process.env.GEMINI_API_KEY || '');

export async function POST(request: NextRequest) {
    try {
        const { message, coinSymbol, coinName, currentPrice, priceChange, language } = await request.json();

        const isVi = language === 'vi';

        const prompt = isVi ? `
        **QUY TẮC QUAN TRỌNG NHẤT (MUST FOLLOW):**
        - **NẾU NGƯỜI DÙNG HỎI TIẾNG ANH -> BẮT BUỘC TRẢ LỜI TIẾNG ANH.**
        - **NẾU NGƯỜI DÙNG HỎI TIẾNG VIỆT -> TRẢ LỜI TIẾNG VIỆT.**
        - Không quan trọng ngôn ngữ của prompt này, hãy ưu tiên ngôn ngữ của câu hỏi.

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

        **QUAN TRỌNG:**
        Với mọi câu trả lời liên quan đến xu hướng giá hoặc lời khuyên đầu tư, BẮT BUỘC phải kết thúc bằng dòng disclaimer sau (in nghiêng):
        *> Lưu ý: Đây là nhận định dựa trên dữ liệu tham khảo, có thể đúng hoặc sai. Quyết định đầu tư thuộc về bạn, vui lòng cân nhắc kỹ lưỡng trước khi hành động.*
        ` : `
        **MOST IMPORTANT RULE (MUST FOLLOW):**
        - **IF USER ASKS IN VIETNAMESE -> YOU MUST ANSWER IN VIETNAMESE.**
        - **IF USER ASKS IN ENGLISH -> ANSWER IN ENGLISH.**
        - Prioritize the language of the user's question above all else.

        You are a crypto market analysis expert (CryptoCheck AI).
        The user is asking about: ${coinName} (${coinSymbol}).
        Current Price: $${currentPrice} (24h Change: ${priceChange}%).

        User Question: "${message}"

        Please provide a concise, data-driven answer (under 200 words).
        **FORMAT REQUIREMENTS (MANDATORY):**
        - Use \`###\` for main headers (e.g., ### Market Overview). **DO NOT NUMBER THEM**.
        - Use \`>\` (blockquote) for summaries or key points.
        - **PREFER paragraphs** for detailed explanation.
        - Use \`-\` lists only for short items.
        - **Bold** key terms.
        - Separate paragraphs clearly.

        Finally, suggest 3 short follow-up questions (under 10 words).
        Mandatory format:
        ---QUESTIONS---
        Question 1
        Question 2
        Question 3

        If the question is unrelated to crypto, politely steer it back to market topics.
        Style: Professional, objective, helpful.

        **IMPORTANT:**
        For any answer regarding price trends or investment expectations, YOU MUST end with this disclaimer (italicized):
        *> Note: This analysis is based on reference data and may not be accurate. Investment decisions are yours alone; please consider carefully before acting.*
        `;

        const model = genAI.getGenerativeModel({ model: 'gemini-2.5-flash' });
        const result = await model.generateContentStream(prompt);

        const stream = new ReadableStream({
            async start(controller) {
                const encoder = new TextEncoder();
                try {
                    for await (const chunk of result.stream) {
                        const text = chunk.text();
                        if (text) {
                            controller.enqueue(encoder.encode(text));
                        }
                    }
                    controller.close();
                } catch (error) {
                    console.error('Stream processing error:', error);
                    controller.error(error);
                }
            }
        });

        return new Response(stream, {
            headers: { 'Content-Type': 'text/plain; charset=utf-8' }
        });

    } catch (error) {
        console.error('Chat API Error:', error);
        return NextResponse.json({
            error: `Xin lỗi, có lỗi xảy ra: ${error instanceof Error ? error.message : String(error)}`
        }, { status: 500 });
    }
}

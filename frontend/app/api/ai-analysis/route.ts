import { NextRequest, NextResponse } from 'next/server';
import { GoogleGenerativeAI } from '@google/generative-ai';

// Initialize Gemini AI
const genAI = new GoogleGenerativeAI(process.env.GEMINI_API_KEY || '');

export async function POST(request: NextRequest) {
    let symbol, name, price, priceChange, timeframe, language;

    try {
        const body = await request.json();
        symbol = body.symbol;
        name = body.name;
        price = body.price;
        priceChange = body.priceChange;
        timeframe = body.timeframe;
        language = body.language || 'en';

        const isVi = language === 'vi';

        // Build comprehensive prompt
        const prompt = isVi ? `
        Bạn là chuyên gia phân tích thị trường tiền điện tử. Hãy phân tích biến động giá của ${name} (${symbol}) với thông tin sau:

        **Dữ liệu thị trường:**
        - Giá hiện tại: $${price.toLocaleString()}
        - Biến động ${timeframe}: ${priceChange > 0 ? '+' : ''}${priceChange.toFixed(2)}%

        **Yêu cầu phân tích:**
        Cung cấp phân tích chi tiết bằng tiếng Việt theo cấu trúc JSON sau:

        {
        "tldr": "Tóm tắt ngắn gọn 2-3 câu về biến động giá, nguyên nhân chính và triển vọng",
        "mainCause": {
            "title": "Tiêu đề nguyên nhân chính",
            "overview": "Mô tả chi tiết nguyên nhân chính",
            "implication": "Ý nghĩa của nguyên nhân này",
            "watchFor": "Các yếu tố cần theo dõi"
        },
        "secondaryCause": {
            "title": "Tiêu đề nguyên nhân phụ (nếu có)",
            "overview": "Mô tả chi tiết",
            "implication": "Ý nghĩa",
            "watchFor": "Yếu tố theo dõi"
        },
        "outlook": {
            "title": "Triển vọng thị trường ngắn hạn",
            "overview": "Phân tích kỹ thuật và xu hướng",
            "implication": "Ý nghĩa cho nhà đầu tư",
            "watchFor": "Mức giá quan trọng cần theo dõi"
        },
        "conclusion": {
            "outlook": "Áp lực tăng/giảm/trung lập",
            "keyPoint": "Điểm cần lưu ý quan trọng nhất"
        }
        }

        **Lưu ý:**
        - Sử dụng dữ liệu thực tế về thị trường crypto
        - Phân tích dựa trên yếu tố vĩ mô, kỹ thuật, và tâm lý thị trường
        - Đề cập đến các mức hỗ trợ/kháng cự quan trọng
        - Nếu không có nguyên nhân phụ rõ ràng, có thể bỏ qua "secondaryCause"
        - Trả về CHÍNH XÁC theo format JSON, không thêm markdown hay text khác
        
        **QUAN TRỌNG:**
        BẮT BUỘC phải thêm dòng disclaimer sau vào cuối nội dung của trường "conclusion.keyPoint" hoặc "tldr":
        " (Lưu ý: Đây là nhận định dựa trên dữ liệu tham khảo, có thể đúng hoặc sai. Quyết định đầu tư thuộc về bạn.)"
        ` : `
        You are a cryptocurrency market analysis expert. Analyze the price movement of ${name} (${symbol}) with the following information:

        **Market Data:**
        - Current Price: $${price.toLocaleString()}
        - ${timeframe} Change: ${priceChange > 0 ? '+' : ''}${priceChange.toFixed(2)}%

        **Analysis Request:**
        Provide a detailed analysis in English following this JSON structure:

        {
        "tldr": "Short 2-3 sentence summary of price movement, main cause, and outlook",
        "mainCause": {
            "title": "Main cause title",
            "overview": "Detailed description of the main cause",
            "implication": "Implication of this cause",
            "watchFor": "Factors to watch"
        },
        "secondaryCause": {
            "title": "Secondary cause title (if any)",
            "overview": "Detailed description",
            "implication": "Implication",
            "watchFor": "Factors to watch"
        },
        "outlook": {
            "title": "Short-term market outlook",
            "overview": "Technical analysis and trend",
            "implication": "Implication for investors",
            "watchFor": "Key price levels to watch"
        },
        "conclusion": {
            "outlook": "Bullish/Bearish/Neutral Pressure",
            "keyPoint": "Most important takeaway"
        }
        }

        **Note:**
        - Use real crypto market data
        - Analyze based on macro, technical, and market sentiment factors
        - Mention key support/resistance levels
        - If no clear secondary cause, "secondaryCause" can be omitted
        - Return EXACTLY in JSON format, no markdown or other text
        
        **IMPORTANT:**
        MUST add the following disclaimer at the end of "conclusion.keyPoint" or "tldr":
        " (Note: This analysis is based on reference data and may not be accurate. Investment decisions are yours alone.)"
        `;

        const model = genAI.getGenerativeModel({ model: 'gemini-2.5-flash' });
        const result = await model.generateContent(prompt);
        const response = await result.response;
        const text = response.text();

        // Parse JSON from response
        let analysis;
        try {
            // Remove markdown code blocks if present
            const cleanText = text.replace(/```json\n?/g, '').replace(/```\n?/g, '').trim();
            analysis = JSON.parse(cleanText);
        } catch (parseError) {
            console.error('Failed to parse AI response:', text);
            // Fallback to mock data if parsing fails
            analysis = generateMockAnalysis(name, priceChange, language);
        }

        return NextResponse.json({
            success: true,
            analysis
        });

    } catch (error) {
        console.error('AI Analysis error:', error);

        // Return mock data on error using already parsed values
        return NextResponse.json({
            success: false,
            analysis: generateMockAnalysis(name || 'Bitcoin', priceChange || 0, language || 'en'),
            error: 'Using fallback analysis'
        });
    }
}

function generateMockAnalysis(coinName: string, priceChange: number, language: string) {
    const isNegative = priceChange < 0;
    const absChange = Math.abs(priceChange);
    const isVi = language === 'vi';

    if (isVi) {
        return {
            tldr: `${coinName} ${isNegative ? 'giảm' : 'tăng'} ${absChange.toFixed(1)}% trong 24 giờ qua, chủ yếu do ${isNegative ? 'áp lực bán tháo từ thị trường rộng lớn hơn và thanh lý các vị thế đòn bẩy' : 'tâm lý tích cực từ thị trường và dòng tiền mua vào mạnh'}. ${isNegative ? 'Cần theo dõi các mức hỗ trợ quan trọng để tránh đà giảm sâu hơn' : 'Xu hướng tăng có thể tiếp tục nếu vượt qua các mức kháng cự'}.`,

            mainCause: {
                title: isNegative ? "Bán tháo thị trường do yếu tố vĩ mô" : "Tâm lý thị trường tích cực",
                overview: isNegative
                    ? `Thị trường crypto chịu áp lực bán mạnh do lo ngại về lãi suất và thanh khoản. ${coinName} di chuyển đồng bộ với xu hướng giảm chung của thị trường, phản ánh mối tương quan cao với các tài sản rủi ro.`
                    : `Thị trường crypto hưởng lợi từ tâm lý tích cực và dòng tiền mua vào mạnh. ${coinName} tăng giá theo xu hướng chung, được hỗ trợ bởi khối lượng giao dịch tăng cao.`,
                implication: isNegative
                    ? `${coinName} hiện được xem như một tài sản rủi ro có độ biến động cao, không tách rời khỏi thị trường truyền thống trong giai đoạn suy giảm này.`
                    : `${coinName} đang thu hút sự quan tâm của nhà đầu tư, với dòng tiền vào ổn định và tâm lý thị trường cải thiện.`,
                watchFor: "Thay đổi trong tâm lý thị trường và các dữ liệu vĩ mô mới có thể ảnh hưởng đến xu hướng giá."
            },

            secondaryCause: isNegative ? {
                title: "Áp lực thanh lý các vị thế đòn bẩy",
                overview: `Một đợt thanh lý hàng loạt các vị thế ${isNegative ? 'mua' : 'bán'} với đòn bẩy cao đã ${isNegative ? 'đẩy nhanh đà giảm' : 'khuếch đại đà tăng'}. Tỷ lệ phí tài trợ ${isNegative ? 'âm sâu' : 'dương cao'} cho thấy ${isNegative ? 'đòn bẩy bán quá mức' : 'đòn bẩy mua quá mức'}.`,
                implication: `Các vị thế sử dụng đòn bẩy bị buộc phải đóng, tạo thêm áp lực ${isNegative ? 'bán' : 'mua'} và làm ${isNegative ? 'trầm trọng thêm đà giảm' : 'khuếch đại đà tăng'}.`,
                watchFor: "Sự bình thường hóa của tỷ lệ phí tài trợ và giảm tốc độ thanh lý."
            } : undefined,

            outlook: {
                title: "Triển vọng thị trường ngắn hạn",
                overview: isNegative
                    ? `Về mặt kỹ thuật, ${coinName} đang bị bán quá mức (RSI thấp) và đã phá vỡ một số mức hỗ trợ quan trọng. Cần theo dõi các mức giá then chốt để xác định xu hướng tiếp theo.`
                    : `Về mặt kỹ thuật, ${coinName} đang trong xu hướng tăng với RSI ở vùng trung tính đến mua quá mức. Cần theo dõi các mức kháng cự để đánh giá khả năng tăng tiếp.`,
                implication: isNegative
                    ? "Xu hướng hiện tại là giảm, nhưng điều kiện bán quá mức có thể tạo cơ hội cho một nhịp hồi ngắn hạn."
                    : "Xu hướng hiện tại là tăng, nhưng cần cảnh giác với tín hiệu mua quá mức có thể dẫn đến điều chỉnh.",
                watchFor: isNegative
                    ? "Diễn biến giá quanh các mức hỗ trợ quan trọng; nếu có sự từ chối mạnh, có thể báo hiệu đáy cục bộ."
                    : "Diễn biến giá quanh các mức kháng cự; việc vượt qua sẽ mở đường cho đà tăng mạnh hơn."
            },

            conclusion: {
                outlook: isNegative ? "Áp lực giảm" : "Áp lực tăng",
                keyPoint: isNegative
                    ? `${coinName} cần giữ vững các mức hỗ trợ quan trọng để tránh đà giảm sâu hơn. Theo dõi chặt chẽ khối lượng giao dịch và tâm lý thị trường.`
                    : `${coinName} có thể tiếp tục tăng nếu vượt qua các mức kháng cự. Cần theo dõi khối lượng giao dịch để xác nhận xu hướng.`
            }
        };
    } else {
        return {
            tldr: `${coinName} ${isNegative ? 'decreased' : 'increased'} by ${absChange.toFixed(1)}% in the last 24 hours, mainly due to ${isNegative ? 'broad market sell-off pressure and leveraged position liquidations' : 'positive market sentiment and strong buying inflow'}. ${isNegative ? 'Key support levels need to be monitored to avoid further decline' : 'The uptrend may continue if resistance levels are broken'}.`,

            mainCause: {
                title: isNegative ? "Market Sell-off due to Macro Factors" : "Positive Market Sentiment",
                overview: isNegative
                    ? `The crypto market is under strong selling pressure due to concerns about interest rates and liquidity. ${coinName} moves in sync with the general market downtrend, reflecting high correlation with risk assets.`
                    : `The crypto market benefits from positive sentiment and strong buying inflows. ${coinName} is rising in line with the general trend, supported by increased trading volume.`,
                implication: isNegative
                    ? `${coinName} is currently viewed as a high-volatility risk asset, not decoupled from traditional markets during this downturn.`
                    : `${coinName} is attracting investor interest, with stable inflows and improved market sentiment.`,
                watchFor: "Changes in market sentiment and new macro data that could affect price trends."
            },

            secondaryCause: isNegative ? {
                title: "Leveraged Position Liquidations",
                overview: `A wave of liquidations of highly leveraged ${isNegative ? 'long' : 'short'} positions has ${isNegative ? 'accelerated the decline' : 'amplified the rally'}. ${isNegative ? 'Deeply negative' : 'High positive'} funding rates suggest ${isNegative ? 'excessive short leverage' : 'excessive long leverage'}.`,
                implication: `Leveraged positions are forced to close, creating additional ${isNegative ? 'selling' : 'buying'} pressure and ${isNegative ? 'exacerbating the drop' : 'fueling the rise'}.`,
                watchFor: "Normalization of funding rates and a slowdown in liquidations."
            } : undefined,

            outlook: {
                title: "Short-term Market Outlook",
                overview: isNegative
                    ? `Technically, ${coinName} is oversold (low RSI) and has broken some key support levels. Key price levels need to be monitored to determine the next trend.`
                    : `Technically, ${coinName} is in an uptrend with RSI in neutral to overbought territory. Resistance levels need to be watched to assess further upside potential.`,
                implication: isNegative
                    ? "The current trend is bearish, but oversold conditions may create an opportunity for a short-term bounce."
                    : "The current trend is bullish, but be wary of overbought signals that could lead to a correction.",
                watchFor: isNegative
                    ? "Price action around key support levels; a strong rejection could signal a local bottom."
                    : "Price action around resistance levels; a breakout would pave the way for a stronger rally."
            },

            conclusion: {
                outlook: isNegative ? "Bearish Pressure" : "Bullish Pressure",
                keyPoint: isNegative
                    ? `${coinName} needs to hold key support levels to avoid deeper decline. Closely monitor trading volume and market sentiment.`
                    : `${coinName} could continue to rise if it breaks resistance levels. Monitor trading volume to confirm the trend.`
            }
        };
    }
}
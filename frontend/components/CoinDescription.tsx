import { useState } from 'react';
import { useLanguage } from '@/context/LanguageContext';
import { ChevronDown, ChevronUp } from 'lucide-react';

interface CoinDescriptionProps {
    description: string;
    coinName: string;
}

export default function CoinDescription({ description, coinName }: CoinDescriptionProps) {
    const { t, language } = useLanguage();
    const [isExpanded, setIsExpanded] = useState(false);

    // Threshold for truncating text (characters)
    const TRUNCATE_LENGTH = 300;
    const shouldTruncate = description.length > TRUNCATE_LENGTH;

    const displayedText = isExpanded ? description : description.slice(0, TRUNCATE_LENGTH) + (shouldTruncate ? '...' : '');

    return (
        <div className="bg-[#111] border border-white/5 rounded-2xl p-6 relative overflow-hidden group hover:border-white/10 transition-all duration-300">
            <h2 className="text-xl font-bold text-white mb-4 flex items-center gap-2">
                <span className="bg-clip-text text-transparent bg-gradient-to-r from-blue-400 to-purple-400">
                    {t.analysis_page.about} {coinName}
                </span>
            </h2>

            <div className={`relative transition-all duration-500 ${isExpanded ? 'max-h-[1000px]' : 'max-h-[200px]'}`}>
                {displayedText.split('\n').map((paragraph, index) => (
                    paragraph.trim() && (
                        <p key={index} className="text-gray-400 text-sm leading-relaxed text-justify mb-4 last:mb-0">
                            {paragraph}
                        </p>
                    )
                ))}

                {/* Gradient overlay when collapsed */}
                {!isExpanded && shouldTruncate && (
                    <div className="absolute bottom-0 left-0 w-full h-16 bg-gradient-to-t from-[#111] to-transparent pointer-events-none" />
                )}
            </div>

            {shouldTruncate && (
                <button
                    onClick={() => setIsExpanded(!isExpanded)}
                    className="mt-4 flex items-center gap-2 text-xs font-bold text-blue-400 hover:text-blue-300 transition-colors uppercase tracking-wider bg-blue-500/10 px-4 py-2 rounded-lg hover:bg-blue-500/20"
                >
                    {isExpanded ? (
                        <>
                            {language === 'vi' ? 'Thu gọn' : 'Show Less'} <ChevronUp className="w-3 h-3" />
                        </>
                    ) : (
                        <>
                            {language === 'vi' ? 'Xem thêm' : 'Read More'} <ChevronDown className="w-3 h-3" />
                        </>
                    )}
                </button>
            )}
        </div>
    );
}

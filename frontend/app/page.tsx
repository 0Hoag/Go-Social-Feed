import HeroSection from "@/components/HeroSection";
import TokenSearch from "@/components/TokenSearch";
import FeaturesGrid from "@/components/FeaturesGrid";
import HowItWorks from "@/components/HowItWorks";
import SecurityIssues from "@/components/SecurityIssues";
import CTASection from "@/components/CTASection";

export default function Home() {
  return (
    <main className="min-h-screen bg-[#050505] text-gray-200 font-sans selection:bg-cyan-500/20 selection:text-cyan-200">
      {/* Hero Section */}
      <HeroSection />

      {/* Token Search & Trending */}
      <TokenSearch />

      {/* Features Grid */}
      <FeaturesGrid />

      {/* How It Works */}
      <HowItWorks />

      {/* Security Issues */}
      <SecurityIssues />

      {/* Final CTA Section */}
      <CTASection />
    </main>
  );
}

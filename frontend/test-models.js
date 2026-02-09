const { GoogleGenerativeAI } = require("@google/generative-ai");

async function listModels() {
    const apiKey = process.argv[2];
    if (!apiKey) {
        console.error("Please provide API KEY as argument");
        process.exit(1);
    }

    const genAI = new GoogleGenerativeAI(apiKey);
    try {
        // Note: List models is on the specific client or manager usually? 
        // Actually SDK might not expose listModels directly on genAI instance easily in older versions?
        // Let's check imports.
        // In newer SDK: genAI.getGenerativeModel...
        // There isn't a direct listModels helper in the main class in some versions.
        // We might need to use fetch directly to be sure.

        // Using fetch to call the API directly for listing models
        const url = `https://generativelanguage.googleapis.com/v1beta/models?key=${apiKey}`;
        const response = await fetch(url);
        const data = await response.json();

        if (data.models) {
            console.log("Available Models:");
            data.models.forEach(m => {
                if (m.supportedGenerationMethods && m.supportedGenerationMethods.includes("generateContent")) {
                    console.log(`- ${m.name}`);
                }
            });
        } else {
            console.log("No models found or error:", data);
        }

    } catch (error) {
        console.error("Error:", error);
    }
}

listModels();

import React from "react";
import PromptInput from "@/components/PromptInput";
import ImageGallery from "@/components/ImageGallery";

function App() {
    const [images, setImages] = React.useState<string[]>([]);
    const [loading, setLoading] = React.useState(false);

    const handleGenerate = async (
        prompt: string,
        model: string,
        count: number
    ) => {
        setLoading(true);
        try {
            // TODO: Call backend API to generate images
            console.log("Generating images:", { prompt, model, count });
            // setImages([...]) - will be implemented with API integration
        } catch (error) {
            console.error("Failed to generate images:", error);
        } finally {
            setLoading(false);
        }
    };

    return (
        <div className="min-h-screen bg-gradient-to-br from-primary to-secondary">
            <header className="text-center py-12 text-white">
                <h1 className="text-5xl font-bold mb-2 animate-float drop-shadow-lg">
                    🎲 Image Gacha Generator
                </h1>
                <p className="text-xl opacity-90">
                    Generate AI images with a gacha twist - every spin is
                    different!
                </p>
            </header>
            <main className="flex-1 max-w-4xl w-full mx-auto px-4 py-8 space-y-8">
                <PromptInput onGenerate={handleGenerate} loading={loading} />
                <ImageGallery images={images} loading={loading} />
            </main>
        </div>
    );
}

export default App;

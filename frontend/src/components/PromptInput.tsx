import React from "react";

interface PromptInputProps {
    onGenerate: (prompt: string, model: string, count: number) => void;
    loading: boolean;
}

const PromptInput: React.FC<PromptInputProps> = ({ onGenerate, loading }) => {
    const [prompt, setPrompt] = React.useState("");
    const [model, setModel] = React.useState("stable-diffusion-v2.1");
    const [count, setCount] = React.useState(4);
    const [showAdvanced, setShowAdvanced] = React.useState(false);
    const [steps, setSteps] = React.useState(30);
    const [cfgScale, setCfgScale] = React.useState(7.5);
    const [negPrompt, setNegPrompt] = React.useState("");

    const handleSubmit = (e: React.FormEvent) => {
        e.preventDefault();
        if (prompt.trim()) {
            onGenerate(prompt, model, count);
        }
    };

    return (
        <form
            className="bg-white rounded-lg p-8 shadow-lg"
            onSubmit={handleSubmit}
        >
            <div className="mb-6">
                <label
                    htmlFor="prompt"
                    className="block font-semibold text-gray-800 mb-2"
                >
                    Your Prompt
                </label>
                <textarea
                    id="prompt"
                    value={prompt}
                    onChange={(e) => setPrompt(e.target.value)}
                    placeholder="Describe the image you want to generate..."
                    disabled={loading}
                    rows={3}
                    className="w-full px-3 py-2 border-2 border-gray-300 rounded-lg font-normal focus:outline-none focus:border-primary focus:ring-2 focus:ring-primary/20 disabled:opacity-50 resize-vertical"
                />
            </div>

            <div className="grid grid-cols-2 gap-4 mb-6">
                <div>
                    <label
                        htmlFor="model"
                        className="block font-semibold text-gray-800 mb-2"
                    >
                        Model
                    </label>
                    <select
                        id="model"
                        value={model}
                        onChange={(e) => setModel(e.target.value)}
                        disabled={loading}
                        className="w-full px-3 py-2 border-2 border-gray-300 rounded-lg font-normal focus:outline-none focus:border-primary focus:ring-2 focus:ring-primary/20 disabled:opacity-50"
                    >
                        <option value="stable-diffusion-v1.5">
                            Stable Diffusion v1.5
                        </option>
                        <option value="stable-diffusion-v2.1">
                            Stable Diffusion v2.1
                        </option>
                        <option value="dall-e-3">DALL-E 3</option>
                    </select>
                </div>

                <div>
                    <label
                        htmlFor="count"
                        className="block font-semibold text-gray-800 mb-2"
                    >
                        Number of Images
                    </label>
                    <div className="flex gap-2">
                        <button
                            type="button"
                            onClick={() => setCount(Math.max(1, count - 1))}
                            disabled={loading || count <= 1}
                            className="w-10 h-10 border-2 border-primary bg-white text-primary rounded-lg font-bold hover:bg-primary hover:text-white disabled:opacity-50 disabled:cursor-not-allowed transition-all"
                        >
                            −
                        </button>
                        <input
                            type="number"
                            id="count"
                            value={count}
                            onChange={(e) =>
                                setCount(
                                    Math.min(
                                        10,
                                        Math.max(
                                            1,
                                            parseInt(e.target.value) || 1
                                        )
                                    )
                                )
                            }
                            min="1"
                            max="10"
                            disabled={loading}
                            className="flex-1 px-3 py-2 border-2 border-gray-300 rounded-lg text-center focus:outline-none focus:border-primary focus:ring-2 focus:ring-primary/20 disabled:opacity-50"
                        />
                        <button
                            type="button"
                            onClick={() => setCount(Math.min(10, count + 1))}
                            disabled={loading || count >= 10}
                            className="w-10 h-10 border-2 border-primary bg-white text-primary rounded-lg font-bold hover:bg-primary hover:text-white disabled:opacity-50 disabled:cursor-not-allowed transition-all"
                        >
                            +
                        </button>
                    </div>
                </div>
            </div>

            <button
                type="button"
                className="text-primary font-semibold mb-4 hover:text-secondary transition-colors"
                onClick={() => setShowAdvanced(!showAdvanced)}
            >
                {showAdvanced ? "▼" : "▶"} Advanced Options
            </button>

            {showAdvanced && (
                <div className="bg-gray-50 rounded-lg p-6 mb-6 border-l-4 border-primary">
                    <div className="mb-6">
                        <label
                            htmlFor="negPrompt"
                            className="block font-semibold text-gray-800 mb-2"
                        >
                            Negative Prompt
                        </label>
                        <textarea
                            id="negPrompt"
                            value={negPrompt}
                            onChange={(e) => setNegPrompt(e.target.value)}
                            placeholder="What to avoid in the image..."
                            disabled={loading}
                            rows={2}
                            className="w-full px-3 py-2 border-2 border-gray-300 rounded-lg font-normal focus:outline-none focus:border-primary focus:ring-2 focus:ring-primary/20 disabled:opacity-50 resize-vertical"
                        />
                    </div>

                    <div className="grid grid-cols-2 gap-4">
                        <div>
                            <label
                                htmlFor="steps"
                                className="block font-semibold text-gray-800 mb-2"
                            >
                                Steps: {steps}
                            </label>
                            <input
                                type="range"
                                id="steps"
                                value={steps}
                                onChange={(e) =>
                                    setSteps(parseInt(e.target.value))
                                }
                                min="1"
                                max="150"
                                disabled={loading}
                                className="w-full cursor-pointer"
                            />
                        </div>

                        <div>
                            <label
                                htmlFor="cfgScale"
                                className="block font-semibold text-gray-800 mb-2"
                            >
                                CFG Scale: {cfgScale.toFixed(1)}
                            </label>
                            <input
                                type="range"
                                id="cfgScale"
                                value={cfgScale}
                                onChange={(e) =>
                                    setCfgScale(parseFloat(e.target.value))
                                }
                                min="1"
                                max="20"
                                step="0.5"
                                disabled={loading}
                                className="w-full cursor-pointer"
                            />
                        </div>
                    </div>
                </div>
            )}

            <button
                type="submit"
                disabled={loading || !prompt.trim()}
                className="w-full py-3 bg-gradient-to-r from-primary to-secondary text-white rounded-lg font-bold text-lg shadow-md hover:shadow-lg hover:-translate-y-0.5 disabled:opacity-60 disabled:cursor-not-allowed transition-all"
            >
                {loading ? "🔄 Generating..." : "🎲 Spin the Gacha!"}
            </button>
        </form>
    );
};

export default PromptInput;

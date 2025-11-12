import React from "react";

interface ImageGalleryProps {
    images: string[];
    loading: boolean;
}

const ImageGallery: React.FC<ImageGalleryProps> = ({ images, loading }) => {
    if (loading) {
        return (
            <div className="flex flex-col items-center justify-center min-h-80 bg-white rounded-lg shadow-lg">
                <div className="w-12 h-12 border-4 border-gray-200 border-t-primary rounded-full animate-spin"></div>
                <p className="mt-4 text-lg text-primary font-semibold">
                    Generating your gacha images...
                </p>
            </div>
        );
    }

    if (images.length === 0) {
        return (
            <div className="flex items-center justify-center min-h-80 bg-white rounded-lg shadow-lg">
                <p className="text-2xl text-gray-400 text-center">
                    🎨 Your generated images will appear here
                </p>
            </div>
        );
    }

    return (
        <div className="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-4 gap-6">
            {images.map((image, index) => (
                <div
                    key={index}
                    className="relative rounded-lg overflow-hidden bg-white shadow-md hover:shadow-xl hover:-translate-y-1 transition-all cursor-pointer group"
                >
                    <img
                        src={image}
                        alt={`Generated ${index + 1}`}
                        className="w-full h-full object-cover aspect-square"
                    />
                    <div className="absolute bottom-0 left-0 right-0 bg-gradient-to-t from-black/70 to-transparent p-4 transform translate-y-full group-hover:translate-y-0 transition-transform">
                        <p className="text-white font-semibold text-sm">
                            Image {index + 1}
                        </p>
                    </div>
                </div>
            ))}
        </div>
    );
};

export default ImageGallery;

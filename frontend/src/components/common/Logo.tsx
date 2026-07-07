import { BookOpen } from "lucide-react";

export const Logo = () => {
    return (
        <div className="inline-block">
            <div 
                className="w-10 h-10 bg-blue-600 rounded-xl flex items-center justify-center font-bold text-xl shadow-lg shadow-indigo-500/30 hover:shadow-xl hover:shadow-indigo-500/40 hover:scale-105 transition-all duration-300 ease-in-out cursor-pointer"
            >
                <BookOpen size={24} className="text-white" />
            </div>
        </div>
    );
};
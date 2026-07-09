import { BookOpen } from "lucide-react";

export const Logo = () => {
    return (
            <div className="w-10 h-10 rounded-lg bg-primary flex items-center justify-center group-hover:scale-105 transition-transform">
              <BookOpen className="w-6 h-6 text-primary-foreground" />
            </div>
    );
};
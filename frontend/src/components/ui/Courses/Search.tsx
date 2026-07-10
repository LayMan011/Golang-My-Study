import { Search } from 'lucide-react';

export const SearchBar = ({searchQuery, setSearchQuery}: {searchQuery: string, setSearchQuery: React.Dispatch<React.SetStateAction<string>>}) => {
    return (
        <div className="relative">
            <Search className="absolute left-4 top-1/2 -translate-y-1/2 w-5 h-5 text-muted-foreground" />
            <input
            type="text"
            placeholder="Поиск курсов..."
            value={searchQuery}
            onChange={(e) => setSearchQuery(e.target.value)}
            className="w-full pl-12 pr-4 py-4 bg-card border-2 border-border rounded-xl text-foreground placeholder:text-muted-foreground focus:border-primary focus:outline-none transition-colors"
            />
        </div>
    )
}
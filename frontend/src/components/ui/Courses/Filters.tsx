import { Filter, ChevronDown } from 'lucide-react';

const levels = [
  { id: 'all', name: 'Все уровни' },
  { id: 'beginner', name: 'Начальный' },
  { id: 'intermediate', name: 'Средний' },
  { id: 'advanced', name: 'Продвинутый' },
];

const priceRanges = [
  { id: 'all', name: 'Любая цена' },
  { id: 'free', name: 'Бесплатные' },
  { id: 'low', name: 'До 3000 ₽' },
  { id: 'medium', name: '3000 - 5000 ₽' },
  { id: 'high', name: 'Более 5000 ₽' },
];

export const Filters = ({
    showFilters,
    setShowFilters,
    selectedLevel,
    setSelectedLevel,
    selectedPrice,
    setSelectedPrice,

    setSelectedSubject,
    setSearchQuery,
}: {
    showFilters: boolean,
    setShowFilters: React.Dispatch<React.SetStateAction<boolean>>,
    selectedLevel: string,
    setSelectedLevel: React.Dispatch<React.SetStateAction<string>>,
    selectedPrice: string,
    setSelectedPrice: React.Dispatch<React.SetStateAction<string>>,

    setSelectedSubject: React.Dispatch<React.SetStateAction<string>>,
    setSearchQuery: React.Dispatch<React.SetStateAction<string>>
}) => {
    return (
        <aside className="lg:w-64 shrink-0">
            <div className="lg:sticky lg:top-24">
            <button
                onClick={() => setShowFilters(!showFilters)}
                className="lg:hidden w-full flex items-center justify-between px-4 py-3 bg-card border border-border rounded-lg mb-4"
            >
                <span className="flex items-center gap-2 font-medium text-foreground">
                <Filter className="w-5 h-5" />
                Фильтры
                </span>
                <ChevronDown className={`w-5 h-5 transition-transform ${showFilters ? 'rotate-180' : ''}`} />
            </button>

            <div className={`space-y-6 ${showFilters ? 'block' : 'hidden lg:block'}`}>
                {/* Level Filter */}
                <div className="bg-card border border-border rounded-xl p-4">
                <h3 className="font-semibold text-foreground mb-3">Уровень сложности</h3>
                <div className="space-y-2">
                    {levels.map(level => (
                    <label key={level.id} className="flex items-center gap-2 cursor-pointer">
                        <input
                        type="radio"
                        name="level"
                        checked={selectedLevel === level.id}
                        onChange={() => setSelectedLevel(level.id)}
                        className="w-4 h-4 text-primary accent-primary"
                        />
                        <span className="text-sm text-foreground">{level.name}</span>
                    </label>
                    ))}
                </div>
                </div>

                {/* Price Filter */}
                <div className="bg-card border border-border rounded-xl p-4">
                <h3 className="font-semibold text-foreground mb-3">Цена</h3>
                <div className="space-y-2">
                    {priceRanges.map(range => (
                    <label key={range.id} className="flex items-center gap-2 cursor-pointer">
                        <input
                        type="radio"
                        name="price"
                        checked={selectedPrice === range.id}
                        onChange={() => setSelectedPrice(range.id)}
                        className="w-4 h-4 text-primary accent-primary"
                        />
                        <span className="text-sm text-foreground">{range.name}</span>
                    </label>
                    ))}
                </div>
                </div>

                <button
                onClick={() => {
                    setSelectedSubject('all');
                    setSelectedLevel('all');
                    setSelectedPrice('all');
                    setSearchQuery('');
                }}
                className="w-full px-4 py-2 bg-muted text-foreground rounded-lg hover:bg-muted/80 transition-colors"
                >
                Сбросить фильтры
                </button>
            </div>
            </div>
        </aside>
    )
}
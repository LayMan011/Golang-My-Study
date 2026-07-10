

const sortOptions = [
  { id: 'popular', name: 'По популярности' },
  { id: 'rating', name: 'По рейтингу' },
  { id: 'price-low', name: 'Сначала дешевые' },
  { id: 'price-high', name: 'Сначала дорогие' },
  { id: 'new', name: 'Новинки' },
];

interface filteredCourses {
    id: number;
    subject: string;
    subjectName: string;
    title: string;
    description: string;
    rating: number;
    reviews: number;
    price: number;
    students: number;
    level: string;
    duration: string;
    format: string;
}

export const Sort = ({sortBy, setSortBy, filteredCourses }: { sortBy: string, setSortBy: React.Dispatch<React.SetStateAction<string>>, filteredCourses: filteredCourses[]}) => {
    return (
        <div className="flex items-center justify-between mb-6">
        <div className="text-sm text-muted-foreground">
            Найдено курсов: <span className="font-semibold text-foreground">{filteredCourses.length}</span>
        </div>
        <div className="flex items-center gap-2">
            <label htmlFor="sort" className="text-sm text-muted-foreground">
            Сортировка:
            </label>
            <select
            id="sort"
            value={sortBy}
            onChange={(e) => setSortBy(e.target.value)}
            className="px-3 py-1.5 bg-card border border-border rounded-lg text-sm text-foreground focus:border-primary focus:outline-none"
            >
            {sortOptions.map(option => (
                <option key={option.id} value={option.id}>
                {option.name}
                </option>
            ))}
            </select>
        </div>
        </div>
    )
}

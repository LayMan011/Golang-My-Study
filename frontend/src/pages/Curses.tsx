import { useState } from 'react';
import { motion } from 'motion/react';
import { Search, Filter, Star, Users, ChevronDown } from 'lucide-react';

const allCourses = [
  {
    id: 1,
    subject: 'russian',
    subjectName: 'Русский язык',
    title: 'Полная подготовка к ЕГЭ',
    description: 'Комплексный курс по всем разделам экзамена',
    rating: 4.9,
    reviews: 234,
    price: 4990,
    students: 1234,
    level: 'advanced',
    duration: '6 месяцев',
    format: 'video',
  },
  {
    id: 2,
    subject: 'english',
    subjectName: 'Английский язык',
    title: 'Подготовка к ЕГЭ за 3 месяца',
    description: 'Интенсивный курс для быстрой подготовки',
    rating: 4.8,
    reviews: 189,
    price: 5990,
    students: 987,
    level: 'intermediate',
    duration: '3 месяца',
    format: 'video',
  },
  {
    id: 3,
    subject: 'math',
    subjectName: 'Математика',
    title: 'Профильная математика',
    description: 'Углубленная подготовка к профильному ЕГЭ',
    rating: 4.9,
    reviews: 456,
    price: 5490,
    students: 2134,
    level: 'advanced',
    duration: '8 месяцев',
    format: 'video',
  },
  {
    id: 4,
    subject: 'physics',
    subjectName: 'Физика',
    title: 'Механика и термодинамика',
    description: 'Детальный разбор сложных тем',
    rating: 4.7,
    reviews: 123,
    price: 4490,
    students: 654,
    level: 'intermediate',
    duration: '5 месяцев',
    format: 'video',
  },
  {
    id: 5,
    subject: 'chemistry',
    subjectName: 'Химия',
    title: 'Органическая химия',
    description: 'Полный разбор органической химии',
    rating: 4.8,
    reviews: 167,
    price: 4790,
    students: 789,
    level: 'intermediate',
    duration: '4 месяца',
    format: 'video',
  },
  {
    id: 6,
    subject: 'biology',
    subjectName: 'Биология',
    title: 'Полный курс подготовки',
    description: 'Все темы по биологии для ЕГЭ',
    rating: 4.9,
    reviews: 201,
    price: 4990,
    students: 1098,
    level: 'advanced',
    duration: '6 месяцев',
    format: 'video',
  },
  {
    id: 7,
    subject: 'history',
    subjectName: 'История',
    title: 'История России',
    description: 'От древности до современности',
    rating: 4.7,
    reviews: 145,
    price: 4290,
    students: 567,
    level: 'beginner',
    duration: '5 месяцев',
    format: 'text',
  },
  {
    id: 8,
    subject: 'social',
    subjectName: 'Обществознание',
    title: 'Обществознание: теория и практика',
    description: 'Подготовка с нуля до высоких баллов',
    rating: 4.8,
    reviews: 198,
    price: 4490,
    students: 876,
    level: 'beginner',
    duration: '4 месяца',
    format: 'mixed',
  },
];

const subjects = [
  { id: 'all', name: 'Все предметы', color: 'primary' },
  { id: 'russian', name: 'Русский язык', color: 'subject-russian' },
  { id: 'english', name: 'Английский язык', color: 'subject-english' },
  { id: 'math', name: 'Математика', color: 'subject-math' },
  { id: 'physics', name: 'Физика', color: 'subject-physics' },
  { id: 'chemistry', name: 'Химия', color: 'subject-chemistry' },
  { id: 'biology', name: 'Биология', color: 'subject-biology' },
  { id: 'history', name: 'История', color: 'subject-history' },
  { id: 'social', name: 'Обществознание', color: 'subject-social' },
];

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

const sortOptions = [
  { id: 'popular', name: 'По популярности' },
  { id: 'rating', name: 'По рейтингу' },
  { id: 'price-low', name: 'Сначала дешевые' },
  { id: 'price-high', name: 'Сначала дорогие' },
  { id: 'new', name: 'Новинки' },
];

export const Courses = () => {
  const [searchQuery, setSearchQuery] = useState('');
  const [selectedSubject, setSelectedSubject] = useState('all');
  const [selectedLevel, setSelectedLevel] = useState('all');
  const [selectedPrice, setSelectedPrice] = useState('all');
  const [sortBy, setSortBy] = useState('popular');
  const [showFilters, setShowFilters] = useState(false);

  const filteredCourses = allCourses.filter(course => {
    const matchesSearch = course.title.toLowerCase().includes(searchQuery.toLowerCase()) ||
      course.description.toLowerCase().includes(searchQuery.toLowerCase());
    const matchesSubject = selectedSubject === 'all' || course.subject === selectedSubject;
    const matchesLevel = selectedLevel === 'all' || course.level === selectedLevel;
    const matchesPrice = selectedPrice === 'all' ||
      (selectedPrice === 'free' && course.price === 0) ||
      (selectedPrice === 'low' && course.price < 3000) ||
      (selectedPrice === 'medium' && course.price >= 3000 && course.price <= 5000) ||
      (selectedPrice === 'high' && course.price > 5000);

    return matchesSearch && matchesSubject && matchesLevel && matchesPrice;
  });

  return (
    <div className="min-h-screen bg-background">
      {/* Search Hero */}
      <section className="bg-linear-to-br from-primary/10 via-background to-accent/5 py-12">
        <div className="container mx-auto px-4 sm:px-6 lg:px-8">
          <div className="max-w-3xl mx-auto">
            <h1 className="text-3xl sm:text-4xl font-bold text-foreground mb-6 text-center">
              Найди свой идеальный курс
            </h1>
            
            {/* Search Bar */}
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

            {/* Subject Pills */}
            <div className="flex flex-wrap gap-2 mt-6">
              {subjects.map(subject => (
                <button
                  key={subject.id}
                  onClick={() => setSelectedSubject(subject.id)}
                  className={`px-4 py-2 rounded-full text-sm font-medium transition-all ${
                    selectedSubject === subject.id
                      ? 'text-white shadow-md scale-105'
                      : 'bg-card text-muted-foreground hover:bg-muted border border-border'
                  }`}
                  style={{
                    backgroundColor: selectedSubject === subject.id ? `var(--${subject.color})` : undefined,
                  }}
                >
                  {subject.name}
                </button>
              ))}
            </div>
          </div>
        </div>
      </section>

      {/* Filters and Results */}
      <section className="py-8">
        <div className="container mx-auto px-4 sm:px-6 lg:px-8">
          <div className="flex flex-col lg:flex-row gap-8">
            {/* Sidebar Filters */}
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

            {/* Results */}
            <div className="flex-1">
              {/* Sort and Count */}
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

              {/* Course Grid */}
              <div className="grid md:grid-cols-2 gap-6">
                {filteredCourses.map((course, index) => (
                  <motion.div
                    key={course.id}
                    initial={{ opacity: 0, y: 20 }}
                    animate={{ opacity: 1, y: 0 }}
                    transition={{ delay: index * 0.05 }}
                    whileHover={{ y: -4 }}
                    className="bg-card rounded-2xl overflow-hidden shadow-sm hover:shadow-lg transition-all border border-border"
                  >
                    <div className="p-6">
                      <div
                        className="inline-block px-3 py-1 rounded-full text-xs font-semibold text-white mb-4"
                        style={{ backgroundColor: `var(--subject-${course.subject})` }}
                      >
                        {course.subjectName}
                      </div>
                      <h3 className="text-xl font-semibold text-foreground mb-2">
                        {course.title}
                      </h3>
                      <p className="text-sm text-muted-foreground mb-4">{course.description}</p>
                      
                      <div className="flex items-center gap-4 mb-4 text-sm">
                        <div className="flex items-center gap-1">
                          <Star className="w-4 h-4 text-accent fill-accent" />
                          <span className="font-semibold text-foreground">{course.rating}</span>
                          <span className="text-muted-foreground">({course.reviews})</span>
                        </div>
                        <div className="flex items-center gap-1 text-muted-foreground">
                          <Users className="w-4 h-4" />
                          <span>{course.students}</span>
                        </div>
                      </div>

                      <div className="flex flex-wrap gap-2 mb-4">
                        <span className="px-2 py-1 bg-muted rounded text-xs text-muted-foreground">
                          {course.duration}
                        </span>
                        <span className="px-2 py-1 bg-muted rounded text-xs text-muted-foreground capitalize">
                          {course.level === 'beginner' ? 'Начальный' : course.level === 'intermediate' ? 'Средний' : 'Продвинутый'}
                        </span>
                      </div>

                      <div className="flex items-center justify-between pt-4 border-t border-border">
                        <div className="text-2xl font-bold text-foreground">
                          {course.price.toLocaleString('ru-RU')} ₽
                        </div>
                        <button className="px-4 py-2 rounded-lg bg-primary text-primary-foreground font-medium hover:bg-primary/90 transition-colors">
                          Подробнее
                        </button>
                      </div>
                    </div>
                  </motion.div>
                ))}
              </div>

              {/* Pagination */}
              {filteredCourses.length > 6 && (
                <div className="flex justify-center gap-2 mt-8">
                  <button className="px-4 py-2 rounded-lg bg-card border border-border text-foreground hover:bg-muted transition-colors">
                    Предыдущая
                  </button>
                  <button className="px-4 py-2 rounded-lg bg-primary text-primary-foreground font-medium">
                    1
                  </button>
                  <button className="px-4 py-2 rounded-lg bg-card border border-border text-foreground hover:bg-muted transition-colors">
                    2
                  </button>
                  <button className="px-4 py-2 rounded-lg bg-card border border-border text-foreground hover:bg-muted transition-colors">
                    3
                  </button>
                  <button className="px-4 py-2 rounded-lg bg-card border border-border text-foreground hover:bg-muted transition-colors">
                    Следующая
                  </button>
                </div>
              )}

              {filteredCourses.length === 0 && (
                <div className="text-center py-12">
                  <p className="text-lg text-muted-foreground mb-4">
                    Курсы не найдены
                  </p>
                  <button
                    onClick={() => {
                      setSelectedSubject('all');
                      setSelectedLevel('all');
                      setSelectedPrice('all');
                      setSearchQuery('');
                    }}
                    className="px-6 py-2 bg-primary text-primary-foreground rounded-lg hover:bg-primary/90 transition-colors"
                  >
                    Сбросить фильтры
                  </button>
                </div>
              )}
            </div>
          </div>
        </div>
      </section>
    </div>
  );
}

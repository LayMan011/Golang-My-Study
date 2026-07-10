import { useState } from 'react';

import { SearchBar, SubjectSerchHelpers, Filters, Sort, FiltersCourses, Pagination } from '@/components/ui/Courses'

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

export const Courses = () => {
  const [searchQuery, setSearchQuery] = useState('');
  const [selectedSubject, setSelectedSubject] = useState('all');
  const [selectedLevel, setSelectedLevel] = useState('all');
  const [selectedPrice, setSelectedPrice] = useState('all');
  const [sortBy, setSortBy] = useState('popular');
  const [showFilters, setShowFilters] = useState(false);

  const [currentPage, setCurrentPage] = useState(1);

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

  const itemsPerPage = 10;
  const totalPages = Math.ceil(filteredCourses.length / itemsPerPage);

  const handlePageChange = (page: number) => {
    setCurrentPage(page);  // Обновляем состояние
    window.scrollTo({ top: 0, behavior: 'smooth' });
  };

  return (
    <div className="min-h-screen bg-background">
      {/* Search Hero */}
      <section className="bg-linear-to-br from-primary/10 via-background to-accent/5 py-12">
        <div className="container mx-auto px-4 sm:px-6 lg:px-8">
          <div className="max-w-3xl mx-auto">
            <h1 className="text-3xl sm:text-4xl font-bold text-foreground mb-6 text-center">
              Найди свой идеальный курс
            </h1>
            
            <SearchBar searchQuery={searchQuery} setSearchQuery={setSearchQuery} />

            <SubjectSerchHelpers selectedSubject={selectedSubject} setSelectedSubject={setSelectedSubject} />
          </div>
        </div>
      </section>

      {/* Filters and Results */}
      <section className="py-8">
        <div className="container mx-auto px-4 sm:px-6 lg:px-8">
          <div className="flex flex-col lg:flex-row gap-8">
            <Filters 
              showFilters={showFilters}
              setShowFilters={setShowFilters}
              selectedLevel={selectedLevel}
              setSelectedLevel={setSelectedLevel}
              selectedPrice={selectedPrice}
              setSelectedPrice={setSelectedPrice}
              setSelectedSubject={setSelectedSubject}
              setSearchQuery={setSearchQuery} 
            />

            {/* Results */}
            <div className="flex-1">
              <Sort sortBy={sortBy} setSortBy={setSortBy} filteredCourses={filteredCourses} />

              <FiltersCourses filteredCourses={filteredCourses} />

              {/* Pagination */}
              <Pagination currentPage={currentPage} totalPages={totalPages} onPageChange={handlePageChange} />

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

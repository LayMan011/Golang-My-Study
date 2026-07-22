import { useState, useEffect } from 'react';

import { SearchBar, SubjectSerchHelpers, Filters, Sort, FiltersCourses, Pagination } from '@/components/ui/Courses'

interface Course {
  id: number;
  subject: string;
  title: string;
  description: string;
  rating: number;
  number_of_ratings: number;
  price: number;
  number_of_users: number;
  level: string;
  duration: string;
  format: string;
}

const getCourses = async (): Promise<Course[]> => {
  try {
    const response = await fetch('http://localhost:5050/api/v1/themes');
    
    if (!response.ok) {
      throw new Error(`Ошибка: ${response.status}`);
    }

    const courses: Course[] = await response.json();
    return courses;
  } catch (error) {
    console.error('Ошибка загрузки:', error);
    return [];
  }
};

const subjectColors: Record<string, string> = {
  'Русский язык': 'subject-russian',
  'Английский язык': 'subject-english',
  'Математика': 'subject-math',
  'Физика': 'subject-physics',
  'Химия': 'subject-chemistry',
  'Биология': 'subject-biology',
  'История': 'subject-history',
  'Обществознание': 'subject-society',
  'Информатика': 'subject-it',
  'Литература': 'subject-literature',
};

const addColorsToCourses = (courses: Omit<Course, 'color'>[]): Course[] => {
  return courses.map(course => ({
    ...course,
    color: subjectColors[course.subject],
  }));
};

const Courses = () => {
  const [searchQuery, setSearchQuery] = useState('');
  const [selectedSubject, setSelectedSubject] = useState('all');
  const [selectedLevel, setSelectedLevel] = useState('all');
  const [selectedPrice, setSelectedPrice] = useState('all');
  const [sortBy, setSortBy] = useState('popular');
  const [showFilters, setShowFilters] = useState(false);

  const [currentPage, setCurrentPage] = useState(1);
  const [courses, setCourses] = useState<Course[]>([]);
  // const [loading, setLoading] = useState(true);

  useEffect(() => {
    const fetchCourses = async () => {
      // setLoading(true);
      const data = await getCourses();
      setCourses(addColorsToCourses(data));
      // setLoading(false);
    };

    fetchCourses();
  }, []);

  const filteredCourses = courses.filter(course => {
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

  const sortedCourses = [...filteredCourses].sort((a, b) => {
    switch (sortBy) {
      case 'rating':
        return b.rating - a.rating;
      case 'price-low':
        return a.price - b.price;
      case 'price-high':
        return b.price - a.price;
      case 'new':
        return b.id - a.id;
      case 'popular':
      default:
        return b.number_of_users - a.number_of_users;
    }
  });

  const itemsPerPage = 10;
  const totalPages = Math.ceil(filteredCourses.length / itemsPerPage);

  const paginatedCourses = sortedCourses.slice(
    (currentPage - 1) * itemsPerPage,
    currentPage * itemsPerPage,
  );

  const handlePageChange = (page: number) => {
    setCurrentPage(page);
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
              <Sort sortBy={sortBy} setSortBy={setSortBy} filteredCourses={paginatedCourses} />

              <FiltersCourses filteredCourses={paginatedCourses} />

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

export default Courses
import { motion } from 'motion/react';
import { ArrowRight, Users, Star } from 'lucide-react';
import { Link } from 'react-router';
import {useState, useEffect} from 'react'

interface Course {
  id: number
  title: string
  subject: string
  rating: number
  number_of_ratings: number
  price: number
  number_of_users: number
  color: string 
}

const getCoursesLimit6 = async (): Promise<Course[]> => {
  try {
    const response = await fetch('http://localhost:5050/api/v1/themes?limit=6');
    
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


export const PopularCurses = () => {
  const [courses, setCourses] = useState<Course[]>([]);
  const [loading, setLoading] = useState(true);

  useEffect(() => {
    const fetchCourses = async () => {
      setLoading(true);
      const data = await getCoursesLimit6();
      setCourses(addColorsToCourses(data));
      setLoading(false);
    };

    fetchCourses();
  }, []);

  if (loading) {
    return (
      <section className="py-20 bg-muted/30">
        <div className="container mx-auto px-4">
          <div className="grid md:grid-cols-2 lg:grid-cols-3 gap-6">
            {[1, 2, 3, 4, 5, 6].map((index) => (
              <div key={index} className="bg-card rounded-2xl p-6 border border-border">
                <div className="h-6 w-24 bg-muted animate-pulse rounded-full mb-4" />
                <div className="h-6 w-48 bg-muted animate-pulse rounded mb-4" />
                <div className="h-4 w-20 bg-muted animate-pulse rounded" />
              </div>
            ))}
          </div>
        </div>
      </section>
    );
  }

  return (
    <section className="py-20 bg-muted/30">
      <div className="container mx-auto px-4 sm:px-6 lg:px-8">
        <div className="text-center mb-12">
          <h2 className="text-3xl sm:text-4xl font-bold text-foreground mb-4">
            Популярные курсы
          </h2>
          <p className="text-lg text-muted-foreground max-w-2xl mx-auto">
            Самые востребованные курсы от наших преподавателей
          </p>
        </div>

        <div className="grid md:grid-cols-2 lg:grid-cols-3 gap-6 mb-8">
          {courses.map((course, index) => (
            <motion.div
              key={course.id}
              initial={{ opacity: 0, y: 20 }}
              whileInView={{ opacity: 1, y: 0 }}
              viewport={{ once: true }}
              transition={{ delay: index * 0.1 }}
              whileHover={{ y: -4 }}
              className="bg-card rounded-2xl overflow-hidden shadow-sm hover:shadow-lg border border-border"
            >
              <div className="p-6">
                <div
                  className="inline-block px-3 py-1 rounded-full text-xs font-semibold text-white mb-4"
                  style={{ backgroundColor: `var(--${course.color})` }}
                >
                  {course.subject}
                </div>
                <h3 className="text-xl font-semibold text-foreground mb-4">
                  {course.title}
                </h3>
                <div className="flex items-center gap-4 mb-4">
                  <div className="flex items-center gap-1">
                    <Star className="w-4 h-4 text-accent fill-accent" />
                    <span className="font-semibold text-foreground">{course.rating}</span>
                    <span className="text-sm text-muted-foreground">({course.number_of_ratings})</span>
                  </div>
                  <div className="flex items-center gap-1 text-sm text-muted-foreground">
                    <Users className="w-4 h-4" />
                    <span>{course.number_of_users}</span>
                  </div>
                </div>
                <div className="flex items-center justify-between pt-4 border-t border-border">
                  <div>
                    <div className="text-2xl font-bold text-foreground">
                      {course.price.toLocaleString('ru-RU')} ₽
                    </div>
                  </div>
                  <button className="px-4 py-2 rounded-lg bg-primary text-primary-foreground font-medium hover:bg-primary/90 transition-colors">
                    Подробнее
                  </button>
                </div>
              </div>
            </motion.div>
          ))}
        </div>

        <div className="text-center">
          <Link
            to="/courses"
            className="inline-flex items-center gap-2 px-6 py-3 rounded-lg bg-card border-2 border-border text-foreground font-semibold hover:bg-muted transition-colors"
          >
            Смотреть все курсы
            <ArrowRight className="w-5 h-5" />
          </Link>
        </div>
      </div>
    </section>
  )
}
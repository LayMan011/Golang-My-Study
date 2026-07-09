import { motion } from 'motion/react';
import { ArrowRight, Users, Star } from 'lucide-react';
import { Link } from 'react-router';

const courses = [
  {
    id: 1,
    subject: 'Русский язык',
    title: 'Полная подготовка к ЕГЭ',
    rating: 4.9,
    reviews: 234,
    price: 4990,
    students: 1234,
    color: 'subject-russian',
  },
  {
    id: 2,
    subject: 'Английский язык',
    title: 'Подготовка к ЕГЭ за 3 месяца',
    rating: 4.8,
    reviews: 189,
    price: 5990,
    students: 987,
    color: 'subject-english',
  },
  {
    id: 3,
    subject: 'Математика',
    title: 'Профильная математика',
    rating: 4.9,
    reviews: 456,
    price: 5490,
    students: 2134,
    color: 'subject-math',
  },
  {
    id: 4,
    subject: 'Физика',
    title: 'Механика и термодинамика',
    rating: 4.7,
    reviews: 123,
    price: 4490,
    students: 654,
    color: 'subject-physics',
  },
  {
    id: 5,
    subject: 'Химия',
    title: 'Органическая химия',
    rating: 4.8,
    reviews: 167,
    price: 4790,
    students: 789,
    color: 'subject-chemistry',
  },
  {
    id: 6,
    subject: 'Биология',
    title: 'Полный курс подготовки',
    rating: 4.9,
    reviews: 201,
    price: 4990,
    students: 1098,
    color: 'subject-biology',
  },
];

export const PopularCurses = () => {
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
                className="bg-card rounded-2xl overflow-hidden shadow-sm hover:shadow-lg transition-all border border-border"
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
                      <span className="text-sm text-muted-foreground">({course.reviews})</span>
                    </div>
                    <div className="flex items-center gap-1 text-sm text-muted-foreground">
                      <Users className="w-4 h-4" />
                      <span>{course.students}</span>
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
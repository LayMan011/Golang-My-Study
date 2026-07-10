import { motion } from 'motion/react';
import { Star, Users } from 'lucide-react';

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

export const FiltersCourses = ({ filteredCourses }: {filteredCourses: filteredCourses[] }) => {
    return (
        <div className="grid md:grid-cols-2 gap-6">
        {filteredCourses.map((course, index) => (
            <motion.div
            key={course.id}
            initial={{ opacity: 0, y: 20 }}
            animate={{ opacity: 1, y: 0 }}
            transition={{ delay: index * 0.05 }}
            whileHover={{ y: -4 }}
            className="bg-card rounded-2xl overflow-hidden shadow-sm hover:shadow-lg border border-border"
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
    )
}
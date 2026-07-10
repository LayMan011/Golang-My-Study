import { motion } from 'motion/react';
import { Award } from 'lucide-react';

export const CompletedTab = ({ completedCourses }: {
    completedCourses: {
        id: number;
        subject: string;
        subjectName: string;
        title: string;
        score: number;
        completedDate: string;
        certificate: boolean;
    }[]
}) => {
    return (
        <motion.div
            key="completed"
            initial={{ opacity: 0, y: 20 }}
            animate={{ opacity: 1, y: 0 }}
            transition={{ duration: 0.3 }}
            className="grid md:grid-cols-2 gap-6"
        >
            {completedCourses.map((course, index) => (
            <motion.div
                key={course.id}
                initial={{ opacity: 0, y: 20 }}
                animate={{ opacity: 1, y: 0 }}
                transition={{ delay: index * 0.1 }}
                className="bg-card rounded-2xl p-6 border border-border shadow-sm hover:shadow-md transition-shadow relative overflow-hidden"
            >
                {/* Success Badge */}
                <div className="absolute top-4 right-4">
                <div className="w-12 h-12 bg-green-500/10 rounded-full flex items-center justify-center">
                    <Award className="w-6 h-6 text-green-500" />
                </div>
                </div>

                <div
                className="inline-block px-3 py-1 rounded-full text-xs font-semibold text-white mb-3"
                style={{ backgroundColor: `var(--subject-${course.subject})` }}
                >
                {course.subjectName}
                </div>
                <h3 className="text-xl font-semibold text-foreground mb-4 pr-12">
                {course.title}
                </h3>

                <div className="space-y-3">
                <div className="flex items-center justify-between py-3 border-t border-border">
                    <span className="text-sm text-muted-foreground">Результат</span>
                    <span className="text-2xl font-bold text-foreground">{course.score} баллов</span>
                </div>
                <div className="flex items-center justify-between">
                    <span className="text-sm text-muted-foreground">Завершено</span>
                    <span className="text-sm text-foreground">{course.completedDate}</span>
                </div>
                {course.certificate && (
                    <button className="w-full mt-4 px-4 py-2 bg-accent/10 text-accent rounded-lg hover:bg-accent/20 transition-colors font-medium border border-accent/20">
                    Скачать сертификат
                    </button>
                )}
                </div>
            </motion.div>
            ))}
        </motion.div>
    )
}
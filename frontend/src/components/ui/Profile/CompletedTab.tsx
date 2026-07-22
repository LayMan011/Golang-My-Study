import { motion } from 'motion/react';
import { Award } from 'lucide-react';
import type { UserCourseProgress } from '@/types';

export const CompletedTab = ({
  completedCourses,
}: {
  completedCourses: UserCourseProgress[];
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
            style={{
              backgroundColor: `var(--subject-${course.subject})`,
            }}
          >
            <SubjectName subcject={course.subject} />
          </div>

          <h3 className="text-xl font-semibold text-foreground mb-4 pr-12">
            {course.title}
          </h3>

          <div className="space-y-3">
            <div className="flex items-center justify-between">
              <span className="text-sm text-muted-foreground">
                Завершено
              </span>
              <span className="text-sm text-foreground">
                {course.completed_at
                  ? new Date(course.completed_at).toLocaleDateString(
                      'ru-RU',
                    )
                  : '—'}
              </span>
            </div>

            <div className="flex items-center justify-between">
              <span className="text-sm text-muted-foreground">
                Результат
              </span>
              <span className="text-sm font-semibold text-foreground">
                {course.percentages}%
              </span>
            </div>
          </div>
        </motion.div>
      ))}
    </motion.div>
  );
};

function SubjectName({ subcject }: { subcject: string }) {
  switch (subcject) {
    case 'russian':
      return 'Русский язык';
    case 'english':
      return 'Английский язык';
    case 'math':
      return 'Математика';
    case 'physics':
      return 'Физика';
    case 'chemistry':
      return 'Химия';
    case 'biology':
      return 'Биология';
    case 'history':
      return 'История';
    case 'social':
      return 'Обществознание';
  }

  return 'Общий';
}
import { motion } from 'motion/react';
import * as Progress from '@radix-ui/react-progress';
import type { UserCourseProgress } from '@/types';

export const ProgressTab = ({
  inProgressCourses,
}: {
  inProgressCourses: UserCourseProgress[];
}) => {
  return (
    <motion.div
      key="progress"
      initial={{ opacity: 0, y: 20 }}
      animate={{ opacity: 1, y: 0 }}
      transition={{ duration: 0.3 }}
      className="space-y-6"
    >
      {inProgressCourses.map((course, index) => (
        <motion.div
          key={course.id}
          initial={{ opacity: 0, y: 20 }}
          animate={{ opacity: 1, y: 0 }}
          transition={{ delay: index * 0.1 }}
          className="bg-card rounded-2xl p-6 border border-border shadow-sm hover:shadow-md transition-shadow"
        >
          <div className="flex flex-col md:flex-row md:items-center justify-between gap-4 mb-4">
            <div className="flex-1">
              <div
                className="inline-block px-3 py-1 rounded-full text-xs font-semibold text-white mb-3"
                style={{
                  backgroundColor: `var(--subject-${course.subject})`,
                }}
              >
                <SubjectName subcject={course.subject} />
              </div>
              <h3 className="text-xl font-semibold text-foreground mb-2">
                {course.title}
              </h3>
              <p className="text-sm text-muted-foreground">
                {/* TODO: форматировать added_at в “2 часа назад” */}
                Последняя активность:{' '}
                {new Date(course.addition_at).toLocaleString('ru-RU')}
              </p>
            </div>
            <button className="px-6 py-3 bg-primary text-primary-foreground rounded-lg hover:bg-primary/90 transition-colors font-medium">
              Продолжить
            </button>
          </div>

          {/* Progress Bar */}
          <div className="space-y-2">
            <div className="flex items-center justify-between text-sm">
              <span className="text-muted-foreground">
                {course.completed_lessons} из {course.total_lessons} уроков
              </span>
              <span className="font-semibold text-foreground">
                {course.percentages}%
              </span>
            </div>
            <Progress.Root
              className="relative h-3 overflow-hidden rounded-full bg-muted"
              value={course.percentages}
            >
              <Progress.Indicator
                className="h-full bg-accent transition-all duration-300 ease-out rounded-full"
                style={{
                  width: `${course.percentages}%`,
                }}
              />
            </Progress.Root>
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
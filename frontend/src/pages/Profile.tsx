import { useState } from 'react';
import { motion } from 'motion/react';
import { LogOut, Award, TrendingUp, Calendar } from 'lucide-react';
import * as Progress from '@radix-ui/react-progress';

const inProgressCourses = [
  {
    id: 1,
    subject: 'russian',
    subjectName: 'Русский язык',
    title: 'Полная подготовка к ЕГЭ',
    progress: 65,
    totalLessons: 48,
    completedLessons: 31,
    lastActivity: '2 часа назад',
  },
  {
    id: 2,
    subject: 'math',
    subjectName: 'Математика',
    title: 'Профильная математика',
    progress: 42,
    totalLessons: 60,
    completedLessons: 25,
    lastActivity: '1 день назад',
  },
  {
    id: 3,
    subject: 'english',
    subjectName: 'Английский язык',
    title: 'Подготовка к ЕГЭ за 3 месяца',
    progress: 78,
    totalLessons: 36,
    completedLessons: 28,
    lastActivity: '5 часов назад',
  },
];

const completedCourses = [
  {
    id: 4,
    subject: 'physics',
    subjectName: 'Физика',
    title: 'Механика и термодинамика',
    score: 92,
    completedDate: '15 июня 2026',
    certificate: true,
  },
  {
    id: 5,
    subject: 'chemistry',
    subjectName: 'Химия',
    title: 'Органическая химия',
    score: 88,
    completedDate: '3 мая 2026',
    certificate: true,
  },
];

export const Profile = () => {
  const [activeTab, setActiveTab] = useState<'progress' | 'completed'>('progress');

  return (
    <div className="min-h-screen bg-background">
      {/* Profile Header */}
      <section className="bg-linear-to-br from-primary/10 via-background to-accent/5 py-12 border-b border-border">
        <div className="container mx-auto px-4 sm:px-6 lg:px-8">
          <div className="flex flex-col md:flex-row items-center md:items-start gap-6">
            {/* Avatar */}
            <motion.div
              initial={{ scale: 0.8, opacity: 0 }}
              animate={{ scale: 1, opacity: 1 }}
              transition={{ duration: 0.3 }}
              className="w-32 h-32 rounded-full bg-linear-to-br from-primary to-accent flex items-center justify-center text-4xl font-bold text-white shadow-xl"
            >
              ИИ
            </motion.div>

            {/* User Info */}
            <div className="flex-1 text-center md:text-left">
              <motion.h1
                initial={{ opacity: 0, y: 20 }}
                animate={{ opacity: 1, y: 0 }}
                transition={{ delay: 0.1 }}
                className="text-3xl font-bold text-foreground mb-2"
              >
                Иван Иванов
              </motion.h1>
              <motion.p
                initial={{ opacity: 0, y: 20 }}
                animate={{ opacity: 1, y: 0 }}
                transition={{ delay: 0.2 }}
                className="text-muted-foreground mb-4"
              >
                ivan.ivanov@email.com
              </motion.p>
              <motion.div
                initial={{ opacity: 0, y: 20 }}
                animate={{ opacity: 1, y: 0 }}
                transition={{ delay: 0.3 }}
                className="flex flex-wrap gap-4 justify-center md:justify-start"
              >
                <div className="flex items-center gap-2 text-sm">
                  <Calendar className="w-4 h-4 text-muted-foreground" />
                  <span className="text-muted-foreground">На платформе с января 2026</span>
                </div>
              </motion.div>
            </div>

            {/* Logout Button */}
            <motion.button
              initial={{ opacity: 0, x: 20 }}
              animate={{ opacity: 1, x: 0 }}
              transition={{ delay: 0.4 }}
              className="flex items-center gap-2 px-6 py-3 bg-destructive text-destructive-foreground rounded-lg hover:bg-destructive/90 transition-colors shadow-md"
            >
              <LogOut className="w-5 h-5" />
              Выйти
            </motion.button>
          </div>

          {/* Stats */}
          <motion.div
            initial={{ opacity: 0, y: 20 }}
            animate={{ opacity: 1, y: 0 }}
            transition={{ delay: 0.5 }}
            className="grid grid-cols-2 md:grid-cols-4 gap-4 mt-8"
          >
            <div className="bg-card rounded-xl p-4 border border-border">
              <div className="flex items-center gap-3 mb-2">
                <div className="w-10 h-10 bg-primary/10 rounded-lg flex items-center justify-center">
                  <TrendingUp className="w-5 h-5 text-primary" />
                </div>
                <div>
                  <div className="text-2xl font-bold text-foreground">3</div>
                  <div className="text-xs text-muted-foreground">В процессе</div>
                </div>
              </div>
            </div>
            <div className="bg-card rounded-xl p-4 border border-border">
              <div className="flex items-center gap-3 mb-2">
                <div className="w-10 h-10 bg-accent/10 rounded-lg flex items-center justify-center">
                  <Award className="w-5 h-5 text-accent" />
                </div>
                <div>
                  <div className="text-2xl font-bold text-foreground">2</div>
                  <div className="text-xs text-muted-foreground">Завершено</div>
                </div>
              </div>
            </div>
            <div className="bg-card rounded-xl p-4 border border-border">
              <div className="flex items-center gap-3 mb-2">
                <div className="w-10 h-10 bg-green-500/10 rounded-lg flex items-center justify-center">
                  <TrendingUp className="w-5 h-5 text-green-500" />
                </div>
                <div>
                  <div className="text-2xl font-bold text-foreground">90</div>
                  <div className="text-xs text-muted-foreground">Средний балл</div>
                </div>
              </div>
            </div>
            <div className="bg-card rounded-xl p-4 border border-border">
              <div className="flex items-center gap-3 mb-2">
                <div className="w-10 h-10 bg-purple-500/10 rounded-lg flex items-center justify-center">
                  <Award className="w-5 h-5 text-purple-500" />
                </div>
                <div>
                  <div className="text-2xl font-bold text-foreground">84</div>
                  <div className="text-xs text-muted-foreground">Уроков пройдено</div>
                </div>
              </div>
            </div>
          </motion.div>
        </div>
      </section>

      {/* Courses Section */}
      <section className="py-12">
        <div className="container mx-auto px-4 sm:px-6 lg:px-8">
          {/* Tabs */}
          <div className="flex gap-4 mb-8 border-b border-border">
            <button
              onClick={() => setActiveTab('progress')}
              className={`relative px-6 py-3 font-semibold transition-colors ${
                activeTab === 'progress'
                  ? 'text-primary'
                  : 'text-muted-foreground hover:text-foreground'
              }`}
            >
              В процессе
              {activeTab === 'progress' && (
                <motion.div
                  layoutId="activeProfileTab"
                  className="absolute bottom-0 left-0 right-0 h-0.5 bg-primary"
                  transition={{ type: 'spring', stiffness: 500, damping: 30 }}
                />
              )}
            </button>
            <button
              onClick={() => setActiveTab('completed')}
              className={`relative px-6 py-3 font-semibold transition-colors ${
                activeTab === 'completed'
                  ? 'text-primary'
                  : 'text-muted-foreground hover:text-foreground'
              }`}
            >
              Завершённые
              {activeTab === 'completed' && (
                <motion.div
                  layoutId="activeProfileTab"
                  className="absolute bottom-0 left-0 right-0 h-0.5 bg-primary"
                  transition={{ type: 'spring', stiffness: 500, damping: 30 }}
                />
              )}
            </button>
          </div>

          {/* In Progress Tab */}
          {activeTab === 'progress' && (
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
                        style={{ backgroundColor: `var(--subject-${course.subject})` }}
                      >
                        {course.subjectName}
                      </div>
                      <h3 className="text-xl font-semibold text-foreground mb-2">
                        {course.title}
                      </h3>
                      <p className="text-sm text-muted-foreground">
                        Последняя активность: {course.lastActivity}
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
                        {course.completedLessons} из {course.totalLessons} уроков
                      </span>
                      <span className="font-semibold text-foreground">{course.progress}%</span>
                    </div>
                    <Progress.Root
                      className="relative h-3 overflow-hidden rounded-full bg-muted"
                      value={course.progress}
                    >
                      <Progress.Indicator
                        className="h-full bg-accent transition-all duration-300 ease-out rounded-full"
                        style={{
                          width: `${course.progress}%`,
                        }}
                      />
                    </Progress.Root>
                  </div>
                </motion.div>
              ))}
            </motion.div>
          )}

          {/* Completed Tab */}
          {activeTab === 'completed' && (
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
          )}
        </div>
      </section>
    </div>
  );
}

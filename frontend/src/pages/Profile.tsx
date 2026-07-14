import { useState } from 'react';

import { ProfileHeader, Tabs, ProgressTab, CompletedTab } from '@/components/ui/Profile'

const inProgressCourses = [
  {
    id: 1,
    subject: 'russian',
    title: 'Полная подготовка к ЕГЭ',
    percentages: 65,
    totalLessons: 48,
    completedLessons: 31,
    lastActivity: '2 часа назад',
  },
  {
    id: 2,
    subject: 'math',
    title: 'Профильная математика',
    percentages: 42,
    totalLessons: 60,
    completedLessons: 25,
    lastActivity: '1 день назад',
  },
  {
    id: 3,
    subject: 'english',
    title: 'Подготовка к ЕГЭ за 3 месяца',
    percentages: 78,
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
      <ProfileHeader fullname={"Иванов Иван"} mail={"ivan@example.com"} register={"11.02.2026"} noCompletedCourses={inProgressCourses.length} CompletedCourses={completedCourses.length} />

      {/* Courses Section */}
      <section className="py-12">
        <div className="container mx-auto px-4 sm:px-6 lg:px-8">
          <Tabs activeTab={activeTab} setActiveTab={setActiveTab} />

          {activeTab === 'progress' && <ProgressTab inProgressCourses={inProgressCourses} />}

          {activeTab === 'completed' && <CompletedTab completedCourses={completedCourses} />}
        </div>
      </section>
    </div>
  );
}

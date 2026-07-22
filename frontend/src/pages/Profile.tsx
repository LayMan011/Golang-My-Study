import { useEffect, useState } from 'react';
import { useNavigate } from 'react-router';

import {
  ProfileHeader,
  Tabs,
  ProgressTab,
  CompletedTab,
  ProfileCoursesSkeleton
} from '@/components/ui/Profile';
import { Preloader } from "@/components/common";
import type { UserCourseProgress } from '@/types';

type Theme = {
  id: string;
  version: number;
  completed: boolean;
  addition_at: string;
  completed_at: string | null;
  percentages: number;
  theme_id: string;
  user_id: string;
  total_lessons: number;
  completed_lessons: number;
};

type UserProfile = {
  id: number;
  version: number;
  full_name: string;
  phone_number: string | null;
  email: string;
  created_at: string;
};

type ThemeUser = {
  id: number;
  version: number;

  title: string;
  description: string | null;
  createdAt: Date | string;
  subject: string;
  rating: number | null;
  allRatings: number;
  numberOfRatings: number;
  numberOfUsers: number;
  price: number;

  level: string;
  duration: string;
  format: string;

  authorUserId: number;
};

const EmptyCoursesState = () => {
  const navigate = useNavigate();

  return (
    <section className="py-12">
      <div className="container mx-auto px-4 sm:px-6 lg:px-8">
        <div className="bg-card border border-dashed border-border rounded-2xl p-8 text-center">
          <h2 className="text-xl font-semibold text-foreground mb-2">
            Курсы ещё не добавлены
          </h2>
          <p className="text-sm text-muted-foreground mb-4">
            Как только вы начнёте обучение, здесь появятся ваши курсы и
            прогресс.
          </p>
          <button onClick={() => navigate('/courses')} className="px-6 py-3 bg-primary text-primary-foreground rounded-lg hover:bg-primary/90 transition-colors font-medium">
            Найти курс
          </button>
        </div>
      </div>
    </section>
  )
};

const Profile = () => {
  const [activeTab, setActiveTab] = useState<'progress' | 'completed'>('progress');

  const [user, setUser] = useState<UserProfile | null>(null);
  const [courses, setCourses] = useState<UserCourseProgress[]>([]);
  const [loading, setLoading] = useState(true);
  const [error, setError] = useState<string | null>(null);
  const [hasCourses, setHasCourses] = useState<boolean>(false);

  useEffect(() => {
    const userId = localStorage.getItem('user_id');
    const accessToken = localStorage.getItem('access_token');

    if (!userId || !accessToken) {
      setError('Пользователь не авторизован');
      setLoading(false);
      return;
    }

    const fetchData = async () => {
      try {
        setLoading(true);

        const userRes = await fetch(`http://localhost:5050/api/v1/users/${userId}`);
        if (!userRes.ok) {
          throw new Error('Ошибка загрузки профиля');
        }
        const userData: UserProfile = await userRes.json();
        setUser(userData);

        const coursesRes = await fetch(
          `http://localhost:5050/api/v1/themes_user/user/${userId}`,
        );

        if (!coursesRes.ok) {
          if (coursesRes.status === 404) {
            setHasCourses(false);
          } else {
            throw new Error('Ошибка загрузки курсов');
          }
        } else {
          const coursesData: Theme[] = await coursesRes.json();
          const rawCourses = Array.isArray(coursesData) ? coursesData : [coursesData];

          if (rawCourses.length === 0) {
            setHasCourses(false);
          } else {
            setHasCourses(true);
          }

          const themesRes = await fetch(`http://localhost:5050/api/v1/themes`);
          if (!themesRes.ok) {
            throw new Error('Ошибка загрузки тем');
          }
          const themesData: ThemeUser[] = await themesRes.json();

          const merged: UserCourseProgress[] = rawCourses.map((course) => {
            const themeUser = themesData.find(
              (theme) => theme.id === Number(course.theme_id),
            );

            return {
              id: course.id,
              version: course.version,
              completed: course.completed,
              addition_at: course.addition_at,
              completed_at: course.completed_at,
              percentages: course.percentages,
              theme_id: course.theme_id,
              user_id: course.user_id,
              total_lessons: course.total_lessons,
              completed_lessons: course.completed_lessons,

              title: themeUser?.title ?? '',
              subject: themeUser?.subject ?? '',
            };
          });

          setCourses(merged);
        }
      } catch (err) {
        setError(err instanceof Error ? err.message : 'Неизвестная ошибка');
      } finally {
        setLoading(false);
      }
    };

    fetchData();
  }, []);

  const inProgressCourses = courses.filter((c) => !c.completed);
  const completedCourses = courses.filter((c) => c.completed);

  if (loading && !user) {
    return (
      <div className="min-h-screen bg-background flex items-center justify-center">
        <Preloader />
      </div>
    );
  }

  if (error || !user) {
    return (
      <div className="min-h-screen bg-background flex items-center justify-center">
        <span className="text-destructive">
          {error || 'Ошибка загрузки данных'}
        </span>
      </div>
    );
  }

  return (
    <div className="min-h-screen bg-background">
      <ProfileHeader
        fullname={user.full_name}
        mail={user.email}
        register={user.created_at}
        noCompletedCourses={inProgressCourses.length}
        CompletedCourses={completedCourses.length}
      />

      {loading && user && <ProfileCoursesSkeleton />}

      {!loading && (!hasCourses || courses.length === 0) && <EmptyCoursesState />}

      {!loading && hasCourses && courses.length > 0 && (
        <section className="py-12">
          <div className="container mx-auto px-4 sm:px-6 lg:px-8">
            <Tabs activeTab={activeTab} setActiveTab={setActiveTab} />

            {activeTab === 'progress' && (
              <ProgressTab inProgressCourses={inProgressCourses} />
            )}

            {activeTab === 'completed' && (
              <CompletedTab completedCourses={completedCourses} />
            )}
          </div>
        </section>
      )}
    </div>
  );
};

export default Profile;
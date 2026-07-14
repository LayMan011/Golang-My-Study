export interface Course {
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

export interface Course_User {
    id: number,
    subject: string,
    title: string,
    progress: number,
    totalLessons: number,
    completedLessons: number,
    lastActivity: number,
}

export interface inProgressCourses {
    id: number;
    subject: string;
    title: string;
    percentages: number;
    totalLessons: number;
    completedLessons: number;
    lastActivity: string;
}

export interface completedCourses {
    id: number;
    subject: string;
    title: string;
    completedDate: string;
}
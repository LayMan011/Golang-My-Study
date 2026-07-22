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

export type UserCourseProgress = {
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
  
  title: string;
  subject: string;
};
import { motion } from 'motion/react';
import { CheckCircle, TrendingUp, Users, Award } from 'lucide-react';
import {useState, useEffect} from 'react'

const achievements = [
  { icon: Users, value: '12,000+', label: 'Учеников' },
  { icon: Award, value: '95%', label: 'Поступили в вузы' },
  { icon: TrendingUp, value: '87', label: 'Средний балл ЕГЭ' },
  { icon: CheckCircle, value: '500+', label: 'Курсов' },
];

const updateAchievements = async () => {
  try {
    const responseUsers = await fetch('http://localhost:5050/api/v1/users');
    const dataUsers = await responseUsers.json();

    const responseCourses = await fetch('http://localhost:5050/api/v1/themes');
    const dataCourses = await responseCourses.json();

    const updatedAchievements = achievements.map((item) => {
      if (item.label === 'Учеников') {
        return { ...item, value: `${dataUsers.length}+` };
      }
      if (item.label === 'Курсов') {
        return { ...item, value: `${dataCourses.length}+` };
      }
      return item;
    });

    return updatedAchievements;
  } catch (error) {
    console.error('Ошибка обновления:', error);
    return achievements;
  }
};

export const Icons = () => {
  const [achievementsData, setAchievementsData] = useState(achievements);
  const [loading, setLoading] = useState(true);

  useEffect(() => {
    const fetchData = async () => {
      setLoading(true);
      const updated = await updateAchievements();
      setAchievementsData(updated);
      setLoading(false);
    };

    fetchData();
  }, []);

  if (loading) {
    return (
      <section className="py-12 bg-card border-y border-border">
        <div className="container mx-auto px-4 sm:px-6 lg:px-8">
          <div className="grid grid-cols-2 md:grid-cols-4 gap-8">
            {achievementsData.map((_, index) => (
              <motion.div
                key={index}
                initial={{ opacity: 0, y: 20 }}
                whileInView={{ opacity: 1, y: 0 }}
                viewport={{ once: true }}
                transition={{ delay: index * 0.1 }}
                className="text-center"
              >
                <div className="w-8 h-8 mx-auto mb-2 bg-muted animate-pulse rounded" />
                <div className="h-8 w-20 mx-auto bg-muted animate-pulse rounded mb-1" />
                <div className="h-4 w-16 mx-auto bg-muted animate-pulse rounded" />
              </motion.div>
            ))}
          </div>
        </div>
      </section>
    )
  }

  return (
    <section className="py-12 bg-card border-y border-border">
      <div className="container mx-auto px-4 sm:px-6 lg:px-8">
        <div className="grid grid-cols-2 md:grid-cols-4 gap-8">
          {achievementsData.map((item, index) => (
            <motion.div
              key={index}
              initial={{ opacity: 0, y: 20 }}
              whileInView={{ opacity: 1, y: 0 }}
              viewport={{ once: true }}
              transition={{ delay: index * 0.1 }}
              className="text-center"
            >
              <item.icon className="w-8 h-8 text-primary mx-auto mb-2" />
              <div className="text-3xl font-bold text-foreground mb-1">{item.value}</div>
              <div className="text-sm text-muted-foreground">{item.label}</div>
            </motion.div>
          ))}
        </div>
      </div>
    </section>
  );
};
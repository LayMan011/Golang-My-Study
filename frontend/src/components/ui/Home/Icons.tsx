import { motion } from 'motion/react';
import { CheckCircle, TrendingUp, Users, Award } from 'lucide-react';

const achievements = [
  { icon: Users, value: '12,000+', label: 'Учеников' },
  { icon: Award, value: '95%', label: 'Поступили в вузы' },
  { icon: TrendingUp, value: '87', label: 'Средний балл ЕГЭ' },
  { icon: CheckCircle, value: '500+', label: 'Курсов' },
];

export const Icons = () => {
    return (
      <section className="py-12 bg-card border-y border-border">
        <div className="container mx-auto px-4 sm:px-6 lg:px-8">
          <div className="grid grid-cols-2 md:grid-cols-4 gap-8">
            {achievements.map((item, index) => (
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
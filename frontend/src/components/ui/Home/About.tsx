import { motion } from 'motion/react';

const steps = [
  {
    number: 1,
    title: 'Выберите курс',
    description: 'Подберите курс по нужному предмету и уровню сложности',
  },
  {
    number: 2,
    title: 'Учитесь в своем темпе',
    description: 'Проходите уроки, решайте задачи и тесты когда удобно',
  },
  {
    number: 3,
    title: 'Отслеживайте прогресс',
    description: 'Следите за своими результатами и улучшайте слабые места',
  },
  {
    number: 4,
    title: 'Сдайте ЕГЭ на 100 баллов',
    description: 'Подготовьтесь к экзамену на максимальный результат',
  },
];

export const About = () => {
    return (
      <section className="py-20 bg-background">
        <div className="container mx-auto px-4 sm:px-6 lg:px-8">
          <div className="text-center mb-12">
            <h2 className="text-3xl sm:text-4xl font-bold text-foreground mb-4">
              Как это работает
            </h2>
            <p className="text-lg text-muted-foreground max-w-2xl mx-auto">
              Простой процесс от выбора курса до успешной сдачи экзамена
            </p>
          </div>
          
          <div className="grid md:grid-cols-2 lg:grid-cols-4 gap-8">
            {steps.map((step, index) => (
              <motion.div
                key={index}
                initial={{ opacity: 0, y: 20 }}
                whileInView={{ opacity: 1, y: 0 }}
                viewport={{ once: true }}
                transition={{ delay: index * 0.1 }}
                className="relative"
              >
                <div className="bg-card rounded-2xl p-6 shadow-sm hover:shadow-md transition-shadow border border-border">
                  <div className="w-12 h-12 rounded-full bg-primary text-primary-foreground flex items-center justify-center text-xl font-bold mb-4">
                    {step.number}
                  </div>
                  <h3 className="text-xl font-semibold text-foreground mb-2">
                    {step.title}
                  </h3>
                  <p className="text-muted-foreground">{step.description}</p>
                </div>
              </motion.div>
            ))}
          </div>
        </div>
      </section>
    )
}

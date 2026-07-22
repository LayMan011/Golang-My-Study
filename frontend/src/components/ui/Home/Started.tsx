import { motion } from 'motion/react';
import { ArrowRight } from 'lucide-react';
import { Link } from 'react-router';

export const Started = () => {
    return (
      <section className="relative overflow-hidden bg-linear-to-br from-primary/10 via-background to-accent/5 py-20 sm:py-28">
        <div className="container mx-auto px-4 sm:px-6 lg:px-8">
          <div className="max-w-3xl mx-auto text-center">
            <motion.div
              initial={{ opacity: 0, y: 20 }}
              animate={{ opacity: 1, y: 0 }}
              transition={{ duration: 0.5 }}
            >
              <h1 className="text-4xl sm:text-5xl lg:text-6xl font-bold text-foreground mb-6">
                Подготовься к ЕГЭ <br />
                <span className="text-primary">на максимум</span>
              </h1>
              <p className="text-lg sm:text-xl text-muted-foreground mb-8">
                Более 100 курсов от лучших преподавателей. Учись в своем темпе и достигай высоких результатов
              </p>
              <div className="flex flex-col sm:flex-row gap-4 justify-center">
                <Link
                  to="/courses"
                  className="inline-flex items-center justify-center gap-2 px-8 py-4 rounded-lg bg-accent text-accent-foreground font-semibold hover:bg-accent/90 transition-all hover:scale-105 shadow-lg hover:shadow-xl"
                >
                  Начать бесплатно
                  <ArrowRight className="w-5 h-5" />
                </Link>
                <button className="inline-flex items-center justify-center gap-2 px-8 py-4 rounded-lg bg-card border-2 border-border text-foreground font-semibold hover:bg-muted transition-colors">
                  Смотреть курсы
                </button>
              </div>
            </motion.div>
          </div>
        </div>
        
        {/* Decorative elements */}
        <div className="absolute top-20 left-10 w-20 h-20 bg-primary/10 rounded-full blur-3xl" />
        <div className="absolute bottom-20 right-10 w-32 h-32 bg-accent/10 rounded-full blur-3xl" />
      </section>
    )
}
import { Moon, Sun } from 'lucide-react';
import { useTheme } from '@/hooks';
import { motion } from 'motion/react';

export function ThemeToggle() {
  const { theme, toggleTheme } = useTheme();

  return (
    <motion.button
      onClick={toggleTheme}
      className="relative w-14 h-8 rounded-full bg-muted flex items-center justify-between px-1.5 hover:bg-muted/80 transition-colors"
      whileTap={{ scale: 0.95 }}
      aria-label="Toggle theme"
    >
      <Sun className="w-4 h-4 text-accent" />
      <Moon className="w-4 h-4 text-primary" />
      <motion.div
        className="absolute w-6 h-6 bg-card rounded-full shadow-md"
        initial={false}
        animate={{
          x: theme === 'light' ? 0 : 24,
        }}
        transition={{ type: 'spring', stiffness: 500, damping: 30 }}
      />
    </motion.button>
  );
}
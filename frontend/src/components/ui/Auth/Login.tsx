import { motion } from 'motion/react';
import { Mail, Lock, Eye, EyeOff } from 'lucide-react';

type LoginProps = {
  handleLogin: (e: React.FormEvent<HTMLFormElement>) => Promise<void>;
  email: string;
  setEmail: React.Dispatch<React.SetStateAction<string>>;
  emailPassword: string;
  setEmailPassword: React.Dispatch<React.SetStateAction<string>>;
  showPassword: boolean;
  setShowPassword: React.Dispatch<React.SetStateAction<boolean>>;
  error?: string;
  loading?: boolean;
};

export const Login = ({
  handleLogin,
  email,
  setEmail,
  emailPassword,
  setEmailPassword,
  showPassword,
  setShowPassword,
  loading = false,
}: LoginProps) => {
  return (
    <motion.form
      key="login"
      initial={{ opacity: 0, x: -20 }}
      animate={{ opacity: 1, x: 0 }}
      exit={{ opacity: 0, x: 20 }}
      onSubmit={handleLogin}
      className="space-y-6"
    >
      <div>
        <h2 className="text-2xl font-bold text-foreground mb-2">
          С возвращением!
        </h2>
        <p className="text-muted-foreground">
          Войдите, чтобы продолжить обучение
        </p>
      </div>

      <div>
        <label htmlFor="email" className="block text-sm font-medium text-foreground mb-2">
          Почта
        </label>
        <div className="relative">
          <Mail className="absolute left-3 top-1/2 -translate-y-1/2 w-5 h-5 text-muted-foreground" />
          <input
            id="email"
            type="text"
            value={email}
            onChange={(e) => setEmail(e.target.value)}
            placeholder="your_email"
            className="w-full pl-10 pr-4 py-3 bg-background border border-border rounded-lg text-foreground placeholder:text-muted-foreground focus:border-primary focus:outline-none transition-colors"
            required
          />
        </div>
      </div>

      <div>
        <label htmlFor="email-password" className="block text-sm font-medium text-foreground mb-2">
          Пароль
        </label>
        <div className="relative">
          <Lock className="absolute left-3 top-1/2 -translate-y-1/2 w-5 h-5 text-muted-foreground" />
          <input
            id="email-password"
            type={showPassword ? 'text' : 'password'}
            value={emailPassword}
            onChange={(e) => setEmailPassword(e.target.value)}
            placeholder="••••••••"
            className="w-full pl-10 pr-12 py-3 bg-background border border-border rounded-lg text-foreground placeholder:text-muted-foreground focus:border-primary focus:outline-none transition-colors"
            required
          />
          <button
            type="button"
            onClick={() => setShowPassword(!showPassword)}
            className="absolute right-3 top-1/2 -translate-y-1/2 text-muted-foreground hover:text-foreground"
          >
            {showPassword ? <EyeOff className="w-5 h-5" /> : <Eye className="w-5 h-5" />}
          </button>
        </div>
      </div>

      <div className="flex items-center justify-between text-sm">
        <label className="flex items-center gap-2 cursor-pointer">
          <input type="checkbox" className="w-4 h-4 rounded accent-primary" />
          <span className="text-muted-foreground">Запомнить меня</span>
        </label>
        <a href="#" className="text-primary hover:underline">
          Забыли пароль?
        </a>
      </div>

      <button
        type="submit"
        disabled={loading}
        className="w-full py-3 bg-accent text-accent-foreground font-semibold rounded-lg hover:bg-accent/90 transition-colors shadow-md hover:shadow-lg disabled:opacity-60"
      >
        {loading ? 'Вход...' : 'Войти'}
      </button>
    </motion.form>
  );
};
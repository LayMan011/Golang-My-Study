import { motion } from 'motion/react';
import { Mail, Lock, Eye, EyeOff } from 'lucide-react';

export const Login = ({
    handleLogin,
    loginEmail,
    setLoginEmail,
    loginPassword,
    setLoginPassword,
    showPassword,
    setShowPassword,
}: {
    handleLogin: (e: React.FormEvent<Element>) => void,
    loginEmail: string,
    setLoginEmail: React.Dispatch<React.SetStateAction<string>>,
    loginPassword: string,
    setLoginPassword: React.Dispatch<React.SetStateAction<string>>,
    showPassword: boolean,
    setShowPassword: React.Dispatch<React.SetStateAction<boolean>>,
}) => {
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

            {/* Email */}
            <div>
                <label htmlFor="login-email" className="block text-sm font-medium text-foreground mb-2">
                Email
                </label>
                <div className="relative">
                <Mail className="absolute left-3 top-1/2 -translate-y-1/2 w-5 h-5 text-muted-foreground" />
                <input
                    id="login-email"
                    type="email"
                    value={loginEmail}
                    onChange={(e) => setLoginEmail(e.target.value)}
                    placeholder="your@email.com"
                    className="w-full pl-10 pr-4 py-3 bg-background border border-border rounded-lg text-foreground placeholder:text-muted-foreground focus:border-primary focus:outline-none transition-colors"
                    required
                />
                </div>
            </div>

            {/* Password */}
            <div>
                <label htmlFor="login-password" className="block text-sm font-medium text-foreground mb-2">
                Пароль
                </label>
                <div className="relative">
                <Lock className="absolute left-3 top-1/2 -translate-y-1/2 w-5 h-5 text-muted-foreground" />
                <input
                    id="login-password"
                    type={showPassword ? 'text' : 'password'}
                    value={loginPassword}
                    onChange={(e) => setLoginPassword(e.target.value)}
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
                <input
                    type="checkbox"
                    className="w-4 h-4 rounded accent-primary"
                />
                <span className="text-muted-foreground">Запомнить меня</span>
                </label>
                <a href="#" className="text-primary hover:underline">
                Забыли пароль?
                </a>
            </div>

            <button
                type="submit"
                className="w-full py-3 bg-accent text-accent-foreground font-semibold rounded-lg hover:bg-accent/90 transition-colors shadow-md hover:shadow-lg"
            >
                Войти
            </button>
        </motion.form>
    )
}
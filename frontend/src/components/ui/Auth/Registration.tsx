import { motion } from 'motion/react';
import { Mail, Lock, User, Calendar, CheckCircle, XCircle, Eye, EyeOff } from 'lucide-react';

export const Registration = ({
    handleRegister,

    fullName,
    setFullName,

    registerEmail,
    setRegisterEmail,

    emailValid,
    validateEmail,

    birthDate,
    setBirthDate,

    registerPassword,
    setRegisterPassword,

    passwordValid,
    validatePassword,

    showPassword,
    setShowPassword,

    confirmPassword,
    setConfirmPassword,

    showConfirmPassword,
    setShowConfirmPassword,

    passwordMatch,
    checkPasswordMatch,
}: {
    handleRegister: (e: React.FormEvent<Element>) => void,

    fullName: string,
    setFullName: React.Dispatch<React.SetStateAction<string>>,

    registerEmail: string,
    setRegisterEmail: React.Dispatch<React.SetStateAction<string>>,

    emailValid: boolean | null,
    validateEmail: (email: string) => boolean,

    birthDate: string,
    setBirthDate: React.Dispatch<React.SetStateAction<string>>,

    registerPassword: string,
    setRegisterPassword: React.Dispatch<React.SetStateAction<string>>,

    passwordValid: boolean | null,
    validatePassword: (password: string) => boolean,

    showPassword: boolean,
    setShowPassword: React.Dispatch<React.SetStateAction<boolean>>,

    confirmPassword: string,
    setConfirmPassword: React.Dispatch<React.SetStateAction<string>>,

    showConfirmPassword: boolean,
    setShowConfirmPassword: React.Dispatch<React.SetStateAction<boolean>>,

    passwordMatch: boolean | null,
    checkPasswordMatch: (pass: string, confirm: string) => boolean,
}) => {
    return (
        <motion.form
        key="register"
        initial={{ opacity: 0, x: 20 }}
        animate={{ opacity: 1, x: 0 }}
        exit={{ opacity: 0, x: -20 }}
        transition={{ duration: 0.3 }}
        onSubmit={handleRegister}
        className="space-y-6"
        >
        <div>
            <h2 className="text-2xl font-bold text-foreground mb-2">
            Создать аккаунт
            </h2>
            <p className="text-muted-foreground">
            Присоединяйтесь к тысячам учеников
            </p>
        </div>

        {/* Full Name */}
        <div>
            <label htmlFor="full-name" className="block text-sm font-medium text-foreground mb-2">
            Полное имя
            </label>
            <div className="relative">
            <User className="absolute left-3 top-1/2 -translate-y-1/2 w-5 h-5 text-muted-foreground" />
            <input
                id="full-name"
                type="text"
                value={fullName}
                onChange={(e) => setFullName(e.target.value)}
                placeholder="Иван Иванов"
                className="w-full pl-10 pr-4 py-3 bg-background border border-border rounded-lg text-foreground placeholder:text-muted-foreground focus:border-primary focus:outline-none transition-colors"
                required
            />
            </div>
        </div>

        {/* Email */}
        <div>
            <label htmlFor="register-email" className="block text-sm font-medium text-foreground mb-2">
            Email
            </label>
            <div className="relative">
            <Mail className="absolute left-3 top-1/2 -translate-y-1/2 w-5 h-5 text-muted-foreground" />
            <input
                id="register-email"
                type="email"
                value={registerEmail}
                onChange={(e) => {
                setRegisterEmail(e.target.value);
                validateEmail(e.target.value);
                }}
                placeholder="your@email.com"
                className={`w-full pl-10 pr-12 py-3 bg-background border rounded-lg text-foreground placeholder:text-muted-foreground focus:outline-none transition-colors ${
                emailValid === null
                    ? 'border-border focus:border-primary'
                    : emailValid
                    ? 'border-green-500 focus:border-green-500'
                    : 'border-red-500 focus:border-red-500'
                }`}
                required
            />
            {emailValid !== null && (
                <div className="absolute right-3 top-1/2 -translate-y-1/2">
                {emailValid ? (
                    <CheckCircle className="w-5 h-5 text-green-500" />
                ) : (
                    <XCircle className="w-5 h-5 text-red-500" />
                )}
                </div>
            )}
            </div>
            {emailValid === false && (
            <p className="text-xs text-red-500 mt-1">Введите корректный email</p>
            )}
        </div>

        {/* Birth Date */}
        <div>
            <label htmlFor="birth-date" className="block text-sm font-medium text-foreground mb-2">
            Дата рождения
            </label>
            <div className="relative">
            <Calendar className="absolute left-3 top-1/2 -translate-y-1/2 w-5 h-5 text-muted-foreground" />
            <input
                id="birth-date"
                type="date"
                value={birthDate}
                onChange={(e) => setBirthDate(e.target.value)}
                className="w-full pl-10 pr-4 py-3 bg-background border border-border rounded-lg text-foreground focus:border-primary focus:outline-none transition-colors"
                required
            />
            </div>
        </div>

        {/* Password */}
        <div>
            <label htmlFor="register-password" className="block text-sm font-medium text-foreground mb-2">
            Пароль
            </label>
            <div className="relative">
            <Lock className="absolute left-3 top-1/2 -translate-y-1/2 w-5 h-5 text-muted-foreground" />
            <input
                id="register-password"
                type={showPassword ? 'text' : 'password'}
                value={registerPassword}
                onChange={(e) => {
                setRegisterPassword(e.target.value);
                validatePassword(e.target.value);
                if (confirmPassword) {
                    checkPasswordMatch(e.target.value, confirmPassword);
                }
                }}
                placeholder="••••••••"
                className={`w-full pl-10 pr-12 py-3 bg-background border rounded-lg text-foreground placeholder:text-muted-foreground focus:outline-none transition-colors ${
                passwordValid === null
                    ? 'border-border focus:border-primary'
                    : passwordValid
                    ? 'border-green-500 focus:border-green-500'
                    : 'border-red-500 focus:border-red-500'
                }`}
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
            {passwordValid === false && (
            <p className="text-xs text-red-500 mt-1">Минимум 8 символов</p>
            )}
        </div>

        {/* Confirm Password */}
        <div>
            <label htmlFor="confirm-password" className="block text-sm font-medium text-foreground mb-2">
            Повторите пароль
            </label>
            <div className="relative">
            <Lock className="absolute left-3 top-1/2 -translate-y-1/2 w-5 h-5 text-muted-foreground" />
            <input
                id="confirm-password"
                type={showConfirmPassword ? 'text' : 'password'}
                value={confirmPassword}
                onChange={(e) => {
                setConfirmPassword(e.target.value);
                checkPasswordMatch(registerPassword, e.target.value);
                }}
                placeholder="••••••••"
                className={`w-full pl-10 pr-12 py-3 bg-background border rounded-lg text-foreground placeholder:text-muted-foreground focus:outline-none transition-colors ${
                passwordMatch === null
                    ? 'border-border focus:border-primary'
                    : passwordMatch
                    ? 'border-green-500 focus:border-green-500'
                    : 'border-red-500 focus:border-red-500'
                }`}
                required
            />
            <button
                type="button"
                onClick={() => setShowConfirmPassword(!showConfirmPassword)}
                className="absolute right-3 top-1/2 -translate-y-1/2 text-muted-foreground hover:text-foreground"
            >
                {showConfirmPassword ? <EyeOff className="w-5 h-5" /> : <Eye className="w-5 h-5" />}
            </button>
            </div>
            {passwordMatch === false && (
            <p className="text-xs text-red-500 mt-1">Пароли не совпадают</p>
            )}
        </div>

        <button
            type="submit"
            className="w-full py-3 bg-accent text-accent-foreground font-semibold rounded-lg hover:bg-accent/90 transition-colors shadow-md hover:shadow-lg"
        >
            Зарегистрироваться
        </button>
        </motion.form>
    )
}
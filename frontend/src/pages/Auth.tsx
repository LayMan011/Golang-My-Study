import { useState } from 'react';
import { motion } from 'motion/react';
import { Mail, Lock, User, Calendar, CheckCircle, XCircle, Eye, EyeOff } from 'lucide-react';

export const Auth = () => {
  const [activeTab, setActiveTab] = useState<'login' | 'register'>('login');
  const [showPassword, setShowPassword] = useState(false);
  const [showConfirmPassword, setShowConfirmPassword] = useState(false);

  // Login state
  const [loginEmail, setLoginEmail] = useState('');
  const [loginPassword, setLoginPassword] = useState('');

  // Register state
  const [registerEmail, setRegisterEmail] = useState('');
  const [registerPassword, setRegisterPassword] = useState('');
  const [confirmPassword, setConfirmPassword] = useState('');
  const [birthDate, setBirthDate] = useState('');
  const [fullName, setFullName] = useState('');

  // Validation states
  const [emailValid, setEmailValid] = useState<boolean | null>(null);
  const [passwordValid, setPasswordValid] = useState<boolean | null>(null);
  const [passwordMatch, setPasswordMatch] = useState<boolean | null>(null);

  const validateEmail = (email: string) => {
    const regex = /^[^\s@]+@[^\s@]+\.[^\s@]+$/;
    const isValid = regex.test(email);
    setEmailValid(email.length > 0 ? isValid : null);
    return isValid;
  };

  const validatePassword = (password: string) => {
    const isValid = password.length >= 8;
    setPasswordValid(password.length > 0 ? isValid : null);
    return isValid;
  };

  const checkPasswordMatch = (pass: string, confirm: string) => {
    const isMatch = pass === confirm && pass.length > 0;
    setPasswordMatch(confirm.length > 0 ? isMatch : null);
    return isMatch;
  };

  const handleLogin = (e: React.FormEvent) => {
    e.preventDefault();
    console.log('Login:', { email: loginEmail, password: loginPassword });
  };

  const handleRegister = (e: React.FormEvent) => {
    e.preventDefault();
    console.log('Register:', {
      email: registerEmail,
      password: registerPassword,
      fullName,
      birthDate,
    });
  };

  return (
    <div className="min-h-screen bg-background flex items-center justify-center py-12 px-4 sm:px-6 lg:px-8">
      <div className="w-full max-w-md">
        <motion.div
          initial={{ opacity: 0, y: 20 }}
          animate={{ opacity: 1, y: 0 }}
          transition={{ duration: 0.5 }}
          className="bg-card rounded-2xl shadow-xl border border-border overflow-hidden"
        >
          {/* Tabs */}
          <div className="flex border-b border-border">
            <button
              onClick={() => setActiveTab('login')}
              className={`flex-1 py-4 font-semibold text-center transition-colors relative ${
                activeTab === 'login'
                  ? 'text-primary'
                  : 'text-muted-foreground hover:text-foreground'
              }`}
            >
              Вход
              {activeTab === 'login' && (
                <motion.div
                  layoutId="activeAuthTab"
                  className="absolute bottom-0 left-0 right-0 h-0.5 bg-primary"
                  transition={{ type: 'spring', stiffness: 500, damping: 30 }}
                />
              )}
            </button>
            <button
              onClick={() => setActiveTab('register')}
              className={`flex-1 py-4 font-semibold text-center transition-colors relative ${
                activeTab === 'register'
                  ? 'text-primary'
                  : 'text-muted-foreground hover:text-foreground'
              }`}
            >
              Регистрация
              {activeTab === 'register' && (
                <motion.div
                  layoutId="activeAuthTab"
                  className="absolute bottom-0 left-0 right-0 h-0.5 bg-primary"
                  transition={{ type: 'spring', stiffness: 500, damping: 30 }}
                />
              )}
            </button>
          </div>

          {/* Content */}
          <div className="p-8">
            {activeTab === 'login' ? (
              <motion.form
                key="login"
                initial={{ opacity: 0, x: -20 }}
                animate={{ opacity: 1, x: 0 }}
                exit={{ opacity: 0, x: 20 }}
                transition={{ duration: 0.3 }}
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
            ) : (
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
            )}

            {/* Social Auth */}
            <div className="mt-6">
              <div className="relative">
                <div className="absolute inset-0 flex items-center">
                  <div className="w-full border-t border-border"></div>
                </div>
                <div className="relative flex justify-center text-sm">
                  <span className="px-2 bg-card text-muted-foreground">Или продолжите с</span>
                </div>
              </div>

              <div className="mt-6 grid grid-cols-2 gap-3">
                <button className="flex items-center justify-center gap-2 px-4 py-2 bg-background border border-border rounded-lg text-foreground hover:bg-muted transition-colors">
                  <svg className="w-5 h-5" viewBox="0 0 24 24">
                    <path fill="currentColor" d="M22.56 12.25c0-.78-.07-1.53-.2-2.25H12v4.26h5.92c-.26 1.37-1.04 2.53-2.21 3.31v2.77h3.57c2.08-1.92 3.28-4.74 3.28-8.09z"/>
                    <path fill="currentColor" d="M12 23c2.97 0 5.46-.98 7.28-2.66l-3.57-2.77c-.98.66-2.23 1.06-3.71 1.06-2.86 0-5.29-1.93-6.16-4.53H2.18v2.84C3.99 20.53 7.7 23 12 23z"/>
                    <path fill="currentColor" d="M5.84 14.09c-.22-.66-.35-1.36-.35-2.09s.13-1.43.35-2.09V7.07H2.18C1.43 8.55 1 10.22 1 12s.43 3.45 1.18 4.93l2.85-2.22.81-.62z"/>
                    <path fill="currentColor" d="M12 5.38c1.62 0 3.06.56 4.21 1.64l3.15-3.15C17.45 2.09 14.97 1 12 1 7.7 1 3.99 3.47 2.18 7.07l3.66 2.84c.87-2.6 3.3-4.53 6.16-4.53z"/>
                  </svg>
                  Google
                </button>
                <button className="flex items-center justify-center gap-2 px-4 py-2 bg-background border border-border rounded-lg text-foreground hover:bg-muted transition-colors">
                  <svg className="w-5 h-5" fill="currentColor" viewBox="0 0 24 24">
                    <path d="M12 0c-6.626 0-12 5.373-12 12 0 5.302 3.438 9.8 8.207 11.387.599.111.793-.261.793-.577v-2.234c-3.338.726-4.033-1.416-4.033-1.416-.546-1.387-1.333-1.756-1.333-1.756-1.089-.745.083-.729.083-.729 1.205.084 1.839 1.237 1.839 1.237 1.07 1.834 2.807 1.304 3.492.997.107-.775.418-1.305.762-1.604-2.665-.305-5.467-1.334-5.467-5.931 0-1.311.469-2.381 1.236-3.221-.124-.303-.535-1.524.117-3.176 0 0 1.008-.322 3.301 1.23.957-.266 1.983-.399 3.003-.404 1.02.005 2.047.138 3.006.404 2.291-1.552 3.297-1.23 3.297-1.23.653 1.653.242 2.874.118 3.176.77.84 1.235 1.911 1.235 3.221 0 4.609-2.807 5.624-5.479 5.921.43.372.823 1.102.823 2.222v3.293c0 .319.192.694.801.576 4.765-1.589 8.199-6.086 8.199-11.386 0-6.627-5.373-12-12-12z"/>
                  </svg>
                  GitHub
                </button>
              </div>
            </div>
          </div>
        </motion.div>
      </div>
    </div>
  );
}

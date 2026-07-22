import toast from 'react-hot-toast';
import { useEffect, useState } from 'react';
import { motion, AnimatePresence } from 'motion/react';
import { useNavigate } from 'react-router-dom';

import { Tabs, Login, Registration } from '@/components/ui/Auth';
import { useAuthModal, useAuth } from '@/hooks';

type LoginResponse = {
  user_id: string;
  access_token: string;
  refresh_token: string;
};

type Err = {
  error: string;
  message: string;
}

interface ErrorWithDetails extends Error {
  statusCode?: number;
  errorData?: Err;
  raw?: string;
}

const toasterError = (errorMessage: string) => {
  toast.error(`${errorMessage}`, {
    style: {
      background: '#FEE2E2',
      color: '#991B1B',
      border: '1px solid #FCA5A5',
    },
  })
}

const filtersForError = (errorMessage: string) => {
  if (errorMessage === "failed to decode and validate HTTP request") return "Некорректный ввод данных"
  if (errorMessage === "incorrect password entered") return "Неверный пароль"
  return errorMessage
}

export const AuthModal = () => {
  const { isOpen, close } = useAuthModal();
  const { login: authLogin } = useAuth();
  const navigate = useNavigate();

  const [activeTab, setActiveTab] = useState<'login' | 'register'>('login');
  const [showPassword, setShowPassword] = useState(false);
  const [showConfirmPassword, setShowConfirmPassword] = useState(false);

  const [email, setEmail] = useState('');
  const [emailPassword, setEmailPassword] = useState('');

  const [registerEmail, setRegisterEmail] = useState('');
  const [registerPassword, setRegisterPassword] = useState('');
  const [confirmPassword, setConfirmPassword] = useState('');
  const [fullName, setFullName] = useState('');

  const [emailValid, setEmailValid] = useState<boolean | null>(null);
  const [passwordValid, setPasswordValid] = useState<boolean | null>(null);
  const [passwordMatch, setPasswordMatch] = useState<boolean | null>(null);

  const [loading, setLoading] = useState(false);

  useEffect(() => {
    if (!isOpen) return;

    const { overflow } = document.body.style;
    document.body.style.overflow = 'hidden';

    return () => {
      document.body.style.overflow = overflow || '';
    };
  }, [isOpen]);

  const validateEmail = (value: string) => {
    const trimmed = value.trim();
    if (trimmed.length === 0) {
      setEmailValid(null);
      return false;
    }
    const emailPattern = /^[^\s@]+@[^\s@]+\.[^\s@]+$/;
    const isValid = emailPattern.test(trimmed);
    setEmailValid(isValid);
    return isValid;
  };

  const validatePassword = (password: string) => {
    const isValid = password.length >= 8;
    setPasswordValid(password.length > 0 ? isValid : null);
    return isValid;
  };

  const validateFullname = (fullname: string) => {
    const isValid = fullname.length >= 3;
    return isValid;
  };

  const checkPasswordMatch = (pass: string, confirm: string) => {
    const isMatch = pass === confirm && pass.length > 0;
    setPasswordMatch(confirm.length > 0 ? isMatch : null);
    return isMatch;
  };

  const clearInputs = () => {
    setEmail('');
    setEmailPassword('');
    setRegisterEmail('');
    setRegisterPassword('');
    setConfirmPassword('');
    setFullName('');

    setEmailValid(null);
    setPasswordValid(null);
    setPasswordMatch(null);
  }

  const handleLogin = async (e: React.FormEvent<HTMLFormElement>) => {
    e.preventDefault();
    setLoading(true);

    try {
      if (!validateEmail(email)) {
        toasterError('Введите корректный email');
        setLoading(false);
        return;
      }

      if (!validatePassword(emailPassword)) {
        toasterError('Пароль должен содержать минимум 8 символов');
        setLoading(false);
        return;
      }

      const response = await fetch(
        'http://localhost:5050/api/v1/users/login',
        {
          method: 'POST',
          headers: {
            'Content-Type': 'application/json',
          },
          body: JSON.stringify({
            email: email,
            password: emailPassword,
          }),
        },
      );

      const raw = await response.text();

      if (!response.ok) {
        let errorMsg = `Ошибка входа: ${response.status}`;
        let errorData: Err | undefined;

        try {
          errorData = JSON.parse(raw) as Err;
          errorMsg = errorData.message || errorData.error || raw;
        } catch {
          errorMsg = raw || errorMsg;
        }

        const error: ErrorWithDetails  = new Error(errorMsg);
        error.statusCode = response.status;
        error.errorData = errorData;
        error.raw = raw;
        
        throw error;
      }

      const data: LoginResponse = JSON.parse(raw);

      authLogin(
        data.user_id,
        data.access_token,
        data.refresh_token,
      );

      clearInputs();

      close();
      navigate('/', { replace: true });
    } catch (err) {
        let errorMessage = 'Неизвестная ошибка';

        if (err instanceof Error) {
          errorMessage = err.message;
        }

        toasterError(filtersForError(errorMessage));
    } finally {
      setLoading(false);
    }
  };

  const handleRegister = async (e: React.FormEvent) => {
    e.preventDefault();
    setLoading(true);

    try {
      if (!validateFullname(fullName)) {
        toasterError('Имя пользователя должно содержать минимум 3 символа');
        setLoading(false);
        return;
      }

      if (!validateEmail(registerEmail)) {
        toasterError('Введите корректный email');
        setLoading(false);
        return;
      }

      if (!validatePassword(registerPassword)) {
        toasterError('Пароль должен содержать минимум 8 символов');
        setLoading(false);
        return;
      }

      if (!checkPasswordMatch(registerPassword, confirmPassword)) {
        toasterError('Пароли не совпадают');
        setLoading(false);
        return;
      }

      const response = await fetch(
        'http://localhost:5050/api/v1/users',
        {
          method: 'POST',
          headers: {
            'Content-Type': 'application/json',
          },
          body: JSON.stringify({
            email: registerEmail,
            password: registerPassword,
            full_name: fullName,
          }),
        },
      );

      const raw = await response.text();

      if (!response.ok) {
        let errorMsg = `Ошибка входа: ${response.status}`;
        let errorData: Err | undefined;

        try {
          errorData = JSON.parse(raw) as Err;
          errorMsg = errorData.message || errorData.error || raw;
        } catch {
          errorMsg = raw || errorMsg;
        }

        const error: ErrorWithDetails  = new Error(errorMsg);
        error.statusCode = response.status;
        error.errorData = errorData;
        error.raw = raw;
        
        throw error;
      }

      const loginRes = await fetch(
        'http://localhost:5050/api/v1/users/login',
        {
          method: 'POST',
          headers: {
            'Content-Type': 'application/json',
          },
          body: JSON.stringify({
            email: registerEmail,
            password: registerPassword,
          }),
        },
      );

      const loginRaw = await loginRes.text();

      if (!loginRes.ok) {
        let errorMsg = `Ошибка входа: ${response.status}`;
        let errorData: Err | undefined;

        try {
          errorData = JSON.parse(raw) as Err;
          errorMsg = errorData.message || errorData.error || raw;
        } catch {
          errorMsg = raw || errorMsg;
        }

        const error: ErrorWithDetails  = new Error(errorMsg);
        error.statusCode = response.status;
        error.errorData = errorData;
        error.raw = raw;
        
        throw error;
      }

      const loginData: LoginResponse = JSON.parse(loginRaw);

      authLogin(
        loginData.user_id,
        loginData.access_token,
        loginData.refresh_token,
      );

      clearInputs();

      close();
      navigate('/', { replace: true });
    } catch (err) {
        let errorMessage = 'Неизвестная ошибка';

        if (err instanceof Error) {
          errorMessage = err.message;
        }

        toasterError(filtersForError(errorMessage));
    } finally {
      setLoading(false);
    }
  };

  return (
    <AnimatePresence>
      {isOpen && (
        <motion.div
          className="fixed inset-0 z-60 flex items-center justify-center bg-black/50 backdrop-blur-sm"
          initial={{ opacity: 0 }}
          animate={{ opacity: 1 }}
          exit={{ opacity: 0 }}
          onClick={close}
        >
          <motion.div
            initial={{ opacity: 0, scale: 0.96, y: 8 }}
            animate={{ opacity: 1, scale: 1, y: 0 }}
            exit={{ opacity: 0, scale: 0.96, y: 8 }}
            transition={{ type: 'spring', stiffness: 260, damping: 24 }}
            className="w-full max-w-md bg-card rounded-2xl shadow-xl border border-border overflow-hidden"
            onClick={(e) => e.stopPropagation()}
          >
            <Tabs activeTab={activeTab} setActiveTab={setActiveTab} />

            <div className="p-8">
              {activeTab === 'login' ? (
                <Login
                  handleLogin={handleLogin}
                  email={email}
                  setEmail={setEmail}
                  emailPassword={emailPassword}
                  setEmailPassword={setEmailPassword}
                  showPassword={showPassword}
                  setShowPassword={setShowPassword}
                  loading={loading}
                />
              ) : (
                <Registration
                  handleRegister={handleRegister}
                  fullName={fullName}
                  setFullName={setFullName}
                  registerEmail={registerEmail}
                  setRegisterEmail={setRegisterEmail}
                  emailValid={emailValid}
                  validateEmail={validateEmail}
                  registerPassword={registerPassword}
                  setRegisterPassword={setRegisterPassword}
                  passwordValid={passwordValid}
                  validatePassword={validatePassword}
                  showPassword={showPassword}
                  setShowPassword={setShowPassword}
                  confirmPassword={confirmPassword}
                  setConfirmPassword={setConfirmPassword}
                  showConfirmPassword={showConfirmPassword}
                  setShowConfirmPassword={setShowConfirmPassword}
                  passwordMatch={passwordMatch}
                  checkPasswordMatch={checkPasswordMatch}
                />
              )}
            </div>
          </motion.div>
        </motion.div>
      )}
    </AnimatePresence>
  );
};
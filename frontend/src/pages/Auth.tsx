import { useState } from 'react';
import { motion } from 'motion/react';

import { Tabs, Login, Registration, Social } from '@/components/ui/Auth'

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
          <Tabs activeTab={activeTab} setActiveTab={setActiveTab} />

          <div className="p-8">
            {activeTab === 'login' ? (
              <Login
                handleLogin={handleLogin}
                loginEmail={loginEmail}
                setLoginEmail={setLoginEmail}
                loginPassword={loginPassword}
                setLoginPassword={setLoginPassword}
                showPassword={showPassword}
                setShowPassword={setShowPassword}
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
                birthDate={birthDate}
                setBirthDate={setBirthDate}
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

            <Social />
          </div>
        </motion.div>
      </div>
    </div>
  );
}

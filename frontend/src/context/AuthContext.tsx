import {
  createContext,
  useState,
  useEffect,
} from 'react';
import toast from 'react-hot-toast';
import type { ReactNode } from 'react';

type AuthContextType = {
  user_id: string | null;
  isAuthenticated: boolean;
  login: (userID: string, access_token: string, refresh_token: string) => void;
  logout: () => void;
};

export const AuthContext = createContext<AuthContextType | undefined>(undefined);

export function AuthProvider({ children }: { children: ReactNode }) {
  const [user_id, setUserID] = useState<string | null>(null);
  const [access_token, setAccess_token] = useState<string | null>(null);
  const [refresh_token, setRefresh_token] = useState<string | null>(null);

  useEffect(() => {
    const savedToken = localStorage.getItem('token');
    const savedUserID = localStorage.getItem('user_id');

    if (savedToken && savedUserID) {
      setUserID(savedUserID);
      setAccess_token(access_token);
      setRefresh_token(refresh_token);
    }
  }, []);

  const login = (userID: string, access_token: string, refresh_token: string) => {
    setUserID(userID);
    setAccess_token(access_token);
    setRefresh_token(refresh_token);

    localStorage.setItem('user_id', userID);
    localStorage.setItem('access_token', access_token);
    localStorage.setItem('refresh_token', refresh_token);

    toast.success("Успешный вход")
  };

  const logout = () => {
    setUserID(null);
    setAccess_token(null);
    setRefresh_token(null);

    localStorage.removeItem('user_id');
    localStorage.removeItem('access_token');
    localStorage.removeItem('refresh_token');

    toast.success('Вы вышли из аккаунта');
  };

  return (
    <AuthContext.Provider
      value={{
        user_id,
        isAuthenticated: !!user_id && !!access_token && !!refresh_token,
        login,
        logout,
      }}
    >
      {children}
    </AuthContext.Provider>
  );
}
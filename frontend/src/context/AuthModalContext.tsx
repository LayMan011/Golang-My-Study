import {
  createContext,
  useState,
} from 'react';
import type { ReactNode } from 'react';

type AuthModalContextType = {
  isOpen: boolean;
  open: () => void;
  close: () => void;
};

export const AuthModalContext = createContext<AuthModalContextType | undefined>(
  undefined,
);

export function AuthModalProvider({ children }: { children: ReactNode }) {
  const [isOpen, setIsOpen] = useState(false);

  const open = () => setIsOpen(true);
  const close = () => setIsOpen(false);

  return (
    <AuthModalContext.Provider value={{ isOpen, open, close }}>
      {children}
    </AuthModalContext.Provider>
  );
}
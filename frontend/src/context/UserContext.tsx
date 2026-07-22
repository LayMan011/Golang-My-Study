import { createContext } from "react";
import type { User } from "@/types";

export interface UserContextType {
    user: User | undefined;
    isAuthenticated: boolean;
    isLoadingUserData: boolean;
    isErrorUserData: boolean;
}

export const UserContext = createContext<UserContextType | null>(null);
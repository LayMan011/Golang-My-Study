import { UserContext } from "@/context";
import { useContext } from "react";

export const useUser = () => {
    const context = useContext(UserContext);
    if (!context) throw new Error("Context can be used only in UserProvider");
    return context;
};
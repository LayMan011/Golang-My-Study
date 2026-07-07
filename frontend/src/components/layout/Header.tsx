import { Link, NavLink } from "react-router-dom";
import { Logo } from "@/components/common";

export const Header = () => {
    return (
        <div className="flex items-center gap-4 pt-2 pb-2">
            <Link to={"/"} className="cursor-pointer flex items-center gap-1">
                <Logo />
                <p className="text-lg font-bold">ЕГЭ Платформа</p>
            </Link>
            <nav className="flex items-center gap-4">
                <NavLink
                    to="/"
                    className={({ isActive }) =>
                        `${isActive ? "text-blue-600 underline underline-offset-4 decoration-2" : "text-gray-600 hover:text-blue-600"} text-lg font-medium rounded-lg transition-colors`
                    }
                >
                    Главная
                </NavLink>
                <NavLink
                    to="/curs"
                    className={({ isActive }) =>
                        `${isActive ? "text-blue-600 underline underline-offset-4" : "text-gray-600 hover:text-blue-600"} text-lg font-medium rounded-lg transition-colors`
                    }
                >
                    Курсы
                </NavLink>
            </nav>        
        </div>
    );
};
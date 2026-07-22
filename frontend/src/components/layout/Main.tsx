import { Outlet } from "react-router-dom";

export const Main = () => {
    return (
        <div className="min-h-[calc(100vh-10px)] grow">
            <div className="overflow-x-clip">
                <Outlet />
            </div>
        </div>
    );
};
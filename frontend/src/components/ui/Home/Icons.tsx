import { Users, Award, TrendingUp, CircleCheck } from "lucide-react";

export const Icons = () => {
    return (
        <div className="flex items-center justify-between pb-8 pt-8 max-w-7xl mx-auto">
            <div className="flex flex-col items-center text-center">
                <Users size={34} className="text-blue-600" />
                <p className="font-bold text-2xl">12000+</p>
                <p className="text-gray-600">Учеников</p>
            </div>
            <div className="flex flex-col items-center text-center">
                <Award size={34} className="text-blue-600" />
                <p className="font-bold text-2xl">95%</p>
                <p className="text-gray-600">Поступили в вузы</p>
            </div>
            <div className="flex flex-col items-center text-center">
                <TrendingUp size={34} className="text-blue-600" />
                <p className="font-bold text-2xl">87</p>
                <p className="text-gray-600">Средний балл ЕГЭ</p>
            </div>
            <div className="flex flex-col items-center text-center">
                <CircleCheck size={34} className="text-blue-600" />
                <p className="font-bold text-2xl">500+</p>
                <p className="text-gray-600">Курсов</p>
            </div>
        </div>
    );
};
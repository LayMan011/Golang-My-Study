export const Started = () => {
    return (
        <div className="bg-blue-50 flex flex-col items-center justify-center gap-6 pb-8 pt-8">
            <h1 className="font-bold text-6xl text-center">
                Подготовься к ЕГЭ
                <br />
                <span className="text-blue-600">на максимум</span>
            </h1>
            <p className="text-2xl text-gray-600">
                Более 500 курсов от лучших преподавателей.
                <br />
                Учись в своем темпе и достигай высоких результатов
            </p>
            <button className="h-20 w-40 border-4 border-gray-200 rounded-lg font-medium text-lg hover:bg-gray-200  hover:scale-110 transition-all duration-300 ease-in-out cursor-pointer">
                Смотреть курсы
            </button>
        </div>
    )
}
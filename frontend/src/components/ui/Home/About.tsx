

export const About = () => {
    return (
        <div className="bg-blue-50 flex flex-col items-center justify-center gap-4 pb-8 pt-8">
            <h1 className="font-bold text-4xl text-center mt-6">
                Как это работает
            </h1>
            <p className="text-xl text-gray-600">
                Простой процесс от выбора курса до успешной сдачи экзамена
            </p>
            <div className="flex flex-column">
                {CARDS.map((card) => (
                    <Card
                        key={card.id}
                        id={card.id}
                        title={card.title}
                        description={card.description}
                        />
                ))}
            </div>
        </div>
    )
}

interface CardI {
    id: string,
    title: string,
    description: string,
}

const CARDS: CardI[] = [
    {
        id: "1",
        title: "Выберите курс",
        description: "Подберите курс по нужному предмету и уровню сложности",
    },
    {
        id: "2",
        title: "Учитесь в своем темпе",
        description: "Проходите уроки, решайте задачи и тесты когда удобно",
    },
    {
        id: "3",
        title: "Отслеживайте прогресс",
        description: "Следите за своими результатами и улучшайте слабые места",
    },
    {
        id: "4",
        title: "Сдайте ЕГЭ на 100 баллов",
        description: "Подготовьтесь к экзамену на максимальный результат",
    },
]

const Card = ({ id, title, description }: CardI) => {
    return (
        <div className="flex flex-col items-start text-left px-4 py-2 mx-8 my-6 bg-white rounded-xl shadow-md hover:shadow-xl transition-all duration-300 hover:-translate-y-1 w-86">
            <div className="h-14 w-14 flex items-center justify-center bg-blue-600 text-white text-xl font-bold rounded-full mb-4">
                {id}
            </div>
            <p className="font-bold text-2xl text-gray-800 mb-2">{title}</p>
            <p className="text-gray-500 text-lg leading-relaxed">{description}</p>
        </div>
    )
}
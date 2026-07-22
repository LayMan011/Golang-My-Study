export const Preloader = () => (
    <div className="flex flex-col gap-2 items-center justify-center py-15">
        <div className="relative">
            <div className="w-16 h-16 border-4 border-indigo-200 border-t-blue-600 rounded-full animate-spin" />
            <div className="absolute inset-0 w-16 h-16 border-4 border-transparent border-r-blue-400 rounded-full animate-spin" style={{ animationDirection: 'reverse', animationDuration: '1s' }}></div>
        </div>
        <span>Загрузка...</span>
    </div>
)
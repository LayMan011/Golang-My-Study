import { motion } from 'motion/react';

export const Tabs = ({activeTab, setActiveTab}: { activeTab: 'login' | 'register', setActiveTab: React.Dispatch<React.SetStateAction<"login" | "register">>}) => {
    return (
        <div className="flex border-b border-border">
            <button
                onClick={() => setActiveTab('login')}
                className={`flex-1 py-4 font-semibold text-center transition-colors relative ${
                activeTab === 'login'
                    ? 'text-primary'
                    : 'text-muted-foreground hover:text-foreground'
                }`}
            >
                Вход
                {activeTab === 'login' && (
                <motion.div
                    layoutId="activeAuthTab"
                    className="absolute bottom-0 left-0 right-0 h-0.5 bg-primary"
                    transition={{ type: 'spring', stiffness: 500, damping: 30 }}
                />
                )}
            </button>
            <button
                onClick={() => setActiveTab('register')}
                className={`flex-1 py-4 font-semibold text-center transition-colors relative ${
                activeTab === 'register'
                    ? 'text-primary'
                    : 'text-muted-foreground hover:text-foreground'
                }`}
            >
                Регистрация
                {activeTab === 'register' && (
                <motion.div
                    layoutId="activeAuthTab"
                    className="absolute bottom-0 left-0 right-0 h-0.5 bg-primary"
                    transition={{ type: 'spring', stiffness: 500, damping: 30 }}
                />
                )}
            </button>
        </div>
    )
}
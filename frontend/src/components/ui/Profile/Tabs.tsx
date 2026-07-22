import { motion } from 'motion/react';

export const Tabs = ({ activeTab, setActiveTab }: { activeTab: 'progress' | 'completed', setActiveTab: React.Dispatch<React.SetStateAction<"progress" | "completed">>}) => {
    return (
        <div className="flex gap-4 mb-8 border-b border-border">
        <button
            onClick={() => setActiveTab('progress')}
            className={`relative px-6 py-3 font-semibold transition-colors ${
            activeTab === 'progress'
                ? 'text-primary'
                : 'text-muted-foreground hover:text-foreground'
            }`}
        >
            В процессе
            {activeTab === 'progress' && (
            <motion.div
                layoutId="activeProfileTab"
                className="absolute bottom-0 left-0 right-0 h-0.5 bg-primary"
                transition={{ type: 'spring', stiffness: 500, damping: 30 }}
            />
            )}
        </button>
        <button
            onClick={() => setActiveTab('completed')}
            className={`relative px-6 py-3 font-semibold transition-colors ${
            activeTab === 'completed'
                ? 'text-primary'
                : 'text-muted-foreground hover:text-foreground'
            }`}
        >
            Завершённые
            {activeTab === 'completed' && (
            <motion.div
                layoutId="activeProfileTab"
                className="absolute bottom-0 left-0 right-0 h-0.5 bg-primary"
                transition={{ type: 'spring', stiffness: 500, damping: 30 }}
            />
            )}
        </button>
        </div>
    )
}
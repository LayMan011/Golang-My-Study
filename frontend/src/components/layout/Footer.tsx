import { Link, useLocation } from "react-router-dom";
// import { Mail, Github, Twitter, Youtube, BookOpen } from "lucide-react";

// const SOCIALS = [
//     { href: "https://github.com/LayMan011", icon: Github, label: "GitHub" },
//     { href: "https://twitter.com", icon: Twitter, label: "Twitter" },
//     { href: "https://youtube.com", icon: Youtube, label: "YouTube" },
//     { href: "mailto:support@progress.com", icon: Mail, label: "Email" },
// ];

const NAV_SECTIONS = [
    {
        title: "О платформе",
        links: [
            { to: "/about", label: "О нас" },
            { to: "/blog", label: "Блог" },
        ]
    },
    {
        title: "Курсы",
        links: [
            { to: "/courses", label: "Все курсы" },
            { to: "/courses?popular=true", label: "Популярные" },
            { to: "/courses?new=true", label: "Новинки" },
        ]
    },
    {
        title: "Поддержка",
        links: [
            { to: "/help", label: "Помощь" },
            { to: "/contacts", label: "Контакты" },
        ]
    }
];

export const Footer = () => {
    const location = useLocation();
    const isChatPage = location.pathname.includes("/messages/");
    
    if (isChatPage) return null;

    return (
        <footer className="flex">
            <div className="flex flex-wrap justify-between gap-5 w-full mt-auto">
                {NAV_SECTIONS.map((block, index) => (
                    <Block 
                        key={index}
                        title={block.title} 
                        links={block.links}
                    />
                ))}
            </div>
        </footer>
    );
};

const Block = ({ title, links }: { title: string; links: { to: string; label: string }[] }) => {
    return (
        <div className="flex-1 min-w-[150px]">
            <h3 className="text-lg font-semibold mb-2.5">{title}</h3>
            {links.map((link, index) => (
                <Link 
                    key={index}
                    to={link.to}
                    className="block my-1 opacity-80 hover:opacity-100 transition-opacity"
                >
                    {link.label}
                </Link>
            ))}
        </div>
    );
};
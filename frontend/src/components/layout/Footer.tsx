import { useLocation } from "react-router-dom";
// import { Mail, Github, Twitter, Youtube, BookOpen } from "lucide-react";

// const SOCIALS = [
//     { href: "https://github.com/LayMan011", icon: Github, label: "GitHub" },
//     { href: "https://twitter.com", icon: Twitter, label: "Twitter" },
//     { href: "https://youtube.com", icon: Youtube, label: "YouTube" },
//     { href: "mailto:support@progress.com", icon: Mail, label: "Email" },
// ];

export const Footer = () => {
    const location = useLocation();
    const isChatPage = location.pathname.includes("/messages/");
    
    if (isChatPage) return null;

    return (
      <footer className="bg-card border-t border-border py-12">
        <div className="container mx-auto px-4 sm:px-6 lg:px-8">
          <div className="grid md:grid-cols-4 gap-8 mb-8">
            <div>
              <h3 className="font-semibold text-foreground mb-4">О платформе</h3>
              <ul className="space-y-2 text-sm text-muted-foreground">
                <li><a href="#" className="hover:text-foreground transition-colors">О нас</a></li>
                <li><a href="#" className="hover:text-foreground transition-colors">Блог</a></li>
                <li><a href="#" className="hover:text-foreground transition-colors">Карьера</a></li>
              </ul>
            </div>
            <div>
              <h3 className="font-semibold text-foreground mb-4">Курсы</h3>
              <ul className="space-y-2 text-sm text-muted-foreground">
                <li><a href="#" className="hover:text-foreground transition-colors">Все курсы</a></li>
                <li><a href="#" className="hover:text-foreground transition-colors">Популярные</a></li>
                <li><a href="#" className="hover:text-foreground transition-colors">Новинки</a></li>
              </ul>
            </div>
            <div>
              <h3 className="font-semibold text-foreground mb-4">Поддержка</h3>
              <ul className="space-y-2 text-sm text-muted-foreground">
                <li><a href="#" className="hover:text-foreground transition-colors">Помощь</a></li>
                <li><a href="#" className="hover:text-foreground transition-colors">Контакты</a></li>
                <li><a href="#" className="hover:text-foreground transition-colors">FAQ</a></li>
              </ul>
            </div>
            <div>
              <h3 className="font-semibold text-foreground mb-4">Правовая информация</h3>
              <ul className="space-y-2 text-sm text-muted-foreground">
                <li><a href="#" className="hover:text-foreground transition-colors">Условия использования</a></li>
                <li><a href="#" className="hover:text-foreground transition-colors">Политика конфиденциальности</a></li>
              </ul>
            </div>
          </div>
          <div className="pt-8 border-t border-border text-center text-sm text-muted-foreground">
            <p>© 2026 ЕГЭ Платформа. Все права защищены.</p>
          </div>
        </div>
      </footer>
    );
};

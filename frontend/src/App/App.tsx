import { Routes, Route } from "react-router-dom";
import { ThemeProvider } from '@/context';
import { Footer, Header, Main } from "@/components/layout/index";
import { Home, Courses, Auth, Profile } from "@/pages";
import { Toaster } from "react-hot-toast";

export function App() {
    return (<>
    <ThemeProvider >
        <Toaster toastOptions={{ duration: 1000 }} reverseOrder={false} />
        <Header />
        <Routes >
            <Route element={<Main />}>
                <Route path="/" index element={<Home />} />
                <Route path="/courses" element={<Courses />} />
                <Route path="/auth" element={<Auth />} />
                <Route path="/profile" element={<Profile />} />
            </Route>
        </Routes>
        <Footer />
    </ThemeProvider>
    </>)
}

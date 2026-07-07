import { Footer, Header, Main } from "@/components/layout/index";
import { Home } from "@/pages/Home";
import { Toaster } from "react-hot-toast";
import { Routes, Route } from "react-router-dom";

export function App() {
    return (<>
        <Toaster toastOptions={{ duration: 1000 }} reverseOrder={false} />
        <Header />
        <Routes >
            <Route element={<Main />}>
                <Route path="/" index element={<Home />} />
            </Route>
        </Routes>
        <Footer />
    </>)
}
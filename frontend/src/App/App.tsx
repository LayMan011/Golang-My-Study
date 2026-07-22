import { Routes, Route } from "react-router-dom";
import { ThemeProvider } from '@/context';
import { Footer, Header, Main } from "@/components/layout";
import { Preloader } from "@/components/common";
import { Home, AuthModal } from "@/pages";
import { Toaster } from "react-hot-toast";
import { lazy, Suspense } from "react";

const Courses = lazy(() => import("@/pages/Cuorses"));
const Profile = lazy(() => import("@/pages/Profile"));

export function App() {
  return (
    <ThemeProvider>
      <Toaster toastOptions={{ duration: 3000 }} reverseOrder={false} />
      <Header />

      <Routes>
        <Route element={<Main />}>
          <Route path="/" index element={<Home />} />

          <Route
            path="/courses"
            element={
              <Suspense fallback={<Preloader />}>
                <Courses />
              </Suspense>
            }
          />

          <Route
            path="/profile"
            element={
              <Suspense fallback={<Preloader />}>
                <Profile />
              </Suspense>
            }
          />
        </Route>
      </Routes>

      <AuthModal />

      <Footer />
    </ThemeProvider>
  );
}
import { Outlet, createBrowserRouter } from "react-router-dom";
import { HomePage, LoginPage } from "@/pages";
import AuthProvider from "@/contexts/auth_context";
import { AppBar } from "@/modules/navbars";
import OnboardingPage from "./pages/onboarding";

const router = createBrowserRouter([
    {
        path: "/",
        element: (
            <AuthProvider>
                <AppBar />
                <div className="w-full h-[54px] md:h-[64px]" />
                <Outlet />
            </AuthProvider>
        ),
        children: [
            {
                path: "",
                element: <HomePage />,
            },
        ],
    },
    {
        path: "/auth",
        children: [
            {
                path: "",
                element: <LoginPage />,
            },
        ],
    },
    {
        path: "/onboarding",
        element: (
            <AuthProvider>
                <OnboardingPage />
            </AuthProvider>
        ),
    },
]);

export default router;

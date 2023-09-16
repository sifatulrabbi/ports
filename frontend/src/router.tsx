import { createBrowserRouter } from "react-router-dom";
import { HomePage, LoginPage } from "./pages";
import App from "./app";

const router = createBrowserRouter([
    {
        path: "/",
        element: <App />,
        children: [
            {
                path: "",
                element: <HomePage />,
            },
            {
                path: "/login",
                element: <LoginPage />,
            },
        ],
    },
]);

export default router;

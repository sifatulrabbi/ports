import React from "react";
import { Outlet } from "react-router-dom";
import AuthProvider from "./contexts/auth_context";

const App: React.FC = () => {
    return (
        <AuthProvider>
            <Outlet />
        </AuthProvider>
    );
};

export default App;

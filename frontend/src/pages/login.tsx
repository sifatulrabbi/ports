import React from "react";
import { Typography, Button } from "@mui/material";
import { useAuth0 } from "@auth0/auth0-react";

const LoginPage: React.FC = () => {
    const { loginWithRedirect } = useAuth0();

    return (
        <div className="w-full flex flex-col p-8 justify-center items-center h-[100vh] gap-4">
            <Typography variant="h4">Ports</Typography>
            <Button
                onClick={() => loginWithRedirect()}
                variant="contained"
                size="large"
                className="w-max mx-auto"
            >
                Login
            </Button>
        </div>
    );
};

export default LoginPage;

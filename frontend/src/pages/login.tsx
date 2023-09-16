import React from "react";
import { Button } from "@mui/material";

const LoginPage: React.FC = () => {
    const handleSubmit = async (e: React.SyntheticEvent<HTMLFormElement>) => {
        e.preventDefault();
    };

    return (
        <div className="tw-w-full tw-flex tw-flex-col tw-p-4">
            <form
                action="submit"
                className="w-full flex flex-col max-w-[600px]"
                onSubmit={handleSubmit}
            >
                <Button type="submit" size="large" variant="contained">
                    Login
                </Button>
            </form>
        </div>
    );
};

export default LoginPage;

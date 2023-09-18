import React from "react";
import { useNavigate } from "react-router-dom";
import { useAuth0 } from "@auth0/auth0-react";
import { CircularProgress, Typography } from "@mui/material";
import { useHttp } from "@/hooks";

type AuthCtx = null;

const AuthContext = React.createContext<AuthCtx>(null);

// eslint-disable-next-line react-refresh/only-export-components
export const useAuthContext = () => React.useContext(AuthContext);

const AuthProvider: React.FC<{ children?: React.ReactNode }> = ({
    children,
}) => {
    const [loading, setLoading] = React.useState(false);
    const { user, isAuthenticated, isLoading, logout } = useAuth0();
    const { https } = useHttp();

    const navigate = useNavigate();

    React.useEffect(() => {
        if (isLoading) return;
        if (!isAuthenticated || !user) navigate("/auth");
        else silentLogin();
    }, [isLoading, isAuthenticated, user]);

    const silentLogin = async () => {
        try {
            setLoading(true);
            const res = await (
                await https()
            ).get("/users/profile", {
                params: {
                    email: user?.email,
                },
            });
            console.log(res.data);
        } catch (err: any) {
            if (
                err.response &&
                err.response.data &&
                err.response.data.status_code === 404
            ) {
                navigate("/onboarding");
            } else if (isAuthenticated) {
                await logout({
                    logoutParams: {
                        returnTo: window.location.origin + "/auth",
                    },
                });
            } else navigate("/auth");
        } finally {
            setLoading(false);
        }
    };

    if (loading || isLoading || !isAuthenticated || !user) {
        return (
            <div className="fixed top-1/2 left-1/2 -translate-x-1/2 -translate-y-1/2 w-full max-w-[85vw] sm:max-w-[400px] h-full max-h-[300px] bg-white flex flex-col justify-center items-center rounded gap-6">
                <Typography variant="h5">Authenticating...</Typography>
                <CircularProgress />
            </div>
        );
    }

    return <AuthContext.Provider value={null}>{children}</AuthContext.Provider>;
};

export default AuthProvider;

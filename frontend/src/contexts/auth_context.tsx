import React from "react";
import { useNavigate } from "react-router-dom";
import { useAuth0 } from "@auth0/auth0-react";
import { CircularProgress, Typography } from "@mui/material";

const AuthContext = React.createContext(null);

// eslint-disable-next-line react-refresh/only-export-components
export const useAuthContext = () => React.useContext(AuthContext);

const AuthProvider: React.FC<{ children?: React.ReactNode }> = ({
    children,
}) => {
    const [loading, setLoading] = React.useState(false);
    const { user, isAuthenticated, isLoading, getAccessTokenSilently } =
        useAuth0();

    const navigate = useNavigate();

    React.useEffect(() => {
        if (isLoading) return;
        if (!isAuthenticated || !user) navigate("/auth");
        else silentLogin();
    }, [isLoading, isAuthenticated, user]);

    const silentLogin = async () => {
        try {
            setLoading(true);
            const token = await getAccessTokenSilently({ cacheMode: "on" });
            console.log(token && "got a token");
            setLoading(false);
        } catch (err) {
            console.error(err);
            navigate("/auth");
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

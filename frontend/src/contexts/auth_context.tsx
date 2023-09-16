import React from "react";
import { useNavigate } from "react-router-dom";

const AuthContext = React.createContext({ user: null });

// eslint-disable-next-line react-refresh/only-export-components
export const useAuthContext = () => React.useContext(AuthContext);

const AuthProvider: React.FC<{ children?: React.ReactNode }> = ({
    children,
}) => {
    const [loading, setLoading] = React.useState(false);
    const [user, setUser] = React.useState(null);

    const navigate = useNavigate();

    React.useEffect(() => {
        if (user) return;
        silentLogin();
        // eslint-disable-next-line react-hooks/exhaustive-deps
    }, [user]);

    const silentLogin = async () => {
        try {
            setUser(null);
            setLoading(false);
            navigate("/login");
        } catch (err) {
            console.error(err);
            navigate("/login");
        }
    };

    return (
        <AuthContext.Provider value={{ user }}>
            {!loading && children}
            {loading && <div className="w-full"></div>}
        </AuthContext.Provider>
    );
};

export default AuthProvider;

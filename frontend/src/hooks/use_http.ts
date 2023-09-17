import { useAuth0 } from "@auth0/auth0-react";
import axios from "axios";

const useHttps = () => {
    const { getAccessTokenSilently } = useAuth0();

    const https = async () => {
        const token = await getAccessTokenSilently({ cacheMode: "on" });
        if (!token) throw new Error("User not logged in!");
        const client = axios.create({
            baseURL: "https://localhost:8000",
            headers: {
                "Authorization": token,
                "Content-Type": "application/json",
                "Accept": "application/json",
            },
        });
        return client;
    };

    const http = () => {
        const client = axios.create({
            baseURL: "https://localhost:8000",
            headers: {
                "Content-Type": "application/json",
                "Accept": "application/json",
            },
        });
        return client;
    };

    return { https, http };
};

export default useHttps;

import React from "react";
import ReactDOM from "react-dom/client";
import { RouterProvider } from "react-router-dom";
import {
    createTheme,
    StyledEngineProvider,
    ThemeProvider,
} from "@mui/material/styles";
import { CssBaseline } from "@mui/material";
import { Auth0Provider } from "@auth0/auth0-react";
import router from "@/router.tsx";

import "@fontsource/roboto/300.css";
import "@fontsource/roboto/400.css";
import "@fontsource/roboto/500.css";
import "@fontsource/roboto/700.css";
import "./styles/index.scss";

const rootElement = document.getElementById("root");
const root = ReactDOM.createRoot(rootElement!);

// All `Portal`-related components need to have the the main app wrapper element as a container
// so that the are in the subtree under the element used in the `important` option of the Tailwind's config.
const theme = createTheme({
    components: {
        MuiPopover: {
            defaultProps: {
                container: rootElement,
            },
        },
        MuiPopper: {
            defaultProps: {
                container: rootElement,
            },
        },
    },
});

root.render(
    <React.StrictMode>
        <StyledEngineProvider injectFirst>
            <ThemeProvider theme={theme}>
                <CssBaseline />
                <Auth0Provider
                    domain="sifatul.us.auth0.com"
                    clientId="qZJbS9PON3T090pNuN03FMxcfLMIx7o0"
                    authorizationParams={{
                        redirect_uri: window.location.origin,
                    }}
                >
                    <RouterProvider router={router} />
                </Auth0Provider>
            </ThemeProvider>
        </StyledEngineProvider>
    </React.StrictMode>,
);

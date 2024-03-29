import React from "react";
import {
    AppBar as MUIAppBar,
    Toolbar,
    IconButton,
    Typography,
    Box,
    Button,
    Divider,
    List,
    ListItem,
    ListItemButton,
    ListItemText,
    Drawer,
} from "@mui/material";
import MenuIcon from "@mui/icons-material/Menu";
import { useAuth0 } from "@auth0/auth0-react";

const navItems: any[] = [];

const AppBar: React.FC = () => {
    const [mobileOpen, setMobileOpen] = React.useState(false);

    const { logout } = useAuth0();

    const handleLogout = () =>
        logout({
            logoutParams: {
                returnTo: window.location.origin + "/auth",
            },
        });

    const handleDrawerToggle = () => {
        setMobileOpen((prevState) => !prevState);
    };

    const drawer = (
        <Box onClick={handleDrawerToggle} sx={{ textAlign: "center" }}>
            <Typography variant="h6" sx={{ my: 2 }}>
                Ports
            </Typography>
            <Divider />
            <List>
                {navItems.map((item) => (
                    <ListItem key={item.name + item.path} disablePadding>
                        <ListItemButton sx={{ textAlign: "center" }}>
                            <ListItemText primary={item.name} />
                        </ListItemButton>
                    </ListItem>
                ))}
                <ListItem disablePadding>
                    <ListItemButton sx={{ textAlign: "center" }}>
                        <ListItemText primary="Logout" onClick={handleLogout} />
                    </ListItemButton>
                </ListItem>
            </List>
        </Box>
    );

    return (
        <>
            <MUIAppBar component="nav">
                <Toolbar>
                    <Typography
                        variant="h6"
                        component="div"
                        sx={{
                            flexGrow: 1,
                        }}
                    >
                        Ports
                    </Typography>
                    <Box sx={{ display: { xs: "none", sm: "block" } }}>
                        {navItems.map((item) => (
                            <Button
                                key={"nav-links-" + item.name + item.path}
                                sx={{ color: "#fff" }}
                            >
                                {item.name}
                            </Button>
                        ))}
                        <Button sx={{ color: "#fff" }} onClick={handleLogout}>
                            Logout
                        </Button>
                    </Box>
                    <IconButton
                        color="inherit"
                        aria-label="open drawer"
                        edge="start"
                        onClick={handleDrawerToggle}
                        sx={{ display: { sm: "none" } }}
                    >
                        <MenuIcon />
                    </IconButton>
                </Toolbar>
            </MUIAppBar>

            <Drawer
                variant="temporary"
                open={mobileOpen}
                onClose={handleDrawerToggle}
                ModalProps={{
                    keepMounted: true, // Better open performance on mobile.
                }}
                sx={{
                    "display": { xs: "block", sm: "none" },
                    "& .MuiDrawer-paper": {
                        boxSizing: "border-box",
                        width: "240px",
                    },
                }}
            >
                {drawer}
            </Drawer>
        </>
    );
};

export default AppBar;

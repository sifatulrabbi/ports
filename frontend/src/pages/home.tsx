import React from "react";
import {
    Box,
    Button,
    CircularProgress,
    Divider,
    Typography,
    TextField,
} from "@mui/material";
import { User, useAuth0 } from "@auth0/auth0-react";
import EditIcon from "@mui/icons-material/Edit";
import SaveIcon from "@mui/icons-material/Save";
import { useHttp } from "@/hooks";

const HomePage: React.FC = () => {
    const [editProfile, setEditProfile] = React.useState(false);
    const [savingProfile, setSavingProfile] = React.useState(false);
    const [name, setName] = React.useState("");
    const [title, setTitle] = React.useState("");

    const { user } = useAuth0();
    const { https } = useHttp();

    const handleSubmit = async (e: React.SyntheticEvent) => {
        e.preventDefault();
        try {
            setSavingProfile(true);
            const payload = {
                name,
                title,
            };
            const res = await (await https()).put("/users", payload);
            console.log(res);
            handleResetForm();
        } catch (err) {
            console.error(err);
        } finally {
            setSavingProfile(false);
        }
    };

    const handleResetForm = async () => {
        setName("");
        setTitle("");
    };

    const InfoSection = (p: { label: string; value: string }) => (
        <div className="w-full flex flex-col">
            <Typography variant="body2">{p.label}</Typography>
            <Typography variant="h6">{p.value}</Typography>
        </div>
    );

    const ProfileInfoSection = (
        <>
            <InfoSection label="Name" value={(user as User).name || ""} />
            <InfoSection label="Email" value={(user as User).email || ""} />
            <InfoSection label="Title" value="Full Stack Developer" />
        </>
    );

    const ProfileEditForm = (
        <form
            action="submit"
            onSubmit={handleSubmit}
            onReset={handleResetForm}
            className="w-full flex flex-col justify-start gap-4"
        >
            <TextField
                value={name}
                onChange={(e) => setName(e.currentTarget.value)}
                name="user-name"
                id="user-name"
                required
                type="text"
                label="Name"
                variant="filled"
            />
            <TextField
                value={title}
                onChange={(e) => setTitle(e.currentTarget.value)}
                name="user-title"
                id="user-title"
                required
                type="text"
                label="Title"
                variant="filled"
            />
        </form>
    );

    return (
        <div className="w-full p-6 flex flex-col justify-start items-start">
            <div className="w-full flex flex-col gap-2 md:px-[7vw] lg:px-[15vw]">
                {!editProfile ? ProfileInfoSection : ProfileEditForm}
                {!editProfile && (
                    <Button
                        onClick={() => setEditProfile(true)}
                        variant="contained"
                        className="mt-2 max-w-max gap-2"
                    >
                        <EditIcon fontSize="small" />
                        Edit Profile
                    </Button>
                )}
                {editProfile && (
                    <Box
                        sx={{
                            display: "flex",
                            alignItems: "center",
                            justifyContent: "start",
                            gap: "1rem",
                        }}
                    >
                        <Button
                            type="submit"
                            variant="contained"
                            className="max-w-max gap-2"
                            disabled={savingProfile}
                        >
                            {savingProfile ? (
                                <CircularProgress
                                    size="1rem"
                                    sx={{ color: "#fff" }}
                                />
                            ) : (
                                <SaveIcon fontSize="small" />
                            )}
                            Save
                        </Button>
                        <Button
                            onClick={() => {
                                handleResetForm();
                                setEditProfile(false);
                            }}
                            variant="outlined"
                            color="error"
                            disabled={savingProfile}
                            type="reset"
                        >
                            Cancel
                        </Button>
                    </Box>
                )}
                <Divider className="mt-4" />
            </div>
        </div>
    );
};

export default HomePage;

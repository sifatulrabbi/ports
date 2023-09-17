import React from "react";
import { Typography, TextField, Button } from "@mui/material";
import ArrowRightIcon from "@mui/icons-material/ArrowRight";
import { useHttp } from "@/hooks";
import { useNavigate } from "react-router-dom";

const OnboardingPage: React.FC = () => {
    const [name, setName] = React.useState("");
    const [title, setTitle] = React.useState("");

    const { http } = useHttp();
    const navigate = useNavigate();

    const handleSubmit = async (e: React.SyntheticEvent<HTMLFormElement>) => {
        e.preventDefault();
        try {
            const res = await http().post("/users", {
                name,
                title,
            });
            if (res.data.id) {
                navigate("/");
            }
        } catch (err) {
            console.error(err);
        }
    };

    const handleReset = () => {};

    return (
        <div className="w-full flex flex-col justify-center items-center min-h-[100vh]">
            <Typography variant="subtitle1">Welcome to</Typography>
            <Typography variant="h4">Ports</Typography>
            <form
                action="submit"
                onSubmit={handleSubmit}
                onReset={handleReset}
                className="w-full max-w-[500px] p-6 bg-white flex flex-col gap-4"
            >
                <TextField
                    label="Name"
                    id="user-name"
                    name="user-name"
                    required
                    value={name}
                    onChange={(e) => setName(e.currentTarget.value)}
                    variant="filled"
                    placeholder="i.e. Sifatul Rabbi"
                />
                <TextField
                    label="Title/Occupation"
                    id="user-title"
                    name="user-title"
                    required
                    value={title}
                    onChange={(e) => setTitle(e.currentTarget.value)}
                    variant="filled"
                    placeholder="i.e. Full Stack Developer"
                />
                <Button type="submit" variant="contained" size="large">
                    Create Profile
                    <ArrowRightIcon />
                </Button>
            </form>
        </div>
    );
};

export default OnboardingPage;

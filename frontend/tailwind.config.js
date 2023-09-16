/** @type {import('tailwindcss').Config} */
export default {
    content: ["./src/**/*.{html,js,jsx,tsx,ts}"],
    important: true,
    // prefix: "tw-",
    theme: {
        extend: {},
    },
    corePlugins: {
        // Remove the Tailwind CSS preflight styles so it can use Material UI's preflight instead (CssBaseline).
        preflight: false,
    },
    plugins: [],
};

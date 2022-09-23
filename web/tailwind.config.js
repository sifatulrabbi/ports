module.exports = {
	content: ["./src/**/*.{tsx,ts,js,jsx}"],
	plugins: [require("@tailwindcss/typography"), require("daisyui")],
	daisyui: {
		theme: ["winter"],
	},
}

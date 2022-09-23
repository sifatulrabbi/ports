const API_URL = process.env.NEXT_PUBLIC_SERVER_URL

if (!API_URL) {
	console.error("Unable to find required env vars")
	process.exit(1)
}

export const configs = {
	API_URL,
}

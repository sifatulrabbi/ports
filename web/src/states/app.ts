import { atom } from "recoil"
import type { User } from "../types"

export const showAppSidebarState = atom({
	key: "showAppSidebarState",
	default: false,
})

export const userState = atom<User | null>({
	key: "userState",
	default: {
		id: "sifatul-rabbi",
		email: "mdsifatulislam.rabbi@gmail.com",
		name: "Sifatul Rabbi",
		avatar: "",
		bio: "",
		title: "Full Stack Developer",
		timezone: "+6",
	},
})

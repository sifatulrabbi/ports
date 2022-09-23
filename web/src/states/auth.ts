import {atom} from "recoil"
import {IAuth} from "@/interfaces"

export const authState = atom<IAuth | null>({
	key: "authStateKey",
	default: null,
})

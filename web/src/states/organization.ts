import { atom } from "recoil"
import type { User } from "../types"
import { mockUsers } from "../assets/mock"

export const organizationMembersState = atom<User[]>({
	key: "organizationMembersState",
	default: mockUsers,
})

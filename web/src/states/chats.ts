import { atom } from "recoil"
import type { Message, Room } from "../types"
import { mockMessages, mockRoom } from "../assets/mock"

export const messagesState = atom<Message[]>({
	key: "messagesState",
	default: [...mockMessages],
})

export const currentRoomState = atom<Room | null>({
	key: "currentRoomState",
	default: mockRoom,
})

export const roomsState = atom<Room[]>({
	key: "roomsState",
	default: [mockRoom],
})

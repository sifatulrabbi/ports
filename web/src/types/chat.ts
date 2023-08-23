import type { User } from "./user"

type UnixTimestamp = number

type Richtext = `<p>${string}</p>`

type LinkString = `http${"s" | ""}://${string}`

export type Message = {
	id: string
	sender_id: string
	room_id: string
	created_at: UnixTimestamp
	updated_at: UnixTimestamp
	body: MessageBody
	replied_to: Message | null
	metadata: MessageMetadata
}

export type MessageBody = {
	richtext: Richtext
	audios: LinkString[]
	videos: LinkString[]
	images: LinkString[]
}

export type MessageMetadata = {
	sender_ip: string
	receiver_ip: string
	api_version: string
}

export type Room = {
	id: string
	name: string
	description: string
	participant_ids: string[]
	participants: User[]
	created_at: UnixTimestamp
	updated_at: UnixTimestamp
}

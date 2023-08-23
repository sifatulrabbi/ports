export type User = {
	id: string
	email: string
	name: string
	avatar: string
	bio: string
	title: string
	timezone: string
}

export type Organization = {
	id: string
	email: string
	name: string
	avatar: string
	bio: string
	timezone: string
	users: User[]
}

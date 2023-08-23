import { Message, Room, User } from "../types"

export const mockMessages: Message[] = [
	{
		id: "7fds7fas",
		room_id: "test-room-1",
		sender_id: "temujin-hideyoshi",
		body: {
			richtext: `<p>
<strong>🌟 Welcome to Ports! 🌟</strong>
<br/>
Before you dive into the conversations, here are a few guidelines to ensure a friendly and positive experience for everyone:</li>
<ul>
<li>Respect Everyone: All members are expected to be respectful to others. No hate speech, bullying, or discrimination will be tolerated.</li>
<li>Stay On Topic: Keep the conversation relevant to the subject of the chat room. If you wish to discuss something off-topic, consider creating a new chat or joining another room.</li>
<li>No Spamming: Please refrain from sending repetitive messages, promotional content, or unsolicited links.</li>
<li>Protect Your Privacy: Avoid sharing personal information like your email, home address, and phone number in public chats.</li>
<li>Listen to the Moderators: Our team of moderators work hard to keep the chat environment safe and welcoming. Please follow their instructions and guidance.</li>
<li>Report Issues: If you encounter a disruptive user or inappropriate content, please report it immediately using the "Report" button or by messaging a moderator.</li>
<li>Enjoy and Engage: Our chat room is a platform for sharing ideas, making friends, and having fun. Engage, participate, and make the most of it!</li>
</ul>
Thanks for being a part of our community. Let's keep it a positive and enjoyable space for everyone!
<br/>
<code>Happy chatting! 🚀</code>
</p>`,
			audios: [],
			videos: [],
			images: [],
		},
		created_at: Date.now() - 1000 * 300,
		updated_at: Date.now() - 1000 * 200,
		replied_to: null,
		metadata: {
			api_version: "",
			receiver_ip: "",
			sender_ip: "",
		},
	},
	{
		id: "fadsfdafa",
		room_id: "test-room-1",
		sender_id: "temujin-hideyoshi",
		body: {
			richtext: "<p>Hello everyone</p>",
			audios: [],
			videos: [],
			images: [],
		},
		created_at: Date.now(),
		updated_at: Date.now(),
		replied_to: null,
		metadata: {
			api_version: "",
			receiver_ip: "",
			sender_ip: "",
		},
	},
	{
		id: "1wqrqrqwr",
		room_id: "test-room-1",
		sender_id: "hashirama-senju",
		body: {
			richtext: "<p>This is a test message.</p>",
			audios: [],
			videos: [],
			images: [],
		},
		created_at: Date.now() - 1000 * 300,
		updated_at: Date.now() - 1000 * 200,
		replied_to: null,
		metadata: {
			api_version: "",
			receiver_ip: "",
			sender_ip: "",
		},
	},
	{
		id: "3hhgffgfdg",
		room_id: "test-room-1",
		sender_id: "temujin-hideyoshi",
		body: {
			richtext: "<p>Welcome to the room chat</p>",
			audios: [],
			videos: [],
			images: [],
		},
		created_at: Date.now() - 1000 * 300,
		updated_at: Date.now() - 1000 * 200,
		replied_to: null,
		metadata: {
			api_version: "",
			receiver_ip: "",
			sender_ip: "",
		},
	},
]

export const mockUsers: User[] = [
	{
		id: "hashirama-senju",
		email: "islammasraful@gmail.com",
		name: "Hashirama Senju",
		avatar: "",
		title: "Frontend Engineer",
		bio: "",
		timezone: "+9",
	},
	{
		id: "temujin-hideyoshi",
		email: "sifatuli.r@gmail.com",
		name: "Temujin Hideyoshi",
		avatar: "",
		title: "Backend Engineer",
		bio: "",
		timezone: "+9",
	},
	{
		id: "sifatul-rabbi",
		email: "mdsifatulislam.rabbi@gmail.com",
		name: "Sifatul Rabbi",
		avatar: "",
		bio: "",
		title: "Full Stack Developer",
		timezone: "+6",
	},
]

export const mockRoom: Room = {
	id: "test-room-1",
	name: "Test Room",
	description: "Test room description",
	participant_ids: [],
	participants: [...mockUsers],
	created_at: Date.now(),
	updated_at: Date.now(),
}

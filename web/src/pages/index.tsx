import type { NextPage } from "next"
import { useState } from "react"
import { useRecoilValue } from "recoil"
import { messagesState } from "../states"
import MessageBubble from "../modules/chats/MessageBubble"
import MessageInput from "../modules/chats/MessageInput"

const Home: NextPage = () => {
	const messages = useRecoilValue(messagesState)
	const [focused, setFocused] = useState(false)

	return (
		<>
			<div
				className={`w-full bg-white pt-[76px] px-4 flex flex-col justify-start items-start gap-6
				${focused ? "pb-[38vh]" : "pb-[112px]"}
				lg:max-w-[800px] lg:mx-auto lg:border-x`}
			>
				{messages.map((m) => (
					<MessageBubble key={m.id} message={m} />
				))}
			</div>
			<MessageInput focused={focused} setFocused={setFocused} />
		</>
	)
}

export default Home

import React, { useState, useEffect } from "react"
import { useRecoilValue } from "recoil"
import Image from "next/image"

import { currentRoomState } from "../../states"
import type { User, Message } from "../../types"
import dayjs from "dayjs"

type Props = {
	message: Message
}

const MessageBubble: React.FC<Props> = ({ message }) => {
	const [author, setAuthor] = useState<User | null>(null)
	const room = useRecoilValue(currentRoomState)

	useEffect(() => {
		const user = room?.participants.find((p) => p.id === message.sender_id)
		setAuthor(user || null)
	}, [message, room])

	if (!author) return <></>

	return (
		<div className="w-full flex flex-row justify-start items-start gap-2">
			<div className="w-8 h-8 min-w-8 min-h-8 rounded-lg overflow-hidden bg-gray-200">
				{author.avatar && (
					<Image
						src={author.avatar}
						height="24px"
						width="24px"
						alt=""
					/>
				)}
			</div>

			<div className="w-full flex flex-col justify-start items-start">
				<h5 className="text-base leading-[1.1] font-bold">
					{author.name}
				</h5>
				<span className="font-normal text-xs text-slate-400 w-full flex justify-between">
					{author.title}
					<span>{dayjs(message.created_at).format("hh:mma")}</span>
				</span>

				<div className="w-full prose prose-p:my-0">
					<p
						dangerouslySetInnerHTML={{
							__html: message.body.richtext,
						}}
					></p>
				</div>
			</div>
		</div>
	)
}

export default MessageBubble

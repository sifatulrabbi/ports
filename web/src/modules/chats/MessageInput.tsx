import React, { SyntheticEvent, useEffect, useRef, useState } from "react"
import { v4 } from "uuid"
import {
	FaArrowRight,
	FaImage,
	FaMicrophone,
	FaTag,
	FaVideo,
} from "react-icons/fa"
import { currentRoomState, messagesState, userState } from "../../states"
import { useRecoilValue, useSetRecoilState } from "recoil"
import { Message } from "../../types"

type Props = {
	focused: boolean
	setFocused: React.Dispatch<React.SetStateAction<boolean>>
}

const MessageInput: React.FC<Props> = ({ focused, setFocused }) => {
	const user = useRecoilValue(userState)
	const setMessages = useSetRecoilState(messagesState)
	const room = useRecoilValue(currentRoomState)

	const parentEl = useRef<HTMLFormElement>(null)

	const [text, setText] = useState("")

	useEffect(() => {
		if (focused) window.addEventListener("click", handleOutsideClick)
		else window.removeEventListener("click", handleOutsideClick)
		return () => {
			window.removeEventListener("click", handleOutsideClick)
		}
	}, [focused])

	const handleOutsideClick = (e: MouseEvent) => {
		if (!parentEl.current) return
		if (parentEl.current.contains(e.target as Node)) return
		setFocused(false)
	}

	const handleSubmit = async (e: SyntheticEvent<HTMLFormElement>) => {
		e.preventDefault()
		if (!user || !room) return
		try {
			const msg: Message = {
				id: v4(),
				body: {
					richtext: `<p>${text}</p>`,
					audios: [],
					images: [],
					videos: [],
				},
				created_at: Date.now(),
				updated_at: Date.now(),
				metadata: {
					api_version: "",
					receiver_ip: "",
					sender_ip: "",
				},
				replied_to: null,
				room_id: room.id,
				sender_id: user.id,
			}
			setMessages((prev) => [...prev, msg])
			setText("")
			setFocused(false)
			document.documentElement.scrollTo({
				top: document.documentElement.clientHeight + window.innerHeight,
				behavior: "smooth",
			})
			// eslint-disable-next-line @typescript-eslint/no-explicit-any
		} catch (err: any) {
			console.log(err.response?.data || err)
		}
	}

	return (
		<form
			ref={parentEl}
			onSubmit={handleSubmit}
			className="w-full fixed flex flex-col justify-start items-start bottom-0 left-0 right-0 bg-white p-4 border-t gap-4"
		>
			<textarea
				name="message-input"
				required
				className={`dui--textarea dui--textarea-bordered text-base resize-none w-full transition-[height] 
                ${!focused ? "h-[48px]" : "h-[25vh]"}
                lg:max-w-[800px] lg:mx-auto`}
				placeholder="Write a message..."
				value={text}
				onChange={(e) => setText(e.currentTarget.value)}
				onFocus={() => setFocused(true)}
			></textarea>

			{focused && (
				<div className="w-full flex flex-row justify-start items-center gap-2 lg:max-w-[800px] lg:mx-auto">
					<button
						type="button"
						className="dui--btn dui--btn-square dui--btn-sm dui--btn-active dui--btn-ghost"
					>
						<FaImage />
					</button>
					<button
						type="button"
						className="dui--btn dui--btn-square dui--btn-sm dui--btn-active dui--btn-ghost"
					>
						<FaVideo />
					</button>
					<button
						type="button"
						className="dui--btn dui--btn-square dui--btn-sm dui--btn-active dui--btn-ghost"
					>
						<FaMicrophone />
					</button>
					<button
						type="button"
						className="dui--btn dui--btn-square dui--btn-sm dui--btn-active dui--btn-ghost"
					>
						<FaTag />
					</button>
					<button
						type="submit"
						className="dui--btn dui--btn-square dui--btn-primary dui--btn-active dui--btn-sm ml-auto"
					>
						<FaArrowRight />
					</button>
				</div>
			)}
		</form>
	)
}

export default MessageInput

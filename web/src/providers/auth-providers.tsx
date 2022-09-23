import React, {useEffect, useState} from "react"
import {useRouter} from "next/router"
import axios from "axios"
import {STORED_SESSION_KEY} from "@/constants"
import {getCoreApiUrl} from "@/utils"

import {useSetRecoilState} from "recoil"
import {authState} from "@/states"

interface Props {
	children?: React.ReactNode
}

export const AuthProvider: React.FC<Props> = ({children}) => {
	const setAuth = useSetRecoilState(authState)
	const [loading, setLoading] = useState(false)
	const router = useRouter()

	useEffect(() => {
		if (router.pathname === "/login" || router.pathname === "/register") {
			return
		}

		const session = sessionStorage.getItem(STORED_SESSION_KEY)
		if (!session) {
			setAuth(null)
			router.replace("/login")
			return
		}

		;(async () => {
			try {
				setLoading(true)
				// Get user's credentials.
				const res = await axios.get(getCoreApiUrl(`/auth/profile`), {
					headers: {Authorization: "Bearer " + session},
				})
				if (!res.data.success) {
					setAuth(null)
					console.error(res.data.message)
					router.replace("/login")
				} else {
					setAuth(res.data.data)
					setLoading(false)
				}
			} catch (err) {
				console.error(err)
				router.replace("/login")
			}
		})()
	}, [])

	if (loading) return <div>Loading...</div>
	return <>{children}</>
}

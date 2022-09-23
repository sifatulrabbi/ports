import "../../styles/globals.scss"
import type {AppProps} from "next/app"
import {AuthProvider} from "@/providers"
import {RecoilRoot} from "recoil"

function MyApp({Component, pageProps}: AppProps) {
	return (
		<RecoilRoot>
			<AuthProvider>
				<Component {...pageProps} />
			</AuthProvider>
		</RecoilRoot>
	)
}

export default MyApp

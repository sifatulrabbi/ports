import "../../styles/globals.scss"
import type { AppProps } from "next/app"
import { RecoilRoot } from "recoil"

import AppNavbar from "../modules/appnavbar"
import AppSidebar from "../modules/appsidebar"

function MyApp({ Component, pageProps }: AppProps) {
	return (
		<RecoilRoot>
			<AppNavbar />
			<AppSidebar />
			<Component {...pageProps} />
		</RecoilRoot>
	)
}

export default MyApp

import React from "react"
import { useSetRecoilState } from "recoil"
import { showAppSidebarState } from "../../states"

const AppNavbar: React.FC = () => {
	const setShowAppSidebar = useSetRecoilState(showAppSidebarState)

	return (
		<nav className="fixed z-[999] top-0 right-0 left-0">
			<div className="dui--navbar bg-base-100 border-b">
				<div className="flex-none">
					<button
						onClick={() => setShowAppSidebar((p) => !p)}
						className="dui--btn dui--btn-square dui--btn-ghost"
					>
						<svg
							xmlns="http://www.w3.org/2000/svg"
							fill="none"
							viewBox="0 0 24 24"
							className="inline-block w-5 h-5 stroke-current"
						>
							<path
								strokeLinecap="round"
								strokeLinejoin="round"
								strokeWidth="2"
								d="M4 6h16M4 12h16M4 18h16"
							></path>
						</svg>
					</button>
				</div>
				<div className="flex-1">
					<a className="dui--btn dui--btn-ghost normal-case text-xl">
						Ports
					</a>
				</div>
			</div>
		</nav>
	)
}

export default AppNavbar

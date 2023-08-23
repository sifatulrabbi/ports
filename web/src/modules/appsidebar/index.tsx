import React from "react"
import { useRecoilState } from "recoil"
import { showAppSidebarState } from "../../states"

const AppSidebar: React.FC = () => {
	const [showSidebar, setShowSidebar] = useRecoilState(showAppSidebarState)

	return (
		<div
			className={`fixed top-[60px] left-0 bottom-0 z-[900] w-full ${
				showSidebar ? "" : "-translate-x-full"
			}`}
		>
			<div className="dui--drawer">
				<input
					id="app-sidebar"
					type="checkbox"
					checked={showSidebar}
					onChange={() => setShowSidebar((p) => !p)}
					className="dui--drawer-toggle"
				/>

				<div className="dui--drawer-side">
					<label
						htmlFor="app-sidebar"
						className="dui--drawer-overlay"
					></label>
					<ul className="dui--menu p-4 w-80 h-full bg-base-200 text-base-content">
						{/* Sidebar content here */}
						<li>
							<a>Sidebar Item 1</a>
						</li>
						<li>
							<a>Sidebar Item 2</a>
						</li>
					</ul>
				</div>
			</div>
		</div>
	)
}

export default AppSidebar

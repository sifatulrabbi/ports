import type {NextPage} from "next"
import React from "react"

const Login: NextPage = () => {
	const handleSubmit = async (e: React.SyntheticEvent<HTMLFormElement>) => {
		e.preventDefault()
	}

	return (
		<div className="w-full min-h-screen grid grid-cols-1 md:grid-cols-[20%_75%]">
			<div className="drawer w-full h-full">
				<aside className="drawer-side p-4 flex flex-col items-center bg-primary/5">
					<article className="prose">
						<h1 className="font-black mb-2">Ports</h1>
						<p className="leading-[1.4] max-w-[150px] mt-0">Your private drive</p>
					</article>
				</aside>
			</div>
			<div className="flex items-center justify-start">
				<form onSubmit={handleSubmit} className="w-full flex flex-col gap-4 justify-start items-start p-6">
					<article className="prose">
						<h3>Login</h3>
					</article>
					<div className="form-control w-full max-w-lg">
						<label className="label">
							<span className="label-text">Username</span>
						</label>
						<input type="text" placeholder="Type here" className="input input-bordered w-full max-w-lg" />
					</div>
					<div className="form-control w-full max-w-lg">
						<label className="label">
							<span className="label-text">Password</span>
						</label>
						<input
							type="password"
							placeholder="Type here"
							className="input input-bordered w-full max-w-lg"
						/>
					</div>
					<button type="submit" className="btn btn-primary">
						Login
					</button>
				</form>
			</div>
		</div>
	)
}

export default Login

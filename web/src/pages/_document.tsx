import Document, { Html, Head, Main, NextScript } from "next/document"

class MyDocument extends Document {
	render() {
		return (
			<Html lang="en" data-theme="winter">
				<Head>
					<meta charSet="UTF-8" />
					<meta
						name="description"
						content="Ports, connect with your team mates"
					/>
					<meta name="robots" content="index, follow" />
					<link rel="icon" href="/favicon.png" />
				</Head>
				<body>
					<Main />
					<NextScript />
				</body>
			</Html>
		)
	}
}

export default MyDocument

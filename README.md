# Port

### Authentication and authorization

Use an OAuth-like authentication flow.

-   Login will require the user to enter their username and password
-   The server sends back a refresh token, and an access token
-   Use the access token to send requests. Save the refresh token to exchange for access tokens in the future

### The registration process

**Schema**

```go
type User struct {
	// The username has to be unique
	Username  string `json:"password"`
	// Password is the user's password, min-length 8 and max-length 30
	Password  string `json:"password"`
	CreatedAt int    `json:"createdAt"`
}
```

-   Send a `POST` request to the server with the `username` and `password`
-   Server will
    -   Verify the username is unique
    -   Verify the password is the valid (min-length 8 and max-length 30)
    -   Save the information on MongoDB
    -   Send back a request successful message

### The login process

**Auth response schema**

```go
type AuthResponse struct {
	RefreshToken string `json:"refreshToken"`
	AccessToken  string `json:"accessToken"`
}
```

The response will also set a cookie on the client’s end `PUN = {encoded username}`. The cookie is important because it will verify if the user is using any CLI or real browser/HTTP client.

As the login process is complete the server will store the generated refresh token on MongoDB.

**Auth session’s document schema**

```go
type AuthSession struct {
	// Session initiation date
	Iat          int    `bson:"iat"`
	// Session's refresh token
	RefreshToken string `bosn:"refreshToken"`
	// The document will also store the IP address to ensure security
	IP           string `bson:"ip"`
	Username     string `bson:"username"`
}
```

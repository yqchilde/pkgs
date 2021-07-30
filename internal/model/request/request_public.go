package request

type SignUp struct {
	Email       string `json:"email"`
	Username    string `json:"username"`
	Password    string `json:"password"`
	AuthorityID string `json:"authority_id"`
}

type SignIn struct {
	Email    string `json:"email"`
	Username string `json:"username"`
	Password string `json:"password"`
}

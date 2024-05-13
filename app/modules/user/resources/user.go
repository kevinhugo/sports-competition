package resources

type UserLogin struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type UserLoginResponse struct {
	AccessToken string `json:"access_token"`
}

type GetUserIdentityResponse struct {
	Username    string `json:"username"`
	FirstName   string `json:"first_name"`
	Proficiency uint   `json:"proficiency"`
}

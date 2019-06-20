package usermodels

type UserCredentials struct {
	Username string
	Password string
	Email    string
}

type UserSignInRequest struct {
	Username string
	Password string
}

type UserKeys struct {
	Username    string
	AccessToken string
	RenewToken  string
}
type User struct {
	UserID   string
	Password []byte
	Username string
	Email    string
}

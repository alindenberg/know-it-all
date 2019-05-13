package usermodels

type UserRequest struct {
	Username string
	Password string
	Email	string
}

type UserSignInRequest struct {
	Username string
	Password string
}

type User struct {
	UserID string
	Password []byte
	Username string
	Email	string
}

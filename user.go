package GoProject4

type User struct {
	Id       int    `json:"-"`
	Email    string `json:"email"`
	Username string `json:"username"`
	Password string `json:"password"`
}

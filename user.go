package GoProject4

type User struct {
	Id       int    `json:"-" db:"id"`
	Email    string `json:"email" db:"email" binding:"required"`
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

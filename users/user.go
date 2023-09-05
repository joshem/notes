package users

type User struct {
	ID       string `json:"id" gorm:"primary_key"`
	Username string `json:"username"`
	Password string `json:"password"`
}

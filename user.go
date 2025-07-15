package butterplanner

type User struct {
	Id        int    `json:"-"`
	Name      string `json:"name" binding:"required"`
	Last_name string `json:"last_name" binding:"required"`
	Username  string `json:"username" binding:"required"`
	Password  string `json:"password" binding:"required"`
}

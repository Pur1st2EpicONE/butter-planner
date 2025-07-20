package butterplanner

type User struct {
	Id        int    `form:"-" json:"-"`
	Name      string `form:"name" json:"name" binding:"required"`
	Last_name string `form:"last_name" json:"last_name" binding:"required"`
	Username  string `form:"username" json:"username" binding:"required"`
	Password  string `form:"password" json:"password" binding:"required"`
}

type LoginPassword struct {
	Username string `form:"username" json:"username" binding:"required"`
	Password string `form:"password" json:"password" binding:"required"`
}

type Note struct {
	Id      int    `json:"-"`
	UserId  int    `json:"-"`
	Title   string `json:"title"`
	Content string `json:"content"`
}

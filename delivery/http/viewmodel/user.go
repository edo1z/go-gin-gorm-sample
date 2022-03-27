package viewmodel

type UserName struct {
	Name string `form:"name"`
}

type UserID struct {
	ID int `uri:"id" binding:"required"`
}

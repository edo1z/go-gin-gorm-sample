package viewmodel

type CategoryName struct {
	Name string `form:"name"`
}

type CategoryID struct {
	ID int `uri:"id" binding:"required"`
}

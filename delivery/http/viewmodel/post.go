package viewmodel

type PostTitle struct {
	Title string `form:"title"`
}

type PostID struct {
	ID int `uri:"id" binding:"required"`
}

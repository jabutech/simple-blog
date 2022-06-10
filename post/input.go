package post

type CreatePostInput struct {
	Title  string `json:"title" binding:"required"`
	UserId string `uri:"user_id" binding:"required"`
}

type UpdatePostInput struct {
	Title  string `json:"title" binding:"required"`
	PostId string `uri:"post_id" binding:"required"`
}

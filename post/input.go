package post

type CreatePostInput struct {
	Title  string `json:"title" binding:"required"`
	UserId string `uri:"user_id"`
}

type UpdatePostInput struct {
	Title  string `json:"title" binding:"required"`
	PostId string `uri:"post_id"`
	UserId string `uri:"user_id"`
}

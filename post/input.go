package post

type CreatePostInput struct {
	Title  string `json:"title" binding:"required"`
	UserId string `json:"user_id"`
}

package post

type CreateOrUpdatePostInput struct {
	Title  string `json:"title" binding:"required"`
	UserId string `json:"user_id" binding:"required"`
}

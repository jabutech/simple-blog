package user

// GetIdUserInput for get id from uri
type GetIdUserInput struct {
	Id string `uri:"id" binding:"required"`
}

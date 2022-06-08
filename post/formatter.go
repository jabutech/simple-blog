package post

// Formatter for post create or update
type PostCreateOrUpdateFormatter struct {
	Id    string `json:"id"`
	Title string `json:"title"`
}

func FormatPostCreateOrUpdate(post Post) PostCreateOrUpdateFormatter {
	formatter := PostCreateOrUpdateFormatter{}
	formatter.Id = post.Id
	formatter.Title = post.Title

	return formatter
}

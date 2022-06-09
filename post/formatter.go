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

// Structure post formatter
type PostFormatter struct {
	Id     string `json:"id"`
	Author string `json:"author"`
	Title  string `json:"title"`
}

// Formatter for single post
func FormatPost(post Post) PostFormatter {
	formatter := PostFormatter{}
	formatter.Id = post.Id
	formatter.Author = post.User.Fullname
	formatter.Title = post.Title
	return formatter
}

// Formater for multiple posts
func FormatPosts(posts []Post) []PostFormatter {
	// If data not available, return empty array
	if len(posts) == 0 {
		return []PostFormatter{}
	}

	var PostFormatters []PostFormatter
	// Do loop posts, and append to var `PostFormatters` with use PostFormatter
	for _, user := range posts {
		formatter := FormatPost(user)
		PostFormatters = append(PostFormatters, formatter)
	}

	// Return all posts
	return PostFormatters
}

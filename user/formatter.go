package user

type UserFormatter struct {
	Id       string `json:"id"`
	Fullname string `json:"fullname"`
	Email    string `json:"email"`
	IsAdmin  bool   `json:"is_admin"`
}

// FormatUser for format sigle data user
func FormatUser(user User) UserFormatter {
	formatter := UserFormatter{}
	formatter.Id = user.ID
	formatter.Fullname = user.Fullname
	formatter.Email = user.Email

	isAdmin := false
	if user.IsAdmin == 1 {
		isAdmin = true
	}

	formatter.IsAdmin = isAdmin

	return formatter
}

// Formatter for list users
func FormatUsers(users []User) []UserFormatter {
	// If data not available, return empty array
	if len(users) == 0 {
		return []UserFormatter{}
	}

	var userFormatters []UserFormatter
	// Do loop user, and append to var `userFormatters` with use UserFormatter
	for _, user := range users {
		formatter := FormatUser(user)
		userFormatters = append(userFormatters, formatter)
	}

	// Return all users
	return userFormatters
}

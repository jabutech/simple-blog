package user

type Service interface {
	GetUsers(fullname string, email string) ([]User, error)
	GetUserById(id string) (User, error)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{repository}
}

// GetUsers find all users data
func (s *service) GetUsers(fullname string, email string) ([]User, error) {

	// If parameter fullname not empty string
	if fullname != "" {
		// Find user by email
		users, err := s.repository.FindByFullname(fullname)
		if err != nil {
			return users, err
		}

		// If success, return
		return users, nil
	}

	// If parameter email not empty string
	if email != "" {
		users := []User{}
		// Find user by email
		user, err := s.repository.FindByEmail(email)
		if user.ID != "" {
			users = append(users, user)
		}
		if err != nil {
			return users, err
		}

		// If success, return
		return users, nil
	}

	// Find all user with repository
	users, err := s.repository.FindAll()
	if err != nil {
		return users, err
	}

	// If success, return users
	return users, nil
}

func (s *service) GetUserById(id string) (User, error) {
	user, err := s.repository.FindById(id)
	if err != nil {
		return user, err
	}

	return user, nil
}

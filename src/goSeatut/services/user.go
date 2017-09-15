package services

var users map[int]*User

func init() {
	users = make(map[int]*User)
	user := &User{
		ID: 1,
		FirstName: "Big",
		LastName: "Boss",
		Roles: []string{AdministratorRole},
	}
	user2 := &User{
		ID: 2,
		FirstName: "Test",
		LastName: "User",
		Roles: []string{},
	}

	users[user.ID] = user
	users[user2.ID] = user2
}

type User struct {
	ID int
	FirstName string
	LastName string
	Roles []string
}

func (u *User) HasRole(roleName string) bool {
	for _, role := range u.Roles {
		if role == roleName {
			return true
		}
	}
	return false
}

// UserService provides a CRUD interface for Users
type UserService interface {
	Create(*User) error 
	Read(int) (*User, error)
	Update(*User) error
	Delete (int) error
}

// NewUserService() 
func NewUserService() UserService {
	return &userService{}
}

type userService struct {}

func (s *userService) Create(u *User) error {
	users[u.ID] = u
	return nil
}

func (s *userService) Read(id int) (*User, error) {
	return users[id], nil
}

func (s *userService) Update(u *User) error {
	users[u.ID] = u
	return nil
}

func (s *userService) Delete(id int) error {
	delete(users, id)
	return nil
}
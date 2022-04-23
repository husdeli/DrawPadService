package user

type User struct {
	Id        string `json:"id"`
	FirstName string `json:"firstName"`
}

type UserService interface {
	GetUserById(id string) (User, error)
	GetUsers() ([]User, error)
	CreateUser() (User, error)
	UpdateUser(id string) (User, error)
	PatchUser(id string) (User, error)
	DeleteUser(id string) error
}

type UsersRepository interface {
	GetUserById(id string) (User, error)
	GetUsers() ([]User, error)
	CreateUser() (User, error)
	UpdateUser(id string) (User, error)
	PatchUser(id string) (User, error)
	DeleteUser(id string) error
}

type userService struct {
	repository UsersRepository
}

func NewService(repository UsersRepository) UserService {
	return &userService{repository}
}

func (service *userService) GetUserById(id string) (User, error) {
	return service.repository.GetUserById(id)
}

func (service *userService) GetUsers() ([]User, error) {
	return service.repository.GetUsers()
}

func (service *userService) CreateUser() (User, error) {
	return service.repository.CreateUser()
}

func (service *userService) UpdateUser(id string) (User, error) {
	return service.repository.UpdateUser(id)
}

func (service *userService) PatchUser(id string) (User, error) {
	return service.repository.PatchUser(id)
}

func (service *userService) DeleteUser(id string) error {
	return service.repository.DeleteUser(id)
}

package user

// Interface implementation check
var _ UsersRepository = NewRepository()

// Implementation
type usersRepository struct{}

func NewRepository() *usersRepository {
	return &usersRepository{}
}

func (repository *usersRepository) GetUsers() ([]User, error) {
	return []User{{"1", "Dima"}, {"2", "Pasha"}, {"3", "lex"}}, nil
}

func (repository *usersRepository) GetUserById(id string) (User, error) {
	return User{"1", "Dima"}, nil
}

func (repository *usersRepository) CreateUser() (User, error) {
	return User{"1", "Dima"}, nil
}

func (repository *usersRepository) UpdateUser(id string) (User, error) {
	return User{"1", "Dima"}, nil
}

func (repository *usersRepository) PatchUser(id string) (User, error) {
	return User{"1", "Dima"}, nil
}

func (repository *usersRepository) DeleteUser(id string) error {
	return nil
}

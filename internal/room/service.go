package room

type RoomsService interface {
	CreateRoom() (model RoomModel, err error)
	GetRoom(id string) (model RoomModel, err error)
}

type RoomsRepository interface {
	CreateRoom() (model RoomModel, err error)
	GetRoom(id string) (model RoomModel, err error)
}

type roomService struct {
	repository RoomsRepository
}

func NewService(repository RoomsRepository) RoomsService {
	return &roomService{repository}
}

func (service roomService) CreateRoom() (RoomModel, error) {
	return service.repository.CreateRoom()
}

func (service roomService) GetRoom(id string) (RoomModel, error) {
	return service.repository.GetRoom(id)
}

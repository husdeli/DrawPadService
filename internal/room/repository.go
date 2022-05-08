package room

// Interface implementation check
var _ RoomsRepository = NewRepository()

// Implementation
type roomsRepository struct{}

func NewRepository() *roomsRepository {
	return &roomsRepository{}
}

func (repository *roomsRepository) CreateRoom() (RoomModel, error) {
	return RoomModel{Id: "testId"}, nil
}

func (repository *roomsRepository) GetRoom(id string) (RoomModel, error) {
	return RoomModel{Id: "testId"}, nil
}

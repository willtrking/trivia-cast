package room


type RoomManager interface {
    CreateRoom() (*Room, error)
    JoinRoom(roomID, playerName string) (*Player, *Room, error)
}

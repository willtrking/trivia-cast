package local

import (
    "errors"
    "github.com/satori/go.uuid"
    "github.com/willtrking/trivia-cast/go/server/room"
    "sync"
)

func NewRoomManager() room.RoomManager {
    return &RoomManager{}
}

type RoomManager struct {
    watchersLock sync.RWMutex
    roomLock sync.Mutex
    // room ID : Room
    rooms map[string]*room.Room

    // room ID : watcher number :  Watchers
    watchers map[string]map[int]*room.ChangeWatcher
    watcherNum int
}

func (r *RoomManager) generateRoom() *room.Room {

    u := uuid.NewV4()

    return &room.Room{
        ID: u.String(),
    }

}

func (r *RoomManager) generatePlayer() *room.Player {

    u := uuid.NewV4()

    return &room.Player{
        ID: u.String(),
    }

}


func (r *RoomManager) CreateRoom() (*room.Room, error) {

    r.roomLock.Lock()
    defer r.roomLock.Unlock()

    g := r.generateRoom()

    r.rooms[g.ID] = g
    r.watchers[g.ID] = make(map[int]*room.ChangeWatcher)

    return g, nil


}

func (r *RoomManager) JoinRoom(roomID, playerName string) (*room.Player, *room.Room, error) {

    r.roomLock.Lock()
    defer r.roomLock.Unlock()

    room, hasRoom := r.rooms[roomID]
    if hasRoom == false {

        return nil, nil, errors.New("Unknown room ID")

    }

    player := r.generatePlayer()
    player.Name = playerName



    err := room.AddPlayer(player)
    if err != nil {
        return nil, nil, err
    }

    return player, room, nil


}

func (r *RoomManager) WatchRoom(roomID string) (*room.ChangeWatcher, error) {

    r.roomLock.Lock()
    defer r.roomLock.Unlock()

    if _, hasMap := r.watchers[roomID]; hasMap == false {
        return nil, errors.New("Unknown room ID")
    }

    r.watcherNum++

    r.watchers[roomID][r.watcherNum] = &room.ChangeWatcher{
        RoomID:roomID,
    }
    r.watchers[roomID][r.watcherNum].SetNumber(r.watcherNum)
    r.watchers[roomID][r.watcherNum].Start()

    return r.watchers[roomID][r.watcherNum], nil

}


func (r *RoomManager) StopWatching(c *room.ChangeWatcher) {

    r.roomLock.Lock()
    defer r.roomLock.Unlock()


    if _, hasWatchers := r.watchers[c.RoomID]; hasWatchers {
        num := c.Number()
        delete(r.watchers[c.RoomID], num)
    }
}






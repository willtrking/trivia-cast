package server

import (
    "github.com/willtrking/trivia-cast/go/protobuf"
    "github.com/willtrking/trivia-cast/go/server/room"
)

type GRPCServer struct {
    RoomManager room.RoomManager
}

func (g *GRPCServer) CreateRoom(req *protobuf.CreateRoomReq, stream protobuf.BaseGame_CreateRoomServer) error {
    room, err := g.RoomManager.CreateRoom()
    if err != nil {
        return err
    }

    room.AddPlayer()
}

func (g *GRPCServer) JoinRoom(player *protobuf.Player, stream protobuf.BaseGame_JoinRoomServer) error {

}

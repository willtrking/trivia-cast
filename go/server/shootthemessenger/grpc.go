package shootthemessenger

import (
    "context"
    "github.com/willtrking/trivia-cast/go/protobuf/stm"
)

type GRPCServer struct {

}

func (g *GRPCServer) BeginGame(ctx context.Context, req *stm.BeginGameReq) (*stm.GameState, error) {

}

func (g *GRPCServer)PlayerStream(stream stm.ShootTheMessenger_PlayerStreamServer) error {

}



package room


type Round struct {
    Number uint64
    Winner *Player
    Players []*Player
}

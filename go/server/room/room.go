package room

import "errors"

type Room struct {
    ID string
    // Map of players by ID
    Players map[string]*Player
    CurrentRound *Round
    PreviousRounds []*Round
}

func (r *Room) AddPlayer(p *Player) error {

    if _, hasPlayer := r.Players[p.ID]; hasPlayer == false {
        return errors.New("Duplicate player ID")
    }

    if r.NameExists(p.Name) {
        return errors.New("Name already exists")
    }

    r.Players[p.ID] = p
}

// Check if a player exists with a certain name
func (r *Room) NameExists(name string) bool {

    for _, p := range r.Players {
        if p.Name == name {
            return true
        }
    }

    return false
}


//type State

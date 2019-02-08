//
//  Player.swift
//  TriviaCast
//
//  Created by Brophy on 2/8/19.
//  Copyright © 2019 TriviaCast, Inc. All rights reserved.
//

import Foundation

enum PlayerState {
    case unknown
    case joined
    case playing
    case disconnected
}

class Player {
    let id: String
    let name: String
    var score: Int
    var state: PlayerState
    
    init(id: String, name: String, score: Int, state: PlayerState) {
        self.id = id
        self.name = name
        self.score = score
        self.state = state
    }
}

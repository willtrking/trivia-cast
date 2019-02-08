//
//  GameType.swift
//  TriviaCast
//
//  Created by Brophy on 2/6/19.
//  Copyright © 2019 TriviaCast, Inc. All rights reserved.
//

import Foundation

class GameType {
    let name: String
    let rules: [String]
    
    init(name: String, rules: [String]) {
        self.name = name
        self.rules = rules
    }
}

class ShootTheMessenger: GameType {
    init() {
        super.init(
            name: "Shoot The Messenger",
            rules: [
                "Start the Game",
                "DL the app / website",
                "Join the game",
                "Play the game"
            ]
        )
    }
}

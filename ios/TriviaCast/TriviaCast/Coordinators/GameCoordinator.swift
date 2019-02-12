//
//  GameCoordinator.swift
//  TriviaCast
//
//  Created by Brophy on 2/6/19.
//  Copyright © 2019 TriviaCast, Inc. All rights reserved.
//

import Foundation
import UIKit

class GameCoordinator: NSObject, RootCoordinating, ViewControllerListener {
    
    var childCoordinators: [Coordinating] = []
    var rootViewController: UIViewController
    
    private var gameType: GameType = ShootTheMessenger()
    
    private let navigationController: UINavigationController
    
    init(navigationController: UINavigationController) {
        self.navigationController = navigationController
        rootViewController = navigationController
        super.init()
    }
    
    func start() {
        routeToGameSplashScreen()
    }
    
    func routeToGameSplashScreen() {
        let splashScreen = TitleMediaActionsViewController()
        
        splashScreen.setTitle(gameType.name)
        
        let beginButtonAction = ButtonViewModel(
            title: "Begin",
            action: { [weak self] in
                self?.routeToJoinGameLobby()
            }
        )
        let howToPlayButtonAction = ButtonViewModel(
            title: "How to Play",
            action: { [weak self] in
                self?.routeToGameRules()
            }
        )
        splashScreen.setActionButtons([beginButtonAction, howToPlayButtonAction])
        
        
        if let url = Bundle.main.url(forResource: "PlaceholderPreview", withExtension: ".mp4") {
            let mediaView = MediaView()
            mediaView.setContent(url)
            splashScreen.setMediaView(mediaView)
            mediaView.play()
        } else {
            splashScreen.setMediaView(nil)
        }
        
        navigationController.pushViewController(splashScreen, animated: true)
    }
    
    func routeToGameRules() {
        let gameRules = TitleDetailViewController()
        gameRules.setTitle("Game Rules")
        gameRules.setBulletPoints(gameType.rules)
        gameRules.listener = self
        navigationController.pushViewController(gameRules, animated: true)
    }
    
    func dismissGameRules() {
        popCurrentViewController()
    }
    
    func routeToJoinGameLobby() {
        let joinScreen = DoubleMediaBottomTitleViewController()
        
        joinScreen.setTitle("Join now @ bears.nz/bears")
        
        let userScreen = CircularIconsView()
        
        let testViews = (0..<1).map { _ -> UIView in
            let view = UserIconView()
            view.backgroundColor = UIColor.red
            return view
        }
        
        userScreen.setIcons(testViews)
        
        joinScreen.setRightMediaView(userScreen)
        
        navigationController.pushViewController(joinScreen, animated: true)
    }
    
    func dismissJoinGameLobby() {
        
    }
    
    // MARK: - ViewControllerListener
    
    func didTapMenuButton() {
        popCurrentViewController()
    }
    
    // MARK: - Private Interface
    
    private func popCurrentViewController() {
        navigationController.popViewController(animated: true)
    }
}

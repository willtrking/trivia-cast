//
//  AppCoordinator.swift
//  TriviaCast
//
//  Created by Brophy on 2/5/19.
//  Copyright © 2019 TriviaCast, Inc. All rights reserved.
//

import Foundation
import UIKit

protocol AppCoordinating: RootCoordinating {
    
}

class AppCoordinator: NSObject, AppCoordinating {
    
    var childCoordinators: [Coordinating] = []
    
    var rootViewController: UIViewController {
        return self.navigationController
    }

    private let window: UIWindow
    
    private lazy var navigationController: UINavigationController = {
        let navigationController = UINavigationController()
        navigationController.isNavigationBarHidden = true
        return navigationController
    }()
    
    init(window: UIWindow) {
        self.window = window
        super.init()
        
        window.rootViewController = rootViewController
        window.makeKeyAndVisible()
    }
    
    func start() {
        let viewController = TitleDetailActionViewController()
        viewController.setTitle("TriviaCast!")
        viewController.setDetail("Copyright 2019 by some drunk guys")
        viewController.setPrimaryAction("Start") { [weak self] in
            self?.routeToGameSelect()
        }
        navigationController.setViewControllers([viewController], animated: true)
    }
    
    // MARK: - Actions
    
    func routeToGameSelect() {
        let gameCoordinator = GameCoordinator(navigationController: navigationController)
        self.childCoordinators.append(gameCoordinator)
        gameCoordinator.start()
    }
}

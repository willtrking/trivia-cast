//
//  BaseCoordinator.swift
//  TriviaCast
//
//  Created by Brophy on 2/5/19.
//  Copyright © 2019 TriviaCast, Inc. All rights reserved.
//

import Foundation
import UIKit

protocol Coordinating {
    var childCoordinators: [Coordinating] { get set }
    
    func start()
}

protocol RootViewControllerProviding: class {
    var rootViewController: UIViewController { get }
}

typealias RootCoordinating = Coordinating & RootViewControllerProviding

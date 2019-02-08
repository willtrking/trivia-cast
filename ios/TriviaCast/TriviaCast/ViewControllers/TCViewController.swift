//
//  TCViewController.swift
//  TriviaCast
//
//  Created by Brophy on 2/7/19.
//  Copyright © 2019 TriviaCast, Inc. All rights reserved.
//

import Foundation
import UIKit

protocol ViewControllerListener: class {
    func didTapMenuButton()
}

class TCViewController: UIViewController {
    
    weak var listener: ViewControllerListener?
    
    override func viewDidLoad() {
        super.viewDidLoad()
        setupTapGestures()
    }
    
    func setupTapGestures() {
        let backGesture = UITapGestureRecognizer(target: self, action: #selector(handleMenuTap(_:)))
        backGesture.allowedPressTypes = [NSNumber(value: UIPress.PressType.menu.rawValue)]
        self.view.addGestureRecognizer(backGesture)
    }
    
    @objc func handleMenuTap(_ gestureRecognizer: UITapGestureRecognizer) {
        listener?.didTapMenuButton()
    }
}

//
//  CircularIconsView.swift
//  TriviaCast
//
//  Created by Brophy on 2/8/19.
//  Copyright © 2019 TriviaCast, Inc. All rights reserved.
//

import Foundation
import UIKit

class CircularIconsView: UIView {
    
    private var icons: [UIView] = []
    
    init() {
        super.init(frame: CGRect.zero)
        build()
        configure()
    }
    
    required init?(coder aDecoder: NSCoder) {
        fatalError("init(coder:) has not been implemented")
    }
    
    func setIcons(_ icons: [UIView]) {
        for icon in icons {
            icon.removeFromSuperview()
        }
        self.icons = icons
        
        guard icons.count > 0 else {
            return
        }
        guard icons.count > 1 else {
            addSubview(icons[0])
            icons[0].snp.makeConstraints { make in
                make.center.equalTo(self)
                make.height.width.equalTo(100.0)
            }
            return
        }
        
        let angleIncrement = (2.0 * Double.pi) / Double(icons.count)
        var initialAngle = 0.0
        
        for icon in icons {
            let xMultiplier = 1.0 - cos(initialAngle) * 0.8
            let yMultiplier = 1.0 - sin(initialAngle) * 0.8
            
            addSubview(icon)
            icon.snp.makeConstraints { make in
                if xMultiplier > 0 {
                    make.centerX.equalTo(self).multipliedBy(xMultiplier)
                } else {
                    make.centerX.equalTo(0)
                }
                if yMultiplier > 0 {
                    make.centerY.equalTo(self).multipliedBy(yMultiplier)
                } else {
                    make.centerY.equalTo(0)
                }
                make.height.width.equalTo(100.0)
            }
            initialAngle += angleIncrement
        }
    }
    
    // MARK: - Private Interface
    
    private func build() {

    }
    
    private func configure() {
        
    }
}

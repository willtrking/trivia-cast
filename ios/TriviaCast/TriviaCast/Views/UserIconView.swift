//
//  UserIconView.swift
//  TriviaCast
//
//  Created by Brophy on 2/8/19.
//  Copyright © 2019 TriviaCast, Inc. All rights reserved.
//

import Foundation
import UIKit
import SnapKit

enum UserStatus {
    case notReady
    case ready
    
    var color: UIColor {
        switch self {
        case .notReady:
            return UIColor.red
        case .ready:
            return UIColor.green
        }
    }
    
    var image: UIImage {
        switch self {
        case .notReady:
            return UIImage(imageLiteralResourceName: "status_icon_ready")
                .withRenderingMode(.alwaysTemplate)
        case .ready:
            return UIImage(imageLiteralResourceName: "status_icon_ready")
                .withRenderingMode(.alwaysTemplate)
        }
    }
}

class UserIconView: UIView {
    
    private let playerNameLabel: UILabel = UILabel()
    private let profileImageView: UIImageView = UIImageView()
    private let statusIconImageView: UIImageView = UIImageView()
    
    init() {
        super.init(frame: CGRect.zero)
        build()
        configure()
    }
    
    required init?(coder aDecoder: NSCoder) {
        fatalError("init(coder:) has not been implemented")
    }
    
    override func layoutSubviews() {
        super.layoutSubviews()
        profileImageView.layer.cornerRadius = profileImageView.frame.size.height / 2.0
        statusIconImageView.layer.cornerRadius = statusIconImageView.frame.size.height / 2.0
    }
    
    func setStatus(_ status: UserStatus?) {
        guard let status = status else {
            statusIconImageView.isHidden = true
            return
        }
        statusIconImageView.tintColor = status.color
        statusIconImageView.image = status.image
    }
    
    func setPlayer(_ player: Player) {
        // TODO: Set image
        playerNameLabel.text = player.name
    }
    
    // MARK: - Private Interface
    
    private func build() {
        addSubview(profileImageView)
        addSubview(statusIconImageView)
        addSubview(playerNameLabel)
        
        profileImageView.clipsToBounds = true
        
        statusIconImageView.clipsToBounds = true
        
        playerNameLabel.font = TCFont.body
        playerNameLabel.textAlignment = .center
        playerNameLabel.numberOfLines = 0
    }
    
    private func configure() {
        
        playerNameLabel.snp.makeConstraints { make in
            make.bottom.centerX.equalTo(self)
        }
        
        profileImageView.snp.makeConstraints { make in
            make.bottom.equalTo(playerNameLabel.snp.top)
            make.top.equalTo(self)
            make.width.equalTo(profileImageView.snp.height)
        }

        statusIconImageView.snp.makeConstraints { make in
            make.leading.bottom.equalTo(profileImageView)
            make.width.equalTo(profileImageView).multipliedBy(0.2)
            make.height.equalTo(statusIconImageView.snp.width)
        }
    }
}

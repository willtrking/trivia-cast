//
//  TitleDetailActionViewController.swift
//  TriviaCast
//
//  Created by Brophy on 2/5/19.
//  Copyright © 2019 TriviaCast, Inc. All rights reserved.
//

import Foundation
import UIKit
import TVUIKit

import SnapKit

class TitleDetailActionViewController: TCViewController {
    
    private let titleLabel: UILabel = UILabel()
    private let detailLabel: UILabel = UILabel()
    private let primaryActionButton: UIButton = UIButton(type: UIButton.ButtonType.roundedRect)
    
    private var primaryAction: ButtonAction?
    
    init() {
        super.init(nibName: nil, bundle: nil)
        build()
        configure()
        setupGestures()
    }
    
    required init?(coder aDecoder: NSCoder) {
        fatalError("init(coder:) has not been implemented")
    }
    
    func setTitle(_ title: String?) {
        titleLabel.text = title
    }
    
    func setDetail(_ detail: String?) {
        detailLabel.text = detail
    }
    
    func setPrimaryAction(_ title: String?, action: ButtonAction?) {
        primaryActionButton.setTitle(title, for: .normal)
        primaryAction = action
    }
    
    // MARK: - Private Interface
    
    private func build() {
        view.addSubview(titleLabel)
        view.addSubview(detailLabel)
        view.addSubview(primaryActionButton)
        
        detailLabel.numberOfLines = 0
        detailLabel.textAlignment = .center
        detailLabel.font = TCFont.body
        
        titleLabel.numberOfLines = 0
        titleLabel.textAlignment = .center
        print(titleLabel.font.pointSize)
        titleLabel.font = TCFont.h1
        
        primaryActionButton.isUserInteractionEnabled = true
        primaryActionButton.clipsToBounds = false
        primaryActionButton.titleLabel?.font = TCFont.h4
    }
    
    private func configure() {
        detailLabel.snp.makeConstraints { make in
            make.leading.trailing.equalTo(view)
            make.bottom.equalTo(view).offset(-SpacingUnit.u2)
        }
        
        primaryActionButton.snp.makeConstraints { make in
            make.centerX.equalTo(view)
            make.bottom.equalTo(detailLabel.snp.top).offset(-SpacingUnit.u4)
        }
        
        titleLabel.snp.makeConstraints { make in
            make.leading.trailing.equalTo(view)
            make.centerY.equalTo(view)
//            make.top.equalTo(view).offset(SpacingUnit.u2)
        }
    }
    
    private func setupGestures() {
        primaryActionButton.addTarget(
            self,
            action: #selector(handlePrimaryAction(_:)),
            for: .primaryActionTriggered
        )
    }
    
    // MARK: - Button Gesture stuff
    
    @objc func handlePrimaryAction(_ button: UIButton) {
        primaryAction?()
    }
}

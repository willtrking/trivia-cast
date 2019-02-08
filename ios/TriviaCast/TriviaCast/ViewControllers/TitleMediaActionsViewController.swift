//
//  TitleMediaActionsViewController.swift
//  TriviaCast
//
//  Created by Brophy on 2/6/19.
//  Copyright © 2019 TriviaCast, Inc. All rights reserved.
//

import Foundation
import UIKit

import SnapKit

class TitleMediaActionsViewController: TCViewController {
    private let titleLabel: UILabel = UILabel()
    private let mediaContainer: UIView = UIView()
    private let buttonStackView: UIStackView = UIStackView()
    
    private var actionButtons: [UIButton] = []
    private var actions: [ButtonAction] = []
    private var mediaView: UIView?
    
    
    override func viewDidLoad() {
        super.viewDidLoad()
        build()
        configure()
    }
    
    func setMediaView(_ view: UIView?) {
        self.mediaView?.removeFromSuperview()
        self.mediaView = view
        
        if let view = view {
            mediaContainer.addSubview(view)
            view.snp.makeConstraints { make in
                make.edges.equalTo(mediaContainer)
            }
        }
    }
    
    func setTitle(_ title: String?) {
        titleLabel.text = title
    }
    
    func setActionButtons(_ actionButtons: [ButtonViewModel]) {
        for button in self.actionButtons {
            button.removeFromSuperview()
        }
        self.actionButtons = []
        self.actions = []
        
        let view = UIView()
        view.setContentHuggingPriority(UILayoutPriority.defaultLow, for: .vertical)
        view.setContentHuggingPriority(UILayoutPriority.defaultLow, for: .horizontal)
        buttonStackView.addArrangedSubview(view)
        for (index, model) in actionButtons.enumerated() {
            let button = UIButton(type: UIButton.ButtonType.roundedRect)
            button.setTitle(model.title, for: .normal)
            button.tag = index
            button.addTarget(self, action: #selector(handleAction(_:)), for: .primaryActionTriggered)
            buttonStackView.addArrangedSubview(button)
            self.actionButtons.append(button)
            actions.append(model.action)
        }
    }
    
    // MARK: - Private Interface
    
    private func build() {
        view.addSubview(titleLabel)
        view.addSubview(mediaContainer)
        view.addSubview(buttonStackView)
        
        titleLabel.font = TCFont.h1
        titleLabel.textAlignment = .center
        
        buttonStackView.axis = .vertical
        buttonStackView.alignment = .center
        buttonStackView.spacing = SpacingUnit.u4
    }
    
    private func configure() {
        titleLabel.snp.makeConstraints { make in
            make.leading.trailing.equalTo(view)
            make.top.equalTo(view).offset(SpacingUnit.u2)
        }
        
        mediaContainer.snp.makeConstraints { make in
            make.top.equalTo(titleLabel.snp.bottom).offset(SpacingUnit.u4)
            make.trailing.bottom.equalTo(view).offset(-SpacingUnit.u4)
            make.leading.equalTo(view.snp.centerX)
        }
        
        buttonStackView.snp.makeConstraints { make in
            make.leading.equalTo(view).offset(SpacingUnit.u4)
            make.top.equalTo(mediaContainer)
            make.trailing.equalTo(view.snp.centerX)
            make.bottom.equalTo(mediaContainer).offset(-SpacingUnit.u2)
        }
    }
    
    @objc func handleAction(_ button: UIButton) {
        guard button.tag < actions.count else {
            return
        }
        actions[button.tag]()
    }
}
